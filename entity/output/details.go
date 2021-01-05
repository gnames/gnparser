package output

type Uninomial struct {
	Value      string      `json:"uninomial"`
	Rank       string      `json:"rank,omitempty"`
	Parent     string      `json:"parent,omitempty"`
	Authorship *Authorship `json:"authorship,omitempty"`
}

type Species struct {
	Genus      string      `json:"genus"`
	SubGenus   string      `json:"subGenus,omitempty"`
	Species    string      `json:"species"`
	Authorship *Authorship `json:"authorship,omitempty"`
}

type InfraSpecies struct {
	Species
	InfraSpecies []InfraSpeciesElem `json:"infraSpecies,omitempty"`
}

type InfraSpeciesElem struct {
	Value      string      `json:"value"`
	Rank       string      `json:"rank,omitempty"`
	Authorship *Authorship `json:"authorship,omitempty"`
}

type Comparison struct {
	Genus             string      `json:"genus"`
	Species           string      `json:"species,omitempty"`
	SpeciesAuthorship *Authorship `json:"authorship,omitempty"`
	CompMarker        string      `json:"comparisonMarker"`
}

type Approximation struct {
	Genus             string      `json:"genus"`
	Species           string      `json:"species,omitempty"`
	SpeciesAuthorship *Authorship `json:"authorship,omitempty"`
	ApproxMarker      string      `json:"approximationMarker,omitempty"`
	Ignored           string      `json:"ignored,omitempty"`
}

type DetailsHybridFormula struct {
	HybridFormula []Details `json:"hybridFormula"`
}

func (DetailsHybridFormula) isDetails() {}

type DetailsUninomial struct {
	Uninomial Uninomial `json:"uninomial"`
}

func (DetailsUninomial) isDetails() {}

type DetailsSpecies struct {
	Species Species `json:"species"`
}

func (DetailsSpecies) isDetails() {}

type DetailsInfraSpecies struct {
	InfraSpecies InfraSpecies `json:"infraSpecies"`
}

func (DetailsInfraSpecies) isDetails() {}

type DetailsComparison struct {
	Comparison Comparison `json:"comparison"`
}

func (DetailsComparison) isDetails() {}

type DetailsApproximation struct {
	Approximation Approximation `json:"approximation"`
}

func (DetailsApproximation) isDetails() {}
