package gobible

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
)

// Open a bible file and return a pointer to a Bible struct
func NewBible(file string) *Bible {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bible Bible
	err = json.NewDecoder(f).Decode(&bible)
	if err != nil {
		log.Fatal(err)
	}

	return &bible
}

// Create a Bible struct from a jSON string
func NewBibleFromString(bibleJSON string) *Bible {
	var bible Bible
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
func NewBibleKJV() *Bible {
	return NewBibleFromString(KJV)
}
