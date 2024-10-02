package library

import (
	"lib_app/storage"
	"lib_app/types"
)

type Library interface {
	AddBook(newBook types.Book)
	Search(title string) (types.Book, bool)
	SetStorage(newStorage storage.Storage)
	SetIdGen(newIdGen func(string) int)
}
