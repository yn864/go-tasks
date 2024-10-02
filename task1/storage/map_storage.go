package storage

import (
	"lib_app/types"
)

type MapStorage struct {
	booksContainer map[int]types.Book
}

func CreateMapStorage() MapStorage {
	return MapStorage{make(map[int]types.Book)}
}

func (s *MapStorage) Search(id int) (types.Book, bool) {
	book, status := s.booksContainer[id]
	return book, status
}

func (s *MapStorage) AddBook(book types.Book, id int) {
	s.booksContainer[id] = book
}

func (s *MapStorage) RebuildId(idRelations map[int]int) {
	for oldId, newId := range idRelations {
		s.booksContainer[oldId] = s.booksContainer[newId]
	}
}

func (s *MapStorage) Migrate(newStorage *Storage) {
	for id, book := range s.booksContainer {
		(*newStorage).AddBook(book, id)
	}
	s.booksContainer = nil
}