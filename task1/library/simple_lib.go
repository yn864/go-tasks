package library

import (
	"lib_app/storage"
	"lib_app/types"
)

type SimpleLib struct {
	storage storage.Storage
	currIdGen   func(string) int
	idContainer map[string]int
}

func NewSimpleLib(newStorage storage.Storage, newIdGen func(string) int) SimpleLib {
	return SimpleLib{newStorage, newIdGen, make(map[string]int)}
}

func (l *SimpleLib) AddBook(newBook types.Book) {
	newId := l.currIdGen(newBook.Title)
	l.idContainer[newBook.Title] = newId
	(l.storage).AddBook(newBook, newId)
}

func (l *SimpleLib) Search(title string) (types.Book, bool) {
	return (l.storage).Search(l.idContainer[title])
}

func (l *SimpleLib) SetIdGen(newIdGen func(string) int) {
	l.currIdGen = newIdGen
	idRelations := make(map[int]int, len(l.idContainer))
	for key, value := range l.idContainer {
		idRelations[value] = newIdGen(key)
	}
	(l.storage).RebuildId(idRelations)
}
