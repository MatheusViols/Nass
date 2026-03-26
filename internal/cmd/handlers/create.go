package handlers


import (
	"net/http"
	"encoding/json"

	"nass/internal/cmd/models"
)


func Create(res http.ResponseWriter, req *http.Request) {
	/*
	ERROR: Something wrong with parsing JSON to a Note
	Read more about parsing:

	pkg.go.dev/encoding/json#Encode
	pkg.go.dev/encoding/json#Marshal
	*/

	var newNote models.Note

	err := json.NewDecoder(req.Body).Decode(&newNote)
	if err != nil {
		res.WriteHeader(400)
		json.NewEncoder(res).Encode("Error: Could not create a new note")

		return
	}

	models.Notes.Add(newNote)
	res.WriteHeader(200)
	json.NewEncoder(res).Encode("Sucesso!")
}