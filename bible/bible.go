package bible

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
	name = NormalizeBookName(name)
	for _, book := range b.Books {
		if book.Name == name {
			return &book
		}
	}
	// if we get here, lets try to find the book by number.
	for index, book := range bibleNumberToEnglishBook {
		if book == name {
			return &b.Books[index]
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
