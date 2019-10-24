package pb

import (
	"gitlab.com/gogna/gnparser/output"
	"gitlab.com/gogna/gnparser/stemmer"
)

func ToPB(o *output.Output) *Parsed {
	po := &Parsed{
		Parsed:         o.Parsed,
		Quality:        int32(o.Quality),
		QualityWarning: qualityWarning(o),
		Verbatim:       o.Verbatim,
		Id:             o.NameStringID,
		Canonical:      canonicalName(o),
		Hybrid:         o.Hybrid,
		Normalized:     o.Normalized,
		Positions:      positions(o),
		Bacteria:       o.Bacteria,
		Tail:           o.Tail,
		ParserVersion:  o.ParserVersion,
	}
	if o.Virus {
		po.NameType = NameType_VIRUS
	} else if o.Surrogate {
		po.NameType = NameType_SURROGATE
	}

	details(po, o)
	return po
}

func canonicalName(o *output.Output) *Canonical {
	var cn *Canonical
	if o.CanonicalName == nil {
		return cn
	}
	cn = &Canonical{
		Stem:   stemmer.StemCanonical(o.CanonicalName.Simple),
		Simple: o.CanonicalName.Simple,
		Full:   o.CanonicalName.Full,
	}
	return cn
}

func positions(o *output.Output) []*Position {
	if len(o.Positions) == 0 {
		var p []*Position
		return p
	}
	res := make([]*Position, len(o.Positions))
	for i, v := range o.Positions {
		res[i] = &Position{
			Type:  v.Type,
			Start: int32(v.Start),
			End:   int32(v.End),
		}
	}
	return res
}

func qualityWarning(o *output.Output) []*QualityWarning {
	if len(o.Warnings) == 0 {
		var w []*QualityWarning
		return w
	}
	res := make([]*QualityWarning, len(o.Warnings))
	for i, v := range o.Warnings {
		res[i] = &QualityWarning{
			Quality: int32(v.Quality),
			Message: v.Message,
		}
	}
	return res
}
