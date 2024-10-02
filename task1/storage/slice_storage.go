package storage

import (
	"lib_app/types"
)

type SliceStorageObject struct {
	book types.Book
	id   int
}

type SliceStorage struct {
	booksContainer []SliceStorageObject
}

func CreateSliceStorage() SliceStorage {
	return SliceStorage{[]SliceStorageObject{}}
}

func (s *SliceStorage) Search(id int) (types.Book, bool) {
	for _, object := range s.booksContainer {
		if object.id == id {
			return object.book, true
		}
	}
	return types.Book{}, false
}

func (s *SliceStorage) AddBook(book types.Book, id int) {
	s.booksContainer = append(s.booksContainer, SliceStorageObject{book, id})
}

func (s *SliceStorage) RebuildId(idRelations map[int]int) {
	for _, object := range s.booksContainer {
		object.id = idRelations[object.id]
	}
}

func (s *SliceStorage) Migrate(newStorage *Storage) {
	for _, object := range s.booksContainer {
		(*newStorage).AddBook(object.book, object.id)
	}
	s.booksContainer = nil
}