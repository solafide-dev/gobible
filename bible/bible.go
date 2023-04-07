package bible

// Bible is the main struct that holds all the data
type Bible struct {
	Version Version `json:"version"`
	Extra   Extra   `json:"extra"` // optional
	Books   []Book  `json:"books"`
}

// Version holds the name and abbreviation of the bible version
type Version struct {
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
}

type Extra struct {
	Preface  string `json:"preface"`  // Markdown Preface
	Foreword string `json:"foreword"` // Markdown Forword
}

// GetBook returns a pointer to a Book struct by the name of the book
func (b *Bible) GetBook(name string) *Book {
	name = NormalizeBookName(name)
	for k := range b.Books {
		book := &b.Books[k]
		if book.Name == name {
			return book
		}
	}

	// Lookup by english name and book number
	for _, book := range BooksTable {
		if book.Name == name {
			for k := range b.Books {
				book2 := &b.Books[k]
				if book2.Number == book.BookNumber {
					return book2
				}
			}
		}
	}

	return nil
}

// For dealing with books that are not part of the biblical canon
func (b *Bible) AddBook(name string) *Book {
	// find the highest book number
	var max int
	for _, book := range b.Books {
		if book.Number > max {
			max = book.Number
		}
	}
	b.Books = append(b.Books, Book{
		Number: max + 1,
		Name:   name,
	})
	return &b.Books[len(b.Books)-1]
}

func (b *Bible) GetName() string {
	return b.Version.Name
}

func (b *Bible) GetAbbrev() string {
	return b.Version.Abbrev
}
