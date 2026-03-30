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

	var noteTitle string = req.FormValue("title")
	if noteTitle == "" {
		http.Redirect(res, req, "Notes can't have empty titles", 400)
		return
	}

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