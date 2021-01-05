package parser

import (
	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/preprocess"
)

func (p *Engine) PreprocessAndParse(
	s, ver string,
	keepHTML bool,
) ScientificNameNode {
	defer func() {
		p.sn.addVerbatim(s)
		p.sn.parserVersion = ver
	}()
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
		p.newNotParsedScientificNameNode(preproc)
	}

	p.Buffer = string(preproc.Body)
	p.fullReset()
	if tagsOrEntities {
		p.addWarn(o.HTMLTagsEntitiesWarn)
	}
	if len(preproc.Tail) > 0 {
		p.addWarn(o.TailWarn)
	}
	if preproc.Underscore {
		p.addWarn(o.SpaceNonStandardWarn)
	}
	err := p.Parse()
	if err != nil {
		p.error = err
		p.newNotParsedScientificNameNode(preproc)
		return p.sn
	}

	p.OutputAST()
	p.newScientificNameNode()
	if len(preproc.Tail) > 0 {
		p.sn.tail += string(preproc.Tail)
	}
	return p.sn
}
