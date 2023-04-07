package bible

// Book holds the name and chapters of a book
type Book struct {
	Number   int       `json:"number"`
	Name     string    `json:"name"`
	Notes    string    `json:"notes"` // translation notes that might be relevant to store
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

func (b *Book) GetChapterOrCreate(number int) *Chapter {
	chapter := b.GetChapter(number)
	if chapter == nil {
		// Not found? Add it.
		b.Chapters = append(b.Chapters, Chapter{
			Number: number,
		})
		return &b.Chapters[len(b.Chapters)-1]
	}
	return chapter
}

func (b *Book) GetChapter(number int) *Chapter {
	for k := range b.Chapters {
		chapter := &b.Chapters[k]
		if chapter.Number == number {
			return chapter
		}
	}
	return nil
}
