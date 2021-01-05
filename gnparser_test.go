package gnparser_test

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/config"
	output "github.com/gnames/gnparser/entity/output"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	name     string
	jsonData string
}

func TestParseName(t *testing.T) {
	cfg := config.NewConfig(
		config.OptWithDetails(true),
		config.OptFormat("compact"),
		config.OptIsTest(true),
	)
	gnp := gnparser.NewGNParser(cfg)
	data := getTestData(t)
	for _, v := range data {
		parsed := gnp.ParseName(v.name)
		json := parsed.Output(gnp.Format())
		assert.Equal(t, json, v.jsonData, v.name)
	}
}

func getTestData(t *testing.T) []testData {
	var res []testData
	path := filepath.Join("testdata", "test_data.md")
	f, err := os.Open(path)
	assert.Nil(t, err)
	scanner := bufio.NewScanner(f)
	var isName bool
	var count int
	var datum testData
	for scanner.Scan() {
		line := scanner.Text()
		if !isName {
			if strings.HasPrefix(line, "Name: ") {
				isName = true
				datum.name = line[6:]
			}
			continue
		}
		count++
		if count == 7 {
			datum.jsonData = line
			res = append(res, datum)
			isName = false
			count = 0
			datum = testData{}
		}
	}

	assert.Nil(t, scanner.Err())
	return res
}

// BenchmarkParse checks parsing event speed. Run it with:
// `go test -bench=. -benchmem -count=10 -run=XXX > bench.txt && benchstat bench.txt`
func BenchmarkParse(b *testing.B) {
	path := filepath.Join("testdata", "200k-lines.txt")
	count := 1000
	test := make([]string, count)
	cfgJSON := config.NewConfig(config.OptFormat("compact"))
	gnpJSON := gnparser.NewGNParser(cfgJSON)
	cfgDet := config.NewConfig(config.OptFormat("compact"), config.OptWithDetails(true))
	gnpDet := gnparser.NewGNParser(cfgDet)
	cfgCSV := config.NewConfig(config.OptFormat("csv"))
	gnpCSV := gnparser.NewGNParser(cfgCSV)
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if count == 0 {
			break
		}
		test = append(test, scanner.Text())
		count--
	}
	b.Run("Parse to object once", func(b *testing.B) {
		var p output.Parsed
		for i := 0; i < b.N; i++ {
			p = gnpCSV.ParseName("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("Parse to object once with Init", func(b *testing.B) {
		var p output.Parsed
		for i := 0; i < b.N; i++ {
			gnp := gnparser.NewGNParser(cfgCSV)
			p = gnp.ParseName("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("Parse to object", func(b *testing.B) {
		var p output.Parsed
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p = gnpCSV.ParseName(v)
			}
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})

	b.Run("Parse to JSON", func(b *testing.B) {
		var s string
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p := gnpJSON.ParseName(v)
				s = p.Output(gnpJSON.Format())
				if err != nil {
					panic(err)
				}
			}
		}
		_ = fmt.Sprintf("%d", len(s))
	})

	b.Run("Parse to JSON (Details)", func(b *testing.B) {
		var s string
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p := gnpJSON.ParseName(v)
				s = p.Output(gnpDet.Format())
			}
		}
		_ = fmt.Sprintf("%d", len(s))
	})

	b.Run("Parse to CSV", func(b *testing.B) {
		var s string
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p := gnpCSV.ParseName(v)
				s = p.Output(gnpCSV.Format())
			}
		}
		_ = fmt.Sprintf("%d", len(s))
	})
}
