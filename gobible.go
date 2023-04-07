package gobible

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"

	bible "github.com/gobible/gobible/bible"
)

// Create an empty Bible
func NewEmpty() *bible.Bible {
	return &bible.Bible{}
}

// Open a bible file and return a pointer to a Bible struct
func New(file string) *bible.Bible {
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
func NewFromString(bibleJSON string) *bible.Bible {
	var bible bible.Bible
	err := json.Unmarshal([]byte(bibleJSON), &bible)
	if err != nil {
		log.Fatal(err)
	}

	return &bible
}

//go:embed data/WEB.json
var WEB string

//go:embed data/KJV.json
var KJV string

// Return a pointer to a Bible struct with the King James Version
// The WEB is embedded in the binary by default
func LoadInternal(version string) *bible.Bible {
	switch version {
	case "WEB":
		return NewFromString(WEB)
	case "KJV":
		return NewFromString(KJV)
	}
	log.Fatalf("Version %s not found", version)
	return &bible.Bible{}
}
