package output

import (
	"errors"
	"strings"
)

// WordType designates semantic meaning of a word.
type WordType int

const (
	UnknownType WordType = iota
	ComparisonType
	ApproxType
	AuthorWordType
	AuthorWordFiliusType
	GenusType
	InfraSpEpithetType
	HybridCharType
	RankType
	SpEpithetType
	SubGenusType
	SuperSpType
	UninomialType
	YearApproximateType
	YearType
)

var wordTypeMap = map[WordType]string{
	UnknownType:          "word",
	ComparisonType:       "comparisonMarker",
	ApproxType:           "approximationMarker",
	AuthorWordType:       "authorWord",
	AuthorWordFiliusType: "authorWordFilius",
	GenusType:            "genus",
	HybridCharType:       "hybridChar",
	InfraSpEpithetType:   "infraspecificEpithet",
	RankType:             "rank",
	SpEpithetType:        "specificEpithet",
	SubGenusType:         "infragenericEpithet",
	UninomialType:        "uninomial",
	YearApproximateType:  "approximateYear",
	YearType:             "year",
}

var wordTypeStrMap = func() map[string]WordType {
	res := make(map[string]WordType)
	for k, v := range wordTypeMap {
		res[v] = k
	}
	return res
}()

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
