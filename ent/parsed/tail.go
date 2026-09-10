package parsed

// Types of concept relations implied by concept-alignment annotations.
const (
	// RelationBroader means that the concept of the name is broader than
	// the reference concept (sensu lato).
	RelationBroader = "broader"

	// RelationNarrower means that the concept of the name is narrower than
	// the reference concept (sensu stricto).
	RelationNarrower = "narrower"

	// RelationSameAs means that the concept of the name is the concept of
	// the cited author (sensu <author>).
	RelationSameAs = "same_as"

	// RelationMisapplicationOf means that the name is used in a sense that
	// excludes the concept of the cited author (auct. non <author>).
	RelationMisapplicationOf = "misapplication_of"
)

// TailAnnotations contains annotations recognized in the tail of a
// name-string. It is set only if tail parsing is enabled (see
// gnparser.OptWithTail) and at least one annotation is recognized.
//
// Annotations are recognized from left to right. Recognition stops at the
// first unrecognized element, the rest of the tail stays in Parsed.Tail.
type TailAnnotations struct {
	// Verbatim is the tail before recognized annotations were removed
	// from it.
	Verbatim string `json:"verbatim"`

	// Sensu contains concept-alignment annotations: sensu lato,
	// sensu stricto, sensu <author>, sensu auct., auct.,
	// auct. non <author>, pro parte.
	Sensu []SensuAnnotation `json:"sensu,omitempty"`

	// Status contains nomenclatural status annotations, for example
	// nom. nud., nom. illeg., comb. nov.
	Status []StatusAnnotation `json:"status,omitempty"`

	// Publication contains publication and provenance modifiers:
	// ined., hort., fide <author>, emend. <author>, ex <author>.
	Publication []PublicationAnnotation `json:"publication,omitempty"`
}

// SensuAnnotation is a concept-alignment annotation.
type SensuAnnotation struct {
	// Verbatim is the annotation as it appears in the name-string, without
	// the cited author, for example "s. lat.".
	Verbatim string `json:"verbatim"`

	// Normalized is the normalized annotation without the cited author,
	// for example "sensu lato".
	Normalized string `json:"normalized"`

	// Author is the author cited by the annotation, for example "Smith, 1850"
	// in "sensu Smith, 1850". It is the author of the referenced concept,
	// and it is never copied to Parsed.Authorship.
	Author string `json:"author,omitempty"`

	// ConceptRelation is the relation implied by the annotation. It is nil
	// for annotations that do not imply a single relation (auct.,
	// sensu auct., pro parte).
	ConceptRelation *ConceptRelation `json:"conceptRelation,omitempty"`
}

// StatusAnnotation is a nomenclatural status annotation.
type StatusAnnotation struct {
	// Verbatim is the annotation as it appears in the name-string, for
	// example "nom. nud.".
	Verbatim string `json:"verbatim"`

	// Normalized is the normalized annotation, for example "nomen nudum".
	Normalized string `json:"normalized"`
}

// PublicationAnnotation is a publication or provenance modifier.
type PublicationAnnotation struct {
	// Verbatim is the annotation as it appears in the name-string, without
	// the cited author, for example "em.".
	Verbatim string `json:"verbatim"`

	// Normalized is the normalized annotation without the cited author,
	// for example "emend.".
	Normalized string `json:"normalized"`

	// Author is the author cited by the annotation, for example "Jones" in
	// "fide Jones".
	Author string `json:"author,omitempty"`
}

// ConceptRelation is an RCC5 relation implied by a concept-alignment
// annotation. It relates the concept of the name to a reference concept:
// the concept of the cited author, or an implicit concept, if no author
// is cited. It is a hint for concept alignment. Resolving the relation
// requires a comparator concept, which is outside of the name-string.
type ConceptRelation struct {
	// Type describes the relation. It is one of "broader", "narrower",
	// "same_as", "misapplication_of".
	Type string `json:"type"`

	// RCC5 is the RCC5 operator of the relation: ">", "<", "==", "|".
	RCC5 string `json:"rcc5"`

	// ReferenceAuthor is the author of the reference concept, if the
	// annotation cites one. It is empty when the reference concept is
	// implicit.
	ReferenceAuthor string `json:"referenceAuthor,omitempty"`
}
