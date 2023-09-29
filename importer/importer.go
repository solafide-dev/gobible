package importer

import bible "github.com/solafide-dev/gobible/bible"

type Importer interface {
	Import(string) (*bible.Bible, error)
}

type ImporterFactory func() Importer

var importers = map[string]ImporterFactory{}

func RegisterImporter(name string, factory ImporterFactory) {
	importers[name] = factory
}

func GetImporter(name string) Importer {
	return importers[name]()
}

func GetImporterNames() []string {
	names := []string{}
	for name := range importers {
		names = append(names, name)
	}
	return names
}

func Import(name, file string) (*bible.Bible, error) {
	importer := GetImporter(name)
	return importer.Import(file)
}

func init() {
	RegisterImporter("osis", NewOsisImporter)
	RegisterImporter("gobible", NewGoBibleImporter)
}
