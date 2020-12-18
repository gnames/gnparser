package pb_test

import (
	"bytes"
	fmt "fmt"
	"io/ioutil"
	"log"

	. "github.com/gnames/gnparser/pb"
	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/gnames/gnparser"
)

type TestData struct {
	Entries []TestEntry `json:"entries"`
}

type TestEntry struct {
	Name        string   `json:"name"`
	NameType    string   `json:"name_type"`
	Cardinality int32    `json:"cardinality"`
	AllAuth     []string `json:"all_auth"`
}

var _ = Describe("PB", func() {
	DescribeTable("ToPB",
		func(name, nameType string, cardinality int32, allAuth []string) {
			gnp := gnparser.NewGNparser()
			o := gnp.ParseToObject(name)
			Expect(o.Cardinality).To(Equal(cardinality))
			Expect(NameType_name[int32(o.NameType)]).To(Equal(nameType))
			Expect(o.Cardinality).To(Equal(cardinality))
			if len(allAuth) > 0 {
				Expect(o.Authorship.AllAuthors).To(Equal(allAuth))
			} else {
				Expect(o.Authorship).To(BeNil())
			}
		}, pbEntries()...)
})

func pbEntries() []TableEntry {
	var td TestData
	var entries []TableEntry
	data, err := ioutil.ReadFile("../testdata/test_pb.json")
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(data)
	err = jsoniter.NewDecoder(r).Decode(&td)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range td.Entries {
		testName := fmt.Sprintf("PB: %s", v.Name)
		te := Entry(testName, v.Name, v.NameType, v.Cardinality, v.AllAuth)
		entries = append(entries, te)
	}
	return entries
}
