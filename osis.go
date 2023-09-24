package gobible

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
	"strings"

	bible "github.com/simpleworship/gobible/bible"
)

type Osis struct {
	XMLName        xml.Name `xml:"osis"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xsi            string   `xml:"xsi,attr"`
	OsisText       struct {
		Text        string `xml:",chardata"`
		OsisRefWork string `xml:"osisRefWork,attr"`
		OsisIDWork  string `xml:"osisIDWork,attr"`
		Lang        string `xml:"lang,attr"`
		Header      struct {
			Text         string `xml:",chardata"`
			RevisionDesc struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				P    string `xml:"p"`
			} `xml:"revisionDesc"`
			Work struct {
				Text        string `xml:",chardata"`
				OsisWork    string `xml:"osisWork,attr"`
				Title       string `xml:"title"`
				Contributor string `xml:"contributor"`
				Creator     []struct {
					Text string `xml:",chardata"`
					Role string `xml:"role,attr"`
				} `xml:"creator"`
				Subject     string `xml:"subject"`
				Date        string `xml:"date"`
				Description string `xml:"description"`
				Publisher   string `xml:"publisher"`
				Type        struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"type"`
				Identifier struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"identifier"`
				Source   string `xml:"source"`
				Language struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"language"`
				Relation  string `xml:"relation"`
				Coverage  string `xml:"coverage"`
				Rights    string `xml:"rights"`
				Scope     string `xml:"scope"`
				RefSystem string `xml:"refSystem"`
			} `xml:"work"`
		} `xml:"header"`
		Div []struct {
			Text    string `xml:",chardata"`
			OsisID  string `xml:"osisID,attr"`
			Type    string `xml:"type,attr"`
			Chapter []struct {
				Text   string `xml:",chardata"`
				OsisID string `xml:"osisID,attr"`
				Verse  []struct {
					Text   string `xml:",chardata"`
					OsisID string `xml:"osisID,attr"`
				} `xml:"verse"`
			} `xml:"chapter"`
		} `xml:"div"`
	} `xml:"osisText"`
}

// Open an OSIS file, parse it, and return a Bible object
func NewOSIS(filename string) *bible.Bible {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var osis Osis
	err = xml.NewDecoder(file).Decode(&osis)
	if err != nil {
		log.Fatal(err)
	}

	b := bible.Bible{}

	b.Version = bible.Version{
		Name:   osis.OsisText.Header.Work.Title,
		Abbrev: osis.OsisText.Header.Work.OsisWork,
	}

	for _, book := range osis.OsisText.Div {
		if book.Type == "book" {
			// Verify the book is in the Protestant Bible
			if !bible.BookIsValid(book.OsisID) {
				continue
			}

			bb := bible.Book{}
			bb.Name = bible.NormalizeBookName(book.OsisID)
			bb.Number = bible.BookNumberFromName(bb.Name)
			for _, chapter := range book.Chapter {
				cc := bible.Chapter{}
				splitId := strings.Split(chapter.OsisID, ".")
				cc.Number, _ = strconv.Atoi(splitId[len(splitId)-1])
				cc.Name = bb.Name + " " + strconv.Itoa(cc.Number)
				for _, verse := range chapter.Verse {
					vv := bible.Verse{}
					splitVerseId := strings.Split(verse.OsisID, ".")
					vv.Number, _ = strconv.Atoi(splitVerseId[len(splitVerseId)-1])
					vv.Text = verse.Text
					cc.Verses = append(cc.Verses, vv)
				}
				bb.Chapters = append(bb.Chapters, cc)
			}
			b.Books = append(b.Books, bb)
		}
	}

	return &b
}
