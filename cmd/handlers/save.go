package handlers


import (
	"net/http"
	"encoding/json"

	"nass/cmd/models"
)

type SavedMessage struct {
	Message string
}


func Save(res http.ResponseWriter, req *http.Request) {

	var sm = SavedMessage{}

	var bodyData models.NoteDTO
	
	err := json.NewDecoder(req.Body).Decode(&bodyData)
	if err != nil {
		sm.Message = "Internal Error"
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(sm)
		return
	}

	var noteTitle string = bodyData.Title
	var noteBody string = bodyData.Body 

	if noteTitle == "" {
		sm.Message = "Notes can't have an empty title"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(sm)
		return
	}

	var newNote = models.NoteDTO{
		Title: noteTitle,
		Body: noteBody,
	}

	
	existentNote, err := models.Notes.Search(noteTitle)
	if err != nil {
		err = models.Notes.Add(newNote)
		if err != nil {
			sm.Message = "Could not save the new note"
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(sm)

			return
		}

		sm.Message = "Successfuly added the new note"

	} else {
		err = models.Notes.Edit(existentNote.Title, newNote)
		if err != nil {
			sm.Message = "Could not edit" + existentNote.Title
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(sm)

			return
		}

		sm.Message = "Successfuly edited the note"
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(sm)

}