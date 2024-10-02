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
	book := types.Book{Title: "Don Quixote", Author: "Miguel de Cervantes", PageCount: 1072}
	anotherBook := types.Book{Title: "The Lord of the Rings", Author: "John Ronald Reuel Tolkien", PageCount: 1178}

	sliceStorage := storage.CreateSliceStorage()
	library := library.BuildSimpleLib(&sliceStorage, RandomIdGen)

	library.AddBook(book)
	library.AddBook(anotherBook)

	searchResult1, status1 := library.Search("Don Quixote")
	_, status2 := library.Search("Evgenii Onegin")

	if status1 {
		fmt.Println(searchResult1.Author)
	}

	if !status2 {
		fmt.Println("couldn't find it((")
	}

	library.SetIdGen(TimeIdGen)

	searchResult3, status3 := library.Search("The Lord of the Rings")

	if status3 {
		fmt.Println(searchResult3.PageCount)
	}

	mapStorage := storage.CreateMapStorage()
	library.SetStorage(&mapStorage)

	oneMoreBook := types.Book{Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez", PageCount: 448}
	library.AddBook(oneMoreBook)

	searchResult4, status4 := library.Search("One Hundred Years of Solitude")

	if status4 {
		fmt.Println(searchResult4.Author)
	}
}
