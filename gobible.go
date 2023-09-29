package gobible

import (
	_ "embed"
	"encoding/json"

	"github.com/pkg/errors"
	bible "github.com/solafide-dev/gobible/bible"
	importer "github.com/solafide-dev/gobible/importer"
)

type GoBible struct {
	bibles map[string]*bible.Bible
	loaded []string
}

// Load a translation into the GoBible struct in GoBible format
func (g *GoBible) Load(filename string) error {
	return g.load(filename, "gobible")
}

func (g *GoBible) Unload(t string) error {
	if _, ok := g.bibles[t]; ok {
		delete(g.bibles, t)
		return nil
	}
	return errors.New("translation " + t + " not found to unload")
}

// Load a translation into the GoBible struct in an alternative supported format
func (g *GoBible) LoadFormat(filename string, format string) error {
	return g.load(filename, format)
}

// Load a string containing a GoBible JSON string
func (g *GoBible) LoadString(bibleJSON string) error {
	var bible bible.Bible
	err := json.Unmarshal([]byte(bibleJSON), &bible)
	if err != nil {
		return err
	}

	if bible.Version.Abbrev == "" {
		return errors.New("bible abbreviation not found in JSON")
	}

	// check if the bible is already loaded
	if _, ok := g.bibles[bible.Version.Abbrev]; ok {
		return errors.New("bible of type " + bible.Version.Abbrev + " already loaded")
	}
	g.loaded = append(g.loaded, bible.Version.Abbrev)
	g.bibles[bible.Version.Abbrev] = &bible
	return nil
}

func (g *GoBible) GetBibleJSON(abbrev string) (string, error) {
	b, ok := g.bibles[abbrev]
	if !ok {
		return "", errors.New("bible not found")
	}

	bJSON, err := json.Marshal(b)
	if err != nil {
		return "", err
	}

	return string(bJSON), nil
}

func (g *GoBible) GetTranslation(t string) (*bible.Bible, error) {
	if _, ok := g.bibles[t]; !ok {
		return nil, errors.New("bible not found")
	}

	return g.bibles[t], nil
}

func (g *GoBible) GetTranslations() []string {
	return g.loaded
}

// internal load function using the importer package
func (g *GoBible) load(filename string, format string) error {
	formats := importer.GetImporterNames()
	for _, f := range formats {
		if f == format {
			// load the bible
			b, err := importer.Import(format, filename)
			if err != nil {
				return err
			}
			// check if the bible is already loaded
			if _, ok := g.bibles[b.Version.Abbrev]; ok {
				return errors.New("bible of type " + b.Version.Abbrev + " already loaded")
			}
			g.loaded = append(g.loaded, b.Version.Abbrev)
			g.bibles[b.Version.Abbrev] = b
			return nil
		}
	}
	return errors.New("invalid format " + format)
}

// Create an empty Bible
func NewGoBible() *GoBible {
	return &GoBible{
		bibles: map[string]*bible.Bible{},
	}
}
