package parsed

import (
	"errors"
	"strings"
)

// WordType designates semantic meaning of a word.
type WordType int

const (
	UnknownType WordType = iota
	ComparisonMarkerType
	CultivarType
	ApproxMarkerType
	AuthorWordType
	AuthorWordFiliusType
	CandidatusType
	GenusType
	InfraspEpithetType
	HybridCharType
	GraftChimeraCharType
	RankType
	SpEpithetType
	SubgenusType
	SuperspType
	UninomialType
	YearApproximateType
	YearType
)

var wordTypeMap = map[WordType]string{
	UnknownType:          "WORD",
	ComparisonMarkerType: "COMPARISON_MARKER",
	CultivarType:					"CULTIVAR",
	ApproxMarkerType:     "APPROXIMATION_MARKER",
	AuthorWordType:       "AUTHOR_WORD",
	AuthorWordFiliusType: "AUTHOR_WORD_FILIUS",
	CandidatusType:       "CANDIDATUS",
	GenusType:            "GENUS",
	HybridCharType:       "HYBRID_CHAR",
	GraftChimeraCharType: "GRAFT_CHIMERA_CHAR",
	InfraspEpithetType:   "INFRASPECIES",
	RankType:             "RANK",
	SpEpithetType:        "SPECIES",
	SubgenusType:         "INFRA_GENUS",
	UninomialType:        "UNINOMIAL",
	YearApproximateType:  "APPROXIMATE_YEAR",
	YearType:             "YEAR",
}

var wordTypeStrMap = func() map[string]WordType {
	res := make(map[string]WordType)
	for k, v := range wordTypeMap {
		res[v] = k
	}
	return res
}()

// String is an implementation of fmt.Stringer interface.
func (wt WordType) String() string {
	return wordTypeMap[wt]
}

// MarshalJSON implements json.Marshaler.
func (wt WordType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + wt.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller.
func (wt *WordType) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	// strings.Trim seems to be ~10 time faster here than
	// json-iter Unmarshal
	s := strings.Trim(string(bs), `"`)
	*wt, ok = wordTypeStrMap[s]
	if !ok {
		err = errors.New("cannot decode WordType")
	}
	return err
}
