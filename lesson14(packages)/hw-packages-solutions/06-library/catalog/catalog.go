package catalog

import "library/book"

// FindByAuthor возвращает все книги указанного автора.
// Пакет catalog импортирует пакет book - так один свой пакет
// использует другой свой пакет внутри того же модуля.
func FindByAuthor(books []book.Book, author string) []book.Book {
	var result []book.Book
	for _, b := range books {
		if b.Author == author {
			result = append(result, b)
		}
	}
	return result
}
