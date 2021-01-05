package output

import (
	"errors"
	"strings"
)

type Annotation int

const (
	NoAnnot Annotation = iota
	SurrogateAnnot
	ComparisonAnnot
	ApproximationAnnot
	BOLDAnnot
	HybridAnnot
	NamedHybridAnnot
	HybridFormulaAnnot
	NothoHybridAnnot
)

var annotMap = map[Annotation]string{
	NoAnnot:            "",
	SurrogateAnnot:     "SURROGATE",
	ComparisonAnnot:    "COMPARISON",
	ApproximationAnnot: "APPROXIMATION",
	BOLDAnnot:          "BOLD_SURROGATE",
	HybridAnnot:        "HYBRID",
	NamedHybridAnnot:   "NAMED_HYBRID",
	HybridFormulaAnnot: "HYBRID_FORMULA",
	NothoHybridAnnot:   "NOTHO_HYBRID",
}

var annotStrMap = func() map[string]Annotation {
	res := make(map[string]Annotation)
	for k, v := range annotMap {
		res[v] = k
	}
	return res
}()

func (a Annotation) String() string {
	return annotMap[a]
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int is null.
func (a Annotation) MarshalJSON() ([]byte, error) {
	return []byte("\"" + a.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller.
func (a *Annotation) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	// strings.Trim seems to be ~10 time faster here than
	// json-iter Unmarshal
	s := strings.Trim(string(bs), `"`)
	*a, ok = annotStrMap[s]
	if !ok {
		err = errors.New("cannot decode Annotation")
	}
	return err
}
