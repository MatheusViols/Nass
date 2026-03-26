package models 

type NoteList struct {
	list map[int]Note
}

var nextId int = 1
var Notes = new(NoteList)



func (nts *NoteList) GetAll() map[int]Note {
	return nts.list
}

func (nts *NoteList) Add(newNote Note) {
	nts.list[nextId] = newNote
	nextId++
}

func (nts *NoteList) Del(id int) {
	delete(nts.list, id)
}

