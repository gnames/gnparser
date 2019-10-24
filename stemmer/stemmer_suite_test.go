package stemmer_test

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var stemsDict map[string]string

func TestStemmer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stemmer Suite")
}

var _ = BeforeSuite(func() {
	stemsDict = stemData()
})

func stemData() map[string]string {
	res := make(map[string]string)
	path := filepath.Join("..", "testdata", "stems.txt")
	f, err := os.Open(path)
	Expect(err).To(BeNil())
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		l := strings.TrimSpace(scan.Text())
		ws := regexp.MustCompile(`\s+`).Split(l, 2)
		res[ws[0]] = ws[1]
	}

	Expect(scan.Err()).To(BeNil())

	return res
}
