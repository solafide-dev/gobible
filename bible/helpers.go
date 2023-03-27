package bible

import "strings"

// Helper Function that can normalize a book name
// (e.g. "Ezk" -> "Ezekiel)
func NormalizeBookName(bookIn string) string {
	for _, book := range booksTable {
		if strings.EqualFold(bookIn, book.Name) {
			return book.Name
		}
		for _, altName := range book.Alts {
			if strings.EqualFold(bookIn, altName) {
				return book.Name
			}
		}
	}
	return bookIn
}

// Helper Function that can validate a book name
func BookIsValid(bookIn string) bool {
	for _, book := range booksTable {
		if strings.EqualFold(bookIn, book.Name) {
			return true
		}
		for _, altName := range book.Alts {
			if strings.EqualFold(bookIn, altName) {
				return true
			}
		}
	}
	return false
}

// Takes a normalized (see NormalizeBookName) book name and returns the book number
func BookNumberFromName(normalized string) int {
	for _, book := range booksTable {
		if book.Name == normalized {
			return book.BookNumber
		}
	}
	return -1
}
