// Package dict provides lookup data for gnparser.
package dict

import (
	"bufio"
	"embed"
	"fmt"
	"log/slog"
	"os"
)

//go:embed data
var data embed.FS

// Dict contains loaded dictionaries
var Dict *Dictionary = LoadDictionary()

// Dictionary contains dictionaries used for detecting information
// about scientific names
type Dictionary struct {
	// Bacteria contains bacterial genera, where boolean value is true if
	// we are aware of homonyms from other codes.
	Bacteria map[string]bool
	// AuthorICN contains family names of ICN authors of genera names.
	// This list is used to detect ICN name-strings so we can parse a word in
	// parenthesis after genus word as an author instead of subgenus.
	AuthorICN map[string]struct{}
}

// LoadDictionary creates dictionary from text files.
func LoadDictionary() *Dictionary {
	d := Dictionary{
		Bacteria:  readBacterialData(),
		AuthorICN: readAuthorICNData(),
	}
	return &d
}

func readBacterialData() map[string]bool {
	m := make(map[string]bool)
	scanBacterialFile("bacteria_genera.txt", false, m)
	scanBacterialFile("bacteria_genera_homonyms.txt", true, m)
	return m
}

func readAuthorICNData() map[string]struct{} {
	m := make(map[string]struct{})
	scanAuthorICNFIle("genera_auth_icn.txt", m)
	return m
}

func scanAuthorICNFIle(path string, m map[string]struct{}) {
	path = fmt.Sprintf("data/%s", path)
	f, err := data.Open(path)
	if err != nil {
		slog.Error("Cannot open authors' file", "error", err, "path", path)
		os.Exit(1)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m[sc.Text()] = struct{}{}
	}
}

func scanBacterialFile(path string, isHomonym bool, m map[string]bool) {
	path = fmt.Sprintf("data/%s", path)
	f, err := data.Open(path)
	if err != nil {
		slog.Error("Cannot open bacteria file", "error", err, "path", path)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m[sc.Text()] = isHomonym
	}
}
