package gobible

// Bible is the main struct that holds all the data
type Bible struct {
	Version Version `json:"version"`
	Books   []Book  `json:"books"`
}

// Version holds the name and abbreviation of the bible version
type Version struct {
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
}

// GetBook returns a pointer to a Book struct by the name of the book
func (b *Bible) GetBook(name string) *Book {
	for _, book := range b.Books {
		if book.Name == name {
			return &book
		}
	}
	return nil
}

func (b *Bible) GetName() string {
	return b.Version.Name
}

func (b *Bible) GetAbbrev() string {
	return b.Version.Abbrev
}

// Book holds the name and chapters of a book
type Book struct {
	Name     string    `json:"name"`
	Number   int       `json:"number"`
	Chapters []Chapter `json:"chapters"`
}

func (b *Book) GetName() string {
	return b.Name
}

// Use the book number to get the English name of the book
func (b *Book) GetEnglishName() string {
	return bibleNumberToEnglishBook[b.Number]
}

// Chapter holds the number and verses of a chapter
type Chapter struct {
	Number int     `json:"number"`
	Name   string  `json:"name"`
	Verses []Verse `json:"verses"`
}

// Verse holds the number and text of a verse
type Verse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func (b *Book) GetChapter(number int) *Chapter {
	for _, chapter := range b.Chapters {
		if chapter.Number == number {
			return &chapter
		}
	}
	return nil
}

func (c *Chapter) GetVerseCount() int {
	return len(c.Verses)
}

func (c *Chapter) GetVerse(number int) *Verse {
	for _, verse := range c.Verses {
		if verse.Number == number {
			return &verse
		}
	}
	return nil
}

func (c *Chapter) GetVerses(start, end int) []Verse {
	var verses []Verse
	for _, verse := range c.Verses {
		if verse.Number >= start && verse.Number <= end {
			verses = append(verses, verse)
		}
	}
	return verses
}
