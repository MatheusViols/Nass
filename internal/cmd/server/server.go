package main


import (
	"net/http"

	"nass/internal/cmd/handlers"
)


func main() {

	http.HandleFunc("/", handlers.Show)
	http.HandleFunc("/new", handlers.Create)
//	http.HandleFunc("/att", handlers.Actualize)
//	http.HandleFunc("/del", handlers.Delete)

	http.ListenAndServe(":3000", nil)
}