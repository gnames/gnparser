package parsed

import (
	"strings"

	"github.com/gnames/gnparser/ent/stemmer"
)

// RestoreAmbiguous method is used for cases where specific or infra-specific
// epithets had to be changed to be parsed sucessfully. Such situation
// arises when an epithet is the same as some word that is also an
// annotation, a prefix/suffix of an author name etc.
func (p *Parsed) RestoreAmbiguous(epithet, subst string) {
	stem := stemmer.Stem(epithet).Stem
	stemSubst := stemmer.Stem(subst).Stem
	p.Normalized = restoreString(p.Normalized, epithet, subst)
	p.Canonical.Full = restoreString(p.Canonical.Full, epithet, subst)
	p.Canonical.Simple = restoreString(p.Canonical.Simple, epithet, subst)
	p.Canonical.Stemmed = restoreString(p.Canonical.Stemmed, stem, stemSubst)

	for i := range p.Words {
		p.Words[i].Verbatim = restoreWord(p.Words[i].Verbatim, epithet, subst)
		p.Words[i].Normalized = restoreWord(p.Words[i].Normalized, epithet, subst)
	}

	if sp, ok := p.Details.(DetailsSpecies); ok {
		sp.Species.Species = restoreWord(sp.Species.Species, epithet, subst)
		p.Details = sp
	}
}

func restoreString(s, epithet, subst string) string {
	words := strings.Split(s, " ")
	for i := range words {
		if strings.HasPrefix(words[i], subst) {
			words[i] = epithet + words[i][len(epithet):]
			return strings.Join(words, " ")
		}
	}
	return s
}

func restoreWord(w, epithet, subst string) string {
	if strings.HasPrefix(w, subst) {
		return epithet + w[len(epithet):]
	}
	return w
}
