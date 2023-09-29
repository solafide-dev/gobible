package gobible

import (
	"errors"
	"strconv"
	"strings"

	bible "github.com/solafide-dev/gobible/bible"
)

type References []Reference

type Reference struct {
	Book        string
	Chapter     int
	Verse       int
	Translation string
	VerseRef    *bible.Verse
}

func (r *Reference) String() string {
	return r.VerseRef.Text + " - " + r.Book + " " + strconv.Itoa(r.Chapter) + ":" + strconv.Itoa(r.Verse) + " " + r.Translation
}

// Take a string and return a Refrences object
// Example Inputs:
//
//	John 3:16
//	John 3:16-18
//	1 Cor 1:2-5
//	1 Cor 1:2-2:5
//
// Also supports translations tagged at the end:
//
//	John 3:16 ESV
//	John 3:16 KJV
func (g *GoBible) ParseReference(reference string) (References, error) {
	// if we have no loaded bibles, return an error
	if len(g.bibles) == 0 {
		return nil, errors.New("no bibles loaded")
	}

	parts := strings.Split(reference, " ")
	if len(parts) < 2 {
		return nil, errors.New("invalid reference format")
	}

	// by default, we'll use the first translation loaded

	translation := g.loaded[0]

	// lets check if the last part happens to be a translation
	// if it is, we'll remove it from the parts slice
	// and set it as the translation for the reference
	potentialTranslation := strings.ToLower(parts[len(parts)-1])

	// check if potnetialTranslation starts with a number
	if _, err := strconv.Atoi(string(potentialTranslation[0])); err != nil {
		// nix this off the end of the parts slice
		parts = parts[:len(parts)-1]

		for t, _ := range g.bibles {
			if strings.ToLower(t) == potentialTranslation {
				translation = t
				break
			}
		}
	}

	b := g.bibles[translation]

	book := strings.Join(parts[:len(parts)-1], " ")
	chapterVersePart := parts[len(parts)-1]
	chapterVerseParts := strings.Split(chapterVersePart, "-")

	//log.Println("book", book)
	//log.Println("chapterVersePart", chapterVersePart)
	//log.Println("chapterVerseParts", chapterVerseParts)

	var startChapter, endChapter, startVerse, endVerse int
	chapterParts := strings.Split(chapterVerseParts[0], ":")
	if len(chapterParts) != 2 {
		return nil, errors.New("invalid chapter format")
	}

	//log.Println("chapterParts", chapterParts)

	startChapter, err := strconv.Atoi(chapterParts[0])
	if err != nil {
		return nil, err
	}

	//log.Println("startChapter", startChapter)

	startVerse, err = strconv.Atoi(chapterParts[1])
	if err != nil {
		return nil, err
	}

	//log.Println("startVerse", startVerse)

	//log.Println("len(chapterVerseParts)", len(chapterVerseParts))

	if len(chapterVerseParts) > 1 {
		chapterParts := strings.Split(chapterVerseParts[1], ":")
		if len(chapterParts) == 2 {
			endChapter, err = strconv.Atoi(chapterParts[0])
			if err != nil {
				return nil, err
			}
			endVerse, err = strconv.Atoi(chapterParts[1])
			if err != nil {
				return nil, err
			}
		} else if len(chapterParts) == 1 {
			endChapter = startChapter
			endVerse, err = strconv.Atoi(chapterParts[0])
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("invalid reference format")
		}
	} else {
		endChapter = startChapter
		endVerse = startVerse
		if err != nil {
			return nil, err
		}
	}

	if startChapter > endChapter {
		return nil, errors.New("invalid reference format: start chapter greater than end chapter")
	}

	var references []Reference
	for chapter := startChapter; chapter <= endChapter; chapter++ {
		//log.Println("chapter", chapter)
		start, end := startVerse, endVerse

		//log.Println("start", start)
		//log.Println("end", end)

		if chapter > startChapter {
			start = 1
		}
		if chapter < endChapter {
			end = b.GetBook(book).GetChapter(chapter).GetVerseCount()
			if err != nil {
				return nil, err
			}
		}
		verses := make([]int, end-start+1)

		for i := range verses {
			verses[i] = start + i
		}

		//log.Println("verses", verses)

		book = bible.NormalizeBookName(book)

		for _, verse := range verses {
			ref := Reference{
				Book:        book,
				Chapter:     chapter,
				Verse:       verse,
				Translation: translation,
				VerseRef:    b.GetBook(book).GetChapter(chapter).GetVerse(verse),
			}
			references = append(references, ref)
		}
	}

	return references, nil
}
