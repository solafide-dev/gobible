package importer

import (
	"encoding/json"
	"os"

	bible "github.com/solafide-dev/gobible/bible"
)

type GoBible struct {
}

func (g *GoBible) Import(filename string) (*bible.Bible, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := bible.Bible{}

	err = json.NewDecoder(file).Decode(&b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func NewGoBibleImporter() Importer {
	return &GoBible{}
}
