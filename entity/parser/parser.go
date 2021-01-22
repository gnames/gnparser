package parser

import (
	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/preprocess"
	"github.com/gnames/gnparser/entity/str"
)

// PreprocessAndParse takes a string and returns back the Abstract
// Syntax Tree of the scientific names. The AST is later used to
// create the final output.
func (p *Engine) PreprocessAndParse(
	s, ver string,
	keepHTML bool,
) ScientificNameNode {

	originalString := s
	tagsOrEntities := false
	if !keepHTML {
		s = preprocess.StripTags(s)
		if originalString != s {
			tagsOrEntities = true
		}
	}
	preproc := preprocess.Preprocess([]byte(s))

	defer func() {
		if len(preproc.Tail) > 0 {
			p.sn.tail += string(preproc.Tail)
		}
		if len(p.sn.tail) > 0 {
			p.addWarn(o.TailWarn)
			if str.IsBoldSurrogate(p.sn.tail) {
				p.sn.cardinality = 0
				annot := o.BOLDAnnot
				p.sn.surrogate = &annot
			}
		}
		p.sn.warnings = p.warnings
		p.sn.addVerbatim(originalString)
		p.sn.parserVersion = ver
	}()

	if preproc.NoParse {
		p.newNotParsedScientificNameNode(preproc)
		return p.sn
	}

	p.Buffer = string(preproc.Body)
	p.fullReset()
	if tagsOrEntities {
		p.addWarn(o.HTMLTagsEntitiesWarn)
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
	return p.sn
}
