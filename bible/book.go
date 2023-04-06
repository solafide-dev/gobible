package bible

// Book holds the name and chapters of a book
type Book struct {
	Number   int       `json:"number"`
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
}

func (b *Book) GetName() string {
	return b.Name
}

// Use the book number to get the English name of the book
func (b *Book) GetEnglishName() string {
	for _, book := range BooksTable {
		if book.BookNumber == b.Number {
			return book.Name
		}
	}
	return ""
}

func (b *Book) GetChapter(number int) *Chapter {
	for _, chapter := range b.Chapters {
		if chapter.Number == number {
			return &chapter
		}
	}
	return nil
}
