package main

import (
	"net/http"

	"nass/cmd/handlers"
)


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/edit/", handlers.Edit)
	http.HandleFunc("/save/", handlers.Save)
	

	http.ListenAndServe(":8000", nil)
}