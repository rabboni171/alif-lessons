package main

import (
	"fmt"

	"library/book"
	"library/catalog"
)

func main() {
	books := []book.Book{
		{Title: "Мастер и Маргарита", Author: "Булгаков", Year: 1967},
		{Title: "Собачье сердце", Author: "Булгаков", Year: 1925},
		{Title: "Преступление и наказание", Author: "Достоевский", Year: 1866},
		{Title: "Идиот", Author: "Достоевский", Year: 1869},
		{Title: "Отцы и дети", Author: "Тургенев", Year: 1862},
	}

	found := catalog.FindByAuthor(books, "Достоевский")
	fmt.Println("Книги Достоевского:")
	for _, b := range found {
		fmt.Printf("  %s (%d)\n", b.Title, b.Year)
	}
}
