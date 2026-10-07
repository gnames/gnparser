# Global Names Parser Test With Zoological Code

<!-- TOC GFM -->

* [Introduction](#introduction)
* [Tests](#tests)
  * [Subspecies rank](#subspecies-rank)
  * [Infraspecific ranks other than subspecies](#infraspecific-ranks-other-than-subspecies)
  * [Uncommon ranks](#uncommon-ranks)
  * [More than one infraspecific epithet](#more-than-one-infraspecific-epithet)

<!-- /TOC -->

## Introduction

These tests run with the `-n zoo` (ICZN) setting. ICZN allows only one
infraspecific epithet, and its rank can only be subspecies.

## Tests

### Subspecies rank

<!-- These names follow ICZN rules and get no warnings. -->

Name: Aus bus Smith, 1890

Canonical: Aus bus

Authorship: Smith 1890

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":1,"verbatim":"Aus bus Smith, 1890","normalized":"Aus bus Smith 1890","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1890","normalized":"1890","wordType":"YEAR","start":15,"end":19}],"id":"5c759118-28e6-5192-b9bd-a1554f25de5b","parserVersion":"test_version"}
```

Name: Aus bus cus Smith, 1890

Canonical: Aus bus cus

Authorship: Smith 1890

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":1,"verbatim":"Aus bus cus Smith, 1890","normalized":"Aus bus cus Smith 1890","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus cus"},"cardinality":3,"authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":8,"end":11},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":12,"end":17},{"verbatim":"1890","normalized":"1890","wordType":"YEAR","start":19,"end":23}],"id":"e401b3ca-489a-5c02-9c82-a454678e876e","parserVersion":"test_version"}
```

Name: Aus bus subsp. cus Smith, 1890

Canonical: Aus bus subsp. cus

Authorship: Smith 1890

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":1,"verbatim":"Aus bus subsp. cus Smith, 1890","normalized":"Aus bus subsp. cus Smith 1890","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus subsp. cus"},"cardinality":3,"rank":"subsp.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"subsp.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":8,"end":14},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":15,"end":18},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"1890","normalized":"1890","wordType":"YEAR","start":26,"end":30}],"id":"3fcc41ac-d1c8-5bb9-a125-063d08c8d8a0","parserVersion":"test_version"}
```

Name: Aus bus ssp. cus

Canonical: Aus bus subsp. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":1,"verbatim":"Aus bus ssp. cus","normalized":"Aus bus subsp. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus subsp. cus"},"cardinality":3,"rank":"subsp.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"subsp."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"ssp.","normalized":"subsp.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16}],"id":"782e05a5-26ce-59b0-a2fe-16a11e7ece04","parserVersion":"test_version"}
```

### Infraspecific ranks other than subspecies

<!-- ICZN recognizes only subspecies below species rank (Art. 45.6). Names
published before 1961 as variety or form might be subspecific
(Art. 45.6.4), but canonical forms stay the same, only a warning is added. -->

Name: Aus bus var. cus Smith, 1890

Canonical: Aus bus var. cus

Authorship: Smith 1890

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus var. cus Smith, 1890","normalized":"Aus bus var. cus Smith 1890","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus var. cus"},"cardinality":3,"rank":"var.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"var.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":17,"end":22},{"verbatim":"1890","normalized":"1890","wordType":"YEAR","start":24,"end":28}],"id":"24481018-0278-5775-bfda-9f70847c4671","parserVersion":"test_version"}
```

Name: Aus bus variety cus

Canonical: Aus bus var. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus variety cus","normalized":"Aus bus var. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus var. cus"},"cardinality":3,"rank":"var.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"var."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"variety","normalized":"var.","wordType":"RANK","start":8,"end":15},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":16,"end":19}],"id":"b585efc5-8d60-5d39-889d-b26219c43357","parserVersion":"test_version"}
```

Name: Aus bus f. cus

Canonical: Aus bus f. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus f. cus","normalized":"Aus bus f. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus f. cus"},"cardinality":3,"rank":"f.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"f."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":8,"end":10},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":11,"end":14}],"id":"a642e78e-7a8f-503f-8ec1-2d0935dbfb3a","parserVersion":"test_version"}
```

Name: Aus bus forma cus

Canonical: Aus bus f. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus forma cus","normalized":"Aus bus f. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus f. cus"},"cardinality":3,"rank":"f.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"f."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"forma","normalized":"f.","wordType":"RANK","start":8,"end":13},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":14,"end":17}],"id":"1651395d-78d4-57af-a798-943dc08b4801","parserVersion":"test_version"}
```

Name: Aus bus ab. cus

Canonical: Aus bus ab. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus ab. cus","normalized":"Aus bus ab. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus ab. cus"},"cardinality":3,"rank":"ab.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"ab."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"ab.","normalized":"ab.","wordType":"RANK","start":8,"end":11},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":12,"end":15}],"id":"ea7111fd-50ef-5c62-ba24-98b056d6f4ca","parserVersion":"test_version"}
```

Name: Aus bus morph cus

Canonical: Aus bus morph cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus morph cus","normalized":"Aus bus morph cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus morph cus"},"cardinality":3,"rank":"morph","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"morph"}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"morph","normalized":"morph","wordType":"RANK","start":8,"end":13},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":14,"end":17}],"id":"b440d0eb-8a45-5734-a8d2-b4e6a80c89e6","parserVersion":"test_version"}
```

Name: Aus bus race cus

Canonical: Aus bus race cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus race cus","normalized":"Aus bus race cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus race cus"},"cardinality":3,"rank":"race","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"race"}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"race","normalized":"race","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16}],"id":"4db206e7-faae-5d32-b26b-920304b3bf4a","parserVersion":"test_version"}
```

Name: Aus bus agamosp. cus

Canonical: Aus bus agamosp. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus agamosp. cus","normalized":"Aus bus agamosp. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus agamosp. cus"},"cardinality":3,"rank":"agamosp.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"agamosp."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"agamosp.","normalized":"agamosp.","wordType":"RANK","start":8,"end":16},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":17,"end":20}],"id":"694a51e7-e34d-5478-bf4a-a02ead4d2ae8","parserVersion":"test_version"}
```

Name: Aus bus nothosubsp. cus

Canonical: Aus bus nothosubsp. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aus bus nothosubsp. cus","normalized":"Aus bus nothosubsp. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus nothosubsp. cus"},"cardinality":3,"rank":"nothosubsp.","hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"nothosubsp."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"nothosubsp.","normalized":"nothosubsp.","wordType":"RANK","start":8,"end":19},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":20,"end":23}],"id":"1b4b32a9-fc1d-585f-a042-993a5ae61a04","parserVersion":"test_version"}
```

Name: Aus bus var. β cus

Canonical: Aus bus var. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"},{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"}],"verbatim":"Aus bus var. β cus","normalized":"Aus bus var. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus var. cus"},"cardinality":3,"rank":"var.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"var."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":15,"end":18}],"id":"ab865be9-61b6-5cb0-8d3b-4115e2ab0852","parserVersion":"test_version"}
```

### Uncommon ranks

<!-- Uncommon ranks keep their own quality 3 warning. -->

Name: Aus bus natio cus

Canonical: Aus bus natio cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Aus bus natio cus","normalized":"Aus bus natio cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus natio cus"},"cardinality":3,"rank":"natio","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"natio"}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"natio","normalized":"natio","wordType":"RANK","start":8,"end":13},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":14,"end":17}],"id":"1b846287-3701-5344-9664-3d99f291c150","parserVersion":"test_version"}
```

Name: Aus bus mut. cus

Canonical: Aus bus mut. cus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Aus bus mut. cus","normalized":"Aus bus mut. cus","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus mut. cus"},"cardinality":3,"rank":"mut.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"mut."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"mut.","normalized":"mut.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16}],"id":"f70ab520-1611-5c9f-9f18-2d86ae998248","parserVersion":"test_version"}
```

### More than one infraspecific epithet

<!-- Subspecies name in ICZN is a trinomen (Art. 5.2), so names with more than
one infraspecific epithet get a warning. -->

Name: Aus bus cus dus

Canonical: Aus bus cus dus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"More than one infraspecific epithet (ICZN Art. 5.2)"}],"verbatim":"Aus bus cus dus","normalized":"Aus bus cus dus","canonical":{"stemmed":"Aus bus cus dus","simple":"Aus bus cus dus","full":"Aus bus cus dus"},"cardinality":4,"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus"},{"value":"dus"}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":8,"end":11},{"verbatim":"dus","normalized":"dus","wordType":"INFRASPECIES","start":12,"end":15}],"id":"8d69ecde-0554-567c-af0e-6f2dcb457355","parserVersion":"test_version"}
```

Name: Aus bus subsp. cus var. dus

Canonical: Aus bus subsp. cus var. dus

Authorship:

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"},{"quality":2,"warning":"More than one infraspecific epithet (ICZN Art. 5.2)"}],"verbatim":"Aus bus subsp. cus var. dus","normalized":"Aus bus subsp. cus var. dus","canonical":{"stemmed":"Aus bus cus dus","simple":"Aus bus cus dus","full":"Aus bus subsp. cus var. dus"},"cardinality":4,"rank":"var.","details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"subsp."},{"value":"dus","rank":"var."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":8,"end":14},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":15,"end":18},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":19,"end":23},{"verbatim":"dus","normalized":"dus","wordType":"INFRASPECIES","start":24,"end":27}],"id":"51957e9d-77dd-5c69-878f-4d34f29a4fd1","parserVersion":"test_version"}
```

Name: Aus bus var. cus f. dus Smith, 1890

Canonical: Aus bus var. cus f. dus

Authorship: Smith 1890

```json
{"parsed":true,"nomenclaturalCodeSetting":"ICZN","quality":2,"qualityWarnings":[{"quality":2,"warning":"Infraspecific rank other than subspecies (ICZN Art. 45.6)"},{"quality":2,"warning":"More than one infraspecific epithet (ICZN Art. 5.2)"}],"verbatim":"Aus bus var. cus f. dus Smith, 1890","normalized":"Aus bus var. cus f. dus Smith 1890","canonical":{"stemmed":"Aus bus cus dus","simple":"Aus bus cus dus","full":"Aus bus var. cus f. dus"},"cardinality":4,"rank":"f.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"var."},{"value":"dus","rank":"f.","authorship":{"verbatim":"Smith, 1890","normalized":"Smith 1890","year":"1890","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1890"}}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":17,"end":19},{"verbatim":"dus","normalized":"dus","wordType":"INFRASPECIES","start":20,"end":23},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":24,"end":29},{"verbatim":"1890","normalized":"1890","wordType":"YEAR","start":31,"end":35}],"id":"f344a0f2-ff65-544c-b6c7-9fa826d37006","parserVersion":"test_version"}
```
