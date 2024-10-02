package library

import (
	"lib_app/storage"
	"lib_app/types"
)

type SimpleLib struct {
	currStorage *storage.Storage
	currIdGen   func(string) int
	idContainer map[string]int
}

func BuildSimpleLib(newStorage storage.Storage, newIdGen func(string) int) SimpleLib {
	return SimpleLib{&newStorage, newIdGen, make(map[string]int)}
}

func (l *SimpleLib) AddBook(newBook types.Book) {
	newId := l.currIdGen(newBook.Title)
	l.idContainer[newBook.Title] = newId
	(*l.currStorage).AddBook(newBook, newId)
}

func (l *SimpleLib) Search(title string) (types.Book, bool) {
	return (*l.currStorage).Search(l.idContainer[title])
}

func (l *SimpleLib) SetStorage(newStorage storage.Storage) {
	(*l.currStorage).Migrate(&newStorage)
	l.currStorage = &newStorage
}

func (l *SimpleLib) SetIdGen(newIdGen func(string) int) {
	l.currIdGen = newIdGen
	idRelations := make(map[int]int, len(l.idContainer))
	for key, value := range l.idContainer {
		idRelations[value] = newIdGen(key)
	}
	(*l.currStorage).RebuildId(idRelations)
}
