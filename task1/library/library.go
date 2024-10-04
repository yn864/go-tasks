package library

import (
	"lib_app/types"
)

type Library interface {
	AddBook(newBook types.Book)
	Search(title string) (types.Book, bool)
	SetIdGen(newIdGen func(string) int)
}
