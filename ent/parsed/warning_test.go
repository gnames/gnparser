package parsed_test

import (
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/stretchr/testify/assert"
)

func TestStringWarn(t *testing.T) {
	data := []struct {
		annot parsed.Warning
		res   string
	}{
		{parsed.TailWarn, "Unparsed tail"},
	}

	for i := range data {
		assert.Equal(t, data[i].res, data[i].annot.String())
	}
}

func TestJSONWarn(t *testing.T) {
	type dataOb struct {
		Field1 string         `json:"f1"`
		Warn   parsed.Warning `json:"warning"`
		Field2 []int          `json:"f2"`
	}
	data := []struct {
		dob dataOb
		res string
	}{
		{dataOb{"Tail", parsed.TailWarn, []int{}},
			`{"f1":"Tail","warning":"Unparsed tail","f2":[]}`},
		{dataOb{"AuthEx", parsed.AuthExWarn, []int{2, 3, 4}},
			`{"f1":"AuthEx","warning":"` + "`ex`" + ` authors are not required (ICZN only)","f2":[2,3,4]}`},
	}
	enc := gnfmt.GNjson{}
	var dob dataOb
	for i := range data {
		res, err := enc.Encode(data[i].dob)
		assert.Nil(t, err)
		assert.Equal(t, data[i].res, string(res))
		err = enc.Decode(res, &dob)
		assert.Nil(t, err)
		assert.Equal(t, data[i].dob.Warn, dob.Warn)
	}
}
