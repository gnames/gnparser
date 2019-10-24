package output

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
	grm "gitlab.com/gogna/gnparser/grammar"
	"gitlab.com/gogna/gnparser/stemmer"
)

type Output struct {
	Parsed        bool          `json:"parsed"`
	Quality       int           `json:"quality"`
	Warnings      []Warning     `json:"qualityWarnings,omitempty"`
	Verbatim      string        `json:"verbatim"`
	Normalized    string        `json:"normalized,omitempty"`
	CanonicalName *canonical    `json:"canonicalName,omitempty"`
	Authorship    string        `json:"authorship,omitempty"`
	Details       []interface{} `json:"details,omitempty"`
	Positions     []pos         `json:"positions,omitempty"`
	Surrogate     bool          `json:"surrogate"`
	Virus         bool          `json:"virus"`
	Hybrid        bool          `json:"hybrid"`
	Bacteria      bool          `json:"bacteria"`
	Tail          string        `json:"unparsedTail,omitempty"`
	NameStringID  string        `json:"nameStringId"`
	ParserVersion string        `json:"parserVersion"`
}

func NewOutput(sn *grm.ScientificNameNode) *Output {
	var co *canonical
	var quality int
	var au string
	var ws []Warning
	var ps []pos
	var hybrid bool
	var parsed bool
	det := sn.Details()
	c := sn.Canonical()
	if c != nil {
		co = &canonical{
			Full:   c.ValueRanked,
			Simple: c.Value,
			Stem:   stemmer.StemCanonical(c.Value),
		}
		ws, quality = qualityAndWarnings(sn.Warnings)
		ps = convertPos(sn.Pos())
		hybrid = sn.Hybrid
		parsed = true
		lastAuthorship := sn.LastAuthorship()
		if lastAuthorship != nil {
			au = lastAuthorship.Value
		}

	}

	o := Output{
		Parsed:        parsed,
		Quality:       quality,
		Warnings:      ws,
		Verbatim:      sn.Verbatim,
		NameStringID:  sn.VerbatimID,
		Surrogate:     sn.Surrogate,
		CanonicalName: co,
		Virus:         sn.Virus,
		Hybrid:        hybrid,
		Normalized:    sn.Value(),
		Positions:     ps,
		Bacteria:      sn.Bacteria,
		Tail:          sn.Tail,
		Details:       det,
		Authorship:    au,
		ParserVersion: sn.ParserVersion,
	}
	return &o
}

func qualityAndWarnings(ws []grm.Warning) ([]Warning, int) {
	warns := prepareWarnings(ws)
	quality := 1
	if len(warns) > 0 {
		quality = warns[0].Quality
	}
	return warns, quality
}

// ToJSON converts Output to JSON representation.
func (o *Output) ToJSON(pretty bool) ([]byte, error) {
	if pretty {
		return jsoniter.MarshalIndent(o, "", "  ")
	}
	return jsoniter.Marshal(o)
}

// FromJSON converts JSON representation of Outout to Output object.
func FromJSON(data []byte) (Output, error) {
	var o Output
	r := bytes.NewReader(data)
	err := jsoniter.NewDecoder(r).Decode(&o)
	return o, err
}

type canonical struct {
	Full   string `json:"full"`
	Simple string `json:"simple"`
	Stem   string `json:"stem"`
}

type pos struct {
	Type  string
	Start int
	End   int
}

func convertPos(pp []grm.Pos) []pos {
	res := make([]pos, len(pp))
	for i, v := range pp {
		t, ok := wordTypeMap[v.Type]
		if !ok {
			t = "??????"
		}
		res[i] = pos{Type: t, Start: v.Start, End: v.End}
	}
	return res
}

func (p *pos) MarshalJSON() ([]byte, error) {
	arr := []interface{}{p.Type, p.Start, p.End}
	return jsoniter.Marshal(arr)
}

func (p *pos) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	err := jsoniter.Unmarshal(bs, &arr)
	if err != nil {
		return err
	}
	p.Type = arr[0].(string)
	p.Start = int(arr[1].(float64))
	p.End = int(arr[2].(float64))
	return nil
}
