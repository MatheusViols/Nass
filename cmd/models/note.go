package models

type Note struct {
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