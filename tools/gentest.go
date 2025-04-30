//go:build ignore
// +build ignore

// Generates a new test_data_new.txt file out of test_data.txt using current
// parser output. We need to do this in cases when parser output is modified.
// Run `go run gentest.go`
package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
)

func genTestData() error {
	testFiles := []string{"test_data", "test_data_cultivars"}
	for _, v := range testFiles {
		err := newTestFile(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func newTestFile(file string) error {
	enc := gnfmt.GNjson{}
	path := filepath.Join("..", "testdata", file+".md")
	outPath := filepath.Join("..", "testdata", file+"_new.md")
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	w, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	defer w.Close()

	sc := bufio.NewScanner(f)
	opts := []gnparser.Option{gnparser.OptIsTest(true), gnparser.OptWithDetails(true)}
	if file == "test_data_cultivars" {
		opts = append(opts, gnparser.OptCode(nomcode.Cultivars))
	}
	cfg := gnparser.NewConfig(opts...)
	gnp := gnparser.New(cfg)
	var res parsed.Parsed
	isName := false
	var count int
	var can, au, nameString string
	var jsonData []byte
	for sc.Scan() {
		line := sc.Text()
		if !isName {
			w.Write([]byte(line + "\n"))
			if strings.HasPrefix(line, "Name: ") {
				isName = true
				nameString = line[6:]
				res = gnp.ParseName(nameString)
				jsonData, _ = enc.Encode(res)
				if res.Parsed {
					can = res.Canonical.Full
					if res.Authorship != nil {
						au = res.Authorship.Normalized
					}
				}
			}
			continue
		}
		count++
		switch count {
		case 2: // Canonical: name_here
			can = strings.TrimRight("Canonical: "+can, " ")
			w.Write([]byte(can + "\n"))
		case 4: // Authorship
			au = strings.TrimRight("Authorship: "+au, " ")
			w.Write([]byte(au + "\n"))
		case 7:
			w.Write(jsonData)
			w.Write([]byte("\n"))
			count = 0
			isName = false
			can, au = "", ""
			jsonData = []byte("")
		default:
			w.Write([]byte(line + "\n"))
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	genTestData()
}
