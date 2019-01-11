package grammar

// WordType designates semantic meaning of a word.
type WordType int

const (
	UnknownType WordType = iota
	ComparisonType
	ApproxType
	AuthorWordType
	AuthorWordExType
	AuthorWordEmendType
	AuthorWordFiliusType
	GenusType
	InfraSpEpithetType
	HybridCharType
	RankType
	RankUniType
	SpEpithetType
	SubGenusType
	SuperSpType
	UninomialType
	YearApproximateType
	YearType
)
