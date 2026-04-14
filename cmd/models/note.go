package models

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
}

func (nt *Note) DTO() (NoteDTO) {
	return NoteDTO{
		Title: nt.Title,
		Body: nt.Body,
	}
}