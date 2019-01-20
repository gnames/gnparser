package dict

import (
	"bufio"
	"log"

	"gitlab.com/gogna/gnparser/fs"
)

// Dict contains loaded dictionaries
var Dict *Dictionary = LoadDictionary()

// Dictionary contains dictionaries used for detecting information
// about scientific names
type Dictionary struct {
	// Bacteria contains bacterial genera, where boolean value is true if
	// we are aware of homonyms from other codes.
	Bacteria map[string]bool
}

// LoadDictionary creates dictionary from text files.
func LoadDictionary() *Dictionary {
	d := Dictionary{
		Bacteria: readBacterialData(),
	}
	return &d
}

func readBacterialData() map[string]bool {
	m := make(map[string]bool)
	scanFile("bacteria_genera.txt", false, m)
	scanFile("bacteria_genera_homonyms.txt", true, m)
	return m
}

func scanFile(path string, isHomonym bool, m map[string]bool) {
	f, err := fs.Files.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m[sc.Text()] = isHomonym
	}
}
