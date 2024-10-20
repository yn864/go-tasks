package main

import (
	"fmt"
	"lib_app/library"
	"lib_app/storage"
	"lib_app/types"
	"math/rand"
	"time"
)

func RandomIdGen(title string) int {
	return rand.Int()
}

func TimeIdGen(title string) int {
	return int(time.Now().Unix())
}

func main() {
	book := types.Book{Title: "Don Quixote", Author: "Miguel de Cervantes"}
	anotherBook := types.Book{Title: "The Lord of the Rings", Author: "John Ronald Reuel Tolkien"}

	sliceStorage := storage.NewSliceStorage()
	library1 := library.NewSimpleLib(&sliceStorage, RandomIdGen)

	library1.AddBook(book)
	library1.AddBook(anotherBook)

	book1, err1 := library1.Search("Don Quixote")
	_, err2 := library1.Search("Evgenii Onegin")

	if err1 {
		fmt.Println(book1.Author)
	}

	if !err2 {
		fmt.Println("couldn't find it((")
	}

	library1.SetIdGen(TimeIdGen)

	book3, err3 := library1.Search("The Lord of the Rings")

	if err3 {
		fmt.Println(book3.Title)
	}

	mapStorage := storage.NewMapStorage()
	library2 := library.NewSimpleLib(&mapStorage, RandomIdGen)

	oneMoreBook := types.Book{Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez"}
	library2.AddBook(oneMoreBook)

	book4, err4 := library2.Search("One Hundred Years of Solitude")

	if err4 {
		fmt.Println(book4.Author)
	}
}
