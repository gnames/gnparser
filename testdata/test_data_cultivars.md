# Global Names Parser Test With Cultivars

<!-- TOC GFM -->

* [Introduction](#introduction)
* [Tests](#tests)
  * [Binomials with cultivars](#binomials-with-cultivars)
  * [Names with cultivars in apostrophes](#names-with-cultivars-in-apostrophes)
  * [Names with cultivars in single quotes](#names-with-cultivars-in-single-quotes)
  * [Names with cultivars in double straight quotes](#names-with-cultivars-in-double-straight-quotes)
  * [Hybrid formulae with cultivars](#hybrid-formulae-with-cultivars)
  * [Uninomials with cultivars](#uninomials-with-cultivars)
  * [Graft-chimeras](#graft-chimeras)

<!-- /TOC -->

## Introduction

These tests run with the -C/--cultivar flag enabled, which adds the cultivar
epithet into the normalized and canonical names and enables the parsing of graft-chimeras

## Tests

### Binomials with cultivars

Name: Sarracenia flava 'Maxima'

Canonical: Sarracenia flava ‘Maxima’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Sarracenia flava 'Maxima'","normalized":"Sarracenia flava ‘Maxima’","canonical":{"stemmed":"Sarracenia flau ‘Maxima’","simple":"Sarracenia flava ‘Maxima’","full":"Sarracenia flava ‘Maxima’"},"cardinality":3,"cultivar":true,"details":{"species":{"genus":"Sarracenia","species":"flava","cultivar":"‘Maxima’"}},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"flava","normalized":"flava","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Maxima","normalized":"‘Maxima’","wordType":"CULTIVAR","start":18,"end":24}],"id":"39178008-65ee-5de3-af88-63ffdd67e00b","parserVersion":"test_version"}
```

Name: Phyllostachys vivax cv aureocaulis

Canonical: Phyllostachys vivax ‘aureocaulis’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Phyllostachys vivax cv aureocaulis","normalized":"Phyllostachys vivax ‘aureocaulis’","canonical":{"stemmed":"Phyllostachys uiuax ‘aureocaulis’","simple":"Phyllostachys vivax ‘aureocaulis’","full":"Phyllostachys vivax ‘aureocaulis’"},"cardinality":3,"cultivar":true,"details":{"species":{"genus":"Phyllostachys","species":"vivax","cultivar":"‘aureocaulis’"}},"words":[{"verbatim":"Phyllostachys","normalized":"Phyllostachys","wordType":"GENUS","start":0,"end":13},{"verbatim":"vivax","normalized":"vivax","wordType":"SPECIES","start":14,"end":19},{"verbatim":"aureocaulis","normalized":"‘aureocaulis’","wordType":"CULTIVAR","start":23,"end":34}],"id":"56f7057d-9c5c-5ac7-bc7a-f631fb58f5d6","parserVersion":"test_version"}
```

Name: Ligusticum sinense cv 'chuanxiong' S.H. Qiu & et al.

Canonical: Ligusticum sinense ‘chuanxiong’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ligusticum sinense cv 'chuanxiong' S.H. Qiu \u0026 et al.","normalized":"Ligusticum sinense ‘chuanxiong’","canonical":{"stemmed":"Ligusticum sinens ‘chuanxiong’","simple":"Ligusticum sinense ‘chuanxiong’","full":"Ligusticum sinense ‘chuanxiong’"},"cardinality":3,"cultivar":true,"tail":" S.H. Qiu \u0026 et al.","details":{"species":{"genus":"Ligusticum","species":"sinense","cultivar":"‘chuanxiong’"}},"words":[{"verbatim":"Ligusticum","normalized":"Ligusticum","wordType":"GENUS","start":0,"end":10},{"verbatim":"sinense","normalized":"sinense","wordType":"SPECIES","start":11,"end":18},{"verbatim":"chuanxiong","normalized":"‘chuanxiong’","wordType":"CULTIVAR","start":23,"end":33}],"id":"73f015c2-6679-5428-b418-6f4487af419d","parserVersion":"test_version"}
```

Name: Anthurium 'Ace of Spades'

Canonical: Anthurium ‘Ace of Spades’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Anthurium 'Ace of Spades'","normalized":"Anthurium ‘Ace of Spades’","canonical":{"stemmed":"Anthurium ‘Ace of Spades’","simple":"Anthurium ‘Ace of Spades’","full":"Anthurium ‘Ace of Spades’"},"cardinality":2,"cultivar":true,"details":{"uninomial":{"uninomial":"Anthurium","cultivar":"‘Ace of Spades’"}},"words":[{"verbatim":"Anthurium","normalized":"Anthurium","wordType":"UNINOMIAL","start":0,"end":9},{"verbatim":"Ace of Spades","normalized":"‘Ace of Spades’","wordType":"CULTIVAR","start":11,"end":24}],"id":"3adaf031-08f2-576e-b9af-616bf328473e","parserVersion":"test_version"}
```

### Names with cultivars in apostrophes

Name: Sarracenia flava 'Maxima'

Canonical: Sarracenia flava ‘Maxima’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Sarracenia flava 'Maxima'","normalized":"Sarracenia flava ‘Maxima’","canonical":{"stemmed":"Sarracenia flau ‘Maxima’","simple":"Sarracenia flava ‘Maxima’","full":"Sarracenia flava ‘Maxima’"},"cardinality":3,"cultivar":true,"details":{"species":{"genus":"Sarracenia","species":"flava","cultivar":"‘Maxima’"}},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"flava","normalized":"flava","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Maxima","normalized":"‘Maxima’","wordType":"CULTIVAR","start":18,"end":24}],"id":"39178008-65ee-5de3-af88-63ffdd67e00b","parserVersion":"test_version"}
```

### Names with cultivars in single quotes

Name: Colocasia esculenta ‘Black Magic’

Canonical: Colocasia esculenta ‘Black Magic’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Colocasia esculenta ‘Black Magic’","normalized":"Colocasia esculenta ‘Black Magic’","canonical":{"stemmed":"Colocasia esculent ‘Black Magic’","simple":"Colocasia esculenta ‘Black Magic’","full":"Colocasia esculenta ‘Black Magic’"},"cardinality":3,"cultivar":true,"details":{"species":{"genus":"Colocasia","species":"esculenta","cultivar":"‘Black Magic’"}},"words":[{"verbatim":"Colocasia","normalized":"Colocasia","wordType":"GENUS","start":0,"end":9},{"verbatim":"esculenta","normalized":"esculenta","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Black Magic","normalized":"‘Black Magic’","wordType":"CULTIVAR","start":21,"end":32}],"id":"9a74485c-86d2-5bc6-a796-a634bdf03a9e","parserVersion":"test_version"}
```

### Names with cultivars in double straight quotes

Name: Amorphophallus konjac "Nightstick"

Canonical: Amorphophallus konjac ‘Nightstick’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Amorphophallus konjac \"Nightstick\"","normalized":"Amorphophallus konjac ‘Nightstick’","canonical":{"stemmed":"Amorphophallus koniac ‘Nightstick’","simple":"Amorphophallus konjac ‘Nightstick’","full":"Amorphophallus konjac ‘Nightstick’"},"cardinality":3,"cultivar":true,"details":{"species":{"genus":"Amorphophallus","species":"konjac","cultivar":"‘Nightstick’"}},"words":[{"verbatim":"Amorphophallus","normalized":"Amorphophallus","wordType":"GENUS","start":0,"end":14},{"verbatim":"konjac","normalized":"konjac","wordType":"SPECIES","start":15,"end":21},{"verbatim":"Nightstick","normalized":"‘Nightstick’","wordType":"CULTIVAR","start":23,"end":33}],"id":"eaa0c523-412c-55d5-a3fb-88c9476362c7","parserVersion":"test_version"}
```

### Hybrid formulae with cultivars

Name: Sarracenia alata 'Black Tube' x Sarracenia leucophylla

Canonical: Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Sarracenia alata 'Black Tube' x Sarracenia leucophylla","normalized":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla","canonical":{"stemmed":"Sarracenia alat ‘Black Tube’ × Sarracenia leucophyll","simple":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla","full":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla"},"cardinality":0,"cultivar":true,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Sarracenia","species":"alata","cultivar":"‘Black Tube’"}},{"species":{"genus":"Sarracenia","species":"leucophylla"}}]},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"alata","normalized":"alata","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Black Tube","normalized":"‘Black Tube’","wordType":"CULTIVAR","start":18,"end":28},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":30,"end":31},{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":32,"end":42},{"verbatim":"leucophylla","normalized":"leucophylla","wordType":"SPECIES","start":43,"end":54}],"id":"17b9d0fb-76f0-510e-8c13-bf48033d50dd","parserVersion":"test_version"}
```

Name: Sarracenia alata cv Black Tube x Sarracenia leucophylla

Canonical: Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Sarracenia alata cv Black Tube x Sarracenia leucophylla","normalized":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla","canonical":{"stemmed":"Sarracenia alat ‘Black Tube’ × Sarracenia leucophyll","simple":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla","full":"Sarracenia alata ‘Black Tube’ × Sarracenia leucophylla"},"cardinality":0,"cultivar":true,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Sarracenia","species":"alata","cultivar":"‘Black Tube’"}},{"species":{"genus":"Sarracenia","species":"leucophylla"}}]},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"alata","normalized":"alata","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Black Tube","normalized":"‘Black Tube’","wordType":"CULTIVAR","start":20,"end":30},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":31,"end":32},{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":33,"end":43},{"verbatim":"leucophylla","normalized":"leucophylla","wordType":"SPECIES","start":44,"end":55}],"id":"1b978ba7-efc5-550f-a598-7830114514b1","parserVersion":"test_version"}
```

Name: Sarracenia alata cv Black Tube x Sarracenia flava 'Copper Lid'

Canonical: Sarracenia alata ‘Black Tube’ × Sarracenia flava ‘Copper Lid’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Sarracenia alata cv Black Tube x Sarracenia flava 'Copper Lid'","normalized":"Sarracenia alata ‘Black Tube’ × Sarracenia flava ‘Copper Lid’","canonical":{"stemmed":"Sarracenia alat ‘Black Tube’ × Sarracenia flau ‘Copper Lid’","simple":"Sarracenia alata ‘Black Tube’ × Sarracenia flava ‘Copper Lid’","full":"Sarracenia alata ‘Black Tube’ × Sarracenia flava ‘Copper Lid’"},"cardinality":0,"cultivar":true,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Sarracenia","species":"alata","cultivar":"‘Black Tube’"}},{"species":{"genus":"Sarracenia","species":"flava","cultivar":"‘Copper Lid’"}}]},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"alata","normalized":"alata","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Black Tube","normalized":"‘Black Tube’","wordType":"CULTIVAR","start":20,"end":30},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":31,"end":32},{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":33,"end":43},{"verbatim":"flava","normalized":"flava","wordType":"SPECIES","start":44,"end":49},{"verbatim":"Copper Lid","normalized":"‘Copper Lid’","wordType":"CULTIVAR","start":51,"end":61}],"id":"260dea27-b2c9-5231-bebf-b149999e053a","parserVersion":"test_version"}
```

### Uninomials with cultivars

Name: Rhododendron cv Cilpinense

Canonical: Rhododendron ‘Cilpinense’

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Rhododendron cv Cilpinense","normalized":"Rhododendron ‘Cilpinense’","canonical":{"stemmed":"Rhododendron ‘Cilpinense’","simple":"Rhododendron ‘Cilpinense’","full":"Rhododendron ‘Cilpinense’"},"cardinality":2,"cultivar":true,"details":{"uninomial":{"uninomial":"Rhododendron","cultivar":"‘Cilpinense’"}},"words":[{"verbatim":"Rhododendron","normalized":"Rhododendron","wordType":"UNINOMIAL","start":0,"end":12},{"verbatim":"Cilpinense","normalized":"‘Cilpinense’","wordType":"CULTIVAR","start":16,"end":26}],"id":"abd299df-e4b2-533c-86eb-a4a5e273b934","parserVersion":"test_version"}
```

Name: Spathiphyllum Schott “Mauna Loa”

Canonical: Spathiphyllum ‘Mauna Loa’

Authorship: Schott

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":1,"verbatim":"Spathiphyllum Schott “Mauna Loa”","normalized":"Spathiphyllum Schott ‘Mauna Loa’","canonical":{"stemmed":"Spathiphyllum ‘Mauna Loa’","simple":"Spathiphyllum ‘Mauna Loa’","full":"Spathiphyllum ‘Mauna Loa’"},"cardinality":2,"authorship":{"verbatim":"Schott","normalized":"Schott","authors":["Schott"],"originalAuth":{"authors":["Schott"]}},"cultivar":true,"details":{"uninomial":{"uninomial":"Spathiphyllum","cultivar":"‘Mauna Loa’","authorship":{"verbatim":"Schott","normalized":"Schott","authors":["Schott"],"originalAuth":{"authors":["Schott"]}}}},"words":[{"verbatim":"Spathiphyllum","normalized":"Spathiphyllum","wordType":"UNINOMIAL","start":0,"end":13},{"verbatim":"Schott","normalized":"Schott","wordType":"AUTHOR_WORD","start":14,"end":20},{"verbatim":"Mauna Loa","normalized":"‘Mauna Loa’","wordType":"CULTIVAR","start":22,"end":31}],"id":"fb8afb5b-67b8-5bcc-8492-773cc40d3bb9","parserVersion":"test_version"}
```

### Graft-chimeras

Name: + Crataegomespilus

Canonical: + Crataegomespilus

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Named graft-chimera"}],"verbatim":"+ Crataegomespilus","normalized":"+ Crataegomespilus","canonical":{"stemmed":"Crataegomespilus","simple":"Crataegomespilus","full":"+ Crataegomespilus"},"cardinality":1,"hybrid":"NAMED_GRAFT_CHIMERA","details":{"uninomial":{"uninomial":"Crataegomespilus"}},"words":[{"verbatim":"+","normalized":"+","wordType":"GRAFT_CHIMERA_CHAR","start":0,"end":1},{"verbatim":"Crataegomespilus","normalized":"Crataegomespilus","wordType":"UNINOMIAL","start":2,"end":18}],"id":"408e8fc7-fa27-53a6-9eff-37cb779724e4","parserVersion":"test_version"}
```

Name: +Crataegomespilus

Canonical: + Crataegomespilus

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":3,"qualityWarnings":[{"quality":3,"warning":"Graft-chimera char is not separated by space"},{"quality":2,"warning":"Named graft-chimera"}],"verbatim":"+Crataegomespilus","normalized":"+ Crataegomespilus","canonical":{"stemmed":"Crataegomespilus","simple":"Crataegomespilus","full":"+ Crataegomespilus"},"cardinality":1,"hybrid":"NAMED_GRAFT_CHIMERA","details":{"uninomial":{"uninomial":"Crataegomespilus"}},"words":[{"verbatim":"+","normalized":"+","wordType":"GRAFT_CHIMERA_CHAR","start":0,"end":1},{"verbatim":"Crataegomespilus","normalized":"Crataegomespilus","wordType":"UNINOMIAL","start":1,"end":17}],"id":"c2c50c08-1f62-547f-8fab-50359caf0b31","parserVersion":"test_version"}
```

Name: Cytisus purpureus + Laburnum anagyroides

Canonical: Cytisus purpureus + Laburnum anagyroides

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Graft-chimera formula"}],"verbatim":"Cytisus purpureus + Laburnum anagyroides","normalized":"Cytisus purpureus + Laburnum anagyroides","canonical":{"stemmed":"Cytisus purpure + Laburnum anagyroid","simple":"Cytisus purpureus + Laburnum anagyroides","full":"Cytisus purpureus + Laburnum anagyroides"},"cardinality":0,"hybrid":"GRAFT_CHIMERA_FORMULA","details":{"graftChimeraFormula":[{"species":{"genus":"Cytisus","species":"purpureus"}},{"species":{"genus":"Laburnum","species":"anagyroides"}}]},"words":[{"verbatim":"Cytisus","normalized":"Cytisus","wordType":"GENUS","start":0,"end":7},{"verbatim":"purpureus","normalized":"purpureus","wordType":"SPECIES","start":8,"end":17},{"verbatim":"+","normalized":"+","wordType":"GRAFT_CHIMERA_CHAR","start":18,"end":19},{"verbatim":"Laburnum","normalized":"Laburnum","wordType":"GENUS","start":20,"end":28},{"verbatim":"anagyroides","normalized":"anagyroides","wordType":"SPECIES","start":29,"end":40}],"id":"a8f8ace8-ba1a-5371-b9d5-73efce81d52c","parserVersion":"test_version"}
```

Name: Crataegus + Mespilus

Canonical: Crataegus + Mespilus

Authorship:

```json
{"parsed":true,"nomenclaturalCode":"ICNCP","quality":2,"qualityWarnings":[{"quality":2,"warning":"Graft-chimera formula"}],"verbatim":"Crataegus + Mespilus","normalized":"Crataegus + Mespilus","canonical":{"stemmed":"Crataegus + Mespilus","simple":"Crataegus + Mespilus","full":"Crataegus + Mespilus"},"cardinality":0,"hybrid":"GRAFT_CHIMERA_FORMULA","details":{"graftChimeraFormula":[{"uninomial":{"uninomial":"Crataegus"}},{"uninomial":{"uninomial":"Mespilus"}}]},"words":[{"verbatim":"Crataegus","normalized":"Crataegus","wordType":"UNINOMIAL","start":0,"end":9},{"verbatim":"+","normalized":"+","wordType":"GRAFT_CHIMERA_CHAR","start":10,"end":11},{"verbatim":"Mespilus","normalized":"Mespilus","wordType":"UNINOMIAL","start":12,"end":20}],"id":"d651cd82-9b00-53dd-9d59-6af66ab62046","parserVersion":"test_version"}
```
