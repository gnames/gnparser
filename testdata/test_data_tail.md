# Global Names Parser Test With Tail Parsing

<!-- markdownlint-disable -->

<!-- TOC GFM -->

* [Introduction](#introduction)
* [Tests](#tests)
  * [Sensu lato](#sensu-lato)
  * [Sensu stricto](#sensu-stricto)
  * [Sensu with an author](#sensu-with-an-author)
  * [Auct and misapplied names](#auct-and-misapplied-names)
  * [Pro parte](#pro-parte)
  * [Concept author without name author](#concept-author-without-name-author)
  * [Concept qualifiers without author](#concept-qualifiers-without-author)
  * [Nomenclatural status](#nomenclatural-status)
  * [Publication and provenance](#publication-and-provenance)
  * [Several annotations](#several-annotations)
  * [Other name types](#other-name-types)
  * [Partially recognized tails](#partially-recognized-tails)
  * [Unrecognized tails](#unrecognized-tails)

<!-- /TOC -->

## Introduction

These tests run with the -t/--tail flag enabled (`gnparser.OptWithTail(true)`)
and with details. Annotations recognized in the tail of a name-string
(concept-alignment annotations, nomenclatural status, publication modifiers)
are moved to `tailAnnotations`, and the quality is recalculated. The part of
the tail that is not recognized stays in `tail`.

Most names are fabricated (`Aus bus Smith, 1850`), so the tests do not depend
on changes in taxonomy.

## Tests

### Sensu lato

Name: Aus bus Smith, 1850 sensu lato

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 sensu lato","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" sensu lato","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"2ee58b7d-5167-585e-916a-03ebe81e8381","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s.l.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s.l.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.l.","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"1924c1f5-b5af-5b1c-9a3b-5faee0211c18","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s. l.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s. l.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s. l.","sensu":[{"verbatim":"s. l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"9f482b3e-4796-5c43-a4de-15ccb707aa8c","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s.lat.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s.lat.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.lat.","sensu":[{"verbatim":"s.lat.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"b3465c97-c167-5b5b-be83-d59ce37e3dd4","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s. lat.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s. lat.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s. lat.","sensu":[{"verbatim":"s. lat.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"49f99e68-f06b-5eae-8b76-4a1e7b045aa9","parserVersion":"test_version"}
```

Name: Quercus robur L. s.l.

Canonical: Quercus robur

Authorship: L.

```json
{"parsed":true,"quality":1,"verbatim":"Quercus robur L. s.l.","normalized":"Quercus robur L.","canonical":{"stemmed":"Quercus robur","simple":"Quercus robur","full":"Quercus robur"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tailAnnotations":{"verbatim":" s.l.","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Quercus","species":"robur","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"words":[{"verbatim":"Quercus","normalized":"Quercus","wordType":"GENUS","start":0,"end":7},{"verbatim":"robur","normalized":"robur","wordType":"SPECIES","start":8,"end":13},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":14,"end":16}],"id":"81cfd547-0d8c-50e4-aeb1-7372d81350cf","parserVersion":"test_version"}
```

### Sensu stricto

Name: Aus bus Smith, 1850 sensu stricto

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 sensu stricto","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" sensu stricto","sensu":[{"verbatim":"sensu stricto","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"10416a8e-7a13-5552-8918-a8cb3bf9afde","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s.s.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s.s.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.s.","sensu":[{"verbatim":"s.s.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"4fcc463c-b506-568f-8e1e-e5a3e5910628","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s. s.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s. s.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s. s.","sensu":[{"verbatim":"s. s.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"5a9489ea-587c-5255-9aee-e53897724399","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s.str.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s.str.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.str.","sensu":[{"verbatim":"s.str.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"24a14987-c18e-542f-8468-b257cb83fd54","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s. str.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 s. str.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s. str.","sensu":[{"verbatim":"s. str.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"433ed32f-510d-5f08-932a-a2ced6b831ae","parserVersion":"test_version"}
```

Name: Cus dus (Jones, 1900) (s.str.)

Canonical: Cus dus

Authorship: (Jones 1900)

```json
{"parsed":true,"quality":1,"verbatim":"Cus dus (Jones, 1900) (s.str.)","normalized":"Cus dus (Jones 1900)","canonical":{"stemmed":"Cus dus","simple":"Cus dus","full":"Cus dus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Jones, 1900)","normalized":"(Jones 1900)","year":"1900","authors":["Jones"],"originalAuth":{"authors":["Jones"],"year":{"year":"1900"}}},"tailAnnotations":{"verbatim":" (s.str.)","sensu":[{"verbatim":"s.str.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Cus","species":"dus","authorship":{"verbatim":"(Jones, 1900)","normalized":"(Jones 1900)","year":"1900","authors":["Jones"],"originalAuth":{"authors":["Jones"],"year":{"year":"1900"}}}}},"words":[{"verbatim":"Cus","normalized":"Cus","wordType":"GENUS","start":0,"end":3},{"verbatim":"dus","normalized":"dus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Jones","normalized":"Jones","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1900","normalized":"1900","wordType":"YEAR","start":16,"end":20}],"id":"00761316-8f3a-54b1-9cfc-1afe630c9dd4","parserVersion":"test_version"}
```

Name: Homo sapiens Linnaeus, 1758 sensu stricto

Canonical: Homo sapiens

Authorship: Linnaeus 1758

```json
{"parsed":true,"quality":1,"verbatim":"Homo sapiens Linnaeus, 1758 sensu stricto","normalized":"Homo sapiens Linnaeus 1758","canonical":{"stemmed":"Homo sapiens","simple":"Homo sapiens","full":"Homo sapiens"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Linnaeus, 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tailAnnotations":{"verbatim":" sensu stricto","sensu":[{"verbatim":"sensu stricto","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Homo","species":"sapiens","authorship":{"verbatim":"Linnaeus, 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Homo","normalized":"Homo","wordType":"GENUS","start":0,"end":4},{"verbatim":"sapiens","normalized":"sapiens","wordType":"SPECIES","start":5,"end":12},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":13,"end":21},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":23,"end":27}],"id":"4c1665bb-fdec-5ce2-a1cb-ddbb1077df06","parserVersion":"test_version"}
```

### Sensu with an author

Name: Aus bus Smith, 1850 sensu Jones, 1900

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 sensu Jones, 1900","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" sensu Jones, 1900","sensu":[{"verbatim":"sensu","normalized":"sensu","author":"Jones, 1900","conceptRelation":{"type":"same_as","rcc5":"==","referenceAuthor":"Jones, 1900"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"c231f640-b145-5d52-92ba-84ac60e7eaea","parserVersion":"test_version"}
```

Name: Cus dus (Jones, 1900) Brown, sensu White & Black, 1950

Canonical: Cus dus

Authorship: (Jones 1900) Brown

```json
{"parsed":true,"quality":1,"verbatim":"Cus dus (Jones, 1900) Brown, sensu White \u0026 Black, 1950","normalized":"Cus dus (Jones 1900) Brown","canonical":{"stemmed":"Cus dus","simple":"Cus dus","full":"Cus dus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Jones, 1900) Brown","normalized":"(Jones 1900) Brown","year":"1900","authors":["Jones","Brown"],"originalAuth":{"authors":["Jones"],"year":{"year":"1900"}},"combinationAuth":{"authors":["Brown"]}},"tailAnnotations":{"verbatim":", sensu White \u0026 Black, 1950","sensu":[{"verbatim":"sensu","normalized":"sensu","author":"White \u0026 Black, 1950","conceptRelation":{"type":"same_as","rcc5":"==","referenceAuthor":"White \u0026 Black, 1950"}}]},"details":{"species":{"genus":"Cus","species":"dus","authorship":{"verbatim":"(Jones, 1900) Brown","normalized":"(Jones 1900) Brown","year":"1900","authors":["Jones","Brown"],"originalAuth":{"authors":["Jones"],"year":{"year":"1900"}},"combinationAuth":{"authors":["Brown"]}}}},"words":[{"verbatim":"Cus","normalized":"Cus","wordType":"GENUS","start":0,"end":3},{"verbatim":"dus","normalized":"dus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Jones","normalized":"Jones","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1900","normalized":"1900","wordType":"YEAR","start":16,"end":20},{"verbatim":"Brown","normalized":"Brown","wordType":"AUTHOR_WORD","start":22,"end":27}],"id":"21fb30bc-b43e-56a4-95b0-f0b1e9fdb636","parserVersion":"test_version"}
```

### Auct and misapplied names

Name: Aus bus auct.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus auct.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" auct.","sensu":[{"verbatim":"auct.","normalized":"auct."}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"bca4c848-54e0-573f-ada9-31b436071691","parserVersion":"test_version"}
```

Name: Aus bus Auct.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Auct.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" Auct.","sensu":[{"verbatim":"Auct.","normalized":"auct."}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"a3d62f99-7554-5f5c-b078-d9845b219a9a","parserVersion":"test_version"}
```

Name: Aus bus sensu auct.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus sensu auct.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu auct.","sensu":[{"verbatim":"sensu auct.","normalized":"sensu auct."}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"4a30633d-76e0-513b-ab80-73876edd5424","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 auct. non Jones, 1900

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 auct. non Jones, 1900","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" auct. non Jones, 1900","sensu":[{"verbatim":"auct. non","normalized":"auct. non","author":"Jones, 1900","conceptRelation":{"type":"misapplication_of","rcc5":"|","referenceAuthor":"Jones, 1900"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"a570fee1-87c0-52e6-9087-0d424f58f7d6","parserVersion":"test_version"}
```

Name: Aus bus auct., non Smith, 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus auct., non Smith, 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" auct., non Smith, 1850","sensu":[{"verbatim":"auct., non","normalized":"auct. non","author":"Smith, 1850","conceptRelation":{"type":"misapplication_of","rcc5":"|","referenceAuthor":"Smith, 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"2e4d4426-797b-5333-bb35-74750b0e32bc","parserVersion":"test_version"}
```

Name: Aus bus sensu auct., non (Smith) Jones

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus sensu auct., non (Smith) Jones","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu auct., non (Smith) Jones","sensu":[{"verbatim":"sensu auct., non","normalized":"sensu auct. non","author":"(Smith) Jones","conceptRelation":{"type":"misapplication_of","rcc5":"|","referenceAuthor":"(Smith) Jones"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"f627fbb9-ba04-54c4-b1b2-b932e0fdc3d1","parserVersion":"test_version"}
```

### Pro parte

Name: Aus bus Smith, 1850 pro parte

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 pro parte","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" pro parte","sensu":[{"verbatim":"pro parte","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"32077c18-6ff7-5f65-9e3c-c129d1b3f5a8","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850, p.p.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850, p.p.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":", p.p.","sensu":[{"verbatim":"p.p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"de0d5f2a-08e8-54b9-9c08-3bdb3e0b88d3","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 p. p.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 p. p.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" p. p.","sensu":[{"verbatim":"p. p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"131be0d4-6aca-56ec-983f-0e5371cb3101","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 pro p.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 pro p.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" pro p.","sensu":[{"verbatim":"pro p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"56503df8-1dbf-56a0-839f-9211daabcd1e","parserVersion":"test_version"}
```

### Concept author without name author

Name: Aus bus sensu lato Smith, 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ambiguity: name author or concept author"}],"verbatim":"Aus bus sensu lato Smith, 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu lato Smith, 1850","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","author":"Smith, 1850","conceptRelation":{"type":"broader","rcc5":"\u003e","referenceAuthor":"Smith, 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"19281e9e-70e2-5aaa-b6c7-3b44bcefe628","parserVersion":"test_version"}
```

Name: Aus bus s. str. Smith, 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ambiguity: name author or concept author"}],"verbatim":"Aus bus s. str. Smith, 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" s. str. Smith, 1850","sensu":[{"verbatim":"s. str.","normalized":"sensu stricto","author":"Smith, 1850","conceptRelation":{"type":"narrower","rcc5":"\u003c","referenceAuthor":"Smith, 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"6023af21-8718-5402-8394-ee77dc988bc9","parserVersion":"test_version"}
```

Name: Aus bus s.l. (Smith) Jones

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ambiguity: name author or concept author"}],"verbatim":"Aus bus s.l. (Smith) Jones","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" s.l. (Smith) Jones","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","author":"(Smith) Jones","conceptRelation":{"type":"broader","rcc5":"\u003e","referenceAuthor":"(Smith) Jones"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"eefc9b92-1ba7-5bf4-8c5d-768fc9d6c744","parserVersion":"test_version"}
```

Name: Aus bus sensu Smith, 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus sensu Smith, 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu Smith, 1850","sensu":[{"verbatim":"sensu","normalized":"sensu","author":"Smith, 1850","conceptRelation":{"type":"same_as","rcc5":"==","referenceAuthor":"Smith, 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"e1bd779a-9660-5c74-97c9-1bb4b7d3a3e5","parserVersion":"test_version"}
```

Name: Aus bus sensu. Smith & Jones 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus sensu. Smith \u0026 Jones 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu. Smith \u0026 Jones 1850","sensu":[{"verbatim":"sensu.","normalized":"sensu","author":"Smith \u0026 Jones 1850","conceptRelation":{"type":"same_as","rcc5":"==","referenceAuthor":"Smith \u0026 Jones 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"562c4a87-c705-5b9b-88b5-2592f1f2a4ac","parserVersion":"test_version"}
```

Name: Aus bus auct. non Smith, 1850

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus auct. non Smith, 1850","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" auct. non Smith, 1850","sensu":[{"verbatim":"auct. non","normalized":"auct. non","author":"Smith, 1850","conceptRelation":{"type":"misapplication_of","rcc5":"|","referenceAuthor":"Smith, 1850"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"caa48e0c-b7a5-58d8-a786-e146564f5b2e","parserVersion":"test_version"}
```

### Concept qualifiers without author

Name: Aus bus sensu lato

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Concept qualifier without author"}],"verbatim":"Aus bus sensu lato","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu lato","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"8db5ea5d-6993-5c09-bb72-5312a03cdef3","parserVersion":"test_version"}
```

Name: Aus bus s.str.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Concept qualifier without author"}],"verbatim":"Aus bus s.str.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" s.str.","sensu":[{"verbatim":"s.str.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"aa567a4c-0f5a-55c1-aef0-c854ab7f2508","parserVersion":"test_version"}
```

Name: Aus bus p.p.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Concept qualifier without author"}],"verbatim":"Aus bus p.p.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" p.p.","sensu":[{"verbatim":"p.p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"5725f011-10de-5a56-b0b2-7614bdc8ce2b","parserVersion":"test_version"}
```

### Nomenclatural status

Name: Aus bus Smith, 1850 nom. dub.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. dub.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. dub.","status":[{"verbatim":"nom. dub.","normalized":"nomen dubium"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"b9fcd700-3fc6-542e-89ce-3b38f7a9e127","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen dubium

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen dubium","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen dubium","status":[{"verbatim":"nomen dubium","normalized":"nomen dubium"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"1b5bb0b9-23d0-5924-a234-cdfc856b7d9a","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nom. nud.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. nud.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. nud.","status":[{"verbatim":"nom. nud.","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"9843fda6-928d-5b2c-a0c9-600cd5b7d796","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen nudum

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen nudum","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen nudum","status":[{"verbatim":"nomen nudum","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"fd6f580f-57a3-54eb-ada0-be4cf8d16466","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850, nom. illeg.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850, nom. illeg.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":", nom. illeg.","status":[{"verbatim":"nom. illeg.","normalized":"nomen illegitimum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"afa4f10c-06db-54e2-ad0a-6e3a8acf8b80","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen illegitimum

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen illegitimum","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen illegitimum","status":[{"verbatim":"nomen illegitimum","normalized":"nomen illegitimum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"4df98360-01dc-5633-9f08-f49f4e1497bb","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nom. cons.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. cons.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. cons.","status":[{"verbatim":"nom. cons.","normalized":"nomen conservandum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"e16fa3b8-9ddf-5af6-af4d-516e062451e1","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen conservandum

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen conservandum","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen conservandum","status":[{"verbatim":"nomen conservandum","normalized":"nomen conservandum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"3321a908-a05d-5d16-b1c6-97f0d7650d5d","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nom. rej.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. rej.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. rej.","status":[{"verbatim":"nom. rej.","normalized":"nomen rejiciendum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"35b8617b-a0b0-5ed3-8898-3f16b82d977c","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen rejiciendum

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen rejiciendum","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen rejiciendum","status":[{"verbatim":"nomen rejiciendum","normalized":"nomen rejiciendum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"e3d1473a-71a3-57ae-933c-6779a9ec4502","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nom. prov.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. prov.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. prov.","status":[{"verbatim":"nom. prov.","normalized":"nomen provisorium"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"65c094d7-67fa-5aa1-a8ac-50f130223652","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nomen provisorium

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nomen provisorium","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nomen provisorium","status":[{"verbatim":"nomen provisorium","normalized":"nomen provisorium"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"8eabbecc-5f98-5d19-beb0-48c06d57ae08","parserVersion":"test_version"}
```

Name: Aus bus (Smith, 1850) stat. nov.

Canonical: Aus bus

Authorship: (Smith 1850)

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus (Smith, 1850) stat. nov.","normalized":"Aus bus (Smith 1850)","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" stat. nov.","status":[{"verbatim":"stat. nov.","normalized":"status novus"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":16,"end":20}],"id":"556d42a1-53be-5511-af11-488335cb17c5","parserVersion":"test_version"}
```

Name: Aus bus (Smith, 1850) comb. nov.

Canonical: Aus bus

Authorship: (Smith 1850)

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus (Smith, 1850) comb. nov.","normalized":"Aus bus (Smith 1850)","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" comb. nov.","status":[{"verbatim":"comb. nov.","normalized":"combinatio nova"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":16,"end":20}],"id":"23f647b8-663c-508f-b013-a206d82f9ed9","parserVersion":"test_version"}
```

Name: Aus bus (Smith, 1850) stat. rev.

Canonical: Aus bus

Authorship: (Smith 1850)

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus (Smith, 1850) stat. rev.","normalized":"Aus bus (Smith 1850)","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" stat. rev.","status":[{"verbatim":"stat. rev.","normalized":"status restitutus"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":16,"end":20}],"id":"1202ae09-cb39-54a2-b9cf-7b6d013c6429","parserVersion":"test_version"}
```

Name: Aus bus Smith [nom. nud.]

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith [nom. nud.]","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" [nom. nud.]","status":[{"verbatim":"nom. nud.","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"93559177-579e-53d4-8e81-47f552553240","parserVersion":"test_version"}
```

Name: Aus bus Smith (nom. nud.)

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith (nom. nud.)","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" (nom. nud.)","status":[{"verbatim":"nom. nud.","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"8d8e0baf-9ea7-5792-b713-0e3cf31fe72b","parserVersion":"test_version"}
```

### Publication and provenance

Name: Aus bus Smith ined.

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith ined.","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" ined.","publication":[{"verbatim":"ined.","normalized":"ineditus"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"c7098455-17a7-5d44-8deb-319c5dbcd0ff","parserVersion":"test_version"}
```

Name: Aus bus hort.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus hort.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" hort.","publication":[{"verbatim":"hort.","normalized":"hortorum"}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"00d14270-5b19-56d8-b886-1b1205ffbb3b","parserVersion":"test_version"}
```

Name: Aus bus hort. ex Smith

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus hort. ex Smith","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" hort. ex Smith","publication":[{"verbatim":"hort.","normalized":"hortorum"},{"verbatim":"ex","normalized":"ex","author":"Smith"}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"d5e5e5ff-c8de-5bf7-ba7e-4b01d3d37519","parserVersion":"test_version"}
```

Name: Aus bus Smith fide Jones, 1900

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith fide Jones, 1900","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" fide Jones, 1900","publication":[{"verbatim":"fide","normalized":"fide","author":"Jones, 1900"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"03ed9bf1-b24d-5f68-8bc9-505d816a2ad6","parserVersion":"test_version"}
```

Name: Aus bus Smith em. Jones

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith em. Jones","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" em. Jones","publication":[{"verbatim":"em.","normalized":"emend.","author":"Jones"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"4f1a6edd-2ec3-5512-b1f0-a044afad1bfb","parserVersion":"test_version"}
```

### Several annotations

Name: Aus bus Smith, 1850 sensu lato p.p.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 sensu lato p.p.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" sensu lato p.p.","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}},{"verbatim":"p.p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"6a4fbbe5-3546-561f-9c24-adbdc266b9cb","parserVersion":"test_version"}
```

Name: Aus bus s.l. p.p.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Concept qualifier without author"}],"verbatim":"Aus bus s.l. p.p.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" s.l. p.p.","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}},{"verbatim":"p.p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"41cbe221-a1c1-551f-bcc7-2354574180f1","parserVersion":"test_version"}
```

Name: Aus bus sensu lato Smith, 1850 p.p.

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ambiguity: name author or concept author"}],"verbatim":"Aus bus sensu lato Smith, 1850 p.p.","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tailAnnotations":{"verbatim":" sensu lato Smith, 1850 p.p.","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","author":"Smith, 1850","conceptRelation":{"type":"broader","rcc5":"\u003e","referenceAuthor":"Smith, 1850"}},{"verbatim":"p.p.","normalized":"pro parte"}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"f801fd5b-39f5-5dbb-8860-8804f3a4eec8","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 nom. nud. p.p.

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith, 1850 nom. nud. p.p.","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" nom. nud. p.p.","sensu":[{"verbatim":"p.p.","normalized":"pro parte"}],"status":[{"verbatim":"nom. nud.","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"40464ae2-5f4c-5a1f-9eaa-dd107b926c4e","parserVersion":"test_version"}
```

Name: Aus bus (Smith, 1850) comb. nov., nom. illeg.

Canonical: Aus bus

Authorship: (Smith 1850)

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus (Smith, 1850) comb. nov., nom. illeg.","normalized":"Aus bus (Smith 1850)","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" comb. nov., nom. illeg.","status":[{"verbatim":"comb. nov.","normalized":"combinatio nova"},{"verbatim":"nom. illeg.","normalized":"nomen illegitimum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":16,"end":20}],"id":"9207b837-f9d3-5d4d-a544-248e19fedecd","parserVersion":"test_version"}
```

Name: Aus bus (Smith, 1850) comb. nov., status novus

Canonical: Aus bus

Authorship: (Smith 1850)

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus (Smith, 1850) comb. nov., status novus","normalized":"Aus bus (Smith 1850)","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" comb. nov., status novus","status":[{"verbatim":"comb. nov.","normalized":"combinatio nova"},{"verbatim":"status novus","normalized":"status novus"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"(Smith, 1850)","normalized":"(Smith 1850)","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":9,"end":14},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":16,"end":20}],"id":"609a7c5b-1181-541a-a74f-27dfca912630","parserVersion":"test_version"}
```

Name: Aus bus Smith ined., fide Jones

Canonical: Aus bus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Smith ined., fide Jones","normalized":"Aus bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" ined., fide Jones","publication":[{"verbatim":"ined.","normalized":"ineditus"},{"verbatim":"fide","normalized":"fide","author":"Jones"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13}],"id":"a84e8798-ef17-5e7d-9abb-b601c1e08d9d","parserVersion":"test_version"}
```

### Other name types

Name: Aus Smith, 1850 s.str.

Canonical: Aus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus Smith, 1850 s.str.","normalized":"Aus Smith 1850","canonical":{"stemmed":"Aus","simple":"Aus","full":"Aus"},"cardinality":1,"authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.str.","sensu":[{"verbatim":"s.str.","normalized":"sensu stricto","conceptRelation":{"type":"narrower","rcc5":"\u003c"}}]},"details":{"uninomial":{"uninomial":"Aus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"UNINOMIAL","start":0,"end":3},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":4,"end":9},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":11,"end":15}],"id":"e4c39efa-e6d3-5f3a-b5fe-472b3e301ba3","parserVersion":"test_version"}
```

Name: Aus (Bus) cus Smith, 1850 s.l.

Canonical: Aus cus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus (Bus) cus Smith, 1850 s.l.","normalized":"Aus (Bus) cus Smith 1850","canonical":{"stemmed":"Aus cus","simple":"Aus cus","full":"Aus cus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.l.","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","subgenus":"Bus","species":"cus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"Bus","normalized":"Bus","wordType":"INFRA_GENUS","start":5,"end":8},{"verbatim":"cus","normalized":"cus","wordType":"SPECIES","start":10,"end":13},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":14,"end":19},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":21,"end":25}],"id":"2d381c38-1bb3-59e9-a4c3-9de3d1bbe79f","parserVersion":"test_version"}
```

Name: Aus bus cus Smith, 1850 s.l.

Canonical: Aus bus cus

Authorship: Smith 1850

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus cus Smith, 1850 s.l.","normalized":"Aus bus cus Smith 1850","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus cus"},"cardinality":3,"authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tailAnnotations":{"verbatim":" s.l.","sensu":[{"verbatim":"s.l.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":8,"end":11},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":12,"end":17},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":19,"end":23}],"id":"e524f6f6-a269-592d-b195-16874b72a2e7","parserVersion":"test_version"}
```

Name: Aus bus var. cus Smith sensu lato

Canonical: Aus bus var. cus

Authorship: Smith

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus var. cus Smith sensu lato","normalized":"Aus bus var. cus Smith","canonical":{"stemmed":"Aus bus cus","simple":"Aus bus cus","full":"Aus bus var. cus"},"cardinality":3,"rank":"var.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"tailAnnotations":{"verbatim":" sensu lato","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"infraspecies":{"genus":"Aus","species":"bus","infraspecies":[{"value":"cus","rank":"var.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":8,"end":12},{"verbatim":"cus","normalized":"cus","wordType":"INFRASPECIES","start":13,"end":16},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":17,"end":22}],"id":"18620dc2-4247-544f-b2b2-b796b33fc6a5","parserVersion":"test_version"}
```

Name: Aus × bus Smith nom. nud.

Canonical: Aus × bus

Authorship: Smith

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aus × bus Smith nom. nud.","normalized":"Aus × bus Smith","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus × bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}},"hybrid":"NAMED_HYBRID","tailAnnotations":{"verbatim":" nom. nud.","status":[{"verbatim":"nom. nud.","normalized":"nomen nudum"}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith","normalized":"Smith","authors":["Smith"],"originalAuth":{"authors":["Smith"]}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":4,"end":5},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":6,"end":9},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":10,"end":15}],"id":"805821bc-59d8-563f-a8ab-f834bdffd54d","parserVersion":"test_version"}
```

### Partially recognized tails

Name: Aus bus Smith, 1850 sensu lato foo

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus Smith, 1850 sensu lato foo","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tail":" foo","tailAnnotations":{"verbatim":" sensu lato foo","sensu":[{"verbatim":"sensu lato","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"cdf0e54e-1bc0-5a2e-9e77-48a1c08536fa","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 s.lat. - Aus bus

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus Smith, 1850 s.lat. - Aus bus","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tail":" - Aus bus","tailAnnotations":{"verbatim":" s.lat. - Aus bus","sensu":[{"verbatim":"s.lat.","normalized":"sensu lato","conceptRelation":{"type":"broader","rcc5":"\u003e"}}]},"details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"865ddf85-50d6-5155-9238-7a963c3b8fe5","parserVersion":"test_version"}
```

Name: Aus bus sensu Smith, non Jones

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus sensu Smith, non Jones","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tail":", non Jones","tailAnnotations":{"verbatim":" sensu Smith, non Jones","sensu":[{"verbatim":"sensu","normalized":"sensu","author":"Smith","conceptRelation":{"type":"same_as","rcc5":"==","referenceAuthor":"Smith"}}]},"details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"64df9ef5-0995-59fc-bc7c-596c6d8489ca","parserVersion":"test_version"}
```

### Unrecognized tails

Name: Aus bus Smith, 1850 non Jones, 1900

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus Smith, 1850 non Jones, 1900","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tail":" non Jones, 1900","details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"80518545-1663-5a04-8dc1-650ced5a30b4","parserVersion":"test_version"}
```

Name: Aus bus sensu

Canonical: Aus bus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus sensu","normalized":"Aus bus","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","tail":" sensu","details":{"species":{"genus":"Aus","species":"bus"}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7}],"id":"f122e07f-a470-596d-afb5-1fa26caf4fd9","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 p.p.B

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus Smith, 1850 p.p.B","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tail":" p.p.B","details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"ab6731ef-3dfe-50df-885a-8f694860b7a8","parserVersion":"test_version"}
```

Name: Aus bus Smith, 1850 ined.?

Canonical: Aus bus

Authorship: Smith 1850

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aus bus Smith, 1850 ined.?","normalized":"Aus bus Smith 1850","canonical":{"stemmed":"Aus bus","simple":"Aus bus","full":"Aus bus"},"cardinality":2,"rank":"sp.","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}},"tail":" ined.?","details":{"species":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Smith, 1850","normalized":"Smith 1850","year":"1850","authors":["Smith"],"originalAuth":{"authors":["Smith"],"year":{"year":"1850"}}}}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"1850","normalized":"1850","wordType":"YEAR","start":15,"end":19}],"id":"cb3b3968-5c32-5885-86a1-1cbd4b905837","parserVersion":"test_version"}
```
