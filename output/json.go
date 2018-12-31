package output

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
	grm "gitlab.com/gogna/gnparser/grammar"
)

type Output struct {
	Quality       int           `json:"quality,omitempty"`
	Parsed        bool          `json:"parsed"`
	Verbatim      string        `json:"verbatim"`
	Surrogate     bool          `json:"surrogate"`
	Warnings      []Warning     `json:"qualityWarnings,omitempty"`
	Normalized    string        `json:"normalized"`
	CanonicalName *canonical    `json:"canonicalName,omitempty"`
	Virus         bool          `json:"virus"`
	Positions     []pos         `json:"positions,omitempty"`
	NameStringID  string        `json:"nameStringId"`
	ParserVersion string        `json:"parserVersion"`
	Hybrid        bool          `json:"hybrid"`
	Details       []interface{} `json:"details"`
	Bacteria      bool          `json:"bacteria"`
	Tail          string        `json:"unparsedTail,omitempty"`
}

func NewOutput(sn *grm.ScientificNameNode) *Output {
	var co *canonical
	var quality int
	var ws []Warning
	var ps []pos
	var parsed bool
	det := sn.Details()
	c, hybrid := sn.Canonical()
	if c != nil {
		co = &canonical{Value: c.Value, ValueRanked: c.ValueRanked}
		ws, quality = qualityAndWarnings(sn.Warnings)
		ps = convertPos(sn.Pos())
		parsed = true
	}

	o := Output{
		Parsed:        parsed,
		Quality:       quality,
		Warnings:      ws,
		Verbatim:      sn.Verbatim,
		NameStringID:  sn.VerbatimID,
		CanonicalName: co,
		Hybrid:        hybrid,
		Normalized:    sn.Value(),
		Positions:     ps,
		Tail:          sn.Tail,
		Details:       det,
		ParserVersion: "test_version",
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
	Value       string `json:"value"`
	ValueRanked string `json:"valueRanked"`
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
