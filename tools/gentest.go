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

	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/config"
	"github.com/gnames/gnparser/entity/parsed"
)

func genTestData() error {
	enc := encode.GNjson{}
	path := filepath.Join("..", "testdata", "test_data.md")
	outPath := filepath.Join("..", "testdata", "test_data_new.md")
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
	opts := []config.Option{config.OptIsTest(true), config.OptWithDetails(true)}
	cfg := config.New(opts...)
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
