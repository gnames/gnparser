package gnparser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestGnparser is part of ``ginkgo`` package and is exposed because we want to
// test some private libraries.
func TestGNparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GNparser Suite")
}

type testRecord struct {
	NameString string
	Parsed     string
	Compact    string
	Simple     string
}

func testData() ([]testRecord, error) {
	var tests []testRecord
	var test testRecord
	empty := regexp.MustCompile(`^\s*$`)
	comment := regexp.MustCompile(`^\s*#`)
	path := filepath.Join("testdata", "test_data.txt")
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	count := 0
	for sc.Scan() {
		line := sc.Text()
		if empty.MatchString(line) || comment.MatchString(line) {
			continue
		}
		count++
		switch count {
		case 1:
			test = testRecord{NameString: line}
		case 2:
			test.Parsed = line
		case 3:
			test.Compact = line
		case 4:
			test.Simple = line
			tests = append(tests, test)
			count = 0
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	makeBigFile(tests)
	return tests, nil
}

func makeBigFile(t []testRecord) error {
	path := filepath.Join("testdata", "200k-lines.txt")
	iterNum := 200000 / len(t)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := iterNum; i > 0; i-- {
		for _, v := range t {
			name := fmt.Sprintf("%s\n", v.NameString)
			_, err := f.Write([]byte(name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
