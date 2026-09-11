package parser

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gnames/gnparser/ent/parsed"
)

// tailKind is a category of a tail annotation.
type tailKind int

const (
	sensuKind tailKind = iota
	statusKind
	publicationKind
)

// authorUse tells if a tail annotation cites an author.
type authorUse int

const (
	noAuthor authorUse = iota
	optionalAuthor
	requiredAuthor
)

// tailMarker describes an annotation recognized by tail parsing.
type tailMarker struct {
	re         *regexp.Regexp
	normalized string
	kind       tailKind
	author     authorUse
	// relation is a type of a concept relation implied by the annotation.
	relation string
}

// tailMarkers are tried in order, so a marker that starts with another
// marker has to go first ("sensu lato" before "sensu").
var tailMarkers = []tailMarker{
	sensu(`(?i:sensu\s+lato)|s\.\s?l\.|s\.\s?lat\.`, "sensu lato",
		optionalAuthor, parsed.RelationBroader),
	sensu(`(?i:sensu\s+stricto)|s\.\s?s\.|s\.\s?str\.`, "sensu stricto",
		optionalAuthor, parsed.RelationNarrower),
	sensu(`(?i:sensu\s+auct\.?,?\s+non)`, "sensu auct. non",
		requiredAuthor, parsed.RelationMisapplicationOf),
	sensu(`(?i:sensu\s+auct\.?)`, "sensu auct.", noAuthor, ""),
	sensu(`(?i:auct\.?,?\s+non)`, "auct. non",
		requiredAuthor, parsed.RelationMisapplicationOf),
	sensu(`(?i:auct\.?)`, "auct.", noAuthor, ""),
	sensu(`(?i:sensu\.?)`, "sensu", requiredAuthor, parsed.RelationSameAs),
	sensu(`(?i:pro\s+parte|pro\s+p\.|p\.\s?p\.)`, "pro parte", noAuthor, ""),

	status(`(?i:nom\.\s?dub\.|nomen\s+dubium)`, "nomen dubium"),
	status(`(?i:nom\.\s?nud\.|nomen\s+nudum)`, "nomen nudum"),
	status(`(?i:nom\.\s?illeg\.|nomen\s+illegitimum)`, "nomen illegitimum"),
	status(`(?i:nom\.\s?cons\.|nomen\s+conservandum)`, "nomen conservandum"),
	status(`(?i:nom\.\s?rej\.|nomen\s+rejiciendum)`, "nomen rejiciendum"),
	status(`(?i:nom\.\s?prov\.|nomen\s+provisorium)`, "nomen provisorium"),
	status(`(?i:stat\.\s?nov\.|status\s+novus)`, "status novus"),
	status(`(?i:comb\.\s?nov\.|combinatio\s+nova)`, "combinatio nova"),
	status(`(?i:stat\.\s?rev\.|status\s+restitutus)`, "status restitutus"),

	publication(`(?i:ined\.)`, "ineditus", noAuthor),
	publication(`(?i:hort\.?)`, "hortorum", noAuthor),
	publication(`(?i:fide)`, "fide", requiredAuthor),
	publication(`(?i:emend\.?|em\.)`, "emend.", requiredAuthor),
	publication(`(?i:ex)`, "ex", requiredAuthor),
}

// relationRCC5 maps types of concept relations to RCC5 operators.
var relationRCC5 = map[string]string{
	parsed.RelationBroader:          ">",
	parsed.RelationNarrower:         "<",
	parsed.RelationSameAs:           "==",
	parsed.RelationMisapplicationOf: "|",
}

// scopeAnnots are concept-alignment annotations that extend or restrict
// a concept. They are interpreted relative to an author.
var scopeAnnots = map[string]struct{}{
	"sensu lato":    {},
	"sensu stricto": {},
	"pro parte":     {},
}

func tailRe(s string) *regexp.Regexp {
	return regexp.MustCompile(`^(?:` + s + `)`)
}

func sensu(re, norm string, au authorUse, rel string) tailMarker {
	return tailMarker{
		re: tailRe(re), normalized: norm, kind: sensuKind,
		author: au, relation: rel,
	}
}

func status(re, norm string) tailMarker {
	return tailMarker{re: tailRe(re), normalized: norm, kind: statusKind}
}

func publication(re, norm string, au authorUse) tailMarker {
	return tailMarker{
		re: tailRe(re), normalized: norm, kind: publicationKind, author: au,
	}
}

// tailMatch is an annotation found in a tail.
type tailMatch struct {
	marker   *tailMarker
	verbatim string
	author   string
	end      int
}

// ParseTail recognizes annotations in the tail of a parsed name-string.
// Recognized annotations are moved from Tail to TailAnnotations, and
// quality warnings are recalculated. Tail keeps the part that was not
// recognized. If nothing is recognized, the result is returned unchanged.
//
// The result is passed by value, so taking its address does not move
// the result of every name-string parsing to the heap.
func (p *Engine) ParseTail(res parsed.Parsed) parsed.Parsed {
	if !res.Parsed || res.Tail == "" {
		return res
	}
	ta, residual := p.parseTail(res.Tail)
	if ta == nil {
		return res
	}
	res.TailAnnotations = ta
	res.Tail = residual

	ws := make(map[parsed.Warning]struct{}, len(res.QualityWarnings)+1)
	for _, v := range res.QualityWarnings {
		ws[v.Warning] = struct{}{}
	}
	if residual == "" {
		delete(ws, parsed.TailWarn)
	}
	for _, w := range tailWarnings(ta, res.Authorship != nil) {
		ws[w] = struct{}{}
	}
	res.QualityWarnings = prepareWarnings(ws)
	res.ParseQuality = 1
	if len(res.QualityWarnings) > 0 {
		res.ParseQuality = res.QualityWarnings[0].Quality
	}
	return res
}

// parseTail returns annotations found in the tail and the part of the tail
// that was not recognized. If no annotations are found, it returns nil.
func (p *Engine) parseTail(tail string) (*parsed.TailAnnotations, string) {
	ta := &parsed.TailAnnotations{Verbatim: tail}
	var found bool
	var pos int
	for {
		i := skipTailSeps(tail, pos)
		if i == len(tail) {
			pos = i
			break
		}
		m, ok := p.tailItem(tail, i)
		if !ok {
			break
		}
		addTailAnnotation(ta, m)
		found = true
		pos = m.end
	}
	if !found {
		return nil, tail
	}
	return ta, tail[pos:]
}

// tailItem recognizes an annotation at s[i:]. The annotation can be
// enclosed in parentheses or square brackets.
func (p *Engine) tailItem(s string, i int) (tailMatch, bool) {
	var closing byte
	switch s[i] {
	case '(':
		closing = ')'
	case '[':
		closing = ']'
	}
	if closing == 0 {
		return p.matchTailMarker(s, i)
	}

	m, ok := p.matchTailMarker(s, skipSpaces(s, i+1))
	if !ok {
		return m, false
	}
	end := skipSpaces(s, m.end)
	if end == len(s) || s[end] != closing || !tailBoundary(s, end+1) {
		return tailMatch{}, false
	}
	m.end = end + 1
	return m, true
}

// matchTailMarker finds a marker at s[i:] and, if the marker cites an author,
// the author that follows the marker.
func (p *Engine) matchTailMarker(s string, i int) (tailMatch, bool) {
	for k := range tailMarkers {
		m := &tailMarkers[k]
		end, ok := m.matchAt(s, i)
		if !ok {
			continue
		}
		res := tailMatch{marker: m, verbatim: s[i:end], end: end}
		if m.author == noAuthor {
			return res, true
		}
		if auEnd := p.tailAuthor(s, end); auEnd > end {
			res.author = strings.TrimSpace(s[end:auEnd])
			res.end = auEnd
			return res, true
		}
		if m.author == optionalAuthor {
			return res, true
		}
	}
	return tailMatch{}, false
}

// matchAt returns the end of the marker, if s[i:] starts with it.
func (m *tailMarker) matchAt(s string, i int) (int, bool) {
	loc := m.re.FindStringIndex(s[i:])
	if loc == nil || !tailBoundary(s, i+loc[1]) {
		return 0, false
	}
	return i + loc[1], true
}

// tailAuthor returns the end of an authorship that follows s[:i] after
// a space. If there is no such authorship, it returns i.
func (p *Engine) tailAuthor(s string, i int) int {
	j := skipSpaces(s, i)
	if j == i || j == len(s) {
		return i
	}
	for k := range tailMarkers {
		if _, ok := tailMarkers[k].matchAt(s, j); ok {
			return i
		}
	}
	n := p.authorshipLen(s[j:])
	if n == 0 {
		return i
	}
	return j + n
}

// authorshipLen returns the length in bytes of an authorship at the start
// of s, or 0 if s does not start with an authorship. It uses the Authorship
// rule of the name-string grammar, so it has to run only after the result
// of the name-string parsing is created.
func (p *Engine) authorshipLen(s string) int {
	p.Buffer = s
	p.fullReset()
	if err := p.Parse(int(ruleAuthorship)); err != nil {
		return 0
	}
	var end uint32
	for _, t := range p.Tokens() {
		if t.pegRule == ruleAuthorship && t.end > end {
			end = t.end
		}
	}
	return p.runeToByteOffset[end]
}

func addTailAnnotation(ta *parsed.TailAnnotations, m tailMatch) {
	switch m.marker.kind {
	case sensuKind:
		sa := parsed.SensuAnnotation{
			Verbatim:   m.verbatim,
			Normalized: m.marker.normalized,
			Author:     m.author,
		}
		if rel := m.marker.relation; rel != "" {
			sa.ConceptRelation = &parsed.ConceptRelation{
				Type:            rel,
				RCC5:            relationRCC5[rel],
				ReferenceAuthor: m.author,
			}
		}
		ta.Sensu = append(ta.Sensu, sa)
	case statusKind:
		ta.Status = append(ta.Status, parsed.StatusAnnotation{
			Verbatim:   m.verbatim,
			Normalized: m.marker.normalized,
		})
	case publicationKind:
		ta.Publication = append(ta.Publication, parsed.PublicationAnnotation{
			Verbatim:   m.verbatim,
			Normalized: m.marker.normalized,
			Author:     m.author,
		})
	}
}

// sensuAuthorAnnots are annotations whose cited author might also be the
// author of the name, if the name has no authorship.
var sensuAuthorAnnots = map[string]struct{}{
	"sensu":         {},
	"sensu lato":    {},
	"sensu stricto": {},
}

// tailWarnings returns warnings about recognized annotations. An author
// after sensu, sensu lato or sensu stricto is ambiguous if the name has no
// authorship: it might be the author of the name or of the concept. An
// annotation that extends or restricts a concept is not anchored if neither
// the name nor any annotation provides an author.
func tailWarnings(
	ta *parsed.TailAnnotations,
	hasAuthorship bool,
) []parsed.Warning {
	var res []parsed.Warning
	var scope, cited, ambiguous bool
	for _, v := range ta.Publication {
		if v.Author != "" {
			cited = true
		}
	}
	for _, v := range ta.Sensu {
		if _, ok := scopeAnnots[v.Normalized]; ok {
			scope = true
		}
		if v.Author == "" {
			continue
		}
		cited = true
		if _, ok := sensuAuthorAnnots[v.Normalized]; ok && !hasAuthorship {
			ambiguous = true
		}
	}
	if ambiguous {
		res = append(res, parsed.SensuAuthorAmbiguousWarn)
	}
	if scope && !hasAuthorship && !cited {
		res = append(res, parsed.SensuNoAuthorWarn)
	}
	return res
}

// tailBoundary checks if an annotation can end at s[i].
func tailBoundary(s string, i int) bool {
	if i >= len(s) {
		return true
	}
	r, _ := utf8.DecodeRuneInString(s[i:])
	return unicode.IsSpace(r) || strings.ContainsRune(",;)]", r)
}

// skipTailSeps returns the position of the first character after i that
// does not separate annotations.
func skipTailSeps(s string, i int) int {
	for i < len(s) {
		r, size := utf8.DecodeRuneInString(s[i:])
		if !unicode.IsSpace(r) && r != ',' && r != ';' {
			break
		}
		i += size
	}
	return i
}

func skipSpaces(s string, i int) int {
	for i < len(s) {
		r, size := utf8.DecodeRuneInString(s[i:])
		if !unicode.IsSpace(r) {
			break
		}
		i += size
	}
	return i
}
