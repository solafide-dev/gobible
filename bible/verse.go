package bible

// Verse holds the number and text of a verse
type Verse struct {
	Number     int          `json:"number"`
	Text       string       `json:"text"`
	Footnotes  []Footnote   `json:"footnotes,omitempty"`  // optional
	Formatting []Formatting `json:"formatting,omitempty"` // optional
}

type Footnote struct {
	Marker string `json:"number"`
	Text   string `json:"text"`
	Start  int    `json:"start"`
	End    int    `json:"end"`
}

type FormattingType string

// define valid FormattingType values
const (
	FormatBold      FormattingType = "bold"
	FormatItalic    FormattingType = "italic"
	FormatUnderline FormattingType = "underline"
	FormatRedLetter FormattingType = "red-letter"
	FormatSmallCaps FormattingType = "small-caps"
)

type Formatting struct {
	Type  FormattingType `json:"type"`
	Start int            `json:"start"`
	End   int            `json:"end"`
}
