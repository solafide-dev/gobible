package bible

// Verse holds the number and text of a verse
type Verse struct {
	Number    int           `json:"number"`
	Text      BibleMarkdown `json:"text"`
	Footnotes []Footnote    `json:"footnotes,omitempty"` // optional
}

// Footnote holds the text and location of a footnote
type Footnote struct {
	Text  BibleMarkdown `json:"text"`
	Start int           `json:"start"`
	End   int           `json:"end"`
}
