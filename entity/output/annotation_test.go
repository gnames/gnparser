package output_test

import (
	"testing"

	"github.com/gnames/gnlib/encode"
	out "github.com/gnames/gnparser/entity/output"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	data := []struct {
		annot out.Annotation
		res   string
	}{
		{out.None, ""},
		{out.Comparison, "Comparison"},
		{out.ApproxSurrogate, "Approx. surrogate"},
		{out.Surrogate, "Surrogate"},
	}

	for i := range data {
		assert.Equal(t, data[i].annot.String(), data[i].res)
	}
}

func TestJSON(t *testing.T) {
	type dataOb struct {
		Field1 string         `json:"f1"`
		Annot  out.Annotation `json:"annot,omitempty"`
		Field2 []int          `json:"f2"`
	}
	data := []struct {
		dob dataOb
		res string
	}{
		{dataOb{"None", out.None, []int{}},
			`{"f1":"None","f2":[]}`},
		{dataOb{"Comparison", out.Comparison, []int{2, 3, 4}},
			`{"f1":"Comparison","annot":"Comparison","f2":[2,3,4]}`},
	}
	enc := encode.GNjson{}
	var dob dataOb
	for i := range data {
		res, err := enc.Encode(data[i].dob)
		assert.Nil(t, err)
		assert.Equal(t, string(res), data[i].res)
		err = enc.Decode(res, &dob)
		assert.Equal(t, dob.Annot, data[i].dob.Annot)
	}
}
