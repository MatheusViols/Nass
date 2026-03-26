package handlers


import (
	"net/http"
	"encoding/json"

	"nass/internal/cmd/models"
)

func Show(res http.ResponseWriter, req *http.Request) {
	notes := models.Notes.GetAll()

	res.WriteHeader(200)
	json.NewEncoder(res).Encode(notes)
}


