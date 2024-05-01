package gnparser_test

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnsys"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	name     string
	jsonData string
}

func TestParseName(t *testing.T) {
	cfg := gnparser.NewConfig(
		gnparser.OptWithDetails(true),
		gnparser.OptFormat("compact"),
		gnparser.OptIsTest(true),
	)
	gnp := gnparser.New(cfg)
	data := getTestData(t, "test_data.md")
	for _, v := range data {
		parsed := gnp.ParseName(v.name)
		json := parsed.Output(gnp.Format())
		assert.Equal(t, v.jsonData, json, v.name)
	}
}

func TestPool(t *testing.T) {
	assert := assert.New(t)
	cfg := gnparser.NewConfig()
	pool := gnparser.NewPool(cfg, 3)
	gnp := <-pool
	assert.NotNil(gnp)
	gnp2 := <-pool
	assert.NotNil(gnp2)
	gnp3 := <-pool
	assert.NotNil(gnp3)
	pd := gnp3.ParseName("Abarema clypearia")
	assert.True(pd.Parsed)
	pool <- gnp
	pool <- gnp2
	pool <- gnp3
}

func TestParseNameCultivars(t *testing.T) {
	cfg := gnparser.NewConfig(
		gnparser.OptWithDetails(true),
		gnparser.OptWithCultivars(true),
		gnparser.OptFormat("compact"),
		gnparser.OptIsTest(true),
	)
	gnp := gnparser.New(cfg)
	data := getTestData(t, "test_data_cultivars.md")
	for _, v := range data {
		parsed := gnp.ParseName(v.name)
		json := parsed.Output(gnp.Format())
		assert.Equal(t, v.jsonData, json, v.name)
	}
}

func TestParseLowCaseName(t *testing.T) {
	tests := []struct {
		msg, in, out string
		quality      int
	}{
		{"Caps", "Pardosa moesta", "Pardosa moesta", 1},
		{"LowCaps", "pardosa moesta", "Pardosa moesta", 4},
		{"Deutsch", "überweisen", "", 0},
	}
	cfg := gnparser.NewConfig(
		gnparser.OptWithCapitaliation(true),
	)
	gnp := gnparser.New(cfg)
	for _, v := range tests {
		parsed := gnp.ParseName(v.in)
		if v.out != "" {
			assert.Equal(t, v.out, parsed.Canonical.Simple, v.msg)
		} else {
			assert.Nil(t, parsed.Canonical)
		}
		assert.Equal(t, v.quality, parsed.ParseQuality, v.msg)
	}
}

func TestParsePreserveDiaereses(t *testing.T) {
	tests := []struct {
		msg, in, normalized, canonical string
		quality                        int
	}{
		{"DiaeresisInGenus", "Leptochloöpsis virgata", "Leptochloöpsis virgata", "Leptochloöpsis virgata", 1},
		{"DiaeresisInSpEpithet", "Hieracium samoënsicum", "Hieracium samoënsicum", "Hieracium samoënsicum", 1},
		{"DiaeresisInInfraSpEpithet", "Hieracium macilentum subsp. samoënsicum", "Hieracium macilentum subsp. samoënsicum", "Hieracium macilentum samoënsicum", 1},
		{"TransliteratesDiactiric", "Anthurium gudiñoi", "Anthurium gudinoi", "Anthurium gudinoi", 1},
	}
	cfg := gnparser.NewConfig(
		gnparser.OptWithPreserveDiaereses(true),
	)
	gnp := gnparser.New(cfg)
	for _, v := range tests {
		parsed := gnp.ParseName(v.in)
		assert.Equal(t, v.canonical, parsed.Canonical.Simple, v.msg)
		assert.Equal(t, v.normalized, parsed.Normalized, v.msg)
		assert.Equal(t, v.quality, parsed.ParseQuality, v.msg)
	}
}

func TestWordNormalizeByType(t *testing.T) {
	tests := []struct {
		msg, word, norm string
		wType           parsed.WordType
	}{
		{"b.", "B.", "b.", parsed.GenusType},
		{"betula", "Betula", "betula", parsed.GenusType},
		{"alba", "alba", "alb", parsed.SpEpithetType},
		{"Linn", "Linn.", "linn.", parsed.AuthorWordType},
		{"yr", "1888", "1888", parsed.YearType},
	}

	for _, v := range tests {
		res := parsed.NormalizeByType(v.word, v.wType)
		assert.Equal(t, v.norm, res, v.msg)
	}
}

func TestOutputRestore(t *testing.T) {
	name := "Homo zapiens Linn. 1758"
	cfg := gnparser.NewConfig(gnparser.OptWithDetails(true))
	gnp := gnparser.New(cfg)
	res := gnp.ParseName(name)
	res.RestoreAmbiguous("sapiens", "zapiens")
	assert.Equal(t, "Homo zapiens Linn. 1758", res.Verbatim)
	assert.Equal(t, "Homo sapiens Linn. 1758", res.Normalized)
	assert.Equal(t, "Homo sapiens", res.Canonical.Full)
	assert.Equal(t, "Homo sapiens", res.Canonical.Simple)
	assert.Equal(t, "Homo sapiens", res.Canonical.Stemmed)
	assert.Equal(t, "sapiens", res.Words[1].Verbatim)
	assert.Equal(t, "sapiens", res.Words[1].Normalized)
	sp, ok := res.Details.(parsed.DetailsSpecies)
	assert.True(t, ok)
	assert.Equal(t, "sapiens", sp.Species.Species)
}

func getTestData(t *testing.T, filename string) []testData {
	var res []testData
	path := filepath.Join("testdata", filename)
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

func Example() {
	names := []string{"Pardosa moesta Banks, 1892", "Bubo bubo"}
	cfg := gnparser.NewConfig()
	gnp := gnparser.New(cfg)
	res := gnp.ParseNames(names)
	fmt.Println(res[0].Authorship.Verbatim)
	fmt.Println(res[1].Canonical.Simple)
	fmt.Println(parsed.HeaderCSV(gnp.Format()))
	fmt.Println(res[0].Output(gnp.Format()))
	// Output:
	// Banks, 1892
	// Bubo bubo
	// Id,Verbatim,Cardinality,CanonicalStem,CanonicalSimple,CanonicalFull,Authorship,Year,Quality
	// e2fdf10b-6a36-5cc7-b6ca-be4d3b34b21f,"Pardosa moesta Banks, 1892",2,Pardosa moest,Pardosa moesta,Pardosa moesta,"Banks, 1892",1892,1
}

// BenchmarkParse checks parsing event speed. Run it with:
// `go test -bench=. -benchmem -count=10 -run=XXX > bench.txt && benchstat bench.txt`
func BenchmarkParse(b *testing.B) {
	path := filepath.Join("testdata", "200k-lines.txt")
	check200kFile(path)
	count := 1000
	test := make([]string, count)
	cfgJSON := gnparser.NewConfig(gnparser.OptFormat("compact"))
	gnpJSON := gnparser.New(cfgJSON)
	cfgDet := gnparser.NewConfig(gnparser.OptFormat("compact"), gnparser.OptWithDetails(true))
	gnpDet := gnparser.New(cfgDet)
	cfgCSV := gnparser.NewConfig(gnparser.OptFormat("csv"))
	gnpCSV := gnparser.New(cfgCSV)
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
		var p parsed.Parsed
		for i := 0; i < b.N; i++ {
			p = gnpCSV.ParseName("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("Parse to object once with Init", func(b *testing.B) {
		var p parsed.Parsed
		for i := 0; i < b.N; i++ {
			gnp := gnparser.New(cfgCSV)
			p = gnp.ParseName("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("Parse to object", func(b *testing.B) {
		var p parsed.Parsed
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

func check200kFile(path string) {
	exists, err := gnsys.FileExists(path)
	if exists && err == nil {
		return
	}

	names := getNames()
	iterNum := 200000 / len(names)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := iterNum; i > 0; i-- {
		for i := range names {
			name := fmt.Sprintf("%s\n", names[i])
			_, err := f.Write([]byte(name))
			if err != nil {
				panic(err)
			}
		}
	}
}

func getNames() []string {
	var err error
	path := filepath.Join("testdata", "test_data.md")
	var names []string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Name: ") {
			names = append(names, line[6:])
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return names
}
