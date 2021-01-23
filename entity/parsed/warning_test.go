package parsed_test

import (
	"testing"

	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnparser/entity/parsed"
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
		assert.Equal(t, data[i].annot.String(), data[i].res)
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
			`{"f1":"AuthEx","warning":"Ex authors are not required","f2":[2,3,4]}`},
	}
	enc := encode.GNjson{}
	var dob dataOb
	for i := range data {
		res, err := enc.Encode(data[i].dob)
		assert.Nil(t, err)
		assert.Equal(t, string(res), data[i].res)
		err = enc.Decode(res, &dob)
		assert.Nil(t, err)
		assert.Equal(t, dob.Warn, data[i].dob.Warn)
	}
}
