package parsed_test

import (
	"testing"

	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnparser/entity/parsed"
	"github.com/stretchr/testify/assert"
)

func TestStringAnnot(t *testing.T) {
	data := []struct {
		annot parsed.Annotation
		res   string
	}{
		{parsed.NoAnnot, ""},
		{parsed.ComparisonAnnot, "COMPARISON"},
		{parsed.ApproximationAnnot, "APPROXIMATION"},
		{parsed.SurrogateAnnot, "SURROGATE"},
	}

	for i := range data {
		assert.Equal(t, data[i].annot.String(), data[i].res)
	}
}

func TestJSONAnnot(t *testing.T) {
	type dataOb struct {
		Field1 string            `json:"f1"`
		Annot  parsed.Annotation `json:"annot"`
		Field2 []int             `json:"f2"`
	}
	data := []struct {
		dob dataOb
		res string
	}{
		{dataOb{"None", parsed.NoAnnot, []int{}},
			`{"f1":"None","annot":"","f2":[]}`},
		{dataOb{"Comparison", parsed.ComparisonAnnot, []int{2, 3, 4}},
			`{"f1":"Comparison","annot":"COMPARISON","f2":[2,3,4]}`},
	}
	enc := encode.GNjson{}
	var dob dataOb
	for i := range data {
		res, err := enc.Encode(data[i].dob)
		assert.Nil(t, err)
		assert.Equal(t, string(res), data[i].res)
		err = enc.Decode(res, &dob)
		assert.Nil(t, err)
		assert.Equal(t, dob.Annot, data[i].dob.Annot)
	}
}
