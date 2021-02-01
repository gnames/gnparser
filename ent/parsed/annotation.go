package parsed

import (
	"errors"
	"strings"
)

// Annotations are additional descriptions of a name type.
type Annotation int

const (
	// NoAnnot is absence of additional descriptions.
	NoAnnot Annotation = iota
	// SurrogateAnnot is a miscellaneous informal name.
	SurrogateAnnot
	// ComparisonAnnot name with comparison marker (cf.).
	ComparisonAnnot
	// ApproximationAnnot is a name with approximation annotation (sp., spp etc.)
	ApproximationAnnot
	// BOLDAnnot is a surrogate name created by BOLD project.
	BOLDAnnot
	// HybridAnnot is a miscellaneous hybrid name.
	HybridAnnot
	// NameHybridAnnot is a stable hybrid in botany with registered name.
	NamedHybridAnnot
	// HybridFormulaAnnot is a hybrid created by combination of 2 or more names.
	HybridFormulaAnnot
	// NothoHybridAnnot is a hybrid with notho- 'ranks'.
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

// String is an implementation of fmt.Stringer interface.
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
