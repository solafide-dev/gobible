package bible

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Reference struct {
	Book        string
	Chapter     int
	VerseNumber int
	Verse       Verse
}

// If the book reference is a short code, attempt to convert it to the full name and update the reference
func (r *Reference) NormalizeBook() {
	r.Book = NormalizeBookName(r.Book)
}

// Take a string and return a slice of Reference structs
// Example Inputs:
// 	John 3:16
// 	John 3:16-18
//  1Cor1.2-5
//  1Cor1.2-5,7
//  1Cor1.2-5,7-9
//  1Cor1.2-2Cor3.2
//  Matthew 5:1, John 3:16

func (b *Bible) ParseReference(reference string) ([]Reference, error) {
	parts := strings.Split(reference, " ")
	if len(parts) < 2 {
		return nil, errors.New("invalid reference format")
	}

	book := strings.Join(parts[:len(parts)-1], " ")
	chapterVersePart := parts[len(parts)-1]
	chapterVerseParts := strings.Split(chapterVersePart, "-")

	log.Println("book", book)
	log.Println("chapterVersePart", chapterVersePart)
	log.Println("chapterVerseParts", chapterVerseParts)

	var startChapter, endChapter, startVerse, endVerse int
	chapterParts := strings.Split(chapterVerseParts[0], ":")
	if len(chapterParts) != 2 {
		return nil, errors.New("invalid chapter format")
	}

	log.Println("chapterParts", chapterParts)

	startChapter, err := strconv.Atoi(chapterParts[0])
	if err != nil {
		return nil, err
	}

	log.Println("startChapter", startChapter)

	startVerse, err = strconv.Atoi(chapterParts[1])
	if err != nil {
		return nil, err
	}

	log.Println("startVerse", startVerse)

	log.Println("len(chapterVerseParts)", len(chapterVerseParts))

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
		log.Println("chapter", chapter)
		start, end := startVerse, endVerse

		log.Println("start", start)
		log.Println("end", end)

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

		log.Println("verses", verses)

		for _, verse := range verses {
			ref := Reference{
				Book:        book,
				Chapter:     chapter,
				VerseNumber: verse,
			}
			ref.NormalizeBook()
			references = append(references, ref)
		}
	}

	return references, nil
}
