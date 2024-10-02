package storage

import "lib_app/types"

type Storage interface {
	Search(id int) (types.Book, bool)
	AddBook(book types.Book, id int)
	RebuildId(idRelations map[int]int)
	Migrate(newStorage *Storage)
}
