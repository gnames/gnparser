package parsed_test

import (
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser/ent/parsed"
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
		assert.Equal(t, data[i].res, data[i].annot.String())
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
	enc := gnfmt.GNjson{}
	var dob dataOb
	for i := range data {
		res, err := enc.Encode(data[i].dob)
		assert.Nil(t, err)
		assert.Equal(t, data[i].res, string(res))
		err = enc.Decode(res, &dob)
		assert.Nil(t, err)
		assert.Equal(t, data[i].dob.Annot, dob.Annot)
	}
}
