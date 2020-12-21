package parser

import (
	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/preprocess"
)

func (p *Engine) PreprocessAndParse(
	s, ver string,
	keepHTML bool,
) *ScientificNameNode {
	tagsOrEntities := false
	if !keepHTML {
		orig := s
		s = preprocess.StripTags(s)
		if orig != s {
			tagsOrEntities = true
		}
	}
	preproc := preprocess.Preprocess([]byte(s))

	if preproc.NoParse {
		p.NewNotParsedScientificNameNode(preproc)
		return p.SN
	}

	p.Buffer = string(preproc.Body)
	p.FullReset()
	if tagsOrEntities {
		p.AddWarn(o.HTMLTagsEntitiesWarn)
	}
	if len(preproc.Tail) > 0 {
		p.AddWarn(o.TailWarn)
	}
	if preproc.Underscore {
		p.AddWarn(o.SpaceNonStandardWarn)
	}
	err := p.Parse()
	if err != nil {
		p.Error = err
		p.NewNotParsedScientificNameNode(preproc)
		return p.SN
	}

	p.OutputAST()
	p.NewScientificNameNode()
	if len(preproc.Tail) > 0 {
		p.SN.Tail += string(preproc.Tail)
	}
	p.SN.AddVerbatim(s)
	p.SN.ParserVersion = ver
	return p.SN
}
