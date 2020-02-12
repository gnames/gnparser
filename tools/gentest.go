// +build ignore

// Generates a new test_data_new.txt file out of test_data.txt using current
// parser output. We need to do this in cases when parser output is modified.
// Run `go run gentest.go`
package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"

	"gitlab.com/gogna/gnparser"
	"gitlab.com/gogna/gnparser/output"
)

func genTestData() error {
	var nameString string
	empty := regexp.MustCompile(`^\s*$`)
	comment := regexp.MustCompile(`^\s*#`)
	path := filepath.Join("..", "testdata", "test_data.txt")
	outPath := filepath.Join("..", "testdata", "test_data_new.txt")
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
	gnp := gnparser.NewGNparser(gnparser.OptIsTest())
	count := 0
	for sc.Scan() {
		line := sc.Text()
		if empty.MatchString(line) || comment.MatchString(line) {
			w.Write([]byte(line + "\n"))
			continue
		}
		count++
		switch count {
		case 1:
			nameString = line
			w.Write([]byte(nameString + "\n"))
			gnp.Parse(nameString)
			res := gnp.ParsedName()
			w.Write([]byte(res + "\n"))
			bs, err := gnp.ToJSON()
			if err != nil {
				return err
			}
			w.Write(bs)
			w.Write([]byte("\n"))
			sl := gnp.ToSlice()
			res = output.ToCSV(sl) + "\n"
			w.Write([]byte(res))
		case 4:
			count = 0
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
