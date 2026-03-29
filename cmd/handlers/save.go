package handlers


import (
	"net/http"
	"encoding/json"

	"nass/cmd/models"
)


func Save(res http.ResponseWriter, req *http.Request) {
	var noteTitle string = req.URL.Path[len("/save/"):]
	var noteBody string = req.FormValue("body")

	if noteTitle == "" {
		http.Error(res, "Note's title can't be empty", 400)
		return
	}

	newNote := models.NoteDTO{
		Title: noteTitle,
		Body: noteBody,
	}

	existentNote, err := models.Notes.Search(noteTitle)
	if err != nil {
		err = models.Notes.Add(newNote)
		if err == nil {
			res.WriteHeader(200)
			json.NewEncoder(res).Encode("Succesfuly created the new note")

			return
		} 
	} else {
		err := models.Notes.Edit(existentNote.Title, newNote)
		if err == nil {
			res.WriteHeader(200)
			json.NewEncoder(res).Encode("Succesfuly edited the existent note")

			return
		}	
	}

}