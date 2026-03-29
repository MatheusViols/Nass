package models

import (
	"time"
	"errors"
)

type NotesList struct {
	list []Note
}

var Notes = NotesList{
	list: make([]Note, 0),
}

func (nts *NotesList) Search(searchTitle string) (*Note, error) {
	for _, actualNote :=  range nts.list {
		if actualNote.Title == searchTitle {
			return &actualNote, nil
		}
	}

	return nil, errors.New("Could not find a note for the title" + searchTitle)
}

func (nts *NotesList) Add(newNote NoteDTO) error {
	existentNote, _ := Notes.Search(newNote.Title)
	if existentNote != nil {
		return errors.New("Can't create a note if it already exists")
	}

	
	note := Note{
		Title: newNote.Title,
		Body: newNote.Body,
		Date: time.Now().Format(time.DateOnly),
	}

	nts.list = append(nts.list, note)

	return nil
}

func (nts *NotesList) Edit(noteTitle string, newNote NoteDTO) error {
	existentNote, err := Notes.Search(noteTitle)
	if err != nil {
		return errors.New(noteTitle + " does not exist")
	}

	existentNote.Title = newNote.Title
	existentNote.Body = newNote.Body
	existentNote.Date = time.Now().Format(time.DateOnly)

	return nil
}