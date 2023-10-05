package bible

// Chapter holds the number and verses of a chapter
type Chapter struct {
	Number  int           `json:"number"`
	Name    string        `json:"name"`
	Title   BibleMarkdown `json:"title,omitempty"`
	Content BibleMarkdown `json:"content,omitempty"` // optional content that is not a verse, in markdown format
	Notes   BibleMarkdown `json:"notes,omitempty"`   // optional notes that are not a verse, in markdown format
	Verses  []Verse       `json:"verses"`
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
	for k := range c.Verses {
		verse := &c.Verses[k]
		if verse.Number >= start && verse.Number <= end {
			verses = append(verses, *verse)
		}
	}
	return verses
}

// Set chapter title
func (c *Chapter) SetTitle(title BibleMarkdown) {
	c.Title = title
}

// Set verse text, optionally adding the verse if its missing
func (c *Chapter) SetVerse(number int, text BibleMarkdown) {
	for k := range c.Verses {
		verse := &c.Verses[k]
		if verse.Number == number {
			verse.Text = text
			return
		}
	}
	c.Verses = append(c.Verses, Verse{
		Number: number,
		Text:   text,
	})
}
