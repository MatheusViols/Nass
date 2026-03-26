package models


type Note struct {
	Title string `json:"title"`	
	Body  string `json:"body"`

	Date  string `json:"date"`
}