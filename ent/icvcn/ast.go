package icvcn

import (
	"strings"

	"github.com/gnames/gnparser/ent/parsed"
)

// blacklist contains non-ICVCN names that happen to match ICVCN suffix patterns.
// These are legitimate botanical/zoological genera that should not be parsed
// as ICVCN names when they appear as uninomials.
var blacklist = map[string]struct{}{
	"Calviria":     {},
	"Caviria":      {},
	"Corvira":      {},
	"Dravira":      {},
	"Elvira":       {},
	"Eugivira":     {},
	"Euvira":       {},
	"Givira":       {},
	"Kaviria":      {},
	"Lussanvira":   {},
	"Mahavira":     {},
	"Miracavira":   {},
	"Navira":       {},
	"Paracalviria": {},
	"Roselviria":   {},
	"Rovira":       {},
	"Selviria":     {},
}

func Parse(inp string) *Parsed {
	// Check blacklist before attempting to parse
	trimmed := strings.TrimSpace(inp)
	if _, ok := blacklist[trimmed]; ok {
		return &Parsed{
			Input:  inp,
			Parsed: false,
		}
	}

	p := &Parser{Buffer: inp}
	p.Init()
	res := p.ParseToStruct()
	return &res
}

// ParseToStruct parses the input and returns a ParsedICVCN structure.
// If parsing fails, the Error field will contain the error and Parsed will be false.
func (p *Parser) ParseToStruct() Parsed {
	res := Parsed{
		Input: p.Buffer,
	}

	// Attempt to parse the input
	err := p.Parse()
	if err != nil {
		res.Error = err
		res.Parsed = false
		return res
	}

	// If parse succeeded, walk the AST
	return p.walkAST()
}

func (p Parser) walkAST() Parsed {
	res := Parsed{
		Input:  p.Buffer,
		Parsed: true,
	}

	// Get the root of the AST
	root := p.AST()
	if root == nil {
		res.Parsed = false
		return res
	}

	// Navigate to the VirusName node
	node := root
	for node != nil {
		switch node.pegRule {
		case ruleVirusName:
			p.walkVirusName(node, &res)
			return res
		}
		node = node.up
	}

	return res
}

func (p Parser) walkVirusName(n *node32, res *Parsed) {
	if n.up == nil {
		return
	}

	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleSpecies:
			p.walkSpecies(n, res)
			return
		case ruleUninomial:
			p.walkUninomial(n, res)
			return
		}
		n = n.next
	}
}

func (p Parser) walkSpecies(n *node32, res *Parsed) {
	if n.up == nil {
		return
	}

	var genusNode, epithetNode *node32

	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleGenus:
			genusNode = n
		case ruleSpeciesEpithet:
			epithetNode = n
		}
		n = n.next
	}

	if genusNode != nil {
		res.Genus = p.nodeValue(genusNode)
		res.Rank = Genus
		res.Uninomial = res.Genus

		// Add genus to Words slice
		word := parsed.Word{
			Verbatim:   res.Genus,
			Normalized: res.Genus,
			Type:       parsed.GenusIcvcnType,
			Start:      int(genusNode.begin),
			End:        int(genusNode.end),
		}
		res.Words = append(res.Words, word)
	}

	if epithetNode != nil {
		res.Species = p.nodeValue(epithetNode)
		res.Rank = Species

		// Add species epithet to Words slice
		word := parsed.Word{
			Verbatim:   res.Species,
			Normalized: res.Species,
			Type:       parsed.SpeciesIcvcnType,
			Start:      int(epithetNode.begin),
			End:        int(epithetNode.end),
		}
		res.Words = append(res.Words, word)
	}
}

func (p Parser) walkUninomial(n *node32, res *Parsed) {
	if n.up == nil {
		return
	}

	n = n.up
	if n == nil {
		return
	}

	res.Uninomial = p.nodeValue(n)

	// Determine the rank and word type based on the rule
	var wordType parsed.WordType
	switch n.pegRule {
	case ruleRealm:
		res.Rank = Realm
		wordType = parsed.RealmIcvcnType
	case ruleSubrealm:
		res.Rank = Subrealm
		wordType = parsed.SubrealmIcvcnType
	case ruleKingdom:
		res.Rank = Kingdom
		wordType = parsed.KingdomIcvcnType
	case ruleSubkingdom:
		res.Rank = Subkingdom
		wordType = parsed.SubkingdomIcvcnType
	case rulePhylum:
		res.Rank = Phylum
		wordType = parsed.PhylumIcvcnType
	case ruleSubphylum:
		res.Rank = Subphylum
		wordType = parsed.SubphylumIcvcnType
	case ruleClass:
		res.Rank = Class
		wordType = parsed.ClassIcvcnType
	case ruleSubclass:
		res.Rank = Subclass
		wordType = parsed.SubclassIcvcnType
	case ruleOrder:
		res.Rank = Order
		wordType = parsed.OrderIcvcnType
	case ruleSuborder:
		res.Rank = Suborder
		wordType = parsed.SuborderIcvcnType
	case ruleFamily:
		res.Rank = Family
		wordType = parsed.FamilyIcvcnType
	case ruleSubfamily:
		res.Rank = Subfamily
		wordType = parsed.SubfamilyIcvcnType
	case ruleGenus:
		res.Rank = Genus
		res.Genus = res.Uninomial
		wordType = parsed.GenusIcvcnType
	default:
		wordType = parsed.UninomialType
	}

	// Add to Words slice
	word := parsed.Word{
		Verbatim:   res.Uninomial,
		Normalized: res.Uninomial,
		Type:       wordType,
		Start:      int(n.begin),
		End:        int(n.end),
	}
	res.Words = append(res.Words, word)
}

func (p Parser) nodeValue(n *node32) string {
	if n == nil {
		return ""
	}
	return string(p.buffer[n.begin:n.end])
}
