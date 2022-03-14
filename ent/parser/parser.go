// Package parser provides entities and methods to perform Parsing
// Expression Grammer parsing on scientific names.
package parser

import (
	"bytes"
	"fmt"

	"github.com/gnames/gnparser/ent/internal/preprocess"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/str"
)

// Debug takes a string, parsers it, and returns a byte representation of
// the node tree
func (p *Engine) Debug(s string) []byte {
	ppr := preprocess.Preprocess(p.preParser, []byte(s))
	var b bytes.Buffer
	if ppr.NoParse || ppr.Virus {
		b.WriteString("\n*** Preprocessing: NO PARSE ***\n")
		b.WriteString(fmt.Sprintf("\n%s\n", s))
		return b.Bytes()
	}
	p.Buffer = string(ppr.Body)
	fmt.Println(p.Buffer)
	p.fullReset()
	p.parse()
	p.outputAST()
	b.WriteString("\n*** Complete Syntax Tree ***\n")
	p.AST().PrettyPrint(&b, p.Buffer)
	b.WriteString("\n*** Output Syntax Tree ***\n")
	p.PrintOutputSyntaxTree(&b)
	return b.Bytes()
}

// PreprocessAndParse takes a string and returns back the Abstract
// Syntax Tree of the scientific names. The AST is later used to
// create the final output.
func (p *Engine) PreprocessAndParse(
	s, ver string,
	keepHTML bool,
	capitalize bool,
	enableCultivars bool,
	preserveDiaereses bool,
) ScientificNameNode {

	p.enableCultivars = enableCultivars
	p.preserveDiaereses = preserveDiaereses

	originalString := s
	var tagsOrEntities, lowCase bool
	if !keepHTML {
		s = preprocess.StripTags(s)
		if originalString != s {
			tagsOrEntities = true
		}
	}

	if capitalize {
		s = str.CapitalizeName(s)
		if s != originalString {
			lowCase = true
		}
	}

	preproc := preprocess.Preprocess(p.preParser, []byte(s))

	defer func() {
		p.sn.daggerChar = preproc.DaggerChar
		if len(preproc.Tail) > 0 {
			p.sn.tail += string(preproc.Tail)
		}
		if len(p.sn.tail) > 0 {
			p.addWarn(parsed.TailWarn)
			if str.IsBoldSurrogate(p.sn.tail) {
				p.sn.cardinality = 0
				annot := parsed.BOLDAnnot
				p.sn.surrogate = &annot
			}
		}

		p.sn.ambiguousEpithet = preproc.Ambiguous.Orig
		p.sn.ambiguousModif = preproc.Ambiguous.Subst

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
		p.addWarn(parsed.HTMLTagsEntitiesWarn)
	}

	if lowCase {
		p.addWarn(parsed.LowCaseWarn)
	}

	if preproc.Underscore {
		p.addWarn(parsed.SpaceNonStandardWarn)
	}
	err := p.Parse()

	if err != nil {
		p.error = err
		p.newNotParsedScientificNameNode(preproc)
		return p.sn
	}

	p.outputAST()
	p.newScientificNameNode()
	return p.sn
}
