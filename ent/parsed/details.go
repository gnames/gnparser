package parsed

// Uninomial are details for names with cardinality 1.
type Uninomial struct {
	// Value is the uninomial name.
	Value string `json:"uninomial"`
	// Rank of the uninomial in a combination name, for example
	// "Pereskia subg. Maihuenia Philippi ex F.A.C.Weber, 1898"
	Rank string `json:"rank,omitempty"`
	// Cultivar is a value of a cultivar of a uninomial.
	Cultivar string `json:"cultivar,omitempty"`
	// Parent of a uninomial in a combination name.
	Parent string `json:"parent,omitempty"`
	// Authorship of the uninomial.
	Authorship *Authorship `json:"authorship,omitempty"`
}

// Species are details for binomial names with cardinality 2.
type Species struct {
	// Genus is a value of a genus of a binomial.
	Genus string `json:"genus"`
	// Subgenus is a value of subgenus of binomial.
	Subgenus string `json:"subgenus,omitempty"`
	// Species is a value of a specific epithet.
	Species string `json:"species"`
	// Cultivar is a value of a cultivar of a binomial.
	Cultivar string `json:"cultivar,omitempty"`
	// Authorship of the binomial.
	Authorship *Authorship `json:"authorship,omitempty"`
}

// Infraspecies are details for names with cardinality higher than 2.
type Infraspecies struct {
	// Species are details for the binomial part of a name.
	Species
	// Infraspecies is a slice of infraspecific epithets of a name.
	Infraspecies []InfraspeciesElem `json:"infraspecies,omitempty"`
}

// InfraspeciesElem are details for an infraspecific epithet of an
// Infraspecies name.
type InfraspeciesElem struct {
	// Value of an infraspecific epithet.
	Value string `json:"value"`
	// Rank of the infraspecific epithet.
	Rank string `json:"rank,omitempty"`
	// Authorship of the infraspecific epithet.
	Authorship *Authorship `json:"authorship,omitempty"`
}

// Comparison are details for a surrogate comparison name.
type Comparison struct {
	// Genus is the genus of a name.
	Genus string `json:"genus"`
	// Species is a specific epithet of a name.
	Species string `json:"species,omitempty"`
	// Cultivar is a value of a cultivar of a binomial.
	Cultivar string `json:"cultivar,omitempty"`
	// SpeciesAuthorship the authorship of Species.
	SpeciesAuthorship *Authorship `json:"authorship,omitempty"`
	// CompMarker, usually "cf.".
	CompMarker string `json:"comparisonMarker"`
}

// Approximation are details for a surrogate approximation name.
type Approximation struct {
	// Genus is the genus of a name.
	Genus string `json:"genus"`
	// Species is a specific epithet of a name.
	Species string `json:"species,omitempty"`
	// Cultivar is a value of a cultivar of a binomial.
	Cultivar string `json:"cultivar,omitempty"`
	// SpeciesAuthorship the authorship of Species.
	SpeciesAuthorship *Authorship `json:"authorship,omitempty"`
	// ApproxMarker describes what kind of approximation it is (sp., spp. etc.).
	ApproxMarker string `json:"approximationMarker,omitempty"`
	// Part of a name after ApproxMarker.
	Ignored string `json:"ignored,omitempty"`
}

// DetailsHybridFormula are details for a hybrid formula names.
type DetailsHybridFormula struct {
	HybridFormula []Details `json:"hybridFormula"`
}

// DetailsGraftChimeraFormula are details for a graft-chimera formula names.
type DetailsGraftChimeraFormula struct {
	GraftChimeraFormula []Details `json:"graftChimeraFormula"`
}

// isDetails implements Details interface.
func (DetailsHybridFormula) isDetails() {}


// isDetails implements Details interface.
func (DetailsGraftChimeraFormula) isDetails() {}

// DetailsUninomial are Uninomial details.
type DetailsUninomial struct {
	// Uninomial details.
	Uninomial Uninomial `json:"uninomial"`
}

// isDetails implements Details interface.
func (DetailsUninomial) isDetails() {}

// DetailsSpecies are binomial details.
type DetailsSpecies struct {
	// Species is details for binomial names.
	Species Species `json:"species"`
}

// isDetails implements Details interface.
func (DetailsSpecies) isDetails() {}

// DetailsInfraspecies are multinomial details.
type DetailsInfraspecies struct {
	// Infraspecies details.
	Infraspecies Infraspecies `json:"infraspecies"`
}

// isDetails implements Details interface.
func (DetailsInfraspecies) isDetails() {}

// DetailsComparison are details for comparison surrogate names.
type DetailsComparison struct {
	// Comparison details.
	Comparison Comparison `json:"comparison"`
}

// isDetails implements Details interface.
func (DetailsComparison) isDetails() {}

// DetailsApproximation are details for approximation surrogate names.
type DetailsApproximation struct {
	// Approximation details.
	Approximation Approximation `json:"approximation"`
}

// isDetails implements Details interface.
func (DetailsApproximation) isDetails() {}
