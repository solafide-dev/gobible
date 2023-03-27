package bible

import "strings"

// Helper Function that can normalize a book name
// (e.g. "Ezk" -> "Ezekiel)
func NormalizeBookName(bookIn string) string {
	for book, _ := range bookEnglishShortNames {
		if strings.EqualFold(bookIn, book) {
			return book
		}
	}
	// No match found, we must traverse the book names and find the closest match
	for book, names := range bookEnglishShortNames {
		for _, name := range names {
			if strings.EqualFold(bookIn, name) {
				return book
			}
		}
	}
	return bookIn
}

// Helper Function that can validate a book name
func BookIsValid(bookIn string) bool {
	for book, _ := range bookEnglishShortNames {
		if strings.EqualFold(bookIn, book) {
			return true
		}
	}
	// No match found, we must traverse the book names and find the closest match
	for _, names := range bookEnglishShortNames {
		for _, name := range names {
			if strings.EqualFold(bookIn, name) {
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
