package gobible

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"

	bible "github.com/applehat/gobible/bible"
)

// Open a bible file and return a pointer to a Bible struct
func NewBible(file string) *bible.Bible {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bible bible.Bible
	err = json.NewDecoder(f).Decode(&bible)
	if err != nil {
		log.Fatal(err)
	}

	return &bible
}

// Create a Bible struct from a jSON string
func NewBibleFromString(bibleJSON string) *bible.Bible {
	var bible bible.Bible
	err := json.Unmarshal([]byte(bibleJSON), &bible)
	if err != nil {
		log.Fatal(err)
	}

	return &bible
}

//go:embed data/KJV.json
var KJV string

// Return a pointer to a Bible struct with the King James Version
// The KJV is embedded in the binary by default
func NewBibleKJV() *bible.Bible {
	return NewBibleFromString(KJV)
}
