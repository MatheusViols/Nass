package handlers


import (
	"net/http"
	"html/template"

	"nass/cmd/models"
)

func Edit(res http.ResponseWriter, req *http.Request) {
	htmlTemplate, err :=  template.ParseFiles("../../templates/edit.html")
	if err != nil {
		http.Error(res, "Error: Internal server error", 500)
	}

	var noteTitle string = req.URL.Path[len("/edit/"):]

	noteToEdit, err := models.Notes.Search(noteTitle)
	if err != nil {
		newNote := models.NoteDTO{
			Title: noteTitle,
		}
		htmlTemplate.Execute(res, newNote)

		return 
	} 

	htmlTemplate.Execute(res, noteToEdit)		
}