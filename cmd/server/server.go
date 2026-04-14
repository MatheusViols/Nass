package main

import (
	"net/http"
	"database/sql"

	"nass/cmd/handlers"
)

var db *sql.DB

func main() {
	fs := http.FileServer(http.Dir("../../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/edit/", handlers.Edit)
	http.HandleFunc("/save/", handlers.Save)


	http.ListenAndServe(":8000", nil)
}