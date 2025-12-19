package icvcn

type Rank int

const (
	UnknownRank Rank = iota
	Realm
	Subrealm
	Kingdom
	Subkingdom
	Phylum
	Subphylum
	Class
	Subclass
	Order
	Suborder
	Family
	Subfamily
	Genus
	Species
)

var RankToStr = map[Rank]string{
	Realm:      "Realm",
	Subrealm:   "Subrealm",
	Kingdom:    "Kingdom",
	Subkingdom: "Subkingdom",
	Phylum:     "Phylum",
	Subphylum:  "Subphylum",
	Class:      "Class",
	Subclass:   "Subclass",
	Order:      "Order",
	Suborder:   "Suborder",
	Family:     "Family",
	Subfamily:  "Subfamily",
	Genus:      "Genus",
	Species:    "Species",
}

func (r Rank) String() string {
	return RankToStr[r]
}
