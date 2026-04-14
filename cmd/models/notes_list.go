package models

import (
	"time"
	"errors"
	"database/sql"
)

type NotesList struct {
	db *sql.DB
}

var Notes = NotesList{
	db: connect(),
}

func (nts *NotesList) Search(searchTitle string) (*Note, error) {
	var note Note = Note{}

	row := nts.db.QueryRow("SELECT * FROM notes WHERE title = ?", searchTitle)
	err :=  row.Scan(&note.ID, &note.Title, &note.Body, &note.Date);

	if err != nil {
		return nil, errors.New("Could not find a note for the title" + searchTitle)
	}

	return &note, nil
}

func (nts *NotesList) Add(newNote NoteDTO) error {
	existentNote, err := nts.Search(newNote.Title)
	if existentNote != nil {
		return errors.New("Trying to add an already existent note")
	}

	actualDate := time.Now().Format(time.DateOnly)

	cmd := "INSERT INTO notes (title, body, edit_date) VALUES (?, ?, ?)"

	_, err = nts.db.Exec(cmd, newNote.Title, newNote.Body, actualDate)
	if err != nil {
		return errors.New("Could not add the new note")
	}

	return nil
}

func (nts *NotesList) Edit(noteTitle string, newNote NoteDTO) error {
	existentNote, err := nts.Search(noteTitle)
	if existentNote == nil {
		return errors.New("Could not find a note for the title" + noteTitle)
	}

	actualDate := time.Now().Format(time.DateOnly)

	cmd := "UPDATE notes SET title = ? , body = ? , edit_date = ? WHERE title = ?"

	_, err = nts.db.Exec(cmd, newNote.Title, newNote.Body, actualDate, noteTitle)
	if err != nil {
		return errors.New("Could not edit the note" + noteTitle)
	}

	return nil
}