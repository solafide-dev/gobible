package bible

// Chapter holds the number and verses of a chapter
type Chapter struct {
	Number int     `json:"number"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Verses []Verse `json:"verses"`
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
