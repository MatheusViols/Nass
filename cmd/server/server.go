package main

import (
	"net/http"

	"nass/cmd/handlers"
)


func main() {
	/*
	TODO: 
	Definir FileServer.
	Definir handler para o diretório static.
	*/

	fs := http.FileServer(http.Dir("../../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/edit/", handlers.Edit)
	http.HandleFunc("/save/", handlers.Save)


	

	http.ListenAndServe(":8000", nil)
}