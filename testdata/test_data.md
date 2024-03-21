# Global Names Parser Test

<!-- markdownlint-disable -->

<!-- vim-markdown-toc GFM -->

* [Introduction](#introduction)
* [Tests](#tests)
  * [Uninomials without authorship](#uninomials-without-authorship)
  * [Uninomials with authorship](#uninomials-with-authorship)
  * [Two-letter genus names (legacy genera, not allowed anymore)](#two-letter-genus-names-legacy-genera-not-allowed-anymore)
  * [Combination of two uninomials](#combination-of-two-uninomials)
  * [ICN names that look like combined uninomials for ICZN](#icn-names-that-look-like-combined-uninomials-for-iczn)
  * [Binomials without authorship](#binomials-without-authorship)
  * [Binomials with authorship](#binomials-with-authorship)
  * [Binomials with an abbreviated genus](#binomials-with-an-abbreviated-genus)
  * [Binomials with abbreviated subgenus](#binomials-with-abbreviated-subgenus)
  * [Binomials with basionym and combination authors](#binomials-with-basionym-and-combination-authors)
  * [Exceptions with Binomials](#exceptions-with-binomials)
  * [Binomials with Mc and Mac authors](#binomials-with-mc-and-mac-authors)
  * [Infraspecies without rank (ICZN)](#infraspecies-without-rank-iczn)
  * [Legacy ICZN names with rank](#legacy-iczn-names-with-rank)
  * [Infraspecies with rank (ICN)](#infraspecies-with-rank-icn)
  * [Infraspecies multiple (ICN)](#infraspecies-multiple-icn)
  * [Infraspecies with greek letters (ICN)](#infraspecies-with-greek-letters-icn)
  * [Names with the dagger char '†'](#names-with-the-dagger-char-)
  * [Hybrids with notho- ranks](#hybrids-with-notho--ranks)
  * [Named hybrids](#named-hybrids)
  * [Hybrid formulae](#hybrid-formulae)
  * [Graft-chimeras](#graft-chimeras)
  * [Genus with hyphen (allowed by ICN)](#genus-with-hyphen-allowed-by-icn)
  * [Misspeled name](#misspeled-name)
  * [A 'basionym' author in parenthesis (basionym is an ICN term)](#a-basionym-author-in-parenthesis-basionym-is-an-icn-term)
  * [Infrageneric epithets (ICZN)](#infrageneric-epithets-iczn)
  * [Names with multiple dashes in specific epithet](#names-with-multiple-dashes-in-specific-epithet)
  * [Genus with question mark](#genus-with-question-mark)
  * [Epithets with a period character](#epithets-with-a-period-character)
  * [Epithets starting with non-](#epithets-starting-with-non-)
  * [Epithets starting with authors' prefixes (de, di, la, von etc.)](#epithets-starting-with-authors-prefixes-de-di-la-von-etc)
  * [Authorship missing one parenthesis](#authorship-missing-one-parenthesis)
  * [Unknown authorship](#unknown-authorship)
  * [Treating apud (with)](#treating-apud-with)
  * [Names with ex authors (we follow ICZN convention)](#names-with-ex-authors-we-follow-iczn-convention)
  * [Empty spaces](#empty-spaces)
  * [Names with a dash](#names-with-a-dash)
  * [Authorship with 'degli'](#authorship-with-degli)
  * [Authorship with filius (son of)](#authorship-with-filius-son-of)
  * [Names with emend (rectified by) authorship](#names-with-emend-rectified-by-authorship)
  * [Names with an unparsed "tail"](#names-with-an-unparsed-tail)
  * [Abbreviated words after a name](#abbreviated-words-after-a-name)
  * [Epithets starting with numeric value (not allowed anymore)](#epithets-starting-with-numeric-value-not-allowed-anymore)
  * [Non-ASCII UTF-8 characters in a name](#non-ascii-utf-8-characters-in-a-name)
  * [Epithets with an apostrophe](#epithets-with-an-apostrophe)
  * [Authors with an apostrophe](#authors-with-an-apostrophe)
  * [Digraph unicode characters](#digraph-unicode-characters)
  * [Old style s (ſ)](#old-style-s-)
  * [Miscellaneous diacritics](#miscellaneous-diacritics)
  * [Open Nomenclature ('approximate' names)](#open-nomenclature-approximate-names)
  * [Surrogate Name-Strings](#surrogate-name-strings)
  * [Virus-like "normal" names](#virus-like-normal-names)
  * [Viruses, plasmids, prions etc.](#viruses-plasmids-prions-etc)
  * [Name-strings with RNA](#name-strings-with-rna)
  * [Epithet prioni is not a prion](#epithet-prioni-is-not-a-prion)
  * [Names with "satellite" as a substring](#names-with-satellite-as-a-substring)
  * [Bacterial genus](#bacterial-genus)
  * [Bacteria genus homonym](#bacteria-genus-homonym)
  * [Bacteria with pathovar rank](#bacteria-with-pathovar-rank)
  * ["Stray" ex is not parsed as species](#stray-ex-is-not-parsed-as-species)
  * [Authorship in upper case](#authorship-in-upper-case)
  * [Numbers and letters separated with '-' are not parsed as authors](#numbers-and-letters-separated-with---are-not-parsed-as-authors)
  * [Double parenthesis](#double-parenthesis)
  * [Numbers at the start/middle of names](#numbers-at-the-startmiddle-of-names)
  * [Year without authorship](#year-without-authorship)
  * [Year range](#year-range)
  * [Year with page number](#year-with-page-number)
  * [Year in square brackets](#year-in-square-brackets)
  * [Names with broken conversion between encodings](#names-with-broken-conversion-between-encodings)
  * [UTF-8 0xA0 character (NO_BREAK_SPACE)](#utf-8-0xa0-character-no_break_space)
  * [UTF-8 0x3000 character (IDEOGRAPHIC_SPACE)](#utf-8-0x3000-character-ideographic_space)
  * [Punctuation in the end](#punctuation-in-the-end)
  * [Names with 'ex' as sp. epithet](#names-with-ex-as-sp-epithet)
  * [Names with Spanish 'y' instead of '&'](#names-with-spanish-y-instead-of-)
  * [Normalize atypical dashes](#normalize-atypical-dashes)
  * [Discard apostrophes at the start and end of words](#discard-apostrophes-at-the-start-and-end-of-words)
  * [Discard apostrophe with dash (rare, needs further investigation)](#discard-apostrophe-with-dash-rare-needs-further-investigation)
  * [Possible canonical](#possible-canonical)
  * [Treating `& al.` as `et al.`](#treating--al-as-et-al)
  * [Authors do not start with apostrophe](#authors-do-not-start-with-apostrophe)
  * [Epithets do not start or end with a dash](#epithets-do-not-start-or-end-with-a-dash)
  * [Names that contain "of"](#names-that-contain-of)
  * [Cultivars](#cultivars)
  * ["Open taxonomy" with ranks unfinished](#open-taxonomy-with-ranks-unfinished)
  * [Ignoring serovar/serotype](#ignoring-serovarserotype)
  * [Ignoring sensu sec](#ignoring-sensu-sec)
  * [Unparseable hort. annotations](#unparseable-hort-annotations)
  * [Removing nomenclatural annotations](#removing-nomenclatural-annotations)
  * [Misc annotations](#misc-annotations)
  * [Horticultural annotation](#horticultural-annotation)
  * [Names with "mihi"](#names-with-mihi)
  * [Exceptions with "mihi"](#exceptions-with-mihi)
  * [Exceptions from ranks (rank-line epithets)](#exceptions-from-ranks-rank-line-epithets)
  * [Exceptions from author prefixes (prefix-like epithets)](#exceptions-from-author-prefixes-prefix-like-epithets)
  * [Exceptions from author suffixes (suffix-like epithets)](#exceptions-from-author-suffixes-suffix-like-epithets)
  * [Not parsed OCR errors to get better precision/recall ratio](#not-parsed-ocr-errors-to-get-better-precisionrecall-ratio)
  * [No parsing -- Genera abbreviated to 3 letters (too rare)](#no-parsing----genera-abbreviated-to-3-letters-too-rare)
  * [No parsing -- incertae sedis](#no-parsing----incertae-sedis)
  * [No parsing -- bacterium, Candidatus](#no-parsing----bacterium-candidatus)
  * [No parsing -- 'Not', 'None', 'Unidentified'  phrases](#no-parsing----not-none-unidentified--phrases)
  * [No parsing -- genus with apostrophe](#no-parsing----genus-with-apostrophe)
  * [No parsing -- CamelCase 'genus' word](#no-parsing----camelcase-genus-word)
  * [No parsing -- phytoplasma](#no-parsing----phytoplasma)
  * [No parsing symbiont](#no-parsing-symbiont)
  * [Names with spec., nov spec](#names-with-spec-nov-spec)
  * [HTML tags and entities](#html-tags-and-entities)
  * [Underscores instead of spaces](#underscores-instead-of-spaces)

<!-- vim-markdown-toc -->

## Introduction

This test consists of a line-delimited input (scientific name), detailed
parsed output in JSON format and simplified parsed output in
pipe-delimited format

Test Structure

The first line in every test is a scientific name to parse
The second line corresponds to detailed JSON output from the gnparser
The third line corresponds to pipe-delimited "simple" output. Simple output
consists of name-string UUID v5, verbatim name-string, canonical form without
ranks, canonical form with ranks, authorship of the most junior clade, year,
parsing quality number

[Parsing quality](https://github.com/gnames/gnparser/quality.md)

1: parsed without problems

2: parsed with minor problems,

3: parsed with significant problems

4: parsed with severe problems

0: parsing failed

## Tests

### Uninomials without authorship

Name: Pseudocercospora

Canonical: Pseudocercospora

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pseudocercospora"}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"UNINOMIAL","start":0,"end":16}],"id":"9c1167ca-79e7-53de-b4c3-fcdb68410527","parserVersion":"test_version"}
```

### Uninomials with authorship

Name: Tremoctopus violaceus delle Chiaje, 1830

Canonical: Tremoctopus violaceus

Authorship: delle Chiaje 1830

```json
{"parsed":true,"quality":1,"verbatim":"Tremoctopus violaceus delle Chiaje, 1830","normalized":"Tremoctopus violaceus delle Chiaje 1830","canonical":{"stemmed":"Tremoctopus uiolace","simple":"Tremoctopus violaceus","full":"Tremoctopus violaceus"},"cardinality":2,"authorship":{"verbatim":"delle Chiaje, 1830","normalized":"delle Chiaje 1830","year":"1830","authors":["delle Chiaje"],"originalAuth":{"authors":["delle Chiaje"],"year":{"year":"1830"}}},"details":{"species":{"genus":"Tremoctopus","species":"violaceus","authorship":{"verbatim":"delle Chiaje, 1830","normalized":"delle Chiaje 1830","year":"1830","authors":["delle Chiaje"],"originalAuth":{"authors":["delle Chiaje"],"year":{"year":"1830"}}}}},"words":[{"verbatim":"Tremoctopus","normalized":"Tremoctopus","wordType":"GENUS","start":0,"end":11},{"verbatim":"violaceus","normalized":"violaceus","wordType":"SPECIES","start":12,"end":21},{"verbatim":"delle","normalized":"delle","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"Chiaje","normalized":"Chiaje","wordType":"AUTHOR_WORD","start":28,"end":34},{"verbatim":"1830","normalized":"1830","wordType":"YEAR","start":36,"end":40}],"id":"0543be2c-c14c-57e3-9529-570446ee1de4","parserVersion":"test_version"}
```

Name: Protis hydrothermica ten Hove & Zibrowius, 1986

Canonical: Protis hydrothermica

Authorship: ten Hove & Zibrowius 1986

```json
{"parsed":true,"quality":1,"verbatim":"Protis hydrothermica ten Hove \u0026 Zibrowius, 1986","normalized":"Protis hydrothermica ten Hove \u0026 Zibrowius 1986","canonical":{"stemmed":"Protis hydrothermic","simple":"Protis hydrothermica","full":"Protis hydrothermica"},"cardinality":2,"authorship":{"verbatim":"ten Hove \u0026 Zibrowius, 1986","normalized":"ten Hove \u0026 Zibrowius 1986","year":"1986","authors":["ten Hove","Zibrowius"],"originalAuth":{"authors":["ten Hove","Zibrowius"],"year":{"year":"1986"}}},"details":{"species":{"genus":"Protis","species":"hydrothermica","authorship":{"verbatim":"ten Hove \u0026 Zibrowius, 1986","normalized":"ten Hove \u0026 Zibrowius 1986","year":"1986","authors":["ten Hove","Zibrowius"],"originalAuth":{"authors":["ten Hove","Zibrowius"],"year":{"year":"1986"}}}}},"words":[{"verbatim":"Protis","normalized":"Protis","wordType":"GENUS","start":0,"end":6},{"verbatim":"hydrothermica","normalized":"hydrothermica","wordType":"SPECIES","start":7,"end":20},{"verbatim":"ten","normalized":"ten","wordType":"AUTHOR_WORD","start":21,"end":24},{"verbatim":"Hove","normalized":"Hove","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"Zibrowius","normalized":"Zibrowius","wordType":"AUTHOR_WORD","start":32,"end":41},{"verbatim":"1986","normalized":"1986","wordType":"YEAR","start":43,"end":47}],"id":"ef360f20-b14a-5eb2-a9ce-a5089956758b","parserVersion":"test_version"}
```

Name: Cladoniicola staurospora Diederich, van den Boom & Aptroot 2001

Canonical: Cladoniicola staurospora

Authorship: Diederich, van den Boom & Aptroot 2001

```json
{"parsed":true,"quality":1,"verbatim":"Cladoniicola staurospora Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Cladoniicola staurospora Diederich, van den Boom \u0026 Aptroot 2001","canonical":{"stemmed":"Cladoniicola staurospor","simple":"Cladoniicola staurospora","full":"Cladoniicola staurospora"},"cardinality":2,"authorship":{"verbatim":"Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Diederich, van den Boom \u0026 Aptroot 2001","year":"2001","authors":["Diederich","van den Boom","Aptroot"],"originalAuth":{"authors":["Diederich","van den Boom","Aptroot"],"year":{"year":"2001"}}},"details":{"species":{"genus":"Cladoniicola","species":"staurospora","authorship":{"verbatim":"Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Diederich, van den Boom \u0026 Aptroot 2001","year":"2001","authors":["Diederich","van den Boom","Aptroot"],"originalAuth":{"authors":["Diederich","van den Boom","Aptroot"],"year":{"year":"2001"}}}}},"words":[{"verbatim":"Cladoniicola","normalized":"Cladoniicola","wordType":"GENUS","start":0,"end":12},{"verbatim":"staurospora","normalized":"staurospora","wordType":"SPECIES","start":13,"end":24},{"verbatim":"Diederich","normalized":"Diederich","wordType":"AUTHOR_WORD","start":25,"end":34},{"verbatim":"van","normalized":"van","wordType":"AUTHOR_WORD","start":36,"end":39},{"verbatim":"den","normalized":"den","wordType":"AUTHOR_WORD","start":40,"end":43},{"verbatim":"Boom","normalized":"Boom","wordType":"AUTHOR_WORD","start":44,"end":48},{"verbatim":"Aptroot","normalized":"Aptroot","wordType":"AUTHOR_WORD","start":51,"end":58},{"verbatim":"2001","normalized":"2001","wordType":"YEAR","start":59,"end":63}],"id":"e59e3b01-311d-5dda-88e7-7e821440f5ee","parserVersion":"test_version"}
```

Name: Stagonospora polyspora M.T. Lucas & Sousa da Câmara 1934

Canonical: Stagonospora polyspora

Authorship: M. T. Lucas & Sousa da Câmara 1934

```json
{"parsed":true,"quality":1,"verbatim":"Stagonospora polyspora M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"Stagonospora polyspora M. T. Lucas \u0026 Sousa da Câmara 1934","canonical":{"stemmed":"Stagonospora polyspor","simple":"Stagonospora polyspora","full":"Stagonospora polyspora"},"cardinality":2,"authorship":{"verbatim":"M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}},"details":{"species":{"genus":"Stagonospora","species":"polyspora","authorship":{"verbatim":"M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}}}},"words":[{"verbatim":"Stagonospora","normalized":"Stagonospora","wordType":"GENUS","start":0,"end":12},{"verbatim":"polyspora","normalized":"polyspora","wordType":"SPECIES","start":13,"end":22},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Lucas","normalized":"Lucas","wordType":"AUTHOR_WORD","start":28,"end":33},{"verbatim":"Sousa","normalized":"Sousa","wordType":"AUTHOR_WORD","start":36,"end":41},{"verbatim":"da","normalized":"da","wordType":"AUTHOR_WORD","start":42,"end":44},{"verbatim":"Câmara","normalized":"Câmara","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"1934","normalized":"1934","wordType":"YEAR","start":52,"end":56}],"id":"f03d53d7-2db1-591f-8727-6b77c0af2e0c","parserVersion":"test_version"}
```

Name: Stagonospora polyspora M.T. Lucas et Sousa da Câmara 1934

Canonical: Stagonospora polyspora

Authorship: M. T. Lucas & Sousa da Câmara 1934

```json
{"parsed":true,"quality":1,"verbatim":"Stagonospora polyspora M.T. Lucas et Sousa da Câmara 1934","normalized":"Stagonospora polyspora M. T. Lucas \u0026 Sousa da Câmara 1934","canonical":{"stemmed":"Stagonospora polyspor","simple":"Stagonospora polyspora","full":"Stagonospora polyspora"},"cardinality":2,"authorship":{"verbatim":"M.T. Lucas et Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}},"details":{"species":{"genus":"Stagonospora","species":"polyspora","authorship":{"verbatim":"M.T. Lucas et Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}}}},"words":[{"verbatim":"Stagonospora","normalized":"Stagonospora","wordType":"GENUS","start":0,"end":12},{"verbatim":"polyspora","normalized":"polyspora","wordType":"SPECIES","start":13,"end":22},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Lucas","normalized":"Lucas","wordType":"AUTHOR_WORD","start":28,"end":33},{"verbatim":"Sousa","normalized":"Sousa","wordType":"AUTHOR_WORD","start":37,"end":42},{"verbatim":"da","normalized":"da","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"Câmara","normalized":"Câmara","wordType":"AUTHOR_WORD","start":46,"end":52},{"verbatim":"1934","normalized":"1934","wordType":"YEAR","start":53,"end":57}],"id":"a8a48393-0ca9-5916-83e3-fb32b7b0c422","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii U. Braun & Crous 2003

Canonical: Pseudocercospora dendrobii

Authorship: U. Braun & Crous 2003

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii U. Braun \u0026 Crous 2003","normalized":"Pseudocercospora dendrobii U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"U. Braun \u0026 Crous 2003","normalized":"U. Braun \u0026 Crous 2003","year":"2003","authors":["U. Braun","Crous"],"originalAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"U. Braun \u0026 Crous 2003","normalized":"U. Braun \u0026 Crous 2003","year":"2003","authors":["U. Braun","Crous"],"originalAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"U.","normalized":"U.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":30,"end":35},{"verbatim":"Crous","normalized":"Crous","wordType":"AUTHOR_WORD","start":38,"end":43},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":44,"end":48}],"id":"afd958fc-82a5-5551-951b-a725a49d3df0","parserVersion":"test_version"}
```

Name: Abaxisotima acuminata (Wang, Yuwen & Xiangwei Liu 1996)

Canonical: Abaxisotima acuminata

Authorship: (Wang, Yuwen & Xiangwei Liu 1996)

```json
{"parsed":true,"quality":1,"verbatim":"Abaxisotima acuminata (Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"Abaxisotima acuminata (Wang, Yuwen \u0026 Xiangwei Liu 1996)","canonical":{"stemmed":"Abaxisotima acuminat","simple":"Abaxisotima acuminata","full":"Abaxisotima acuminata"},"cardinality":2,"authorship":{"verbatim":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","year":"1996","authors":["Wang","Yuwen","Xiangwei Liu"],"originalAuth":{"authors":["Wang","Yuwen","Xiangwei Liu"],"year":{"year":"1996"}}},"details":{"species":{"genus":"Abaxisotima","species":"acuminata","authorship":{"verbatim":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","year":"1996","authors":["Wang","Yuwen","Xiangwei Liu"],"originalAuth":{"authors":["Wang","Yuwen","Xiangwei Liu"],"year":{"year":"1996"}}}}},"words":[{"verbatim":"Abaxisotima","normalized":"Abaxisotima","wordType":"GENUS","start":0,"end":11},{"verbatim":"acuminata","normalized":"acuminata","wordType":"SPECIES","start":12,"end":21},{"verbatim":"Wang","normalized":"Wang","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"Yuwen","normalized":"Yuwen","wordType":"AUTHOR_WORD","start":29,"end":34},{"verbatim":"Xiangwei","normalized":"Xiangwei","wordType":"AUTHOR_WORD","start":37,"end":45},{"verbatim":"Liu","normalized":"Liu","wordType":"AUTHOR_WORD","start":46,"end":49},{"verbatim":"1996","normalized":"1996","wordType":"YEAR","start":50,"end":54}],"id":"5eecff7d-181c-508c-832d-df4619b8b027","parserVersion":"test_version"}
```

Name: Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi & L. Tang, 2009

Canonical: Aboilomimus sichuanensis ornatus

Authorship: Liu, Xiang-wei, M. Zhou, W Bi & L. Tang 2009

```json
{"parsed":true,"quality":1,"verbatim":"Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","canonical":{"stemmed":"Aboilomimus sichuanens ornat","simple":"Aboilomimus sichuanensis ornatus","full":"Aboilomimus sichuanensis ornatus"},"cardinality":3,"authorship":{"verbatim":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","year":"2009","authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"originalAuth":{"authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"year":{"year":"2009"}}},"details":{"infraspecies":{"genus":"Aboilomimus","species":"sichuanensis","infraspecies":[{"value":"ornatus","authorship":{"verbatim":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","year":"2009","authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"originalAuth":{"authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"year":{"year":"2009"}}}}]}},"words":[{"verbatim":"Aboilomimus","normalized":"Aboilomimus","wordType":"GENUS","start":0,"end":11},{"verbatim":"sichuanensis","normalized":"sichuanensis","wordType":"SPECIES","start":12,"end":24},{"verbatim":"ornatus","normalized":"ornatus","wordType":"INFRASPECIES","start":25,"end":32},{"verbatim":"Liu","normalized":"Liu","wordType":"AUTHOR_WORD","start":33,"end":36},{"verbatim":"Xiang-wei","normalized":"Xiang-wei","wordType":"AUTHOR_WORD","start":38,"end":47},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"Zhou","normalized":"Zhou","wordType":"AUTHOR_WORD","start":52,"end":56},{"verbatim":"W","normalized":"W","wordType":"AUTHOR_WORD","start":58,"end":59},{"verbatim":"Bi","normalized":"Bi","wordType":"AUTHOR_WORD","start":60,"end":62},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":65,"end":67},{"verbatim":"Tang","normalized":"Tang","wordType":"AUTHOR_WORD","start":68,"end":72},{"verbatim":"2009","normalized":"2009","wordType":"YEAR","start":74,"end":78}],"id":"25ac4ba8-6595-5ab3-8463-f99f738bf4e4","parserVersion":"test_version"}
```
Name: Pseudocercospora Speg.

Canonical: Pseudocercospora

Authorship: Speg.

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora Speg.","normalized":"Pseudocercospora Speg.","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg.","normalized":"Speg.","authors":["Speg."],"originalAuth":{"authors":["Speg."]}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg.","normalized":"Speg.","authors":["Speg."],"originalAuth":{"authors":["Speg."]}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"UNINOMIAL","start":0,"end":16},{"verbatim":"Speg.","normalized":"Speg.","wordType":"AUTHOR_WORD","start":17,"end":22}],"id":"ccc7780b-c68b-53c6-9166-6b2d4902923e","parserVersion":"test_version"}
```

Name: Döringina Ihering 1929 (synonym)

Canonical: Doeringina

Authorship: Ihering 1929

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Döringina Ihering 1929 (synonym)","normalized":"Doeringina Ihering 1929","canonical":{"stemmed":"Doeringina","simple":"Doeringina","full":"Doeringina"},"cardinality":1,"authorship":{"verbatim":"Ihering 1929","normalized":"Ihering 1929","year":"1929","authors":["Ihering"],"originalAuth":{"authors":["Ihering"],"year":{"year":"1929"}}},"tail":" (synonym)","details":{"uninomial":{"uninomial":"Doeringina","authorship":{"verbatim":"Ihering 1929","normalized":"Ihering 1929","year":"1929","authors":["Ihering"],"originalAuth":{"authors":["Ihering"],"year":{"year":"1929"}}}}},"words":[{"verbatim":"Döringina","normalized":"Doeringina","wordType":"UNINOMIAL","start":0,"end":9},{"verbatim":"Ihering","normalized":"Ihering","wordType":"AUTHOR_WORD","start":10,"end":17},{"verbatim":"1929","normalized":"1929","wordType":"YEAR","start":18,"end":22}],"id":"95eb9081-5fe5-5497-be3d-ef0ce65a472c","parserVersion":"test_version"}
```

Name: Pseudocercospora Speg., Francis Jack.-Drake.

Canonical: Pseudocercospora

Authorship: Speg. & Francis Jack.-Drake.

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora Speg., Francis Jack.-Drake.","normalized":"Pseudocercospora Speg. \u0026 Francis Jack.-Drake.","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg., Francis Jack.-Drake.","normalized":"Speg. \u0026 Francis Jack.-Drake.","authors":["Speg.","Francis Jack.-Drake."],"originalAuth":{"authors":["Speg.","Francis Jack.-Drake."]}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg., Francis Jack.-Drake.","normalized":"Speg. \u0026 Francis Jack.-Drake.","authors":["Speg.","Francis Jack.-Drake."],"originalAuth":{"authors":["Speg.","Francis Jack.-Drake."]}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"UNINOMIAL","start":0,"end":16},{"verbatim":"Speg.","normalized":"Speg.","wordType":"AUTHOR_WORD","start":17,"end":22},{"verbatim":"Francis","normalized":"Francis","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"Jack.-Drake.","normalized":"Jack.-Drake.","wordType":"AUTHOR_WORD","start":32,"end":44}],"id":"25b015c7-a099-5bf6-91a9-cc8fde31f388","parserVersion":"test_version"}
```

Name: Aaaba de Laubenfels, 1936

Canonical: Aaaba

Authorship: de Laubenfels 1936

```json
{"parsed":true,"quality":1,"verbatim":"Aaaba de Laubenfels, 1936","normalized":"Aaaba de Laubenfels 1936","canonical":{"stemmed":"Aaaba","simple":"Aaaba","full":"Aaaba"},"cardinality":1,"authorship":{"verbatim":"de Laubenfels, 1936","normalized":"de Laubenfels 1936","year":"1936","authors":["de Laubenfels"],"originalAuth":{"authors":["de Laubenfels"],"year":{"year":"1936"}}},"details":{"uninomial":{"uninomial":"Aaaba","authorship":{"verbatim":"de Laubenfels, 1936","normalized":"de Laubenfels 1936","year":"1936","authors":["de Laubenfels"],"originalAuth":{"authors":["de Laubenfels"],"year":{"year":"1936"}}}}},"words":[{"verbatim":"Aaaba","normalized":"Aaaba","wordType":"UNINOMIAL","start":0,"end":5},{"verbatim":"de","normalized":"de","wordType":"AUTHOR_WORD","start":6,"end":8},{"verbatim":"Laubenfels","normalized":"Laubenfels","wordType":"AUTHOR_WORD","start":9,"end":19},{"verbatim":"1936","normalized":"1936","wordType":"YEAR","start":21,"end":25}],"id":"abead069-293d-5299-badd-c10c0f5545fb","parserVersion":"test_version"}
```

Name: Abbottia F. von Mueller, 1875

Canonical: Abbottia

Authorship: F. von Mueller 1875

```json
{"parsed":true,"quality":1,"verbatim":"Abbottia F. von Mueller, 1875","normalized":"Abbottia F. von Mueller 1875","canonical":{"stemmed":"Abbottia","simple":"Abbottia","full":"Abbottia"},"cardinality":1,"authorship":{"verbatim":"F. von Mueller, 1875","normalized":"F. von Mueller 1875","year":"1875","authors":["F. von Mueller"],"originalAuth":{"authors":["F. von Mueller"],"year":{"year":"1875"}}},"details":{"uninomial":{"uninomial":"Abbottia","authorship":{"verbatim":"F. von Mueller, 1875","normalized":"F. von Mueller 1875","year":"1875","authors":["F. von Mueller"],"originalAuth":{"authors":["F. von Mueller"],"year":{"year":"1875"}}}}},"words":[{"verbatim":"Abbottia","normalized":"Abbottia","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":9,"end":11},{"verbatim":"von","normalized":"von","wordType":"AUTHOR_WORD","start":12,"end":15},{"verbatim":"Mueller","normalized":"Mueller","wordType":"AUTHOR_WORD","start":16,"end":23},{"verbatim":"1875","normalized":"1875","wordType":"YEAR","start":25,"end":29}],"id":"34738de5-0112-56f0-85f2-0f4e815161b5","parserVersion":"test_version"}
```

Name: Abella von Heyden, 1826

Canonical: Abella

Authorship: von Heyden 1826

```json
{"parsed":true,"quality":1,"verbatim":"Abella von Heyden, 1826","normalized":"Abella von Heyden 1826","canonical":{"stemmed":"Abella","simple":"Abella","full":"Abella"},"cardinality":1,"authorship":{"verbatim":"von Heyden, 1826","normalized":"von Heyden 1826","year":"1826","authors":["von Heyden"],"originalAuth":{"authors":["von Heyden"],"year":{"year":"1826"}}},"details":{"uninomial":{"uninomial":"Abella","authorship":{"verbatim":"von Heyden, 1826","normalized":"von Heyden 1826","year":"1826","authors":["von Heyden"],"originalAuth":{"authors":["von Heyden"],"year":{"year":"1826"}}}}},"words":[{"verbatim":"Abella","normalized":"Abella","wordType":"UNINOMIAL","start":0,"end":6},{"verbatim":"von","normalized":"von","wordType":"AUTHOR_WORD","start":7,"end":10},{"verbatim":"Heyden","normalized":"Heyden","wordType":"AUTHOR_WORD","start":11,"end":17},{"verbatim":"1826","normalized":"1826","wordType":"YEAR","start":19,"end":23}],"id":"7dc5b624-1232-5072-bc4c-8eebde6c48b2","parserVersion":"test_version"}
```

Name: Micropleura v Linstow 1906

Canonical: Micropleura

Authorship: v Linstow 1906

```json
{"parsed":true,"quality":1,"verbatim":"Micropleura v Linstow 1906","normalized":"Micropleura v Linstow 1906","canonical":{"stemmed":"Micropleura","simple":"Micropleura","full":"Micropleura"},"cardinality":1,"authorship":{"verbatim":"v Linstow 1906","normalized":"v Linstow 1906","year":"1906","authors":["v Linstow"],"originalAuth":{"authors":["v Linstow"],"year":{"year":"1906"}}},"details":{"uninomial":{"uninomial":"Micropleura","authorship":{"verbatim":"v Linstow 1906","normalized":"v Linstow 1906","year":"1906","authors":["v Linstow"],"originalAuth":{"authors":["v Linstow"],"year":{"year":"1906"}}}}},"words":[{"verbatim":"Micropleura","normalized":"Micropleura","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"v","normalized":"v","wordType":"AUTHOR_WORD","start":12,"end":13},{"verbatim":"Linstow","normalized":"Linstow","wordType":"AUTHOR_WORD","start":14,"end":21},{"verbatim":"1906","normalized":"1906","wordType":"YEAR","start":22,"end":26}],"id":"94f99223-2631-52a9-9497-a29452387980","parserVersion":"test_version"}
```

Name: Pseudocercospora Speg. 1910

Canonical: Pseudocercospora

Authorship: Speg. 1910

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora Speg. 1910","normalized":"Pseudocercospora Speg. 1910","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg. 1910","normalized":"Speg. 1910","year":"1910","authors":["Speg."],"originalAuth":{"authors":["Speg."],"year":{"year":"1910"}}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg. 1910","normalized":"Speg. 1910","year":"1910","authors":["Speg."],"originalAuth":{"authors":["Speg."],"year":{"year":"1910"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"UNINOMIAL","start":0,"end":16},{"verbatim":"Speg.","normalized":"Speg.","wordType":"AUTHOR_WORD","start":17,"end":22},{"verbatim":"1910","normalized":"1910","wordType":"YEAR","start":23,"end":27}],"id":"eac97817-869a-5400-8b1e-0a125876189d","parserVersion":"test_version"}
```

Name: Pseudocercospora Spegazzini, 1910

Canonical: Pseudocercospora

Authorship: Spegazzini 1910

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora Spegazzini, 1910","normalized":"Pseudocercospora Spegazzini 1910","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Spegazzini, 1910","normalized":"Spegazzini 1910","year":"1910","authors":["Spegazzini"],"originalAuth":{"authors":["Spegazzini"],"year":{"year":"1910"}}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Spegazzini, 1910","normalized":"Spegazzini 1910","year":"1910","authors":["Spegazzini"],"originalAuth":{"authors":["Spegazzini"],"year":{"year":"1910"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"UNINOMIAL","start":0,"end":16},{"verbatim":"Spegazzini","normalized":"Spegazzini","wordType":"AUTHOR_WORD","start":17,"end":27},{"verbatim":"1910","normalized":"1910","wordType":"YEAR","start":29,"end":33}],"id":"6cc2922a-1f1d-5a40-90a7-b155fd16b233","parserVersion":"test_version"}
```

Name: Rhynchonellidae d'Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"quality":1,"verbatim":"Rhynchonellidae d'Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d'Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d'Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"words":[{"verbatim":"Rhynchonellidae","normalized":"Rhynchonellidae","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"d'Orbigny","normalized":"d'Orbigny","wordType":"AUTHOR_WORD","start":16,"end":25},{"verbatim":"1847","normalized":"1847","wordType":"YEAR","start":26,"end":30}],"id":"f3b90050-32f2-5009-ae9d-705fc58e45c4","parserVersion":"test_version"}
```

Name: Rhynchonellidae d‘Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Rhynchonellidae d‘Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d‘Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d‘Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"words":[{"verbatim":"Rhynchonellidae","normalized":"Rhynchonellidae","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"d‘Orbigny","normalized":"d'Orbigny","wordType":"AUTHOR_WORD","start":16,"end":25},{"verbatim":"1847","normalized":"1847","wordType":"YEAR","start":26,"end":30}],"id":"8a72add4-b276-5a92-ad30-a4c8bc03598a","parserVersion":"test_version"}
```

Name: Rhynchonellidae d’Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Rhynchonellidae d’Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d’Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d’Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"words":[{"verbatim":"Rhynchonellidae","normalized":"Rhynchonellidae","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"d’Orbigny","normalized":"d'Orbigny","wordType":"AUTHOR_WORD","start":16,"end":25},{"verbatim":"1847","normalized":"1847","wordType":"YEAR","start":26,"end":30}],"id":"cc9b39b8-b4d0-5e8e-9ffe-866454d3e49a","parserVersion":"test_version"}
```

Name: Ataladoris Iredale & O'Donoghue 1923

Canonical: Ataladoris

Authorship: Iredale & O'Donoghue 1923

```json
{"parsed":true,"quality":1,"verbatim":"Ataladoris Iredale \u0026 O'Donoghue 1923","normalized":"Ataladoris Iredale \u0026 O'Donoghue 1923","canonical":{"stemmed":"Ataladoris","simple":"Ataladoris","full":"Ataladoris"},"cardinality":1,"authorship":{"verbatim":"Iredale \u0026 O'Donoghue 1923","normalized":"Iredale \u0026 O'Donoghue 1923","year":"1923","authors":["Iredale","O'Donoghue"],"originalAuth":{"authors":["Iredale","O'Donoghue"],"year":{"year":"1923"}}},"details":{"uninomial":{"uninomial":"Ataladoris","authorship":{"verbatim":"Iredale \u0026 O'Donoghue 1923","normalized":"Iredale \u0026 O'Donoghue 1923","year":"1923","authors":["Iredale","O'Donoghue"],"originalAuth":{"authors":["Iredale","O'Donoghue"],"year":{"year":"1923"}}}}},"words":[{"verbatim":"Ataladoris","normalized":"Ataladoris","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"Iredale","normalized":"Iredale","wordType":"AUTHOR_WORD","start":11,"end":18},{"verbatim":"O'Donoghue","normalized":"O'Donoghue","wordType":"AUTHOR_WORD","start":21,"end":31},{"verbatim":"1923","normalized":"1923","wordType":"YEAR","start":32,"end":36}],"id":"dbb90380-0552-5237-82ef-8a8b07e42049","parserVersion":"test_version"}
```

Name: Anteplana le Renard 1995

Canonical: Anteplana

Authorship: le Renard 1995

```json
{"parsed":true,"quality":1,"verbatim":"Anteplana le Renard 1995","normalized":"Anteplana le Renard 1995","canonical":{"stemmed":"Anteplana","simple":"Anteplana","full":"Anteplana"},"cardinality":1,"authorship":{"verbatim":"le Renard 1995","normalized":"le Renard 1995","year":"1995","authors":["le Renard"],"originalAuth":{"authors":["le Renard"],"year":{"year":"1995"}}},"details":{"uninomial":{"uninomial":"Anteplana","authorship":{"verbatim":"le Renard 1995","normalized":"le Renard 1995","year":"1995","authors":["le Renard"],"originalAuth":{"authors":["le Renard"],"year":{"year":"1995"}}}}},"words":[{"verbatim":"Anteplana","normalized":"Anteplana","wordType":"UNINOMIAL","start":0,"end":9},{"verbatim":"le","normalized":"le","wordType":"AUTHOR_WORD","start":10,"end":12},{"verbatim":"Renard","normalized":"Renard","wordType":"AUTHOR_WORD","start":13,"end":19},{"verbatim":"1995","normalized":"1995","wordType":"YEAR","start":20,"end":24}],"id":"6920744c-27e9-546f-96d9-c8859544ef78","parserVersion":"test_version"}
```

Name: Candinia le Renard, Sabelli & Taviani 1996

Canonical: Candinia

Authorship: le Renard, Sabelli & Taviani 1996

```json
{"parsed":true,"quality":1,"verbatim":"Candinia le Renard, Sabelli \u0026 Taviani 1996","normalized":"Candinia le Renard, Sabelli \u0026 Taviani 1996","canonical":{"stemmed":"Candinia","simple":"Candinia","full":"Candinia"},"cardinality":1,"authorship":{"verbatim":"le Renard, Sabelli \u0026 Taviani 1996","normalized":"le Renard, Sabelli \u0026 Taviani 1996","year":"1996","authors":["le Renard","Sabelli","Taviani"],"originalAuth":{"authors":["le Renard","Sabelli","Taviani"],"year":{"year":"1996"}}},"details":{"uninomial":{"uninomial":"Candinia","authorship":{"verbatim":"le Renard, Sabelli \u0026 Taviani 1996","normalized":"le Renard, Sabelli \u0026 Taviani 1996","year":"1996","authors":["le Renard","Sabelli","Taviani"],"originalAuth":{"authors":["le Renard","Sabelli","Taviani"],"year":{"year":"1996"}}}}},"words":[{"verbatim":"Candinia","normalized":"Candinia","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"le","normalized":"le","wordType":"AUTHOR_WORD","start":9,"end":11},{"verbatim":"Renard","normalized":"Renard","wordType":"AUTHOR_WORD","start":12,"end":18},{"verbatim":"Sabelli","normalized":"Sabelli","wordType":"AUTHOR_WORD","start":20,"end":27},{"verbatim":"Taviani","normalized":"Taviani","wordType":"AUTHOR_WORD","start":30,"end":37},{"verbatim":"1996","normalized":"1996","wordType":"YEAR","start":38,"end":42}],"id":"2a92b7b1-4da8-5571-98de-9cd225526081","parserVersion":"test_version"}
```

Name: Polypodium le Sourdianum Fourn.

Canonical: Polypodium

Authorship: le Sourdianum Fourn.

```json
{"parsed":true,"quality":1,"verbatim":"Polypodium le Sourdianum Fourn.","normalized":"Polypodium le Sourdianum Fourn.","canonical":{"stemmed":"Polypodium","simple":"Polypodium","full":"Polypodium"},"cardinality":1,"authorship":{"verbatim":"le Sourdianum Fourn.","normalized":"le Sourdianum Fourn.","authors":["le Sourdianum Fourn."],"originalAuth":{"authors":["le Sourdianum Fourn."]}},"details":{"uninomial":{"uninomial":"Polypodium","authorship":{"verbatim":"le Sourdianum Fourn.","normalized":"le Sourdianum Fourn.","authors":["le Sourdianum Fourn."],"originalAuth":{"authors":["le Sourdianum Fourn."]}}}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"le","normalized":"le","wordType":"AUTHOR_WORD","start":11,"end":13},{"verbatim":"Sourdianum","normalized":"Sourdianum","wordType":"AUTHOR_WORD","start":14,"end":24},{"verbatim":"Fourn.","normalized":"Fourn.","wordType":"AUTHOR_WORD","start":25,"end":31}],"id":"ea72f0d9-2f8a-5ba0-95c7-986075eda321","parserVersion":"test_version"}
```

### Two-letter genus names (legacy genera, not allowed anymore)

Name: Ca Dyar 1914

Canonical: Ca

Authorship: Dyar 1914

```json
{"parsed":true,"quality":1,"verbatim":"Ca Dyar 1914","normalized":"Ca Dyar 1914","canonical":{"stemmed":"Ca","simple":"Ca","full":"Ca"},"cardinality":1,"authorship":{"verbatim":"Dyar 1914","normalized":"Dyar 1914","year":"1914","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1914"}}},"details":{"uninomial":{"uninomial":"Ca","authorship":{"verbatim":"Dyar 1914","normalized":"Dyar 1914","year":"1914","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1914"}}}}},"words":[{"verbatim":"Ca","normalized":"Ca","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Dyar","normalized":"Dyar","wordType":"AUTHOR_WORD","start":3,"end":7},{"verbatim":"1914","normalized":"1914","wordType":"YEAR","start":8,"end":12}],"id":"ccb4663f-3d9a-5447-ab28-13e453738075","parserVersion":"test_version"}
```

Name: Ea Distant 1911

Canonical: Ea

Authorship: Distant 1911

```json
{"parsed":true,"quality":1,"verbatim":"Ea Distant 1911","normalized":"Ea Distant 1911","canonical":{"stemmed":"Ea","simple":"Ea","full":"Ea"},"cardinality":1,"authorship":{"verbatim":"Distant 1911","normalized":"Distant 1911","year":"1911","authors":["Distant"],"originalAuth":{"authors":["Distant"],"year":{"year":"1911"}}},"details":{"uninomial":{"uninomial":"Ea","authorship":{"verbatim":"Distant 1911","normalized":"Distant 1911","year":"1911","authors":["Distant"],"originalAuth":{"authors":["Distant"],"year":{"year":"1911"}}}}},"words":[{"verbatim":"Ea","normalized":"Ea","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Distant","normalized":"Distant","wordType":"AUTHOR_WORD","start":3,"end":10},{"verbatim":"1911","normalized":"1911","wordType":"YEAR","start":11,"end":15}],"id":"c5a5643f-452f-5c51-91eb-42789ed6f3a4","parserVersion":"test_version"}
```

Name: Do

Canonical: Do

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Do","normalized":"Do","canonical":{"stemmed":"Do","simple":"Do","full":"Do"},"cardinality":1,"details":{"uninomial":{"uninomial":"Do"}},"words":[{"verbatim":"Do","normalized":"Do","wordType":"UNINOMIAL","start":0,"end":2}],"id":"f3b54204-8d34-5c7d-94a3-85d89dd99d86","parserVersion":"test_version"}
```

Name: Ge Nicéville 1895

Canonical: Ge

Authorship: Nicéville 1895

```json
{"parsed":true,"quality":1,"verbatim":"Ge Nicéville 1895","normalized":"Ge Nicéville 1895","canonical":{"stemmed":"Ge","simple":"Ge","full":"Ge"},"cardinality":1,"authorship":{"verbatim":"Nicéville 1895","normalized":"Nicéville 1895","year":"1895","authors":["Nicéville"],"originalAuth":{"authors":["Nicéville"],"year":{"year":"1895"}}},"details":{"uninomial":{"uninomial":"Ge","authorship":{"verbatim":"Nicéville 1895","normalized":"Nicéville 1895","year":"1895","authors":["Nicéville"],"originalAuth":{"authors":["Nicéville"],"year":{"year":"1895"}}}}},"words":[{"verbatim":"Ge","normalized":"Ge","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Nicéville","normalized":"Nicéville","wordType":"AUTHOR_WORD","start":3,"end":12},{"verbatim":"1895","normalized":"1895","wordType":"YEAR","start":13,"end":17}],"id":"ba4f0f90-1df5-5054-a17b-15938a942d88","parserVersion":"test_version"}
```

Name: Ia Thomas 1902

Canonical: Ia

Authorship: Thomas 1902

```json
{"parsed":true,"quality":1,"verbatim":"Ia Thomas 1902","normalized":"Ia Thomas 1902","canonical":{"stemmed":"Ia","simple":"Ia","full":"Ia"},"cardinality":1,"authorship":{"verbatim":"Thomas 1902","normalized":"Thomas 1902","year":"1902","authors":["Thomas"],"originalAuth":{"authors":["Thomas"],"year":{"year":"1902"}}},"details":{"uninomial":{"uninomial":"Ia","authorship":{"verbatim":"Thomas 1902","normalized":"Thomas 1902","year":"1902","authors":["Thomas"],"originalAuth":{"authors":["Thomas"],"year":{"year":"1902"}}}}},"words":[{"verbatim":"Ia","normalized":"Ia","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Thomas","normalized":"Thomas","wordType":"AUTHOR_WORD","start":3,"end":9},{"verbatim":"1902","normalized":"1902","wordType":"YEAR","start":10,"end":14}],"id":"9826997c-1d52-5de2-8b7b-facdc9fb73f2","parserVersion":"test_version"}
```

Name: Io Lea 1831

Canonical: Io

Authorship: Lea 1831

```json
{"parsed":true,"quality":1,"verbatim":"Io Lea 1831","normalized":"Io Lea 1831","canonical":{"stemmed":"Io","simple":"Io","full":"Io"},"cardinality":1,"authorship":{"verbatim":"Lea 1831","normalized":"Lea 1831","year":"1831","authors":["Lea"],"originalAuth":{"authors":["Lea"],"year":{"year":"1831"}}},"details":{"uninomial":{"uninomial":"Io","authorship":{"verbatim":"Lea 1831","normalized":"Lea 1831","year":"1831","authors":["Lea"],"originalAuth":{"authors":["Lea"],"year":{"year":"1831"}}}}},"words":[{"verbatim":"Io","normalized":"Io","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Lea","normalized":"Lea","wordType":"AUTHOR_WORD","start":3,"end":6},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":7,"end":11}],"id":"3cc533a5-4f2c-5aec-ba30-85a27548aa95","parserVersion":"test_version"}
```

Name: Io Blanchard 1852

Canonical: Io

Authorship: Blanchard 1852

```json
{"parsed":true,"quality":1,"verbatim":"Io Blanchard 1852","normalized":"Io Blanchard 1852","canonical":{"stemmed":"Io","simple":"Io","full":"Io"},"cardinality":1,"authorship":{"verbatim":"Blanchard 1852","normalized":"Blanchard 1852","year":"1852","authors":["Blanchard"],"originalAuth":{"authors":["Blanchard"],"year":{"year":"1852"}}},"details":{"uninomial":{"uninomial":"Io","authorship":{"verbatim":"Blanchard 1852","normalized":"Blanchard 1852","year":"1852","authors":["Blanchard"],"originalAuth":{"authors":["Blanchard"],"year":{"year":"1852"}}}}},"words":[{"verbatim":"Io","normalized":"Io","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Blanchard","normalized":"Blanchard","wordType":"AUTHOR_WORD","start":3,"end":12},{"verbatim":"1852","normalized":"1852","wordType":"YEAR","start":13,"end":17}],"id":"4de7e503-a5a5-5309-bc6c-cbaf90a9199b","parserVersion":"test_version"}
```

Name: Ix Bergroth 1916

Canonical: Ix

Authorship: Bergroth 1916

```json
{"parsed":true,"quality":1,"verbatim":"Ix Bergroth 1916","normalized":"Ix Bergroth 1916","canonical":{"stemmed":"Ix","simple":"Ix","full":"Ix"},"cardinality":1,"authorship":{"verbatim":"Bergroth 1916","normalized":"Bergroth 1916","year":"1916","authors":["Bergroth"],"originalAuth":{"authors":["Bergroth"],"year":{"year":"1916"}}},"details":{"uninomial":{"uninomial":"Ix","authorship":{"verbatim":"Bergroth 1916","normalized":"Bergroth 1916","year":"1916","authors":["Bergroth"],"originalAuth":{"authors":["Bergroth"],"year":{"year":"1916"}}}}},"words":[{"verbatim":"Ix","normalized":"Ix","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Bergroth","normalized":"Bergroth","wordType":"AUTHOR_WORD","start":3,"end":11},{"verbatim":"1916","normalized":"1916","wordType":"YEAR","start":12,"end":16}],"id":"981228e8-45fe-5b7b-ab78-4793cae51602","parserVersion":"test_version"}
```

Name: Lo Seale 1906

Canonical: Lo

Authorship: Seale 1906

```json
{"parsed":true,"quality":1,"verbatim":"Lo Seale 1906","normalized":"Lo Seale 1906","canonical":{"stemmed":"Lo","simple":"Lo","full":"Lo"},"cardinality":1,"authorship":{"verbatim":"Seale 1906","normalized":"Seale 1906","year":"1906","authors":["Seale"],"originalAuth":{"authors":["Seale"],"year":{"year":"1906"}}},"details":{"uninomial":{"uninomial":"Lo","authorship":{"verbatim":"Seale 1906","normalized":"Seale 1906","year":"1906","authors":["Seale"],"originalAuth":{"authors":["Seale"],"year":{"year":"1906"}}}}},"words":[{"verbatim":"Lo","normalized":"Lo","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Seale","normalized":"Seale","wordType":"AUTHOR_WORD","start":3,"end":8},{"verbatim":"1906","normalized":"1906","wordType":"YEAR","start":9,"end":13}],"id":"8d9cb022-3458-5473-aa5a-91da319d5d78","parserVersion":"test_version"}
```

Name: Oa Girault 1929

Canonical: Oa

Authorship: Girault 1929

```json
{"parsed":true,"quality":1,"verbatim":"Oa Girault 1929","normalized":"Oa Girault 1929","canonical":{"stemmed":"Oa","simple":"Oa","full":"Oa"},"cardinality":1,"authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}},"details":{"uninomial":{"uninomial":"Oa","authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}}}},"words":[{"verbatim":"Oa","normalized":"Oa","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Girault","normalized":"Girault","wordType":"AUTHOR_WORD","start":3,"end":10},{"verbatim":"1929","normalized":"1929","wordType":"YEAR","start":11,"end":15}],"id":"14647a9c-70c8-55a8-b2a7-1fc47c39732b","parserVersion":"test_version"}
```

Name: Oo

Canonical: Oo

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Oo","normalized":"Oo","canonical":{"stemmed":"Oo","simple":"Oo","full":"Oo"},"cardinality":1,"details":{"uninomial":{"uninomial":"Oo"}},"words":[{"verbatim":"Oo","normalized":"Oo","wordType":"UNINOMIAL","start":0,"end":2}],"id":"2b54a44f-a680-5a5c-9b68-4206605f2145","parserVersion":"test_version"}
```

Name: Nu

Canonical: Nu

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Nu","normalized":"Nu","canonical":{"stemmed":"Nu","simple":"Nu","full":"Nu"},"cardinality":1,"details":{"uninomial":{"uninomial":"Nu"}},"words":[{"verbatim":"Nu","normalized":"Nu","wordType":"UNINOMIAL","start":0,"end":2}],"id":"9d068582-37e2-5223-87d0-86562388b35c","parserVersion":"test_version"}
```

Name: Ra Whitley 1931

Canonical: Ra

Authorship: Whitley 1931

```json
{"parsed":true,"quality":1,"verbatim":"Ra Whitley 1931","normalized":"Ra Whitley 1931","canonical":{"stemmed":"Ra","simple":"Ra","full":"Ra"},"cardinality":1,"authorship":{"verbatim":"Whitley 1931","normalized":"Whitley 1931","year":"1931","authors":["Whitley"],"originalAuth":{"authors":["Whitley"],"year":{"year":"1931"}}},"details":{"uninomial":{"uninomial":"Ra","authorship":{"verbatim":"Whitley 1931","normalized":"Whitley 1931","year":"1931","authors":["Whitley"],"originalAuth":{"authors":["Whitley"],"year":{"year":"1931"}}}}},"words":[{"verbatim":"Ra","normalized":"Ra","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Whitley","normalized":"Whitley","wordType":"AUTHOR_WORD","start":3,"end":10},{"verbatim":"1931","normalized":"1931","wordType":"YEAR","start":11,"end":15}],"id":"72b5b436-6381-5939-b8d1-7f04bb2a82bb","parserVersion":"test_version"}
```

Name: Ty Bory de St. Vincent 1827

Canonical: Ty

Authorship: Bory de St. Vincent 1827

```json
{"parsed":true,"quality":1,"verbatim":"Ty Bory de St. Vincent 1827","normalized":"Ty Bory de St. Vincent 1827","canonical":{"stemmed":"Ty","simple":"Ty","full":"Ty"},"cardinality":1,"authorship":{"verbatim":"Bory de St. Vincent 1827","normalized":"Bory de St. Vincent 1827","year":"1827","authors":["Bory de St. Vincent"],"originalAuth":{"authors":["Bory de St. Vincent"],"year":{"year":"1827"}}},"details":{"uninomial":{"uninomial":"Ty","authorship":{"verbatim":"Bory de St. Vincent 1827","normalized":"Bory de St. Vincent 1827","year":"1827","authors":["Bory de St. Vincent"],"originalAuth":{"authors":["Bory de St. Vincent"],"year":{"year":"1827"}}}}},"words":[{"verbatim":"Ty","normalized":"Ty","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Bory","normalized":"Bory","wordType":"AUTHOR_WORD","start":3,"end":7},{"verbatim":"de","normalized":"de","wordType":"AUTHOR_WORD","start":8,"end":10},{"verbatim":"St.","normalized":"St.","wordType":"AUTHOR_WORD","start":11,"end":14},{"verbatim":"Vincent","normalized":"Vincent","wordType":"AUTHOR_WORD","start":15,"end":22},{"verbatim":"1827","normalized":"1827","wordType":"YEAR","start":23,"end":27}],"id":"1d05b120-8f75-58ab-bdf7-c181fdf1bc3c","parserVersion":"test_version"}
```

Name: Ua Girault 1929

Canonical: Ua

Authorship: Girault 1929

```json
{"parsed":true,"quality":1,"verbatim":"Ua Girault 1929","normalized":"Ua Girault 1929","canonical":{"stemmed":"Ua","simple":"Ua","full":"Ua"},"cardinality":1,"authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}},"details":{"uninomial":{"uninomial":"Ua","authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}}}},"words":[{"verbatim":"Ua","normalized":"Ua","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Girault","normalized":"Girault","wordType":"AUTHOR_WORD","start":3,"end":10},{"verbatim":"1929","normalized":"1929","wordType":"YEAR","start":11,"end":15}],"id":"aee3fe77-1797-5172-82f1-5ee233108c15","parserVersion":"test_version"}
```

Name: Aa Baker 1940

Canonical: Aa

Authorship: Baker 1940

```json
{"parsed":true,"quality":1,"verbatim":"Aa Baker 1940","normalized":"Aa Baker 1940","canonical":{"stemmed":"Aa","simple":"Aa","full":"Aa"},"cardinality":1,"authorship":{"verbatim":"Baker 1940","normalized":"Baker 1940","year":"1940","authors":["Baker"],"originalAuth":{"authors":["Baker"],"year":{"year":"1940"}}},"details":{"uninomial":{"uninomial":"Aa","authorship":{"verbatim":"Baker 1940","normalized":"Baker 1940","year":"1940","authors":["Baker"],"originalAuth":{"authors":["Baker"],"year":{"year":"1940"}}}}},"words":[{"verbatim":"Aa","normalized":"Aa","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Baker","normalized":"Baker","wordType":"AUTHOR_WORD","start":3,"end":8},{"verbatim":"1940","normalized":"1940","wordType":"YEAR","start":9,"end":13}],"id":"101d126d-c14a-5043-a1d8-72bc6a9f4dcf","parserVersion":"test_version"}
```

Name: Ja Uéno 1955

Canonical: Ja

Authorship: Uéno 1955

```json
{"parsed":true,"quality":1,"verbatim":"Ja Uéno 1955","normalized":"Ja Uéno 1955","canonical":{"stemmed":"Ja","simple":"Ja","full":"Ja"},"cardinality":1,"authorship":{"verbatim":"Uéno 1955","normalized":"Uéno 1955","year":"1955","authors":["Uéno"],"originalAuth":{"authors":["Uéno"],"year":{"year":"1955"}}},"details":{"uninomial":{"uninomial":"Ja","authorship":{"verbatim":"Uéno 1955","normalized":"Uéno 1955","year":"1955","authors":["Uéno"],"originalAuth":{"authors":["Uéno"],"year":{"year":"1955"}}}}},"words":[{"verbatim":"Ja","normalized":"Ja","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Uéno","normalized":"Uéno","wordType":"AUTHOR_WORD","start":3,"end":7},{"verbatim":"1955","normalized":"1955","wordType":"YEAR","start":8,"end":12}],"id":"45f6eba8-1063-590d-bc4a-9f9ffdef4a10","parserVersion":"test_version"}
```

Name: Zu Walters & Fitch 1960

Canonical: Zu

Authorship: Walters & Fitch 1960

```json
{"parsed":true,"quality":1,"verbatim":"Zu Walters \u0026 Fitch 1960","normalized":"Zu Walters \u0026 Fitch 1960","canonical":{"stemmed":"Zu","simple":"Zu","full":"Zu"},"cardinality":1,"authorship":{"verbatim":"Walters \u0026 Fitch 1960","normalized":"Walters \u0026 Fitch 1960","year":"1960","authors":["Walters","Fitch"],"originalAuth":{"authors":["Walters","Fitch"],"year":{"year":"1960"}}},"details":{"uninomial":{"uninomial":"Zu","authorship":{"verbatim":"Walters \u0026 Fitch 1960","normalized":"Walters \u0026 Fitch 1960","year":"1960","authors":["Walters","Fitch"],"originalAuth":{"authors":["Walters","Fitch"],"year":{"year":"1960"}}}}},"words":[{"verbatim":"Zu","normalized":"Zu","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Walters","normalized":"Walters","wordType":"AUTHOR_WORD","start":3,"end":10},{"verbatim":"Fitch","normalized":"Fitch","wordType":"AUTHOR_WORD","start":13,"end":18},{"verbatim":"1960","normalized":"1960","wordType":"YEAR","start":19,"end":23}],"id":"c8724802-7dfb-5743-9988-a5f11b4c57b5","parserVersion":"test_version"}
```

Name: La Bleszynski 1966

Canonical: La

Authorship: Bleszynski 1966

```json
{"parsed":true,"quality":1,"verbatim":"La Bleszynski 1966","normalized":"La Bleszynski 1966","canonical":{"stemmed":"La","simple":"La","full":"La"},"cardinality":1,"authorship":{"verbatim":"Bleszynski 1966","normalized":"Bleszynski 1966","year":"1966","authors":["Bleszynski"],"originalAuth":{"authors":["Bleszynski"],"year":{"year":"1966"}}},"details":{"uninomial":{"uninomial":"La","authorship":{"verbatim":"Bleszynski 1966","normalized":"Bleszynski 1966","year":"1966","authors":["Bleszynski"],"originalAuth":{"authors":["Bleszynski"],"year":{"year":"1966"}}}}},"words":[{"verbatim":"La","normalized":"La","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Bleszynski","normalized":"Bleszynski","wordType":"AUTHOR_WORD","start":3,"end":13},{"verbatim":"1966","normalized":"1966","wordType":"YEAR","start":14,"end":18}],"id":"002f2de4-3661-5c8f-9175-cc1d1a9d6467","parserVersion":"test_version"}
```

Name: Qu Durkoop

Canonical: Qu

Authorship: Durkoop

```json
{"parsed":true,"quality":1,"verbatim":"Qu Durkoop","normalized":"Qu Durkoop","canonical":{"stemmed":"Qu","simple":"Qu","full":"Qu"},"cardinality":1,"authorship":{"verbatim":"Durkoop","normalized":"Durkoop","authors":["Durkoop"],"originalAuth":{"authors":["Durkoop"]}},"details":{"uninomial":{"uninomial":"Qu","authorship":{"verbatim":"Durkoop","normalized":"Durkoop","authors":["Durkoop"],"originalAuth":{"authors":["Durkoop"]}}}},"words":[{"verbatim":"Qu","normalized":"Qu","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Durkoop","normalized":"Durkoop","wordType":"AUTHOR_WORD","start":3,"end":10}],"id":"b4d879fa-028f-5b03-ad38-cc3a0765779a","parserVersion":"test_version"}
```

Name: As Slipinski 1982

Canonical: As

Authorship: Slipinski 1982

```json
{"parsed":true,"quality":1,"verbatim":"As Slipinski 1982","normalized":"As Slipinski 1982","canonical":{"stemmed":"As","simple":"As","full":"As"},"cardinality":1,"authorship":{"verbatim":"Slipinski 1982","normalized":"Slipinski 1982","year":"1982","authors":["Slipinski"],"originalAuth":{"authors":["Slipinski"],"year":{"year":"1982"}}},"details":{"uninomial":{"uninomial":"As","authorship":{"verbatim":"Slipinski 1982","normalized":"Slipinski 1982","year":"1982","authors":["Slipinski"],"originalAuth":{"authors":["Slipinski"],"year":{"year":"1982"}}}}},"words":[{"verbatim":"As","normalized":"As","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Slipinski","normalized":"Slipinski","wordType":"AUTHOR_WORD","start":3,"end":12},{"verbatim":"1982","normalized":"1982","wordType":"YEAR","start":13,"end":17}],"id":"55237f82-2126-5579-a8c6-385c0eb7ed8e","parserVersion":"test_version"}
```

Name: Ba Solem 1983

Canonical: Ba

Authorship: Solem 1983

```json
{"parsed":true,"quality":1,"verbatim":"Ba Solem 1983","normalized":"Ba Solem 1983","canonical":{"stemmed":"Ba","simple":"Ba","full":"Ba"},"cardinality":1,"authorship":{"verbatim":"Solem 1983","normalized":"Solem 1983","year":"1983","authors":["Solem"],"originalAuth":{"authors":["Solem"],"year":{"year":"1983"}}},"details":{"uninomial":{"uninomial":"Ba","authorship":{"verbatim":"Solem 1983","normalized":"Solem 1983","year":"1983","authors":["Solem"],"originalAuth":{"authors":["Solem"],"year":{"year":"1983"}}}}},"words":[{"verbatim":"Ba","normalized":"Ba","wordType":"UNINOMIAL","start":0,"end":2},{"verbatim":"Solem","normalized":"Solem","wordType":"AUTHOR_WORD","start":3,"end":8},{"verbatim":"1983","normalized":"1983","wordType":"YEAR","start":9,"end":13}],"id":"452f1a8e-711a-5b9c-906c-f475015229dd","parserVersion":"test_version"}
```

### Combination of two uninomials

Name: Agaricus tr. Hypholoma Fr.

Canonical: Agaricus trib. Hypholoma

Authorship: Fr.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Agaricus tr. Hypholoma Fr.","normalized":"Agaricus trib. Hypholoma Fr.","canonical":{"stemmed":"Hypholoma","simple":"Hypholoma","full":"Agaricus trib. Hypholoma"},"cardinality":1,"authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}},"details":{"uninomial":{"uninomial":"Hypholoma","rank":"trib.","parent":"Agaricus","authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}}}},"words":[{"verbatim":"Agaricus","normalized":"Agaricus","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"tr.","normalized":"trib.","wordType":"RANK","start":9,"end":12},{"verbatim":"Hypholoma","normalized":"Hypholoma","wordType":"UNINOMIAL","start":13,"end":22},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":23,"end":26}],"id":"e00a0fc2-e0b2-53e4-9ec3-3d6f793c772f","parserVersion":"test_version"}
```

Name: Agaricus tr Hypholoma Fr.

Canonical: Agaricus trib. Hypholoma

Authorship: Fr.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Agaricus tr Hypholoma Fr.","normalized":"Agaricus trib. Hypholoma Fr.","canonical":{"stemmed":"Hypholoma","simple":"Hypholoma","full":"Agaricus trib. Hypholoma"},"cardinality":1,"authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}},"details":{"uninomial":{"uninomial":"Hypholoma","rank":"trib.","parent":"Agaricus","authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}}}},"words":[{"verbatim":"Agaricus","normalized":"Agaricus","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"tr","normalized":"trib.","wordType":"RANK","start":9,"end":11},{"verbatim":"Hypholoma","normalized":"Hypholoma","wordType":"UNINOMIAL","start":12,"end":21},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":22,"end":25}],"id":"2e90eb8f-6371-5c1f-9ec1-e753a30f0e86","parserVersion":"test_version"}
```

Name: Agaricus subtr. Oesypii Fr.

Canonical: Agaricus subtrib. Oesypii

Authorship: Fr.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Agaricus subtr. Oesypii Fr.","normalized":"Agaricus subtrib. Oesypii Fr.","canonical":{"stemmed":"Oesypii","simple":"Oesypii","full":"Agaricus subtrib. Oesypii"},"cardinality":1,"authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}},"details":{"uninomial":{"uninomial":"Oesypii","rank":"subtrib.","parent":"Agaricus","authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}}}},"words":[{"verbatim":"Agaricus","normalized":"Agaricus","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"subtr.","normalized":"subtrib.","wordType":"RANK","start":9,"end":15},{"verbatim":"Oesypii","normalized":"Oesypii","wordType":"UNINOMIAL","start":16,"end":23},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":24,"end":27}],"id":"b58fd6e6-71d1-5889-9ada-519ffb0efd7c","parserVersion":"test_version"}
```

Name: Agaricus subtr Oesypii Fr.

Canonical: Agaricus subtrib. Oesypii

Authorship: Fr.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Agaricus subtr Oesypii Fr.","normalized":"Agaricus subtrib. Oesypii Fr.","canonical":{"stemmed":"Oesypii","simple":"Oesypii","full":"Agaricus subtrib. Oesypii"},"cardinality":1,"authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}},"details":{"uninomial":{"uninomial":"Oesypii","rank":"subtrib.","parent":"Agaricus","authorship":{"verbatim":"Fr.","normalized":"Fr.","authors":["Fr."],"originalAuth":{"authors":["Fr."]}}}},"words":[{"verbatim":"Agaricus","normalized":"Agaricus","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"subtr","normalized":"subtrib.","wordType":"RANK","start":9,"end":14},{"verbatim":"Oesypii","normalized":"Oesypii","wordType":"UNINOMIAL","start":15,"end":22},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":23,"end":26}],"id":"20360ca1-6217-5d48-915d-e5515832e2b2","parserVersion":"test_version"}
```

Name: Poaceae subtrib. Scolochloinae Soreng

Canonical: Poaceae subtrib. Scolochloinae

Authorship: Soreng

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Poaceae subtrib. Scolochloinae Soreng","normalized":"Poaceae subtrib. Scolochloinae Soreng","canonical":{"stemmed":"Scolochloinae","simple":"Scolochloinae","full":"Poaceae subtrib. Scolochloinae"},"cardinality":1,"authorship":{"verbatim":"Soreng","normalized":"Soreng","authors":["Soreng"],"originalAuth":{"authors":["Soreng"]}},"details":{"uninomial":{"uninomial":"Scolochloinae","rank":"subtrib.","parent":"Poaceae","authorship":{"verbatim":"Soreng","normalized":"Soreng","authors":["Soreng"],"originalAuth":{"authors":["Soreng"]}}}},"words":[{"verbatim":"Poaceae","normalized":"Poaceae","wordType":"UNINOMIAL","start":0,"end":7},{"verbatim":"subtrib.","normalized":"subtrib.","wordType":"RANK","start":8,"end":16},{"verbatim":"Scolochloinae","normalized":"Scolochloinae","wordType":"UNINOMIAL","start":17,"end":30},{"verbatim":"Soreng","normalized":"Soreng","wordType":"AUTHOR_WORD","start":31,"end":37}],"id":"d10510a7-ad50-587a-8411-e03d30d44214","parserVersion":"test_version"}
```

Name: Zygophyllaceae subfam. Tribuloideae D.M.Porter

Canonical: Zygophyllaceae subfam. Tribuloideae

Authorship: D. M. Porter

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Zygophyllaceae subfam. Tribuloideae D.M.Porter","normalized":"Zygophyllaceae subfam. Tribuloideae D. M. Porter","canonical":{"stemmed":"Tribuloideae","simple":"Tribuloideae","full":"Zygophyllaceae subfam. Tribuloideae"},"cardinality":1,"authorship":{"verbatim":"D.M.Porter","normalized":"D. M. Porter","authors":["D. M. Porter"],"originalAuth":{"authors":["D. M. Porter"]}},"details":{"uninomial":{"uninomial":"Tribuloideae","rank":"subfam.","parent":"Zygophyllaceae","authorship":{"verbatim":"D.M.Porter","normalized":"D. M. Porter","authors":["D. M. Porter"],"originalAuth":{"authors":["D. M. Porter"]}}}},"words":[{"verbatim":"Zygophyllaceae","normalized":"Zygophyllaceae","wordType":"UNINOMIAL","start":0,"end":14},{"verbatim":"subfam.","normalized":"subfam.","wordType":"RANK","start":15,"end":22},{"verbatim":"Tribuloideae","normalized":"Tribuloideae","wordType":"UNINOMIAL","start":23,"end":35},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"Porter","normalized":"Porter","wordType":"AUTHOR_WORD","start":40,"end":46}],"id":"c60c1ff6-8e9d-5817-b49c-5845a5eaa9f5","parserVersion":"test_version"}
```

Name: Cordia (Adans.) Kuntze sect. Salimori

Canonical: Cordia sect. Salimori

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Cordia (Adans.) Kuntze sect. Salimori","normalized":"Cordia sect. Salimori","canonical":{"stemmed":"Salimori","simple":"Salimori","full":"Cordia sect. Salimori"},"cardinality":1,"details":{"uninomial":{"uninomial":"Salimori","rank":"sect.","parent":"Cordia"}},"words":[{"verbatim":"Cordia","normalized":"Cordia","wordType":"UNINOMIAL","start":0,"end":6},{"verbatim":"Adans.","normalized":"Adans.","wordType":"AUTHOR_WORD","start":8,"end":14},{"verbatim":"Kuntze","normalized":"Kuntze","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"sect.","normalized":"sect.","wordType":"RANK","start":23,"end":28},{"verbatim":"Salimori","normalized":"Salimori","wordType":"UNINOMIAL","start":29,"end":37}],"id":"48d5dbbe-50ff-50ae-a1f8-1cf4b3e2144b","parserVersion":"test_version"}
```

Name: Cordia sect. Salimori (Adans.) Kuntz

Canonical: Cordia sect. Salimori

Authorship: (Adans.) Kuntz

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Cordia sect. Salimori (Adans.) Kuntz","normalized":"Cordia sect. Salimori (Adans.) Kuntz","canonical":{"stemmed":"Salimori","simple":"Salimori","full":"Cordia sect. Salimori"},"cardinality":1,"authorship":{"verbatim":"(Adans.) Kuntz","normalized":"(Adans.) Kuntz","authors":["Adans.","Kuntz"],"originalAuth":{"authors":["Adans."]},"combinationAuth":{"authors":["Kuntz"]}},"details":{"uninomial":{"uninomial":"Salimori","rank":"sect.","parent":"Cordia","authorship":{"verbatim":"(Adans.) Kuntz","normalized":"(Adans.) Kuntz","authors":["Adans.","Kuntz"],"originalAuth":{"authors":["Adans."]},"combinationAuth":{"authors":["Kuntz"]}}}},"words":[{"verbatim":"Cordia","normalized":"Cordia","wordType":"UNINOMIAL","start":0,"end":6},{"verbatim":"sect.","normalized":"sect.","wordType":"RANK","start":7,"end":12},{"verbatim":"Salimori","normalized":"Salimori","wordType":"UNINOMIAL","start":13,"end":21},{"verbatim":"Adans.","normalized":"Adans.","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"Kuntz","normalized":"Kuntz","wordType":"AUTHOR_WORD","start":31,"end":36}],"id":"337ef30d-f5da-5194-8bca-5354b262a05c","parserVersion":"test_version"}
```

Name: Poaceae supertrib. Arundinarodae L.Liu

Canonical: Poaceae supertrib. Arundinarodae

Authorship: L. Liu

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Poaceae supertrib. Arundinarodae L.Liu","normalized":"Poaceae supertrib. Arundinarodae L. Liu","canonical":{"stemmed":"Arundinarodae","simple":"Arundinarodae","full":"Poaceae supertrib. Arundinarodae"},"cardinality":1,"authorship":{"verbatim":"L.Liu","normalized":"L. Liu","authors":["L. Liu"],"originalAuth":{"authors":["L. Liu"]}},"details":{"uninomial":{"uninomial":"Arundinarodae","rank":"supertrib.","parent":"Poaceae","authorship":{"verbatim":"L.Liu","normalized":"L. Liu","authors":["L. Liu"],"originalAuth":{"authors":["L. Liu"]}}}},"words":[{"verbatim":"Poaceae","normalized":"Poaceae","wordType":"UNINOMIAL","start":0,"end":7},{"verbatim":"supertrib.","normalized":"supertrib.","wordType":"RANK","start":8,"end":18},{"verbatim":"Arundinarodae","normalized":"Arundinarodae","wordType":"UNINOMIAL","start":19,"end":32},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"Liu","normalized":"Liu","wordType":"AUTHOR_WORD","start":35,"end":38}],"id":"c589a60b-1273-5b0b-93ea-25919d86647d","parserVersion":"test_version"}
```

Name: Alchemilla subsect. Sericeae A.Plocek

Canonical: Alchemilla subsect. Sericeae

Authorship: A. Plocek

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Alchemilla subsect. Sericeae A.Plocek","normalized":"Alchemilla subsect. Sericeae A. Plocek","canonical":{"stemmed":"Sericeae","simple":"Sericeae","full":"Alchemilla subsect. Sericeae"},"cardinality":1,"authorship":{"verbatim":"A.Plocek","normalized":"A. Plocek","authors":["A. Plocek"],"originalAuth":{"authors":["A. Plocek"]}},"details":{"uninomial":{"uninomial":"Sericeae","rank":"subsect.","parent":"Alchemilla","authorship":{"verbatim":"A.Plocek","normalized":"A. Plocek","authors":["A. Plocek"],"originalAuth":{"authors":["A. Plocek"]}}}},"words":[{"verbatim":"Alchemilla","normalized":"Alchemilla","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"subsect.","normalized":"subsect.","wordType":"RANK","start":11,"end":19},{"verbatim":"Sericeae","normalized":"Sericeae","wordType":"UNINOMIAL","start":20,"end":28},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Plocek","normalized":"Plocek","wordType":"AUTHOR_WORD","start":31,"end":37}],"id":"bedd1b9c-91dd-5ad9-9cd6-0504b85aae30","parserVersion":"test_version"}
```

Name: Hymenophyllum subgen. Hymenoglossum (Presl) R.M.Tryon & A.Tryon

Canonical: Hymenophyllum subgen. Hymenoglossum

Authorship: (Presl) R. M. Tryon & A. Tryon

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Hymenophyllum subgen. Hymenoglossum (Presl) R.M.Tryon \u0026 A.Tryon","normalized":"Hymenophyllum subgen. Hymenoglossum (Presl) R. M. Tryon \u0026 A. Tryon","canonical":{"stemmed":"Hymenoglossum","simple":"Hymenoglossum","full":"Hymenophyllum subgen. Hymenoglossum"},"cardinality":1,"authorship":{"verbatim":"(Presl) R.M.Tryon \u0026 A.Tryon","normalized":"(Presl) R. M. Tryon \u0026 A. Tryon","authors":["Presl","R. M. Tryon","A. Tryon"],"originalAuth":{"authors":["Presl"]},"combinationAuth":{"authors":["R. M. Tryon","A. Tryon"]}},"details":{"uninomial":{"uninomial":"Hymenoglossum","rank":"subgen.","parent":"Hymenophyllum","authorship":{"verbatim":"(Presl) R.M.Tryon \u0026 A.Tryon","normalized":"(Presl) R. M. Tryon \u0026 A. Tryon","authors":["Presl","R. M. Tryon","A. Tryon"],"originalAuth":{"authors":["Presl"]},"combinationAuth":{"authors":["R. M. Tryon","A. Tryon"]}}}},"words":[{"verbatim":"Hymenophyllum","normalized":"Hymenophyllum","wordType":"UNINOMIAL","start":0,"end":13},{"verbatim":"subgen.","normalized":"subgen.","wordType":"RANK","start":14,"end":21},{"verbatim":"Hymenoglossum","normalized":"Hymenoglossum","wordType":"UNINOMIAL","start":22,"end":35},{"verbatim":"Presl","normalized":"Presl","wordType":"AUTHOR_WORD","start":37,"end":42},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":46,"end":48},{"verbatim":"Tryon","normalized":"Tryon","wordType":"AUTHOR_WORD","start":48,"end":53},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":56,"end":58},{"verbatim":"Tryon","normalized":"Tryon","wordType":"AUTHOR_WORD","start":58,"end":63}],"id":"22ea4710-3a2a-5526-a42e-7c7ff508ee79","parserVersion":"test_version"}
```

Name: Pereskia subg. Maihuenia Philippi ex F.A.C.Weber, 1898

Canonical: Pereskia subgen. Maihuenia

Authorship: Philippi ex F. A. C. Weber 1898

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Pereskia subg. Maihuenia Philippi ex F.A.C.Weber, 1898","normalized":"Pereskia subgen. Maihuenia Philippi ex F. A. C. Weber 1898","canonical":{"stemmed":"Maihuenia","simple":"Maihuenia","full":"Pereskia subgen. Maihuenia"},"cardinality":1,"authorship":{"verbatim":"Philippi ex F.A.C.Weber, 1898","normalized":"Philippi ex F. A. C. Weber 1898","year":"1898","authors":["Philippi","F. A. C. Weber"],"originalAuth":{"authors":["Philippi"],"exAuthors":{"authors":["F. A. C. Weber"],"year":{"year":"1898"}}}},"details":{"uninomial":{"uninomial":"Maihuenia","rank":"subgen.","parent":"Pereskia","authorship":{"verbatim":"Philippi ex F.A.C.Weber, 1898","normalized":"Philippi ex F. A. C. Weber 1898","year":"1898","authors":["Philippi","F. A. C. Weber"],"originalAuth":{"authors":["Philippi"],"exAuthors":{"authors":["F. A. C. Weber"],"year":{"year":"1898"}}}}}},"words":[{"verbatim":"Pereskia","normalized":"Pereskia","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"subg.","normalized":"subgen.","wordType":"RANK","start":9,"end":14},{"verbatim":"Maihuenia","normalized":"Maihuenia","wordType":"UNINOMIAL","start":15,"end":24},{"verbatim":"Philippi","normalized":"Philippi","wordType":"AUTHOR_WORD","start":25,"end":33},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"Weber","normalized":"Weber","wordType":"AUTHOR_WORD","start":43,"end":48},{"verbatim":"1898","normalized":"1898","wordType":"YEAR","start":50,"end":54}],"id":"344bd8c1-a4d2-5120-a738-0903aafad63d","parserVersion":"test_version"}
```

Name: Aconitum ser. Tangutica W.T. Wang

Canonical: Aconitum ser. Tangutica

Authorship: W. T. Wang

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aconitum ser. Tangutica W.T. Wang","normalized":"Aconitum ser. Tangutica W. T. Wang","canonical":{"stemmed":"Tangutica","simple":"Tangutica","full":"Aconitum ser. Tangutica"},"cardinality":1,"authorship":{"verbatim":"W.T. Wang","normalized":"W. T. Wang","authors":["W. T. Wang"],"originalAuth":{"authors":["W. T. Wang"]}},"details":{"uninomial":{"uninomial":"Tangutica","rank":"ser.","parent":"Aconitum","authorship":{"verbatim":"W.T. Wang","normalized":"W. T. Wang","authors":["W. T. Wang"],"originalAuth":{"authors":["W. T. Wang"]}}}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"ser.","normalized":"ser.","wordType":"RANK","start":9,"end":13},{"verbatim":"Tangutica","normalized":"Tangutica","wordType":"UNINOMIAL","start":14,"end":23},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"Wang","normalized":"Wang","wordType":"AUTHOR_WORD","start":29,"end":33}],"id":"8f5d7bd0-90a1-556d-a8ef-1a440b157c34","parserVersion":"test_version"}
```

Name: Calathus (Lindrothius) KURNAKOV 1961

Canonical: Calathus subgen. Lindrothius

Authorship: Kurnakov 1961

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Calathus (Lindrothius) KURNAKOV 1961","normalized":"Calathus subgen. Lindrothius Kurnakov 1961","canonical":{"stemmed":"Lindrothius","simple":"Lindrothius","full":"Calathus subgen. Lindrothius"},"cardinality":1,"authorship":{"verbatim":"KURNAKOV 1961","normalized":"Kurnakov 1961","year":"1961","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1961"}}},"details":{"uninomial":{"uninomial":"Lindrothius","rank":"subgen.","parent":"Calathus","authorship":{"verbatim":"KURNAKOV 1961","normalized":"Kurnakov 1961","year":"1961","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1961"}}}}},"words":[{"verbatim":"Calathus","normalized":"Calathus","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"Lindrothius","normalized":"Lindrothius","wordType":"UNINOMIAL","start":10,"end":21},{"verbatim":"KURNAKOV","normalized":"Kurnakov","wordType":"AUTHOR_WORD","start":23,"end":31},{"verbatim":"1961","normalized":"1961","wordType":"YEAR","start":32,"end":36}],"id":"aa113505-61a1-58fe-92f3-8fd511dcfd61","parserVersion":"test_version"}
```

Name: Eucalyptus subser. Regulares Brooker

Canonical: Eucalyptus subser. Regulares

Authorship: Brooker

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Eucalyptus subser. Regulares Brooker","normalized":"Eucalyptus subser. Regulares Brooker","canonical":{"stemmed":"Regulares","simple":"Regulares","full":"Eucalyptus subser. Regulares"},"cardinality":1,"authorship":{"verbatim":"Brooker","normalized":"Brooker","authors":["Brooker"],"originalAuth":{"authors":["Brooker"]}},"details":{"uninomial":{"uninomial":"Regulares","rank":"subser.","parent":"Eucalyptus","authorship":{"verbatim":"Brooker","normalized":"Brooker","authors":["Brooker"],"originalAuth":{"authors":["Brooker"]}}}},"words":[{"verbatim":"Eucalyptus","normalized":"Eucalyptus","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"subser.","normalized":"subser.","wordType":"RANK","start":11,"end":18},{"verbatim":"Regulares","normalized":"Regulares","wordType":"UNINOMIAL","start":19,"end":28},{"verbatim":"Brooker","normalized":"Brooker","wordType":"AUTHOR_WORD","start":29,"end":36}],"id":"783aa15c-f54f-5233-b792-16774a21a34d","parserVersion":"test_version"}
```

Name: Rosa div. Caninae Lindl.

Canonical: Rosa div. Caninae

Authorship: Lindl.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Rosa div. Caninae Lindl.","normalized":"Rosa div. Caninae Lindl.","canonical":{"stemmed":"Caninae","simple":"Caninae","full":"Rosa div. Caninae"},"cardinality":1,"authorship":{"verbatim":"Lindl.","normalized":"Lindl.","authors":["Lindl."],"originalAuth":{"authors":["Lindl."]}},"details":{"uninomial":{"uninomial":"Caninae","rank":"div.","parent":"Rosa","authorship":{"verbatim":"Lindl.","normalized":"Lindl.","authors":["Lindl."],"originalAuth":{"authors":["Lindl."]}}}},"words":[{"verbatim":"Rosa","normalized":"Rosa","wordType":"UNINOMIAL","start":0,"end":4},{"verbatim":"div.","normalized":"div.","wordType":"RANK","start":5,"end":9},{"verbatim":"Caninae","normalized":"Caninae","wordType":"UNINOMIAL","start":10,"end":17},{"verbatim":"Lindl.","normalized":"Lindl.","wordType":"AUTHOR_WORD","start":18,"end":24}],"id":"e48a933f-93e2-5839-aae9-33b83bc046d1","parserVersion":"test_version"}
```

Name: Rosa div Caninae Lindl.

Canonical: Rosa div Caninae

Authorship: Lindl.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Rosa div Caninae Lindl.","normalized":"Rosa div Caninae Lindl.","canonical":{"stemmed":"Caninae","simple":"Caninae","full":"Rosa div Caninae"},"cardinality":1,"authorship":{"verbatim":"Lindl.","normalized":"Lindl.","authors":["Lindl."],"originalAuth":{"authors":["Lindl."]}},"details":{"uninomial":{"uninomial":"Caninae","rank":"div","parent":"Rosa","authorship":{"verbatim":"Lindl.","normalized":"Lindl.","authors":["Lindl."],"originalAuth":{"authors":["Lindl."]}}}},"words":[{"verbatim":"Rosa","normalized":"Rosa","wordType":"UNINOMIAL","start":0,"end":4},{"verbatim":"div","normalized":"div","wordType":"RANK","start":5,"end":8},{"verbatim":"Caninae","normalized":"Caninae","wordType":"UNINOMIAL","start":9,"end":16},{"verbatim":"Lindl.","normalized":"Lindl.","wordType":"AUTHOR_WORD","start":17,"end":23}],"id":"39b7a4e3-9184-5994-bbb8-b1508c420f7e","parserVersion":"test_version"}
```

Name: Aaleniella (Danocythere)

Canonical: Aaleniella subgen. Danocythere

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aaleniella (Danocythere)","normalized":"Aaleniella subgen. Danocythere","canonical":{"stemmed":"Danocythere","simple":"Danocythere","full":"Aaleniella subgen. Danocythere"},"cardinality":1,"details":{"uninomial":{"uninomial":"Danocythere","rank":"subgen.","parent":"Aaleniella"}},"words":[{"verbatim":"Aaleniella","normalized":"Aaleniella","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"Danocythere","normalized":"Danocythere","wordType":"UNINOMIAL","start":12,"end":23}],"id":"8b7eddb1-b9a4-5cca-8fa8-25527e25d8df","parserVersion":"test_version"}
```

### ICN names that look like combined uninomials for ICZN

Name: Clathrotropis (Bentham) Harms in Dalla Torre & Harms, 1901

Canonical: Clathrotropis

Authorship: (Bentham) Harms ex Dalla Torre & Harms 1901

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Clathrotropis (Bentham) Harms in Dalla Torre \u0026 Harms, 1901","normalized":"Clathrotropis (Bentham) Harms ex Dalla Torre \u0026 Harms 1901","canonical":{"stemmed":"Clathrotropis","simple":"Clathrotropis","full":"Clathrotropis"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Bentham) Harms ex Dalla Torre \u0026 Harms 1901","authors":["Bentham","Harms","Dalla Torre"],"originalAuth":{"authors":["Bentham"]},"combinationAuth":{"authors":["Harms"],"exAuthors":{"authors":["Dalla Torre","Harms"],"year":{"year":"1901"}}}},"details":{"uninomial":{"uninomial":"Clathrotropis","authorship":{"verbatim":"","normalized":"(Bentham) Harms ex Dalla Torre \u0026 Harms 1901","authors":["Bentham","Harms","Dalla Torre"],"originalAuth":{"authors":["Bentham"]},"combinationAuth":{"authors":["Harms"],"exAuthors":{"authors":["Dalla Torre","Harms"],"year":{"year":"1901"}}}}}},"words":[{"verbatim":"Clathrotropis","normalized":"Clathrotropis","wordType":"UNINOMIAL","start":0,"end":13},{"verbatim":"Bentham","normalized":"Bentham","wordType":"AUTHOR_WORD","start":15,"end":22},{"verbatim":"Harms","normalized":"Harms","wordType":"AUTHOR_WORD","start":24,"end":29},{"verbatim":"Dalla","normalized":"Dalla","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"Torre","normalized":"Torre","wordType":"AUTHOR_WORD","start":39,"end":44},{"verbatim":"Harms","normalized":"Harms","wordType":"AUTHOR_WORD","start":47,"end":52},{"verbatim":"1901","normalized":"1901","wordType":"YEAR","start":54,"end":58}],"id":"6b730cea-e81b-53ba-a511-caaa233b9b84","parserVersion":"test_version"}
```

Name: Humiriastrum (Urban) Cuatrecasas, 1961

Canonical: Humiriastrum

Authorship: (Urban) Cuatrecasas 1961

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Humiriastrum (Urban) Cuatrecasas, 1961","normalized":"Humiriastrum (Urban) Cuatrecasas 1961","canonical":{"stemmed":"Humiriastrum","simple":"Humiriastrum","full":"Humiriastrum"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Urban) Cuatrecasas 1961","authors":["Urban","Cuatrecasas"],"originalAuth":{"authors":["Urban"]},"combinationAuth":{"authors":["Cuatrecasas"],"year":{"year":"1961"}}},"details":{"uninomial":{"uninomial":"Humiriastrum","authorship":{"verbatim":"","normalized":"(Urban) Cuatrecasas 1961","authors":["Urban","Cuatrecasas"],"originalAuth":{"authors":["Urban"]},"combinationAuth":{"authors":["Cuatrecasas"],"year":{"year":"1961"}}}}},"words":[{"verbatim":"Humiriastrum","normalized":"Humiriastrum","wordType":"UNINOMIAL","start":0,"end":12},{"verbatim":"Urban","normalized":"Urban","wordType":"AUTHOR_WORD","start":14,"end":19},{"verbatim":"Cuatrecasas","normalized":"Cuatrecasas","wordType":"AUTHOR_WORD","start":21,"end":32},{"verbatim":"1961","normalized":"1961","wordType":"YEAR","start":34,"end":38}],"id":"98f8aa31-1cc3-59c2-a4f2-ebf18e0929ab","parserVersion":"test_version"}
```

Name: Pampocactus (Doweld) Doweld

Canonical: Pampocactus

Authorship: (Doweld) Doweld

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Pampocactus (Doweld) Doweld","normalized":"Pampocactus (Doweld) Doweld","canonical":{"stemmed":"Pampocactus","simple":"Pampocactus","full":"Pampocactus"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Doweld) Doweld","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]},"combinationAuth":{"authors":["Doweld"]}},"details":{"uninomial":{"uninomial":"Pampocactus","authorship":{"verbatim":"","normalized":"(Doweld) Doweld","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]},"combinationAuth":{"authors":["Doweld"]}}}},"words":[{"verbatim":"Pampocactus","normalized":"Pampocactus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Doweld","normalized":"Doweld","wordType":"AUTHOR_WORD","start":13,"end":19},{"verbatim":"Doweld","normalized":"Doweld","wordType":"AUTHOR_WORD","start":21,"end":27}],"id":"82494c70-6400-51a3-b786-2a8a747f8305","parserVersion":"test_version"}
```

Name: Pampocactus (Doweld)

Canonical: Pampocactus

Authorship: (Doweld)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Pampocactus (Doweld)","normalized":"Pampocactus (Doweld)","canonical":{"stemmed":"Pampocactus","simple":"Pampocactus","full":"Pampocactus"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Doweld)","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]}},"details":{"uninomial":{"uninomial":"Pampocactus","authorship":{"verbatim":"","normalized":"(Doweld)","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]}}}},"words":[{"verbatim":"Pampocactus","normalized":"Pampocactus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Doweld","normalized":"Doweld","wordType":"AUTHOR_WORD","start":13,"end":19}],"id":"3ed64c9a-ec8a-52c9-a913-eae09b6c71b9","parserVersion":"test_version"}
```

Name: Drepanolejeunea (Spruce) (Steph.)

Canonical: Drepanolejeunea

Authorship: (Spruce)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Drepanolejeunea (Spruce) (Steph.)","normalized":"Drepanolejeunea (Spruce)","canonical":{"stemmed":"Drepanolejeunea","simple":"Drepanolejeunea","full":"Drepanolejeunea"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Spruce)","authors":["Spruce"],"originalAuth":{"authors":["Spruce"]}},"tail":"(Steph.)","details":{"uninomial":{"uninomial":"Drepanolejeunea","authorship":{"verbatim":"","normalized":"(Spruce)","authors":["Spruce"],"originalAuth":{"authors":["Spruce"]}}}},"words":[{"verbatim":"Drepanolejeunea","normalized":"Drepanolejeunea","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"Spruce","normalized":"Spruce","wordType":"AUTHOR_WORD","start":17,"end":23}],"id":"19265c95-0a2b-5e8a-b2c4-478716e9c9ec","parserVersion":"test_version"}
```


### Binomials without authorship

Name: Notopholia corrusca

Canonical: Notopholia corrusca

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Notopholia corrusca","normalized":"Notopholia corrusca","canonical":{"stemmed":"Notopholia corrusc","simple":"Notopholia corrusca","full":"Notopholia corrusca"},"cardinality":2,"details":{"species":{"genus":"Notopholia","species":"corrusca"}},"words":[{"verbatim":"Notopholia","normalized":"Notopholia","wordType":"GENUS","start":0,"end":10},{"verbatim":"corrusca","normalized":"corrusca","wordType":"SPECIES","start":11,"end":19}],"id":"755cef9c-65e4-598d-abf5-4d4a91be9845","parserVersion":"test_version"}
```

Name: Cyathicula scelobelonium

Canonical: Cyathicula scelobelonium

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Cyathicula scelobelonium","normalized":"Cyathicula scelobelonium","canonical":{"stemmed":"Cyathicula scelobeloni","simple":"Cyathicula scelobelonium","full":"Cyathicula scelobelonium"},"cardinality":2,"details":{"species":{"genus":"Cyathicula","species":"scelobelonium"}},"words":[{"verbatim":"Cyathicula","normalized":"Cyathicula","wordType":"GENUS","start":0,"end":10},{"verbatim":"scelobelonium","normalized":"scelobelonium","wordType":"SPECIES","start":11,"end":24}],"id":"21047543-b5ef-5426-b2b4-bc19f3498407","parserVersion":"test_version"}
```

Name: Pseudocercospora     dendrobii

Canonical: Pseudocercospora dendrobii

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora     dendrobii","normalized":"Pseudocercospora dendrobii","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii"}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":21,"end":30}],"id":"5b320aa4-d417-5eda-be2d-83632e0d3624","parserVersion":"test_version"}
```

Name: Cucurbita pepo

Canonical: Cucurbita pepo

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Cucurbita pepo","normalized":"Cucurbita pepo","canonical":{"stemmed":"Cucurbita pep","simple":"Cucurbita pepo","full":"Cucurbita pepo"},"cardinality":2,"details":{"species":{"genus":"Cucurbita","species":"pepo"}},"words":[{"verbatim":"Cucurbita","normalized":"Cucurbita","wordType":"GENUS","start":0,"end":9},{"verbatim":"pepo","normalized":"pepo","wordType":"SPECIES","start":10,"end":14}],"id":"022e85ce-a786-5478-9799-ac2e0f2cc726","parserVersion":"test_version"}
```

Name: Hirsutëlla mâle

Canonical: Hirsutella male

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hirsutëlla mâle","normalized":"Hirsutella male","canonical":{"stemmed":"Hirsutella mal","simple":"Hirsutella male","full":"Hirsutella male"},"cardinality":2,"details":{"species":{"genus":"Hirsutella","species":"male"}},"words":[{"verbatim":"Hirsutëlla","normalized":"Hirsutella","wordType":"GENUS","start":0,"end":10},{"verbatim":"mâle","normalized":"male","wordType":"SPECIES","start":11,"end":15}],"id":"62cc5704-b486-5aba-882c-dc29f5282179","parserVersion":"test_version"}
```

Name: Aëtosaurus ferratus

Canonical: Aetosaurus ferratus

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Aëtosaurus ferratus","normalized":"Aetosaurus ferratus","canonical":{"stemmed":"Aetosaurus ferrat","simple":"Aetosaurus ferratus","full":"Aetosaurus ferratus"},"cardinality":2,"details":{"species":{"genus":"Aetosaurus","species":"ferratus"}},"words":[{"verbatim":"Aëtosaurus","normalized":"Aetosaurus","wordType":"GENUS","start":0,"end":10},{"verbatim":"ferratus","normalized":"ferratus","wordType":"SPECIES","start":11,"end":19}],"id":"9d95ffa0-0203-541f-854a-77ca7ff187fa","parserVersion":"test_version"}
```

Name: Remera cvancarai

Canonical: Remera cvancarai

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Remera cvancarai","normalized":"Remera cvancarai","canonical":{"stemmed":"Remera cuancara","simple":"Remera cvancarai","full":"Remera cvancarai"},"cardinality":2,"details":{"species":{"genus":"Remera","species":"cvancarai"}},"words":[{"verbatim":"Remera","normalized":"Remera","wordType":"GENUS","start":0,"end":6},{"verbatim":"cvancarai","normalized":"cvancarai","wordType":"SPECIES","start":7,"end":16}],"id":"d5d77ab3-2648-5409-a6c7-e3e20d75c38b","parserVersion":"test_version"}
```

### Binomials with authorship

Name: Cymatium raderi D’Attilio & Myers, 1984

Canonical: Cymatium raderi

Authorship: D'Attilio & Myers 1984

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Cymatium raderi D’Attilio \u0026 Myers, 1984","normalized":"Cymatium raderi D'Attilio \u0026 Myers 1984","canonical":{"stemmed":"Cymatium rader","simple":"Cymatium raderi","full":"Cymatium raderi"},"cardinality":2,"authorship":{"verbatim":"D’Attilio \u0026 Myers, 1984","normalized":"D'Attilio \u0026 Myers 1984","year":"1984","authors":["D'Attilio","Myers"],"originalAuth":{"authors":["D'Attilio","Myers"],"year":{"year":"1984"}}},"details":{"species":{"genus":"Cymatium","species":"raderi","authorship":{"verbatim":"D’Attilio \u0026 Myers, 1984","normalized":"D'Attilio \u0026 Myers 1984","year":"1984","authors":["D'Attilio","Myers"],"originalAuth":{"authors":["D'Attilio","Myers"],"year":{"year":"1984"}}}}},"words":[{"verbatim":"Cymatium","normalized":"Cymatium","wordType":"GENUS","start":0,"end":8},{"verbatim":"raderi","normalized":"raderi","wordType":"SPECIES","start":9,"end":15},{"verbatim":"D’Attilio","normalized":"D'Attilio","wordType":"AUTHOR_WORD","start":16,"end":25},{"verbatim":"Myers","normalized":"Myers","wordType":"AUTHOR_WORD","start":28,"end":33},{"verbatim":"1984","normalized":"1984","wordType":"YEAR","start":35,"end":39}],"id":"b3a9e67a-58b7-5aed-a74c-1f2b57b015d0","parserVersion":"test_version"}
```

Name: Melania testudinaria Von dem Busch, 1842

Canonical: Melania testudinaria

Authorship: Von dem Busch 1842

```json
{"parsed":true,"quality":1,"verbatim":"Melania testudinaria Von dem Busch, 1842","normalized":"Melania testudinaria Von dem Busch 1842","canonical":{"stemmed":"Melania testudinar","simple":"Melania testudinaria","full":"Melania testudinaria"},"cardinality":2,"authorship":{"verbatim":"Von dem Busch, 1842","normalized":"Von dem Busch 1842","year":"1842","authors":["Von dem Busch"],"originalAuth":{"authors":["Von dem Busch"],"year":{"year":"1842"}}},"details":{"species":{"genus":"Melania","species":"testudinaria","authorship":{"verbatim":"Von dem Busch, 1842","normalized":"Von dem Busch 1842","year":"1842","authors":["Von dem Busch"],"originalAuth":{"authors":["Von dem Busch"],"year":{"year":"1842"}}}}},"words":[{"verbatim":"Melania","normalized":"Melania","wordType":"GENUS","start":0,"end":7},{"verbatim":"testudinaria","normalized":"testudinaria","wordType":"SPECIES","start":8,"end":20},{"verbatim":"Von","normalized":"Von","wordType":"AUTHOR_WORD","start":21,"end":24},{"verbatim":"dem","normalized":"dem","wordType":"AUTHOR_WORD","start":25,"end":28},{"verbatim":"Busch","normalized":"Busch","wordType":"AUTHOR_WORD","start":29,"end":34},{"verbatim":"1842","normalized":"1842","wordType":"YEAR","start":36,"end":40}],"id":"77b32062-db7e-59e5-9c7d-cc7d8e98c2e9","parserVersion":"test_version"}
```

Name: Cryptopleura farlowiana (J.Agardh) ver Steeg & Jossly

Canonical: Cryptopleura farlowiana

Authorship: (J. Agardh) ver Steeg & Jossly

```json
{"parsed":true,"quality":1,"verbatim":"Cryptopleura farlowiana (J.Agardh) ver Steeg \u0026 Jossly","normalized":"Cryptopleura farlowiana (J. Agardh) ver Steeg \u0026 Jossly","canonical":{"stemmed":"Cryptopleura farlowian","simple":"Cryptopleura farlowiana","full":"Cryptopleura farlowiana"},"cardinality":2,"authorship":{"verbatim":"(J.Agardh) ver Steeg \u0026 Jossly","normalized":"(J. Agardh) ver Steeg \u0026 Jossly","authors":["J. Agardh","ver Steeg","Jossly"],"originalAuth":{"authors":["J. Agardh"]},"combinationAuth":{"authors":["ver Steeg","Jossly"]}},"details":{"species":{"genus":"Cryptopleura","species":"farlowiana","authorship":{"verbatim":"(J.Agardh) ver Steeg \u0026 Jossly","normalized":"(J. Agardh) ver Steeg \u0026 Jossly","authors":["J. Agardh","ver Steeg","Jossly"],"originalAuth":{"authors":["J. Agardh"]},"combinationAuth":{"authors":["ver Steeg","Jossly"]}}}},"words":[{"verbatim":"Cryptopleura","normalized":"Cryptopleura","wordType":"GENUS","start":0,"end":12},{"verbatim":"farlowiana","normalized":"farlowiana","wordType":"SPECIES","start":13,"end":23},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Agardh","normalized":"Agardh","wordType":"AUTHOR_WORD","start":27,"end":33},{"verbatim":"ver","normalized":"ver","wordType":"AUTHOR_WORD","start":35,"end":38},{"verbatim":"Steeg","normalized":"Steeg","wordType":"AUTHOR_WORD","start":39,"end":44},{"verbatim":"Jossly","normalized":"Jossly","wordType":"AUTHOR_WORD","start":47,"end":53}],"id":"f9b3b9e2-b1f9-56bb-b0bf-fa8eab2c03dd","parserVersion":"test_version"}
```

Name: Pyxilla caput avis J.-J.Brun

Canonical: Pyxilla caput avis

Authorship: J.-J. Brun

```json
{"parsed":true,"quality":1,"verbatim":"Pyxilla caput avis J.-J.Brun","normalized":"Pyxilla caput avis J.-J. Brun","canonical":{"stemmed":"Pyxilla caput au","simple":"Pyxilla caput avis","full":"Pyxilla caput avis"},"cardinality":3,"authorship":{"verbatim":"J.-J.Brun","normalized":"J.-J. Brun","authors":["J.-J. Brun"],"originalAuth":{"authors":["J.-J. Brun"]}},"details":{"infraspecies":{"genus":"Pyxilla","species":"caput","infraspecies":[{"value":"avis","authorship":{"verbatim":"J.-J.Brun","normalized":"J.-J. Brun","authors":["J.-J. Brun"],"originalAuth":{"authors":["J.-J. Brun"]}}}]}},"words":[{"verbatim":"Pyxilla","normalized":"Pyxilla","wordType":"GENUS","start":0,"end":7},{"verbatim":"caput","normalized":"caput","wordType":"SPECIES","start":8,"end":13},{"verbatim":"avis","normalized":"avis","wordType":"INFRASPECIES","start":14,"end":18},{"verbatim":"J.-J.","normalized":"J.-J.","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"Brun","normalized":"Brun","wordType":"AUTHOR_WORD","start":24,"end":28}],"id":"f2cea9a2-23df-520c-b8a7-c25e50608676","parserVersion":"test_version"}
```

Name: Muscicapa randi Amadon & duPont, 1970

Canonical: Muscicapa randi

Authorship: Amadon & duPont 1970

```json
{"parsed":true,"quality":1,"verbatim":"Muscicapa randi Amadon \u0026 duPont, 1970","normalized":"Muscicapa randi Amadon \u0026 duPont 1970","canonical":{"stemmed":"Muscicapa rand","simple":"Muscicapa randi","full":"Muscicapa randi"},"cardinality":2,"authorship":{"verbatim":"Amadon \u0026 duPont, 1970","normalized":"Amadon \u0026 duPont 1970","year":"1970","authors":["Amadon","duPont"],"originalAuth":{"authors":["Amadon","duPont"],"year":{"year":"1970"}}},"details":{"species":{"genus":"Muscicapa","species":"randi","authorship":{"verbatim":"Amadon \u0026 duPont, 1970","normalized":"Amadon \u0026 duPont 1970","year":"1970","authors":["Amadon","duPont"],"originalAuth":{"authors":["Amadon","duPont"],"year":{"year":"1970"}}}}},"words":[{"verbatim":"Muscicapa","normalized":"Muscicapa","wordType":"GENUS","start":0,"end":9},{"verbatim":"randi","normalized":"randi","wordType":"SPECIES","start":10,"end":15},{"verbatim":"Amadon","normalized":"Amadon","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"duPont","normalized":"duPont","wordType":"AUTHOR_WORD","start":25,"end":31},{"verbatim":"1970","normalized":"1970","wordType":"YEAR","start":33,"end":37}],"id":"07e1f6ac-ab5f-5354-a690-69ed7a5394fc","parserVersion":"test_version"}
```

Name: Scytalopus alvarezlopezi Stiles, Laverde-R. & Cadena 2017

Canonical: Scytalopus alvarezlopezi

Authorship: Stiles, Laverde-R. & Cadena 2017

```json
{"parsed":true,"quality":1,"verbatim":"Scytalopus alvarezlopezi Stiles, Laverde-R. \u0026 Cadena 2017","normalized":"Scytalopus alvarezlopezi Stiles, Laverde-R. \u0026 Cadena 2017","canonical":{"stemmed":"Scytalopus aluarezlopez","simple":"Scytalopus alvarezlopezi","full":"Scytalopus alvarezlopezi"},"cardinality":2,"authorship":{"verbatim":"Stiles, Laverde-R. \u0026 Cadena 2017","normalized":"Stiles, Laverde-R. \u0026 Cadena 2017","year":"2017","authors":["Stiles","Laverde-R.","Cadena"],"originalAuth":{"authors":["Stiles","Laverde-R.","Cadena"],"year":{"year":"2017"}}},"details":{"species":{"genus":"Scytalopus","species":"alvarezlopezi","authorship":{"verbatim":"Stiles, Laverde-R. \u0026 Cadena 2017","normalized":"Stiles, Laverde-R. \u0026 Cadena 2017","year":"2017","authors":["Stiles","Laverde-R.","Cadena"],"originalAuth":{"authors":["Stiles","Laverde-R.","Cadena"],"year":{"year":"2017"}}}}},"words":[{"verbatim":"Scytalopus","normalized":"Scytalopus","wordType":"GENUS","start":0,"end":10},{"verbatim":"alvarezlopezi","normalized":"alvarezlopezi","wordType":"SPECIES","start":11,"end":24},{"verbatim":"Stiles","normalized":"Stiles","wordType":"AUTHOR_WORD","start":25,"end":31},{"verbatim":"Laverde-R.","normalized":"Laverde-R.","wordType":"AUTHOR_WORD","start":33,"end":43},{"verbatim":"Cadena","normalized":"Cadena","wordType":"AUTHOR_WORD","start":46,"end":52},{"verbatim":"2017","normalized":"2017","wordType":"YEAR","start":53,"end":57}],"id":"bac0e1d6-411e-5d96-ad73-a3db20b9b1a0","parserVersion":"test_version"}
```

Name: Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger & Coronado-G 1967

Canonical: Carabus hendrichsi

Authorship: Bolvar, Pieltain, Rotger & Coronado-G 1967

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger \u0026 Coronado-G 1967","normalized":"Carabus (Tanaocarabus) hendrichsi Bolvar, Pieltain, Rotger \u0026 Coronado-G 1967","canonical":{"stemmed":"Carabus hendrichs","simple":"Carabus hendrichsi","full":"Carabus hendrichsi"},"cardinality":2,"authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado-G 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado-G 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado-G"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado-G"],"year":{"year":"1967"}}},"details":{"species":{"genus":"Carabus","subgenus":"Tanaocarabus","species":"hendrichsi","authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado-G 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado-G 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado-G"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado-G"],"year":{"year":"1967"}}}}},"words":[{"verbatim":"Carabus","normalized":"Carabus","wordType":"GENUS","start":0,"end":7},{"verbatim":"Tanaocarabus","normalized":"Tanaocarabus","wordType":"INFRA_GENUS","start":9,"end":21},{"verbatim":"hendrichsi","normalized":"hendrichsi","wordType":"SPECIES","start":23,"end":33},{"verbatim":"Bolvar","normalized":"Bolvar","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"Pieltain","normalized":"Pieltain","wordType":"AUTHOR_WORD","start":43,"end":51},{"verbatim":"Rotger","normalized":"Rotger","wordType":"AUTHOR_WORD","start":53,"end":59},{"verbatim":"Coronado-G","normalized":"Coronado-G","wordType":"AUTHOR_WORD","start":62,"end":72},{"verbatim":"1967","normalized":"1967","wordType":"YEAR","start":73,"end":77}],"id":"7d2a6355-6f24-54a4-8a49-4c7510a07192","parserVersion":"test_version"}
```

Name: Nemcia epacridoides (Meissner)Crisp

Canonical: Nemcia epacridoides

Authorship: (Meissner) Crisp

```json
{"parsed":true,"quality":1,"verbatim":"Nemcia epacridoides (Meissner)Crisp","normalized":"Nemcia epacridoides (Meissner) Crisp","canonical":{"stemmed":"Nemcia epacridoid","simple":"Nemcia epacridoides","full":"Nemcia epacridoides"},"cardinality":2,"authorship":{"verbatim":"(Meissner)Crisp","normalized":"(Meissner) Crisp","authors":["Meissner","Crisp"],"originalAuth":{"authors":["Meissner"]},"combinationAuth":{"authors":["Crisp"]}},"details":{"species":{"genus":"Nemcia","species":"epacridoides","authorship":{"verbatim":"(Meissner)Crisp","normalized":"(Meissner) Crisp","authors":["Meissner","Crisp"],"originalAuth":{"authors":["Meissner"]},"combinationAuth":{"authors":["Crisp"]}}}},"words":[{"verbatim":"Nemcia","normalized":"Nemcia","wordType":"GENUS","start":0,"end":6},{"verbatim":"epacridoides","normalized":"epacridoides","wordType":"SPECIES","start":7,"end":19},{"verbatim":"Meissner","normalized":"Meissner","wordType":"AUTHOR_WORD","start":21,"end":29},{"verbatim":"Crisp","normalized":"Crisp","wordType":"AUTHOR_WORD","start":30,"end":35}],"id":"6ea9d43f-33c1-5bed-b9a9-edb164966eb6","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh & W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii Goh \u0026 W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh \u0026 W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh \u0026 W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"Goh","normalized":"Goh","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":35,"end":37},{"verbatim":"Hsieh","normalized":"Hsieh","wordType":"AUTHOR_WORD","start":38,"end":43},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":44,"end":48}],"id":"988fd6ba-0221-5b62-a041-fb81addc4465","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh and W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii Goh and W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh and W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh and W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"Goh","normalized":"Goh","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":35,"end":37},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"Hsieh","normalized":"Hsieh","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":46,"end":50}],"id":"4d701dca-8774-5a5e-9378-11f60c0e735c","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh et W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii Goh et W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh et W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh et W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"Goh","normalized":"Goh","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":34,"end":36},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"Hsieh","normalized":"Hsieh","wordType":"AUTHOR_WORD","start":39,"end":44},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":45,"end":49}],"id":"13175b62-b95b-53b7-8d88-1be6fca794ec","parserVersion":"test_version"}
```

Name: Schottera nicaeënsis (J.V. Lamouroux ex Duby) Guiry & Hollenberg

Canonical: Schottera nicaeensis

Authorship: (J. V. Lamouroux ex Duby) Guiry & Hollenberg

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Schottera nicaeënsis (J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"Schottera nicaeensis (J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","canonical":{"stemmed":"Schottera nicaeens","simple":"Schottera nicaeensis","full":"Schottera nicaeensis"},"cardinality":2,"authorship":{"verbatim":"(J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"(J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","authors":["J. V. Lamouroux","Duby","Guiry","Hollenberg"],"originalAuth":{"authors":["J. V. Lamouroux"],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Guiry","Hollenberg"]}},"details":{"species":{"genus":"Schottera","species":"nicaeensis","authorship":{"verbatim":"(J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"(J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","authors":["J. V. Lamouroux","Duby","Guiry","Hollenberg"],"originalAuth":{"authors":["J. V. Lamouroux"],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Guiry","Hollenberg"]}}}},"words":[{"verbatim":"Schottera","normalized":"Schottera","wordType":"GENUS","start":0,"end":9},{"verbatim":"nicaeënsis","normalized":"nicaeensis","wordType":"SPECIES","start":10,"end":20},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"V.","normalized":"V.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"Lamouroux","normalized":"Lamouroux","wordType":"AUTHOR_WORD","start":27,"end":36},{"verbatim":"Duby","normalized":"Duby","wordType":"AUTHOR_WORD","start":40,"end":44},{"verbatim":"Guiry","normalized":"Guiry","wordType":"AUTHOR_WORD","start":46,"end":51},{"verbatim":"Hollenberg","normalized":"Hollenberg","wordType":"AUTHOR_WORD","start":54,"end":64}],"id":"ffeb3703-63e5-5ff3-b296-582c0c3a3373","parserVersion":"test_version"}
```

Name: Laevapex vazi dos Santos, 1989

Canonical: Laevapex vazi

Authorship: dos Santos 1989

```json
{"parsed":true,"quality":1,"verbatim":"Laevapex vazi dos Santos, 1989","normalized":"Laevapex vazi dos Santos 1989","canonical":{"stemmed":"Laevapex uaz","simple":"Laevapex vazi","full":"Laevapex vazi"},"cardinality":2,"authorship":{"verbatim":"dos Santos, 1989","normalized":"dos Santos 1989","year":"1989","authors":["dos Santos"],"originalAuth":{"authors":["dos Santos"],"year":{"year":"1989"}}},"details":{"species":{"genus":"Laevapex","species":"vazi","authorship":{"verbatim":"dos Santos, 1989","normalized":"dos Santos 1989","year":"1989","authors":["dos Santos"],"originalAuth":{"authors":["dos Santos"],"year":{"year":"1989"}}}}},"words":[{"verbatim":"Laevapex","normalized":"Laevapex","wordType":"GENUS","start":0,"end":8},{"verbatim":"vazi","normalized":"vazi","wordType":"SPECIES","start":9,"end":13},{"verbatim":"dos","normalized":"dos","wordType":"AUTHOR_WORD","start":14,"end":17},{"verbatim":"Santos","normalized":"Santos","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"1989","normalized":"1989","wordType":"YEAR","start":26,"end":30}],"id":"34df1cb6-bba1-5115-8e9c-c27df4005291","parserVersion":"test_version"}
```

Name: Periclimenaeus aurae dos Santos, Calado & Araújo, 2008

Canonical: Periclimenaeus aurae

Authorship: dos Santos, Calado & Araújo 2008

```json
{"parsed":true,"quality":1,"verbatim":"Periclimenaeus aurae dos Santos, Calado \u0026 Araújo, 2008","normalized":"Periclimenaeus aurae dos Santos, Calado \u0026 Araújo 2008","canonical":{"stemmed":"Periclimenaeus aur","simple":"Periclimenaeus aurae","full":"Periclimenaeus aurae"},"cardinality":2,"authorship":{"verbatim":"dos Santos, Calado \u0026 Araújo, 2008","normalized":"dos Santos, Calado \u0026 Araújo 2008","year":"2008","authors":["dos Santos","Calado","Araújo"],"originalAuth":{"authors":["dos Santos","Calado","Araújo"],"year":{"year":"2008"}}},"details":{"species":{"genus":"Periclimenaeus","species":"aurae","authorship":{"verbatim":"dos Santos, Calado \u0026 Araújo, 2008","normalized":"dos Santos, Calado \u0026 Araújo 2008","year":"2008","authors":["dos Santos","Calado","Araújo"],"originalAuth":{"authors":["dos Santos","Calado","Araújo"],"year":{"year":"2008"}}}}},"words":[{"verbatim":"Periclimenaeus","normalized":"Periclimenaeus","wordType":"GENUS","start":0,"end":14},{"verbatim":"aurae","normalized":"aurae","wordType":"SPECIES","start":15,"end":20},{"verbatim":"dos","normalized":"dos","wordType":"AUTHOR_WORD","start":21,"end":24},{"verbatim":"Santos","normalized":"Santos","wordType":"AUTHOR_WORD","start":25,"end":31},{"verbatim":"Calado","normalized":"Calado","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"Araújo","normalized":"Araújo","wordType":"AUTHOR_WORD","start":42,"end":48},{"verbatim":"2008","normalized":"2008","wordType":"YEAR","start":50,"end":54}],"id":"261677a4-e52c-5cdf-95f8-a1138404112c","parserVersion":"test_version"}
```

Name: Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012

Canonical: Nototriton matama

Authorship: Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños & Wake 2012

```json
{"parsed":true,"quality":1,"verbatim":"Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","canonical":{"stemmed":"Nototriton matam","simple":"Nototriton matama","full":"Nototriton matama"},"cardinality":2,"authorship":{"verbatim":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","year":"2012","authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"originalAuth":{"authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"year":{"year":"2012"}}},"details":{"species":{"genus":"Nototriton","species":"matama","authorship":{"verbatim":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","year":"2012","authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"originalAuth":{"authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"year":{"year":"2012"}}}}},"words":[{"verbatim":"Nototriton","normalized":"Nototriton","wordType":"GENUS","start":0,"end":10},{"verbatim":"matama","normalized":"matama","wordType":"SPECIES","start":11,"end":17},{"verbatim":"Boza-Oviedo","normalized":"Boza-Oviedo","wordType":"AUTHOR_WORD","start":18,"end":29},{"verbatim":"Rovito","normalized":"Rovito","wordType":"AUTHOR_WORD","start":31,"end":37},{"verbatim":"Chaves","normalized":"Chaves","wordType":"AUTHOR_WORD","start":39,"end":45},{"verbatim":"García-Rodríguez","normalized":"García-Rodríguez","wordType":"AUTHOR_WORD","start":47,"end":63},{"verbatim":"Artavia","normalized":"Artavia","wordType":"AUTHOR_WORD","start":65,"end":72},{"verbatim":"Bolaños","normalized":"Bolaños","wordType":"AUTHOR_WORD","start":74,"end":81},{"verbatim":"Wake","normalized":"Wake","wordType":"AUTHOR_WORD","start":87,"end":91},{"verbatim":"2012","normalized":"2012","wordType":"YEAR","start":93,"end":97}],"id":"49503e24-3297-57c6-bc6e-c1a68a338fd3","parserVersion":"test_version"}
```

Name: Architectonica offlexa Iredale, 1931

Canonical: Architectonica offlexa

Authorship: Iredale 1931

```json
{"parsed":true,"quality":1,"verbatim":"Architectonica offlexa Iredale, 1931","normalized":"Architectonica offlexa Iredale 1931","canonical":{"stemmed":"Architectonica offlex","simple":"Architectonica offlexa","full":"Architectonica offlexa"},"cardinality":2,"authorship":{"verbatim":"Iredale, 1931","normalized":"Iredale 1931","year":"1931","authors":["Iredale"],"originalAuth":{"authors":["Iredale"],"year":{"year":"1931"}}},"details":{"species":{"genus":"Architectonica","species":"offlexa","authorship":{"verbatim":"Iredale, 1931","normalized":"Iredale 1931","year":"1931","authors":["Iredale"],"originalAuth":{"authors":["Iredale"],"year":{"year":"1931"}}}}},"words":[{"verbatim":"Architectonica","normalized":"Architectonica","wordType":"GENUS","start":0,"end":14},{"verbatim":"offlexa","normalized":"offlexa","wordType":"SPECIES","start":15,"end":22},{"verbatim":"Iredale","normalized":"Iredale","wordType":"AUTHOR_WORD","start":23,"end":30},{"verbatim":"1931","normalized":"1931","wordType":"YEAR","start":32,"end":36}],"id":"d8088d2a-6d20-5ef6-9ec8-68753e2e6da0","parserVersion":"test_version"}
```

Name: Maracanda amoena Mc'Lach

Canonical: Maracanda amoena

Authorship: Mc'Lach

```json
{"parsed":true,"quality":1,"verbatim":"Maracanda amoena Mc'Lach","normalized":"Maracanda amoena Mc'Lach","canonical":{"stemmed":"Maracanda amoen","simple":"Maracanda amoena","full":"Maracanda amoena"},"cardinality":2,"authorship":{"verbatim":"Mc'Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}},"details":{"species":{"genus":"Maracanda","species":"amoena","authorship":{"verbatim":"Mc'Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}}}},"words":[{"verbatim":"Maracanda","normalized":"Maracanda","wordType":"GENUS","start":0,"end":9},{"verbatim":"amoena","normalized":"amoena","wordType":"SPECIES","start":10,"end":16},{"verbatim":"Mc'Lach","normalized":"Mc'Lach","wordType":"AUTHOR_WORD","start":17,"end":24}],"id":"b561edfc-29e8-5e8d-8849-60899356be0d","parserVersion":"test_version"}
```

Name: Maracanda amoena Mc’Lach

Canonical: Maracanda amoena

Authorship: Mc'Lach

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Maracanda amoena Mc’Lach","normalized":"Maracanda amoena Mc'Lach","canonical":{"stemmed":"Maracanda amoen","simple":"Maracanda amoena","full":"Maracanda amoena"},"cardinality":2,"authorship":{"verbatim":"Mc’Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}},"details":{"species":{"genus":"Maracanda","species":"amoena","authorship":{"verbatim":"Mc’Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}}}},"words":[{"verbatim":"Maracanda","normalized":"Maracanda","wordType":"GENUS","start":0,"end":9},{"verbatim":"amoena","normalized":"amoena","wordType":"SPECIES","start":10,"end":16},{"verbatim":"Mc’Lach","normalized":"Mc'Lach","wordType":"AUTHOR_WORD","start":17,"end":24}],"id":"98ddd2f7-2f78-5970-adac-677273dc3caf","parserVersion":"test_version"}
```

Name: Tridentella tangeroae Bruce, 198?

Canonical: Tridentella tangeroae

Authorship: Bruce (198?)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with question mark"}],"verbatim":"Tridentella tangeroae Bruce, 198?","normalized":"Tridentella tangeroae Bruce (198?)","canonical":{"stemmed":"Tridentella tangero","simple":"Tridentella tangeroae","full":"Tridentella tangeroae"},"cardinality":2,"authorship":{"verbatim":"Bruce, 198?","normalized":"Bruce (198?)","year":"(198?)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"198?","isApproximate":true}}},"details":{"species":{"genus":"Tridentella","species":"tangeroae","authorship":{"verbatim":"Bruce, 198?","normalized":"Bruce (198?)","year":"(198?)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"198?","isApproximate":true}}}}},"words":[{"verbatim":"Tridentella","normalized":"Tridentella","wordType":"GENUS","start":0,"end":11},{"verbatim":"tangeroae","normalized":"tangeroae","wordType":"SPECIES","start":12,"end":21},{"verbatim":"Bruce","normalized":"Bruce","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"198?","normalized":"198?","wordType":"APPROXIMATE_YEAR","start":29,"end":33}],"id":"179d63c9-bad4-5e61-bf2e-7261b4aa5066","parserVersion":"test_version"}
```

Name: Calobota acanthoclada (Dinter) Boatwr. & B.-E.van Wyk

Canonical: Calobota acanthoclada

Authorship: (Dinter) Boatwr. & B.-E. van Wyk

```json
{"parsed":true,"quality":1,"verbatim":"Calobota acanthoclada (Dinter) Boatwr. \u0026 B.-E.van Wyk","normalized":"Calobota acanthoclada (Dinter) Boatwr. \u0026 B.-E. van Wyk","canonical":{"stemmed":"Calobota acanthoclad","simple":"Calobota acanthoclada","full":"Calobota acanthoclada"},"cardinality":2,"authorship":{"verbatim":"(Dinter) Boatwr. \u0026 B.-E.van Wyk","normalized":"(Dinter) Boatwr. \u0026 B.-E. van Wyk","authors":["Dinter","Boatwr.","B.-E. van Wyk"],"originalAuth":{"authors":["Dinter"]},"combinationAuth":{"authors":["Boatwr.","B.-E. van Wyk"]}},"details":{"species":{"genus":"Calobota","species":"acanthoclada","authorship":{"verbatim":"(Dinter) Boatwr. \u0026 B.-E.van Wyk","normalized":"(Dinter) Boatwr. \u0026 B.-E. van Wyk","authors":["Dinter","Boatwr.","B.-E. van Wyk"],"originalAuth":{"authors":["Dinter"]},"combinationAuth":{"authors":["Boatwr.","B.-E. van Wyk"]}}}},"words":[{"verbatim":"Calobota","normalized":"Calobota","wordType":"GENUS","start":0,"end":8},{"verbatim":"acanthoclada","normalized":"acanthoclada","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Dinter","normalized":"Dinter","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"Boatwr.","normalized":"Boatwr.","wordType":"AUTHOR_WORD","start":31,"end":38},{"verbatim":"B.-E.","normalized":"B.-E.","wordType":"AUTHOR_WORD","start":41,"end":46},{"verbatim":"van","normalized":"van","wordType":"AUTHOR_WORD","start":46,"end":49},{"verbatim":"Wyk","normalized":"Wyk","wordType":"AUTHOR_WORD","start":50,"end":53}],"id":"67a3d99b-d8d6-5f5d-ae6e-b69df693e879","parserVersion":"test_version"}
```

Name: Zanthopsis bispinosa M'Coy, 1849

Canonical: Zanthopsis bispinosa

Authorship: M'Coy 1849

```json
{"parsed":true,"quality":1,"verbatim":"Zanthopsis bispinosa M'Coy, 1849","normalized":"Zanthopsis bispinosa M'Coy 1849","canonical":{"stemmed":"Zanthopsis bispinos","simple":"Zanthopsis bispinosa","full":"Zanthopsis bispinosa"},"cardinality":2,"authorship":{"verbatim":"M'Coy, 1849","normalized":"M'Coy 1849","year":"1849","authors":["M'Coy"],"originalAuth":{"authors":["M'Coy"],"year":{"year":"1849"}}},"details":{"species":{"genus":"Zanthopsis","species":"bispinosa","authorship":{"verbatim":"M'Coy, 1849","normalized":"M'Coy 1849","year":"1849","authors":["M'Coy"],"originalAuth":{"authors":["M'Coy"],"year":{"year":"1849"}}}}},"words":[{"verbatim":"Zanthopsis","normalized":"Zanthopsis","wordType":"GENUS","start":0,"end":10},{"verbatim":"bispinosa","normalized":"bispinosa","wordType":"SPECIES","start":11,"end":20},{"verbatim":"M'Coy","normalized":"M'Coy","wordType":"AUTHOR_WORD","start":21,"end":26},{"verbatim":"1849","normalized":"1849","wordType":"YEAR","start":28,"end":32}],"id":"88b58b88-d8fd-55d9-a9c4-ddd11459820e","parserVersion":"test_version"}
```

Name: Scilla rupestris v.d. Merwe

Canonical: Scilla rupestris

Authorship: v.d. Merwe

```json
{"parsed":true,"quality":1,"verbatim":"Scilla rupestris v.d. Merwe","normalized":"Scilla rupestris v.d. Merwe","canonical":{"stemmed":"Scilla rupestr","simple":"Scilla rupestris","full":"Scilla rupestris"},"cardinality":2,"authorship":{"verbatim":"v.d. Merwe","normalized":"v.d. Merwe","authors":["v.d. Merwe"],"originalAuth":{"authors":["v.d. Merwe"]}},"details":{"species":{"genus":"Scilla","species":"rupestris","authorship":{"verbatim":"v.d. Merwe","normalized":"v.d. Merwe","authors":["v.d. Merwe"],"originalAuth":{"authors":["v.d. Merwe"]}}}},"words":[{"verbatim":"Scilla","normalized":"Scilla","wordType":"GENUS","start":0,"end":6},{"verbatim":"rupestris","normalized":"rupestris","wordType":"SPECIES","start":7,"end":16},{"verbatim":"v.d.","normalized":"v.d.","wordType":"AUTHOR_WORD","start":17,"end":21},{"verbatim":"Merwe","normalized":"Merwe","wordType":"AUTHOR_WORD","start":22,"end":27}],"id":"72ec3a37-8a80-5a82-97dd-b6a67a52d209","parserVersion":"test_version"}
```

Name: Bembix bidentata v.d.L.

Canonical: Bembix bidentata

Authorship: v.d. L.

```json
{"parsed":true,"quality":1,"verbatim":"Bembix bidentata v.d.L.","normalized":"Bembix bidentata v.d. L.","canonical":{"stemmed":"Bembix bidentat","simple":"Bembix bidentata","full":"Bembix bidentata"},"cardinality":2,"authorship":{"verbatim":"v.d.L.","normalized":"v.d. L.","authors":["v.d. L."],"originalAuth":{"authors":["v.d. L."]}},"details":{"species":{"genus":"Bembix","species":"bidentata","authorship":{"verbatim":"v.d.L.","normalized":"v.d. L.","authors":["v.d. L."],"originalAuth":{"authors":["v.d. L."]}}}},"words":[{"verbatim":"Bembix","normalized":"Bembix","wordType":"GENUS","start":0,"end":6},{"verbatim":"bidentata","normalized":"bidentata","wordType":"SPECIES","start":7,"end":16},{"verbatim":"v.d.","normalized":"v.d.","wordType":"AUTHOR_WORD","start":17,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":21,"end":23}],"id":"6f226f43-dfa0-5d61-8a3f-200b2277fcf2","parserVersion":"test_version"}
```

Name: Pompilus cinctellus v. d. L.

Canonical: Pompilus cinctellus

Authorship: v. d. L.

```json
{"parsed":true,"quality":1,"verbatim":"Pompilus cinctellus v. d. L.","normalized":"Pompilus cinctellus v. d. L.","canonical":{"stemmed":"Pompilus cinctell","simple":"Pompilus cinctellus","full":"Pompilus cinctellus"},"cardinality":2,"authorship":{"verbatim":"v. d. L.","normalized":"v. d. L.","authors":["v. d. L."],"originalAuth":{"authors":["v. d. L."]}},"details":{"species":{"genus":"Pompilus","species":"cinctellus","authorship":{"verbatim":"v. d. L.","normalized":"v. d. L.","authors":["v. d. L."],"originalAuth":{"authors":["v. d. L."]}}}},"words":[{"verbatim":"Pompilus","normalized":"Pompilus","wordType":"GENUS","start":0,"end":8},{"verbatim":"cinctellus","normalized":"cinctellus","wordType":"SPECIES","start":9,"end":19},{"verbatim":"v. d.","normalized":"v. d.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":26,"end":28}],"id":"8954c0f2-eab4-561d-9f94-6cebd4f8024d","parserVersion":"test_version"}
```

Name: Setaphis viridis v. d.G.

Canonical: Setaphis viridis

Authorship: v. d. G.

```json
{"parsed":true,"quality":1,"verbatim":"Setaphis viridis v. d.G.","normalized":"Setaphis viridis v. d. G.","canonical":{"stemmed":"Setaphis uirid","simple":"Setaphis viridis","full":"Setaphis viridis"},"cardinality":2,"authorship":{"verbatim":"v. d.G.","normalized":"v. d. G.","authors":["v. d. G."],"originalAuth":{"authors":["v. d. G."]}},"details":{"species":{"genus":"Setaphis","species":"viridis","authorship":{"verbatim":"v. d.G.","normalized":"v. d. G.","authors":["v. d. G."],"originalAuth":{"authors":["v. d. G."]}}}},"words":[{"verbatim":"Setaphis","normalized":"Setaphis","wordType":"GENUS","start":0,"end":8},{"verbatim":"viridis","normalized":"viridis","wordType":"SPECIES","start":9,"end":16},{"verbatim":"v. d.","normalized":"v. d.","wordType":"AUTHOR_WORD","start":17,"end":22},{"verbatim":"G.","normalized":"G.","wordType":"AUTHOR_WORD","start":22,"end":24}],"id":"19792117-31fc-52d7-9990-e89b67c459d3","parserVersion":"test_version"}
```

Name: Coleophora mendica Baldizzone & v. d.Wolf 2000

Canonical: Coleophora mendica

Authorship: Baldizzone & v. d. Wolf 2000

```json
{"parsed":true,"quality":1,"verbatim":"Coleophora mendica Baldizzone \u0026 v. d.Wolf 2000","normalized":"Coleophora mendica Baldizzone \u0026 v. d. Wolf 2000","canonical":{"stemmed":"Coleophora mendic","simple":"Coleophora mendica","full":"Coleophora mendica"},"cardinality":2,"authorship":{"verbatim":"Baldizzone \u0026 v. d.Wolf 2000","normalized":"Baldizzone \u0026 v. d. Wolf 2000","year":"2000","authors":["Baldizzone","v. d. Wolf"],"originalAuth":{"authors":["Baldizzone","v. d. Wolf"],"year":{"year":"2000"}}},"details":{"species":{"genus":"Coleophora","species":"mendica","authorship":{"verbatim":"Baldizzone \u0026 v. d.Wolf 2000","normalized":"Baldizzone \u0026 v. d. Wolf 2000","year":"2000","authors":["Baldizzone","v. d. Wolf"],"originalAuth":{"authors":["Baldizzone","v. d. Wolf"],"year":{"year":"2000"}}}}},"words":[{"verbatim":"Coleophora","normalized":"Coleophora","wordType":"GENUS","start":0,"end":10},{"verbatim":"mendica","normalized":"mendica","wordType":"SPECIES","start":11,"end":18},{"verbatim":"Baldizzone","normalized":"Baldizzone","wordType":"AUTHOR_WORD","start":19,"end":29},{"verbatim":"v. d.","normalized":"v. d.","wordType":"AUTHOR_WORD","start":32,"end":37},{"verbatim":"Wolf","normalized":"Wolf","wordType":"AUTHOR_WORD","start":37,"end":41},{"verbatim":"2000","normalized":"2000","wordType":"YEAR","start":42,"end":46}],"id":"982affab-249b-5858-8ea1-ba226378c233","parserVersion":"test_version"}
```

Name: Psoronaias semigranosa von dem Busch in Philippi, 1845

Canonical: Psoronaias semigranosa

Authorship: von dem Busch ex Philippi 1845

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Psoronaias semigranosa von dem Busch in Philippi, 1845","normalized":"Psoronaias semigranosa von dem Busch ex Philippi 1845","canonical":{"stemmed":"Psoronaias semigranos","simple":"Psoronaias semigranosa","full":"Psoronaias semigranosa"},"cardinality":2,"authorship":{"verbatim":"von dem Busch in Philippi, 1845","normalized":"von dem Busch ex Philippi 1845","year":"1845","authors":["von dem Busch","Philippi"],"originalAuth":{"authors":["von dem Busch"],"exAuthors":{"authors":["Philippi"],"year":{"year":"1845"}}}},"details":{"species":{"genus":"Psoronaias","species":"semigranosa","authorship":{"verbatim":"von dem Busch in Philippi, 1845","normalized":"von dem Busch ex Philippi 1845","year":"1845","authors":["von dem Busch","Philippi"],"originalAuth":{"authors":["von dem Busch"],"exAuthors":{"authors":["Philippi"],"year":{"year":"1845"}}}}}},"words":[{"verbatim":"Psoronaias","normalized":"Psoronaias","wordType":"GENUS","start":0,"end":10},{"verbatim":"semigranosa","normalized":"semigranosa","wordType":"SPECIES","start":11,"end":22},{"verbatim":"von dem","normalized":"von dem","wordType":"AUTHOR_WORD","start":23,"end":30},{"verbatim":"Busch","normalized":"Busch","wordType":"AUTHOR_WORD","start":31,"end":36},{"verbatim":"Philippi","normalized":"Philippi","wordType":"AUTHOR_WORD","start":40,"end":48},{"verbatim":"1845","normalized":"1845","wordType":"YEAR","start":50,"end":54}],"id":"948809ee-be49-598d-a755-fded9ba496c5","parserVersion":"test_version"}
```

Name: Phora sororcula v d Wulp 1871

Canonical: Phora sororcula

Authorship: v d Wulp 1871

```json
{"parsed":true,"quality":1,"verbatim":"Phora sororcula v d Wulp 1871","normalized":"Phora sororcula v d Wulp 1871","canonical":{"stemmed":"Phora sororcul","simple":"Phora sororcula","full":"Phora sororcula"},"cardinality":2,"authorship":{"verbatim":"v d Wulp 1871","normalized":"v d Wulp 1871","year":"1871","authors":["v d Wulp"],"originalAuth":{"authors":["v d Wulp"],"year":{"year":"1871"}}},"details":{"species":{"genus":"Phora","species":"sororcula","authorship":{"verbatim":"v d Wulp 1871","normalized":"v d Wulp 1871","year":"1871","authors":["v d Wulp"],"originalAuth":{"authors":["v d Wulp"],"year":{"year":"1871"}}}}},"words":[{"verbatim":"Phora","normalized":"Phora","wordType":"GENUS","start":0,"end":5},{"verbatim":"sororcula","normalized":"sororcula","wordType":"SPECIES","start":6,"end":15},{"verbatim":"v d","normalized":"v d","wordType":"AUTHOR_WORD","start":16,"end":19},{"verbatim":"Wulp","normalized":"Wulp","wordType":"AUTHOR_WORD","start":20,"end":24},{"verbatim":"1871","normalized":"1871","wordType":"YEAR","start":25,"end":29}],"id":"dad2ef8b-4f74-5de5-844b-29b6ee09ce68","parserVersion":"test_version"}
```

Name: Aeolothrips andalusiacus zur Strassen 1973

Canonical: Aeolothrips andalusiacus

Authorship: zur Strassen 1973

```json
{"parsed":true,"quality":1,"verbatim":"Aeolothrips andalusiacus zur Strassen 1973","normalized":"Aeolothrips andalusiacus zur Strassen 1973","canonical":{"stemmed":"Aeolothrips andalusiac","simple":"Aeolothrips andalusiacus","full":"Aeolothrips andalusiacus"},"cardinality":2,"authorship":{"verbatim":"zur Strassen 1973","normalized":"zur Strassen 1973","year":"1973","authors":["zur Strassen"],"originalAuth":{"authors":["zur Strassen"],"year":{"year":"1973"}}},"details":{"species":{"genus":"Aeolothrips","species":"andalusiacus","authorship":{"verbatim":"zur Strassen 1973","normalized":"zur Strassen 1973","year":"1973","authors":["zur Strassen"],"originalAuth":{"authors":["zur Strassen"],"year":{"year":"1973"}}}}},"words":[{"verbatim":"Aeolothrips","normalized":"Aeolothrips","wordType":"GENUS","start":0,"end":11},{"verbatim":"andalusiacus","normalized":"andalusiacus","wordType":"SPECIES","start":12,"end":24},{"verbatim":"zur","normalized":"zur","wordType":"AUTHOR_WORD","start":25,"end":28},{"verbatim":"Strassen","normalized":"Strassen","wordType":"AUTHOR_WORD","start":29,"end":37},{"verbatim":"1973","normalized":"1973","wordType":"YEAR","start":38,"end":42}],"id":"1e99cbcb-7fc9-5454-a40b-4786d3e35751","parserVersion":"test_version"}
```

Name: Orthosia kindermannii Fischer v. Roslerstamm, 1837

Canonical: Orthosia kindermannii

Authorship: Fischer v. Roslerstamm 1837

```json
{"parsed":true,"quality":1,"verbatim":"Orthosia kindermannii Fischer v. Roslerstamm, 1837","normalized":"Orthosia kindermannii Fischer v. Roslerstamm 1837","canonical":{"stemmed":"Orthosia kindermann","simple":"Orthosia kindermannii","full":"Orthosia kindermannii"},"cardinality":2,"authorship":{"verbatim":"Fischer v. Roslerstamm, 1837","normalized":"Fischer v. Roslerstamm 1837","year":"1837","authors":["Fischer v. Roslerstamm"],"originalAuth":{"authors":["Fischer v. Roslerstamm"],"year":{"year":"1837"}}},"details":{"species":{"genus":"Orthosia","species":"kindermannii","authorship":{"verbatim":"Fischer v. Roslerstamm, 1837","normalized":"Fischer v. Roslerstamm 1837","year":"1837","authors":["Fischer v. Roslerstamm"],"originalAuth":{"authors":["Fischer v. Roslerstamm"],"year":{"year":"1837"}}}}},"words":[{"verbatim":"Orthosia","normalized":"Orthosia","wordType":"GENUS","start":0,"end":8},{"verbatim":"kindermannii","normalized":"kindermannii","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Fischer","normalized":"Fischer","wordType":"AUTHOR_WORD","start":22,"end":29},{"verbatim":"v.","normalized":"v.","wordType":"AUTHOR_WORD","start":30,"end":32},{"verbatim":"Roslerstamm","normalized":"Roslerstamm","wordType":"AUTHOR_WORD","start":33,"end":44},{"verbatim":"1837","normalized":"1837","wordType":"YEAR","start":46,"end":50}],"id":"53abecc3-4083-5cdc-966c-09648fe9383d","parserVersion":"test_version"}
```

Name: Boreophilia nomensis (Casey, 1910)

Canonical: Boreophilia nomensis

Authorship: (Casey 1910)

```json
{"parsed":true,"quality":1,"verbatim":"Boreophilia nomensis (Casey, 1910)","normalized":"Boreophilia nomensis (Casey 1910)","canonical":{"stemmed":"Boreophilia nomens","simple":"Boreophilia nomensis","full":"Boreophilia nomensis"},"cardinality":2,"authorship":{"verbatim":"(Casey, 1910)","normalized":"(Casey 1910)","year":"1910","authors":["Casey"],"originalAuth":{"authors":["Casey"],"year":{"year":"1910"}}},"details":{"species":{"genus":"Boreophilia","species":"nomensis","authorship":{"verbatim":"(Casey, 1910)","normalized":"(Casey 1910)","year":"1910","authors":["Casey"],"originalAuth":{"authors":["Casey"],"year":{"year":"1910"}}}}},"words":[{"verbatim":"Boreophilia","normalized":"Boreophilia","wordType":"GENUS","start":0,"end":11},{"verbatim":"nomensis","normalized":"nomensis","wordType":"SPECIES","start":12,"end":20},{"verbatim":"Casey","normalized":"Casey","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"1910","normalized":"1910","wordType":"YEAR","start":29,"end":33}],"id":"3a0b09db-6e9b-513d-9d10-50b828c504f6","parserVersion":"test_version"}
```

Name: Nereidavus kulkovi Kul'kov in Kul'kov & Obut, 1973

Canonical: Nereidavus kulkovi

Authorship: Kul'kov ex Kul'kov & Obut 1973

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Nereidavus kulkovi Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Nereidavus kulkovi Kul'kov ex Kul'kov \u0026 Obut 1973","canonical":{"stemmed":"Nereidavus kulkou","simple":"Nereidavus kulkovi","full":"Nereidavus kulkovi"},"cardinality":2,"authorship":{"verbatim":"Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Kul'kov ex Kul'kov \u0026 Obut 1973","year":"1973","authors":["Kul'kov","Obut"],"originalAuth":{"authors":["Kul'kov"],"exAuthors":{"authors":["Kul'kov","Obut"],"year":{"year":"1973"}}}},"details":{"species":{"genus":"Nereidavus","species":"kulkovi","authorship":{"verbatim":"Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Kul'kov ex Kul'kov \u0026 Obut 1973","year":"1973","authors":["Kul'kov","Obut"],"originalAuth":{"authors":["Kul'kov"],"exAuthors":{"authors":["Kul'kov","Obut"],"year":{"year":"1973"}}}}}},"words":[{"verbatim":"Nereidavus","normalized":"Nereidavus","wordType":"GENUS","start":0,"end":10},{"verbatim":"kulkovi","normalized":"kulkovi","wordType":"SPECIES","start":11,"end":18},{"verbatim":"Kul'kov","normalized":"Kul'kov","wordType":"AUTHOR_WORD","start":19,"end":26},{"verbatim":"Kul'kov","normalized":"Kul'kov","wordType":"AUTHOR_WORD","start":30,"end":37},{"verbatim":"Obut","normalized":"Obut","wordType":"AUTHOR_WORD","start":40,"end":44},{"verbatim":"1973","normalized":"1973","wordType":"YEAR","start":46,"end":50}],"id":"4aa8305f-884f-5515-9bdc-f586e037028c","parserVersion":"test_version"}
```

Name: Xylaria potentillae A S. Xu

Canonical: Xylaria potentillae

Authorship: A S. Xu

```json
{"parsed":true,"quality":1,"verbatim":"Xylaria potentillae A S. Xu","normalized":"Xylaria potentillae A S. Xu","canonical":{"stemmed":"Xylaria potentill","simple":"Xylaria potentillae","full":"Xylaria potentillae"},"cardinality":2,"authorship":{"verbatim":"A S. Xu","normalized":"A S. Xu","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"]}},"details":{"species":{"genus":"Xylaria","species":"potentillae","authorship":{"verbatim":"A S. Xu","normalized":"A S. Xu","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"]}}}},"words":[{"verbatim":"Xylaria","normalized":"Xylaria","wordType":"GENUS","start":0,"end":7},{"verbatim":"potentillae","normalized":"potentillae","wordType":"SPECIES","start":8,"end":19},{"verbatim":"A","normalized":"A","wordType":"AUTHOR_WORD","start":20,"end":21},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"Xu","normalized":"Xu","wordType":"AUTHOR_WORD","start":25,"end":27}],"id":"6bc4bb61-e0b9-5c22-a9b6-46c45757f2c2","parserVersion":"test_version"}
```

Name: Pseudocyrtopora el Hajjaji 1987

Canonical: Pseudocyrtopora

Authorship: el Hajjaji 1987

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocyrtopora el Hajjaji 1987","normalized":"Pseudocyrtopora el Hajjaji 1987","canonical":{"stemmed":"Pseudocyrtopora","simple":"Pseudocyrtopora","full":"Pseudocyrtopora"},"cardinality":1,"authorship":{"verbatim":"el Hajjaji 1987","normalized":"el Hajjaji 1987","year":"1987","authors":["el Hajjaji"],"originalAuth":{"authors":["el Hajjaji"],"year":{"year":"1987"}}},"details":{"uninomial":{"uninomial":"Pseudocyrtopora","authorship":{"verbatim":"el Hajjaji 1987","normalized":"el Hajjaji 1987","year":"1987","authors":["el Hajjaji"],"originalAuth":{"authors":["el Hajjaji"],"year":{"year":"1987"}}}}},"words":[{"verbatim":"Pseudocyrtopora","normalized":"Pseudocyrtopora","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"el","normalized":"el","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"Hajjaji","normalized":"Hajjaji","wordType":"AUTHOR_WORD","start":19,"end":26},{"verbatim":"1987","normalized":"1987","wordType":"YEAR","start":27,"end":31}],"id":"61db186c-cbf4-5949-9fd1-79efe7157873","parserVersion":"test_version"}
```

Name: Geositta poeciloptera (zu Wied-Neuwied, 1830)

Canonical: Geositta poeciloptera

Authorship: (zu Wied-Neuwied 1830)

```json
{"parsed":true,"quality":1,"verbatim":"Geositta poeciloptera (zu Wied-Neuwied, 1830)","normalized":"Geositta poeciloptera (zu Wied-Neuwied 1830)","canonical":{"stemmed":"Geositta poecilopter","simple":"Geositta poeciloptera","full":"Geositta poeciloptera"},"cardinality":2,"authorship":{"verbatim":"(zu Wied-Neuwied, 1830)","normalized":"(zu Wied-Neuwied 1830)","year":"1830","authors":["zu Wied-Neuwied"],"originalAuth":{"authors":["zu Wied-Neuwied"],"year":{"year":"1830"}}},"details":{"species":{"genus":"Geositta","species":"poeciloptera","authorship":{"verbatim":"(zu Wied-Neuwied, 1830)","normalized":"(zu Wied-Neuwied 1830)","year":"1830","authors":["zu Wied-Neuwied"],"originalAuth":{"authors":["zu Wied-Neuwied"],"year":{"year":"1830"}}}}},"words":[{"verbatim":"Geositta","normalized":"Geositta","wordType":"GENUS","start":0,"end":8},{"verbatim":"poeciloptera","normalized":"poeciloptera","wordType":"SPECIES","start":9,"end":21},{"verbatim":"zu","normalized":"zu","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"Wied-Neuwied","normalized":"Wied-Neuwied","wordType":"AUTHOR_WORD","start":26,"end":38},{"verbatim":"1830","normalized":"1830","wordType":"YEAR","start":40,"end":44}],"id":"c2abf205-a19a-5bf1-9a95-668101143dd8","parserVersion":"test_version"}
```

Name: Abacetus laevicollis de Chaudoir, 1869

Canonical: Abacetus laevicollis

Authorship: de Chaudoir 1869

```json
{"parsed":true,"quality":1,"verbatim":"Abacetus laevicollis de Chaudoir, 1869","normalized":"Abacetus laevicollis de Chaudoir 1869","canonical":{"stemmed":"Abacetus laeuicoll","simple":"Abacetus laevicollis","full":"Abacetus laevicollis"},"cardinality":2,"authorship":{"verbatim":"de Chaudoir, 1869","normalized":"de Chaudoir 1869","year":"1869","authors":["de Chaudoir"],"originalAuth":{"authors":["de Chaudoir"],"year":{"year":"1869"}}},"details":{"species":{"genus":"Abacetus","species":"laevicollis","authorship":{"verbatim":"de Chaudoir, 1869","normalized":"de Chaudoir 1869","year":"1869","authors":["de Chaudoir"],"originalAuth":{"authors":["de Chaudoir"],"year":{"year":"1869"}}}}},"words":[{"verbatim":"Abacetus","normalized":"Abacetus","wordType":"GENUS","start":0,"end":8},{"verbatim":"laevicollis","normalized":"laevicollis","wordType":"SPECIES","start":9,"end":20},{"verbatim":"de","normalized":"de","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"Chaudoir","normalized":"Chaudoir","wordType":"AUTHOR_WORD","start":24,"end":32},{"verbatim":"1869","normalized":"1869","wordType":"YEAR","start":34,"end":38}],"id":"8d81b939-695f-5a38-86c7-0f6efd1cacf3","parserVersion":"test_version"}
```

Name: Gastrosericus eremorum von Beaumont 1955

Canonical: Gastrosericus eremorum

Authorship: von Beaumont 1955

```json
{"parsed":true,"quality":1,"verbatim":"Gastrosericus eremorum von Beaumont 1955","normalized":"Gastrosericus eremorum von Beaumont 1955","canonical":{"stemmed":"Gastrosericus eremor","simple":"Gastrosericus eremorum","full":"Gastrosericus eremorum"},"cardinality":2,"authorship":{"verbatim":"von Beaumont 1955","normalized":"von Beaumont 1955","year":"1955","authors":["von Beaumont"],"originalAuth":{"authors":["von Beaumont"],"year":{"year":"1955"}}},"details":{"species":{"genus":"Gastrosericus","species":"eremorum","authorship":{"verbatim":"von Beaumont 1955","normalized":"von Beaumont 1955","year":"1955","authors":["von Beaumont"],"originalAuth":{"authors":["von Beaumont"],"year":{"year":"1955"}}}}},"words":[{"verbatim":"Gastrosericus","normalized":"Gastrosericus","wordType":"GENUS","start":0,"end":13},{"verbatim":"eremorum","normalized":"eremorum","wordType":"SPECIES","start":14,"end":22},{"verbatim":"von","normalized":"von","wordType":"AUTHOR_WORD","start":23,"end":26},{"verbatim":"Beaumont","normalized":"Beaumont","wordType":"AUTHOR_WORD","start":27,"end":35},{"verbatim":"1955","normalized":"1955","wordType":"YEAR","start":36,"end":40}],"id":"98df7228-03ef-511c-9f2d-7f91e10c2af5","parserVersion":"test_version"}
```

Name: Agaricus squamula Berk. & M.A. Curtis 1860

Canonical: Agaricus squamula

Authorship: Berk. & M. A. Curtis 1860

```json
{"parsed":true,"quality":1,"verbatim":"Agaricus squamula Berk. \u0026 M.A. Curtis 1860","normalized":"Agaricus squamula Berk. \u0026 M. A. Curtis 1860","canonical":{"stemmed":"Agaricus squamul","simple":"Agaricus squamula","full":"Agaricus squamula"},"cardinality":2,"authorship":{"verbatim":"Berk. \u0026 M.A. Curtis 1860","normalized":"Berk. \u0026 M. A. Curtis 1860","year":"1860","authors":["Berk.","M. A. Curtis"],"originalAuth":{"authors":["Berk.","M. A. Curtis"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Agaricus","species":"squamula","authorship":{"verbatim":"Berk. \u0026 M.A. Curtis 1860","normalized":"Berk. \u0026 M. A. Curtis 1860","year":"1860","authors":["Berk.","M. A. Curtis"],"originalAuth":{"authors":["Berk.","M. A. Curtis"],"year":{"year":"1860"}}}}},"words":[{"verbatim":"Agaricus","normalized":"Agaricus","wordType":"GENUS","start":0,"end":8},{"verbatim":"squamula","normalized":"squamula","wordType":"SPECIES","start":9,"end":17},{"verbatim":"Berk.","normalized":"Berk.","wordType":"AUTHOR_WORD","start":18,"end":23},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"Curtis","normalized":"Curtis","wordType":"AUTHOR_WORD","start":31,"end":37},{"verbatim":"1860","normalized":"1860","wordType":"YEAR","start":38,"end":42}],"id":"153b8745-887a-56ba-ad4a-69c10b0ad513","parserVersion":"test_version"}
```

Name: Peltula coriacea Büdel, Henssen & Wessels 1986

Canonical: Peltula coriacea

Authorship: Büdel, Henssen & Wessels 1986

```json
{"parsed":true,"quality":1,"verbatim":"Peltula coriacea Büdel, Henssen \u0026 Wessels 1986","normalized":"Peltula coriacea Büdel, Henssen \u0026 Wessels 1986","canonical":{"stemmed":"Peltula coriace","simple":"Peltula coriacea","full":"Peltula coriacea"},"cardinality":2,"authorship":{"verbatim":"Büdel, Henssen \u0026 Wessels 1986","normalized":"Büdel, Henssen \u0026 Wessels 1986","year":"1986","authors":["Büdel","Henssen","Wessels"],"originalAuth":{"authors":["Büdel","Henssen","Wessels"],"year":{"year":"1986"}}},"details":{"species":{"genus":"Peltula","species":"coriacea","authorship":{"verbatim":"Büdel, Henssen \u0026 Wessels 1986","normalized":"Büdel, Henssen \u0026 Wessels 1986","year":"1986","authors":["Büdel","Henssen","Wessels"],"originalAuth":{"authors":["Büdel","Henssen","Wessels"],"year":{"year":"1986"}}}}},"words":[{"verbatim":"Peltula","normalized":"Peltula","wordType":"GENUS","start":0,"end":7},{"verbatim":"coriacea","normalized":"coriacea","wordType":"SPECIES","start":8,"end":16},{"verbatim":"Büdel","normalized":"Büdel","wordType":"AUTHOR_WORD","start":17,"end":22},{"verbatim":"Henssen","normalized":"Henssen","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"Wessels","normalized":"Wessels","wordType":"AUTHOR_WORD","start":34,"end":41},{"verbatim":"1986","normalized":"1986","wordType":"YEAR","start":42,"end":46}],"id":"081f5751-4042-597e-bccc-788754ce0248","parserVersion":"test_version"}
```

Name: Tuber liui A S. Xu 1999

Canonical: Tuber liui

Authorship: A S. Xu 1999

```json
{"parsed":true,"quality":1,"verbatim":"Tuber liui A S. Xu 1999","normalized":"Tuber liui A S. Xu 1999","canonical":{"stemmed":"Tuber liu","simple":"Tuber liui","full":"Tuber liui"},"cardinality":2,"authorship":{"verbatim":"A S. Xu 1999","normalized":"A S. Xu 1999","year":"1999","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"],"year":{"year":"1999"}}},"details":{"species":{"genus":"Tuber","species":"liui","authorship":{"verbatim":"A S. Xu 1999","normalized":"A S. Xu 1999","year":"1999","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"],"year":{"year":"1999"}}}}},"words":[{"verbatim":"Tuber","normalized":"Tuber","wordType":"GENUS","start":0,"end":5},{"verbatim":"liui","normalized":"liui","wordType":"SPECIES","start":6,"end":10},{"verbatim":"A","normalized":"A","wordType":"AUTHOR_WORD","start":11,"end":12},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":13,"end":15},{"verbatim":"Xu","normalized":"Xu","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"1999","normalized":"1999","wordType":"YEAR","start":19,"end":23}],"id":"4c79eb26-ae4c-5f4a-b5c5-07722ef1fa4f","parserVersion":"test_version"}
```

Name: Lecanora wetmorei Śliwa 2004

Canonical: Lecanora wetmorei

Authorship: Śliwa 2004

```json
{"parsed":true,"quality":1,"verbatim":"Lecanora wetmorei Śliwa 2004","normalized":"Lecanora wetmorei Śliwa 2004","canonical":{"stemmed":"Lecanora wetmore","simple":"Lecanora wetmorei","full":"Lecanora wetmorei"},"cardinality":2,"authorship":{"verbatim":"Śliwa 2004","normalized":"Śliwa 2004","year":"2004","authors":["Śliwa"],"originalAuth":{"authors":["Śliwa"],"year":{"year":"2004"}}},"details":{"species":{"genus":"Lecanora","species":"wetmorei","authorship":{"verbatim":"Śliwa 2004","normalized":"Śliwa 2004","year":"2004","authors":["Śliwa"],"originalAuth":{"authors":["Śliwa"],"year":{"year":"2004"}}}}},"words":[{"verbatim":"Lecanora","normalized":"Lecanora","wordType":"GENUS","start":0,"end":8},{"verbatim":"wetmorei","normalized":"wetmorei","wordType":"SPECIES","start":9,"end":17},{"verbatim":"Śliwa","normalized":"Śliwa","wordType":"AUTHOR_WORD","start":18,"end":23},{"verbatim":"2004","normalized":"2004","wordType":"YEAR","start":24,"end":28}],"id":"50e874e9-f807-5446-a416-ca459475b1db","parserVersion":"test_version"}
```

Name: Vachonobisium troglophilum Vitali-di Castri, 1963

Canonical: Vachonobisium troglophilum

Authorship: Vitali-di Castri 1963

```json
{"parsed":true,"quality":1,"verbatim":"Vachonobisium troglophilum Vitali-di Castri, 1963","normalized":"Vachonobisium troglophilum Vitali-di Castri 1963","canonical":{"stemmed":"Vachonobisium troglophil","simple":"Vachonobisium troglophilum","full":"Vachonobisium troglophilum"},"cardinality":2,"authorship":{"verbatim":"Vitali-di Castri, 1963","normalized":"Vitali-di Castri 1963","year":"1963","authors":["Vitali-di Castri"],"originalAuth":{"authors":["Vitali-di Castri"],"year":{"year":"1963"}}},"details":{"species":{"genus":"Vachonobisium","species":"troglophilum","authorship":{"verbatim":"Vitali-di Castri, 1963","normalized":"Vitali-di Castri 1963","year":"1963","authors":["Vitali-di Castri"],"originalAuth":{"authors":["Vitali-di Castri"],"year":{"year":"1963"}}}}},"words":[{"verbatim":"Vachonobisium","normalized":"Vachonobisium","wordType":"GENUS","start":0,"end":13},{"verbatim":"troglophilum","normalized":"troglophilum","wordType":"SPECIES","start":14,"end":26},{"verbatim":"Vitali-di","normalized":"Vitali-di","wordType":"AUTHOR_WORD","start":27,"end":36},{"verbatim":"Castri","normalized":"Castri","wordType":"AUTHOR_WORD","start":37,"end":43},{"verbatim":"1963","normalized":"1963","wordType":"YEAR","start":45,"end":49}],"id":"97424f96-2408-53b6-a6bf-a26613eec14c","parserVersion":"test_version"}
```

Name: Hyalesthes angustula Horvßth, 1909

Canonical: Hyalesthes angustula

Authorship: Horvßth 1909

```json
{"parsed":true,"quality":1,"verbatim":"Hyalesthes angustula Horvßth, 1909","normalized":"Hyalesthes angustula Horvßth 1909","canonical":{"stemmed":"Hyalesthes angustul","simple":"Hyalesthes angustula","full":"Hyalesthes angustula"},"cardinality":2,"authorship":{"verbatim":"Horvßth, 1909","normalized":"Horvßth 1909","year":"1909","authors":["Horvßth"],"originalAuth":{"authors":["Horvßth"],"year":{"year":"1909"}}},"details":{"species":{"genus":"Hyalesthes","species":"angustula","authorship":{"verbatim":"Horvßth, 1909","normalized":"Horvßth 1909","year":"1909","authors":["Horvßth"],"originalAuth":{"authors":["Horvßth"],"year":{"year":"1909"}}}}},"words":[{"verbatim":"Hyalesthes","normalized":"Hyalesthes","wordType":"GENUS","start":0,"end":10},{"verbatim":"angustula","normalized":"angustula","wordType":"SPECIES","start":11,"end":20},{"verbatim":"Horvßth","normalized":"Horvßth","wordType":"AUTHOR_WORD","start":21,"end":28},{"verbatim":"1909","normalized":"1909","wordType":"YEAR","start":30,"end":34}],"id":"02058420-6623-5c22-b5ae-bc6a576f72fe","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl (1935h)

Canonical: Platypus bicaudatulus

Authorship: Schedl (1935)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Platypus bicaudatulus Schedl (1935h)","normalized":"Platypus bicaudatulus Schedl (1935)","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl (1935h)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl (1935h)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"1935h","normalized":"1935","wordType":"APPROXIMATE_YEAR","start":30,"end":35}],"id":"5bf2e3f3-46dc-5138-a912-0e0ab2fdb22d","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl (1935)

Canonical: Platypus bicaudatulus

Authorship: Schedl (1935)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Platypus bicaudatulus Schedl (1935)","normalized":"Platypus bicaudatulus Schedl (1935)","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl (1935)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl (1935)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"1935","normalized":"1935","wordType":"APPROXIMATE_YEAR","start":30,"end":34}],"id":"c13ffa95-76e8-5ad1-aec6-311d65dc4dc0","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl 1935

```json
{"parsed":true,"quality":1,"verbatim":"Platypus bicaudatulus Schedl 1935","normalized":"Platypus bicaudatulus Schedl 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl 1935","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl 1935","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"1935","normalized":"1935","wordType":"YEAR","start":29,"end":33}],"id":"d192a4f8-424f-5eba-affb-9855b153ff53","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl, 1935h

Canonical: Platypus bicaudatulus

Authorship: Schedl 1935

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"}],"verbatim":"Platypus bicaudatulus Schedl, 1935h","normalized":"Platypus bicaudatulus Schedl 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl, 1935h","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl, 1935h","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"1935h","normalized":"1935","wordType":"YEAR","start":30,"end":35}],"id":"2f3b49aa-7d42-557b-9949-41df0e6059e8","parserVersion":"test_version"}
```

Name: Rotalina cultrata d'Orb. 1840

Canonical: Rotalina cultrata

Authorship: d'Orb. 1840

```json
{"parsed":true,"quality":1,"verbatim":"Rotalina cultrata d'Orb. 1840","normalized":"Rotalina cultrata d'Orb. 1840","canonical":{"stemmed":"Rotalina cultrat","simple":"Rotalina cultrata","full":"Rotalina cultrata"},"cardinality":2,"authorship":{"verbatim":"d'Orb. 1840","normalized":"d'Orb. 1840","year":"1840","authors":["d'Orb."],"originalAuth":{"authors":["d'Orb."],"year":{"year":"1840"}}},"details":{"species":{"genus":"Rotalina","species":"cultrata","authorship":{"verbatim":"d'Orb. 1840","normalized":"d'Orb. 1840","year":"1840","authors":["d'Orb."],"originalAuth":{"authors":["d'Orb."],"year":{"year":"1840"}}}}},"words":[{"verbatim":"Rotalina","normalized":"Rotalina","wordType":"GENUS","start":0,"end":8},{"verbatim":"cultrata","normalized":"cultrata","wordType":"SPECIES","start":9,"end":17},{"verbatim":"d'Orb.","normalized":"d'Orb.","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"1840","normalized":"1840","wordType":"YEAR","start":25,"end":29}],"id":"085048a9-a6b8-525e-95ad-ae715b8c00ca","parserVersion":"test_version"}
```

Name: Stylosanthes guianensis (Aubl.) Sw. var. robusta L.'t Mannetje

Canonical: Stylosanthes guianensis var. robusta

Authorship: L. 't Mannetje

```json
{"parsed":true,"quality":1,"verbatim":"Stylosanthes guianensis (Aubl.) Sw. var. robusta L.'t Mannetje","normalized":"Stylosanthes guianensis (Aubl.) Sw. var. robusta L. 't Mannetje","canonical":{"stemmed":"Stylosanthes guianens robust","simple":"Stylosanthes guianensis robusta","full":"Stylosanthes guianensis var. robusta"},"cardinality":3,"authorship":{"verbatim":"L.'t Mannetje","normalized":"L. 't Mannetje","authors":["L. 't Mannetje"],"originalAuth":{"authors":["L. 't Mannetje"]}},"details":{"infraspecies":{"genus":"Stylosanthes","species":"guianensis","authorship":{"verbatim":"(Aubl.) Sw.","normalized":"(Aubl.) Sw.","authors":["Aubl.","Sw."],"originalAuth":{"authors":["Aubl."]},"combinationAuth":{"authors":["Sw."]}},"infraspecies":[{"value":"robusta","rank":"var.","authorship":{"verbatim":"L.'t Mannetje","normalized":"L. 't Mannetje","authors":["L. 't Mannetje"],"originalAuth":{"authors":["L. 't Mannetje"]}}}]}},"words":[{"verbatim":"Stylosanthes","normalized":"Stylosanthes","wordType":"GENUS","start":0,"end":12},{"verbatim":"guianensis","normalized":"guianensis","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Aubl.","normalized":"Aubl.","wordType":"AUTHOR_WORD","start":25,"end":30},{"verbatim":"Sw.","normalized":"Sw.","wordType":"AUTHOR_WORD","start":32,"end":35},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":36,"end":40},{"verbatim":"robusta","normalized":"robusta","wordType":"INFRASPECIES","start":41,"end":48},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"'t","normalized":"'t","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"Mannetje","normalized":"Mannetje","wordType":"AUTHOR_WORD","start":54,"end":62}],"id":"fa16f59c-69a2-50cc-a4f6-bf4e8891eb9a","parserVersion":"test_version"}
```

Name: Doxander vittatus entropi (Man in 't Veld & Visser, 1993)

Canonical: Doxander vittatus entropi

Authorship: (Man ex 't Veld & Visser 1993)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Doxander vittatus entropi (Man in 't Veld \u0026 Visser, 1993)","normalized":"Doxander vittatus entropi (Man ex 't Veld \u0026 Visser 1993)","canonical":{"stemmed":"Doxander uittat entrop","simple":"Doxander vittatus entropi","full":"Doxander vittatus entropi"},"cardinality":3,"authorship":{"verbatim":"(Man in 't Veld \u0026 Visser, 1993)","normalized":"(Man ex 't Veld \u0026 Visser 1993)","year":"1993","authors":["Man","'t Veld","Visser"],"originalAuth":{"authors":["Man"],"exAuthors":{"authors":["'t Veld","Visser"],"year":{"year":"1993"}}}},"details":{"infraspecies":{"genus":"Doxander","species":"vittatus","infraspecies":[{"value":"entropi","authorship":{"verbatim":"(Man in 't Veld \u0026 Visser, 1993)","normalized":"(Man ex 't Veld \u0026 Visser 1993)","year":"1993","authors":["Man","'t Veld","Visser"],"originalAuth":{"authors":["Man"],"exAuthors":{"authors":["'t Veld","Visser"],"year":{"year":"1993"}}}}}]}},"words":[{"verbatim":"Doxander","normalized":"Doxander","wordType":"GENUS","start":0,"end":8},{"verbatim":"vittatus","normalized":"vittatus","wordType":"SPECIES","start":9,"end":17},{"verbatim":"entropi","normalized":"entropi","wordType":"INFRASPECIES","start":18,"end":25},{"verbatim":"Man","normalized":"Man","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"'t","normalized":"'t","wordType":"AUTHOR_WORD","start":34,"end":36},{"verbatim":"Veld","normalized":"Veld","wordType":"AUTHOR_WORD","start":37,"end":41},{"verbatim":"Visser","normalized":"Visser","wordType":"AUTHOR_WORD","start":44,"end":50},{"verbatim":"1993","normalized":"1993","wordType":"YEAR","start":52,"end":56}],"id":"1b3da2cb-82db-511d-86f5-4421966e3b65","parserVersion":"test_version"}
```

Name: Elaeagnus triflora Roxb. var. brevilimbatus E.'t Hart

Canonical: Elaeagnus triflora var. brevilimbatus

Authorship: E. 't Hart

```json
{"parsed":true,"quality":1,"verbatim":"Elaeagnus triflora Roxb. var. brevilimbatus E.'t Hart","normalized":"Elaeagnus triflora Roxb. var. brevilimbatus E. 't Hart","canonical":{"stemmed":"Elaeagnus triflor breuilimbat","simple":"Elaeagnus triflora brevilimbatus","full":"Elaeagnus triflora var. brevilimbatus"},"cardinality":3,"authorship":{"verbatim":"E.'t Hart","normalized":"E. 't Hart","authors":["E. 't Hart"],"originalAuth":{"authors":["E. 't Hart"]}},"details":{"infraspecies":{"genus":"Elaeagnus","species":"triflora","authorship":{"verbatim":"Roxb.","normalized":"Roxb.","authors":["Roxb."],"originalAuth":{"authors":["Roxb."]}},"infraspecies":[{"value":"brevilimbatus","rank":"var.","authorship":{"verbatim":"E.'t Hart","normalized":"E. 't Hart","authors":["E. 't Hart"],"originalAuth":{"authors":["E. 't Hart"]}}}]}},"words":[{"verbatim":"Elaeagnus","normalized":"Elaeagnus","wordType":"GENUS","start":0,"end":9},{"verbatim":"triflora","normalized":"triflora","wordType":"SPECIES","start":10,"end":18},{"verbatim":"Roxb.","normalized":"Roxb.","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":25,"end":29},{"verbatim":"brevilimbatus","normalized":"brevilimbatus","wordType":"INFRASPECIES","start":30,"end":43},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"'t","normalized":"'t","wordType":"AUTHOR_WORD","start":46,"end":48},{"verbatim":"Hart","normalized":"Hart","wordType":"AUTHOR_WORD","start":49,"end":53}],"id":"e3b3f47c-856a-5c21-bfa7-ac8c89453232","parserVersion":"test_version"}
```

Name: Laevistrombus guidoi (Man in't Veld & De Turck, 1998)

Canonical: Laevistrombus guidoi

Authorship: (Man in't Veld & De Turck 1998)

```json
{"parsed":true,"quality":1,"verbatim":"Laevistrombus guidoi (Man in't Veld \u0026 De Turck, 1998)","normalized":"Laevistrombus guidoi (Man in't Veld \u0026 De Turck 1998)","canonical":{"stemmed":"Laevistrombus guido","simple":"Laevistrombus guidoi","full":"Laevistrombus guidoi"},"cardinality":2,"authorship":{"verbatim":"(Man in't Veld \u0026 De Turck, 1998)","normalized":"(Man in't Veld \u0026 De Turck 1998)","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Laevistrombus","species":"guidoi","authorship":{"verbatim":"(Man in't Veld \u0026 De Turck, 1998)","normalized":"(Man in't Veld \u0026 De Turck 1998)","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Laevistrombus","normalized":"Laevistrombus","wordType":"GENUS","start":0,"end":13},{"verbatim":"guidoi","normalized":"guidoi","wordType":"SPECIES","start":14,"end":20},{"verbatim":"Man","normalized":"Man","wordType":"AUTHOR_WORD","start":22,"end":25},{"verbatim":"in't","normalized":"in't","wordType":"AUTHOR_WORD","start":26,"end":30},{"verbatim":"Veld","normalized":"Veld","wordType":"AUTHOR_WORD","start":31,"end":35},{"verbatim":"De","normalized":"De","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"Turck","normalized":"Turck","wordType":"AUTHOR_WORD","start":41,"end":46},{"verbatim":"1998","normalized":"1998","wordType":"YEAR","start":48,"end":52}],"id":"e3ff94a0-92d0-5894-8599-f288e92077c8","parserVersion":"test_version"}
```

Name: Strombus guidoi Man in't Veld & De Turck, 1998

Canonical: Strombus guidoi

Authorship: Man in't Veld & De Turck 1998

```json
{"parsed":true,"quality":1,"verbatim":"Strombus guidoi Man in't Veld \u0026 De Turck, 1998","normalized":"Strombus guidoi Man in't Veld \u0026 De Turck 1998","canonical":{"stemmed":"Strombus guido","simple":"Strombus guidoi","full":"Strombus guidoi"},"cardinality":2,"authorship":{"verbatim":"Man in't Veld \u0026 De Turck, 1998","normalized":"Man in't Veld \u0026 De Turck 1998","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Strombus","species":"guidoi","authorship":{"verbatim":"Man in't Veld \u0026 De Turck, 1998","normalized":"Man in't Veld \u0026 De Turck 1998","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Strombus","normalized":"Strombus","wordType":"GENUS","start":0,"end":8},{"verbatim":"guidoi","normalized":"guidoi","wordType":"SPECIES","start":9,"end":15},{"verbatim":"Man","normalized":"Man","wordType":"AUTHOR_WORD","start":16,"end":19},{"verbatim":"in't","normalized":"in't","wordType":"AUTHOR_WORD","start":20,"end":24},{"verbatim":"Veld","normalized":"Veld","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"De","normalized":"De","wordType":"AUTHOR_WORD","start":32,"end":34},{"verbatim":"Turck","normalized":"Turck","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"1998","normalized":"1998","wordType":"YEAR","start":42,"end":46}],"id":"100d3b6e-62d3-51ad-baf6-60408babc574","parserVersion":"test_version"}
```

Name: Strombus vittatus entropi Man in't Veld & Visser, 1993

Canonical: Strombus vittatus entropi

Authorship: Man in't Veld & Visser 1993

```json
{"parsed":true,"quality":1,"verbatim":"Strombus vittatus entropi Man in't Veld \u0026 Visser, 1993","normalized":"Strombus vittatus entropi Man in't Veld \u0026 Visser 1993","canonical":{"stemmed":"Strombus uittat entrop","simple":"Strombus vittatus entropi","full":"Strombus vittatus entropi"},"cardinality":3,"authorship":{"verbatim":"Man in't Veld \u0026 Visser, 1993","normalized":"Man in't Veld \u0026 Visser 1993","year":"1993","authors":["Man in't Veld","Visser"],"originalAuth":{"authors":["Man in't Veld","Visser"],"year":{"year":"1993"}}},"details":{"infraspecies":{"genus":"Strombus","species":"vittatus","infraspecies":[{"value":"entropi","authorship":{"verbatim":"Man in't Veld \u0026 Visser, 1993","normalized":"Man in't Veld \u0026 Visser 1993","year":"1993","authors":["Man in't Veld","Visser"],"originalAuth":{"authors":["Man in't Veld","Visser"],"year":{"year":"1993"}}}}]}},"words":[{"verbatim":"Strombus","normalized":"Strombus","wordType":"GENUS","start":0,"end":8},{"verbatim":"vittatus","normalized":"vittatus","wordType":"SPECIES","start":9,"end":17},{"verbatim":"entropi","normalized":"entropi","wordType":"INFRASPECIES","start":18,"end":25},{"verbatim":"Man","normalized":"Man","wordType":"AUTHOR_WORD","start":26,"end":29},{"verbatim":"in't","normalized":"in't","wordType":"AUTHOR_WORD","start":30,"end":34},{"verbatim":"Veld","normalized":"Veld","wordType":"AUTHOR_WORD","start":35,"end":39},{"verbatim":"Visser","normalized":"Visser","wordType":"AUTHOR_WORD","start":42,"end":48},{"verbatim":"1993","normalized":"1993","wordType":"YEAR","start":50,"end":54}],"id":"c74691e3-0f71-576b-81ea-6173bdae9817","parserVersion":"test_version"}
```

Name: Velutina haliotoides (Linnaeus, 1758),

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758),","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":",","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Velutina","normalized":"Velutina","wordType":"GENUS","start":0,"end":8},{"verbatim":"haliotoides","normalized":"haliotoides","wordType":"SPECIES","start":9,"end":20},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":32,"end":36}],"id":"59093ba7-64a1-53c4-9795-12de7ff9e718","parserVersion":"test_version"}
```

Name: Hennediella microphylla (R.Br.bis) Paris

Canonical: Hennediella microphylla

Authorship: (R. Br. bis) Paris

```json
{"parsed":true,"quality":1,"verbatim":"Hennediella microphylla (R.Br.bis) Paris","normalized":"Hennediella microphylla (R. Br. bis) Paris","canonical":{"stemmed":"Hennediella microphyll","simple":"Hennediella microphylla","full":"Hennediella microphylla"},"cardinality":2,"authorship":{"verbatim":"(R.Br.bis) Paris","normalized":"(R. Br. bis) Paris","authors":["R. Br. bis","Paris"],"originalAuth":{"authors":["R. Br. bis"]},"combinationAuth":{"authors":["Paris"]}},"details":{"species":{"genus":"Hennediella","species":"microphylla","authorship":{"verbatim":"(R.Br.bis) Paris","normalized":"(R. Br. bis) Paris","authors":["R. Br. bis","Paris"],"originalAuth":{"authors":["R. Br. bis"]},"combinationAuth":{"authors":["Paris"]}}}},"words":[{"verbatim":"Hennediella","normalized":"Hennediella","wordType":"GENUS","start":0,"end":11},{"verbatim":"microphylla","normalized":"microphylla","wordType":"SPECIES","start":12,"end":23},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Br.","normalized":"Br.","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"bis","normalized":"bis","wordType":"AUTHOR_WORD","start":30,"end":33},{"verbatim":"Paris","normalized":"Paris","wordType":"AUTHOR_WORD","start":35,"end":40}],"id":"e8cc6d9d-6e6c-53a1-99a9-59f636009ed0","parserVersion":"test_version"}
```

Name: Pseudocercosporella endophytica Crous & H. Sm. ter

Canonical: Pseudocercosporella endophytica

Authorship: Crous & H. Sm. ter

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercosporella endophytica Crous \u0026 H. Sm. ter","normalized":"Pseudocercosporella endophytica Crous \u0026 H. Sm. ter","canonical":{"stemmed":"Pseudocercosporella endophytic","simple":"Pseudocercosporella endophytica","full":"Pseudocercosporella endophytica"},"cardinality":2,"authorship":{"verbatim":"Crous \u0026 H. Sm. ter","normalized":"Crous \u0026 H. Sm. ter","authors":["Crous","H. Sm. ter"],"originalAuth":{"authors":["Crous","H. Sm. ter"]}},"details":{"species":{"genus":"Pseudocercosporella","species":"endophytica","authorship":{"verbatim":"Crous \u0026 H. Sm. ter","normalized":"Crous \u0026 H. Sm. ter","authors":["Crous","H. Sm. ter"],"originalAuth":{"authors":["Crous","H. Sm. ter"]}}}},"words":[{"verbatim":"Pseudocercosporella","normalized":"Pseudocercosporella","wordType":"GENUS","start":0,"end":19},{"verbatim":"endophytica","normalized":"endophytica","wordType":"SPECIES","start":20,"end":31},{"verbatim":"Crous","normalized":"Crous","wordType":"AUTHOR_WORD","start":32,"end":37},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":40,"end":42},{"verbatim":"Sm.","normalized":"Sm.","wordType":"AUTHOR_WORD","start":43,"end":46},{"verbatim":"ter","normalized":"ter","wordType":"AUTHOR_WORD","start":47,"end":50}],"id":"ac52e64e-1cbe-57c8-86e2-6f5887a84da7","parserVersion":"test_version"}
```

Name: Kudoa amazonica Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves & Matos, 2019

Canonical: Kudoa amazonica

Authorship: Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves & Matos 2019

```json
{"parsed":true,"quality":1,"verbatim":"Kudoa amazonica Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos, 2019","normalized":"Kudoa amazonica Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos 2019","canonical":{"stemmed":"Kudoa amazonic","simple":"Kudoa amazonica","full":"Kudoa amazonica"},"cardinality":2,"authorship":{"verbatim":"Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos, 2019","normalized":"Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos 2019","year":"2019","authors":["Velasco","Sindeaux Neto","Videira","de Cássia Silva do Nascimento","Gonçalves","Matos"],"originalAuth":{"authors":["Velasco","Sindeaux Neto","Videira","de Cássia Silva do Nascimento","Gonçalves","Matos"],"year":{"year":"2019"}}},"details":{"species":{"genus":"Kudoa","species":"amazonica","authorship":{"verbatim":"Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos, 2019","normalized":"Velasco, Sindeaux Neto, Videira, de Cássia Silva do Nascimento, Gonçalves \u0026 Matos 2019","year":"2019","authors":["Velasco","Sindeaux Neto","Videira","de Cássia Silva do Nascimento","Gonçalves","Matos"],"originalAuth":{"authors":["Velasco","Sindeaux Neto","Videira","de Cássia Silva do Nascimento","Gonçalves","Matos"],"year":{"year":"2019"}}}}},"words":[{"verbatim":"Kudoa","normalized":"Kudoa","wordType":"GENUS","start":0,"end":5},{"verbatim":"amazonica","normalized":"amazonica","wordType":"SPECIES","start":6,"end":15},{"verbatim":"Velasco","normalized":"Velasco","wordType":"AUTHOR_WORD","start":16,"end":23},{"verbatim":"Sindeaux","normalized":"Sindeaux","wordType":"AUTHOR_WORD","start":25,"end":33},{"verbatim":"Neto","normalized":"Neto","wordType":"AUTHOR_WORD","start":34,"end":38},{"verbatim":"Videira","normalized":"Videira","wordType":"AUTHOR_WORD","start":40,"end":47},{"verbatim":"de","normalized":"de","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"Cássia","normalized":"Cássia","wordType":"AUTHOR_WORD","start":52,"end":58},{"verbatim":"Silva","normalized":"Silva","wordType":"AUTHOR_WORD","start":59,"end":64},{"verbatim":"do","normalized":"do","wordType":"AUTHOR_WORD","start":65,"end":67},{"verbatim":"Nascimento","normalized":"Nascimento","wordType":"AUTHOR_WORD","start":68,"end":78},{"verbatim":"Gonçalves","normalized":"Gonçalves","wordType":"AUTHOR_WORD","start":80,"end":89},{"verbatim":"Matos","normalized":"Matos","wordType":"AUTHOR_WORD","start":92,"end":97},{"verbatim":"2019","normalized":"2019","wordType":"YEAR","start":99,"end":103}],"id":"331fe77e-4a0e-555a-90ef-2874b72e5c7f","parserVersion":"test_version"}
```

Name: Branchinecta papillata Rogers, de los Rios & Zuniga, 2008

Canonical: Branchinecta papillata

Authorship: Rogers, de los Rios & Zuniga 2008

```json
{"parsed":true,"quality":1,"verbatim":"Branchinecta papillata Rogers, de los Rios \u0026 Zuniga, 2008","normalized":"Branchinecta papillata Rogers, de los Rios \u0026 Zuniga 2008","canonical":{"stemmed":"Branchinecta papillat","simple":"Branchinecta papillata","full":"Branchinecta papillata"},"cardinality":2,"authorship":{"verbatim":"Rogers, de los Rios \u0026 Zuniga, 2008","normalized":"Rogers, de los Rios \u0026 Zuniga 2008","year":"2008","authors":["Rogers","de los Rios","Zuniga"],"originalAuth":{"authors":["Rogers","de los Rios","Zuniga"],"year":{"year":"2008"}}},"details":{"species":{"genus":"Branchinecta","species":"papillata","authorship":{"verbatim":"Rogers, de los Rios \u0026 Zuniga, 2008","normalized":"Rogers, de los Rios \u0026 Zuniga 2008","year":"2008","authors":["Rogers","de los Rios","Zuniga"],"originalAuth":{"authors":["Rogers","de los Rios","Zuniga"],"year":{"year":"2008"}}}}},"words":[{"verbatim":"Branchinecta","normalized":"Branchinecta","wordType":"GENUS","start":0,"end":12},{"verbatim":"papillata","normalized":"papillata","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Rogers","normalized":"Rogers","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"de los","normalized":"de los","wordType":"AUTHOR_WORD","start":31,"end":37},{"verbatim":"Rios","normalized":"Rios","wordType":"AUTHOR_WORD","start":38,"end":42},{"verbatim":"Zuniga","normalized":"Zuniga","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"2008","normalized":"2008","wordType":"YEAR","start":53,"end":57}],"id":"220f3428-87b9-5455-9b71-4998c9ccfd00","parserVersion":"test_version"}
```

Name: Echiophis brunneus (Castro-Aguirre & Suárez de los Cobos, 1983)

Canonical: Echiophis brunneus

Authorship: (Castro-Aguirre & Suárez de los Cobos 1983)

```json
{"parsed":true,"quality":1,"verbatim":"Echiophis brunneus (Castro-Aguirre \u0026 Suárez de los Cobos, 1983)","normalized":"Echiophis brunneus (Castro-Aguirre \u0026 Suárez de los Cobos 1983)","canonical":{"stemmed":"Echiophis brunne","simple":"Echiophis brunneus","full":"Echiophis brunneus"},"cardinality":2,"authorship":{"verbatim":"(Castro-Aguirre \u0026 Suárez de los Cobos, 1983)","normalized":"(Castro-Aguirre \u0026 Suárez de los Cobos 1983)","year":"1983","authors":["Castro-Aguirre","Suárez de los Cobos"],"originalAuth":{"authors":["Castro-Aguirre","Suárez de los Cobos"],"year":{"year":"1983"}}},"details":{"species":{"genus":"Echiophis","species":"brunneus","authorship":{"verbatim":"(Castro-Aguirre \u0026 Suárez de los Cobos, 1983)","normalized":"(Castro-Aguirre \u0026 Suárez de los Cobos 1983)","year":"1983","authors":["Castro-Aguirre","Suárez de los Cobos"],"originalAuth":{"authors":["Castro-Aguirre","Suárez de los Cobos"],"year":{"year":"1983"}}}}},"words":[{"verbatim":"Echiophis","normalized":"Echiophis","wordType":"GENUS","start":0,"end":9},{"verbatim":"brunneus","normalized":"brunneus","wordType":"SPECIES","start":10,"end":18},{"verbatim":"Castro-Aguirre","normalized":"Castro-Aguirre","wordType":"AUTHOR_WORD","start":20,"end":34},{"verbatim":"Suárez","normalized":"Suárez","wordType":"AUTHOR_WORD","start":37,"end":43},{"verbatim":"de los","normalized":"de los","wordType":"AUTHOR_WORD","start":44,"end":50},{"verbatim":"Cobos","normalized":"Cobos","wordType":"AUTHOR_WORD","start":51,"end":56},{"verbatim":"1983","normalized":"1983","wordType":"YEAR","start":58,"end":62}],"id":"18d6069c-8c76-5a7a-8400-511777462b09","parserVersion":"test_version"}
```

### Binomials with an abbreviated genus

Name: M. alpium

Canonical: M. alpium

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"}],"verbatim":"M. alpium","normalized":"M. alpium","canonical":{"stemmed":"M. alpi","simple":"M. alpium","full":"M. alpium"},"cardinality":2,"details":{"species":{"genus":"M.","species":"alpium"}},"words":[{"verbatim":"M.","normalized":"M.","wordType":"GENUS","start":0,"end":2},{"verbatim":"alpium","normalized":"alpium","wordType":"SPECIES","start":3,"end":9}],"id":"9001ffb5-eac2-5bb4-8f78-d7b7e3e02bd8","parserVersion":"test_version"}
```

Name: Mo. alpium (Osbeck, 1778)

Canonical: Mo. alpium

Authorship: (Osbeck 1778)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"}],"verbatim":"Mo. alpium (Osbeck, 1778)","normalized":"Mo. alpium (Osbeck 1778)","canonical":{"stemmed":"Mo. alpi","simple":"Mo. alpium","full":"Mo. alpium"},"cardinality":2,"authorship":{"verbatim":"(Osbeck, 1778)","normalized":"(Osbeck 1778)","year":"1778","authors":["Osbeck"],"originalAuth":{"authors":["Osbeck"],"year":{"year":"1778"}}},"details":{"species":{"genus":"Mo.","species":"alpium","authorship":{"verbatim":"(Osbeck, 1778)","normalized":"(Osbeck 1778)","year":"1778","authors":["Osbeck"],"originalAuth":{"authors":["Osbeck"],"year":{"year":"1778"}}}}},"words":[{"verbatim":"Mo.","normalized":"Mo.","wordType":"GENUS","start":0,"end":3},{"verbatim":"alpium","normalized":"alpium","wordType":"SPECIES","start":4,"end":10},{"verbatim":"Osbeck","normalized":"Osbeck","wordType":"AUTHOR_WORD","start":12,"end":18},{"verbatim":"1778","normalized":"1778","wordType":"YEAR","start":20,"end":24}],"id":"1e9437b7-bf45-5b12-8da0-8966c6ea1c5c","parserVersion":"test_version"}
```

### Binomials with abbreviated subgenus

Name: Phalaena (Tin.) guttella Fab.

Canonical: Phalaena guttella

Authorship: Fab.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Abbreviated subgenus"}],"verbatim":"Phalaena (Tin.) guttella Fab.","normalized":"Phalaena (Tin.) guttella Fab.","canonical":{"stemmed":"Phalaena guttell","simple":"Phalaena guttella","full":"Phalaena guttella"},"cardinality":2,"authorship":{"verbatim":"Fab.","normalized":"Fab.","authors":["Fab."],"originalAuth":{"authors":["Fab."]}},"details":{"species":{"genus":"Phalaena","subgenus":"Tin.","species":"guttella","authorship":{"verbatim":"Fab.","normalized":"Fab.","authors":["Fab."],"originalAuth":{"authors":["Fab."]}}}},"words":[{"verbatim":"Phalaena","normalized":"Phalaena","wordType":"GENUS","start":0,"end":8},{"verbatim":"Tin.","normalized":"Tin.","wordType":"INFRA_GENUS","start":10,"end":14},{"verbatim":"guttella","normalized":"guttella","wordType":"SPECIES","start":16,"end":24},{"verbatim":"Fab.","normalized":"Fab.","wordType":"AUTHOR_WORD","start":25,"end":29}],"id":"da5f9d5b-abdf-5451-8dec-53830e05e43c","parserVersion":"test_version"}
```

Name: Gahrliepia (G.) tessellata Traub & Morrow 1955

Canonical: Gahrliepia tessellata

Authorship: Traub & Morrow 1955

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Abbreviated subgenus"}],"verbatim":"Gahrliepia (G.) tessellata Traub \u0026 Morrow 1955","normalized":"Gahrliepia (G.) tessellata Traub \u0026 Morrow 1955","canonical":{"stemmed":"Gahrliepia tessellat","simple":"Gahrliepia tessellata","full":"Gahrliepia tessellata"},"cardinality":2,"authorship":{"verbatim":"Traub \u0026 Morrow 1955","normalized":"Traub \u0026 Morrow 1955","year":"1955","authors":["Traub","Morrow"],"originalAuth":{"authors":["Traub","Morrow"],"year":{"year":"1955"}}},"details":{"species":{"genus":"Gahrliepia","subgenus":"G.","species":"tessellata","authorship":{"verbatim":"Traub \u0026 Morrow 1955","normalized":"Traub \u0026 Morrow 1955","year":"1955","authors":["Traub","Morrow"],"originalAuth":{"authors":["Traub","Morrow"],"year":{"year":"1955"}}}}},"words":[{"verbatim":"Gahrliepia","normalized":"Gahrliepia","wordType":"GENUS","start":0,"end":10},{"verbatim":"G.","normalized":"G.","wordType":"INFRA_GENUS","start":12,"end":14},{"verbatim":"tessellata","normalized":"tessellata","wordType":"SPECIES","start":16,"end":26},{"verbatim":"Traub","normalized":"Traub","wordType":"AUTHOR_WORD","start":27,"end":32},{"verbatim":"Morrow","normalized":"Morrow","wordType":"AUTHOR_WORD","start":35,"end":41},{"verbatim":"1955","normalized":"1955","wordType":"YEAR","start":42,"end":46}],"id":"776bb155-0d31-5a3d-9e87-e10ebf61a746","parserVersion":"test_version"}
```

Name: Bosmina (Eubosmina) coregoni x B. (E.) longispina

Canonical: Bosmina coregoni × Bosmina longispina

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Abbreviated subgenus"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Bosmina (Eubosmina) coregoni x B. (E.) longispina","normalized":"Bosmina (Eubosmina) coregoni × Bosmina (E.) longispina","canonical":{"stemmed":"Bosmina coregon × Bosmina longispin","simple":"Bosmina coregoni × Bosmina longispina","full":"Bosmina coregoni × Bosmina longispina"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Bosmina","subgenus":"Eubosmina","species":"coregoni"}},{"species":{"genus":"Bosmina","subgenus":"E.","species":"longispina"}}]},"words":[{"verbatim":"Bosmina","normalized":"Bosmina","wordType":"GENUS","start":0,"end":7},{"verbatim":"Eubosmina","normalized":"Eubosmina","wordType":"INFRA_GENUS","start":9,"end":18},{"verbatim":"coregoni","normalized":"coregoni","wordType":"SPECIES","start":20,"end":28},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":29,"end":30},{"verbatim":"B.","normalized":"Bosmina","wordType":"GENUS","start":31,"end":33},{"verbatim":"E.","normalized":"E.","wordType":"INFRA_GENUS","start":35,"end":37},{"verbatim":"longispina","normalized":"longispina","wordType":"SPECIES","start":39,"end":49}],"id":"71c160bf-428b-5b51-9d97-0965686033bc","parserVersion":"test_version"}
```

Name: Simia (Cercop.) nasuus Kerr 1792

Canonical: Simia nasuus

Authorship: Kerr 1792

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Abbreviated subgenus"}],"verbatim":"Simia (Cercop.) nasuus Kerr 1792","normalized":"Simia (Cercop.) nasuus Kerr 1792","canonical":{"stemmed":"Simia nasu","simple":"Simia nasuus","full":"Simia nasuus"},"cardinality":2,"authorship":{"verbatim":"Kerr 1792","normalized":"Kerr 1792","year":"1792","authors":["Kerr"],"originalAuth":{"authors":["Kerr"],"year":{"year":"1792"}}},"details":{"species":{"genus":"Simia","subgenus":"Cercop.","species":"nasuus","authorship":{"verbatim":"Kerr 1792","normalized":"Kerr 1792","year":"1792","authors":["Kerr"],"originalAuth":{"authors":["Kerr"],"year":{"year":"1792"}}}}},"words":[{"verbatim":"Simia","normalized":"Simia","wordType":"GENUS","start":0,"end":5},{"verbatim":"Cercop.","normalized":"Cercop.","wordType":"INFRA_GENUS","start":7,"end":14},{"verbatim":"nasuus","normalized":"nasuus","wordType":"SPECIES","start":16,"end":22},{"verbatim":"Kerr","normalized":"Kerr","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"1792","normalized":"1792","wordType":"YEAR","start":28,"end":32}],"id":"2f54aece-f7e0-5ed2-8744-f135ceab1c7f","parserVersion":"test_version"}
```


### Binomials with basionym and combination authors

Name: Yarrowia lipolytica var. lipolytica (Wick., Kurtzman & E.A. Herrm.) Van der Walt & Arx 1981

Canonical: Yarrowia lipolytica var. lipolytica

Authorship: (Wick., Kurtzman & E. A. Herrm.) Van der Walt & Arx 1981

```json
{"parsed":true,"quality":1,"verbatim":"Yarrowia lipolytica var. lipolytica (Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"Yarrowia lipolytica var. lipolytica (Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","canonical":{"stemmed":"Yarrowia lipolytic lipolytic","simple":"Yarrowia lipolytica lipolytica","full":"Yarrowia lipolytica var. lipolytica"},"cardinality":3,"authorship":{"verbatim":"(Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"(Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","authors":["Wick.","Kurtzman","E. A. Herrm.","Van der Walt","Arx"],"originalAuth":{"authors":["Wick.","Kurtzman","E. A. Herrm."]},"combinationAuth":{"authors":["Van der Walt","Arx"],"year":{"year":"1981"}}},"details":{"infraspecies":{"genus":"Yarrowia","species":"lipolytica","infraspecies":[{"value":"lipolytica","rank":"var.","authorship":{"verbatim":"(Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"(Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","authors":["Wick.","Kurtzman","E. A. Herrm.","Van der Walt","Arx"],"originalAuth":{"authors":["Wick.","Kurtzman","E. A. Herrm."]},"combinationAuth":{"authors":["Van der Walt","Arx"],"year":{"year":"1981"}}}}]}},"words":[{"verbatim":"Yarrowia","normalized":"Yarrowia","wordType":"GENUS","start":0,"end":8},{"verbatim":"lipolytica","normalized":"lipolytica","wordType":"SPECIES","start":9,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"lipolytica","normalized":"lipolytica","wordType":"INFRASPECIES","start":25,"end":35},{"verbatim":"Wick.","normalized":"Wick.","wordType":"AUTHOR_WORD","start":37,"end":42},{"verbatim":"Kurtzman","normalized":"Kurtzman","wordType":"AUTHOR_WORD","start":44,"end":52},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":55,"end":57},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":57,"end":59},{"verbatim":"Herrm.","normalized":"Herrm.","wordType":"AUTHOR_WORD","start":60,"end":66},{"verbatim":"Van","normalized":"Van","wordType":"AUTHOR_WORD","start":68,"end":71},{"verbatim":"der","normalized":"der","wordType":"AUTHOR_WORD","start":72,"end":75},{"verbatim":"Walt","normalized":"Walt","wordType":"AUTHOR_WORD","start":76,"end":80},{"verbatim":"Arx","normalized":"Arx","wordType":"AUTHOR_WORD","start":83,"end":86},{"verbatim":"1981","normalized":"1981","wordType":"YEAR","start":87,"end":91}],"id":"e649d828-0ae9-5b5b-b079-1485c9bbf872","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett)U. Braun & Crous     2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett) U. Braun & Crous 2003

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett) U. Braun \u0026 Crous 2003","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"]},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett) U. Braun \u0026 Crous 2003","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"]},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Burnett","normalized":"Burnett","wordType":"AUTHOR_WORD","start":36,"end":43},{"verbatim":"U.","normalized":"U.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":47,"end":52},{"verbatim":"Crous","normalized":"Crous","wordType":"AUTHOR_WORD","start":55,"end":60},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":65,"end":69}],"id":"3c52bc21-3ac9-5be4-9d5f-1f84fe9d3325","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett, 1873)U. Braun & Crous     2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett 1873) U. Braun & Crous 2003

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett 1873) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Burnett","normalized":"Burnett","wordType":"AUTHOR_WORD","start":36,"end":43},{"verbatim":"1873","normalized":"1873","wordType":"YEAR","start":45,"end":49},{"verbatim":"U.","normalized":"U.","wordType":"AUTHOR_WORD","start":50,"end":52},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":53,"end":58},{"verbatim":"Crous","normalized":"Crous","wordType":"AUTHOR_WORD","start":61,"end":66},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":71,"end":75}],"id":"8e5dd168-d7f1-51e4-989c-cedb253d572c","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett 1873)U. Braun & Crous ,    2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett 1873) U. Braun & Crous 2003

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett 1873) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Burnett","normalized":"Burnett","wordType":"AUTHOR_WORD","start":36,"end":43},{"verbatim":"1873","normalized":"1873","wordType":"YEAR","start":44,"end":48},{"verbatim":"U.","normalized":"U.","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":52,"end":57},{"verbatim":"Crous","normalized":"Crous","wordType":"AUTHOR_WORD","start":60,"end":65},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":71,"end":75}],"id":"a35b47c6-6716-5750-ab81-a19aed44143b","parserVersion":"test_version"}
```

Name: Sedella pumila (Benth.) Britton & Rose

Canonical: Sedella pumila

Authorship: (Benth.) Britton & Rose

```json
{"parsed":true,"quality":1,"verbatim":"Sedella pumila (Benth.) Britton \u0026 Rose","normalized":"Sedella pumila (Benth.) Britton \u0026 Rose","canonical":{"stemmed":"Sedella pumil","simple":"Sedella pumila","full":"Sedella pumila"},"cardinality":2,"authorship":{"verbatim":"(Benth.) Britton \u0026 Rose","normalized":"(Benth.) Britton \u0026 Rose","authors":["Benth.","Britton","Rose"],"originalAuth":{"authors":["Benth."]},"combinationAuth":{"authors":["Britton","Rose"]}},"details":{"species":{"genus":"Sedella","species":"pumila","authorship":{"verbatim":"(Benth.) Britton \u0026 Rose","normalized":"(Benth.) Britton \u0026 Rose","authors":["Benth.","Britton","Rose"],"originalAuth":{"authors":["Benth."]},"combinationAuth":{"authors":["Britton","Rose"]}}}},"words":[{"verbatim":"Sedella","normalized":"Sedella","wordType":"GENUS","start":0,"end":7},{"verbatim":"pumila","normalized":"pumila","wordType":"SPECIES","start":8,"end":14},{"verbatim":"Benth.","normalized":"Benth.","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"Britton","normalized":"Britton","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"Rose","normalized":"Rose","wordType":"AUTHOR_WORD","start":34,"end":38}],"id":"393cedba-6ff1-5e5c-83f0-21e32f031ab7","parserVersion":"test_version"}
```

Name: Impatiens nomenyae Eb.Fisch. & Raheliv.

Canonical: Impatiens nomenyae

Authorship: Eb. Fisch. & Raheliv.

```json
{"parsed":true,"quality":1,"verbatim":"Impatiens nomenyae Eb.Fisch. \u0026 Raheliv.","normalized":"Impatiens nomenyae Eb. Fisch. \u0026 Raheliv.","canonical":{"stemmed":"Impatiens nomeny","simple":"Impatiens nomenyae","full":"Impatiens nomenyae"},"cardinality":2,"authorship":{"verbatim":"Eb.Fisch. \u0026 Raheliv.","normalized":"Eb. Fisch. \u0026 Raheliv.","authors":["Eb. Fisch.","Raheliv."],"originalAuth":{"authors":["Eb. Fisch.","Raheliv."]}},"details":{"species":{"genus":"Impatiens","species":"nomenyae","authorship":{"verbatim":"Eb.Fisch. \u0026 Raheliv.","normalized":"Eb. Fisch. \u0026 Raheliv.","authors":["Eb. Fisch.","Raheliv."],"originalAuth":{"authors":["Eb. Fisch.","Raheliv."]}}}},"words":[{"verbatim":"Impatiens","normalized":"Impatiens","wordType":"GENUS","start":0,"end":9},{"verbatim":"nomenyae","normalized":"nomenyae","wordType":"SPECIES","start":10,"end":18},{"verbatim":"Eb.","normalized":"Eb.","wordType":"AUTHOR_WORD","start":19,"end":22},{"verbatim":"Fisch.","normalized":"Fisch.","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"Raheliv.","normalized":"Raheliv.","wordType":"AUTHOR_WORD","start":31,"end":39}],"id":"6452d4ac-738b-5773-8d69-50232e2842a1","parserVersion":"test_version"}
```

Name: Armeria carpetana ssp. carpetana H. del Villar

Canonical: Armeria carpetana subsp. carpetana

Authorship: H. del Villar

```json
{"parsed":true,"quality":1,"verbatim":"Armeria carpetana ssp. carpetana H. del Villar","normalized":"Armeria carpetana subsp. carpetana H. del Villar","canonical":{"stemmed":"Armeria carpetan carpetan","simple":"Armeria carpetana carpetana","full":"Armeria carpetana subsp. carpetana"},"cardinality":3,"authorship":{"verbatim":"H. del Villar","normalized":"H. del Villar","authors":["H. del Villar"],"originalAuth":{"authors":["H. del Villar"]}},"details":{"infraspecies":{"genus":"Armeria","species":"carpetana","infraspecies":[{"value":"carpetana","rank":"subsp.","authorship":{"verbatim":"H. del Villar","normalized":"H. del Villar","authors":["H. del Villar"],"originalAuth":{"authors":["H. del Villar"]}}}]}},"words":[{"verbatim":"Armeria","normalized":"Armeria","wordType":"GENUS","start":0,"end":7},{"verbatim":"carpetana","normalized":"carpetana","wordType":"SPECIES","start":8,"end":17},{"verbatim":"ssp.","normalized":"subsp.","wordType":"RANK","start":18,"end":22},{"verbatim":"carpetana","normalized":"carpetana","wordType":"INFRASPECIES","start":23,"end":32},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"del","normalized":"del","wordType":"AUTHOR_WORD","start":36,"end":39},{"verbatim":"Villar","normalized":"Villar","wordType":"AUTHOR_WORD","start":40,"end":46}],"id":"4b16116e-549d-56bf-959a-ff11edb25021","parserVersion":"test_version"}
```

### Exceptions with Binomials

Name: Navicula bacterium Frenguelli

Canonical: Navicula bacterium

Authorship: Frenguelli

```json
{"parsed":true,"quality":1,"verbatim":"Navicula bacterium Frenguelli","normalized":"Navicula bacterium Frenguelli","canonical":{"stemmed":"Navicula bacteri","simple":"Navicula bacterium","full":"Navicula bacterium"},"cardinality":2,"authorship":{"verbatim":"Frenguelli","normalized":"Frenguelli","authors":["Frenguelli"],"originalAuth":{"authors":["Frenguelli"]}},"details":{"species":{"genus":"Navicula","species":"bacterium","authorship":{"verbatim":"Frenguelli","normalized":"Frenguelli","authors":["Frenguelli"],"originalAuth":{"authors":["Frenguelli"]}}}},"words":[{"verbatim":"Navicula","normalized":"Navicula","wordType":"GENUS","start":0,"end":8},{"verbatim":"bacterium","normalized":"bacterium","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Frenguelli","normalized":"Frenguelli","wordType":"AUTHOR_WORD","start":19,"end":29}],"id":"0c0ce62a-8ea4-569c-b918-46e7f8c942ef","parserVersion":"test_version"}
```

Name: Bottaria nudum (Nyl.) Vain.

Canonical: Bottaria nudum

Authorship: (Nyl.) Vain.

```json
{"parsed":true,"quality":1,"verbatim":"Bottaria nudum (Nyl.) Vain.","normalized":"Bottaria nudum (Nyl.) Vain.","canonical":{"stemmed":"Bottaria nud","simple":"Bottaria nudum","full":"Bottaria nudum"},"cardinality":2,"authorship":{"verbatim":"(Nyl.) Vain.","normalized":"(Nyl.) Vain.","authors":["Nyl.","Vain."],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["Vain."]}},"details":{"species":{"genus":"Bottaria","species":"nudum","authorship":{"verbatim":"(Nyl.) Vain.","normalized":"(Nyl.) Vain.","authors":["Nyl.","Vain."],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["Vain."]}}}},"words":[{"verbatim":"Bottaria","normalized":"Bottaria","wordType":"GENUS","start":0,"end":8},{"verbatim":"nudum","normalized":"nudum","wordType":"SPECIES","start":9,"end":14},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":16,"end":20},{"verbatim":"Vain.","normalized":"Vain.","wordType":"AUTHOR_WORD","start":22,"end":27}],"id":"91799409-de6f-5341-ab24-336da9f6b80b","parserVersion":"test_version"}
```

Name: Turkozelotes attavirus Chatzaki, 2019

Canonical: Turkozelotes attavirus

Authorship: Chatzaki 2019

```json
{"parsed":true,"quality":1,"verbatim":"Turkozelotes attavirus Chatzaki, 2019","normalized":"Turkozelotes attavirus Chatzaki 2019","canonical":{"stemmed":"Turkozelotes attauir","simple":"Turkozelotes attavirus","full":"Turkozelotes attavirus"},"cardinality":2,"authorship":{"verbatim":"Chatzaki, 2019","normalized":"Chatzaki 2019","year":"2019","authors":["Chatzaki"],"originalAuth":{"authors":["Chatzaki"],"year":{"year":"2019"}}},"details":{"species":{"genus":"Turkozelotes","species":"attavirus","authorship":{"verbatim":"Chatzaki, 2019","normalized":"Chatzaki 2019","year":"2019","authors":["Chatzaki"],"originalAuth":{"authors":["Chatzaki"],"year":{"year":"2019"}}}}},"words":[{"verbatim":"Turkozelotes","normalized":"Turkozelotes","wordType":"GENUS","start":0,"end":12},{"verbatim":"attavirus","normalized":"attavirus","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Chatzaki","normalized":"Chatzaki","wordType":"AUTHOR_WORD","start":23,"end":31},{"verbatim":"2019","normalized":"2019","wordType":"YEAR","start":33,"end":37}],"id":"60295698-060d-5ffd-982b-e3c0e0d6a1c7","parserVersion":"test_version"}
```

Name: Phalium (Semicassis) vector R. T. Abbott, 1993

Canonical: Phalium vector

Authorship: R. T. Abbott 1993

```json
{"parsed":true,"quality":1,"verbatim":"Phalium (Semicassis) vector R. T. Abbott, 1993","normalized":"Phalium (Semicassis) vector R. T. Abbott 1993","canonical":{"stemmed":"Phalium uector","simple":"Phalium vector","full":"Phalium vector"},"cardinality":2,"authorship":{"verbatim":"R. T. Abbott, 1993","normalized":"R. T. Abbott 1993","year":"1993","authors":["R. T. Abbott"],"originalAuth":{"authors":["R. T. Abbott"],"year":{"year":"1993"}}},"details":{"species":{"genus":"Phalium","subgenus":"Semicassis","species":"vector","authorship":{"verbatim":"R. T. Abbott, 1993","normalized":"R. T. Abbott 1993","year":"1993","authors":["R. T. Abbott"],"originalAuth":{"authors":["R. T. Abbott"],"year":{"year":"1993"}}}}},"words":[{"verbatim":"Phalium","normalized":"Phalium","wordType":"GENUS","start":0,"end":7},{"verbatim":"Semicassis","normalized":"Semicassis","wordType":"INFRA_GENUS","start":9,"end":19},{"verbatim":"vector","normalized":"vector","wordType":"SPECIES","start":21,"end":27},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":31,"end":33},{"verbatim":"Abbott","normalized":"Abbott","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"1993","normalized":"1993","wordType":"YEAR","start":42,"end":46}],"id":"15589e11-23ac-5896-859c-448018697211","parserVersion":"test_version"}
```

Name: Spirophora bacterium Lendenfeld, 1887

Canonical: Spirophora bacterium

Authorship: Lendenfeld 1887

```json
{"parsed":true,"quality":1,"verbatim":"Spirophora bacterium Lendenfeld, 1887","normalized":"Spirophora bacterium Lendenfeld 1887","canonical":{"stemmed":"Spirophora bacteri","simple":"Spirophora bacterium","full":"Spirophora bacterium"},"cardinality":2,"authorship":{"verbatim":"Lendenfeld, 1887","normalized":"Lendenfeld 1887","year":"1887","authors":["Lendenfeld"],"originalAuth":{"authors":["Lendenfeld"],"year":{"year":"1887"}}},"details":{"species":{"genus":"Spirophora","species":"bacterium","authorship":{"verbatim":"Lendenfeld, 1887","normalized":"Lendenfeld 1887","year":"1887","authors":["Lendenfeld"],"originalAuth":{"authors":["Lendenfeld"],"year":{"year":"1887"}}}}},"words":[{"verbatim":"Spirophora","normalized":"Spirophora","wordType":"GENUS","start":0,"end":10},{"verbatim":"bacterium","normalized":"bacterium","wordType":"SPECIES","start":11,"end":20},{"verbatim":"Lendenfeld","normalized":"Lendenfeld","wordType":"AUTHOR_WORD","start":21,"end":31},{"verbatim":"1887","normalized":"1887","wordType":"YEAR","start":33,"end":37}],"id":"df16a7e2-a81f-578e-9e1c-ce8644fe4a62","parserVersion":"test_version"}
```

### Binomials with Mc and Mac authors

Name: Zygocera norfolkensis McKeown 1938

Canonical: Zygocera norfolkensis

Authorship: McKeown 1938

```json
{"parsed":true,"quality":1,"verbatim":"Zygocera norfolkensis McKeown 1938","normalized":"Zygocera norfolkensis McKeown 1938","canonical":{"stemmed":"Zygocera norfolkens","simple":"Zygocera norfolkensis","full":"Zygocera norfolkensis"},"cardinality":2,"authorship":{"verbatim":"McKeown 1938","normalized":"McKeown 1938","year":"1938","authors":["McKeown"],"originalAuth":{"authors":["McKeown"],"year":{"year":"1938"}}},"details":{"species":{"genus":"Zygocera","species":"norfolkensis","authorship":{"verbatim":"McKeown 1938","normalized":"McKeown 1938","year":"1938","authors":["McKeown"],"originalAuth":{"authors":["McKeown"],"year":{"year":"1938"}}}}},"words":[{"verbatim":"Zygocera","normalized":"Zygocera","wordType":"GENUS","start":0,"end":8},{"verbatim":"norfolkensis","normalized":"norfolkensis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"McKeown","normalized":"McKeown","wordType":"AUTHOR_WORD","start":22,"end":29},{"verbatim":"1938","normalized":"1938","wordType":"YEAR","start":30,"end":34}],"id":"9286faf0-6410-51df-b647-f9f546f610b4","parserVersion":"test_version"}
```

Name: Zygocera norfolkensis MacKeown 1938

Canonical: Zygocera norfolkensis

Authorship: MacKeown 1938

```json
{"parsed":true,"quality":1,"verbatim":"Zygocera norfolkensis MacKeown 1938","normalized":"Zygocera norfolkensis MacKeown 1938","canonical":{"stemmed":"Zygocera norfolkens","simple":"Zygocera norfolkensis","full":"Zygocera norfolkensis"},"cardinality":2,"authorship":{"verbatim":"MacKeown 1938","normalized":"MacKeown 1938","year":"1938","authors":["MacKeown"],"originalAuth":{"authors":["MacKeown"],"year":{"year":"1938"}}},"details":{"species":{"genus":"Zygocera","species":"norfolkensis","authorship":{"verbatim":"MacKeown 1938","normalized":"MacKeown 1938","year":"1938","authors":["MacKeown"],"originalAuth":{"authors":["MacKeown"],"year":{"year":"1938"}}}}},"words":[{"verbatim":"Zygocera","normalized":"Zygocera","wordType":"GENUS","start":0,"end":8},{"verbatim":"norfolkensis","normalized":"norfolkensis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"MacKeown","normalized":"MacKeown","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1938","normalized":"1938","wordType":"YEAR","start":31,"end":35}],"id":"b1fc99c8-6b6c-5208-a897-910c4738286c","parserVersion":"test_version"}
```

Name: Zygocera norfolkensis Mac'Keown 1938

Canonical: Zygocera norfolkensis

Authorship: Mac'Keown 1938

```json
{"parsed":true,"quality":1,"verbatim":"Zygocera norfolkensis Mac'Keown 1938","normalized":"Zygocera norfolkensis Mac'Keown 1938","canonical":{"stemmed":"Zygocera norfolkens","simple":"Zygocera norfolkensis","full":"Zygocera norfolkensis"},"cardinality":2,"authorship":{"verbatim":"Mac'Keown 1938","normalized":"Mac'Keown 1938","year":"1938","authors":["Mac'Keown"],"originalAuth":{"authors":["Mac'Keown"],"year":{"year":"1938"}}},"details":{"species":{"genus":"Zygocera","species":"norfolkensis","authorship":{"verbatim":"Mac'Keown 1938","normalized":"Mac'Keown 1938","year":"1938","authors":["Mac'Keown"],"originalAuth":{"authors":["Mac'Keown"],"year":{"year":"1938"}}}}},"words":[{"verbatim":"Zygocera","normalized":"Zygocera","wordType":"GENUS","start":0,"end":8},{"verbatim":"norfolkensis","normalized":"norfolkensis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Mac'Keown","normalized":"Mac'Keown","wordType":"AUTHOR_WORD","start":22,"end":31},{"verbatim":"1938","normalized":"1938","wordType":"YEAR","start":32,"end":36}],"id":"7da46f00-251c-5e42-b314-756f0f2b4f41","parserVersion":"test_version"}
```

Name: Zygocera norfolkensis Mc'Keown 1938

Canonical: Zygocera norfolkensis

Authorship: Mc'Keown 1938

```json
{"parsed":true,"quality":1,"verbatim":"Zygocera norfolkensis Mc'Keown 1938","normalized":"Zygocera norfolkensis Mc'Keown 1938","canonical":{"stemmed":"Zygocera norfolkens","simple":"Zygocera norfolkensis","full":"Zygocera norfolkensis"},"cardinality":2,"authorship":{"verbatim":"Mc'Keown 1938","normalized":"Mc'Keown 1938","year":"1938","authors":["Mc'Keown"],"originalAuth":{"authors":["Mc'Keown"],"year":{"year":"1938"}}},"details":{"species":{"genus":"Zygocera","species":"norfolkensis","authorship":{"verbatim":"Mc'Keown 1938","normalized":"Mc'Keown 1938","year":"1938","authors":["Mc'Keown"],"originalAuth":{"authors":["Mc'Keown"],"year":{"year":"1938"}}}}},"words":[{"verbatim":"Zygocera","normalized":"Zygocera","wordType":"GENUS","start":0,"end":8},{"verbatim":"norfolkensis","normalized":"norfolkensis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Mc'Keown","normalized":"Mc'Keown","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1938","normalized":"1938","wordType":"YEAR","start":31,"end":35}],"id":"b1dda8e1-2e48-56e7-a508-0a4dd8372a9e","parserVersion":"test_version"}
```

### Infraspecies without rank (ICZN)

Name: Peristernia nassatula forskali Tapparone-Canefri 1875

Canonical: Peristernia nassatula forskali

Authorship: Tapparone-Canefri 1875

```json
{"parsed":true,"quality":1,"verbatim":"Peristernia nassatula forskali Tapparone-Canefri 1875","normalized":"Peristernia nassatula forskali Tapparone-Canefri 1875","canonical":{"stemmed":"Peristernia nassatul forskal","simple":"Peristernia nassatula forskali","full":"Peristernia nassatula forskali"},"cardinality":3,"authorship":{"verbatim":"Tapparone-Canefri 1875","normalized":"Tapparone-Canefri 1875","year":"1875","authors":["Tapparone-Canefri"],"originalAuth":{"authors":["Tapparone-Canefri"],"year":{"year":"1875"}}},"details":{"infraspecies":{"genus":"Peristernia","species":"nassatula","infraspecies":[{"value":"forskali","authorship":{"verbatim":"Tapparone-Canefri 1875","normalized":"Tapparone-Canefri 1875","year":"1875","authors":["Tapparone-Canefri"],"originalAuth":{"authors":["Tapparone-Canefri"],"year":{"year":"1875"}}}}]}},"words":[{"verbatim":"Peristernia","normalized":"Peristernia","wordType":"GENUS","start":0,"end":11},{"verbatim":"nassatula","normalized":"nassatula","wordType":"SPECIES","start":12,"end":21},{"verbatim":"forskali","normalized":"forskali","wordType":"INFRASPECIES","start":22,"end":30},{"verbatim":"Tapparone-Canefri","normalized":"Tapparone-Canefri","wordType":"AUTHOR_WORD","start":31,"end":48},{"verbatim":"1875","normalized":"1875","wordType":"YEAR","start":49,"end":53}],"id":"5aa39b53-32ee-5e9f-aa29-c268a9662fd7","parserVersion":"test_version"}
```

Name: Cypraeovula (Luponia) amphithales perdentata

Canonical: Cypraeovula amphithales perdentata

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Cypraeovula (Luponia) amphithales perdentata","normalized":"Cypraeovula (Luponia) amphithales perdentata","canonical":{"stemmed":"Cypraeovula amphithal perdentat","simple":"Cypraeovula amphithales perdentata","full":"Cypraeovula amphithales perdentata"},"cardinality":3,"details":{"infraspecies":{"genus":"Cypraeovula","subgenus":"Luponia","species":"amphithales","infraspecies":[{"value":"perdentata"}]}},"words":[{"verbatim":"Cypraeovula","normalized":"Cypraeovula","wordType":"GENUS","start":0,"end":11},{"verbatim":"Luponia","normalized":"Luponia","wordType":"INFRA_GENUS","start":13,"end":20},{"verbatim":"amphithales","normalized":"amphithales","wordType":"SPECIES","start":22,"end":33},{"verbatim":"perdentata","normalized":"perdentata","wordType":"INFRASPECIES","start":34,"end":44}],"id":"d05be4e3-a0e3-5af4-9104-7922df1bcb47","parserVersion":"test_version"}
```

Name: Triticum repens vulgäre

Canonical: Triticum repens vulgaere

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Triticum repens vulgäre","normalized":"Triticum repens vulgaere","canonical":{"stemmed":"Triticum repens uulgaer","simple":"Triticum repens vulgaere","full":"Triticum repens vulgaere"},"cardinality":3,"details":{"infraspecies":{"genus":"Triticum","species":"repens","infraspecies":[{"value":"vulgaere"}]}},"words":[{"verbatim":"Triticum","normalized":"Triticum","wordType":"GENUS","start":0,"end":8},{"verbatim":"repens","normalized":"repens","wordType":"SPECIES","start":9,"end":15},{"verbatim":"vulgäre","normalized":"vulgaere","wordType":"INFRASPECIES","start":16,"end":23}],"id":"5fb6ae9c-d7be-5d81-88b8-3c96d4c48a74","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Batsch) K. A. Harrison 1961

```json
{"parsed":true,"quality":1,"verbatim":"Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961","normalized":"Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Batsch) K. A. Harrison 1961","normalized":"(Batsch) K. A. Harrison 1961","authors":["Batsch","K. A. Harrison"],"originalAuth":{"authors":["Batsch"]},"combinationAuth":{"authors":["K. A. Harrison"],"year":{"year":"1961"}}},"details":{"infraspecies":{"genus":"Hydnellum","species":"scrobiculatum","infraspecies":[{"value":"zonatum","authorship":{"verbatim":"(Batsch) K. A. Harrison 1961","normalized":"(Batsch) K. A. Harrison 1961","authors":["Batsch","K. A. Harrison"],"originalAuth":{"authors":["Batsch"]},"combinationAuth":{"authors":["K. A. Harrison"],"year":{"year":"1961"}}}}]}},"words":[{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":0,"end":9},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":10,"end":23},{"verbatim":"zonatum","normalized":"zonatum","wordType":"INFRASPECIES","start":24,"end":31},{"verbatim":"Batsch","normalized":"Batsch","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"K.","normalized":"K.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"Harrison","normalized":"Harrison","wordType":"AUTHOR_WORD","start":47,"end":55},{"verbatim":"1961","normalized":"1961","wordType":"YEAR","start":56,"end":60}],"id":"8368c11a-7c1b-5e82-bdad-a4887bfa81d2","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"quality":1,"verbatim":"Hydnellum scrobiculatum zonatum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum scrobiculatum zonatum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraspecies":{"genus":"Hydnellum","species":"scrobiculatum","infraspecies":[{"value":"zonatum","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"words":[{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":0,"end":9},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":10,"end":23},{"verbatim":"zonatum","normalized":"zonatum","wordType":"INFRASPECIES","start":24,"end":31},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"Hall","normalized":"Hall","wordType":"AUTHOR_WORD","start":44,"end":48},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":53,"end":55},{"verbatim":"Stuntz","normalized":"Stuntz","wordType":"AUTHOR_WORD","start":56,"end":62},{"verbatim":"1972","normalized":"1972","wordType":"YEAR","start":63,"end":67}],"id":"fa3448c6-168e-575f-a6eb-c5adc6f3e89d","parserVersion":"test_version"}
```

Name: Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"quality":1,"verbatim":"Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraspecies":{"genus":"Hydnellum","subgenus":"Hydnellum","species":"scrobiculatum","infraspecies":[{"value":"zonatum","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"words":[{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":0,"end":9},{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"INFRA_GENUS","start":11,"end":20},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":22,"end":35},{"verbatim":"zonatum","normalized":"zonatum","wordType":"INFRASPECIES","start":36,"end":43},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":53,"end":55},{"verbatim":"Hall","normalized":"Hall","wordType":"AUTHOR_WORD","start":56,"end":60},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":63,"end":65},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":65,"end":67},{"verbatim":"Stuntz","normalized":"Stuntz","wordType":"AUTHOR_WORD","start":68,"end":74},{"verbatim":"1972","normalized":"1972","wordType":"YEAR","start":75,"end":79}],"id":"14e5eb1f-82a3-598c-9ada-3a9a20ab54cc","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum

Canonical: Hydnellum scrobiculatum zonatum

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Hydnellum scrobiculatum zonatum","normalized":"Hydnellum scrobiculatum zonatum","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"details":{"infraspecies":{"genus":"Hydnellum","species":"scrobiculatum","infraspecies":[{"value":"zonatum"}]}},"words":[{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":0,"end":9},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":10,"end":23},{"verbatim":"zonatum","normalized":"zonatum","wordType":"INFRASPECIES","start":24,"end":31}],"id":"22af845f-773e-502e-be46-ac73ae5960be","parserVersion":"test_version"}
```

Name: Mus musculus hortulanus

Canonical: Mus musculus hortulanus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Mus musculus hortulanus","normalized":"Mus musculus hortulanus","canonical":{"stemmed":"Mus muscul hortulan","simple":"Mus musculus hortulanus","full":"Mus musculus hortulanus"},"cardinality":3,"details":{"infraspecies":{"genus":"Mus","species":"musculus","infraspecies":[{"value":"hortulanus"}]}},"words":[{"verbatim":"Mus","normalized":"Mus","wordType":"GENUS","start":0,"end":3},{"verbatim":"musculus","normalized":"musculus","wordType":"SPECIES","start":4,"end":12},{"verbatim":"hortulanus","normalized":"hortulanus","wordType":"INFRASPECIES","start":13,"end":23}],"id":"5fd9a4aa-9fa8-5200-909a-6c9ec8a9a088","parserVersion":"test_version"}
```

Name: Ortygospiza atricollis mülleri

Canonical: Ortygospiza atricollis muelleri

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Ortygospiza atricollis mülleri","normalized":"Ortygospiza atricollis muelleri","canonical":{"stemmed":"Ortygospiza atricoll mueller","simple":"Ortygospiza atricollis muelleri","full":"Ortygospiza atricollis muelleri"},"cardinality":3,"details":{"infraspecies":{"genus":"Ortygospiza","species":"atricollis","infraspecies":[{"value":"muelleri"}]}},"words":[{"verbatim":"Ortygospiza","normalized":"Ortygospiza","wordType":"GENUS","start":0,"end":11},{"verbatim":"atricollis","normalized":"atricollis","wordType":"SPECIES","start":12,"end":22},{"verbatim":"mülleri","normalized":"muelleri","wordType":"INFRASPECIES","start":23,"end":30}],"id":"1ee6bf1d-90d8-5c4b-98c1-2646c301d07c","parserVersion":"test_version"}
```

Name: Cortinarius angulatus B gracilescens Fr. 1838

Canonical: Cortinarius angulatus gracilescens

Authorship: Fr. 1838

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Author is too short"}],"verbatim":"Cortinarius angulatus B gracilescens Fr. 1838","normalized":"Cortinarius angulatus B gracilescens Fr. 1838","canonical":{"stemmed":"Cortinarius angulat gracilescens","simple":"Cortinarius angulatus gracilescens","full":"Cortinarius angulatus gracilescens"},"cardinality":3,"authorship":{"verbatim":"Fr. 1838","normalized":"Fr. 1838","year":"1838","authors":["Fr."],"originalAuth":{"authors":["Fr."],"year":{"year":"1838"}}},"details":{"infraspecies":{"genus":"Cortinarius","species":"angulatus","authorship":{"verbatim":"B","normalized":"B","authors":["B"],"originalAuth":{"authors":["B"]}},"infraspecies":[{"value":"gracilescens","authorship":{"verbatim":"Fr. 1838","normalized":"Fr. 1838","year":"1838","authors":["Fr."],"originalAuth":{"authors":["Fr."],"year":{"year":"1838"}}}}]}},"words":[{"verbatim":"Cortinarius","normalized":"Cortinarius","wordType":"GENUS","start":0,"end":11},{"verbatim":"angulatus","normalized":"angulatus","wordType":"SPECIES","start":12,"end":21},{"verbatim":"B","normalized":"B","wordType":"AUTHOR_WORD","start":22,"end":23},{"verbatim":"gracilescens","normalized":"gracilescens","wordType":"INFRASPECIES","start":24,"end":36},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":37,"end":40},{"verbatim":"1838","normalized":"1838","wordType":"YEAR","start":41,"end":45}],"id":"3fb101ad-d05e-5648-993b-bfbb8c76166e","parserVersion":"test_version"}
```

Name: Caulerpa fastigiata confervoides P. L. Crouan & H. M. Crouan ex Weber-van Bosse

Canonical: Caulerpa fastigiata confervoides

Authorship: P. L. Crouan & H. M. Crouan ex Weber-van Bosse

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Caulerpa fastigiata confervoides P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"Caulerpa fastigiata confervoides P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","canonical":{"stemmed":"Caulerpa fastigiat conferuoid","simple":"Caulerpa fastigiata confervoides","full":"Caulerpa fastigiata confervoides"},"cardinality":3,"authorship":{"verbatim":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","authors":["P. L. Crouan","H. M. Crouan","Weber-van Bosse"],"originalAuth":{"authors":["P. L. Crouan","H. M. Crouan"],"exAuthors":{"authors":["Weber-van Bosse"]}}},"details":{"infraspecies":{"genus":"Caulerpa","species":"fastigiata","infraspecies":[{"value":"confervoides","authorship":{"verbatim":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","authors":["P. L. Crouan","H. M. Crouan","Weber-van Bosse"],"originalAuth":{"authors":["P. L. Crouan","H. M. Crouan"],"exAuthors":{"authors":["Weber-van Bosse"]}}}}]}},"words":[{"verbatim":"Caulerpa","normalized":"Caulerpa","wordType":"GENUS","start":0,"end":8},{"verbatim":"fastigiata","normalized":"fastigiata","wordType":"SPECIES","start":9,"end":19},{"verbatim":"confervoides","normalized":"confervoides","wordType":"INFRASPECIES","start":20,"end":32},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"Crouan","normalized":"Crouan","wordType":"AUTHOR_WORD","start":39,"end":45},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":48,"end":50},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"Crouan","normalized":"Crouan","wordType":"AUTHOR_WORD","start":54,"end":60},{"verbatim":"Weber-van","normalized":"Weber-van","wordType":"AUTHOR_WORD","start":64,"end":73},{"verbatim":"Bosse","normalized":"Bosse","wordType":"AUTHOR_WORD","start":74,"end":79}],"id":"8934dbda-1fd2-52c4-af76-8f80e5f02791","parserVersion":"test_version"}
```

Name: Rhinanthus glacialis simplex(Sterneck) J.Dostál

Canonical: Rhinanthus glacialis simplex

Authorship: (Sterneck) J. Dostál

```json
{"parsed":true,"quality":1,"verbatim":"Rhinanthus glacialis simplex(Sterneck) J.Dostál","normalized":"Rhinanthus glacialis simplex (Sterneck) J. Dostál","canonical":{"stemmed":"Rhinanthus glacial simplex","simple":"Rhinanthus glacialis simplex","full":"Rhinanthus glacialis simplex"},"cardinality":3,"authorship":{"verbatim":"(Sterneck) J.Dostál","normalized":"(Sterneck) J. Dostál","authors":["Sterneck","J. Dostál"],"originalAuth":{"authors":["Sterneck"]},"combinationAuth":{"authors":["J. Dostál"]}},"details":{"infraspecies":{"genus":"Rhinanthus","species":"glacialis","infraspecies":[{"value":"simplex","authorship":{"verbatim":"(Sterneck) J.Dostál","normalized":"(Sterneck) J. Dostál","authors":["Sterneck","J. Dostál"],"originalAuth":{"authors":["Sterneck"]},"combinationAuth":{"authors":["J. Dostál"]}}}]}},"words":[{"verbatim":"Rhinanthus","normalized":"Rhinanthus","wordType":"GENUS","start":0,"end":10},{"verbatim":"glacialis","normalized":"glacialis","wordType":"SPECIES","start":11,"end":20},{"verbatim":"simplex","normalized":"simplex","wordType":"INFRASPECIES","start":21,"end":28},{"verbatim":"Sterneck","normalized":"Sterneck","wordType":"AUTHOR_WORD","start":29,"end":37},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"Dostál","normalized":"Dostál","wordType":"AUTHOR_WORD","start":41,"end":47}],"id":"8128607d-0186-5a38-ab02-c0b18f46b3ed","parserVersion":"test_version"}
```

### Legacy ICZN names with rank

Name: Acipenser gueldenstaedti colchicus natio danubicus Movchan, 1967

Canonical: Acipenser gueldenstaedti colchicus natio danubicus

Authorship: Movchan 1967

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Acipenser gueldenstaedti colchicus natio danubicus Movchan, 1967","normalized":"Acipenser gueldenstaedti colchicus natio danubicus Movchan 1967","canonical":{"stemmed":"Acipenser gueldenstaedt colchic danubic","simple":"Acipenser gueldenstaedti colchicus danubicus","full":"Acipenser gueldenstaedti colchicus natio danubicus"},"cardinality":4,"authorship":{"verbatim":"Movchan, 1967","normalized":"Movchan 1967","year":"1967","authors":["Movchan"],"originalAuth":{"authors":["Movchan"],"year":{"year":"1967"}}},"details":{"infraspecies":{"genus":"Acipenser","species":"gueldenstaedti","infraspecies":[{"value":"colchicus"},{"value":"danubicus","rank":"natio","authorship":{"verbatim":"Movchan, 1967","normalized":"Movchan 1967","year":"1967","authors":["Movchan"],"originalAuth":{"authors":["Movchan"],"year":{"year":"1967"}}}}]}},"words":[{"verbatim":"Acipenser","normalized":"Acipenser","wordType":"GENUS","start":0,"end":9},{"verbatim":"gueldenstaedti","normalized":"gueldenstaedti","wordType":"SPECIES","start":10,"end":24},{"verbatim":"colchicus","normalized":"colchicus","wordType":"INFRASPECIES","start":25,"end":34},{"verbatim":"natio","normalized":"natio","wordType":"RANK","start":35,"end":40},{"verbatim":"danubicus","normalized":"danubicus","wordType":"INFRASPECIES","start":41,"end":50},{"verbatim":"Movchan","normalized":"Movchan","wordType":"AUTHOR_WORD","start":51,"end":58},{"verbatim":"1967","normalized":"1967","wordType":"YEAR","start":60,"end":64}],"id":"d572e7a6-bcbd-59ef-bc60-1e5d659fd51c","parserVersion":"test_version"}
```

### Infraspecies with rank (ICN)

Name: Cantharellus sinuosus var. multiplex(A.H.Sm.) Romagn., 1995

Canonical: Cantharellus sinuosus var. multiplex

Authorship: (A. H. Sm.) Romagn. 1995

```json
{"parsed":true,"quality":1,"verbatim":"Cantharellus sinuosus var. multiplex(A.H.Sm.) Romagn., 1995","normalized":"Cantharellus sinuosus var. multiplex (A. H. Sm.) Romagn. 1995","canonical":{"stemmed":"Cantharellus sinuos multiplex","simple":"Cantharellus sinuosus multiplex","full":"Cantharellus sinuosus var. multiplex"},"cardinality":3,"authorship":{"verbatim":"(A.H.Sm.) Romagn., 1995","normalized":"(A. H. Sm.) Romagn. 1995","authors":["A. H. Sm.","Romagn."],"originalAuth":{"authors":["A. H. Sm."]},"combinationAuth":{"authors":["Romagn."],"year":{"year":"1995"}}},"details":{"infraspecies":{"genus":"Cantharellus","species":"sinuosus","infraspecies":[{"value":"multiplex","rank":"var.","authorship":{"verbatim":"(A.H.Sm.) Romagn., 1995","normalized":"(A. H. Sm.) Romagn. 1995","authors":["A. H. Sm.","Romagn."],"originalAuth":{"authors":["A. H. Sm."]},"combinationAuth":{"authors":["Romagn."],"year":{"year":"1995"}}}}]}},"words":[{"verbatim":"Cantharellus","normalized":"Cantharellus","wordType":"GENUS","start":0,"end":12},{"verbatim":"sinuosus","normalized":"sinuosus","wordType":"SPECIES","start":13,"end":21},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":22,"end":26},{"verbatim":"multiplex","normalized":"multiplex","wordType":"INFRASPECIES","start":27,"end":36},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"Sm.","normalized":"Sm.","wordType":"AUTHOR_WORD","start":41,"end":44},{"verbatim":"Romagn.","normalized":"Romagn.","wordType":"AUTHOR_WORD","start":46,"end":53},{"verbatim":"1995","normalized":"1995","wordType":"YEAR","start":55,"end":59}],"id":"46007c97-3458-58c7-aea8-2413b74449d9","parserVersion":"test_version"}
```

Name: Crematogaster impressa st. brazzai Santschi 1937

Canonical: Crematogaster impressa st. brazzai

Authorship: Santschi 1937

```json
{"parsed":true,"quality":1,"verbatim":"Crematogaster impressa st. brazzai Santschi 1937","normalized":"Crematogaster impressa st. brazzai Santschi 1937","canonical":{"stemmed":"Crematogaster impress brazza","simple":"Crematogaster impressa brazzai","full":"Crematogaster impressa st. brazzai"},"cardinality":3,"authorship":{"verbatim":"Santschi 1937","normalized":"Santschi 1937","year":"1937","authors":["Santschi"],"originalAuth":{"authors":["Santschi"],"year":{"year":"1937"}}},"details":{"infraspecies":{"genus":"Crematogaster","species":"impressa","infraspecies":[{"value":"brazzai","rank":"st.","authorship":{"verbatim":"Santschi 1937","normalized":"Santschi 1937","year":"1937","authors":["Santschi"],"originalAuth":{"authors":["Santschi"],"year":{"year":"1937"}}}}]}},"words":[{"verbatim":"Crematogaster","normalized":"Crematogaster","wordType":"GENUS","start":0,"end":13},{"verbatim":"impressa","normalized":"impressa","wordType":"SPECIES","start":14,"end":22},{"verbatim":"st.","normalized":"st.","wordType":"RANK","start":23,"end":26},{"verbatim":"brazzai","normalized":"brazzai","wordType":"INFRASPECIES","start":27,"end":34},{"verbatim":"Santschi","normalized":"Santschi","wordType":"AUTHOR_WORD","start":35,"end":43},{"verbatim":"1937","normalized":"1937","wordType":"YEAR","start":44,"end":48}],"id":"853d0cff-b499-5d38-ae49-75b558f9ddf0","parserVersion":"test_version"}
```

<!-- badly formed name, we do not deal with it for now -->
Name: Cibotium st.-johnii Krajina

Canonical: Cibotium st-johnii

Authorship: Krajina

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Period character is not allowed in canonical"}],"verbatim":"Cibotium st.-johnii Krajina","normalized":"Cibotium st-johnii Krajina","canonical":{"stemmed":"Cibotium st-iohn","simple":"Cibotium st-johnii","full":"Cibotium st-johnii"},"cardinality":2,"authorship":{"verbatim":"Krajina","normalized":"Krajina","authors":["Krajina"],"originalAuth":{"authors":["Krajina"]}},"details":{"species":{"genus":"Cibotium","species":"st-johnii","authorship":{"verbatim":"Krajina","normalized":"Krajina","authors":["Krajina"],"originalAuth":{"authors":["Krajina"]}}}},"words":[{"verbatim":"Cibotium","normalized":"Cibotium","wordType":"GENUS","start":0,"end":8},{"verbatim":"st.-johnii","normalized":"st-johnii","wordType":"SPECIES","start":9,"end":19},{"verbatim":"Krajina","normalized":"Krajina","wordType":"AUTHOR_WORD","start":20,"end":27}],"id":"6b34256d-6c3b-5870-a781-77eeac49b6c4","parserVersion":"test_version"}
```

Name: Camponotus conspicuus st. zonatus

Canonical: Camponotus conspicuus st. zonatus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Camponotus conspicuus st. zonatus","normalized":"Camponotus conspicuus st. zonatus","canonical":{"stemmed":"Camponotus conspicu zonat","simple":"Camponotus conspicuus zonatus","full":"Camponotus conspicuus st. zonatus"},"cardinality":3,"details":{"infraspecies":{"genus":"Camponotus","species":"conspicuus","infraspecies":[{"value":"zonatus","rank":"st."}]}},"words":[{"verbatim":"Camponotus","normalized":"Camponotus","wordType":"GENUS","start":0,"end":10},{"verbatim":"conspicuus","normalized":"conspicuus","wordType":"SPECIES","start":11,"end":21},{"verbatim":"st.","normalized":"st.","wordType":"RANK","start":22,"end":25},{"verbatim":"zonatus","normalized":"zonatus","wordType":"INFRASPECIES","start":26,"end":33}],"id":"67364c72-53e0-54d3-9795-f04fd1938d75","parserVersion":"test_version"}
```

Name: Fagus sylvatica subsp. orientalis (Lipsky) Greuter & Burdet

Canonical: Fagus sylvatica subsp. orientalis

Authorship: (Lipsky) Greuter & Burdet

```json
{"parsed":true,"quality":1,"verbatim":"Fagus sylvatica subsp. orientalis (Lipsky) Greuter \u0026 Burdet","normalized":"Fagus sylvatica subsp. orientalis (Lipsky) Greuter \u0026 Burdet","canonical":{"stemmed":"Fagus syluatic oriental","simple":"Fagus sylvatica orientalis","full":"Fagus sylvatica subsp. orientalis"},"cardinality":3,"authorship":{"verbatim":"(Lipsky) Greuter \u0026 Burdet","normalized":"(Lipsky) Greuter \u0026 Burdet","authors":["Lipsky","Greuter","Burdet"],"originalAuth":{"authors":["Lipsky"]},"combinationAuth":{"authors":["Greuter","Burdet"]}},"details":{"infraspecies":{"genus":"Fagus","species":"sylvatica","infraspecies":[{"value":"orientalis","rank":"subsp.","authorship":{"verbatim":"(Lipsky) Greuter \u0026 Burdet","normalized":"(Lipsky) Greuter \u0026 Burdet","authors":["Lipsky","Greuter","Burdet"],"originalAuth":{"authors":["Lipsky"]},"combinationAuth":{"authors":["Greuter","Burdet"]}}}]}},"words":[{"verbatim":"Fagus","normalized":"Fagus","wordType":"GENUS","start":0,"end":5},{"verbatim":"sylvatica","normalized":"sylvatica","wordType":"SPECIES","start":6,"end":15},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":16,"end":22},{"verbatim":"orientalis","normalized":"orientalis","wordType":"INFRASPECIES","start":23,"end":33},{"verbatim":"Lipsky","normalized":"Lipsky","wordType":"AUTHOR_WORD","start":35,"end":41},{"verbatim":"Greuter","normalized":"Greuter","wordType":"AUTHOR_WORD","start":43,"end":50},{"verbatim":"Burdet","normalized":"Burdet","wordType":"AUTHOR_WORD","start":53,"end":59}],"id":"f0bff1a3-0923-58d1-807f-c5da5b85531e","parserVersion":"test_version"}
```

Name: Tillandsia utriculata subspec. utriculata

Canonical: Tillandsia utriculata subsp. utriculata

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Tillandsia utriculata subspec. utriculata","normalized":"Tillandsia utriculata subsp. utriculata","canonical":{"stemmed":"Tillandsia utriculat utriculat","simple":"Tillandsia utriculata utriculata","full":"Tillandsia utriculata subsp. utriculata"},"cardinality":3,"details":{"infraspecies":{"genus":"Tillandsia","species":"utriculata","infraspecies":[{"value":"utriculata","rank":"subsp."}]}},"words":[{"verbatim":"Tillandsia","normalized":"Tillandsia","wordType":"GENUS","start":0,"end":10},{"verbatim":"utriculata","normalized":"utriculata","wordType":"SPECIES","start":11,"end":21},{"verbatim":"subspec.","normalized":"subsp.","wordType":"RANK","start":22,"end":30},{"verbatim":"utriculata","normalized":"utriculata","wordType":"INFRASPECIES","start":31,"end":41}],"id":"fa612e5d-f697-5227-a5a0-fdb4a1aafe7a","parserVersion":"test_version"}
```

Name: Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.

Canonical: Prunus mexicana var. reticulata

Authorship: (Sarg.) Sarg.

```json
{"parsed":true,"quality":1,"verbatim":"Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.","normalized":"Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.","canonical":{"stemmed":"Prunus mexican reticulat","simple":"Prunus mexicana reticulata","full":"Prunus mexicana var. reticulata"},"cardinality":3,"authorship":{"verbatim":"(Sarg.) Sarg.","normalized":"(Sarg.) Sarg.","authors":["Sarg."],"originalAuth":{"authors":["Sarg."]},"combinationAuth":{"authors":["Sarg."]}},"details":{"infraspecies":{"genus":"Prunus","species":"mexicana","authorship":{"verbatim":"S. Watson","normalized":"S. Watson","authors":["S. Watson"],"originalAuth":{"authors":["S. Watson"]}},"infraspecies":[{"value":"reticulata","rank":"var.","authorship":{"verbatim":"(Sarg.) Sarg.","normalized":"(Sarg.) Sarg.","authors":["Sarg."],"originalAuth":{"authors":["Sarg."]},"combinationAuth":{"authors":["Sarg."]}}}]}},"words":[{"verbatim":"Prunus","normalized":"Prunus","wordType":"GENUS","start":0,"end":6},{"verbatim":"mexicana","normalized":"mexicana","wordType":"SPECIES","start":7,"end":15},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"Watson","normalized":"Watson","wordType":"AUTHOR_WORD","start":19,"end":25},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":26,"end":30},{"verbatim":"reticulata","normalized":"reticulata","wordType":"INFRASPECIES","start":31,"end":41},{"verbatim":"Sarg.","normalized":"Sarg.","wordType":"AUTHOR_WORD","start":43,"end":48},{"verbatim":"Sarg.","normalized":"Sarg.","wordType":"AUTHOR_WORD","start":50,"end":55}],"id":"5ba1cc96-ab40-51b3-951d-f91b5bff1da8","parserVersion":"test_version"}
```

Name: Potamogeton iilinoensis var. ventanicola

Canonical: Potamogeton iilinoensis var. ventanicola

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Potamogeton iilinoensis var. ventanicola","normalized":"Potamogeton iilinoensis var. ventanicola","canonical":{"stemmed":"Potamogeton iilinoens uentanicol","simple":"Potamogeton iilinoensis ventanicola","full":"Potamogeton iilinoensis var. ventanicola"},"cardinality":3,"details":{"infraspecies":{"genus":"Potamogeton","species":"iilinoensis","infraspecies":[{"value":"ventanicola","rank":"var."}]}},"words":[{"verbatim":"Potamogeton","normalized":"Potamogeton","wordType":"GENUS","start":0,"end":11},{"verbatim":"iilinoensis","normalized":"iilinoensis","wordType":"SPECIES","start":12,"end":23},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":24,"end":28},{"verbatim":"ventanicola","normalized":"ventanicola","wordType":"INFRASPECIES","start":29,"end":40}],"id":"edf418ec-98b3-52fb-a8de-26808b61c50f","parserVersion":"test_version"}
```

Name: Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien

Canonical: Potamogeton iilinoensis var. ventanicola

Authorship: (Hicken) Horn af Rantzien

```json
{"parsed":true,"quality":1,"verbatim":"Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien","normalized":"Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien","canonical":{"stemmed":"Potamogeton iilinoens uentanicol","simple":"Potamogeton iilinoensis ventanicola","full":"Potamogeton iilinoensis var. ventanicola"},"cardinality":3,"authorship":{"verbatim":"(Hicken) Horn af Rantzien","normalized":"(Hicken) Horn af Rantzien","authors":["Hicken","Horn af Rantzien"],"originalAuth":{"authors":["Hicken"]},"combinationAuth":{"authors":["Horn af Rantzien"]}},"details":{"infraspecies":{"genus":"Potamogeton","species":"iilinoensis","infraspecies":[{"value":"ventanicola","rank":"var.","authorship":{"verbatim":"(Hicken) Horn af Rantzien","normalized":"(Hicken) Horn af Rantzien","authors":["Hicken","Horn af Rantzien"],"originalAuth":{"authors":["Hicken"]},"combinationAuth":{"authors":["Horn af Rantzien"]}}}]}},"words":[{"verbatim":"Potamogeton","normalized":"Potamogeton","wordType":"GENUS","start":0,"end":11},{"verbatim":"iilinoensis","normalized":"iilinoensis","wordType":"SPECIES","start":12,"end":23},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":24,"end":28},{"verbatim":"ventanicola","normalized":"ventanicola","wordType":"INFRASPECIES","start":29,"end":40},{"verbatim":"Hicken","normalized":"Hicken","wordType":"AUTHOR_WORD","start":42,"end":48},{"verbatim":"Horn","normalized":"Horn","wordType":"AUTHOR_WORD","start":50,"end":54},{"verbatim":"af","normalized":"af","wordType":"AUTHOR_WORD","start":55,"end":57},{"verbatim":"Rantzien","normalized":"Rantzien","wordType":"AUTHOR_WORD","start":58,"end":66}],"id":"e7888abd-4365-5d74-8d5f-a69c8196328e","parserVersion":"test_version"}
```

Name: Triticum repens var. vulgäre

Canonical: Triticum repens var. vulgaere

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Triticum repens var. vulgäre","normalized":"Triticum repens var. vulgaere","canonical":{"stemmed":"Triticum repens uulgaer","simple":"Triticum repens vulgaere","full":"Triticum repens var. vulgaere"},"cardinality":3,"details":{"infraspecies":{"genus":"Triticum","species":"repens","infraspecies":[{"value":"vulgaere","rank":"var."}]}},"words":[{"verbatim":"Triticum","normalized":"Triticum","wordType":"GENUS","start":0,"end":8},{"verbatim":"repens","normalized":"repens","wordType":"SPECIES","start":9,"end":15},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":16,"end":20},{"verbatim":"vulgäre","normalized":"vulgaere","wordType":"INFRASPECIES","start":21,"end":28}],"id":"3421b13b-aaa9-5234-bc1d-9d3fe7a6b19e","parserVersion":"test_version"}
```

Name: Aus bus Linn. var. bus

Canonical: Aus bus var. bus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aus bus Linn. var. bus","normalized":"Aus bus Linn. var. bus","canonical":{"stemmed":"Aus bus bus","simple":"Aus bus bus","full":"Aus bus var. bus"},"cardinality":3,"details":{"infraspecies":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Linn.","normalized":"Linn.","authors":["Linn."],"originalAuth":{"authors":["Linn."]}},"infraspecies":[{"value":"bus","rank":"var."}]}},"words":[{"verbatim":"Aus","normalized":"Aus","wordType":"GENUS","start":0,"end":3},{"verbatim":"bus","normalized":"bus","wordType":"SPECIES","start":4,"end":7},{"verbatim":"Linn.","normalized":"Linn.","wordType":"AUTHOR_WORD","start":8,"end":13},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":14,"end":18},{"verbatim":"bus","normalized":"bus","wordType":"INFRASPECIES","start":19,"end":22}],"id":"2a6e45e2-5737-514b-8055-06f8a878dd36","parserVersion":"test_version"}
```

Name: Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987

Canonical: Agalinis purpurea var. borealis

Authorship: (Berg.) Peterson 1987

```json
{"parsed":true,"quality":1,"verbatim":"Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987","normalized":"Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987","canonical":{"stemmed":"Agalinis purpure boreal","simple":"Agalinis purpurea borealis","full":"Agalinis purpurea var. borealis"},"cardinality":3,"authorship":{"verbatim":"(Berg.) Peterson 1987","normalized":"(Berg.) Peterson 1987","authors":["Berg.","Peterson"],"originalAuth":{"authors":["Berg."]},"combinationAuth":{"authors":["Peterson"],"year":{"year":"1987"}}},"details":{"infraspecies":{"genus":"Agalinis","species":"purpurea","authorship":{"verbatim":"(L.) Briton","normalized":"(L.) Briton","authors":["L.","Briton"],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Briton"]}},"infraspecies":[{"value":"borealis","rank":"var.","authorship":{"verbatim":"(Berg.) Peterson 1987","normalized":"(Berg.) Peterson 1987","authors":["Berg.","Peterson"],"originalAuth":{"authors":["Berg."]},"combinationAuth":{"authors":["Peterson"],"year":{"year":"1987"}}}}]}},"words":[{"verbatim":"Agalinis","normalized":"Agalinis","wordType":"GENUS","start":0,"end":8},{"verbatim":"purpurea","normalized":"purpurea","wordType":"SPECIES","start":9,"end":17},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":19,"end":21},{"verbatim":"Briton","normalized":"Briton","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":30,"end":34},{"verbatim":"borealis","normalized":"borealis","wordType":"INFRASPECIES","start":35,"end":43},{"verbatim":"Berg.","normalized":"Berg.","wordType":"AUTHOR_WORD","start":45,"end":50},{"verbatim":"Peterson","normalized":"Peterson","wordType":"AUTHOR_WORD","start":52,"end":60},{"verbatim":"1987","normalized":"1987","wordType":"YEAR","start":61,"end":65}],"id":"769863cd-7c9d-5d4a-bf5c-fb6903a96431","parserVersion":"test_version"}
```

Name: Callideriphus flavicollis morph. reductus Fuchs 1961

Canonical: Callideriphus flavicollis morph. reductus

Authorship: Fuchs 1961

```json
{"parsed":true,"quality":1,"verbatim":"Callideriphus flavicollis morph. reductus Fuchs 1961","normalized":"Callideriphus flavicollis morph. reductus Fuchs 1961","canonical":{"stemmed":"Callideriphus flauicoll reduct","simple":"Callideriphus flavicollis reductus","full":"Callideriphus flavicollis morph. reductus"},"cardinality":3,"authorship":{"verbatim":"Fuchs 1961","normalized":"Fuchs 1961","year":"1961","authors":["Fuchs"],"originalAuth":{"authors":["Fuchs"],"year":{"year":"1961"}}},"details":{"infraspecies":{"genus":"Callideriphus","species":"flavicollis","infraspecies":[{"value":"reductus","rank":"morph.","authorship":{"verbatim":"Fuchs 1961","normalized":"Fuchs 1961","year":"1961","authors":["Fuchs"],"originalAuth":{"authors":["Fuchs"],"year":{"year":"1961"}}}}]}},"words":[{"verbatim":"Callideriphus","normalized":"Callideriphus","wordType":"GENUS","start":0,"end":13},{"verbatim":"flavicollis","normalized":"flavicollis","wordType":"SPECIES","start":14,"end":25},{"verbatim":"morph.","normalized":"morph.","wordType":"RANK","start":26,"end":32},{"verbatim":"reductus","normalized":"reductus","wordType":"INFRASPECIES","start":33,"end":41},{"verbatim":"Fuchs","normalized":"Fuchs","wordType":"AUTHOR_WORD","start":42,"end":47},{"verbatim":"1961","normalized":"1961","wordType":"YEAR","start":48,"end":52}],"id":"2b01f892-dbb3-5776-870a-c6cb8f09f2bc","parserVersion":"test_version"}
```

Name: Caulerpa cupressoides forma nuda

Canonical: Caulerpa cupressoides f. nuda

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Caulerpa cupressoides forma nuda","normalized":"Caulerpa cupressoides f. nuda","canonical":{"stemmed":"Caulerpa cupressoid nud","simple":"Caulerpa cupressoides nuda","full":"Caulerpa cupressoides f. nuda"},"cardinality":3,"details":{"infraspecies":{"genus":"Caulerpa","species":"cupressoides","infraspecies":[{"value":"nuda","rank":"f."}]}},"words":[{"verbatim":"Caulerpa","normalized":"Caulerpa","wordType":"GENUS","start":0,"end":8},{"verbatim":"cupressoides","normalized":"cupressoides","wordType":"SPECIES","start":9,"end":21},{"verbatim":"forma","normalized":"f.","wordType":"RANK","start":22,"end":27},{"verbatim":"nuda","normalized":"nuda","wordType":"INFRASPECIES","start":28,"end":32}],"id":"805ee92d-001e-5f05-abad-446f683860cb","parserVersion":"test_version"}
```

Name: Chlorocyperus glaber form. fasciculariforme (Lojac.) Soó

Canonical: Chlorocyperus glaber f. fasciculariforme

Authorship: (Lojac.) Soó

```json
{"parsed":true,"quality":1,"verbatim":"Chlorocyperus glaber form. fasciculariforme (Lojac.) Soó","normalized":"Chlorocyperus glaber f. fasciculariforme (Lojac.) Soó","canonical":{"stemmed":"Chlorocyperus glaber fasciculariform","simple":"Chlorocyperus glaber fasciculariforme","full":"Chlorocyperus glaber f. fasciculariforme"},"cardinality":3,"authorship":{"verbatim":"(Lojac.) Soó","normalized":"(Lojac.) Soó","authors":["Lojac.","Soó"],"originalAuth":{"authors":["Lojac."]},"combinationAuth":{"authors":["Soó"]}},"details":{"infraspecies":{"genus":"Chlorocyperus","species":"glaber","infraspecies":[{"value":"fasciculariforme","rank":"f.","authorship":{"verbatim":"(Lojac.) Soó","normalized":"(Lojac.) Soó","authors":["Lojac.","Soó"],"originalAuth":{"authors":["Lojac."]},"combinationAuth":{"authors":["Soó"]}}}]}},"words":[{"verbatim":"Chlorocyperus","normalized":"Chlorocyperus","wordType":"GENUS","start":0,"end":13},{"verbatim":"glaber","normalized":"glaber","wordType":"SPECIES","start":14,"end":20},{"verbatim":"form.","normalized":"f.","wordType":"RANK","start":21,"end":26},{"verbatim":"fasciculariforme","normalized":"fasciculariforme","wordType":"INFRASPECIES","start":27,"end":43},{"verbatim":"Lojac.","normalized":"Lojac.","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"Soó","normalized":"Soó","wordType":"AUTHOR_WORD","start":53,"end":56}],"id":"beee0dba-bef6-5550-954f-c978af09310a","parserVersion":"test_version"}
```

Name: Pteris longifolia fm. stipularis Linnaeus 1753

Canonical: Pteris longifolia f. stipularis

Authorship: Linnaeus 1753

```json
{"parsed":true,"quality":1,"verbatim":"Pteris longifolia fm. stipularis Linnaeus 1753","normalized":"Pteris longifolia f. stipularis Linnaeus 1753","canonical":{"stemmed":"Pteris longifol stipular","simple":"Pteris longifolia stipularis","full":"Pteris longifolia f. stipularis"},"cardinality":3,"authorship":{"verbatim":"Linnaeus 1753","normalized":"Linnaeus 1753","year":"1753","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1753"}}},"details":{"infraspecies":{"genus":"Pteris","species":"longifolia","infraspecies":[{"value":"stipularis","rank":"f.","authorship":{"verbatim":"Linnaeus 1753","normalized":"Linnaeus 1753","year":"1753","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1753"}}}}]}},"words":[{"verbatim":"Pteris","normalized":"Pteris","wordType":"GENUS","start":0,"end":6},{"verbatim":"longifolia","normalized":"longifolia","wordType":"SPECIES","start":7,"end":17},{"verbatim":"fm.","normalized":"f.","wordType":"RANK","start":18,"end":21},{"verbatim":"stipularis","normalized":"stipularis","wordType":"INFRASPECIES","start":22,"end":32},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":33,"end":41},{"verbatim":"1753","normalized":"1753","wordType":"YEAR","start":42,"end":46}],"id":"0460032a-131f-5a9b-9472-db8244752156","parserVersion":"test_version"}
```

Name: Pteris longifolia fm stipularis Linnaeus 1753

Canonical: Pteris longifolia f. stipularis

Authorship: Linnaeus 1753

```json
{"parsed":true,"quality":1,"verbatim":"Pteris longifolia fm stipularis Linnaeus 1753","normalized":"Pteris longifolia f. stipularis Linnaeus 1753","canonical":{"stemmed":"Pteris longifol stipular","simple":"Pteris longifolia stipularis","full":"Pteris longifolia f. stipularis"},"cardinality":3,"authorship":{"verbatim":"Linnaeus 1753","normalized":"Linnaeus 1753","year":"1753","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1753"}}},"details":{"infraspecies":{"genus":"Pteris","species":"longifolia","infraspecies":[{"value":"stipularis","rank":"f.","authorship":{"verbatim":"Linnaeus 1753","normalized":"Linnaeus 1753","year":"1753","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1753"}}}}]}},"words":[{"verbatim":"Pteris","normalized":"Pteris","wordType":"GENUS","start":0,"end":6},{"verbatim":"longifolia","normalized":"longifolia","wordType":"SPECIES","start":7,"end":17},{"verbatim":"fm","normalized":"f.","wordType":"RANK","start":18,"end":20},{"verbatim":"stipularis","normalized":"stipularis","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":32,"end":40},{"verbatim":"1753","normalized":"1753","wordType":"YEAR","start":41,"end":45}],"id":"b79aea4b-52d8-54fb-bca8-e54a779a1ce6","parserVersion":"test_version"}
```

Name: Sphaerotheca    fuliginea    f.     dahliae    Movss.     1967

Canonical: Sphaerotheca fuliginea f. dahliae

Authorship: Movss. 1967

```json
{"parsed":true,"quality":1,"verbatim":"Sphaerotheca    fuliginea    f.     dahliae    Movss.     1967","normalized":"Sphaerotheca fuliginea f. dahliae Movss. 1967","canonical":{"stemmed":"Sphaerotheca fuligine dahli","simple":"Sphaerotheca fuliginea dahliae","full":"Sphaerotheca fuliginea f. dahliae"},"cardinality":3,"authorship":{"verbatim":"Movss.     1967","normalized":"Movss. 1967","year":"1967","authors":["Movss."],"originalAuth":{"authors":["Movss."],"year":{"year":"1967"}}},"details":{"infraspecies":{"genus":"Sphaerotheca","species":"fuliginea","infraspecies":[{"value":"dahliae","rank":"f.","authorship":{"verbatim":"Movss.     1967","normalized":"Movss. 1967","year":"1967","authors":["Movss."],"originalAuth":{"authors":["Movss."],"year":{"year":"1967"}}}}]}},"words":[{"verbatim":"Sphaerotheca","normalized":"Sphaerotheca","wordType":"GENUS","start":0,"end":12},{"verbatim":"fuliginea","normalized":"fuliginea","wordType":"SPECIES","start":16,"end":25},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":29,"end":31},{"verbatim":"dahliae","normalized":"dahliae","wordType":"INFRASPECIES","start":36,"end":43},{"verbatim":"Movss.","normalized":"Movss.","wordType":"AUTHOR_WORD","start":47,"end":53},{"verbatim":"1967","normalized":"1967","wordType":"YEAR","start":58,"end":62}],"id":"bbd48fd4-ceee-5c66-ae42-f7fa43a8ea97","parserVersion":"test_version"}
```

Name: Allophylus amazonicus var amazonicus

Canonical: Allophylus amazonicus var. amazonicus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Allophylus amazonicus var amazonicus","normalized":"Allophylus amazonicus var. amazonicus","canonical":{"stemmed":"Allophylus amazonic amazonic","simple":"Allophylus amazonicus amazonicus","full":"Allophylus amazonicus var. amazonicus"},"cardinality":3,"details":{"infraspecies":{"genus":"Allophylus","species":"amazonicus","infraspecies":[{"value":"amazonicus","rank":"var."}]}},"words":[{"verbatim":"Allophylus","normalized":"Allophylus","wordType":"GENUS","start":0,"end":10},{"verbatim":"amazonicus","normalized":"amazonicus","wordType":"SPECIES","start":11,"end":21},{"verbatim":"var","normalized":"var.","wordType":"RANK","start":22,"end":25},{"verbatim":"amazonicus","normalized":"amazonicus","wordType":"INFRASPECIES","start":26,"end":36}],"id":"4e5c108c-b089-5198-9088-dd58d74d951f","parserVersion":"test_version"}
```

Name: Yarrowia lipolytica variety lipolytic

Canonical: Yarrowia lipolytica var. lipolytic

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Yarrowia lipolytica variety lipolytic","normalized":"Yarrowia lipolytica var. lipolytic","canonical":{"stemmed":"Yarrowia lipolytic lipolytic","simple":"Yarrowia lipolytica lipolytic","full":"Yarrowia lipolytica var. lipolytic"},"cardinality":3,"details":{"infraspecies":{"genus":"Yarrowia","species":"lipolytica","infraspecies":[{"value":"lipolytic","rank":"var."}]}},"words":[{"verbatim":"Yarrowia","normalized":"Yarrowia","wordType":"GENUS","start":0,"end":8},{"verbatim":"lipolytica","normalized":"lipolytica","wordType":"SPECIES","start":9,"end":19},{"verbatim":"variety","normalized":"var.","wordType":"RANK","start":20,"end":27},{"verbatim":"lipolytic","normalized":"lipolytic","wordType":"INFRASPECIES","start":28,"end":37}],"id":"5ecc8759-e1c3-5632-a863-7664625fc58d","parserVersion":"test_version"}
```

Name: Prunus armeniaca convar. budae (Pénzes) Soó

Canonical: Prunus armeniaca convar. budae

Authorship: (Pénzes) Soó

```json
{"parsed":true,"quality":1,"verbatim":"Prunus armeniaca convar. budae (Pénzes) Soó","normalized":"Prunus armeniaca convar. budae (Pénzes) Soó","canonical":{"stemmed":"Prunus armeniac bud","simple":"Prunus armeniaca budae","full":"Prunus armeniaca convar. budae"},"cardinality":3,"authorship":{"verbatim":"(Pénzes) Soó","normalized":"(Pénzes) Soó","authors":["Pénzes","Soó"],"originalAuth":{"authors":["Pénzes"]},"combinationAuth":{"authors":["Soó"]}},"details":{"infraspecies":{"genus":"Prunus","species":"armeniaca","infraspecies":[{"value":"budae","rank":"convar.","authorship":{"verbatim":"(Pénzes) Soó","normalized":"(Pénzes) Soó","authors":["Pénzes","Soó"],"originalAuth":{"authors":["Pénzes"]},"combinationAuth":{"authors":["Soó"]}}}]}},"words":[{"verbatim":"Prunus","normalized":"Prunus","wordType":"GENUS","start":0,"end":6},{"verbatim":"armeniaca","normalized":"armeniaca","wordType":"SPECIES","start":7,"end":16},{"verbatim":"convar.","normalized":"convar.","wordType":"RANK","start":17,"end":24},{"verbatim":"budae","normalized":"budae","wordType":"INFRASPECIES","start":25,"end":30},{"verbatim":"Pénzes","normalized":"Pénzes","wordType":"AUTHOR_WORD","start":32,"end":38},{"verbatim":"Soó","normalized":"Soó","wordType":"AUTHOR_WORD","start":40,"end":43}],"id":"c2133c2d-0486-54cb-a8cb-d355d458e19f","parserVersion":"test_version"}
```

Name: Polypodium pectinatum (L.) f. typica Rosenst.

Canonical: Polypodium pectinatum f. typica

Authorship: Rosenst.

```json
{"parsed":true,"quality":1,"verbatim":"Polypodium pectinatum (L.) f. typica Rosenst.","normalized":"Polypodium pectinatum (L.) f. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum f. typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraspecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"(L.)","normalized":"(L.)","authors":["L."],"originalAuth":{"authors":["L."]}},"infraspecies":[{"value":"typica","rank":"f.","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":27,"end":29},{"verbatim":"typica","normalized":"typica","wordType":"INFRASPECIES","start":30,"end":36},{"verbatim":"Rosenst.","normalized":"Rosenst.","wordType":"AUTHOR_WORD","start":37,"end":45}],"id":"b74dfd6b-c2d5-5e21-a807-f138667f0370","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f. typica Rosenst.

Canonical: Polypodium pectinatum f. typica

Authorship: Rosenst.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium pectinatum L. f. typica Rosenst.","normalized":"Polypodium pectinatum L. f. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum f. typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraspecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraspecies":[{"value":"typica","rank":"f.","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":25,"end":27},{"verbatim":"typica","normalized":"typica","wordType":"INFRASPECIES","start":28,"end":34},{"verbatim":"Rosenst.","normalized":"Rosenst.","wordType":"AUTHOR_WORD","start":35,"end":43}],"id":"68a2dccb-8b41-5a4f-92aa-06ae377b1503","parserVersion":"test_version"}
```

Name: Rubus fruticosus agamosp. chloocladus (W.C.R. Watson) A. & D. Löve

Canonical: Rubus fruticosus agamosp. chloocladus

Authorship: (W. C. R. Watson) A. & D. Löve

```json
{"parsed":true,"quality":1,"verbatim":"Rubus fruticosus agamosp. chloocladus (W.C.R. Watson) A. \u0026 D. Löve","normalized":"Rubus fruticosus agamosp. chloocladus (W. C. R. Watson) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos chlooclad","simple":"Rubus fruticosus chloocladus","full":"Rubus fruticosus agamosp. chloocladus"},"cardinality":3,"authorship":{"verbatim":"(W.C.R. Watson) A. \u0026 D. Löve","normalized":"(W. C. R. Watson) A. \u0026 D. Löve","authors":["W. C. R. Watson","A.","D. Löve"],"originalAuth":{"authors":["W. C. R. Watson"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraspecies":{"genus":"Rubus","species":"fruticosus","infraspecies":[{"value":"chloocladus","rank":"agamosp.","authorship":{"verbatim":"(W.C.R. Watson) A. \u0026 D. Löve","normalized":"(W. C. R. Watson) A. \u0026 D. Löve","authors":["W. C. R. Watson","A.","D. Löve"],"originalAuth":{"authors":["W. C. R. Watson"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"words":[{"verbatim":"Rubus","normalized":"Rubus","wordType":"GENUS","start":0,"end":5},{"verbatim":"fruticosus","normalized":"fruticosus","wordType":"SPECIES","start":6,"end":16},{"verbatim":"agamosp.","normalized":"agamosp.","wordType":"RANK","start":17,"end":25},{"verbatim":"chloocladus","normalized":"chloocladus","wordType":"INFRASPECIES","start":26,"end":37},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"Watson","normalized":"Watson","wordType":"AUTHOR_WORD","start":46,"end":52},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":54,"end":56},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":59,"end":61},{"verbatim":"Löve","normalized":"Löve","wordType":"AUTHOR_WORD","start":62,"end":66}],"id":"c6a80c28-12ab-550e-8255-3b96032ef98c","parserVersion":"test_version"}
```

Name: Rubus fruticosus L. agamossp. discolor (Weihe & Nees) A. & D. Löve

Canonical: Rubus fruticosus agamossp. discolor

Authorship: (Weihe & Nees) A. & D. Löve

```json
{"parsed":true,"quality":1,"verbatim":"Rubus fruticosus L. agamossp. discolor (Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"Rubus fruticosus L. agamossp. discolor (Weihe \u0026 Nees) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos discolor","simple":"Rubus fruticosus discolor","full":"Rubus fruticosus agamossp. discolor"},"cardinality":3,"authorship":{"verbatim":"(Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"(Weihe \u0026 Nees) A. \u0026 D. Löve","authors":["Weihe","Nees","A.","D. Löve"],"originalAuth":{"authors":["Weihe","Nees"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraspecies":{"genus":"Rubus","species":"fruticosus","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraspecies":[{"value":"discolor","rank":"agamossp.","authorship":{"verbatim":"(Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"(Weihe \u0026 Nees) A. \u0026 D. Löve","authors":["Weihe","Nees","A.","D. Löve"],"originalAuth":{"authors":["Weihe","Nees"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"words":[{"verbatim":"Rubus","normalized":"Rubus","wordType":"GENUS","start":0,"end":5},{"verbatim":"fruticosus","normalized":"fruticosus","wordType":"SPECIES","start":6,"end":16},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":17,"end":19},{"verbatim":"agamossp.","normalized":"agamossp.","wordType":"RANK","start":20,"end":29},{"verbatim":"discolor","normalized":"discolor","wordType":"INFRASPECIES","start":30,"end":38},{"verbatim":"Weihe","normalized":"Weihe","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"Nees","normalized":"Nees","wordType":"AUTHOR_WORD","start":48,"end":52},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":54,"end":56},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":59,"end":61},{"verbatim":"Löve","normalized":"Löve","wordType":"AUTHOR_WORD","start":62,"end":66}],"id":"a4265faa-5096-575b-914c-cd9cea4bbb7d","parserVersion":"test_version"}
```

Name: Rubus fruticosus agamovar. graecensis (W.Maurer) A. & D. Löve

Canonical: Rubus fruticosus agamovar. graecensis

Authorship: (W. Maurer) A. & D. Löve

```json
{"parsed":true,"quality":1,"verbatim":"Rubus fruticosus agamovar. graecensis (W.Maurer) A. \u0026 D. Löve","normalized":"Rubus fruticosus agamovar. graecensis (W. Maurer) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos graecens","simple":"Rubus fruticosus graecensis","full":"Rubus fruticosus agamovar. graecensis"},"cardinality":3,"authorship":{"verbatim":"(W.Maurer) A. \u0026 D. Löve","normalized":"(W. Maurer) A. \u0026 D. Löve","authors":["W. Maurer","A.","D. Löve"],"originalAuth":{"authors":["W. Maurer"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraspecies":{"genus":"Rubus","species":"fruticosus","infraspecies":[{"value":"graecensis","rank":"agamovar.","authorship":{"verbatim":"(W.Maurer) A. \u0026 D. Löve","normalized":"(W. Maurer) A. \u0026 D. Löve","authors":["W. Maurer","A.","D. Löve"],"originalAuth":{"authors":["W. Maurer"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"words":[{"verbatim":"Rubus","normalized":"Rubus","wordType":"GENUS","start":0,"end":5},{"verbatim":"fruticosus","normalized":"fruticosus","wordType":"SPECIES","start":6,"end":16},{"verbatim":"agamovar.","normalized":"agamovar.","wordType":"RANK","start":17,"end":26},{"verbatim":"graecensis","normalized":"graecensis","wordType":"INFRASPECIES","start":27,"end":37},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"Maurer","normalized":"Maurer","wordType":"AUTHOR_WORD","start":41,"end":47},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":54,"end":56},{"verbatim":"Löve","normalized":"Löve","wordType":"AUTHOR_WORD","start":57,"end":61}],"id":"9e3158af-63bd-5c94-91d1-f795342709d6","parserVersion":"test_version"}
```

<!-- TODO: the following phrasing can be ambiguous.
Does f mean forma or filius? Currently capturing it as filius
Following rule of thumb -- if f. is given without space, it is most
likely filius. If there is space, it is most likely forma -->
Name: Polypodium pectinatum L.f. typica Rosenst.

Canonical: Polypodium pectinatum typica

Authorship: Rosenst.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium pectinatum L.f. typica Rosenst.","normalized":"Polypodium pectinatum L. fil. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraspecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L.f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"infraspecies":[{"value":"typica","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":24,"end":26},{"verbatim":"typica","normalized":"typica","wordType":"INFRASPECIES","start":27,"end":33},{"verbatim":"Rosenst.","normalized":"Rosenst.","wordType":"AUTHOR_WORD","start":34,"end":42}],"id":"ea87b733-cae3-5a0f-a74d-3d921dcdbeb6","parserVersion":"test_version"}
```

Name: Polypodium lineare C.Chr. f. caudatoattenuatum Takeda

Canonical: Polypodium lineare f. caudatoattenuatum

Authorship: Takeda

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium lineare C.Chr. f. caudatoattenuatum Takeda","normalized":"Polypodium lineare C. Chr. f. caudatoattenuatum Takeda","canonical":{"stemmed":"Polypodium linear caudatoattenuat","simple":"Polypodium lineare caudatoattenuatum","full":"Polypodium lineare f. caudatoattenuatum"},"cardinality":3,"authorship":{"verbatim":"Takeda","normalized":"Takeda","authors":["Takeda"],"originalAuth":{"authors":["Takeda"]}},"details":{"infraspecies":{"genus":"Polypodium","species":"lineare","authorship":{"verbatim":"C.Chr.","normalized":"C. Chr.","authors":["C. Chr."],"originalAuth":{"authors":["C. Chr."]}},"infraspecies":[{"value":"caudatoattenuatum","rank":"f.","authorship":{"verbatim":"Takeda","normalized":"Takeda","authors":["Takeda"],"originalAuth":{"authors":["Takeda"]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"lineare","normalized":"lineare","wordType":"SPECIES","start":11,"end":18},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":19,"end":21},{"verbatim":"Chr.","normalized":"Chr.","wordType":"AUTHOR_WORD","start":21,"end":25},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":26,"end":28},{"verbatim":"caudatoattenuatum","normalized":"caudatoattenuatum","wordType":"INFRASPECIES","start":29,"end":46},{"verbatim":"Takeda","normalized":"Takeda","wordType":"AUTHOR_WORD","start":47,"end":53}],"id":"18cfd931-1ccd-5ea2-823a-71ba9604c783","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. f. albiflorum T.Yamaz.

Canonical: Rhododendron weyrichii f. albiflorum

Authorship: T. Yamaz.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Rhododendron weyrichii Maxim. f. albiflorum T.Yamaz.","normalized":"Rhododendron weyrichii Maxim. f. albiflorum T. Yamaz.","canonical":{"stemmed":"Rhododendron weyrich albiflor","simple":"Rhododendron weyrichii albiflorum","full":"Rhododendron weyrichii f. albiflorum"},"cardinality":3,"authorship":{"verbatim":"T.Yamaz.","normalized":"T. Yamaz.","authors":["T. Yamaz."],"originalAuth":{"authors":["T. Yamaz."]}},"details":{"infraspecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim.","normalized":"Maxim.","authors":["Maxim."],"originalAuth":{"authors":["Maxim."]}},"infraspecies":[{"value":"albiflorum","rank":"f.","authorship":{"verbatim":"T.Yamaz.","normalized":"T. Yamaz.","authors":["T. Yamaz."],"originalAuth":{"authors":["T. Yamaz."]}}}]}},"words":[{"verbatim":"Rhododendron","normalized":"Rhododendron","wordType":"GENUS","start":0,"end":12},{"verbatim":"weyrichii","normalized":"weyrichii","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Maxim.","normalized":"Maxim.","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":30,"end":32},{"verbatim":"albiflorum","normalized":"albiflorum","wordType":"INFRASPECIES","start":33,"end":43},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"Yamaz.","normalized":"Yamaz.","wordType":"AUTHOR_WORD","start":46,"end":52}],"id":"e515f1c8-3b95-5930-bcd1-09176727f0b7","parserVersion":"test_version"}
```

Name: Armeria maaritima (Mill.) Willd. fma. originaria Bern.

Canonical: Armeria maaritima f. originaria

Authorship: Bern.

```json
{"parsed":true,"quality":1,"verbatim":"Armeria maaritima (Mill.) Willd. fma. originaria Bern.","normalized":"Armeria maaritima (Mill.) Willd. f. originaria Bern.","canonical":{"stemmed":"Armeria maaritim originar","simple":"Armeria maaritima originaria","full":"Armeria maaritima f. originaria"},"cardinality":3,"authorship":{"verbatim":"Bern.","normalized":"Bern.","authors":["Bern."],"originalAuth":{"authors":["Bern."]}},"details":{"infraspecies":{"genus":"Armeria","species":"maaritima","authorship":{"verbatim":"(Mill.) Willd.","normalized":"(Mill.) Willd.","authors":["Mill.","Willd."],"originalAuth":{"authors":["Mill."]},"combinationAuth":{"authors":["Willd."]}},"infraspecies":[{"value":"originaria","rank":"f.","authorship":{"verbatim":"Bern.","normalized":"Bern.","authors":["Bern."],"originalAuth":{"authors":["Bern."]}}}]}},"words":[{"verbatim":"Armeria","normalized":"Armeria","wordType":"GENUS","start":0,"end":7},{"verbatim":"maaritima","normalized":"maaritima","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Mill.","normalized":"Mill.","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"Willd.","normalized":"Willd.","wordType":"AUTHOR_WORD","start":26,"end":32},{"verbatim":"fma.","normalized":"f.","wordType":"RANK","start":33,"end":37},{"verbatim":"originaria","normalized":"originaria","wordType":"INFRASPECIES","start":38,"end":48},{"verbatim":"Bern.","normalized":"Bern.","wordType":"AUTHOR_WORD","start":49,"end":54}],"id":"00d88bea-f076-5911-a450-fcfac1fe98bc","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. albiflorum T.Yamaz. f. fakeepithet

Canonical: Rhododendron weyrichii albiflorum f. fakeepithet

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Rhododendron weyrichii Maxim. albiflorum T.Yamaz. f. fakeepithet","normalized":"Rhododendron weyrichii Maxim. albiflorum T. Yamaz. f. fakeepithet","canonical":{"stemmed":"Rhododendron weyrich albiflor fakeepithet","simple":"Rhododendron weyrichii albiflorum fakeepithet","full":"Rhododendron weyrichii albiflorum f. fakeepithet"},"cardinality":4,"details":{"infraspecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim.","normalized":"Maxim.","authors":["Maxim."],"originalAuth":{"authors":["Maxim."]}},"infraspecies":[{"value":"albiflorum","authorship":{"verbatim":"T.Yamaz.","normalized":"T. Yamaz.","authors":["T. Yamaz."],"originalAuth":{"authors":["T. Yamaz."]}}},{"value":"fakeepithet","rank":"f."}]}},"words":[{"verbatim":"Rhododendron","normalized":"Rhododendron","wordType":"GENUS","start":0,"end":12},{"verbatim":"weyrichii","normalized":"weyrichii","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Maxim.","normalized":"Maxim.","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"albiflorum","normalized":"albiflorum","wordType":"INFRASPECIES","start":30,"end":40},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"Yamaz.","normalized":"Yamaz.","wordType":"AUTHOR_WORD","start":43,"end":49},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":50,"end":52},{"verbatim":"fakeepithet","normalized":"fakeepithet","wordType":"INFRASPECIES","start":53,"end":64}],"id":"ad0e299f-cd2c-52f3-9cab-49c70c5814f8","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. albiflorum (T.Yamaz. f.) fakeepithet

Canonical: Rhododendron weyrichii albiflorum fakeepithet

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Rhododendron weyrichii Maxim. albiflorum (T.Yamaz. f.) fakeepithet","normalized":"Rhododendron weyrichii Maxim. albiflorum (T. Yamaz. fil.) fakeepithet","canonical":{"stemmed":"Rhododendron weyrich albiflor fakeepithet","simple":"Rhododendron weyrichii albiflorum fakeepithet","full":"Rhododendron weyrichii albiflorum fakeepithet"},"cardinality":4,"details":{"infraspecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim.","normalized":"Maxim.","authors":["Maxim."],"originalAuth":{"authors":["Maxim."]}},"infraspecies":[{"value":"albiflorum","authorship":{"verbatim":"(T.Yamaz. f.)","normalized":"(T. Yamaz. fil.)","authors":["T. Yamaz. fil."],"originalAuth":{"authors":["T. Yamaz. fil."]}}},{"value":"fakeepithet"}]}},"words":[{"verbatim":"Rhododendron","normalized":"Rhododendron","wordType":"GENUS","start":0,"end":12},{"verbatim":"weyrichii","normalized":"weyrichii","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Maxim.","normalized":"Maxim.","wordType":"AUTHOR_WORD","start":23,"end":29},{"verbatim":"albiflorum","normalized":"albiflorum","wordType":"INFRASPECIES","start":30,"end":40},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":42,"end":44},{"verbatim":"Yamaz.","normalized":"Yamaz.","wordType":"AUTHOR_WORD","start":44,"end":50},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":51,"end":53},{"verbatim":"fakeepithet","normalized":"fakeepithet","wordType":"INFRASPECIES","start":55,"end":66}],"id":"2a7d1bab-b208-5654-9406-f7afc696b00b","parserVersion":"test_version"}
```

Name: Cotoneaster (Pyracantha) rogersiana var.aurantiaca

Canonical: Cotoneaster rogersiana var. aurantiaca

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Cotoneaster (Pyracantha) rogersiana var.aurantiaca","normalized":"Cotoneaster (Pyracantha) rogersiana var. aurantiaca","canonical":{"stemmed":"Cotoneaster rogersian aurantiac","simple":"Cotoneaster rogersiana aurantiaca","full":"Cotoneaster rogersiana var. aurantiaca"},"cardinality":3,"details":{"infraspecies":{"genus":"Cotoneaster","subgenus":"Pyracantha","species":"rogersiana","infraspecies":[{"value":"aurantiaca","rank":"var."}]}},"words":[{"verbatim":"Cotoneaster","normalized":"Cotoneaster","wordType":"GENUS","start":0,"end":11},{"verbatim":"Pyracantha","normalized":"Pyracantha","wordType":"INFRA_GENUS","start":13,"end":23},{"verbatim":"rogersiana","normalized":"rogersiana","wordType":"SPECIES","start":25,"end":35},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":36,"end":40},{"verbatim":"aurantiaca","normalized":"aurantiaca","wordType":"INFRASPECIES","start":40,"end":50}],"id":"86716b35-27ce-5d21-ab18-e8bb0c5d80be","parserVersion":"test_version"}
```

Name: Poa annua fo varia

Canonical: Poa annua f. varia

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Poa annua fo varia","normalized":"Poa annua f. varia","canonical":{"stemmed":"Poa annu uar","simple":"Poa annua varia","full":"Poa annua f. varia"},"cardinality":3,"details":{"infraspecies":{"genus":"Poa","species":"annua","infraspecies":[{"value":"varia","rank":"f."}]}},"words":[{"verbatim":"Poa","normalized":"Poa","wordType":"GENUS","start":0,"end":3},{"verbatim":"annua","normalized":"annua","wordType":"SPECIES","start":4,"end":9},{"verbatim":"fo","normalized":"f.","wordType":"RANK","start":10,"end":12},{"verbatim":"varia","normalized":"varia","wordType":"INFRASPECIES","start":13,"end":18}],"id":"32838647-3c46-509b-a81b-62d24940845f","parserVersion":"test_version"}
```

Name: Physarum globuliferum forma. flavum Leontyev & Dudka

Canonical: Physarum globuliferum f. flavum

Authorship: Leontyev & Dudka

```json
{"parsed":true,"quality":1,"verbatim":"Physarum globuliferum forma. flavum Leontyev \u0026 Dudka","normalized":"Physarum globuliferum f. flavum Leontyev \u0026 Dudka","canonical":{"stemmed":"Physarum globulifer flau","simple":"Physarum globuliferum flavum","full":"Physarum globuliferum f. flavum"},"cardinality":3,"authorship":{"verbatim":"Leontyev \u0026 Dudka","normalized":"Leontyev \u0026 Dudka","authors":["Leontyev","Dudka"],"originalAuth":{"authors":["Leontyev","Dudka"]}},"details":{"infraspecies":{"genus":"Physarum","species":"globuliferum","infraspecies":[{"value":"flavum","rank":"f.","authorship":{"verbatim":"Leontyev \u0026 Dudka","normalized":"Leontyev \u0026 Dudka","authors":["Leontyev","Dudka"],"originalAuth":{"authors":["Leontyev","Dudka"]}}}]}},"words":[{"verbatim":"Physarum","normalized":"Physarum","wordType":"GENUS","start":0,"end":8},{"verbatim":"globuliferum","normalized":"globuliferum","wordType":"SPECIES","start":9,"end":21},{"verbatim":"forma.","normalized":"f.","wordType":"RANK","start":22,"end":28},{"verbatim":"flavum","normalized":"flavum","wordType":"INFRASPECIES","start":29,"end":35},{"verbatim":"Leontyev","normalized":"Leontyev","wordType":"AUTHOR_WORD","start":36,"end":44},{"verbatim":"Dudka","normalized":"Dudka","wordType":"AUTHOR_WORD","start":47,"end":52}],"id":"bbcecb18-4484-528b-a8b9-93e1634d31b5","parserVersion":"test_version"}
```

Name: Homalanthus nutans (Mull.Arg.) Benth. & Hook. f. ex Drake

Canonical: Homalanthus nutans

Authorship: (Mull. Arg.) Benth. & Hook. fil. ex Drake

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Homalanthus nutans (Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"Homalanthus nutans (Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","canonical":{"stemmed":"Homalanthus nutans","simple":"Homalanthus nutans","full":"Homalanthus nutans"},"cardinality":2,"authorship":{"verbatim":"(Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"(Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","authors":["Mull. Arg.","Benth.","Hook. fil.","Drake"],"originalAuth":{"authors":["Mull. Arg."]},"combinationAuth":{"authors":["Benth.","Hook. fil."],"exAuthors":{"authors":["Drake"]}}},"details":{"species":{"genus":"Homalanthus","species":"nutans","authorship":{"verbatim":"(Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"(Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","authors":["Mull. Arg.","Benth.","Hook. fil.","Drake"],"originalAuth":{"authors":["Mull. Arg."]},"combinationAuth":{"authors":["Benth.","Hook. fil."],"exAuthors":{"authors":["Drake"]}}}}},"words":[{"verbatim":"Homalanthus","normalized":"Homalanthus","wordType":"GENUS","start":0,"end":11},{"verbatim":"nutans","normalized":"nutans","wordType":"SPECIES","start":12,"end":18},{"verbatim":"Mull.","normalized":"Mull.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"Arg.","normalized":"Arg.","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"Benth.","normalized":"Benth.","wordType":"AUTHOR_WORD","start":31,"end":37},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":46,"end":48},{"verbatim":"Drake","normalized":"Drake","wordType":"AUTHOR_WORD","start":52,"end":57}],"id":"83c06d35-e323-5750-84fb-f8c184fd1ee4","parserVersion":"test_version"}
```

Name: Calicium furfuraceum * furfuraceum (L.) Pers. 1797

Canonical: Calicium furfuraceum * furfuraceum

Authorship: (L.) Pers. 1797

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Calicium furfuraceum * furfuraceum (L.) Pers. 1797","normalized":"Calicium furfuraceum * furfuraceum (L.) Pers. 1797","canonical":{"stemmed":"Calicium furfurace furfurace","simple":"Calicium furfuraceum furfuraceum","full":"Calicium furfuraceum * furfuraceum"},"cardinality":3,"authorship":{"verbatim":"(L.) Pers. 1797","normalized":"(L.) Pers. 1797","authors":["L.","Pers."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Pers."],"year":{"year":"1797"}}},"details":{"infraspecies":{"genus":"Calicium","species":"furfuraceum","infraspecies":[{"value":"furfuraceum","rank":"*","authorship":{"verbatim":"(L.) Pers. 1797","normalized":"(L.) Pers. 1797","authors":["L.","Pers."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Pers."],"year":{"year":"1797"}}}}]}},"words":[{"verbatim":"Calicium","normalized":"Calicium","wordType":"GENUS","start":0,"end":8},{"verbatim":"furfuraceum","normalized":"furfuraceum","wordType":"SPECIES","start":9,"end":20},{"verbatim":"*","normalized":"*","wordType":"RANK","start":21,"end":22},{"verbatim":"furfuraceum","normalized":"furfuraceum","wordType":"INFRASPECIES","start":23,"end":34},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"Pers.","normalized":"Pers.","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"1797","normalized":"1797","wordType":"YEAR","start":46,"end":50}],"id":"6c5da8ae-cc50-5ce3-835d-d42e16aa0757","parserVersion":"test_version"}
```

Name: Polyrhachis orsyllus nat musculus Forel 1901

Canonical: Polyrhachis orsyllus nat musculus

Authorship: Forel 1901

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Polyrhachis orsyllus nat musculus Forel 1901","normalized":"Polyrhachis orsyllus nat musculus Forel 1901","canonical":{"stemmed":"Polyrhachis orsyll muscul","simple":"Polyrhachis orsyllus musculus","full":"Polyrhachis orsyllus nat musculus"},"cardinality":3,"authorship":{"verbatim":"Forel 1901","normalized":"Forel 1901","year":"1901","authors":["Forel"],"originalAuth":{"authors":["Forel"],"year":{"year":"1901"}}},"details":{"infraspecies":{"genus":"Polyrhachis","species":"orsyllus","infraspecies":[{"value":"musculus","rank":"nat","authorship":{"verbatim":"Forel 1901","normalized":"Forel 1901","year":"1901","authors":["Forel"],"originalAuth":{"authors":["Forel"],"year":{"year":"1901"}}}}]}},"words":[{"verbatim":"Polyrhachis","normalized":"Polyrhachis","wordType":"GENUS","start":0,"end":11},{"verbatim":"orsyllus","normalized":"orsyllus","wordType":"SPECIES","start":12,"end":20},{"verbatim":"nat","normalized":"nat","wordType":"RANK","start":21,"end":24},{"verbatim":"musculus","normalized":"musculus","wordType":"INFRASPECIES","start":25,"end":33},{"verbatim":"Forel","normalized":"Forel","wordType":"AUTHOR_WORD","start":34,"end":39},{"verbatim":"1901","normalized":"1901","wordType":"YEAR","start":40,"end":44}],"id":"3392132e-3dba-5b7e-a7c9-e4a68954c8b2","parserVersion":"test_version"}
```

Name: Acidalia remutaria ab. n. undularia

Canonical: Acidalia remutaria ab. n. undularia

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Acidalia remutaria ab. n. undularia","normalized":"Acidalia remutaria ab. n. undularia","canonical":{"stemmed":"Acidalia remutar undular","simple":"Acidalia remutaria undularia","full":"Acidalia remutaria ab. n. undularia"},"cardinality":3,"details":{"infraspecies":{"genus":"Acidalia","species":"remutaria","infraspecies":[{"value":"undularia","rank":"ab. n."}]}},"words":[{"verbatim":"Acidalia","normalized":"Acidalia","wordType":"GENUS","start":0,"end":8},{"verbatim":"remutaria","normalized":"remutaria","wordType":"SPECIES","start":9,"end":18},{"verbatim":"ab. n.","normalized":"ab. n.","wordType":"RANK","start":19,"end":25},{"verbatim":"undularia","normalized":"undularia","wordType":"INFRASPECIES","start":26,"end":35}],"id":"ac834e3e-b861-5fbf-9cf9-197ad3effb99","parserVersion":"test_version"}
```

Name: Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius, 1912

Canonical: Acmaeops bivittata ab. fusciceps

Authorship: Aurivillius 1912

```json
{"parsed":true,"quality":1,"verbatim":"Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius, 1912","normalized":"Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius 1912","canonical":{"stemmed":"Acmaeops biuittat fusciceps","simple":"Acmaeops bivittata fusciceps","full":"Acmaeops bivittata ab. fusciceps"},"cardinality":3,"authorship":{"verbatim":"Aurivillius, 1912","normalized":"Aurivillius 1912","year":"1912","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1912"}}},"details":{"infraspecies":{"genus":"Acmaeops","subgenus":"Pseudodinoptera","species":"bivittata","infraspecies":[{"value":"fusciceps","rank":"ab.","authorship":{"verbatim":"Aurivillius, 1912","normalized":"Aurivillius 1912","year":"1912","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1912"}}}}]}},"words":[{"verbatim":"Acmaeops","normalized":"Acmaeops","wordType":"GENUS","start":0,"end":8},{"verbatim":"Pseudodinoptera","normalized":"Pseudodinoptera","wordType":"INFRA_GENUS","start":10,"end":25},{"verbatim":"bivittata","normalized":"bivittata","wordType":"SPECIES","start":27,"end":36},{"verbatim":"ab.","normalized":"ab.","wordType":"RANK","start":37,"end":40},{"verbatim":"fusciceps","normalized":"fusciceps","wordType":"INFRASPECIES","start":41,"end":50},{"verbatim":"Aurivillius","normalized":"Aurivillius","wordType":"AUTHOR_WORD","start":51,"end":62},{"verbatim":"1912","normalized":"1912","wordType":"YEAR","start":64,"end":68}],"id":"3f3dfc38-f660-56d6-a4f8-568f84a6878a","parserVersion":"test_version"}
```

### Infraspecies multiple (ICN)

Name: Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum var. zonatum f. parvum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"quality":1,"verbatim":"Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat paru","simple":"Hydnellum scrobiculatum zonatum parvum","full":"Hydnellum scrobiculatum var. zonatum f. parvum"},"cardinality":4,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraspecies":{"genus":"Hydnellum","species":"scrobiculatum","infraspecies":[{"value":"zonatum","rank":"var."},{"value":"parvum","rank":"f.","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"words":[{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":0,"end":9},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":10,"end":23},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":24,"end":28},{"verbatim":"zonatum","normalized":"zonatum","wordType":"INFRASPECIES","start":29,"end":36},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":37,"end":39},{"verbatim":"parvum","normalized":"parvum","wordType":"INFRASPECIES","start":40,"end":46},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":48,"end":54},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":56,"end":58},{"verbatim":"Hall","normalized":"Hall","wordType":"AUTHOR_WORD","start":59,"end":63},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":66,"end":68},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":68,"end":70},{"verbatim":"Stuntz","normalized":"Stuntz","wordType":"AUTHOR_WORD","start":71,"end":77},{"verbatim":"1972","normalized":"1972","wordType":"YEAR","start":78,"end":82}],"id":"805654ed-0115-5f3e-af92-5808f215afbf","parserVersion":"test_version"}
```

Name: Senecio fuchsii C.C.Gmel. subsp. fuchsii var. expansus (Boiss. & Heldr.) Hayek

Canonical: Senecio fuchsii subsp. fuchsii var. expansus

Authorship: (Boiss. & Heldr.) Hayek

```json
{"parsed":true,"quality":1,"verbatim":"Senecio fuchsii C.C.Gmel. subsp. fuchsii var. expansus (Boiss. \u0026 Heldr.) Hayek","normalized":"Senecio fuchsii C. C. Gmel. subsp. fuchsii var. expansus (Boiss. \u0026 Heldr.) Hayek","canonical":{"stemmed":"Senecio fuchs fuchs expans","simple":"Senecio fuchsii fuchsii expansus","full":"Senecio fuchsii subsp. fuchsii var. expansus"},"cardinality":4,"authorship":{"verbatim":"(Boiss. \u0026 Heldr.) Hayek","normalized":"(Boiss. \u0026 Heldr.) Hayek","authors":["Boiss.","Heldr.","Hayek"],"originalAuth":{"authors":["Boiss.","Heldr."]},"combinationAuth":{"authors":["Hayek"]}},"details":{"infraspecies":{"genus":"Senecio","species":"fuchsii","authorship":{"verbatim":"C.C.Gmel.","normalized":"C. C. Gmel.","authors":["C. C. Gmel."],"originalAuth":{"authors":["C. C. Gmel."]}},"infraspecies":[{"value":"fuchsii","rank":"subsp."},{"value":"expansus","rank":"var.","authorship":{"verbatim":"(Boiss. \u0026 Heldr.) Hayek","normalized":"(Boiss. \u0026 Heldr.) Hayek","authors":["Boiss.","Heldr.","Hayek"],"originalAuth":{"authors":["Boiss.","Heldr."]},"combinationAuth":{"authors":["Hayek"]}}}]}},"words":[{"verbatim":"Senecio","normalized":"Senecio","wordType":"GENUS","start":0,"end":7},{"verbatim":"fuchsii","normalized":"fuchsii","wordType":"SPECIES","start":8,"end":15},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":18,"end":20},{"verbatim":"Gmel.","normalized":"Gmel.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":26,"end":32},{"verbatim":"fuchsii","normalized":"fuchsii","wordType":"INFRASPECIES","start":33,"end":40},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":41,"end":45},{"verbatim":"expansus","normalized":"expansus","wordType":"INFRASPECIES","start":46,"end":54},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":56,"end":62},{"verbatim":"Heldr.","normalized":"Heldr.","wordType":"AUTHOR_WORD","start":65,"end":71},{"verbatim":"Hayek","normalized":"Hayek","wordType":"AUTHOR_WORD","start":73,"end":78}],"id":"93ed1df3-5016-56e7-8aa8-3a01df49a11a","parserVersion":"test_version"}
```

Name: Senecio fuchsii C.C.Gmel. subsp. fuchsii var. fuchsii

Canonical: Senecio fuchsii subsp. fuchsii var. fuchsii

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Senecio fuchsii C.C.Gmel. subsp. fuchsii var. fuchsii","normalized":"Senecio fuchsii C. C. Gmel. subsp. fuchsii var. fuchsii","canonical":{"stemmed":"Senecio fuchs fuchs fuchs","simple":"Senecio fuchsii fuchsii fuchsii","full":"Senecio fuchsii subsp. fuchsii var. fuchsii"},"cardinality":4,"details":{"infraspecies":{"genus":"Senecio","species":"fuchsii","authorship":{"verbatim":"C.C.Gmel.","normalized":"C. C. Gmel.","authors":["C. C. Gmel."],"originalAuth":{"authors":["C. C. Gmel."]}},"infraspecies":[{"value":"fuchsii","rank":"subsp."},{"value":"fuchsii","rank":"var."}]}},"words":[{"verbatim":"Senecio","normalized":"Senecio","wordType":"GENUS","start":0,"end":7},{"verbatim":"fuchsii","normalized":"fuchsii","wordType":"SPECIES","start":8,"end":15},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":18,"end":20},{"verbatim":"Gmel.","normalized":"Gmel.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":26,"end":32},{"verbatim":"fuchsii","normalized":"fuchsii","wordType":"INFRASPECIES","start":33,"end":40},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":41,"end":45},{"verbatim":"fuchsii","normalized":"fuchsii","wordType":"INFRASPECIES","start":46,"end":53}],"id":"481c3fc6-6f0c-55fa-b119-64d78d0bde03","parserVersion":"test_version"}
```

Name: Euastrum divergens var. rhodesiense f. coronulum A.M. Scott & Prescott

Canonical: Euastrum divergens var. rhodesiense f. coronulum

Authorship: A. M. Scott & Prescott

```json
{"parsed":true,"quality":1,"verbatim":"Euastrum divergens var. rhodesiense f. coronulum A.M. Scott \u0026 Prescott","normalized":"Euastrum divergens var. rhodesiense f. coronulum A. M. Scott \u0026 Prescott","canonical":{"stemmed":"Euastrum diuergens rhodesiens coronul","simple":"Euastrum divergens rhodesiense coronulum","full":"Euastrum divergens var. rhodesiense f. coronulum"},"cardinality":4,"authorship":{"verbatim":"A.M. Scott \u0026 Prescott","normalized":"A. M. Scott \u0026 Prescott","authors":["A. M. Scott","Prescott"],"originalAuth":{"authors":["A. M. Scott","Prescott"]}},"details":{"infraspecies":{"genus":"Euastrum","species":"divergens","infraspecies":[{"value":"rhodesiense","rank":"var."},{"value":"coronulum","rank":"f.","authorship":{"verbatim":"A.M. Scott \u0026 Prescott","normalized":"A. M. Scott \u0026 Prescott","authors":["A. M. Scott","Prescott"],"originalAuth":{"authors":["A. M. Scott","Prescott"]}}}]}},"words":[{"verbatim":"Euastrum","normalized":"Euastrum","wordType":"GENUS","start":0,"end":8},{"verbatim":"divergens","normalized":"divergens","wordType":"SPECIES","start":9,"end":18},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":19,"end":23},{"verbatim":"rhodesiense","normalized":"rhodesiense","wordType":"INFRASPECIES","start":24,"end":35},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":36,"end":38},{"verbatim":"coronulum","normalized":"coronulum","wordType":"INFRASPECIES","start":39,"end":48},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":49,"end":51},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"Scott","normalized":"Scott","wordType":"AUTHOR_WORD","start":54,"end":59},{"verbatim":"Prescott","normalized":"Prescott","wordType":"AUTHOR_WORD","start":62,"end":70}],"id":"3e5a8eed-9f34-5f2b-95b5-1a45740e4306","parserVersion":"test_version"}
```

### Infraspecies with greek letters (ICN)

Name: Aristotelia fruticosa var. δ. microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var. δ. microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraspecies":{"genus":"Aristotelia","species":"fruticosa","infraspecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"words":[{"verbatim":"Aristotelia","normalized":"Aristotelia","wordType":"GENUS","start":0,"end":11},{"verbatim":"fruticosa","normalized":"fruticosa","wordType":"SPECIES","start":12,"end":21},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":22,"end":26},{"verbatim":"microphylla","normalized":"microphylla","wordType":"INFRASPECIES","start":30,"end":41},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":42,"end":47},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":47,"end":49}],"id":"34378b1d-27ef-5a38-a3ad-b2da249bc9d4","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var. δ microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var. δ microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraspecies":{"genus":"Aristotelia","species":"fruticosa","infraspecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"words":[{"verbatim":"Aristotelia","normalized":"Aristotelia","wordType":"GENUS","start":0,"end":11},{"verbatim":"fruticosa","normalized":"fruticosa","wordType":"SPECIES","start":12,"end":21},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":22,"end":26},{"verbatim":"microphylla","normalized":"microphylla","wordType":"INFRASPECIES","start":29,"end":40},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":41,"end":46},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":46,"end":48}],"id":"d31a653a-8686-5bf4-b657-6164f494e6b4","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var.δ.microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var.δ.microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraspecies":{"genus":"Aristotelia","species":"fruticosa","infraspecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"words":[{"verbatim":"Aristotelia","normalized":"Aristotelia","wordType":"GENUS","start":0,"end":11},{"verbatim":"fruticosa","normalized":"fruticosa","wordType":"SPECIES","start":12,"end":21},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":22,"end":26},{"verbatim":"microphylla","normalized":"microphylla","wordType":"INFRASPECIES","start":28,"end":39},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":45,"end":47}],"id":"c2f051e5-c1a2-52f8-a02f-70510030faa1","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var. δmicrophylla Hook.f.

Canonical: Aristotelia fruticosa

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aristotelia fruticosa var. δmicrophylla Hook.f.","normalized":"Aristotelia fruticosa","canonical":{"stemmed":"Aristotelia fruticos","simple":"Aristotelia fruticosa","full":"Aristotelia fruticosa"},"cardinality":2,"tail":" var. δmicrophylla Hook.f.","details":{"species":{"genus":"Aristotelia","species":"fruticosa"}},"words":[{"verbatim":"Aristotelia","normalized":"Aristotelia","wordType":"GENUS","start":0,"end":11},{"verbatim":"fruticosa","normalized":"fruticosa","wordType":"SPECIES","start":12,"end":21}],"id":"f7749c21-82a6-5c42-ab58-7b3d5a824e96","parserVersion":"test_version"}
```

### Names with the dagger char '†'

Name: Henriksenopterix†

Canonical: Henriksenopterix

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Henriksenopterix†","normalized":"Henriksenopterix","canonical":{"stemmed":"Henriksenopterix","simple":"Henriksenopterix","full":"Henriksenopterix"},"cardinality":1,"daggerChar":true,"details":{"uninomial":{"uninomial":"Henriksenopterix"}},"words":[{"verbatim":"Henriksenopterix","normalized":"Henriksenopterix","wordType":"UNINOMIAL","start":0,"end":16}],"id":"3cf4f556-ddb9-5a65-ab2f-531d387303eb","parserVersion":"test_version"}
```

Name: Henriksenopterix† paucistriata (Henriksen, 1922)

Canonical: Henriksenopterix paucistriata

Authorship: (Henriksen 1922)

```json
{"parsed":true,"quality":1,"verbatim":"Henriksenopterix† paucistriata (Henriksen, 1922)","normalized":"Henriksenopterix paucistriata (Henriksen 1922)","canonical":{"stemmed":"Henriksenopterix paucistriat","simple":"Henriksenopterix paucistriata","full":"Henriksenopterix paucistriata"},"cardinality":2,"authorship":{"verbatim":"(Henriksen, 1922)","normalized":"(Henriksen 1922)","year":"1922","authors":["Henriksen"],"originalAuth":{"authors":["Henriksen"],"year":{"year":"1922"}}},"daggerChar":true,"details":{"species":{"genus":"Henriksenopterix","species":"paucistriata","authorship":{"verbatim":"(Henriksen, 1922)","normalized":"(Henriksen 1922)","year":"1922","authors":["Henriksen"],"originalAuth":{"authors":["Henriksen"],"year":{"year":"1922"}}}}},"words":[{"verbatim":"Henriksenopterix","normalized":"Henriksenopterix","wordType":"GENUS","start":0,"end":16},{"verbatim":"paucistriata","normalized":"paucistriata","wordType":"SPECIES","start":20,"end":32},{"verbatim":"Henriksen","normalized":"Henriksen","wordType":"AUTHOR_WORD","start":34,"end":43},{"verbatim":"1922","normalized":"1922","wordType":"YEAR","start":45,"end":49}],"id":"510f327c-ee88-50fc-a5f7-94df7d05aa90","parserVersion":"test_version"}
```

Name: Heteralocha acutirostris (Gould, 1837) Huia N E†

Canonical: Heteralocha acutirostris

Authorship: (Gould 1837) Huia N E

```json
{"parsed":true,"quality":1,"verbatim":"Heteralocha acutirostris (Gould, 1837) Huia N E†","normalized":"Heteralocha acutirostris (Gould 1837) Huia N E","canonical":{"stemmed":"Heteralocha acutirostr","simple":"Heteralocha acutirostris","full":"Heteralocha acutirostris"},"cardinality":2,"authorship":{"verbatim":"(Gould, 1837) Huia N E","normalized":"(Gould 1837) Huia N E","year":"1837","authors":["Gould","Huia N E"],"originalAuth":{"authors":["Gould"],"year":{"year":"1837"}},"combinationAuth":{"authors":["Huia N E"]}},"daggerChar":true,"details":{"species":{"genus":"Heteralocha","species":"acutirostris","authorship":{"verbatim":"(Gould, 1837) Huia N E","normalized":"(Gould 1837) Huia N E","year":"1837","authors":["Gould","Huia N E"],"originalAuth":{"authors":["Gould"],"year":{"year":"1837"}},"combinationAuth":{"authors":["Huia N E"]}}}},"words":[{"verbatim":"Heteralocha","normalized":"Heteralocha","wordType":"GENUS","start":0,"end":11},{"verbatim":"acutirostris","normalized":"acutirostris","wordType":"SPECIES","start":12,"end":24},{"verbatim":"Gould","normalized":"Gould","wordType":"AUTHOR_WORD","start":26,"end":31},{"verbatim":"1837","normalized":"1837","wordType":"YEAR","start":33,"end":37},{"verbatim":"Huia","normalized":"Huia","wordType":"AUTHOR_WORD","start":39,"end":43},{"verbatim":"N","normalized":"N","wordType":"AUTHOR_WORD","start":44,"end":45},{"verbatim":"E","normalized":"E","wordType":"AUTHOR_WORD","start":46,"end":47}],"id":"197728f8-091b-5378-a505-c73acd6cbefc","parserVersion":"test_version"}
```

<!-- TODO: tail contains 3 empty spaces instead of a dagger -->
Name: Oncorhynchus nerka (Walbaum, 1792) Sockeye salmon F A †?

Canonical: Oncorhynchus nerka salmon

Authorship: F A

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Oncorhynchus nerka (Walbaum, 1792) Sockeye salmon F A †?","normalized":"Oncorhynchus nerka (Walbaum 1792) Sockeye salmon F A","canonical":{"stemmed":"Oncorhynchus nerk salmon","simple":"Oncorhynchus nerka salmon","full":"Oncorhynchus nerka salmon"},"cardinality":3,"authorship":{"verbatim":"F A","normalized":"F A","authors":["F A"],"originalAuth":{"authors":["F A"]}},"daggerChar":true,"tail":"    ?","details":{"infraspecies":{"genus":"Oncorhynchus","species":"nerka","authorship":{"verbatim":"(Walbaum, 1792) Sockeye","normalized":"(Walbaum 1792) Sockeye","year":"1792","authors":["Walbaum","Sockeye"],"originalAuth":{"authors":["Walbaum"],"year":{"year":"1792"}},"combinationAuth":{"authors":["Sockeye"]}},"infraspecies":[{"value":"salmon","authorship":{"verbatim":"F A","normalized":"F A","authors":["F A"],"originalAuth":{"authors":["F A"]}}}]}},"words":[{"verbatim":"Oncorhynchus","normalized":"Oncorhynchus","wordType":"GENUS","start":0,"end":12},{"verbatim":"nerka","normalized":"nerka","wordType":"SPECIES","start":13,"end":18},{"verbatim":"Walbaum","normalized":"Walbaum","wordType":"AUTHOR_WORD","start":20,"end":27},{"verbatim":"1792","normalized":"1792","wordType":"YEAR","start":29,"end":33},{"verbatim":"Sockeye","normalized":"Sockeye","wordType":"AUTHOR_WORD","start":35,"end":42},{"verbatim":"salmon","normalized":"salmon","wordType":"INFRASPECIES","start":43,"end":49},{"verbatim":"F","normalized":"F","wordType":"AUTHOR_WORD","start":50,"end":51},{"verbatim":"A","normalized":"A","wordType":"AUTHOR_WORD","start":52,"end":53}],"id":"fa50e193-9745-5355-acb9-3c5c2179a3d6","parserVersion":"test_version"}
```

### Hybrids with notho- ranks

Name: Crataegus curvisepala nvar. naviculiformis T. Petauer

Canonical: Crataegus curvisepala nvar. naviculiformis

Authorship: T. Petauer

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Crataegus curvisepala nvar. naviculiformis T. Petauer","normalized":"Crataegus curvisepala nvar. naviculiformis T. Petauer","canonical":{"stemmed":"Crataegus curuisepal nauiculiform","simple":"Crataegus curvisepala naviculiformis","full":"Crataegus curvisepala nvar. naviculiformis"},"cardinality":3,"authorship":{"verbatim":"T. Petauer","normalized":"T. Petauer","authors":["T. Petauer"],"originalAuth":{"authors":["T. Petauer"]}},"hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Crataegus","species":"curvisepala","infraspecies":[{"value":"naviculiformis","rank":"nvar.","authorship":{"verbatim":"T. Petauer","normalized":"T. Petauer","authors":["T. Petauer"],"originalAuth":{"authors":["T. Petauer"]}}}]}},"words":[{"verbatim":"Crataegus","normalized":"Crataegus","wordType":"GENUS","start":0,"end":9},{"verbatim":"curvisepala","normalized":"curvisepala","wordType":"SPECIES","start":10,"end":21},{"verbatim":"nvar.","normalized":"nvar.","wordType":"RANK","start":22,"end":27},{"verbatim":"naviculiformis","normalized":"naviculiformis","wordType":"INFRASPECIES","start":28,"end":42},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"Petauer","normalized":"Petauer","wordType":"AUTHOR_WORD","start":46,"end":53}],"id":"f3e2ccac-4844-57a7-8903-4e3b6a0d0851","parserVersion":"test_version"}
```

Name: Aconitum W. Mucher nothosect. Acopellus

Canonical: Aconitum nothosect. Acopellus

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum W. Mucher nothosect. Acopellus","normalized":"Aconitum nothosect. Acopellus","canonical":{"stemmed":"Acopellus","simple":"Acopellus","full":"Aconitum nothosect. Acopellus"},"cardinality":1,"hybrid":"NOTHO_HYBRID","details":{"uninomial":{"uninomial":"Acopellus","rank":"nothosect.","parent":"Aconitum"}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":9,"end":11},{"verbatim":"Mucher","normalized":"Mucher","wordType":"AUTHOR_WORD","start":12,"end":18},{"verbatim":"nothosect.","normalized":"nothosect.","wordType":"RANK","start":19,"end":29},{"verbatim":"Acopellus","normalized":"Acopellus","wordType":"UNINOMIAL","start":30,"end":39}],"id":"815f38e4-2425-551d-b054-4949a457d6a6","parserVersion":"test_version"}
```

Name: Aconitum W. Mucher nothoser. Acotoxicum

Canonical: Aconitum nothoser. Acotoxicum

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum W. Mucher nothoser. Acotoxicum","normalized":"Aconitum nothoser. Acotoxicum","canonical":{"stemmed":"Acotoxicum","simple":"Acotoxicum","full":"Aconitum nothoser. Acotoxicum"},"cardinality":1,"hybrid":"NOTHO_HYBRID","details":{"uninomial":{"uninomial":"Acotoxicum","rank":"nothoser.","parent":"Aconitum"}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":9,"end":11},{"verbatim":"Mucher","normalized":"Mucher","wordType":"AUTHOR_WORD","start":12,"end":18},{"verbatim":"nothoser.","normalized":"nothoser.","wordType":"RANK","start":19,"end":28},{"verbatim":"Acotoxicum","normalized":"Acotoxicum","wordType":"UNINOMIAL","start":29,"end":39}],"id":"6fd8d3d4-bdb6-5fc6-a94d-966af669c7e9","parserVersion":"test_version"}
```

Name: Abies masjoannis nothof. mesoides

Canonical: Abies masjoannis nothof. mesoides

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Abies masjoannis nothof. mesoides","normalized":"Abies masjoannis nothof. mesoides","canonical":{"stemmed":"Abies masioann mesoid","simple":"Abies masjoannis mesoides","full":"Abies masjoannis nothof. mesoides"},"cardinality":3,"hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Abies","species":"masjoannis","infraspecies":[{"value":"mesoides","rank":"nothof."}]}},"words":[{"verbatim":"Abies","normalized":"Abies","wordType":"GENUS","start":0,"end":5},{"verbatim":"masjoannis","normalized":"masjoannis","wordType":"SPECIES","start":6,"end":16},{"verbatim":"nothof.","normalized":"nothof.","wordType":"RANK","start":17,"end":24},{"verbatim":"mesoides","normalized":"mesoides","wordType":"INFRASPECIES","start":25,"end":33}],"id":"5be2cd2f-c81f-5d81-8eaf-54bd231f5230","parserVersion":"test_version"}
```

Name: Aconitum berdaui nothosubsp. walasii (Mitka) Mitka

Canonical: Aconitum berdaui nothosubsp. walasii

Authorship: (Mitka) Mitka

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum berdaui nothosubsp. walasii (Mitka) Mitka","normalized":"Aconitum berdaui nothosubsp. walasii (Mitka) Mitka","canonical":{"stemmed":"Aconitum berdau walas","simple":"Aconitum berdaui walasii","full":"Aconitum berdaui nothosubsp. walasii"},"cardinality":3,"authorship":{"verbatim":"(Mitka) Mitka","normalized":"(Mitka) Mitka","authors":["Mitka"],"originalAuth":{"authors":["Mitka"]},"combinationAuth":{"authors":["Mitka"]}},"hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Aconitum","species":"berdaui","infraspecies":[{"value":"walasii","rank":"nothosubsp.","authorship":{"verbatim":"(Mitka) Mitka","normalized":"(Mitka) Mitka","authors":["Mitka"],"originalAuth":{"authors":["Mitka"]},"combinationAuth":{"authors":["Mitka"]}}}]}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"GENUS","start":0,"end":8},{"verbatim":"berdaui","normalized":"berdaui","wordType":"SPECIES","start":9,"end":16},{"verbatim":"nothosubsp.","normalized":"nothosubsp.","wordType":"RANK","start":17,"end":28},{"verbatim":"walasii","normalized":"walasii","wordType":"INFRASPECIES","start":29,"end":36},{"verbatim":"Mitka","normalized":"Mitka","wordType":"AUTHOR_WORD","start":38,"end":43},{"verbatim":"Mitka","normalized":"Mitka","wordType":"AUTHOR_WORD","start":45,"end":50}],"id":"ba2f82ac-9312-5595-928a-2ba07aebb04f","parserVersion":"test_version"}
```

Name: Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu

Canonical: Aconitum tauricum nothossp. hayekianum

Authorship: (Gáyer) Grintescu

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu","normalized":"Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu","canonical":{"stemmed":"Aconitum tauric hayekian","simple":"Aconitum tauricum hayekianum","full":"Aconitum tauricum nothossp. hayekianum"},"cardinality":3,"authorship":{"verbatim":"(Gáyer) Grintescu","normalized":"(Gáyer) Grintescu","authors":["Gáyer","Grintescu"],"originalAuth":{"authors":["Gáyer"]},"combinationAuth":{"authors":["Grintescu"]}},"hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Aconitum","species":"tauricum","infraspecies":[{"value":"hayekianum","rank":"nothossp.","authorship":{"verbatim":"(Gáyer) Grintescu","normalized":"(Gáyer) Grintescu","authors":["Gáyer","Grintescu"],"originalAuth":{"authors":["Gáyer"]},"combinationAuth":{"authors":["Grintescu"]}}}]}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"GENUS","start":0,"end":8},{"verbatim":"tauricum","normalized":"tauricum","wordType":"SPECIES","start":9,"end":17},{"verbatim":"nothossp.","normalized":"nothossp.","wordType":"RANK","start":18,"end":27},{"verbatim":"hayekianum","normalized":"hayekianum","wordType":"INFRASPECIES","start":28,"end":38},{"verbatim":"Gáyer","normalized":"Gáyer","wordType":"AUTHOR_WORD","start":40,"end":45},{"verbatim":"Grintescu","normalized":"Grintescu","wordType":"AUTHOR_WORD","start":47,"end":56}],"id":"c02c80bf-11b1-59f9-9fed-6627fb954dd8","parserVersion":"test_version"}
```

Name: Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares

Canonical: Aeonium holospathulatum nothovar. sanchezii

Authorship: (Bañares) Bañares

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares","normalized":"Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares","canonical":{"stemmed":"Aeonium holospathulat sanchez","simple":"Aeonium holospathulatum sanchezii","full":"Aeonium holospathulatum nothovar. sanchezii"},"cardinality":3,"authorship":{"verbatim":"(Bañares) Bañares","normalized":"(Bañares) Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]},"combinationAuth":{"authors":["Bañares"]}},"hybrid":"NOTHO_HYBRID","details":{"infraspecies":{"genus":"Aeonium","species":"holospathulatum","infraspecies":[{"value":"sanchezii","rank":"nothovar.","authorship":{"verbatim":"(Bañares) Bañares","normalized":"(Bañares) Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]},"combinationAuth":{"authors":["Bañares"]}}}]}},"words":[{"verbatim":"Aeonium","normalized":"Aeonium","wordType":"GENUS","start":0,"end":7},{"verbatim":"holospathulatum","normalized":"holospathulatum","wordType":"SPECIES","start":8,"end":23},{"verbatim":"nothovar.","normalized":"nothovar.","wordType":"RANK","start":24,"end":33},{"verbatim":"sanchezii","normalized":"sanchezii","wordType":"INFRASPECIES","start":34,"end":43},{"verbatim":"Bañares","normalized":"Bañares","wordType":"AUTHOR_WORD","start":45,"end":52},{"verbatim":"Bañares","normalized":"Bañares","wordType":"AUTHOR_WORD","start":54,"end":61}],"id":"fc173db1-3977-5cad-a96b-472165bb0bbd","parserVersion":"test_version"}
```

Name: Amaranthus ×ozanonii (Contré) Lambinon nothosubsp. ralletii

Canonical: Amaranthus × ozanonii nothosubsp. ralletii

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Amaranthus ×ozanonii (Contré) Lambinon nothosubsp. ralletii","normalized":"Amaranthus × ozanonii (Contré) Lambinon nothosubsp. ralletii","canonical":{"stemmed":"Amaranthus ozanon rallet","simple":"Amaranthus ozanonii ralletii","full":"Amaranthus × ozanonii nothosubsp. ralletii"},"cardinality":3,"hybrid":"NAMED_HYBRID","details":{"infraspecies":{"genus":"Amaranthus","species":"ozanonii (Contré) Lambinon","authorship":{"verbatim":"(Contré) Lambinon","normalized":"(Contré) Lambinon","authors":["Contré","Lambinon"],"originalAuth":{"authors":["Contré"]},"combinationAuth":{"authors":["Lambinon"]}},"infraspecies":[{"value":"ralletii","rank":"nothosubsp."}]}},"words":[{"verbatim":"Amaranthus","normalized":"Amaranthus","wordType":"GENUS","start":0,"end":10},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":11,"end":12},{"verbatim":"ozanonii","normalized":"ozanonii","wordType":"SPECIES","start":12,"end":20},{"verbatim":"Contré","normalized":"Contré","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"Lambinon","normalized":"Lambinon","wordType":"AUTHOR_WORD","start":30,"end":38},{"verbatim":"nothosubsp.","normalized":"nothosubsp.","wordType":"RANK","start":39,"end":50},{"verbatim":"ralletii","normalized":"ralletii","wordType":"INFRASPECIES","start":51,"end":59}],"id":"678535c6-c679-5716-a874-1cf92bca3ce9","parserVersion":"test_version"}
```

Name: Aconitum ×teppneri Mucher ex Starm. nothosubsp. goetzii

Canonical: Aconitum × teppneri nothosubsp. goetzii

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum ×teppneri Mucher ex Starm. nothosubsp. goetzii","normalized":"Aconitum × teppneri Mucher ex Starm. nothosubsp. goetzii","canonical":{"stemmed":"Aconitum teppner goetz","simple":"Aconitum teppneri goetzii","full":"Aconitum × teppneri nothosubsp. goetzii"},"cardinality":3,"hybrid":"NAMED_HYBRID","details":{"infraspecies":{"genus":"Aconitum","species":"teppneri Mucher ex Starm.","authorship":{"verbatim":"Mucher ex Starm.","normalized":"Mucher ex Starm.","authors":["Mucher","Starm."],"originalAuth":{"authors":["Mucher"],"exAuthors":{"authors":["Starm."]}}},"infraspecies":[{"value":"goetzii","rank":"nothosubsp."}]}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"GENUS","start":0,"end":8},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":9,"end":10},{"verbatim":"teppneri","normalized":"teppneri","wordType":"SPECIES","start":10,"end":18},{"verbatim":"Mucher","normalized":"Mucher","wordType":"AUTHOR_WORD","start":19,"end":25},{"verbatim":"Starm.","normalized":"Starm.","wordType":"AUTHOR_WORD","start":29,"end":35},{"verbatim":"nothosubsp.","normalized":"nothosubsp.","wordType":"RANK","start":36,"end":47},{"verbatim":"goetzii","normalized":"goetzii","wordType":"INFRASPECIES","start":48,"end":55}],"id":"2387941b-9e4f-5fb5-a440-74934cc66c4f","parserVersion":"test_version"}
```

Name: Aeonium × proliferum Bañares nothovar. glabrifolium Bañares

Canonical: Aeonium × proliferum nothovar. glabrifolium

Authorship: Bañares

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aeonium × proliferum Bañares nothovar. glabrifolium Bañares","normalized":"Aeonium × proliferum Bañares nothovar. glabrifolium Bañares","canonical":{"stemmed":"Aeonium prolifer glabrifoli","simple":"Aeonium proliferum glabrifolium","full":"Aeonium × proliferum nothovar. glabrifolium"},"cardinality":3,"authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}},"hybrid":"NAMED_HYBRID","details":{"infraspecies":{"genus":"Aeonium","species":"proliferum Bañares","authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}},"infraspecies":[{"value":"glabrifolium","rank":"nothovar.","authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}}}]}},"words":[{"verbatim":"Aeonium","normalized":"Aeonium","wordType":"GENUS","start":0,"end":7},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":8,"end":9},{"verbatim":"proliferum","normalized":"proliferum","wordType":"SPECIES","start":10,"end":20},{"verbatim":"Bañares","normalized":"Bañares","wordType":"AUTHOR_WORD","start":21,"end":28},{"verbatim":"nothovar.","normalized":"nothovar.","wordType":"RANK","start":29,"end":38},{"verbatim":"glabrifolium","normalized":"glabrifolium","wordType":"INFRASPECIES","start":39,"end":51},{"verbatim":"Bañares","normalized":"Bañares","wordType":"AUTHOR_WORD","start":52,"end":59}],"id":"dc38d07a-f949-5c72-9463-d36a4ae96bea","parserVersion":"test_version"}
```

<!-- Very rare people make this mistake. We do not cover it yet.
Agropyron x pseudorepens notho morph. vulpinum (Rydb.) Bowden, 1965
-->

Name: Biscogniauxia nothofagi Whalley, Læssøe & Kile 1990

Canonical: Biscogniauxia nothofagi

Authorship: Whalley, Læssøe & Kile 1990

```json
{"parsed":true,"quality":1,"verbatim":"Biscogniauxia nothofagi Whalley, Læssøe \u0026 Kile 1990","normalized":"Biscogniauxia nothofagi Whalley, Læssøe \u0026 Kile 1990","canonical":{"stemmed":"Biscogniauxia nothofag","simple":"Biscogniauxia nothofagi","full":"Biscogniauxia nothofagi"},"cardinality":2,"authorship":{"verbatim":"Whalley, Læssøe \u0026 Kile 1990","normalized":"Whalley, Læssøe \u0026 Kile 1990","year":"1990","authors":["Whalley","Læssøe","Kile"],"originalAuth":{"authors":["Whalley","Læssøe","Kile"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Biscogniauxia","species":"nothofagi","authorship":{"verbatim":"Whalley, Læssøe \u0026 Kile 1990","normalized":"Whalley, Læssøe \u0026 Kile 1990","year":"1990","authors":["Whalley","Læssøe","Kile"],"originalAuth":{"authors":["Whalley","Læssøe","Kile"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Biscogniauxia","normalized":"Biscogniauxia","wordType":"GENUS","start":0,"end":13},{"verbatim":"nothofagi","normalized":"nothofagi","wordType":"SPECIES","start":14,"end":23},{"verbatim":"Whalley","normalized":"Whalley","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"Læssøe","normalized":"Læssøe","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"Kile","normalized":"Kile","wordType":"AUTHOR_WORD","start":42,"end":46},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":47,"end":51}],"id":"1f8935ad-5ae2-507e-96aa-f0bb1d22245e","parserVersion":"test_version"}
```

### Named hybrids
Name: ×Agropogon P. Fourn. 1934

Canonical: × Agropogon

Authorship: P. Fourn. 1934

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Agropogon P. Fourn. 1934","normalized":"× Agropogon P. Fourn. 1934","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P. Fourn. 1934","normalized":"P. Fourn. 1934","year":"1934","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."],"year":{"year":"1934"}}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P. Fourn. 1934","normalized":"P. Fourn. 1934","year":"1934","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."],"year":{"year":"1934"}}}}},"words":[{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":1,"end":10},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":11,"end":13},{"verbatim":"Fourn.","normalized":"Fourn.","wordType":"AUTHOR_WORD","start":14,"end":20},{"verbatim":"1934","normalized":"1934","wordType":"YEAR","start":21,"end":25}],"id":"f2bb2ddc-003e-5fc0-83b1-038dca1deb52","parserVersion":"test_version"}
```

Name: xAgropogon P. Fourn.

Canonical: × Agropogon

Authorship: P. Fourn.

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"xAgropogon P. Fourn.","normalized":"× Agropogon P. Fourn.","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P. Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P. Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}}}},"words":[{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":1,"end":10},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":11,"end":13},{"verbatim":"Fourn.","normalized":"Fourn.","wordType":"AUTHOR_WORD","start":14,"end":20}],"id":"b36871e3-e412-5b4f-a859-eb09fcf83a8e","parserVersion":"test_version"}
```

Name: XAgropogon P.Fourn.

Canonical: × Agropogon

Authorship: P. Fourn.

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"XAgropogon P.Fourn.","normalized":"× Agropogon P. Fourn.","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P.Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P.Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}}}},"words":[{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":1,"end":10},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":11,"end":13},{"verbatim":"Fourn.","normalized":"Fourn.","wordType":"AUTHOR_WORD","start":13,"end":19}],"id":"f6257985-ad38-5c29-94e2-bb305cab893a","parserVersion":"test_version"}
```

Name: × Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"× Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"words":[{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":2,"end":11}],"id":"b1858609-4fff-5a00-8d2b-0cb354100b10","parserVersion":"test_version"}
```

Name: x Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"x Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"words":[{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":2,"end":11}],"id":"79c27436-a61a-59cd-acf0-51425556e26f","parserVersion":"test_version"}
```

Name: X Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"words":[{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"UNINOMIAL","start":2,"end":11}],"id":"37eac7d5-a258-503b-ae3f-206739be74fa","parserVersion":"test_version"}
```

Name: X Cupressocyparis leylandii

Canonical: × Cupressocyparis leylandii

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Cupressocyparis leylandii","normalized":"× Cupressocyparis leylandii","canonical":{"stemmed":"Cupressocyparis leyland","simple":"Cupressocyparis leylandii","full":"× Cupressocyparis leylandii"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Cupressocyparis","species":"leylandii"}},"words":[{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Cupressocyparis","normalized":"Cupressocyparis","wordType":"GENUS","start":2,"end":17},{"verbatim":"leylandii","normalized":"leylandii","wordType":"SPECIES","start":18,"end":27}],"id":"a6ebd2cf-a021-50fe-b158-8be16844079d","parserVersion":"test_version"}
```

Name: ×Heucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Heucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"words":[{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Heucherella","normalized":"Heucherella","wordType":"GENUS","start":1,"end":12},{"verbatim":"tiarelloides","normalized":"tiarelloides","wordType":"SPECIES","start":13,"end":25}],"id":"6aab4b31-89fb-5a41-97ee-2024becc9169","parserVersion":"test_version"}
```

Name: xHeucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"xHeucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"words":[{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Heucherella","normalized":"Heucherella","wordType":"GENUS","start":1,"end":12},{"verbatim":"tiarelloides","normalized":"tiarelloides","wordType":"SPECIES","start":13,"end":25}],"id":"726d4f33-a175-5449-aea2-0e3c26dc7a0b","parserVersion":"test_version"}
```

Name: x Heucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"x Heucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"words":[{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Heucherella","normalized":"Heucherella","wordType":"GENUS","start":2,"end":13},{"verbatim":"tiarelloides","normalized":"tiarelloides","wordType":"SPECIES","start":14,"end":26}],"id":"da549587-a768-51b6-af26-1bb3c1977b31","parserVersion":"test_version"}
```

Name: XAgroelymus Lapage sect. Agroelinelymus

Canonical: × Agroelymus sect. Agroelinelymus

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Combination of two uninomials"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"XAgroelymus Lapage sect. Agroelinelymus","normalized":"× Agroelymus sect. Agroelinelymus","canonical":{"stemmed":"Agroelinelymus","simple":"Agroelinelymus","full":"× Agroelymus sect. Agroelinelymus"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agroelinelymus","rank":"sect.","parent":"Agroelymus"}},"words":[{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agroelymus","normalized":"Agroelymus","wordType":"UNINOMIAL","start":1,"end":11},{"verbatim":"Lapage","normalized":"Lapage","wordType":"AUTHOR_WORD","start":12,"end":18},{"verbatim":"sect.","normalized":"sect.","wordType":"RANK","start":19,"end":24},{"verbatim":"Agroelinelymus","normalized":"Agroelinelymus","wordType":"UNINOMIAL","start":25,"end":39}],"id":"419d1a5d-64b9-5e0d-87f4-624b19ddab0f","parserVersion":"test_version"}
```

Name: ×Agropogon littoralis (Sm.) C. E. Hubb. 1946

Canonical: × Agropogon littoralis

Authorship: (Sm.) C. E. Hubb. 1946

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Agropogon littoralis (Sm.) C. E. Hubb. 1946","normalized":"× Agropogon littoralis (Sm.) C. E. Hubb. 1946","canonical":{"stemmed":"Agropogon littoral","simple":"Agropogon littoralis","full":"× Agropogon littoralis"},"cardinality":2,"authorship":{"verbatim":"(Sm.) C. E. Hubb. 1946","normalized":"(Sm.) C. E. Hubb. 1946","authors":["Sm.","C. E. Hubb."],"originalAuth":{"authors":["Sm."]},"combinationAuth":{"authors":["C. E. Hubb."],"year":{"year":"1946"}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Agropogon","species":"littoralis","authorship":{"verbatim":"(Sm.) C. E. Hubb. 1946","normalized":"(Sm.) C. E. Hubb. 1946","authors":["Sm.","C. E. Hubb."],"originalAuth":{"authors":["Sm."]},"combinationAuth":{"authors":["C. E. Hubb."],"year":{"year":"1946"}}}}},"words":[{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Agropogon","normalized":"Agropogon","wordType":"GENUS","start":1,"end":10},{"verbatim":"littoralis","normalized":"littoralis","wordType":"SPECIES","start":11,"end":21},{"verbatim":"Sm.","normalized":"Sm.","wordType":"AUTHOR_WORD","start":23,"end":26},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":31,"end":33},{"verbatim":"Hubb.","normalized":"Hubb.","wordType":"AUTHOR_WORD","start":34,"end":39},{"verbatim":"1946","normalized":"1946","wordType":"YEAR","start":40,"end":44}],"id":"66beda81-d796-5d60-be9f-b3188ef730dc","parserVersion":"test_version"}
```

Name: Asplenium X inexpectatum (E.L. Braun 1940) Morton (1956)

Canonical: Asplenium × inexpectatum

Authorship: (E. L. Braun 1940) Morton (1956)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Asplenium X inexpectatum (E.L. Braun 1940) Morton (1956)","normalized":"Asplenium × inexpectatum (E. L. Braun 1940) Morton (1956)","canonical":{"stemmed":"Asplenium inexpectat","simple":"Asplenium inexpectatum","full":"Asplenium × inexpectatum"},"cardinality":2,"authorship":{"verbatim":"(E.L. Braun 1940) Morton (1956)","normalized":"(E. L. Braun 1940) Morton (1956)","year":"1940","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1940"}},"combinationAuth":{"authors":["Morton"],"year":{"year":"1956","isApproximate":true}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Asplenium","species":"inexpectatum (E. L. Braun 1940) Morton (1956)","authorship":{"verbatim":"(E.L. Braun 1940) Morton (1956)","normalized":"(E. L. Braun 1940) Morton (1956)","year":"1940","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1940"}},"combinationAuth":{"authors":["Morton"],"year":{"year":"1956","isApproximate":true}}}}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":10,"end":11},{"verbatim":"inexpectatum","normalized":"inexpectatum","wordType":"SPECIES","start":12,"end":24},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":31,"end":36},{"verbatim":"1940","normalized":"1940","wordType":"YEAR","start":37,"end":41},{"verbatim":"Morton","normalized":"Morton","wordType":"AUTHOR_WORD","start":43,"end":49},{"verbatim":"1956","normalized":"1956","wordType":"APPROXIMATE_YEAR","start":51,"end":55}],"id":"d37e04e4-90bc-5031-b91c-dbb61113bcfa","parserVersion":"test_version"}
```

Name: Salix ×capreola Andersson (1867)

Canonical: Salix × capreola

Authorship: Andersson (1867)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char is not separated by space"},{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Salix ×capreola Andersson (1867)","normalized":"Salix × capreola Andersson (1867)","canonical":{"stemmed":"Salix capreol","simple":"Salix capreola","full":"Salix × capreola"},"cardinality":2,"authorship":{"verbatim":"Andersson (1867)","normalized":"Andersson (1867)","year":"(1867)","authors":["Andersson"],"originalAuth":{"authors":["Andersson"],"year":{"year":"1867","isApproximate":true}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Salix","species":"capreola Andersson (1867)","authorship":{"verbatim":"Andersson (1867)","normalized":"Andersson (1867)","year":"(1867)","authors":["Andersson"],"originalAuth":{"authors":["Andersson"],"year":{"year":"1867","isApproximate":true}}}}},"words":[{"verbatim":"Salix","normalized":"Salix","wordType":"GENUS","start":0,"end":5},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":6,"end":7},{"verbatim":"capreola","normalized":"capreola","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Andersson","normalized":"Andersson","wordType":"AUTHOR_WORD","start":16,"end":25},{"verbatim":"1867","normalized":"1867","wordType":"APPROXIMATE_YEAR","start":27,"end":31}],"id":"9965be0c-0db2-506a-97f7-e709ef950ef7","parserVersion":"test_version"}
```

Name: Polypodium  x vulgare nothosubsp. mantoniae (Rothm.) Schidlay

Canonical: Polypodium × vulgare nothosubsp. mantoniae

Authorship: (Rothm.) Schidlay

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Polypodium  x vulgare nothosubsp. mantoniae (Rothm.) Schidlay","normalized":"Polypodium × vulgare nothosubsp. mantoniae (Rothm.) Schidlay","canonical":{"stemmed":"Polypodium uulgar mantoni","simple":"Polypodium vulgare mantoniae","full":"Polypodium × vulgare nothosubsp. mantoniae"},"cardinality":3,"authorship":{"verbatim":"(Rothm.) Schidlay","normalized":"(Rothm.) Schidlay","authors":["Rothm.","Schidlay"],"originalAuth":{"authors":["Rothm."]},"combinationAuth":{"authors":["Schidlay"]}},"hybrid":"NAMED_HYBRID","details":{"infraspecies":{"genus":"Polypodium","species":"vulgare","infraspecies":[{"value":"mantoniae","rank":"nothosubsp.","authorship":{"verbatim":"(Rothm.) Schidlay","normalized":"(Rothm.) Schidlay","authors":["Rothm.","Schidlay"],"originalAuth":{"authors":["Rothm."]},"combinationAuth":{"authors":["Schidlay"]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":12,"end":13},{"verbatim":"vulgare","normalized":"vulgare","wordType":"SPECIES","start":14,"end":21},{"verbatim":"nothosubsp.","normalized":"nothosubsp.","wordType":"RANK","start":22,"end":33},{"verbatim":"mantoniae","normalized":"mantoniae","wordType":"INFRASPECIES","start":34,"end":43},{"verbatim":"Rothm.","normalized":"Rothm.","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"Schidlay","normalized":"Schidlay","wordType":"AUTHOR_WORD","start":53,"end":61}],"id":"8666c370-8843-5324-a7f3-754ca778d618","parserVersion":"test_version"}
```

Name: Salix x capreola Andersson

Canonical: Salix × capreola

Authorship: Andersson

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Salix x capreola Andersson","normalized":"Salix × capreola Andersson","canonical":{"stemmed":"Salix capreol","simple":"Salix capreola","full":"Salix × capreola"},"cardinality":2,"authorship":{"verbatim":"Andersson","normalized":"Andersson","authors":["Andersson"],"originalAuth":{"authors":["Andersson"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Salix","species":"capreola Andersson","authorship":{"verbatim":"Andersson","normalized":"Andersson","authors":["Andersson"],"originalAuth":{"authors":["Andersson"]}}}},"words":[{"verbatim":"Salix","normalized":"Salix","wordType":"GENUS","start":0,"end":5},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":6,"end":7},{"verbatim":"capreola","normalized":"capreola","wordType":"SPECIES","start":8,"end":16},{"verbatim":"Andersson","normalized":"Andersson","wordType":"AUTHOR_WORD","start":17,"end":26}],"id":"5780473c-18ac-5386-9c3a-f74bbe426624","parserVersion":"test_version"}
```

### Hybrid formulae

Name: Stanhopea tigrina Bateman ex Lindl. x S. ecornuta Lem.

Canonical: Stanhopea tigrina × Stanhopea ecornuta

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Stanhopea tigrina Bateman ex Lindl. x S. ecornuta Lem.","normalized":"Stanhopea tigrina Bateman ex Lindl. × Stanhopea ecornuta Lem.","canonical":{"stemmed":"Stanhopea tigrin × Stanhopea ecornut","simple":"Stanhopea tigrina × Stanhopea ecornuta","full":"Stanhopea tigrina × Stanhopea ecornuta"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Stanhopea","species":"tigrina","authorship":{"verbatim":"Bateman ex Lindl.","normalized":"Bateman ex Lindl.","authors":["Bateman","Lindl."],"originalAuth":{"authors":["Bateman"],"exAuthors":{"authors":["Lindl."]}}}}},{"species":{"genus":"Stanhopea","species":"ecornuta","authorship":{"verbatim":"Lem.","normalized":"Lem.","authors":["Lem."],"originalAuth":{"authors":["Lem."]}}}}]},"words":[{"verbatim":"Stanhopea","normalized":"Stanhopea","wordType":"GENUS","start":0,"end":9},{"verbatim":"tigrina","normalized":"tigrina","wordType":"SPECIES","start":10,"end":17},{"verbatim":"Bateman","normalized":"Bateman","wordType":"AUTHOR_WORD","start":18,"end":25},{"verbatim":"Lindl.","normalized":"Lindl.","wordType":"AUTHOR_WORD","start":29,"end":35},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":36,"end":37},{"verbatim":"S.","normalized":"Stanhopea","wordType":"GENUS","start":38,"end":40},{"verbatim":"ecornuta","normalized":"ecornuta","wordType":"SPECIES","start":41,"end":49},{"verbatim":"Lem.","normalized":"Lem.","wordType":"AUTHOR_WORD","start":50,"end":54}],"id":"80c0a17d-3422-515c-88bc-3a927438df88","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora X Hydnellum scrobiculatum

Canonical: Arthopyrenia hyalospora × Hydnellum scrobiculatum

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Arthopyrenia hyalospora X Hydnellum scrobiculatum","normalized":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","canonical":{"stemmed":"Arthopyrenia hyalospor × Hydnellum scrobiculat","simple":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","full":"Arthopyrenia hyalospora × Hydnellum scrobiculatum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}},{"species":{"genus":"Hydnellum","species":"scrobiculatum"}}]},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":24,"end":25},{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":26,"end":35},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":36,"end":49}],"id":"e78d9299-9fd4-55d2-aeb4-b2864f5bff45","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Banker) D. Hall X Hydnellum scrobiculatum D.E. Stuntz

Canonical: Arthopyrenia hyalospora × Hydnellum scrobiculatum

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Arthopyrenia hyalospora (Banker) D. Hall X Hydnellum scrobiculatum D.E. Stuntz","normalized":"Arthopyrenia hyalospora (Banker) D. Hall × Hydnellum scrobiculatum D. E. Stuntz","canonical":{"stemmed":"Arthopyrenia hyalospor × Hydnellum scrobiculat","simple":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","full":"Arthopyrenia hyalospora × Hydnellum scrobiculatum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Banker) D. Hall","normalized":"(Banker) D. Hall","authors":["Banker","D. Hall"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall"]}}}},{"species":{"genus":"Hydnellum","species":"scrobiculatum","authorship":{"verbatim":"D.E. Stuntz","normalized":"D. E. Stuntz","authors":["D. E. Stuntz"],"originalAuth":{"authors":["D. E. Stuntz"]}}}}]},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":25,"end":31},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"Hall","normalized":"Hall","wordType":"AUTHOR_WORD","start":36,"end":40},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":41,"end":42},{"verbatim":"Hydnellum","normalized":"Hydnellum","wordType":"GENUS","start":43,"end":52},{"verbatim":"scrobiculatum","normalized":"scrobiculatum","wordType":"SPECIES","start":53,"end":66},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":67,"end":69},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":69,"end":71},{"verbatim":"Stuntz","normalized":"Stuntz","wordType":"AUTHOR_WORD","start":72,"end":78}],"id":"a13ac2e0-5eec-569c-af9c-dd8163dbbd72","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora x

Canonical: Arthopyrenia hyalospora ×

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"},{"quality":2,"warning":"Probably incomplete hybrid formula"}],"verbatim":"Arthopyrenia hyalospora x","normalized":"Arthopyrenia hyalospora ×","canonical":{"stemmed":"Arthopyrenia hyalospor ×","simple":"Arthopyrenia hyalospora ×","full":"Arthopyrenia hyalospora ×"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}}]},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":24,"end":25}],"id":"c056b89e-789b-5c28-89e7-e820ea0baebf","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora × ?

Canonical: Arthopyrenia hyalospora ×

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Hybrid formula"},{"quality":2,"warning":"Probably incomplete hybrid formula"}],"verbatim":"Arthopyrenia hyalospora × ?","normalized":"Arthopyrenia hyalospora ×","canonical":{"stemmed":"Arthopyrenia hyalospor ×","simple":"Arthopyrenia hyalospora ×","full":"Arthopyrenia hyalospora ×"},"cardinality":0,"hybrid":"HYBRID_FORMULA","tail":" ?","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}}]},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":24,"end":25}],"id":"638cc013-3821-55c2-b9d3-b2ea3de33ecf","parserVersion":"test_version"}
```

Name: Agrostis L. × Polypogon Desf.

Canonical: Agrostis × Polypogon

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Agrostis L. × Polypogon Desf.","normalized":"Agrostis L. × Polypogon Desf.","canonical":{"stemmed":"Agrostis × Polypogon","simple":"Agrostis × Polypogon","full":"Agrostis × Polypogon"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"uninomial":{"uninomial":"Agrostis","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"uninomial":{"uninomial":"Polypogon","authorship":{"verbatim":"Desf.","normalized":"Desf.","authors":["Desf."],"originalAuth":{"authors":["Desf."]}}}}]},"words":[{"verbatim":"Agrostis","normalized":"Agrostis","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":9,"end":11},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":12,"end":13},{"verbatim":"Polypogon","normalized":"Polypogon","wordType":"UNINOMIAL","start":14,"end":23},{"verbatim":"Desf.","normalized":"Desf.","wordType":"AUTHOR_WORD","start":24,"end":29}],"id":"e914b63f-f19a-5437-ad19-85bfc98a0de2","parserVersion":"test_version"}
```

Name: Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.

Canonical: Agrostis stolonifera × Polypogon monspeliensis

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.","normalized":"Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.","canonical":{"stemmed":"Agrostis stolonifer × Polypogon monspeliens","simple":"Agrostis stolonifera × Polypogon monspeliensis","full":"Agrostis stolonifera × Polypogon monspeliensis"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Agrostis","species":"stolonifera","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"species":{"genus":"Polypogon","species":"monspeliensis","authorship":{"verbatim":"(L.) Desf.","normalized":"(L.) Desf.","authors":["L.","Desf."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Desf."]}}}}]},"words":[{"verbatim":"Agrostis","normalized":"Agrostis","wordType":"GENUS","start":0,"end":8},{"verbatim":"stolonifera","normalized":"stolonifera","wordType":"SPECIES","start":9,"end":20},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":24,"end":25},{"verbatim":"Polypogon","normalized":"Polypogon","wordType":"GENUS","start":26,"end":35},{"verbatim":"monspeliensis","normalized":"monspeliensis","wordType":"SPECIES","start":36,"end":49},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"Desf.","normalized":"Desf.","wordType":"AUTHOR_WORD","start":55,"end":60}],"id":"a2aeb842-18c5-54b4-a4d9-c78bd0445c10","parserVersion":"test_version"}
```

Name: Coeloglossum viride (L.) Hartman x Dactylorhiza majalis (Rchb. f.) P.F. Hunt & Summerhayes ssp. praetermissa (Druce) D.M. Moore & Soó

Canonical: Coeloglossum viride × Dactylorhiza majalis subsp. praetermissa

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Coeloglossum viride (L.) Hartman x Dactylorhiza majalis (Rchb. f.) P.F. Hunt \u0026 Summerhayes ssp. praetermissa (Druce) D.M. Moore \u0026 Soó","normalized":"Coeloglossum viride (L.) Hartman × Dactylorhiza majalis (Rchb. fil.) P. F. Hunt \u0026 Summerhayes subsp. praetermissa (Druce) D. M. Moore \u0026 Soó","canonical":{"stemmed":"Coeloglossum uirid × Dactylorhiza maial praetermiss","simple":"Coeloglossum viride × Dactylorhiza majalis praetermissa","full":"Coeloglossum viride × Dactylorhiza majalis subsp. praetermissa"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Coeloglossum","species":"viride","authorship":{"verbatim":"(L.) Hartman","normalized":"(L.) Hartman","authors":["L.","Hartman"],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Hartman"]}}}},{"infraspecies":{"genus":"Dactylorhiza","species":"majalis","authorship":{"verbatim":"(Rchb. f.) P.F. Hunt \u0026 Summerhayes","normalized":"(Rchb. fil.) P. F. Hunt \u0026 Summerhayes","authors":["Rchb. fil.","P. F. Hunt","Summerhayes"],"originalAuth":{"authors":["Rchb. fil."]},"combinationAuth":{"authors":["P. F. Hunt","Summerhayes"]}},"infraspecies":[{"value":"praetermissa","rank":"subsp.","authorship":{"verbatim":"(Druce) D.M. Moore \u0026 Soó","normalized":"(Druce) D. M. Moore \u0026 Soó","authors":["Druce","D. M. Moore","Soó"],"originalAuth":{"authors":["Druce"]},"combinationAuth":{"authors":["D. M. Moore","Soó"]}}}]}}]},"words":[{"verbatim":"Coeloglossum","normalized":"Coeloglossum","wordType":"GENUS","start":0,"end":12},{"verbatim":"viride","normalized":"viride","wordType":"SPECIES","start":13,"end":19},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"Hartman","normalized":"Hartman","wordType":"AUTHOR_WORD","start":25,"end":32},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":33,"end":34},{"verbatim":"Dactylorhiza","normalized":"Dactylorhiza","wordType":"GENUS","start":35,"end":47},{"verbatim":"majalis","normalized":"majalis","wordType":"SPECIES","start":48,"end":55},{"verbatim":"Rchb.","normalized":"Rchb.","wordType":"AUTHOR_WORD","start":57,"end":62},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":63,"end":65},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":67,"end":69},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":69,"end":71},{"verbatim":"Hunt","normalized":"Hunt","wordType":"AUTHOR_WORD","start":72,"end":76},{"verbatim":"Summerhayes","normalized":"Summerhayes","wordType":"AUTHOR_WORD","start":79,"end":90},{"verbatim":"ssp.","normalized":"subsp.","wordType":"RANK","start":91,"end":95},{"verbatim":"praetermissa","normalized":"praetermissa","wordType":"INFRASPECIES","start":96,"end":108},{"verbatim":"Druce","normalized":"Druce","wordType":"AUTHOR_WORD","start":110,"end":115},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":117,"end":119},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":119,"end":121},{"verbatim":"Moore","normalized":"Moore","wordType":"AUTHOR_WORD","start":122,"end":127},{"verbatim":"Soó","normalized":"Soó","wordType":"AUTHOR_WORD","start":130,"end":133}],"id":"76fc857a-442a-590e-98c6-174aeb199e68","parserVersion":"test_version"}
```

Name: Salix aurita L. × S. caprea L.

Canonical: Salix aurita × Salix caprea

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Salix aurita L. × S. caprea L.","normalized":"Salix aurita L. × Salix caprea L.","canonical":{"stemmed":"Salix aurit × Salix capre","simple":"Salix aurita × Salix caprea","full":"Salix aurita × Salix caprea"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Salix","species":"aurita","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"species":{"genus":"Salix","species":"caprea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}}]},"words":[{"verbatim":"Salix","normalized":"Salix","wordType":"GENUS","start":0,"end":5},{"verbatim":"aurita","normalized":"aurita","wordType":"SPECIES","start":6,"end":12},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":13,"end":15},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":16,"end":17},{"verbatim":"S.","normalized":"Salix","wordType":"GENUS","start":18,"end":20},{"verbatim":"caprea","normalized":"caprea","wordType":"SPECIES","start":21,"end":27},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":28,"end":30}],"id":"a8de3172-b5e8-55c0-b495-b13b7af462d4","parserVersion":"test_version"}
```

Name: Asplenium rhizophyllum X A. ruta-muraria E.L. Braun 1939

Canonical: Asplenium rhizophyllum × Asplenium ruta-muraria

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Asplenium rhizophyllum X A. ruta-muraria E.L. Braun 1939","normalized":"Asplenium rhizophyllum × Asplenium ruta-muraria E. L. Braun 1939","canonical":{"stemmed":"Asplenium rhizophyll × Asplenium ruta-murar","simple":"Asplenium rhizophyllum × Asplenium ruta-muraria","full":"Asplenium rhizophyllum × Asplenium ruta-muraria"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Asplenium","species":"rhizophyllum"}},{"species":{"genus":"Asplenium","species":"ruta-muraria","authorship":{"verbatim":"E.L. Braun 1939","normalized":"E. L. Braun 1939","year":"1939","authors":["E. L. Braun"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1939"}}}}}]},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"rhizophyllum","normalized":"rhizophyllum","wordType":"SPECIES","start":10,"end":22},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":23,"end":24},{"verbatim":"A.","normalized":"Asplenium","wordType":"GENUS","start":25,"end":27},{"verbatim":"ruta-muraria","normalized":"ruta-muraria","wordType":"SPECIES","start":28,"end":40},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":46,"end":51},{"verbatim":"1939","normalized":"1939","wordType":"YEAR","start":52,"end":56}],"id":"1fa2c609-ce9b-5eea-a1b2-187d36b695cb","parserVersion":"test_version"}
```

Name: Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939

Canonical: Asplenium rhizophyllum × Asplenium ruta-muraria

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Incomplete hybrid formula"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939","normalized":"Asplenium rhizophyllum DC. × Asplenium ruta-muraria E. L. Braun 1939","canonical":{"stemmed":"Asplenium rhizophyll × Asplenium ruta-murar","simple":"Asplenium rhizophyllum × Asplenium ruta-muraria","full":"Asplenium rhizophyllum × Asplenium ruta-muraria"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Asplenium","species":"rhizophyllum","authorship":{"verbatim":"DC.","normalized":"DC.","authors":["DC."],"originalAuth":{"authors":["DC."]}}}},{"species":{"genus":"Asplenium","species":"ruta-muraria","authorship":{"verbatim":"E.L. Braun 1939","normalized":"E. L. Braun 1939","year":"1939","authors":["E. L. Braun"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1939"}}}}}]},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"rhizophyllum","normalized":"rhizophyllum","wordType":"SPECIES","start":10,"end":22},{"verbatim":"DC.","normalized":"DC.","wordType":"AUTHOR_WORD","start":23,"end":26},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":27,"end":28},{"verbatim":"ruta-muraria","normalized":"ruta-muraria","wordType":"SPECIES","start":29,"end":41},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":42,"end":44},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":47,"end":52},{"verbatim":"1939","normalized":"1939","wordType":"YEAR","start":53,"end":57}],"id":"dcb8fb0f-8207-5c67-b02b-81c8e03001b2","parserVersion":"test_version"}
```

<!--
TODO Mentha aquatica L. × M. arvensis L. × M. spicata L.|''
TODO Polypodium vulgare subsp. prionodes (Asch.) Rothm. × subsp. vulgare|''
-->

Name: Tilletia caries (Bjerk.) Tul. × T. foetida (Wallr.) Liro.

Canonical: Tilletia caries × Tilletia foetida

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Tilletia caries (Bjerk.) Tul. × T. foetida (Wallr.) Liro.","normalized":"Tilletia caries (Bjerk.) Tul. × Tilletia foetida (Wallr.) Liro.","canonical":{"stemmed":"Tilletia cari × Tilletia foetid","simple":"Tilletia caries × Tilletia foetida","full":"Tilletia caries × Tilletia foetida"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Tilletia","species":"caries","authorship":{"verbatim":"(Bjerk.) Tul.","normalized":"(Bjerk.) Tul.","authors":["Bjerk.","Tul."],"originalAuth":{"authors":["Bjerk."]},"combinationAuth":{"authors":["Tul."]}}}},{"species":{"genus":"Tilletia","species":"foetida","authorship":{"verbatim":"(Wallr.) Liro.","normalized":"(Wallr.) Liro.","authors":["Wallr.","Liro."],"originalAuth":{"authors":["Wallr."]},"combinationAuth":{"authors":["Liro."]}}}}]},"words":[{"verbatim":"Tilletia","normalized":"Tilletia","wordType":"GENUS","start":0,"end":8},{"verbatim":"caries","normalized":"caries","wordType":"SPECIES","start":9,"end":15},{"verbatim":"Bjerk.","normalized":"Bjerk.","wordType":"AUTHOR_WORD","start":17,"end":23},{"verbatim":"Tul.","normalized":"Tul.","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":30,"end":31},{"verbatim":"T.","normalized":"Tilletia","wordType":"GENUS","start":32,"end":34},{"verbatim":"foetida","normalized":"foetida","wordType":"SPECIES","start":35,"end":42},{"verbatim":"Wallr.","normalized":"Wallr.","wordType":"AUTHOR_WORD","start":44,"end":50},{"verbatim":"Liro.","normalized":"Liro.","wordType":"AUTHOR_WORD","start":52,"end":57}],"id":"65d2072c-86e1-5205-a188-0d554dccd0e7","parserVersion":"test_version"}
```

Name: Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × B. oleracea L. subsp. capitata (L.) var. costata DC.

Canonical: Brassica oleracea subsp. capitata convar. fruticosa × Brassica oleracea subsp. capitata var. costata

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × B. oleracea L. subsp. capitata (L.) var. costata DC.","normalized":"Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × Brassica oleracea L. subsp. capitata (L.) var. costata DC.","canonical":{"stemmed":"Brassica olerace capitat fruticos × Brassica olerace capitat costat","simple":"Brassica oleracea capitata fruticosa × Brassica oleracea capitata costata","full":"Brassica oleracea subsp. capitata convar. fruticosa × Brassica oleracea subsp. capitata var. costata"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"infraspecies":{"genus":"Brassica","species":"oleracea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraspecies":[{"value":"capitata","rank":"subsp.","authorship":{"verbatim":"(L.) DC.","normalized":"(L.) DC.","authors":["L.","DC."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["DC."]}}},{"value":"fruticosa","rank":"convar.","authorship":{"verbatim":"(Metzg.) Alef.","normalized":"(Metzg.) Alef.","authors":["Metzg.","Alef."],"originalAuth":{"authors":["Metzg."]},"combinationAuth":{"authors":["Alef."]}}}]}},{"infraspecies":{"genus":"Brassica","species":"oleracea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraspecies":[{"value":"capitata","rank":"subsp.","authorship":{"verbatim":"(L.)","normalized":"(L.)","authors":["L."],"originalAuth":{"authors":["L."]}}},{"value":"costata","rank":"var.","authorship":{"verbatim":"DC.","normalized":"DC.","authors":["DC."],"originalAuth":{"authors":["DC."]}}}]}}]},"words":[{"verbatim":"Brassica","normalized":"Brassica","wordType":"GENUS","start":0,"end":8},{"verbatim":"oleracea","normalized":"oleracea","wordType":"SPECIES","start":9,"end":17},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":18,"end":20},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":21,"end":27},{"verbatim":"capitata","normalized":"capitata","wordType":"INFRASPECIES","start":28,"end":36},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"DC.","normalized":"DC.","wordType":"AUTHOR_WORD","start":42,"end":45},{"verbatim":"convar.","normalized":"convar.","wordType":"RANK","start":46,"end":53},{"verbatim":"fruticosa","normalized":"fruticosa","wordType":"INFRASPECIES","start":54,"end":63},{"verbatim":"Metzg.","normalized":"Metzg.","wordType":"AUTHOR_WORD","start":65,"end":71},{"verbatim":"Alef.","normalized":"Alef.","wordType":"AUTHOR_WORD","start":73,"end":78},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":79,"end":80},{"verbatim":"B.","normalized":"Brassica","wordType":"GENUS","start":81,"end":83},{"verbatim":"oleracea","normalized":"oleracea","wordType":"SPECIES","start":84,"end":92},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":93,"end":95},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":96,"end":102},{"verbatim":"capitata","normalized":"capitata","wordType":"INFRASPECIES","start":103,"end":111},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":113,"end":115},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":117,"end":121},{"verbatim":"costata","normalized":"costata","wordType":"INFRASPECIES","start":122,"end":129},{"verbatim":"DC.","normalized":"DC.","wordType":"AUTHOR_WORD","start":130,"end":133}],"id":"2e0f4d35-ccd2-5d4a-ab42-956932ea8fb0","parserVersion":"test_version"}
```

Name: Ambystoma laterale × A. texanum × A. tigrinum

Canonical: Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Ambystoma laterale × A. texanum × A. tigrinum","normalized":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum","canonical":{"stemmed":"Ambystoma lateral × Ambystoma texan × Ambystoma tigrin","simple":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum","full":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Ambystoma","species":"laterale"}},{"species":{"genus":"Ambystoma","species":"texanum"}},{"species":{"genus":"Ambystoma","species":"tigrinum"}}]},"words":[{"verbatim":"Ambystoma","normalized":"Ambystoma","wordType":"GENUS","start":0,"end":9},{"verbatim":"laterale","normalized":"laterale","wordType":"SPECIES","start":10,"end":18},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":19,"end":20},{"verbatim":"A.","normalized":"Ambystoma","wordType":"GENUS","start":21,"end":23},{"verbatim":"texanum","normalized":"texanum","wordType":"SPECIES","start":24,"end":31},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":32,"end":33},{"verbatim":"A.","normalized":"Ambystoma","wordType":"GENUS","start":34,"end":36},{"verbatim":"tigrinum","normalized":"tigrinum","wordType":"SPECIES","start":37,"end":45}],"id":"ae91df82-158b-5307-83eb-f448044acec5","parserVersion":"test_version"}
```

<!-- NOTE: handle 'X' in author name correctly -->
Name: Pseudocercospora broussonetiae (Chupp & Linder) X.J. Liu & Y.L. Guo 1989

Canonical: Pseudocercospora broussonetiae

Authorship: (Chupp & Linder) X. J. Liu & Y. L. Guo 1989

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora broussonetiae (Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"Pseudocercospora broussonetiae (Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","canonical":{"stemmed":"Pseudocercospora broussoneti","simple":"Pseudocercospora broussonetiae","full":"Pseudocercospora broussonetiae"},"cardinality":2,"authorship":{"verbatim":"(Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"(Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","authors":["Chupp","Linder","X. J. Liu","Y. L. Guo"],"originalAuth":{"authors":["Chupp","Linder"]},"combinationAuth":{"authors":["X. J. Liu","Y. L. Guo"],"year":{"year":"1989"}}},"details":{"species":{"genus":"Pseudocercospora","species":"broussonetiae","authorship":{"verbatim":"(Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"(Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","authors":["Chupp","Linder","X. J. Liu","Y. L. Guo"],"originalAuth":{"authors":["Chupp","Linder"]},"combinationAuth":{"authors":["X. J. Liu","Y. L. Guo"],"year":{"year":"1989"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"broussonetiae","normalized":"broussonetiae","wordType":"SPECIES","start":17,"end":30},{"verbatim":"Chupp","normalized":"Chupp","wordType":"AUTHOR_WORD","start":32,"end":37},{"verbatim":"Linder","normalized":"Linder","wordType":"AUTHOR_WORD","start":40,"end":46},{"verbatim":"X.","normalized":"X.","wordType":"AUTHOR_WORD","start":48,"end":50},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":50,"end":52},{"verbatim":"Liu","normalized":"Liu","wordType":"AUTHOR_WORD","start":53,"end":56},{"verbatim":"Y.","normalized":"Y.","wordType":"AUTHOR_WORD","start":59,"end":61},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":61,"end":63},{"verbatim":"Guo","normalized":"Guo","wordType":"AUTHOR_WORD","start":64,"end":67},{"verbatim":"1989","normalized":"1989","wordType":"YEAR","start":68,"end":72}],"id":"64f92545-9139-5e53-9ba5-c5c9edb51be5","parserVersion":"test_version"}
```

### Graft-chimeras

Name: + Crataegomespilus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"+ Crataegomespilus","cardinality":0,"id":"408e8fc7-fa27-53a6-9eff-37cb779724e4","parserVersion":"test_version"}
```

Name: +Crataegomespilus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"+Crataegomespilus","cardinality":0,"id":"c2c50c08-1f62-547f-8fab-50359caf0b31","parserVersion":"test_version"}
```

Name: Cytisus purpureus + Laburnum anagyroides

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Cytisus purpureus + Laburnum anagyroides","cardinality":0,"id":"a8f8ace8-ba1a-5371-b9d5-73efce81d52c","parserVersion":"test_version"}
```

Name: Crataegus + Mespilus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Crataegus + Mespilus","cardinality":0,"id":"d651cd82-9b00-53dd-9d59-6af66ab62046","parserVersion":"test_version"}
```

### Genus with hyphen (allowed by ICN)

Name: Saxo-Fridericia R. H. Schomb.

Canonical: Saxo-fridericia

Authorship: R. H. Schomb.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Saxo-Fridericia R. H. Schomb.","normalized":"Saxo-fridericia R. H. Schomb.","canonical":{"stemmed":"Saxo-fridericia","simple":"Saxo-fridericia","full":"Saxo-fridericia"},"cardinality":1,"authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}},"details":{"uninomial":{"uninomial":"Saxo-fridericia","authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}}}},"words":[{"verbatim":"Saxo-Fridericia","normalized":"Saxo-fridericia","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":19,"end":21},{"verbatim":"Schomb.","normalized":"Schomb.","wordType":"AUTHOR_WORD","start":22,"end":29}],"id":"f11d6164-5f08-5bb3-8432-5f07d1ee3bd4","parserVersion":"test_version"}
```

Name: Saxo-fridericia R. H. Schomb.

Canonical: Saxo-fridericia

Authorship: R. H. Schomb.

```json
{"parsed":true,"quality":1,"verbatim":"Saxo-fridericia R. H. Schomb.","normalized":"Saxo-fridericia R. H. Schomb.","canonical":{"stemmed":"Saxo-fridericia","simple":"Saxo-fridericia","full":"Saxo-fridericia"},"cardinality":1,"authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}},"details":{"uninomial":{"uninomial":"Saxo-fridericia","authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}}}},"words":[{"verbatim":"Saxo-fridericia","normalized":"Saxo-fridericia","wordType":"UNINOMIAL","start":0,"end":15},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":16,"end":18},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":19,"end":21},{"verbatim":"Schomb.","normalized":"Schomb.","wordType":"AUTHOR_WORD","start":22,"end":29}],"id":"9eac48bf-fbb1-57a3-b171-0b3bfda9757f","parserVersion":"test_version"}
```

Name: Uva-ursi cinerea (Howell) A. Heller

Canonical: Uva-ursi cinerea

Authorship: (Howell) A. Heller

```json
{"parsed":true,"quality":1,"verbatim":"Uva-ursi cinerea (Howell) A. Heller","normalized":"Uva-ursi cinerea (Howell) A. Heller","canonical":{"stemmed":"Uva-ursi cinere","simple":"Uva-ursi cinerea","full":"Uva-ursi cinerea"},"cardinality":2,"authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}},"details":{"species":{"genus":"Uva-ursi","species":"cinerea","authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}}}},"words":[{"verbatim":"Uva-ursi","normalized":"Uva-ursi","wordType":"GENUS","start":0,"end":8},{"verbatim":"cinerea","normalized":"cinerea","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Howell","normalized":"Howell","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"Heller","normalized":"Heller","wordType":"AUTHOR_WORD","start":29,"end":35}],"id":"1f0bc087-ceec-5326-9fa1-2ce3b369bd7d","parserVersion":"test_version"}
```

Name: Uva-Ursi cinerea (Howell) A. Heller

Canonical: Uva-ursi cinerea

Authorship: (Howell) A. Heller

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Uva-Ursi cinerea (Howell) A. Heller","normalized":"Uva-ursi cinerea (Howell) A. Heller","canonical":{"stemmed":"Uva-ursi cinere","simple":"Uva-ursi cinerea","full":"Uva-ursi cinerea"},"cardinality":2,"authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}},"details":{"species":{"genus":"Uva-ursi","species":"cinerea","authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}}}},"words":[{"verbatim":"Uva-Ursi","normalized":"Uva-ursi","wordType":"GENUS","start":0,"end":8},{"verbatim":"cinerea","normalized":"cinerea","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Howell","normalized":"Howell","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"Heller","normalized":"Heller","wordType":"AUTHOR_WORD","start":29,"end":35}],"id":"c89977a6-b948-5d3f-b4f2-d25b4d0b6ea0","parserVersion":"test_version"}
```

Name: Prunus-lauro-cerasus

Canonical: Prunus-lauro-cerasus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Prunus-lauro-cerasus","normalized":"Prunus-lauro-cerasus","canonical":{"stemmed":"Prunus-lauro-cerasus","simple":"Prunus-lauro-cerasus","full":"Prunus-lauro-cerasus"},"cardinality":1,"details":{"uninomial":{"uninomial":"Prunus-lauro-cerasus"}},"words":[{"verbatim":"Prunus-lauro-cerasus","normalized":"Prunus-lauro-cerasus","wordType":"UNINOMIAL","start":0,"end":20}],"id":"e23ffe7a-f6ef-5276-a591-93e328213992","parserVersion":"test_version"}
```

Name: Prunus-Lauro-Cerasus

Canonical: Prunus-lauro-cerasus

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Prunus-Lauro-Cerasus","normalized":"Prunus-lauro-cerasus","canonical":{"stemmed":"Prunus-lauro-cerasus","simple":"Prunus-lauro-cerasus","full":"Prunus-lauro-cerasus"},"cardinality":1,"details":{"uninomial":{"uninomial":"Prunus-lauro-cerasus"}},"words":[{"verbatim":"Prunus-Lauro-Cerasus","normalized":"Prunus-lauro-cerasus","wordType":"UNINOMIAL","start":0,"end":20}],"id":"192bf946-803d-53b4-934d-365a8b2798e4","parserVersion":"test_version"}
```
Name: Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan & Gaussen

Canonical: Tsugo-piceo-picea × crassifolia

Authorship: (Flous) Campo-Duplan & Gaussen

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan \u0026 Gaussen","normalized":"Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan \u0026 Gaussen","canonical":{"stemmed":"Tsugo-piceo-picea crassifol","simple":"Tsugo-piceo-picea crassifolia","full":"Tsugo-piceo-picea × crassifolia"},"cardinality":2,"authorship":{"verbatim":"(Flous) Campo-Duplan \u0026 Gaussen","normalized":"(Flous) Campo-Duplan \u0026 Gaussen","authors":["Flous","Campo-Duplan","Gaussen"],"originalAuth":{"authors":["Flous"]},"combinationAuth":{"authors":["Campo-Duplan","Gaussen"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Tsugo-piceo-picea","species":"crassifolia (Flous) Campo-Duplan \u0026 Gaussen","authorship":{"verbatim":"(Flous) Campo-Duplan \u0026 Gaussen","normalized":"(Flous) Campo-Duplan \u0026 Gaussen","authors":["Flous","Campo-Duplan","Gaussen"],"originalAuth":{"authors":["Flous"]},"combinationAuth":{"authors":["Campo-Duplan","Gaussen"]}}}},"words":[{"verbatim":"Tsugo-piceo-picea","normalized":"Tsugo-piceo-picea","wordType":"GENUS","start":0,"end":17},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":18,"end":19},{"verbatim":"crassifolia","normalized":"crassifolia","wordType":"SPECIES","start":20,"end":31},{"verbatim":"Flous","normalized":"Flous","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"Campo-Duplan","normalized":"Campo-Duplan","wordType":"AUTHOR_WORD","start":40,"end":52},{"verbatim":"Gaussen","normalized":"Gaussen","wordType":"AUTHOR_WORD","start":55,"end":62}],"id":"a00c94bb-566b-5433-a666-d56c1495ca3b","parserVersion":"test_version"}
```
<!-- 3-dashes in genera are not allowed -->

Name: Tsugo-piceo-piceo-picea × crassifolia

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Tsugo-piceo-piceo-picea × crassifolia","cardinality":0,"id":"0ab8c5ed-b224-5c17-9957-298a80cc07be","parserVersion":"test_version"}
```

<!-- Xx- genera are extremely rare -->

Name: De-Filippii Gortani & Merla 1934

Canonical: De-filippii

Authorship: Gortani & Merla 1934

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"De-Filippii Gortani \u0026 Merla 1934","normalized":"De-filippii Gortani \u0026 Merla 1934","canonical":{"stemmed":"De-filippii","simple":"De-filippii","full":"De-filippii"},"cardinality":1,"authorship":{"verbatim":"Gortani \u0026 Merla 1934","normalized":"Gortani \u0026 Merla 1934","year":"1934","authors":["Gortani","Merla"],"originalAuth":{"authors":["Gortani","Merla"],"year":{"year":"1934"}}},"details":{"uninomial":{"uninomial":"De-filippii","authorship":{"verbatim":"Gortani \u0026 Merla 1934","normalized":"Gortani \u0026 Merla 1934","year":"1934","authors":["Gortani","Merla"],"originalAuth":{"authors":["Gortani","Merla"],"year":{"year":"1934"}}}}},"words":[{"verbatim":"De-Filippii","normalized":"De-filippii","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Gortani","normalized":"Gortani","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"Merla","normalized":"Merla","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"1934","normalized":"1934","wordType":"YEAR","start":28,"end":32}],"id":"5b79c27f-b0b2-5e35-a2d9-ace7d9bffce7","parserVersion":"test_version"}
```

Name: Eu-Scalpellum Hoek, 1907

Canonical: Eu-scalpellum

Authorship: Hoek 1907

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Eu-Scalpellum Hoek, 1907","normalized":"Eu-scalpellum Hoek 1907","canonical":{"stemmed":"Eu-scalpellum","simple":"Eu-scalpellum","full":"Eu-scalpellum"},"cardinality":1,"authorship":{"verbatim":"Hoek, 1907","normalized":"Hoek 1907","year":"1907","authors":["Hoek"],"originalAuth":{"authors":["Hoek"],"year":{"year":"1907"}}},"details":{"uninomial":{"uninomial":"Eu-scalpellum","authorship":{"verbatim":"Hoek, 1907","normalized":"Hoek 1907","year":"1907","authors":["Hoek"],"originalAuth":{"authors":["Hoek"],"year":{"year":"1907"}}}}},"words":[{"verbatim":"Eu-Scalpellum","normalized":"Eu-scalpellum","wordType":"UNINOMIAL","start":0,"end":13},{"verbatim":"Hoek","normalized":"Hoek","wordType":"AUTHOR_WORD","start":14,"end":18},{"verbatim":"1907","normalized":"1907","wordType":"YEAR","start":20,"end":24}],"id":"a071e617-ea3d-5792-95ac-29f59136f6be","parserVersion":"test_version"}
```

Name: Eu-hookeria olfersiana (Hornsch.) Hampe

Canonical: Eu-hookeria olfersiana

Authorship: (Hornsch.) Hampe

```json
{"parsed":true,"quality":1,"verbatim":"Eu-hookeria olfersiana (Hornsch.) Hampe","normalized":"Eu-hookeria olfersiana (Hornsch.) Hampe","canonical":{"stemmed":"Eu-hookeria olfersian","simple":"Eu-hookeria olfersiana","full":"Eu-hookeria olfersiana"},"cardinality":2,"authorship":{"verbatim":"(Hornsch.) Hampe","normalized":"(Hornsch.) Hampe","authors":["Hornsch.","Hampe"],"originalAuth":{"authors":["Hornsch."]},"combinationAuth":{"authors":["Hampe"]}},"details":{"species":{"genus":"Eu-hookeria","species":"olfersiana","authorship":{"verbatim":"(Hornsch.) Hampe","normalized":"(Hornsch.) Hampe","authors":["Hornsch.","Hampe"],"originalAuth":{"authors":["Hornsch."]},"combinationAuth":{"authors":["Hampe"]}}}},"words":[{"verbatim":"Eu-hookeria","normalized":"Eu-hookeria","wordType":"GENUS","start":0,"end":11},{"verbatim":"olfersiana","normalized":"olfersiana","wordType":"SPECIES","start":12,"end":22},{"verbatim":"Hornsch.","normalized":"Hornsch.","wordType":"AUTHOR_WORD","start":24,"end":32},{"verbatim":"Hampe","normalized":"Hampe","wordType":"AUTHOR_WORD","start":34,"end":39}],"id":"60824304-4a59-5d99-8af4-97b7f1ae6a20","parserVersion":"test_version"}
```

Name: Le-monniera

Canonical: Le-monniera

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Le-monniera","normalized":"Le-monniera","canonical":{"stemmed":"Le-monniera","simple":"Le-monniera","full":"Le-monniera"},"cardinality":1,"details":{"uninomial":{"uninomial":"Le-monniera"}},"words":[{"verbatim":"Le-monniera","normalized":"Le-monniera","wordType":"UNINOMIAL","start":0,"end":11}],"id":"86091af8-6354-5f2e-94b4-c8a2a3e1fbef","parserVersion":"test_version"}
```

Name: Le-Monniera clitandrifolia (A. Chev.) Lecomte

Canonical: Le-monniera clitandrifolia

Authorship: (A. Chev.) Lecomte

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Le-Monniera clitandrifolia (A. Chev.) Lecomte","normalized":"Le-monniera clitandrifolia (A. Chev.) Lecomte","canonical":{"stemmed":"Le-monniera clitandrifol","simple":"Le-monniera clitandrifolia","full":"Le-monniera clitandrifolia"},"cardinality":2,"authorship":{"verbatim":"(A. Chev.) Lecomte","normalized":"(A. Chev.) Lecomte","authors":["A. Chev.","Lecomte"],"originalAuth":{"authors":["A. Chev."]},"combinationAuth":{"authors":["Lecomte"]}},"details":{"species":{"genus":"Le-monniera","species":"clitandrifolia","authorship":{"verbatim":"(A. Chev.) Lecomte","normalized":"(A. Chev.) Lecomte","authors":["A. Chev.","Lecomte"],"originalAuth":{"authors":["A. Chev."]},"combinationAuth":{"authors":["Lecomte"]}}}},"words":[{"verbatim":"Le-Monniera","normalized":"Le-monniera","wordType":"GENUS","start":0,"end":11},{"verbatim":"clitandrifolia","normalized":"clitandrifolia","wordType":"SPECIES","start":12,"end":26},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"Chev.","normalized":"Chev.","wordType":"AUTHOR_WORD","start":31,"end":36},{"verbatim":"Lecomte","normalized":"Lecomte","wordType":"AUTHOR_WORD","start":38,"end":45}],"id":"b5366fdc-4715-5fb1-8534-890fa67e60ab","parserVersion":"test_version"}
```
Name: Ne-ourbania adendrobium (Rchb.f. ) Fawc. & Rendle

Canonical: Ne-ourbania adendrobium

Authorship: (Rchb. fil.) Fawc. & Rendle

```json
{"parsed":true,"quality":1,"verbatim":"Ne-ourbania adendrobium (Rchb.f. ) Fawc. \u0026 Rendle","normalized":"Ne-ourbania adendrobium (Rchb. fil.) Fawc. \u0026 Rendle","canonical":{"stemmed":"Ne-ourbania adendrobi","simple":"Ne-ourbania adendrobium","full":"Ne-ourbania adendrobium"},"cardinality":2,"authorship":{"verbatim":"(Rchb.f. ) Fawc. \u0026 Rendle","normalized":"(Rchb. fil.) Fawc. \u0026 Rendle","authors":["Rchb. fil.","Fawc.","Rendle"],"originalAuth":{"authors":["Rchb. fil."]},"combinationAuth":{"authors":["Fawc.","Rendle"]}},"details":{"species":{"genus":"Ne-ourbania","species":"adendrobium","authorship":{"verbatim":"(Rchb.f. ) Fawc. \u0026 Rendle","normalized":"(Rchb. fil.) Fawc. \u0026 Rendle","authors":["Rchb. fil.","Fawc.","Rendle"],"originalAuth":{"authors":["Rchb. fil."]},"combinationAuth":{"authors":["Fawc.","Rendle"]}}}},"words":[{"verbatim":"Ne-ourbania","normalized":"Ne-ourbania","wordType":"GENUS","start":0,"end":11},{"verbatim":"adendrobium","normalized":"adendrobium","wordType":"SPECIES","start":12,"end":23},{"verbatim":"Rchb.","normalized":"Rchb.","wordType":"AUTHOR_WORD","start":25,"end":30},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":30,"end":32},{"verbatim":"Fawc.","normalized":"Fawc.","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"Rendle","normalized":"Rendle","wordType":"AUTHOR_WORD","start":43,"end":49}],"id":"51da6d50-05da-50a1-8a68-67811aa38995","parserVersion":"test_version"}
```
<!-- unregistered 2-letter dashed prefixes are not allowed -->

Name: Ph-echinodermata

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Ph-echinodermata","cardinality":0,"id":"776dc8e6-6fda-5682-90e1-f580b29997b6","parserVersion":"test_version"}
```

<!-- Two-dashes genera are rare -->

Name: Prunus-lauro-cerasus

Canonical: Prunus-lauro-cerasus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Prunus-lauro-cerasus","normalized":"Prunus-lauro-cerasus","canonical":{"stemmed":"Prunus-lauro-cerasus","simple":"Prunus-lauro-cerasus","full":"Prunus-lauro-cerasus"},"cardinality":1,"details":{"uninomial":{"uninomial":"Prunus-lauro-cerasus"}},"words":[{"verbatim":"Prunus-lauro-cerasus","normalized":"Prunus-lauro-cerasus","wordType":"UNINOMIAL","start":0,"end":20}],"id":"e23ffe7a-f6ef-5276-a591-93e328213992","parserVersion":"test_version"}
```

Name: Prunus-Lauro-Cerasus

Canonical: Prunus-lauro-cerasus

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Prunus-Lauro-Cerasus","normalized":"Prunus-lauro-cerasus","canonical":{"stemmed":"Prunus-lauro-cerasus","simple":"Prunus-lauro-cerasus","full":"Prunus-lauro-cerasus"},"cardinality":1,"details":{"uninomial":{"uninomial":"Prunus-lauro-cerasus"}},"words":[{"verbatim":"Prunus-Lauro-Cerasus","normalized":"Prunus-lauro-cerasus","wordType":"UNINOMIAL","start":0,"end":20}],"id":"192bf946-803d-53b4-934d-365a8b2798e4","parserVersion":"test_version"}
```

Name: Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan & Gaussen

Canonical: Tsugo-piceo-picea × crassifolia

Authorship: (Flous) Campo-Duplan & Gaussen

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan \u0026 Gaussen","normalized":"Tsugo-piceo-picea × crassifolia (Flous) Campo-Duplan \u0026 Gaussen","canonical":{"stemmed":"Tsugo-piceo-picea crassifol","simple":"Tsugo-piceo-picea crassifolia","full":"Tsugo-piceo-picea × crassifolia"},"cardinality":2,"authorship":{"verbatim":"(Flous) Campo-Duplan \u0026 Gaussen","normalized":"(Flous) Campo-Duplan \u0026 Gaussen","authors":["Flous","Campo-Duplan","Gaussen"],"originalAuth":{"authors":["Flous"]},"combinationAuth":{"authors":["Campo-Duplan","Gaussen"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Tsugo-piceo-picea","species":"crassifolia (Flous) Campo-Duplan \u0026 Gaussen","authorship":{"verbatim":"(Flous) Campo-Duplan \u0026 Gaussen","normalized":"(Flous) Campo-Duplan \u0026 Gaussen","authors":["Flous","Campo-Duplan","Gaussen"],"originalAuth":{"authors":["Flous"]},"combinationAuth":{"authors":["Campo-Duplan","Gaussen"]}}}},"words":[{"verbatim":"Tsugo-piceo-picea","normalized":"Tsugo-piceo-picea","wordType":"GENUS","start":0,"end":17},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":18,"end":19},{"verbatim":"crassifolia","normalized":"crassifolia","wordType":"SPECIES","start":20,"end":31},{"verbatim":"Flous","normalized":"Flous","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"Campo-Duplan","normalized":"Campo-Duplan","wordType":"AUTHOR_WORD","start":40,"end":52},{"verbatim":"Gaussen","normalized":"Gaussen","wordType":"AUTHOR_WORD","start":55,"end":62}],"id":"a00c94bb-566b-5433-a666-d56c1495ca3b","parserVersion":"test_version"}
```

Name: Tsugo-piceo-piceo-picea × crassifolia

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Tsugo-piceo-piceo-picea × crassifolia","cardinality":0,"id":"0ab8c5ed-b224-5c17-9957-298a80cc07be","parserVersion":"test_version"}
```

### Misspeled name

Name: Ambrysus-Stål, 1862

Canonical: Ambrysus-stål

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Apparent genus with capital character after hyphen"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Ambrysus-Stål, 1862","normalized":"Ambrysus-stål","canonical":{"stemmed":"Ambrysus-stål","simple":"Ambrysus-stål","full":"Ambrysus-stål"},"cardinality":1,"tail":", 1862","details":{"uninomial":{"uninomial":"Ambrysus-stål"}},"words":[{"verbatim":"Ambrysus-Stål","normalized":"Ambrysus-stål","wordType":"UNINOMIAL","start":0,"end":13}],"id":"ab9e69c4-9418-5f86-ad51-3bfc87f76016","parserVersion":"test_version"}
```

### A 'basionym' author in parenthesis (basionym is an ICN term)

Name: Zophosis persis (Chatanay, 1914)

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"quality":1,"verbatim":"Zophosis persis (Chatanay, 1914)","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay, 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay, 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"words":[{"verbatim":"Zophosis","normalized":"Zophosis","wordType":"GENUS","start":0,"end":8},{"verbatim":"persis","normalized":"persis","wordType":"SPECIES","start":9,"end":15},{"verbatim":"Chatanay","normalized":"Chatanay","wordType":"AUTHOR_WORD","start":17,"end":25},{"verbatim":"1914","normalized":"1914","wordType":"YEAR","start":27,"end":31}],"id":"b70a2324-4f36-5fef-80b3-5f6ab9c7788d","parserVersion":"test_version"}
```

Name: Zophosis persis (Chatanay 1914)

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"quality":1,"verbatim":"Zophosis persis (Chatanay 1914)","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"words":[{"verbatim":"Zophosis","normalized":"Zophosis","wordType":"GENUS","start":0,"end":8},{"verbatim":"persis","normalized":"persis","wordType":"SPECIES","start":9,"end":15},{"verbatim":"Chatanay","normalized":"Chatanay","wordType":"AUTHOR_WORD","start":17,"end":25},{"verbatim":"1914","normalized":"1914","wordType":"YEAR","start":26,"end":30}],"id":"c6c42947-16b5-5c1c-a889-51392d82a03b","parserVersion":"test_version"}
```

Name: Zophosis persis (Chatanay), 1914

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Misplaced basionym year"}],"verbatim":"Zophosis persis (Chatanay), 1914","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay), 1914","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay), 1914","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"words":[{"verbatim":"Zophosis","normalized":"Zophosis","wordType":"GENUS","start":0,"end":8},{"verbatim":"persis","normalized":"persis","wordType":"SPECIES","start":9,"end":15},{"verbatim":"Chatanay","normalized":"Chatanay","wordType":"AUTHOR_WORD","start":17,"end":25},{"verbatim":"1914","normalized":"1914","wordType":"YEAR","start":28,"end":32}],"id":"3f9b079c-510a-5c0c-9df6-f1660e1b005f","parserVersion":"test_version"}
```

Name: Zophosis quadrilineata (Oliv. )

Canonical: Zophosis quadrilineata

Authorship: (Oliv.)

```json
{"parsed":true,"quality":1,"verbatim":"Zophosis quadrilineata (Oliv. )","normalized":"Zophosis quadrilineata (Oliv.)","canonical":{"stemmed":"Zophosis quadrilineat","simple":"Zophosis quadrilineata","full":"Zophosis quadrilineata"},"cardinality":2,"authorship":{"verbatim":"(Oliv. )","normalized":"(Oliv.)","authors":["Oliv."],"originalAuth":{"authors":["Oliv."]}},"details":{"species":{"genus":"Zophosis","species":"quadrilineata","authorship":{"verbatim":"(Oliv. )","normalized":"(Oliv.)","authors":["Oliv."],"originalAuth":{"authors":["Oliv."]}}}},"words":[{"verbatim":"Zophosis","normalized":"Zophosis","wordType":"GENUS","start":0,"end":8},{"verbatim":"quadrilineata","normalized":"quadrilineata","wordType":"SPECIES","start":9,"end":22},{"verbatim":"Oliv.","normalized":"Oliv.","wordType":"AUTHOR_WORD","start":24,"end":29}],"id":"4d327524-3514-5faf-85fa-e461cbf6c99e","parserVersion":"test_version"}
```

Name: Zophosis quadrilineata (Olivier 1795)

Canonical: Zophosis quadrilineata

Authorship: (Olivier 1795)

```json
{"parsed":true,"quality":1,"verbatim":"Zophosis quadrilineata (Olivier 1795)","normalized":"Zophosis quadrilineata (Olivier 1795)","canonical":{"stemmed":"Zophosis quadrilineat","simple":"Zophosis quadrilineata","full":"Zophosis quadrilineata"},"cardinality":2,"authorship":{"verbatim":"(Olivier 1795)","normalized":"(Olivier 1795)","year":"1795","authors":["Olivier"],"originalAuth":{"authors":["Olivier"],"year":{"year":"1795"}}},"details":{"species":{"genus":"Zophosis","species":"quadrilineata","authorship":{"verbatim":"(Olivier 1795)","normalized":"(Olivier 1795)","year":"1795","authors":["Olivier"],"originalAuth":{"authors":["Olivier"],"year":{"year":"1795"}}}}},"words":[{"verbatim":"Zophosis","normalized":"Zophosis","wordType":"GENUS","start":0,"end":8},{"verbatim":"quadrilineata","normalized":"quadrilineata","wordType":"SPECIES","start":9,"end":22},{"verbatim":"Olivier","normalized":"Olivier","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"1795","normalized":"1795","wordType":"YEAR","start":32,"end":36}],"id":"837cbd42-87a0-573f-9dbf-d089503028ad","parserVersion":"test_version"}
```

### Infrageneric epithets (ICZN)

Name: Hegeter (Hegeter) tenuipunctatus Brullé, 1838

Canonical: Hegeter tenuipunctatus

Authorship: Brullé 1838

```json
{"parsed":true,"quality":1,"verbatim":"Hegeter (Hegeter) tenuipunctatus Brullé, 1838","normalized":"Hegeter (Hegeter) tenuipunctatus Brullé 1838","canonical":{"stemmed":"Hegeter tenuipunctat","simple":"Hegeter tenuipunctatus","full":"Hegeter tenuipunctatus"},"cardinality":2,"authorship":{"verbatim":"Brullé, 1838","normalized":"Brullé 1838","year":"1838","authors":["Brullé"],"originalAuth":{"authors":["Brullé"],"year":{"year":"1838"}}},"details":{"species":{"genus":"Hegeter","subgenus":"Hegeter","species":"tenuipunctatus","authorship":{"verbatim":"Brullé, 1838","normalized":"Brullé 1838","year":"1838","authors":["Brullé"],"originalAuth":{"authors":["Brullé"],"year":{"year":"1838"}}}}},"words":[{"verbatim":"Hegeter","normalized":"Hegeter","wordType":"GENUS","start":0,"end":7},{"verbatim":"Hegeter","normalized":"Hegeter","wordType":"INFRA_GENUS","start":9,"end":16},{"verbatim":"tenuipunctatus","normalized":"tenuipunctatus","wordType":"SPECIES","start":18,"end":32},{"verbatim":"Brullé","normalized":"Brullé","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"1838","normalized":"1838","wordType":"YEAR","start":41,"end":45}],"id":"a5d28cfb-77a8-509c-a7c6-aa598a7cd3d9","parserVersion":"test_version"}
```

Name: Hegeter (Hegeter) intercedens Lindberg H 1950

Canonical: Hegeter intercedens

Authorship: Lindberg H 1950

```json
{"parsed":true,"quality":1,"verbatim":"Hegeter (Hegeter) intercedens Lindberg H 1950","normalized":"Hegeter (Hegeter) intercedens Lindberg H 1950","canonical":{"stemmed":"Hegeter intercedens","simple":"Hegeter intercedens","full":"Hegeter intercedens"},"cardinality":2,"authorship":{"verbatim":"Lindberg H 1950","normalized":"Lindberg H 1950","year":"1950","authors":["Lindberg H"],"originalAuth":{"authors":["Lindberg H"],"year":{"year":"1950"}}},"details":{"species":{"genus":"Hegeter","subgenus":"Hegeter","species":"intercedens","authorship":{"verbatim":"Lindberg H 1950","normalized":"Lindberg H 1950","year":"1950","authors":["Lindberg H"],"originalAuth":{"authors":["Lindberg H"],"year":{"year":"1950"}}}}},"words":[{"verbatim":"Hegeter","normalized":"Hegeter","wordType":"GENUS","start":0,"end":7},{"verbatim":"Hegeter","normalized":"Hegeter","wordType":"INFRA_GENUS","start":9,"end":16},{"verbatim":"intercedens","normalized":"intercedens","wordType":"SPECIES","start":18,"end":29},{"verbatim":"Lindberg","normalized":"Lindberg","wordType":"AUTHOR_WORD","start":30,"end":38},{"verbatim":"H","normalized":"H","wordType":"AUTHOR_WORD","start":39,"end":40},{"verbatim":"1950","normalized":"1950","wordType":"YEAR","start":41,"end":45}],"id":"2486503e-b9fb-547f-a310-944a50d1bce8","parserVersion":"test_version"}
```

<!--
Brachytrypus (B.) grandidieri
-->

Name: Cyprideis (Cyprideis) thessalonike amasyaensis

Canonical: Cyprideis thessalonike amasyaensis

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Cyprideis (Cyprideis) thessalonike amasyaensis","normalized":"Cyprideis (Cyprideis) thessalonike amasyaensis","canonical":{"stemmed":"Cyprideis thessalonik amasyaens","simple":"Cyprideis thessalonike amasyaensis","full":"Cyprideis thessalonike amasyaensis"},"cardinality":3,"details":{"infraspecies":{"genus":"Cyprideis","subgenus":"Cyprideis","species":"thessalonike","infraspecies":[{"value":"amasyaensis"}]}},"words":[{"verbatim":"Cyprideis","normalized":"Cyprideis","wordType":"GENUS","start":0,"end":9},{"verbatim":"Cyprideis","normalized":"Cyprideis","wordType":"INFRA_GENUS","start":11,"end":20},{"verbatim":"thessalonike","normalized":"thessalonike","wordType":"SPECIES","start":22,"end":34},{"verbatim":"amasyaensis","normalized":"amasyaensis","wordType":"INFRASPECIES","start":35,"end":46}],"id":"19945ce1-52ee-5416-af46-0d6f0803b44e","parserVersion":"test_version"}
```

Name: Acanthoderes (acanthoderes) satanas Aurivillius, 1923

Canonical: Acanthoderes satanas

Authorship: Aurivillius 1923

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguity: subgenus or superspecies found"}],"verbatim":"Acanthoderes (acanthoderes) satanas Aurivillius, 1923","normalized":"Acanthoderes satanas Aurivillius 1923","canonical":{"stemmed":"Acanthoderes satan","simple":"Acanthoderes satanas","full":"Acanthoderes satanas"},"cardinality":2,"authorship":{"verbatim":"Aurivillius, 1923","normalized":"Aurivillius 1923","year":"1923","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1923"}}},"details":{"species":{"genus":"Acanthoderes","species":"satanas","authorship":{"verbatim":"Aurivillius, 1923","normalized":"Aurivillius 1923","year":"1923","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1923"}}}}},"words":[{"verbatim":"Acanthoderes","normalized":"Acanthoderes","wordType":"GENUS","start":0,"end":12},{"verbatim":"satanas","normalized":"satanas","wordType":"SPECIES","start":28,"end":35},{"verbatim":"Aurivillius","normalized":"Aurivillius","wordType":"AUTHOR_WORD","start":36,"end":47},{"verbatim":"1923","normalized":"1923","wordType":"YEAR","start":49,"end":53}],"id":"f1082b19-d13f-54a2-95a9-6e342f2a9e6b","parserVersion":"test_version"}
```

<!-- A fake name to illustrate botaincal author instead of subgenus -->
Name: Acanthoderes (Abramov) satanas Aurivillius

Canonical: Acanthoderes satanas

Authorship: Aurivillius

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Acanthoderes (Abramov) satanas Aurivillius","normalized":"Acanthoderes satanas Aurivillius","canonical":{"stemmed":"Acanthoderes satan","simple":"Acanthoderes satanas","full":"Acanthoderes satanas"},"cardinality":2,"authorship":{"verbatim":"Aurivillius","normalized":"Aurivillius","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"]}},"details":{"species":{"genus":"Acanthoderes","species":"satanas","authorship":{"verbatim":"Aurivillius","normalized":"Aurivillius","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"]}}}},"words":[{"verbatim":"Acanthoderes","normalized":"Acanthoderes","wordType":"GENUS","start":0,"end":12},{"verbatim":"satanas","normalized":"satanas","wordType":"SPECIES","start":23,"end":30},{"verbatim":"Aurivillius","normalized":"Aurivillius","wordType":"AUTHOR_WORD","start":31,"end":42}],"id":"8eb2a9be-eb11-537e-8488-eacdb6e2b9e7","parserVersion":"test_version"}
```

### Names with multiple dashes in specific epithet

There are less than 100 of names like this, and only one in CoL with 3 dashes

Name: Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.

Canonical: Athyrium boreo-occidentali-indobharaticola-birianum

Authorship: Fraser-Jenk.

```json
{"parsed":true,"quality":1,"verbatim":"Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.","normalized":"Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.","canonical":{"stemmed":"Athyrium boreo-occidentali-indobharaticola-birian","simple":"Athyrium boreo-occidentali-indobharaticola-birianum","full":"Athyrium boreo-occidentali-indobharaticola-birianum"},"cardinality":2,"authorship":{"verbatim":"Fraser-Jenk.","normalized":"Fraser-Jenk.","authors":["Fraser-Jenk."],"originalAuth":{"authors":["Fraser-Jenk."]}},"details":{"species":{"genus":"Athyrium","species":"boreo-occidentali-indobharaticola-birianum","authorship":{"verbatim":"Fraser-Jenk.","normalized":"Fraser-Jenk.","authors":["Fraser-Jenk."],"originalAuth":{"authors":["Fraser-Jenk."]}}}},"words":[{"verbatim":"Athyrium","normalized":"Athyrium","wordType":"GENUS","start":0,"end":8},{"verbatim":"boreo-occidentali-indobharaticola-birianum","normalized":"boreo-occidentali-indobharaticola-birianum","wordType":"SPECIES","start":9,"end":51},{"verbatim":"Fraser-Jenk.","normalized":"Fraser-Jenk.","wordType":"AUTHOR_WORD","start":52,"end":64}],"id":"6b979652-191f-5d93-ae23-614768ee0be4","parserVersion":"test_version"}
```

Name: Puccinia band-i-amirii Durrieu, 1975

Canonical: Puccinia band-i-amirii

Authorship: Durrieu 1975

```json
{"parsed":true,"quality":1,"verbatim":"Puccinia band-i-amirii Durrieu, 1975","normalized":"Puccinia band-i-amirii Durrieu 1975","canonical":{"stemmed":"Puccinia band-i-amir","simple":"Puccinia band-i-amirii","full":"Puccinia band-i-amirii"},"cardinality":2,"authorship":{"verbatim":"Durrieu, 1975","normalized":"Durrieu 1975","year":"1975","authors":["Durrieu"],"originalAuth":{"authors":["Durrieu"],"year":{"year":"1975"}}},"details":{"species":{"genus":"Puccinia","species":"band-i-amirii","authorship":{"verbatim":"Durrieu, 1975","normalized":"Durrieu 1975","year":"1975","authors":["Durrieu"],"originalAuth":{"authors":["Durrieu"],"year":{"year":"1975"}}}}},"words":[{"verbatim":"Puccinia","normalized":"Puccinia","wordType":"GENUS","start":0,"end":8},{"verbatim":"band-i-amirii","normalized":"band-i-amirii","wordType":"SPECIES","start":9,"end":22},{"verbatim":"Durrieu","normalized":"Durrieu","wordType":"AUTHOR_WORD","start":23,"end":30},{"verbatim":"1975","normalized":"1975","wordType":"YEAR","start":32,"end":36}],"id":"9733e3df-0b03-5e1e-93f9-5931a4e85f12","parserVersion":"test_version"}
```

### Genus with question mark

Name: Ferganoconcha? oblonga

Canonical: Ferganoconcha oblonga

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Uninomial word with question mark"}],"verbatim":"Ferganoconcha? oblonga","normalized":"Ferganoconcha oblonga","canonical":{"stemmed":"Ferganoconcha oblong","simple":"Ferganoconcha oblonga","full":"Ferganoconcha oblonga"},"cardinality":2,"details":{"species":{"genus":"Ferganoconcha","species":"oblonga"}},"words":[{"verbatim":"Ferganoconcha?","normalized":"Ferganoconcha","wordType":"GENUS","start":0,"end":14},{"verbatim":"oblonga","normalized":"oblonga","wordType":"SPECIES","start":15,"end":22}],"id":"487912fd-85c3-556a-a1b1-8fe802e9ccb1","parserVersion":"test_version"}
```

### Epithets with a period character

Name: Macromitrium st.-johnii E. B. Bartram

Canonical: Macromitrium st-johnii

Authorship: E. B. Bartram

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Period character is not allowed in canonical"}],"verbatim":"Macromitrium st.-johnii E. B. Bartram","normalized":"Macromitrium st-johnii E. B. Bartram","canonical":{"stemmed":"Macromitrium st-iohn","simple":"Macromitrium st-johnii","full":"Macromitrium st-johnii"},"cardinality":2,"authorship":{"verbatim":"E. B. Bartram","normalized":"E. B. Bartram","authors":["E. B. Bartram"],"originalAuth":{"authors":["E. B. Bartram"]}},"details":{"species":{"genus":"Macromitrium","species":"st-johnii","authorship":{"verbatim":"E. B. Bartram","normalized":"E. B. Bartram","authors":["E. B. Bartram"],"originalAuth":{"authors":["E. B. Bartram"]}}}},"words":[{"verbatim":"Macromitrium","normalized":"Macromitrium","wordType":"GENUS","start":0,"end":12},{"verbatim":"st.-johnii","normalized":"st-johnii","wordType":"SPECIES","start":13,"end":23},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"B.","normalized":"B.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"Bartram","normalized":"Bartram","wordType":"AUTHOR_WORD","start":30,"end":37}],"id":"219bf25f-d36d-5259-8005-dc3b8a223d0a","parserVersion":"test_version"}
```

### Epithets starting with non-

Name: Peperomia non-alata Trel.

Canonical: Peperomia non-alata

Authorship: Trel.

```json
{"parsed":true,"quality":1,"verbatim":"Peperomia non-alata Trel.","normalized":"Peperomia non-alata Trel.","canonical":{"stemmed":"Peperomia non-alat","simple":"Peperomia non-alata","full":"Peperomia non-alata"},"cardinality":2,"authorship":{"verbatim":"Trel.","normalized":"Trel.","authors":["Trel."],"originalAuth":{"authors":["Trel."]}},"details":{"species":{"genus":"Peperomia","species":"non-alata","authorship":{"verbatim":"Trel.","normalized":"Trel.","authors":["Trel."],"originalAuth":{"authors":["Trel."]}}}},"words":[{"verbatim":"Peperomia","normalized":"Peperomia","wordType":"GENUS","start":0,"end":9},{"verbatim":"non-alata","normalized":"non-alata","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Trel.","normalized":"Trel.","wordType":"AUTHOR_WORD","start":20,"end":25}],"id":"3eb579ac-ab79-5b6a-a63a-eded8f3af476","parserVersion":"test_version"}
```
Name: Hyacinthoides non-scripta (L.) Chouard ex Rothm.

Canonical: Hyacinthoides non-scripta

Authorship: (L.) Chouard ex Rothm.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Hyacinthoides non-scripta (L.) Chouard ex Rothm.","normalized":"Hyacinthoides non-scripta (L.) Chouard ex Rothm.","canonical":{"stemmed":"Hyacinthoides non-script","simple":"Hyacinthoides non-scripta","full":"Hyacinthoides non-scripta"},"cardinality":2,"authorship":{"verbatim":"(L.) Chouard ex Rothm.","normalized":"(L.) Chouard ex Rothm.","authors":["L.","Chouard","Rothm."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Chouard"],"exAuthors":{"authors":["Rothm."]}}},"details":{"species":{"genus":"Hyacinthoides","species":"non-scripta","authorship":{"verbatim":"(L.) Chouard ex Rothm.","normalized":"(L.) Chouard ex Rothm.","authors":["L.","Chouard","Rothm."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Chouard"],"exAuthors":{"authors":["Rothm."]}}}}},"words":[{"verbatim":"Hyacinthoides","normalized":"Hyacinthoides","wordType":"GENUS","start":0,"end":13},{"verbatim":"non-scripta","normalized":"non-scripta","wordType":"SPECIES","start":14,"end":25},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"Chouard","normalized":"Chouard","wordType":"AUTHOR_WORD","start":31,"end":38},{"verbatim":"Rothm.","normalized":"Rothm.","wordType":"AUTHOR_WORD","start":42,"end":48}],"id":"12e44c2c-33f9-5dfb-bc72-6b495577e7b2","parserVersion":"test_version"}
```
Name: Monocelis non-scripta Curini-Galletti, 2014

Canonical: Monocelis non-scripta

Authorship: Curini-Galletti 2014

```json
{"parsed":true,"quality":1,"verbatim":"Monocelis non-scripta Curini-Galletti, 2014","normalized":"Monocelis non-scripta Curini-Galletti 2014","canonical":{"stemmed":"Monocelis non-script","simple":"Monocelis non-scripta","full":"Monocelis non-scripta"},"cardinality":2,"authorship":{"verbatim":"Curini-Galletti, 2014","normalized":"Curini-Galletti 2014","year":"2014","authors":["Curini-Galletti"],"originalAuth":{"authors":["Curini-Galletti"],"year":{"year":"2014"}}},"details":{"species":{"genus":"Monocelis","species":"non-scripta","authorship":{"verbatim":"Curini-Galletti, 2014","normalized":"Curini-Galletti 2014","year":"2014","authors":["Curini-Galletti"],"originalAuth":{"authors":["Curini-Galletti"],"year":{"year":"2014"}}}}},"words":[{"verbatim":"Monocelis","normalized":"Monocelis","wordType":"GENUS","start":0,"end":9},{"verbatim":"non-scripta","normalized":"non-scripta","wordType":"SPECIES","start":10,"end":21},{"verbatim":"Curini-Galletti","normalized":"Curini-Galletti","wordType":"AUTHOR_WORD","start":22,"end":37},{"verbatim":"2014","normalized":"2014","wordType":"YEAR","start":39,"end":43}],"id":"26be3019-a49f-5299-9c86-6363abe6e982","parserVersion":"test_version"}
```

### Epithets starting with authors' prefixes (de, di, la, von etc.)

<!-- There is a danger that such epithets will be interpreted as authors -->

Name: Aspicilia desertorum desertorum

Canonical: Aspicilia desertorum desertorum

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Aspicilia desertorum desertorum","normalized":"Aspicilia desertorum desertorum","canonical":{"stemmed":"Aspicilia desertor desertor","simple":"Aspicilia desertorum desertorum","full":"Aspicilia desertorum desertorum"},"cardinality":3,"details":{"infraspecies":{"genus":"Aspicilia","species":"desertorum","infraspecies":[{"value":"desertorum"}]}},"words":[{"verbatim":"Aspicilia","normalized":"Aspicilia","wordType":"GENUS","start":0,"end":9},{"verbatim":"desertorum","normalized":"desertorum","wordType":"SPECIES","start":10,"end":20},{"verbatim":"desertorum","normalized":"desertorum","wordType":"INFRASPECIES","start":21,"end":31}],"id":"06de3555-3226-5e05-930e-6706044c1f7a","parserVersion":"test_version"}
```

Name: Theope thestias discus

Canonical: Theope thestias discus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Theope thestias discus","normalized":"Theope thestias discus","canonical":{"stemmed":"Theope thesti disc","simple":"Theope thestias discus","full":"Theope thestias discus"},"cardinality":3,"details":{"infraspecies":{"genus":"Theope","species":"thestias","infraspecies":[{"value":"discus"}]}},"words":[{"verbatim":"Theope","normalized":"Theope","wordType":"GENUS","start":0,"end":6},{"verbatim":"thestias","normalized":"thestias","wordType":"SPECIES","start":7,"end":15},{"verbatim":"discus","normalized":"discus","wordType":"INFRASPECIES","start":16,"end":22}],"id":"a254509a-11e4-52f3-bd57-2271d9e1d99b","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus (Dejean, 1831)

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"quality":1,"verbatim":"Ocydromus dalmatinus dalmatinus (Dejean, 1831)","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"(Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraspecies":{"genus":"Ocydromus","species":"dalmatinus","infraspecies":[{"value":"dalmatinus","authorship":{"verbatim":"(Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"words":[{"verbatim":"Ocydromus","normalized":"Ocydromus","wordType":"GENUS","start":0,"end":9},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Dejean","normalized":"Dejean","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":41,"end":45}],"id":"5701cc12-ec23-5015-b426-3d065c94ea0a","parserVersion":"test_version"}
```

Name: Rhipidia gracilirama lassula

Canonical: Rhipidia gracilirama lassula

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Rhipidia gracilirama lassula","normalized":"Rhipidia gracilirama lassula","canonical":{"stemmed":"Rhipidia graciliram lassul","simple":"Rhipidia gracilirama lassula","full":"Rhipidia gracilirama lassula"},"cardinality":3,"details":{"infraspecies":{"genus":"Rhipidia","species":"gracilirama","infraspecies":[{"value":"lassula"}]}},"words":[{"verbatim":"Rhipidia","normalized":"Rhipidia","wordType":"GENUS","start":0,"end":8},{"verbatim":"gracilirama","normalized":"gracilirama","wordType":"SPECIES","start":9,"end":20},{"verbatim":"lassula","normalized":"lassula","wordType":"INFRASPECIES","start":21,"end":28}],"id":"0b40c395-7466-5879-9b16-9a31d38d21a0","parserVersion":"test_version"}
```

### Authorship missing one parenthesis

Name: Ocydromus dalmatinus dalmatinus Dejean, 1831)

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus Dejean, 1831)","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraspecies":{"genus":"Ocydromus","species":"dalmatinus","infraspecies":[{"value":"dalmatinus","authorship":{"verbatim":"Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"words":[{"verbatim":"Ocydromus","normalized":"Ocydromus","wordType":"GENUS","start":0,"end":9},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Dejean","normalized":"Dejean","wordType":"AUTHOR_WORD","start":32,"end":38},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":40,"end":44}],"id":"5de70fe3-959a-5555-afdb-3ab85b91f1d7","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus Dejean, 1831 )

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus Dejean, 1831 )","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"Dejean, 1831 )","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraspecies":{"genus":"Ocydromus","species":"dalmatinus","infraspecies":[{"value":"dalmatinus","authorship":{"verbatim":"Dejean, 1831 )","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"words":[{"verbatim":"Ocydromus","normalized":"Ocydromus","wordType":"GENUS","start":0,"end":9},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Dejean","normalized":"Dejean","wordType":"AUTHOR_WORD","start":32,"end":38},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":40,"end":44}],"id":"88dcc885-3360-5234-9620-371c2ebb636c","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus ( Dejean, 1831 Mill.

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831) Mill.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus ( Dejean, 1831 Mill.","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831) Mill.","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"( Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}},"details":{"infraspecies":{"genus":"Ocydromus","species":"dalmatinus","infraspecies":[{"value":"dalmatinus","authorship":{"verbatim":"( Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}}}]}},"words":[{"verbatim":"Ocydromus","normalized":"Ocydromus","wordType":"GENUS","start":0,"end":9},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Dejean","normalized":"Dejean","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":42,"end":46},{"verbatim":"Mill.","normalized":"Mill.","wordType":"AUTHOR_WORD","start":47,"end":52}],"id":"0e8758a1-2567-543b-bafd-c8f9c81e2f08","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus (Dejean, 1831 Mill.

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831) Mill.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus (Dejean, 1831 Mill.","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831) Mill.","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"(Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}},"details":{"infraspecies":{"genus":"Ocydromus","species":"dalmatinus","infraspecies":[{"value":"dalmatinus","authorship":{"verbatim":"(Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}}}]}},"words":[{"verbatim":"Ocydromus","normalized":"Ocydromus","wordType":"GENUS","start":0,"end":9},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"dalmatinus","normalized":"dalmatinus","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Dejean","normalized":"Dejean","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"1831","normalized":"1831","wordType":"YEAR","start":41,"end":45},{"verbatim":"Mill.","normalized":"Mill.","wordType":"AUTHOR_WORD","start":46,"end":51}],"id":"b3c856b3-16a7-5dfc-abfd-3bba539b634f","parserVersion":"test_version"}
```

### Unknown authorship

Name: Saccharomyces drosophilae anon.

Canonical: Saccharomyces drosophilae

Authorship: anon.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Saccharomyces drosophilae anon.","normalized":"Saccharomyces drosophilae anon.","canonical":{"stemmed":"Saccharomyces drosophil","simple":"Saccharomyces drosophilae","full":"Saccharomyces drosophilae"},"cardinality":2,"authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Saccharomyces","species":"drosophilae","authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}}}},"words":[{"verbatim":"Saccharomyces","normalized":"Saccharomyces","wordType":"GENUS","start":0,"end":13},{"verbatim":"drosophilae","normalized":"drosophilae","wordType":"SPECIES","start":14,"end":25},{"verbatim":"anon.","normalized":"anon.","wordType":"AUTHOR_WORD","start":26,"end":31}],"id":"45e537d2-6833-5429-a58c-178fe37fc3f5","parserVersion":"test_version"}
```

Name: Physalospora rubiginosa (Fr.) anon.

Canonical: Physalospora rubiginosa

Authorship: (Fr.) anon.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Physalospora rubiginosa (Fr.) anon.","normalized":"Physalospora rubiginosa (Fr.) anon.","canonical":{"stemmed":"Physalospora rubiginos","simple":"Physalospora rubiginosa","full":"Physalospora rubiginosa"},"cardinality":2,"authorship":{"verbatim":"(Fr.) anon.","normalized":"(Fr.) anon.","authors":["Fr.","anon."],"originalAuth":{"authors":["Fr."]},"combinationAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Physalospora","species":"rubiginosa","authorship":{"verbatim":"(Fr.) anon.","normalized":"(Fr.) anon.","authors":["Fr.","anon."],"originalAuth":{"authors":["Fr."]},"combinationAuth":{"authors":["anon."]}}}},"words":[{"verbatim":"Physalospora","normalized":"Physalospora","wordType":"GENUS","start":0,"end":12},{"verbatim":"rubiginosa","normalized":"rubiginosa","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":25,"end":28},{"verbatim":"anon.","normalized":"anon.","wordType":"AUTHOR_WORD","start":30,"end":35}],"id":"85151e19-ab25-5ba5-8a19-47a5859c41bb","parserVersion":"test_version"}
```

Name: Tragacantha leporina (?) Kuntze

Canonical: Tragacantha leporina

Authorship: (anon.) Kuntze

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Author as a question mark"},{"quality":3,"warning":"Author is too short"},{"quality":2,"warning":"Author is unknown"}],"verbatim":"Tragacantha leporina (?) Kuntze","normalized":"Tragacantha leporina (anon.) Kuntze","canonical":{"stemmed":"Tragacantha leporin","simple":"Tragacantha leporina","full":"Tragacantha leporina"},"cardinality":2,"authorship":{"verbatim":"(?) Kuntze","normalized":"(anon.) Kuntze","authors":["anon.","Kuntze"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Kuntze"]}},"details":{"species":{"genus":"Tragacantha","species":"leporina","authorship":{"verbatim":"(?) Kuntze","normalized":"(anon.) Kuntze","authors":["anon.","Kuntze"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Kuntze"]}}}},"words":[{"verbatim":"Tragacantha","normalized":"Tragacantha","wordType":"GENUS","start":0,"end":11},{"verbatim":"leporina","normalized":"leporina","wordType":"SPECIES","start":12,"end":20},{"verbatim":"?","normalized":"anon.","wordType":"AUTHOR_WORD","start":22,"end":23},{"verbatim":"Kuntze","normalized":"Kuntze","wordType":"AUTHOR_WORD","start":25,"end":31}],"id":"af91bdc5-b6d3-5841-9a85-174c0afe6c1b","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (auct.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship: (anon.) Baker

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Lachenalia tricolor var. nelsonii (auct.) Baker","normalized":"Lachenalia tricolor var. nelsonii (anon.) Baker","canonical":{"stemmed":"Lachenalia tricolor nelson","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"authorship":{"verbatim":"(auct.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}},"details":{"infraspecies":{"genus":"Lachenalia","species":"tricolor","infraspecies":[{"value":"nelsonii","rank":"var.","authorship":{"verbatim":"(auct.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}}}]}},"words":[{"verbatim":"Lachenalia","normalized":"Lachenalia","wordType":"GENUS","start":0,"end":10},{"verbatim":"tricolor","normalized":"tricolor","wordType":"SPECIES","start":11,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"nelsonii","normalized":"nelsonii","wordType":"INFRASPECIES","start":25,"end":33},{"verbatim":"auct.","normalized":"anon.","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"Baker","normalized":"Baker","wordType":"AUTHOR_WORD","start":42,"end":47}],"id":"f8d5d993-3d39-550f-bb7b-68f5b6e906df","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (anon.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship: (anon.) Baker

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Lachenalia tricolor var. nelsonii (anon.) Baker","normalized":"Lachenalia tricolor var. nelsonii (anon.) Baker","canonical":{"stemmed":"Lachenalia tricolor nelson","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"authorship":{"verbatim":"(anon.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}},"details":{"infraspecies":{"genus":"Lachenalia","species":"tricolor","infraspecies":[{"value":"nelsonii","rank":"var.","authorship":{"verbatim":"(anon.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}}}]}},"words":[{"verbatim":"Lachenalia","normalized":"Lachenalia","wordType":"GENUS","start":0,"end":10},{"verbatim":"tricolor","normalized":"tricolor","wordType":"SPECIES","start":11,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"nelsonii","normalized":"nelsonii","wordType":"INFRASPECIES","start":25,"end":33},{"verbatim":"anon.","normalized":"anon.","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"Baker","normalized":"Baker","wordType":"AUTHOR_WORD","start":42,"end":47}],"id":"4cc8e603-13fb-551f-a637-04378f3321c2","parserVersion":"test_version"}
```

Name: Puya acris anon.

Canonical: Puya acris

Authorship: anon.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Puya acris anon.","normalized":"Puya acris anon.","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Puya","species":"acris","authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}}}},"words":[{"verbatim":"Puya","normalized":"Puya","wordType":"GENUS","start":0,"end":4},{"verbatim":"acris","normalized":"acris","wordType":"SPECIES","start":5,"end":10},{"verbatim":"anon.","normalized":"anon.","wordType":"AUTHOR_WORD","start":11,"end":16}],"id":"2b5243d3-e8a7-5e6c-a2c1-beb2ee5c3020","parserVersion":"test_version"}
```

### Treating apud (with)

Name: Pseudocercospora dendrobii Goh apud W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh apud W. H. Hsieh 1990

```json
{"parsed":true,"quality":1,"verbatim":"Pseudocercospora dendrobii Goh apud W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh apud W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh apud W.H. Hsieh 1990","normalized":"Goh apud W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh apud W.H. Hsieh 1990","normalized":"Goh apud W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":17,"end":26},{"verbatim":"Goh","normalized":"Goh","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"Hsieh","normalized":"Hsieh","wordType":"AUTHOR_WORD","start":41,"end":46},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":47,"end":51}],"id":"4dee6fc8-3be1-520c-9937-5a7342a17241","parserVersion":"test_version"}
```

### Names with ex authors (we follow ICZN convention)

Name: Amathia tricornis Busk ms in Chimonides, 1987

Canonical: Amathia tricornis

Authorship: Busk ex Chimonides 1987

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Amathia tricornis Busk ms in Chimonides, 1987","normalized":"Amathia tricornis Busk ex Chimonides 1987","canonical":{"stemmed":"Amathia tricorn","simple":"Amathia tricornis","full":"Amathia tricornis"},"cardinality":2,"authorship":{"verbatim":"Busk ms in Chimonides, 1987","normalized":"Busk ex Chimonides 1987","year":"1987","authors":["Busk","Chimonides"],"originalAuth":{"authors":["Busk"],"exAuthors":{"authors":["Chimonides"],"year":{"year":"1987"}}}},"details":{"species":{"genus":"Amathia","species":"tricornis","authorship":{"verbatim":"Busk ms in Chimonides, 1987","normalized":"Busk ex Chimonides 1987","year":"1987","authors":["Busk","Chimonides"],"originalAuth":{"authors":["Busk"],"exAuthors":{"authors":["Chimonides"],"year":{"year":"1987"}}}}}},"words":[{"verbatim":"Amathia","normalized":"Amathia","wordType":"GENUS","start":0,"end":7},{"verbatim":"tricornis","normalized":"tricornis","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Busk","normalized":"Busk","wordType":"AUTHOR_WORD","start":18,"end":22},{"verbatim":"Chimonides","normalized":"Chimonides","wordType":"AUTHOR_WORD","start":29,"end":39},{"verbatim":"1987","normalized":"1987","wordType":"YEAR","start":41,"end":45}],"id":"fb349d1f-30f2-5e4a-a454-68159d362d58","parserVersion":"test_version"}
```

Name: Pisania billehousti Souverbie, in Souverbie and Montrouzier, 1864

Canonical: Pisania billehousti

Authorship: Souverbie ex Souverbie & Montrouzier 1864

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Pisania billehousti Souverbie, in Souverbie and Montrouzier, 1864","normalized":"Pisania billehousti Souverbie ex Souverbie \u0026 Montrouzier 1864","canonical":{"stemmed":"Pisania billehoust","simple":"Pisania billehousti","full":"Pisania billehousti"},"cardinality":2,"authorship":{"verbatim":"Souverbie, in Souverbie and Montrouzier, 1864","normalized":"Souverbie ex Souverbie \u0026 Montrouzier 1864","year":"1864","authors":["Souverbie","Montrouzier"],"originalAuth":{"authors":["Souverbie"],"exAuthors":{"authors":["Souverbie","Montrouzier"],"year":{"year":"1864"}}}},"details":{"species":{"genus":"Pisania","species":"billehousti","authorship":{"verbatim":"Souverbie, in Souverbie and Montrouzier, 1864","normalized":"Souverbie ex Souverbie \u0026 Montrouzier 1864","year":"1864","authors":["Souverbie","Montrouzier"],"originalAuth":{"authors":["Souverbie"],"exAuthors":{"authors":["Souverbie","Montrouzier"],"year":{"year":"1864"}}}}}},"words":[{"verbatim":"Pisania","normalized":"Pisania","wordType":"GENUS","start":0,"end":7},{"verbatim":"billehousti","normalized":"billehousti","wordType":"SPECIES","start":8,"end":19},{"verbatim":"Souverbie","normalized":"Souverbie","wordType":"AUTHOR_WORD","start":20,"end":29},{"verbatim":"Souverbie","normalized":"Souverbie","wordType":"AUTHOR_WORD","start":34,"end":43},{"verbatim":"Montrouzier","normalized":"Montrouzier","wordType":"AUTHOR_WORD","start":48,"end":59},{"verbatim":"1864","normalized":"1864","wordType":"YEAR","start":61,"end":65}],"id":"a84244fa-ee95-5b97-a339-2de33cef70de","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Nyl. ex Banker) R.C. Harris

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl. ex Banker) R. C. Harris

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Arthopyrenia hyalospora (Nyl. ex Banker) R.C. Harris","normalized":"Arthopyrenia hyalospora (Nyl. ex Banker) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl. ex Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","Banker","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl. ex Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","Banker","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}}}},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"Harris","normalized":"Harris","wordType":"AUTHOR_WORD","start":46,"end":52}],"id":"ab3998af-53dc-53fd-af8b-fab94dacbcbc","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Nyl. ex. Banker) R.C. Harris

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl. ex Banker) R. C. Harris

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Arthopyrenia hyalospora (Nyl. ex. Banker) R.C. Harris","normalized":"Arthopyrenia hyalospora (Nyl. ex Banker) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl. ex. Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","Banker","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl. ex. Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","Banker","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}}}},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":42,"end":44},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":44,"end":46},{"verbatim":"Harris","normalized":"Harris","wordType":"AUTHOR_WORD","start":47,"end":53}],"id":"166fd290-17f5-5b9f-8f72-86830a9bd152","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora Nyl. ex Banker

Canonical: Arthopyrenia hyalospora

Authorship: Nyl. ex Banker

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Arthopyrenia hyalospora Nyl. ex Banker","normalized":"Arthopyrenia hyalospora Nyl. ex Banker","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"Nyl. ex Banker","normalized":"Nyl. ex Banker","authors":["Nyl.","Banker"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"Nyl. ex Banker","normalized":"Nyl. ex Banker","authors":["Nyl.","Banker"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}}}},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":24,"end":28},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":32,"end":38}],"id":"7744aea4-d071-593a-82bc-059788724d81","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora Nyl. ex. Banker

Canonical: Arthopyrenia hyalospora

Authorship: Nyl. ex Banker

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Arthopyrenia hyalospora Nyl. ex. Banker","normalized":"Arthopyrenia hyalospora Nyl. ex Banker","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"Nyl. ex. Banker","normalized":"Nyl. ex Banker","authors":["Nyl.","Banker"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"Nyl. ex. Banker","normalized":"Nyl. ex Banker","authors":["Nyl.","Banker"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}}}},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":24,"end":28},{"verbatim":"Banker","normalized":"Banker","wordType":"AUTHOR_WORD","start":33,"end":39}],"id":"e9097ad7-7bb6-57a2-bad4-52822e5fd655","parserVersion":"test_version"}
```

Name: Glomopsis lonicerae Peck ex C.J. Gould 1945

Canonical: Glomopsis lonicerae

Authorship: Peck ex C. J. Gould 1945

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Glomopsis lonicerae Peck ex C.J. Gould 1945","normalized":"Glomopsis lonicerae Peck ex C. J. Gould 1945","canonical":{"stemmed":"Glomopsis lonicer","simple":"Glomopsis lonicerae","full":"Glomopsis lonicerae"},"cardinality":2,"authorship":{"verbatim":"Peck ex C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","year":"1945","authors":["Peck","C. J. Gould"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}},"details":{"species":{"genus":"Glomopsis","species":"lonicerae","authorship":{"verbatim":"Peck ex C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","year":"1945","authors":["Peck","C. J. Gould"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}}}},"words":[{"verbatim":"Glomopsis","normalized":"Glomopsis","wordType":"GENUS","start":0,"end":9},{"verbatim":"lonicerae","normalized":"lonicerae","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Peck","normalized":"Peck","wordType":"AUTHOR_WORD","start":20,"end":24},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":30,"end":32},{"verbatim":"Gould","normalized":"Gould","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"1945","normalized":"1945","wordType":"YEAR","start":39,"end":43}],"id":"422687ca-7f4b-5720-8d99-88695f765530","parserVersion":"test_version"}
```

Name: Glomopsis lonicerae Peck ex. C.J. Gould 1945

Canonical: Glomopsis lonicerae

Authorship: Peck ex C. J. Gould 1945

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Glomopsis lonicerae Peck ex. C.J. Gould 1945","normalized":"Glomopsis lonicerae Peck ex C. J. Gould 1945","canonical":{"stemmed":"Glomopsis lonicer","simple":"Glomopsis lonicerae","full":"Glomopsis lonicerae"},"cardinality":2,"authorship":{"verbatim":"Peck ex. C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","year":"1945","authors":["Peck","C. J. Gould"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}},"details":{"species":{"genus":"Glomopsis","species":"lonicerae","authorship":{"verbatim":"Peck ex. C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","year":"1945","authors":["Peck","C. J. Gould"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}}}},"words":[{"verbatim":"Glomopsis","normalized":"Glomopsis","wordType":"GENUS","start":0,"end":9},{"verbatim":"lonicerae","normalized":"lonicerae","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Peck","normalized":"Peck","wordType":"AUTHOR_WORD","start":20,"end":24},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":31,"end":33},{"verbatim":"Gould","normalized":"Gould","wordType":"AUTHOR_WORD","start":34,"end":39},{"verbatim":"1945","normalized":"1945","wordType":"YEAR","start":40,"end":44}],"id":"a9cdd33f-990c-59b6-abc2-de9698d2f085","parserVersion":"test_version"}
```

Name: Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979

Canonical: Acanthobasidium delicatum

Authorship: (Wakef.) Oberw. ex Jülich 1979

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","normalized":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","canonical":{"stemmed":"Acanthobasidium delicat","simple":"Acanthobasidium delicatum","full":"Acanthobasidium delicatum"},"cardinality":2,"authorship":{"verbatim":"(Wakef.) Oberw. ex Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw.","Jülich"],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}},"details":{"species":{"genus":"Acanthobasidium","species":"delicatum","authorship":{"verbatim":"(Wakef.) Oberw. ex Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw.","Jülich"],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}}}},"words":[{"verbatim":"Acanthobasidium","normalized":"Acanthobasidium","wordType":"GENUS","start":0,"end":15},{"verbatim":"delicatum","normalized":"delicatum","wordType":"SPECIES","start":16,"end":25},{"verbatim":"Wakef.","normalized":"Wakef.","wordType":"AUTHOR_WORD","start":27,"end":33},{"verbatim":"Oberw.","normalized":"Oberw.","wordType":"AUTHOR_WORD","start":35,"end":41},{"verbatim":"Jülich","normalized":"Jülich","wordType":"AUTHOR_WORD","start":45,"end":51},{"verbatim":"1979","normalized":"1979","wordType":"YEAR","start":52,"end":56}],"id":"ed0841f3-d063-5341-a1b6-feafe6ffb70d","parserVersion":"test_version"}
```

Name: Acanthobasidium delicatum (Wakef.) Oberw. ex. Jülich 1979

Canonical: Acanthobasidium delicatum

Authorship: (Wakef.) Oberw. ex Jülich 1979

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Acanthobasidium delicatum (Wakef.) Oberw. ex. Jülich 1979","normalized":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","canonical":{"stemmed":"Acanthobasidium delicat","simple":"Acanthobasidium delicatum","full":"Acanthobasidium delicatum"},"cardinality":2,"authorship":{"verbatim":"(Wakef.) Oberw. ex. Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw.","Jülich"],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}},"details":{"species":{"genus":"Acanthobasidium","species":"delicatum","authorship":{"verbatim":"(Wakef.) Oberw. ex. Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw.","Jülich"],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}}}},"words":[{"verbatim":"Acanthobasidium","normalized":"Acanthobasidium","wordType":"GENUS","start":0,"end":15},{"verbatim":"delicatum","normalized":"delicatum","wordType":"SPECIES","start":16,"end":25},{"verbatim":"Wakef.","normalized":"Wakef.","wordType":"AUTHOR_WORD","start":27,"end":33},{"verbatim":"Oberw.","normalized":"Oberw.","wordType":"AUTHOR_WORD","start":35,"end":41},{"verbatim":"Jülich","normalized":"Jülich","wordType":"AUTHOR_WORD","start":46,"end":52},{"verbatim":"1979","normalized":"1979","wordType":"YEAR","start":53,"end":57}],"id":"96adf61e-3316-5a08-afac-2b7cd0430eee","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. ex Duby) Johanson ex Oudem. 1897

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","canonical":{"stemmed":"Mycosphaerella eryng","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Duby","Johanson","Oudem."],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}},"details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Duby","Johanson","Oudem."],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}}}},"words":[{"verbatim":"Mycosphaerella","normalized":"Mycosphaerella","wordType":"GENUS","start":0,"end":14},{"verbatim":"eryngii","normalized":"eryngii","wordType":"SPECIES","start":15,"end":22},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":24,"end":27},{"verbatim":"Duby","normalized":"Duby","wordType":"AUTHOR_WORD","start":31,"end":35},{"verbatim":"Johanson","normalized":"Johanson","wordType":"AUTHOR_WORD","start":37,"end":45},{"verbatim":"Oudem.","normalized":"Oudem.","wordType":"AUTHOR_WORD","start":49,"end":55},{"verbatim":"1897","normalized":"1897","wordType":"YEAR","start":56,"end":60}],"id":"8ca3d249-fe7d-5a10-af03-f21c413e3503","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. ex. Duby) Johanson ex. Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. ex Duby) Johanson ex Oudem. 1897

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Mycosphaerella eryngii (Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","canonical":{"stemmed":"Mycosphaerella eryng","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Duby","Johanson","Oudem."],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}},"details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Duby","Johanson","Oudem."],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}}}},"words":[{"verbatim":"Mycosphaerella","normalized":"Mycosphaerella","wordType":"GENUS","start":0,"end":14},{"verbatim":"eryngii","normalized":"eryngii","wordType":"SPECIES","start":15,"end":22},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":24,"end":27},{"verbatim":"Duby","normalized":"Duby","wordType":"AUTHOR_WORD","start":32,"end":36},{"verbatim":"Johanson","normalized":"Johanson","wordType":"AUTHOR_WORD","start":38,"end":46},{"verbatim":"Oudem.","normalized":"Oudem.","wordType":"AUTHOR_WORD","start":51,"end":57},{"verbatim":"1897","normalized":"1897","wordType":"YEAR","start":58,"end":62}],"id":"201b50d3-507b-56d1-99b4-50ab9120bca9","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. Duby) ex Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. Duby)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Mycosphaerella eryngii (Fr. Duby) ex Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. Duby)","canonical":{"stemmed":"Mycosphaerella eryng","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. Duby)","normalized":"(Fr. Duby)","authors":["Fr. Duby"],"originalAuth":{"authors":["Fr. Duby"]}},"tail":" ex Oudem. 1897","details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. Duby)","normalized":"(Fr. Duby)","authors":["Fr. Duby"],"originalAuth":{"authors":["Fr. Duby"]}}}},"words":[{"verbatim":"Mycosphaerella","normalized":"Mycosphaerella","wordType":"GENUS","start":0,"end":14},{"verbatim":"eryngii","normalized":"eryngii","wordType":"SPECIES","start":15,"end":22},{"verbatim":"Fr.","normalized":"Fr.","wordType":"AUTHOR_WORD","start":24,"end":27},{"verbatim":"Duby","normalized":"Duby","wordType":"AUTHOR_WORD","start":28,"end":32}],"id":"e5a49f2e-c7a2-5ebf-9349-8a36a410ec77","parserVersion":"test_version"}
```

### Empty spaces
Name:     Asplenium       X inexpectatum(E. L. Braun ex Friesner      )Morton

Canonical: Asplenium × inexpectatum

Authorship: (E. L. Braun ex Friesner) Morton

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"    Asplenium       X inexpectatum(E. L. Braun ex Friesner      )Morton","normalized":"Asplenium × inexpectatum (E. L. Braun ex Friesner) Morton","canonical":{"stemmed":"Asplenium inexpectat","simple":"Asplenium inexpectatum","full":"Asplenium × inexpectatum"},"cardinality":2,"authorship":{"verbatim":"(E. L. Braun ex Friesner      )Morton","normalized":"(E. L. Braun ex Friesner) Morton","authors":["E. L. Braun","Friesner","Morton"],"originalAuth":{"authors":["E. L. Braun"],"exAuthors":{"authors":["Friesner"]}},"combinationAuth":{"authors":["Morton"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Asplenium","species":"inexpectatum (E. L. Braun ex Friesner) Morton","authorship":{"verbatim":"(E. L. Braun ex Friesner      )Morton","normalized":"(E. L. Braun ex Friesner) Morton","authors":["E. L. Braun","Friesner","Morton"],"originalAuth":{"authors":["E. L. Braun"],"exAuthors":{"authors":["Friesner"]}},"combinationAuth":{"authors":["Morton"]}}}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":4,"end":13},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":20,"end":21},{"verbatim":"inexpectatum","normalized":"inexpectatum","wordType":"SPECIES","start":22,"end":34},{"verbatim":"E.","normalized":"E.","wordType":"AUTHOR_WORD","start":35,"end":37},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"Braun","normalized":"Braun","wordType":"AUTHOR_WORD","start":41,"end":46},{"verbatim":"Friesner","normalized":"Friesner","wordType":"AUTHOR_WORD","start":50,"end":58},{"verbatim":"Morton","normalized":"Morton","wordType":"AUTHOR_WORD","start":65,"end":71}],"id":"a2c7a7ee-51c9-5f3a-8117-bffd799b39f4","parserVersion":"test_version"}
```

### Names with a dash

Name: Drosophila obscura-x Burla, 1951

Canonical: Drosophila obscura-x

Authorship: Burla 1951

```json
{"parsed":true,"quality":1,"verbatim":"Drosophila obscura-x Burla, 1951","normalized":"Drosophila obscura-x Burla 1951","canonical":{"stemmed":"Drosophila obscura-x","simple":"Drosophila obscura-x","full":"Drosophila obscura-x"},"cardinality":2,"authorship":{"verbatim":"Burla, 1951","normalized":"Burla 1951","year":"1951","authors":["Burla"],"originalAuth":{"authors":["Burla"],"year":{"year":"1951"}}},"details":{"species":{"genus":"Drosophila","species":"obscura-x","authorship":{"verbatim":"Burla, 1951","normalized":"Burla 1951","year":"1951","authors":["Burla"],"originalAuth":{"authors":["Burla"],"year":{"year":"1951"}}}}},"words":[{"verbatim":"Drosophila","normalized":"Drosophila","wordType":"GENUS","start":0,"end":10},{"verbatim":"obscura-x","normalized":"obscura-x","wordType":"SPECIES","start":11,"end":20},{"verbatim":"Burla","normalized":"Burla","wordType":"AUTHOR_WORD","start":21,"end":26},{"verbatim":"1951","normalized":"1951","wordType":"YEAR","start":28,"end":32}],"id":"778f9878-8e47-5c7a-a464-33805b6bf173","parserVersion":"test_version"}
```

Name: Sanogasta x-signata (Keyserling,1891)

Canonical: Sanogasta x-signata

Authorship: (Keyserling 1891)

```json
{"parsed":true,"quality":1,"verbatim":"Sanogasta x-signata (Keyserling,1891)","normalized":"Sanogasta x-signata (Keyserling 1891)","canonical":{"stemmed":"Sanogasta x-signat","simple":"Sanogasta x-signata","full":"Sanogasta x-signata"},"cardinality":2,"authorship":{"verbatim":"(Keyserling,1891)","normalized":"(Keyserling 1891)","year":"1891","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1891"}}},"details":{"species":{"genus":"Sanogasta","species":"x-signata","authorship":{"verbatim":"(Keyserling,1891)","normalized":"(Keyserling 1891)","year":"1891","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1891"}}}}},"words":[{"verbatim":"Sanogasta","normalized":"Sanogasta","wordType":"GENUS","start":0,"end":9},{"verbatim":"x-signata","normalized":"x-signata","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Keyserling","normalized":"Keyserling","wordType":"AUTHOR_WORD","start":21,"end":31},{"verbatim":"1891","normalized":"1891","wordType":"YEAR","start":32,"end":36}],"id":"ffe6799d-387a-53d8-8fdd-be73cdc681b8","parserVersion":"test_version"}
```

Name: Aedes w-albus (Theobald, 1905)

Canonical: Aedes w-albus

Authorship: (Theobald 1905)

```json
{"parsed":true,"quality":1,"verbatim":"Aedes w-albus (Theobald, 1905)","normalized":"Aedes w-albus (Theobald 1905)","canonical":{"stemmed":"Aedes w-alb","simple":"Aedes w-albus","full":"Aedes w-albus"},"cardinality":2,"authorship":{"verbatim":"(Theobald, 1905)","normalized":"(Theobald 1905)","year":"1905","authors":["Theobald"],"originalAuth":{"authors":["Theobald"],"year":{"year":"1905"}}},"details":{"species":{"genus":"Aedes","species":"w-albus","authorship":{"verbatim":"(Theobald, 1905)","normalized":"(Theobald 1905)","year":"1905","authors":["Theobald"],"originalAuth":{"authors":["Theobald"],"year":{"year":"1905"}}}}},"words":[{"verbatim":"Aedes","normalized":"Aedes","wordType":"GENUS","start":0,"end":5},{"verbatim":"w-albus","normalized":"w-albus","wordType":"SPECIES","start":6,"end":13},{"verbatim":"Theobald","normalized":"Theobald","wordType":"AUTHOR_WORD","start":15,"end":23},{"verbatim":"1905","normalized":"1905","wordType":"YEAR","start":25,"end":29}],"id":"7b0dd259-10ae-5b47-95ca-2685d4c323ce","parserVersion":"test_version"}
```

Name: Abryna regis-petri Paiva, 1860

Canonical: Abryna regis-petri

Authorship: Paiva 1860

```json
{"parsed":true,"quality":1,"verbatim":"Abryna regis-petri Paiva, 1860","normalized":"Abryna regis-petri Paiva 1860","canonical":{"stemmed":"Abryna regis-petr","simple":"Abryna regis-petri","full":"Abryna regis-petri"},"cardinality":2,"authorship":{"verbatim":"Paiva, 1860","normalized":"Paiva 1860","year":"1860","authors":["Paiva"],"originalAuth":{"authors":["Paiva"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Abryna","species":"regis-petri","authorship":{"verbatim":"Paiva, 1860","normalized":"Paiva 1860","year":"1860","authors":["Paiva"],"originalAuth":{"authors":["Paiva"],"year":{"year":"1860"}}}}},"words":[{"verbatim":"Abryna","normalized":"Abryna","wordType":"GENUS","start":0,"end":6},{"verbatim":"regis-petri","normalized":"regis-petri","wordType":"SPECIES","start":7,"end":18},{"verbatim":"Paiva","normalized":"Paiva","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"1860","normalized":"1860","wordType":"YEAR","start":26,"end":30}],"id":"27ad601d-bb92-515b-9c45-1faa55cdf7f3","parserVersion":"test_version"}
```

<!--
Abryna- regis|{"name_string_id":"9ff9c1fa-068e-5296-8c39-66e1c58f0660","parsed":false,"parser_version":"test_version","verbatim":"Abryna- regis","normalized":null,"canonical":null,"hybrid":false,"virus":false}
Abryna regis- Paiva, 1860|{"name_string_id":"473b8b63-8d5c-521f-9a68-7aecd5b9a62c","parsed":false,"parser_version":"test_version","verbatim":"Abryna regis- Paiva, 1860","normalized":null,"canonical":null,"hybrid":false,"virus":false}
-->

Name: Solms-laubachia orbiculata Y.C. Lan & T.Y. Cheo

Canonical: Solms-laubachia orbiculata

Authorship: Y. C. Lan & T. Y. Cheo

```json
{"parsed":true,"quality":1,"verbatim":"Solms-laubachia orbiculata Y.C. Lan \u0026 T.Y. Cheo","normalized":"Solms-laubachia orbiculata Y. C. Lan \u0026 T. Y. Cheo","canonical":{"stemmed":"Solms-laubachia orbiculat","simple":"Solms-laubachia orbiculata","full":"Solms-laubachia orbiculata"},"cardinality":2,"authorship":{"verbatim":"Y.C. Lan \u0026 T.Y. Cheo","normalized":"Y. C. Lan \u0026 T. Y. Cheo","authors":["Y. C. Lan","T. Y. Cheo"],"originalAuth":{"authors":["Y. C. Lan","T. Y. Cheo"]}},"details":{"species":{"genus":"Solms-laubachia","species":"orbiculata","authorship":{"verbatim":"Y.C. Lan \u0026 T.Y. Cheo","normalized":"Y. C. Lan \u0026 T. Y. Cheo","authors":["Y. C. Lan","T. Y. Cheo"],"originalAuth":{"authors":["Y. C. Lan","T. Y. Cheo"]}}}},"words":[{"verbatim":"Solms-laubachia","normalized":"Solms-laubachia","wordType":"GENUS","start":0,"end":15},{"verbatim":"orbiculata","normalized":"orbiculata","wordType":"SPECIES","start":16,"end":26},{"verbatim":"Y.","normalized":"Y.","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Lan","normalized":"Lan","wordType":"AUTHOR_WORD","start":32,"end":35},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"Y.","normalized":"Y.","wordType":"AUTHOR_WORD","start":40,"end":42},{"verbatim":"Cheo","normalized":"Cheo","wordType":"AUTHOR_WORD","start":43,"end":47}],"id":"4dce39e2-ffd7-5a1b-bd1a-2bc12049be90","parserVersion":"test_version"}
```

### Authorship with 'degli'

Name: Cestodiscus gemmifer F. S. Castracane degli Antelminelli

Canonical: Cestodiscus gemmifer

Authorship: F. S. Castracane degli Antelminelli

```json
{"parsed":true,"quality":1,"verbatim":"Cestodiscus gemmifer F. S. Castracane degli Antelminelli","normalized":"Cestodiscus gemmifer F. S. Castracane degli Antelminelli","canonical":{"stemmed":"Cestodiscus gemmifer","simple":"Cestodiscus gemmifer","full":"Cestodiscus gemmifer"},"cardinality":2,"authorship":{"verbatim":"F. S. Castracane degli Antelminelli","normalized":"F. S. Castracane degli Antelminelli","authors":["F. S. Castracane degli Antelminelli"],"originalAuth":{"authors":["F. S. Castracane degli Antelminelli"]}},"details":{"species":{"genus":"Cestodiscus","species":"gemmifer","authorship":{"verbatim":"F. S. Castracane degli Antelminelli","normalized":"F. S. Castracane degli Antelminelli","authors":["F. S. Castracane degli Antelminelli"],"originalAuth":{"authors":["F. S. Castracane degli Antelminelli"]}}}},"words":[{"verbatim":"Cestodiscus","normalized":"Cestodiscus","wordType":"GENUS","start":0,"end":11},{"verbatim":"gemmifer","normalized":"gemmifer","wordType":"SPECIES","start":12,"end":20},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"Castracane","normalized":"Castracane","wordType":"AUTHOR_WORD","start":27,"end":37},{"verbatim":"degli","normalized":"degli","wordType":"AUTHOR_WORD","start":38,"end":43},{"verbatim":"Antelminelli","normalized":"Antelminelli","wordType":"AUTHOR_WORD","start":44,"end":56}],"id":"95572f76-8ce0-5ba4-ae63-7492d37d0bed","parserVersion":"test_version"}
```

### Authorship with filius (son of)

Name: Oxytropis minjanensis Rech. f.

Canonical: Oxytropis minjanensis

Authorship: Rech. fil.

```json
{"parsed":true,"quality":1,"verbatim":"Oxytropis minjanensis Rech. f.","normalized":"Oxytropis minjanensis Rech. fil.","canonical":{"stemmed":"Oxytropis minianens","simple":"Oxytropis minjanensis","full":"Oxytropis minjanensis"},"cardinality":2,"authorship":{"verbatim":"Rech. f.","normalized":"Rech. fil.","authors":["Rech. fil."],"originalAuth":{"authors":["Rech. fil."]}},"details":{"species":{"genus":"Oxytropis","species":"minjanensis","authorship":{"verbatim":"Rech. f.","normalized":"Rech. fil.","authors":["Rech. fil."],"originalAuth":{"authors":["Rech. fil."]}}}},"words":[{"verbatim":"Oxytropis","normalized":"Oxytropis","wordType":"GENUS","start":0,"end":9},{"verbatim":"minjanensis","normalized":"minjanensis","wordType":"SPECIES","start":10,"end":21},{"verbatim":"Rech.","normalized":"Rech.","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":28,"end":30}],"id":"6027cbc2-fa15-510b-ab3e-e1fa44cbd551","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl f. 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl fil. 1935

```json
{"parsed":true,"quality":1,"verbatim":"Platypus bicaudatulus Schedl f. 1935","normalized":"Platypus bicaudatulus Schedl fil. 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl f. 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl f. 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":29,"end":31},{"verbatim":"1935","normalized":"1935","wordType":"YEAR","start":32,"end":36}],"id":"05799df9-471e-5c68-92fe-4edcc0a69d29","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl filius 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl fil. 1935

```json
{"parsed":true,"quality":1,"verbatim":"Platypus bicaudatulus Schedl filius 1935","normalized":"Platypus bicaudatulus Schedl fil. 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl filius 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl filius 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}}}},"words":[{"verbatim":"Platypus","normalized":"Platypus","wordType":"GENUS","start":0,"end":8},{"verbatim":"bicaudatulus","normalized":"bicaudatulus","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Schedl","normalized":"Schedl","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"filius","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":29,"end":35},{"verbatim":"1935","normalized":"1935","wordType":"YEAR","start":36,"end":40}],"id":"2b6cd51f-aa0f-58fd-88fa-2e261cedacbb","parserVersion":"test_version"}
```

Name: Fimbristylis ovata (Burm. f.) J. Kern

Canonical: Fimbristylis ovata

Authorship: (Burm. fil.) J. Kern

```json
{"parsed":true,"quality":1,"verbatim":"Fimbristylis ovata (Burm. f.) J. Kern","normalized":"Fimbristylis ovata (Burm. fil.) J. Kern","canonical":{"stemmed":"Fimbristylis ouat","simple":"Fimbristylis ovata","full":"Fimbristylis ovata"},"cardinality":2,"authorship":{"verbatim":"(Burm. f.) J. Kern","normalized":"(Burm. fil.) J. Kern","authors":["Burm. fil.","J. Kern"],"originalAuth":{"authors":["Burm. fil."]},"combinationAuth":{"authors":["J. Kern"]}},"details":{"species":{"genus":"Fimbristylis","species":"ovata","authorship":{"verbatim":"(Burm. f.) J. Kern","normalized":"(Burm. fil.) J. Kern","authors":["Burm. fil.","J. Kern"],"originalAuth":{"authors":["Burm. fil."]},"combinationAuth":{"authors":["J. Kern"]}}}},"words":[{"verbatim":"Fimbristylis","normalized":"Fimbristylis","wordType":"GENUS","start":0,"end":12},{"verbatim":"ovata","normalized":"ovata","wordType":"SPECIES","start":13,"end":18},{"verbatim":"Burm.","normalized":"Burm.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":26,"end":28},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":30,"end":32},{"verbatim":"Kern","normalized":"Kern","wordType":"AUTHOR_WORD","start":33,"end":37}],"id":"01207e0b-8de4-5a4e-99fc-e60b581c0d1c","parserVersion":"test_version"}
```

Name: Carex chordorrhiza Ehrh. ex L. f.

Canonical: Carex chordorrhiza

Authorship: Ehrh. ex L. fil.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Carex chordorrhiza Ehrh. ex L. f.","normalized":"Carex chordorrhiza Ehrh. ex L. fil.","canonical":{"stemmed":"Carex chordorrhiz","simple":"Carex chordorrhiza","full":"Carex chordorrhiza"},"cardinality":2,"authorship":{"verbatim":"Ehrh. ex L. f.","normalized":"Ehrh. ex L. fil.","authors":["Ehrh.","L. fil."],"originalAuth":{"authors":["Ehrh."],"exAuthors":{"authors":["L. fil."]}}},"details":{"species":{"genus":"Carex","species":"chordorrhiza","authorship":{"verbatim":"Ehrh. ex L. f.","normalized":"Ehrh. ex L. fil.","authors":["Ehrh.","L. fil."],"originalAuth":{"authors":["Ehrh."],"exAuthors":{"authors":["L. fil."]}}}}},"words":[{"verbatim":"Carex","normalized":"Carex","wordType":"GENUS","start":0,"end":5},{"verbatim":"chordorrhiza","normalized":"chordorrhiza","wordType":"SPECIES","start":6,"end":18},{"verbatim":"Ehrh.","normalized":"Ehrh.","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":31,"end":33}],"id":"b972d277-3714-5549-9103-869675f490bd","parserVersion":"test_version"}
```

Name: Amelanchier arborea var. arborea (Michx. f.) Fernald

Canonical: Amelanchier arborea var. arborea

Authorship: (Michx. fil.) Fernald

```json
{"parsed":true,"quality":1,"verbatim":"Amelanchier arborea var. arborea (Michx. f.) Fernald","normalized":"Amelanchier arborea var. arborea (Michx. fil.) Fernald","canonical":{"stemmed":"Amelanchier arbore arbore","simple":"Amelanchier arborea arborea","full":"Amelanchier arborea var. arborea"},"cardinality":3,"authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}},"details":{"infraspecies":{"genus":"Amelanchier","species":"arborea","infraspecies":[{"value":"arborea","rank":"var.","authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}}}]}},"words":[{"verbatim":"Amelanchier","normalized":"Amelanchier","wordType":"GENUS","start":0,"end":11},{"verbatim":"arborea","normalized":"arborea","wordType":"SPECIES","start":12,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"arborea","normalized":"arborea","wordType":"INFRASPECIES","start":25,"end":32},{"verbatim":"Michx.","normalized":"Michx.","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":41,"end":43},{"verbatim":"Fernald","normalized":"Fernald","wordType":"AUTHOR_WORD","start":45,"end":52}],"id":"1644869c-3e0c-5e7e-a709-a86dee11b917","parserVersion":"test_version"}
```

Name: Cerastium arvense var. fuegianum Hook. f.

Canonical: Cerastium arvense var. fuegianum

Authorship: Hook. fil.

```json
{"parsed":true,"quality":1,"verbatim":"Cerastium arvense var. fuegianum Hook. f.","normalized":"Cerastium arvense var. fuegianum Hook. fil.","canonical":{"stemmed":"Cerastium aruens fuegian","simple":"Cerastium arvense fuegianum","full":"Cerastium arvense var. fuegianum"},"cardinality":3,"authorship":{"verbatim":"Hook. f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraspecies":{"genus":"Cerastium","species":"arvense","infraspecies":[{"value":"fuegianum","rank":"var.","authorship":{"verbatim":"Hook. f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"words":[{"verbatim":"Cerastium","normalized":"Cerastium","wordType":"GENUS","start":0,"end":9},{"verbatim":"arvense","normalized":"arvense","wordType":"SPECIES","start":10,"end":17},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":18,"end":22},{"verbatim":"fuegianum","normalized":"fuegianum","wordType":"INFRASPECIES","start":23,"end":32},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":39,"end":41}],"id":"f9fb925a-777f-5a2c-892d-bdf11528dbfc","parserVersion":"test_version"}
```

Name: Cerastium arvense var. fuegianum Hook.f.

Canonical: Cerastium arvense var. fuegianum

Authorship: Hook. fil.

```json
{"parsed":true,"quality":1,"verbatim":"Cerastium arvense var. fuegianum Hook.f.","normalized":"Cerastium arvense var. fuegianum Hook. fil.","canonical":{"stemmed":"Cerastium aruens fuegian","simple":"Cerastium arvense fuegianum","full":"Cerastium arvense var. fuegianum"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraspecies":{"genus":"Cerastium","species":"arvense","infraspecies":[{"value":"fuegianum","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"words":[{"verbatim":"Cerastium","normalized":"Cerastium","wordType":"GENUS","start":0,"end":9},{"verbatim":"arvense","normalized":"arvense","wordType":"SPECIES","start":10,"end":17},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":18,"end":22},{"verbatim":"fuegianum","normalized":"fuegianum","wordType":"INFRASPECIES","start":23,"end":32},{"verbatim":"Hook.","normalized":"Hook.","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":38,"end":40}],"id":"35ea20fb-b794-572f-ba90-36c1463e1927","parserVersion":"test_version"}
```

Name: Cerastium arvense ssp. velutinum var. velutinum (Raf.) Britton f.

Canonical: Cerastium arvense subsp. velutinum var. velutinum

Authorship: (Raf.) Britton fil.

```json
{"parsed":true,"quality":1,"verbatim":"Cerastium arvense ssp. velutinum var. velutinum (Raf.) Britton f.","normalized":"Cerastium arvense subsp. velutinum var. velutinum (Raf.) Britton fil.","canonical":{"stemmed":"Cerastium aruens uelutin uelutin","simple":"Cerastium arvense velutinum velutinum","full":"Cerastium arvense subsp. velutinum var. velutinum"},"cardinality":4,"authorship":{"verbatim":"(Raf.) Britton f.","normalized":"(Raf.) Britton fil.","authors":["Raf.","Britton fil."],"originalAuth":{"authors":["Raf."]},"combinationAuth":{"authors":["Britton fil."]}},"details":{"infraspecies":{"genus":"Cerastium","species":"arvense","infraspecies":[{"value":"velutinum","rank":"subsp."},{"value":"velutinum","rank":"var.","authorship":{"verbatim":"(Raf.) Britton f.","normalized":"(Raf.) Britton fil.","authors":["Raf.","Britton fil."],"originalAuth":{"authors":["Raf."]},"combinationAuth":{"authors":["Britton fil."]}}}]}},"words":[{"verbatim":"Cerastium","normalized":"Cerastium","wordType":"GENUS","start":0,"end":9},{"verbatim":"arvense","normalized":"arvense","wordType":"SPECIES","start":10,"end":17},{"verbatim":"ssp.","normalized":"subsp.","wordType":"RANK","start":18,"end":22},{"verbatim":"velutinum","normalized":"velutinum","wordType":"INFRASPECIES","start":23,"end":32},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":33,"end":37},{"verbatim":"velutinum","normalized":"velutinum","wordType":"INFRASPECIES","start":38,"end":47},{"verbatim":"Raf.","normalized":"Raf.","wordType":"AUTHOR_WORD","start":49,"end":53},{"verbatim":"Britton","normalized":"Britton","wordType":"AUTHOR_WORD","start":55,"end":62},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":63,"end":65}],"id":"c7841295-3aa3-5c40-8adf-88d177f74cbe","parserVersion":"test_version"}
```

Name: Jacquemontia spiciflora (Choisy) Hall. fil.

Canonical: Jacquemontia spiciflora

Authorship: (Choisy) Hall. fil.

```json
{"parsed":true,"quality":1,"verbatim":"Jacquemontia spiciflora (Choisy) Hall. fil.","normalized":"Jacquemontia spiciflora (Choisy) Hall. fil.","canonical":{"stemmed":"Jacquemontia spiciflor","simple":"Jacquemontia spiciflora","full":"Jacquemontia spiciflora"},"cardinality":2,"authorship":{"verbatim":"(Choisy) Hall. fil.","normalized":"(Choisy) Hall. fil.","authors":["Choisy","Hall. fil."],"originalAuth":{"authors":["Choisy"]},"combinationAuth":{"authors":["Hall. fil."]}},"details":{"species":{"genus":"Jacquemontia","species":"spiciflora","authorship":{"verbatim":"(Choisy) Hall. fil.","normalized":"(Choisy) Hall. fil.","authors":["Choisy","Hall. fil."],"originalAuth":{"authors":["Choisy"]},"combinationAuth":{"authors":["Hall. fil."]}}}},"words":[{"verbatim":"Jacquemontia","normalized":"Jacquemontia","wordType":"GENUS","start":0,"end":12},{"verbatim":"spiciflora","normalized":"spiciflora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Choisy","normalized":"Choisy","wordType":"AUTHOR_WORD","start":25,"end":31},{"verbatim":"Hall.","normalized":"Hall.","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"fil.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":39,"end":43}],"id":"14a98945-4e97-5c13-a0b9-97741641a6a4","parserVersion":"test_version"}
```

Name: Littorina (Littorina) littorea fa major (Linnaeus, 1758)

Canonical: Littorina littorea f. major

Authorship: (Linnaeus 1758)

(Linnaeus 1758)```json
{"parsed":true,"quality":1,"verbatim":"Littorina (Littorina) littorea fa major (Linnaeus, 1758)","normalized":"Littorina (Littorina) littorea f. major (Linnaeus 1758)","canonical":{"stemmed":"Littorina littore maior","simple":"Littorina littorea major","full":"Littorina littorea f. major"},"cardinality":3,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"infraspecies":{"genus":"Littorina","subgenus":"Littorina","species":"littorea","infraspecies":[{"value":"major","rank":"f.","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}]}},"words":[{"verbatim":"Littorina","normalized":"Littorina","wordType":"GENUS","start":0,"end":9},{"verbatim":"Littorina","normalized":"Littorina","wordType":"INFRA_GENUS","start":11,"end":20},{"verbatim":"littorea","normalized":"littorea","wordType":"SPECIES","start":22,"end":30},{"verbatim":"fa","normalized":"f.","wordType":"RANK","start":31,"end":33},{"verbatim":"major","normalized":"major","wordType":"INFRASPECIES","start":34,"end":39},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":41,"end":49},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":51,"end":55}],"id":"fcd777b2-d8c9-5fe5-9883-ed0affa4a0e2","parserVersion":"test_version"}
```

Name: Amelanchier arborea f. hirsuta (Michx. f.) Fernald

Canonical: Amelanchier arborea f. hirsuta

Authorship: (Michx. fil.) Fernald

```json
{"parsed":true,"quality":1,"verbatim":"Amelanchier arborea f. hirsuta (Michx. f.) Fernald","normalized":"Amelanchier arborea f. hirsuta (Michx. fil.) Fernald","canonical":{"stemmed":"Amelanchier arbore hirsut","simple":"Amelanchier arborea hirsuta","full":"Amelanchier arborea f. hirsuta"},"cardinality":3,"authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}},"details":{"infraspecies":{"genus":"Amelanchier","species":"arborea","infraspecies":[{"value":"hirsuta","rank":"f.","authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}}}]}},"words":[{"verbatim":"Amelanchier","normalized":"Amelanchier","wordType":"GENUS","start":0,"end":11},{"verbatim":"arborea","normalized":"arborea","wordType":"SPECIES","start":12,"end":19},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":20,"end":22},{"verbatim":"hirsuta","normalized":"hirsuta","wordType":"INFRASPECIES","start":23,"end":30},{"verbatim":"Michx.","normalized":"Michx.","wordType":"AUTHOR_WORD","start":32,"end":38},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":39,"end":41},{"verbatim":"Fernald","normalized":"Fernald","wordType":"AUTHOR_WORD","start":43,"end":50}],"id":"f5786fa9-2b40-5ee4-8786-ffe86ed02ab5","parserVersion":"test_version"}
```

Name: Betula pendula fo. dalecarlica (L. f.) C.K. Schneid.

Canonical: Betula pendula f. dalecarlica

Authorship: (L. fil.) C. K. Schneid.

```json
{"parsed":true,"quality":1,"verbatim":"Betula pendula fo. dalecarlica (L. f.) C.K. Schneid.","normalized":"Betula pendula f. dalecarlica (L. fil.) C. K. Schneid.","canonical":{"stemmed":"Betula pendul dalecarlic","simple":"Betula pendula dalecarlica","full":"Betula pendula f. dalecarlica"},"cardinality":3,"authorship":{"verbatim":"(L. f.) C.K. Schneid.","normalized":"(L. fil.) C. K. Schneid.","authors":["L. fil.","C. K. Schneid."],"originalAuth":{"authors":["L. fil."]},"combinationAuth":{"authors":["C. K. Schneid."]}},"details":{"infraspecies":{"genus":"Betula","species":"pendula","infraspecies":[{"value":"dalecarlica","rank":"f.","authorship":{"verbatim":"(L. f.) C.K. Schneid.","normalized":"(L. fil.) C. K. Schneid.","authors":["L. fil.","C. K. Schneid."],"originalAuth":{"authors":["L. fil."]},"combinationAuth":{"authors":["C. K. Schneid."]}}}]}},"words":[{"verbatim":"Betula","normalized":"Betula","wordType":"GENUS","start":0,"end":6},{"verbatim":"pendula","normalized":"pendula","wordType":"SPECIES","start":7,"end":14},{"verbatim":"fo.","normalized":"f.","wordType":"RANK","start":15,"end":18},{"verbatim":"dalecarlica","normalized":"dalecarlica","wordType":"INFRASPECIES","start":19,"end":30},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":32,"end":34},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":35,"end":37},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"K.","normalized":"K.","wordType":"AUTHOR_WORD","start":41,"end":43},{"verbatim":"Schneid.","normalized":"Schneid.","wordType":"AUTHOR_WORD","start":44,"end":52}],"id":"4c4ee33c-9738-5542-b22f-2326996aa6f7","parserVersion":"test_version"}
```

Name: Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.

Canonical: Racomitrium canescens f. ericoides

Authorship: (F. Weber ex Brid.) Mönk.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","normalized":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","canonical":{"stemmed":"Racomitrium canescens ericoid","simple":"Racomitrium canescens ericoides","full":"Racomitrium canescens f. ericoides"},"cardinality":3,"authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Brid.","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}},"details":{"infraspecies":{"genus":"Racomitrium","species":"canescens","infraspecies":[{"value":"ericoides","rank":"f.","authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Brid.","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}}}]}},"words":[{"verbatim":"Racomitrium","normalized":"Racomitrium","wordType":"GENUS","start":0,"end":11},{"verbatim":"canescens","normalized":"canescens","wordType":"SPECIES","start":12,"end":21},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":22,"end":24},{"verbatim":"ericoides","normalized":"ericoides","wordType":"INFRASPECIES","start":25,"end":34},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"Weber","normalized":"Weber","wordType":"AUTHOR_WORD","start":39,"end":44},{"verbatim":"Brid.","normalized":"Brid.","wordType":"AUTHOR_WORD","start":48,"end":53},{"verbatim":"Mönk.","normalized":"Mönk.","wordType":"AUTHOR_WORD","start":55,"end":60}],"id":"45a001f1-749f-5803-bd92-93c6d524e9db","parserVersion":"test_version"}
```

Name: Racomitrium canescens forma ericoides (F. Weber ex Brid.) Mönk.

Canonical: Racomitrium canescens f. ericoides

Authorship: (F. Weber ex Brid.) Mönk.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Racomitrium canescens forma ericoides (F. Weber ex Brid.) Mönk.","normalized":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","canonical":{"stemmed":"Racomitrium canescens ericoid","simple":"Racomitrium canescens ericoides","full":"Racomitrium canescens f. ericoides"},"cardinality":3,"authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Brid.","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}},"details":{"infraspecies":{"genus":"Racomitrium","species":"canescens","infraspecies":[{"value":"ericoides","rank":"f.","authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Brid.","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}}}]}},"words":[{"verbatim":"Racomitrium","normalized":"Racomitrium","wordType":"GENUS","start":0,"end":11},{"verbatim":"canescens","normalized":"canescens","wordType":"SPECIES","start":12,"end":21},{"verbatim":"forma","normalized":"f.","wordType":"RANK","start":22,"end":27},{"verbatim":"ericoides","normalized":"ericoides","wordType":"INFRASPECIES","start":28,"end":37},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"Weber","normalized":"Weber","wordType":"AUTHOR_WORD","start":42,"end":47},{"verbatim":"Brid.","normalized":"Brid.","wordType":"AUTHOR_WORD","start":51,"end":56},{"verbatim":"Mönk.","normalized":"Mönk.","wordType":"AUTHOR_WORD","start":58,"end":63}],"id":"8a58ed91-9a71-5278-9bd1-b8e82188e938","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f., Rosenst.

Canonical: Polypodium pectinatum

Authorship: L. fil. & Rosenst.

```json
{"parsed":true,"quality":1,"verbatim":"Polypodium pectinatum L. f., Rosenst.","normalized":"Polypodium pectinatum L. fil. \u0026 Rosenst.","canonical":{"stemmed":"Polypodium pectinat","simple":"Polypodium pectinatum","full":"Polypodium pectinatum"},"cardinality":2,"authorship":{"verbatim":"L. f., Rosenst.","normalized":"L. fil. \u0026 Rosenst.","authors":["L. fil.","Rosenst."],"originalAuth":{"authors":["L. fil.","Rosenst."]}},"details":{"species":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L. f., Rosenst.","normalized":"L. fil. \u0026 Rosenst.","authors":["L. fil.","Rosenst."],"originalAuth":{"authors":["L. fil.","Rosenst."]}}}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":25,"end":27},{"verbatim":"Rosenst.","normalized":"Rosenst.","wordType":"AUTHOR_WORD","start":29,"end":37}],"id":"bac3cf47-358a-51e2-83a6-6577d0f362af","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f.

Canonical: Polypodium pectinatum

Authorship: L. fil.

```json
{"parsed":true,"quality":1,"verbatim":"Polypodium pectinatum L. f.","normalized":"Polypodium pectinatum L. fil.","canonical":{"stemmed":"Polypodium pectinat","simple":"Polypodium pectinatum","full":"Polypodium pectinatum"},"cardinality":2,"authorship":{"verbatim":"L. f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"details":{"species":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L. f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}}}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":25,"end":27}],"id":"e4c2c98c-79c9-5ee1-865a-300a0c0287ef","parserVersion":"test_version"}
```

Name: Polypodium pectinatum (L. f.) typica Rosent

Canonical: Polypodium pectinatum typica

Authorship: Rosent

```json
{"parsed":true,"quality":1,"verbatim":"Polypodium pectinatum (L. f.) typica Rosent","normalized":"Polypodium pectinatum (L. fil.) typica Rosent","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum typica"},"cardinality":3,"authorship":{"verbatim":"Rosent","normalized":"Rosent","authors":["Rosent"],"originalAuth":{"authors":["Rosent"]}},"details":{"infraspecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"(L. f.)","normalized":"(L. fil.)","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"infraspecies":[{"value":"typica","authorship":{"verbatim":"Rosent","normalized":"Rosent","authors":["Rosent"],"originalAuth":{"authors":["Rosent"]}}}]}},"words":[{"verbatim":"Polypodium","normalized":"Polypodium","wordType":"GENUS","start":0,"end":10},{"verbatim":"pectinatum","normalized":"pectinatum","wordType":"SPECIES","start":11,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":26,"end":28},{"verbatim":"typica","normalized":"typica","wordType":"INFRASPECIES","start":30,"end":36},{"verbatim":"Rosent","normalized":"Rosent","wordType":"AUTHOR_WORD","start":37,"end":43}],"id":"b345d921-7466-50bb-812c-850b1f368c57","parserVersion":"test_version"}
```

### Names with emend (rectified by) authorship

Name: Chlorobium phaeobacteroides Pfennig, 1968 emend. Imhoff, 2003

Canonical: Chlorobium phaeobacteroides

Authorship: Pfennig 1968 emend. Imhoff 2003

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Emend authors are not required"}],"verbatim":"Chlorobium phaeobacteroides Pfennig, 1968 emend. Imhoff, 2003","normalized":"Chlorobium phaeobacteroides Pfennig 1968 emend. Imhoff 2003","canonical":{"stemmed":"Chlorobium phaeobacteroid","simple":"Chlorobium phaeobacteroides","full":"Chlorobium phaeobacteroides"},"cardinality":2,"authorship":{"verbatim":"Pfennig, 1968 emend. Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}},"bacteria":"yes","details":{"species":{"genus":"Chlorobium","species":"phaeobacteroides","authorship":{"verbatim":"Pfennig, 1968 emend. Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}}}},"words":[{"verbatim":"Chlorobium","normalized":"Chlorobium","wordType":"GENUS","start":0,"end":10},{"verbatim":"phaeobacteroides","normalized":"phaeobacteroides","wordType":"SPECIES","start":11,"end":27},{"verbatim":"Pfennig","normalized":"Pfennig","wordType":"AUTHOR_WORD","start":28,"end":35},{"verbatim":"1968","normalized":"1968","wordType":"YEAR","start":37,"end":41},{"verbatim":"Imhoff","normalized":"Imhoff","wordType":"AUTHOR_WORD","start":49,"end":55},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":57,"end":61}],"id":"4513701d-e56b-54d6-84a7-941bf4b62e69","parserVersion":"test_version"}
```

Name: Chlorobium phaeobacteroides Pfennig, 1968 emend Imhoff, 2003

Canonical: Chlorobium phaeobacteroides

Authorship: Pfennig 1968 emend. Imhoff 2003

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"`emend` without a period"},{"quality":2,"warning":"Emend authors are not required"}],"verbatim":"Chlorobium phaeobacteroides Pfennig, 1968 emend Imhoff, 2003","normalized":"Chlorobium phaeobacteroides Pfennig 1968 emend. Imhoff 2003","canonical":{"stemmed":"Chlorobium phaeobacteroid","simple":"Chlorobium phaeobacteroides","full":"Chlorobium phaeobacteroides"},"cardinality":2,"authorship":{"verbatim":"Pfennig, 1968 emend Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}},"bacteria":"yes","details":{"species":{"genus":"Chlorobium","species":"phaeobacteroides","authorship":{"verbatim":"Pfennig, 1968 emend Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}}}},"words":[{"verbatim":"Chlorobium","normalized":"Chlorobium","wordType":"GENUS","start":0,"end":10},{"verbatim":"phaeobacteroides","normalized":"phaeobacteroides","wordType":"SPECIES","start":11,"end":27},{"verbatim":"Pfennig","normalized":"Pfennig","wordType":"AUTHOR_WORD","start":28,"end":35},{"verbatim":"1968","normalized":"1968","wordType":"YEAR","start":37,"end":41},{"verbatim":"Imhoff","normalized":"Imhoff","wordType":"AUTHOR_WORD","start":48,"end":54},{"verbatim":"2003","normalized":"2003","wordType":"YEAR","start":56,"end":60}],"id":"3cbaceda-83c2-5e36-b170-4f13837782dc","parserVersion":"test_version"}
```

### Names with an unparsed "tail"

Name: Morea (Morea) Burt 2342343242 23424322342 23424234

Canonical: Morea subgen. Morea

Authorship: Burt

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Morea (Morea) Burt 2342343242 23424322342 23424234","normalized":"Morea subgen. Morea Burt","canonical":{"stemmed":"Morea","simple":"Morea","full":"Morea subgen. Morea"},"cardinality":1,"authorship":{"verbatim":"Burt","normalized":"Burt","authors":["Burt"],"originalAuth":{"authors":["Burt"]}},"tail":" 2342343242 23424322342 23424234","details":{"uninomial":{"uninomial":"Morea","rank":"subgen.","parent":"Morea","authorship":{"verbatim":"Burt","normalized":"Burt","authors":["Burt"],"originalAuth":{"authors":["Burt"]}}}},"words":[{"verbatim":"Morea","normalized":"Morea","wordType":"UNINOMIAL","start":0,"end":5},{"verbatim":"Morea","normalized":"Morea","wordType":"UNINOMIAL","start":7,"end":12},{"verbatim":"Burt","normalized":"Burt","wordType":"AUTHOR_WORD","start":14,"end":18}],"id":"ca23679f-f3d8-5194-a406-048f970c4020","parserVersion":"test_version"}
```

Name: Nautilus asterizans von

Canonical: Nautilus asterizans

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nautilus asterizans von","normalized":"Nautilus asterizans","canonical":{"stemmed":"Nautilus asterizans","simple":"Nautilus asterizans","full":"Nautilus asterizans"},"cardinality":2,"tail":" von","details":{"species":{"genus":"Nautilus","species":"asterizans"}},"words":[{"verbatim":"Nautilus","normalized":"Nautilus","wordType":"GENUS","start":0,"end":8},{"verbatim":"asterizans","normalized":"asterizans","wordType":"SPECIES","start":9,"end":19}],"id":"0716f658-c952-5415-b2ad-79a39c2b7b0d","parserVersion":"test_version"}
```

Name: Dryopteris X separabilis Small (pro sp.)

Canonical: Dryopteris × separabilis

Authorship: Small

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Dryopteris X separabilis Small (pro sp.)","normalized":"Dryopteris × separabilis Small","canonical":{"stemmed":"Dryopteris separabil","simple":"Dryopteris separabilis","full":"Dryopteris × separabilis"},"cardinality":2,"authorship":{"verbatim":"Small","normalized":"Small","authors":["Small"],"originalAuth":{"authors":["Small"]}},"hybrid":"NAMED_HYBRID","tail":" (pro sp.)","details":{"species":{"genus":"Dryopteris","species":"separabilis Small","authorship":{"verbatim":"Small","normalized":"Small","authors":["Small"],"originalAuth":{"authors":["Small"]}}}},"words":[{"verbatim":"Dryopteris","normalized":"Dryopteris","wordType":"GENUS","start":0,"end":10},{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":11,"end":12},{"verbatim":"separabilis","normalized":"separabilis","wordType":"SPECIES","start":13,"end":24},{"verbatim":"Small","normalized":"Small","wordType":"AUTHOR_WORD","start":25,"end":30}],"id":"34bf83d8-0466-51c4-b95d-70e583ba1c9f","parserVersion":"test_version"}
```

Name: Eulima excellens Verkrüzen fide Paetel, 1887

Canonical: Eulima excellens

Authorship: Verkrüzen

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Eulima excellens Verkrüzen fide Paetel, 1887","normalized":"Eulima excellens Verkrüzen","canonical":{"stemmed":"Eulima excellens","simple":"Eulima excellens","full":"Eulima excellens"},"cardinality":2,"authorship":{"verbatim":"Verkrüzen","normalized":"Verkrüzen","authors":["Verkrüzen"],"originalAuth":{"authors":["Verkrüzen"]}},"tail":" fide Paetel, 1887","details":{"species":{"genus":"Eulima","species":"excellens","authorship":{"verbatim":"Verkrüzen","normalized":"Verkrüzen","authors":["Verkrüzen"],"originalAuth":{"authors":["Verkrüzen"]}}}},"words":[{"verbatim":"Eulima","normalized":"Eulima","wordType":"GENUS","start":0,"end":6},{"verbatim":"excellens","normalized":"excellens","wordType":"SPECIES","start":7,"end":16},{"verbatim":"Verkrüzen","normalized":"Verkrüzen","wordType":"AUTHOR_WORD","start":17,"end":26}],"id":"1e5dd590-289c-5e83-9f93-64f46f334eef","parserVersion":"test_version"}
```

Name: Procamallanus (Spirocamallanus) soodi Lakshmi & Kumari, 2001 nec (Gupta & Masood, 1988)

Canonical: Procamallanus soodi

Authorship: Lakshmi & Kumari 2001

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Procamallanus (Spirocamallanus) soodi Lakshmi \u0026 Kumari, 2001 nec (Gupta \u0026 Masood, 1988)","normalized":"Procamallanus (Spirocamallanus) soodi Lakshmi \u0026 Kumari 2001","canonical":{"stemmed":"Procamallanus sood","simple":"Procamallanus soodi","full":"Procamallanus soodi"},"cardinality":2,"authorship":{"verbatim":"Lakshmi \u0026 Kumari, 2001","normalized":"Lakshmi \u0026 Kumari 2001","year":"2001","authors":["Lakshmi","Kumari"],"originalAuth":{"authors":["Lakshmi","Kumari"],"year":{"year":"2001"}}},"tail":" nec (Gupta \u0026 Masood, 1988)","details":{"species":{"genus":"Procamallanus","subgenus":"Spirocamallanus","species":"soodi","authorship":{"verbatim":"Lakshmi \u0026 Kumari, 2001","normalized":"Lakshmi \u0026 Kumari 2001","year":"2001","authors":["Lakshmi","Kumari"],"originalAuth":{"authors":["Lakshmi","Kumari"],"year":{"year":"2001"}}}}},"words":[{"verbatim":"Procamallanus","normalized":"Procamallanus","wordType":"GENUS","start":0,"end":13},{"verbatim":"Spirocamallanus","normalized":"Spirocamallanus","wordType":"INFRA_GENUS","start":15,"end":30},{"verbatim":"soodi","normalized":"soodi","wordType":"SPECIES","start":32,"end":37},{"verbatim":"Lakshmi","normalized":"Lakshmi","wordType":"AUTHOR_WORD","start":38,"end":45},{"verbatim":"Kumari","normalized":"Kumari","wordType":"AUTHOR_WORD","start":48,"end":54},{"verbatim":"2001","normalized":"2001","wordType":"YEAR","start":56,"end":60}],"id":"c024f8dd-f7e6-5add-869f-3f93e844ad1a","parserVersion":"test_version"}
```

Name: Membranipora minuscula Canu, 1911 non Hincks, 1882

Canonical: Membranipora minuscula

Authorship: Canu 1911

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Membranipora minuscula Canu, 1911 non Hincks, 1882","normalized":"Membranipora minuscula Canu 1911","canonical":{"stemmed":"Membranipora minuscul","simple":"Membranipora minuscula","full":"Membranipora minuscula"},"cardinality":2,"authorship":{"verbatim":"Canu, 1911","normalized":"Canu 1911","year":"1911","authors":["Canu"],"originalAuth":{"authors":["Canu"],"year":{"year":"1911"}}},"tail":" non Hincks, 1882","details":{"species":{"genus":"Membranipora","species":"minuscula","authorship":{"verbatim":"Canu, 1911","normalized":"Canu 1911","year":"1911","authors":["Canu"],"originalAuth":{"authors":["Canu"],"year":{"year":"1911"}}}}},"words":[{"verbatim":"Membranipora","normalized":"Membranipora","wordType":"GENUS","start":0,"end":12},{"verbatim":"minuscula","normalized":"minuscula","wordType":"SPECIES","start":13,"end":22},{"verbatim":"Canu","normalized":"Canu","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"1911","normalized":"1911","wordType":"YEAR","start":29,"end":33}],"id":"80abde40-859e-5909-aedc-928699ec7d05","parserVersion":"test_version"}
```

Name: Proboscina subechinata Canu & Bassler, 1920 non d'Orbigny, 1853

Canonical: Proboscina subechinata

Authorship: Canu & Bassler 1920

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Proboscina subechinata Canu \u0026 Bassler, 1920 non d'Orbigny, 1853","normalized":"Proboscina subechinata Canu \u0026 Bassler 1920","canonical":{"stemmed":"Proboscina subechinat","simple":"Proboscina subechinata","full":"Proboscina subechinata"},"cardinality":2,"authorship":{"verbatim":"Canu \u0026 Bassler, 1920","normalized":"Canu \u0026 Bassler 1920","year":"1920","authors":["Canu","Bassler"],"originalAuth":{"authors":["Canu","Bassler"],"year":{"year":"1920"}}},"tail":" non d'Orbigny, 1853","details":{"species":{"genus":"Proboscina","species":"subechinata","authorship":{"verbatim":"Canu \u0026 Bassler, 1920","normalized":"Canu \u0026 Bassler 1920","year":"1920","authors":["Canu","Bassler"],"originalAuth":{"authors":["Canu","Bassler"],"year":{"year":"1920"}}}}},"words":[{"verbatim":"Proboscina","normalized":"Proboscina","wordType":"GENUS","start":0,"end":10},{"verbatim":"subechinata","normalized":"subechinata","wordType":"SPECIES","start":11,"end":22},{"verbatim":"Canu","normalized":"Canu","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"Bassler","normalized":"Bassler","wordType":"AUTHOR_WORD","start":30,"end":37},{"verbatim":"1920","normalized":"1920","wordType":"YEAR","start":39,"end":43}],"id":"34e075be-fee2-509b-b08b-e024bd2dbd6c","parserVersion":"test_version"}
```

Name: Porina reussi Meneghini in De Amicis, 1885 vide Neviani (1900)

Canonical: Porina reussi

Authorship: Meneghini ex De Amicis 1885

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Porina reussi Meneghini in De Amicis, 1885 vide Neviani (1900)","normalized":"Porina reussi Meneghini ex De Amicis 1885","canonical":{"stemmed":"Porina reuss","simple":"Porina reussi","full":"Porina reussi"},"cardinality":2,"authorship":{"verbatim":"Meneghini in De Amicis, 1885","normalized":"Meneghini ex De Amicis 1885","year":"1885","authors":["Meneghini","De Amicis"],"originalAuth":{"authors":["Meneghini"],"exAuthors":{"authors":["De Amicis"],"year":{"year":"1885"}}}},"tail":" vide Neviani (1900)","details":{"species":{"genus":"Porina","species":"reussi","authorship":{"verbatim":"Meneghini in De Amicis, 1885","normalized":"Meneghini ex De Amicis 1885","year":"1885","authors":["Meneghini","De Amicis"],"originalAuth":{"authors":["Meneghini"],"exAuthors":{"authors":["De Amicis"],"year":{"year":"1885"}}}}}},"words":[{"verbatim":"Porina","normalized":"Porina","wordType":"GENUS","start":0,"end":6},{"verbatim":"reussi","normalized":"reussi","wordType":"SPECIES","start":7,"end":13},{"verbatim":"Meneghini","normalized":"Meneghini","wordType":"AUTHOR_WORD","start":14,"end":23},{"verbatim":"De","normalized":"De","wordType":"AUTHOR_WORD","start":27,"end":29},{"verbatim":"Amicis","normalized":"Amicis","wordType":"AUTHOR_WORD","start":30,"end":36},{"verbatim":"1885","normalized":"1885","wordType":"YEAR","start":38,"end":42}],"id":"e2a85725-9ffb-5e1e-9bdc-9f34648ef1b6","parserVersion":"test_version"}
```

### Abbreviated words after a name

Name: Graphis scripta L. a.b pulverulenta

Canonical: Graphis scripta

Authorship: L.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Graphis scripta L. a.b pulverulenta","normalized":"Graphis scripta L.","canonical":{"stemmed":"Graphis script","simple":"Graphis scripta","full":"Graphis scripta"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" a.b pulverulenta","details":{"species":{"genus":"Graphis","species":"scripta","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"words":[{"verbatim":"Graphis","normalized":"Graphis","wordType":"GENUS","start":0,"end":7},{"verbatim":"scripta","normalized":"scripta","wordType":"SPECIES","start":8,"end":15},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":16,"end":18}],"id":"ecb4751f-7d9e-5868-8ef7-c96f6ef07f2d","parserVersion":"test_version"}
```

Name: Cetraria iberica a.crespo & barreno

Canonical: Cetraria iberica

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Cetraria iberica a.crespo \u0026 barreno","normalized":"Cetraria iberica","canonical":{"stemmed":"Cetraria iberic","simple":"Cetraria iberica","full":"Cetraria iberica"},"cardinality":2,"tail":" a.crespo \u0026 barreno","details":{"species":{"genus":"Cetraria","species":"iberica"}},"words":[{"verbatim":"Cetraria","normalized":"Cetraria","wordType":"GENUS","start":0,"end":8},{"verbatim":"iberica","normalized":"iberica","wordType":"SPECIES","start":9,"end":16}],"id":"233626eb-645c-5ca0-bb8b-6f410a078a85","parserVersion":"test_version"}
```

Name: Lecanora achariana a.l.sm.

Canonical: Lecanora achariana

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lecanora achariana a.l.sm.","normalized":"Lecanora achariana","canonical":{"stemmed":"Lecanora acharian","simple":"Lecanora achariana","full":"Lecanora achariana"},"cardinality":2,"tail":" a.l.sm.","details":{"species":{"genus":"Lecanora","species":"achariana"}},"words":[{"verbatim":"Lecanora","normalized":"Lecanora","wordType":"GENUS","start":0,"end":8},{"verbatim":"achariana","normalized":"achariana","wordType":"SPECIES","start":9,"end":18}],"id":"4393f813-14e9-5a26-aab0-bf7686463c6a","parserVersion":"test_version"}
```

Name: Arthrosporum populorum a.massal.

Canonical: Arthrosporum populorum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arthrosporum populorum a.massal.","normalized":"Arthrosporum populorum","canonical":{"stemmed":"Arthrosporum populor","simple":"Arthrosporum populorum","full":"Arthrosporum populorum"},"cardinality":2,"tail":" a.massal.","details":{"species":{"genus":"Arthrosporum","species":"populorum"}},"words":[{"verbatim":"Arthrosporum","normalized":"Arthrosporum","wordType":"GENUS","start":0,"end":12},{"verbatim":"populorum","normalized":"populorum","wordType":"SPECIES","start":13,"end":22}],"id":"88db792d-7061-512d-9275-b7fe81493665","parserVersion":"test_version"}
```

Name: Eletica laeviceps ab.lateapicalis Pic

Canonical: Eletica laeviceps

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Eletica laeviceps ab.lateapicalis Pic","normalized":"Eletica laeviceps","canonical":{"stemmed":"Eletica laeuiceps","simple":"Eletica laeviceps","full":"Eletica laeviceps"},"cardinality":2,"tail":" ab.lateapicalis Pic","details":{"species":{"genus":"Eletica","species":"laeviceps"}},"words":[{"verbatim":"Eletica","normalized":"Eletica","wordType":"GENUS","start":0,"end":7},{"verbatim":"laeviceps","normalized":"laeviceps","wordType":"SPECIES","start":8,"end":17}],"id":"12389c9a-7aaf-56d1-8b8a-dffd4b74c58f","parserVersion":"test_version"}
```

<!--
Epithets with a whitespace  (rare, only ~50 cases)<
TODO Donatia novae zelandiae Hook.f.
TODO Donatia novae-zelandiae Hook.f.
-->

### Epithets starting with numeric value (not allowed anymore)

Name: Acanthoderes 4-gibbus RILEY Charles Valentine, 1880

Canonical: Acanthoderes quadrigibbus

Authorship: Riley Charles Valentine 1880

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"},{"quality":2,"warning":"Author in upper case"}],"verbatim":"Acanthoderes 4-gibbus RILEY Charles Valentine, 1880","normalized":"Acanthoderes quadrigibbus Riley Charles Valentine 1880","canonical":{"stemmed":"Acanthoderes quadrigibb","simple":"Acanthoderes quadrigibbus","full":"Acanthoderes quadrigibbus"},"cardinality":2,"authorship":{"verbatim":"RILEY Charles Valentine, 1880","normalized":"Riley Charles Valentine 1880","year":"1880","authors":["Riley Charles Valentine"],"originalAuth":{"authors":["Riley Charles Valentine"],"year":{"year":"1880"}}},"details":{"species":{"genus":"Acanthoderes","species":"quadrigibbus","authorship":{"verbatim":"RILEY Charles Valentine, 1880","normalized":"Riley Charles Valentine 1880","year":"1880","authors":["Riley Charles Valentine"],"originalAuth":{"authors":["Riley Charles Valentine"],"year":{"year":"1880"}}}}},"words":[{"verbatim":"Acanthoderes","normalized":"Acanthoderes","wordType":"GENUS","start":0,"end":12},{"verbatim":"4-gibbus","normalized":"quadrigibbus","wordType":"SPECIES","start":13,"end":21},{"verbatim":"RILEY","normalized":"Riley","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"Charles","normalized":"Charles","wordType":"AUTHOR_WORD","start":28,"end":35},{"verbatim":"Valentine","normalized":"Valentine","wordType":"AUTHOR_WORD","start":36,"end":45},{"verbatim":"1880","normalized":"1880","wordType":"YEAR","start":47,"end":51}],"id":"90bb5882-b093-586d-881a-aeabc55f248b","parserVersion":"test_version"}
```

Name: Acrosoma 12-spinosa Keyserling, 1892

Canonical: Acrosoma duodecimspinosa

Authorship: Keyserling 1892

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Acrosoma 12-spinosa Keyserling, 1892","normalized":"Acrosoma duodecimspinosa Keyserling 1892","canonical":{"stemmed":"Acrosoma duodecimspinos","simple":"Acrosoma duodecimspinosa","full":"Acrosoma duodecimspinosa"},"cardinality":2,"authorship":{"verbatim":"Keyserling, 1892","normalized":"Keyserling 1892","year":"1892","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1892"}}},"details":{"species":{"genus":"Acrosoma","species":"duodecimspinosa","authorship":{"verbatim":"Keyserling, 1892","normalized":"Keyserling 1892","year":"1892","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1892"}}}}},"words":[{"verbatim":"Acrosoma","normalized":"Acrosoma","wordType":"GENUS","start":0,"end":8},{"verbatim":"12-spinosa","normalized":"duodecimspinosa","wordType":"SPECIES","start":9,"end":19},{"verbatim":"Keyserling","normalized":"Keyserling","wordType":"AUTHOR_WORD","start":20,"end":30},{"verbatim":"1892","normalized":"1892","wordType":"YEAR","start":32,"end":36}],"id":"d789c68a-4e40-59d8-a763-3ebadac6fdeb","parserVersion":"test_version"}
```

Name: Canuleius 24-spinosus Redtenbacher, 1906

Canonical: Canuleius vigintiquatuorspinosus

Authorship: Redtenbacher 1906

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Canuleius 24-spinosus Redtenbacher, 1906","normalized":"Canuleius vigintiquatuorspinosus Redtenbacher 1906","canonical":{"stemmed":"Canuleius uigintiquatuorspinos","simple":"Canuleius vigintiquatuorspinosus","full":"Canuleius vigintiquatuorspinosus"},"cardinality":2,"authorship":{"verbatim":"Redtenbacher, 1906","normalized":"Redtenbacher 1906","year":"1906","authors":["Redtenbacher"],"originalAuth":{"authors":["Redtenbacher"],"year":{"year":"1906"}}},"details":{"species":{"genus":"Canuleius","species":"vigintiquatuorspinosus","authorship":{"verbatim":"Redtenbacher, 1906","normalized":"Redtenbacher 1906","year":"1906","authors":["Redtenbacher"],"originalAuth":{"authors":["Redtenbacher"],"year":{"year":"1906"}}}}},"words":[{"verbatim":"Canuleius","normalized":"Canuleius","wordType":"GENUS","start":0,"end":9},{"verbatim":"24-spinosus","normalized":"vigintiquatuorspinosus","wordType":"SPECIES","start":10,"end":21},{"verbatim":"Redtenbacher","normalized":"Redtenbacher","wordType":"AUTHOR_WORD","start":22,"end":34},{"verbatim":"1906","normalized":"1906","wordType":"YEAR","start":36,"end":40}],"id":"6dbf79a3-89dd-55ee-aa7d-6394c226cb02","parserVersion":"test_version"}
```

<!-- numeric prefix cannot be more than 2 digits long -->
Name: Canuleius 777-spinosus Redtenbacher, 1906

Canonical: Canuleius

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Canuleius 777-spinosus Redtenbacher, 1906","normalized":"Canuleius","canonical":{"stemmed":"Canuleius","simple":"Canuleius","full":"Canuleius"},"cardinality":1,"tail":" 777-spinosus Redtenbacher, 1906","details":{"uninomial":{"uninomial":"Canuleius"}},"words":[{"verbatim":"Canuleius","normalized":"Canuleius","wordType":"UNINOMIAL","start":0,"end":9}],"id":"40a1b1cd-0437-5ed8-82bf-8bea169cb8b1","parserVersion":"test_version"}
```

Name: Rhynchophorus 13punctatus Herbst, J.F.W., 1795

Canonical: Rhynchophorus tredecimpunctatus

Authorship: Herbst & J. F. W. 1795

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Rhynchophorus 13punctatus Herbst, J.F.W., 1795","normalized":"Rhynchophorus tredecimpunctatus Herbst \u0026 J. F. W. 1795","canonical":{"stemmed":"Rhynchophorus tredecimpunctat","simple":"Rhynchophorus tredecimpunctatus","full":"Rhynchophorus tredecimpunctatus"},"cardinality":2,"authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}},"details":{"species":{"genus":"Rhynchophorus","species":"tredecimpunctatus","authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}}}},"words":[{"verbatim":"Rhynchophorus","normalized":"Rhynchophorus","wordType":"GENUS","start":0,"end":13},{"verbatim":"13punctatus","normalized":"tredecimpunctatus","wordType":"SPECIES","start":14,"end":25},{"verbatim":"Herbst","normalized":"Herbst","wordType":"AUTHOR_WORD","start":26,"end":32},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":34,"end":36},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":36,"end":38},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":38,"end":40},{"verbatim":"1795","normalized":"1795","wordType":"YEAR","start":42,"end":46}],"id":"8724e04d-a1a0-5b5e-9c0e-1c0f586507d6","parserVersion":"test_version"}
```

Name: Rhynchophorus 13.punctatus Herbst, J.F.W., 1795

Canonical: Rhynchophorus tredecimpunctatus

Authorship: Herbst & J. F. W. 1795

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Rhynchophorus 13.punctatus Herbst, J.F.W., 1795","normalized":"Rhynchophorus tredecimpunctatus Herbst \u0026 J. F. W. 1795","canonical":{"stemmed":"Rhynchophorus tredecimpunctat","simple":"Rhynchophorus tredecimpunctatus","full":"Rhynchophorus tredecimpunctatus"},"cardinality":2,"authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}},"details":{"species":{"genus":"Rhynchophorus","species":"tredecimpunctatus","authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}}}},"words":[{"verbatim":"Rhynchophorus","normalized":"Rhynchophorus","wordType":"GENUS","start":0,"end":13},{"verbatim":"13.punctatus","normalized":"tredecimpunctatus","wordType":"SPECIES","start":14,"end":26},{"verbatim":"Herbst","normalized":"Herbst","wordType":"AUTHOR_WORD","start":27,"end":33},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":35,"end":37},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"1795","normalized":"1795","wordType":"YEAR","start":43,"end":47}],"id":"590b3805-23bc-5a94-a7ca-ea89dcfb5ed1","parserVersion":"test_version"}
```

### Non-ASCII UTF-8 characters in a name

Name: Seleuca chûjôi Voss, 1957

Canonical: Seleuca chujoi

Authorship: Voss 1957

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Seleuca chûjôi Voss, 1957","normalized":"Seleuca chujoi Voss 1957","canonical":{"stemmed":"Seleuca chuio","simple":"Seleuca chujoi","full":"Seleuca chujoi"},"cardinality":2,"authorship":{"verbatim":"Voss, 1957","normalized":"Voss 1957","year":"1957","authors":["Voss"],"originalAuth":{"authors":["Voss"],"year":{"year":"1957"}}},"details":{"species":{"genus":"Seleuca","species":"chujoi","authorship":{"verbatim":"Voss, 1957","normalized":"Voss 1957","year":"1957","authors":["Voss"],"originalAuth":{"authors":["Voss"],"year":{"year":"1957"}}}}},"words":[{"verbatim":"Seleuca","normalized":"Seleuca","wordType":"GENUS","start":0,"end":7},{"verbatim":"chûjôi","normalized":"chujoi","wordType":"SPECIES","start":8,"end":14},{"verbatim":"Voss","normalized":"Voss","wordType":"AUTHOR_WORD","start":15,"end":19},{"verbatim":"1957","normalized":"1957","wordType":"YEAR","start":21,"end":25}],"id":"b6244666-9125-5473-8755-bb35ebbea769","parserVersion":"test_version"}
```

Name: Pleurotus ëous (Berk.) Sacc. 1887

Canonical: Pleurotus eous

Authorship: (Berk.) Sacc. 1887

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pleurotus ëous (Berk.) Sacc. 1887","normalized":"Pleurotus eous (Berk.) Sacc. 1887","canonical":{"stemmed":"Pleurotus eo","simple":"Pleurotus eous","full":"Pleurotus eous"},"cardinality":2,"authorship":{"verbatim":"(Berk.) Sacc. 1887","normalized":"(Berk.) Sacc. 1887","authors":["Berk.","Sacc."],"originalAuth":{"authors":["Berk."]},"combinationAuth":{"authors":["Sacc."],"year":{"year":"1887"}}},"details":{"species":{"genus":"Pleurotus","species":"eous","authorship":{"verbatim":"(Berk.) Sacc. 1887","normalized":"(Berk.) Sacc. 1887","authors":["Berk.","Sacc."],"originalAuth":{"authors":["Berk."]},"combinationAuth":{"authors":["Sacc."],"year":{"year":"1887"}}}}},"words":[{"verbatim":"Pleurotus","normalized":"Pleurotus","wordType":"GENUS","start":0,"end":9},{"verbatim":"ëous","normalized":"eous","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Berk.","normalized":"Berk.","wordType":"AUTHOR_WORD","start":16,"end":21},{"verbatim":"Sacc.","normalized":"Sacc.","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"1887","normalized":"1887","wordType":"YEAR","start":29,"end":33}],"id":"fe8c9a43-3480-5598-891d-e2a864781d13","parserVersion":"test_version"}
```

Name: Sténométope laevissimus Bibron 1855

Canonical: Stenometope laevissimus

Authorship: Bibron 1855

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Sténométope laevissimus Bibron 1855","normalized":"Stenometope laevissimus Bibron 1855","canonical":{"stemmed":"Stenometope laeuissim","simple":"Stenometope laevissimus","full":"Stenometope laevissimus"},"cardinality":2,"authorship":{"verbatim":"Bibron 1855","normalized":"Bibron 1855","year":"1855","authors":["Bibron"],"originalAuth":{"authors":["Bibron"],"year":{"year":"1855"}}},"details":{"species":{"genus":"Stenometope","species":"laevissimus","authorship":{"verbatim":"Bibron 1855","normalized":"Bibron 1855","year":"1855","authors":["Bibron"],"originalAuth":{"authors":["Bibron"],"year":{"year":"1855"}}}}},"words":[{"verbatim":"Sténométope","normalized":"Stenometope","wordType":"GENUS","start":0,"end":11},{"verbatim":"laevissimus","normalized":"laevissimus","wordType":"SPECIES","start":12,"end":23},{"verbatim":"Bibron","normalized":"Bibron","wordType":"AUTHOR_WORD","start":24,"end":30},{"verbatim":"1855","normalized":"1855","wordType":"YEAR","start":31,"end":35}],"id":"363ea9fc-ac47-50e5-ae4b-1bfb104a8e34","parserVersion":"test_version"}
```

Name: Choriozopella trägårdhi Lawrence, 1947

Canonical: Choriozopella traegaordhi

Authorship: Lawrence 1947

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Choriozopella trägårdhi Lawrence, 1947","normalized":"Choriozopella traegaordhi Lawrence 1947","canonical":{"stemmed":"Choriozopella traegaordh","simple":"Choriozopella traegaordhi","full":"Choriozopella traegaordhi"},"cardinality":2,"authorship":{"verbatim":"Lawrence, 1947","normalized":"Lawrence 1947","year":"1947","authors":["Lawrence"],"originalAuth":{"authors":["Lawrence"],"year":{"year":"1947"}}},"details":{"species":{"genus":"Choriozopella","species":"traegaordhi","authorship":{"verbatim":"Lawrence, 1947","normalized":"Lawrence 1947","year":"1947","authors":["Lawrence"],"originalAuth":{"authors":["Lawrence"],"year":{"year":"1947"}}}}},"words":[{"verbatim":"Choriozopella","normalized":"Choriozopella","wordType":"GENUS","start":0,"end":13},{"verbatim":"trägårdhi","normalized":"traegaordhi","wordType":"SPECIES","start":14,"end":23},{"verbatim":"Lawrence","normalized":"Lawrence","wordType":"AUTHOR_WORD","start":24,"end":32},{"verbatim":"1947","normalized":"1947","wordType":"YEAR","start":34,"end":38}],"id":"3d02292a-3526-5364-96c9-f73738b9d2fa","parserVersion":"test_version"}
```

Name: Isoëtes asplundii H. P. Fuchs

Canonical: Isoetes asplundii

Authorship: H. P. Fuchs

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Isoëtes asplundii H. P. Fuchs","normalized":"Isoetes asplundii H. P. Fuchs","canonical":{"stemmed":"Isoetes asplund","simple":"Isoetes asplundii","full":"Isoetes asplundii"},"cardinality":2,"authorship":{"verbatim":"H. P. Fuchs","normalized":"H. P. Fuchs","authors":["H. P. Fuchs"],"originalAuth":{"authors":["H. P. Fuchs"]}},"details":{"species":{"genus":"Isoetes","species":"asplundii","authorship":{"verbatim":"H. P. Fuchs","normalized":"H. P. Fuchs","authors":["H. P. Fuchs"],"originalAuth":{"authors":["H. P. Fuchs"]}}}},"words":[{"verbatim":"Isoëtes","normalized":"Isoetes","wordType":"GENUS","start":0,"end":7},{"verbatim":"asplundii","normalized":"asplundii","wordType":"SPECIES","start":8,"end":17},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":18,"end":20},{"verbatim":"P.","normalized":"P.","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"Fuchs","normalized":"Fuchs","wordType":"AUTHOR_WORD","start":24,"end":29}],"id":"8d713775-782a-5083-92a1-ddaf4af9d785","parserVersion":"test_version"}
```

Name: Cerambyx thomæ GMELIN J. F., 1790

Canonical: Cerambyx thomae

Authorship: Gmelin J. F. 1790

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Cerambyx thomæ GMELIN J. F., 1790","normalized":"Cerambyx thomae Gmelin J. F. 1790","canonical":{"stemmed":"Cerambyx thom","simple":"Cerambyx thomae","full":"Cerambyx thomae"},"cardinality":2,"authorship":{"verbatim":"GMELIN J. F., 1790","normalized":"Gmelin J. F. 1790","year":"1790","authors":["Gmelin J. F."],"originalAuth":{"authors":["Gmelin J. F."],"year":{"year":"1790"}}},"details":{"species":{"genus":"Cerambyx","species":"thomae","authorship":{"verbatim":"GMELIN J. F., 1790","normalized":"Gmelin J. F. 1790","year":"1790","authors":["Gmelin J. F."],"originalAuth":{"authors":["Gmelin J. F."],"year":{"year":"1790"}}}}},"words":[{"verbatim":"Cerambyx","normalized":"Cerambyx","wordType":"GENUS","start":0,"end":8},{"verbatim":"thomæ","normalized":"thomae","wordType":"SPECIES","start":9,"end":14},{"verbatim":"GMELIN","normalized":"Gmelin","wordType":"AUTHOR_WORD","start":15,"end":21},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"1790","normalized":"1790","wordType":"YEAR","start":29,"end":33}],"id":"f9689237-693f-5d6e-b62e-b1622214863e","parserVersion":"test_version"}
```

Name: Campethera cailliautii fülleborni

Canonical: Campethera cailliautii fuelleborni

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Campethera cailliautii fülleborni","normalized":"Campethera cailliautii fuelleborni","canonical":{"stemmed":"Campethera cailliaut fuelleborn","simple":"Campethera cailliautii fuelleborni","full":"Campethera cailliautii fuelleborni"},"cardinality":3,"details":{"infraspecies":{"genus":"Campethera","species":"cailliautii","infraspecies":[{"value":"fuelleborni"}]}},"words":[{"verbatim":"Campethera","normalized":"Campethera","wordType":"GENUS","start":0,"end":10},{"verbatim":"cailliautii","normalized":"cailliautii","wordType":"SPECIES","start":11,"end":22},{"verbatim":"fülleborni","normalized":"fuelleborni","wordType":"INFRASPECIES","start":23,"end":33}],"id":"6a47e9c0-908f-5141-be93-76490af08606","parserVersion":"test_version"}
```

Name: Östrupia Heiden ex Hustedt, 1935

Canonical: Oestrupia

Authorship: Heiden ex Hustedt 1935

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Östrupia Heiden ex Hustedt, 1935","normalized":"Oestrupia Heiden ex Hustedt 1935","canonical":{"stemmed":"Oestrupia","simple":"Oestrupia","full":"Oestrupia"},"cardinality":1,"authorship":{"verbatim":"Heiden ex Hustedt, 1935","normalized":"Heiden ex Hustedt 1935","year":"1935","authors":["Heiden","Hustedt"],"originalAuth":{"authors":["Heiden"],"exAuthors":{"authors":["Hustedt"],"year":{"year":"1935"}}}},"details":{"uninomial":{"uninomial":"Oestrupia","authorship":{"verbatim":"Heiden ex Hustedt, 1935","normalized":"Heiden ex Hustedt 1935","year":"1935","authors":["Heiden","Hustedt"],"originalAuth":{"authors":["Heiden"],"exAuthors":{"authors":["Hustedt"],"year":{"year":"1935"}}}}}},"words":[{"verbatim":"Östrupia","normalized":"Oestrupia","wordType":"UNINOMIAL","start":0,"end":8},{"verbatim":"Heiden","normalized":"Heiden","wordType":"AUTHOR_WORD","start":9,"end":15},{"verbatim":"Hustedt","normalized":"Hustedt","wordType":"AUTHOR_WORD","start":19,"end":26},{"verbatim":"1935","normalized":"1935","wordType":"YEAR","start":28,"end":32}],"id":"940aba5b-2334-5846-98ba-ce29c7305734","parserVersion":"test_version"}
```

### Epithets with an apostrophe

Name: Solanum tuberosum f. wila-k'oyu Ochoa

Canonical: Solanum tuberosum f. wila-koyu

Authorship: Ochoa

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Solanum tuberosum f. wila-k'oyu Ochoa","normalized":"Solanum tuberosum f. wila-koyu Ochoa","canonical":{"stemmed":"Solanum tuberos wila-koy","simple":"Solanum tuberosum wila-koyu","full":"Solanum tuberosum f. wila-koyu"},"cardinality":3,"authorship":{"verbatim":"Ochoa","normalized":"Ochoa","authors":["Ochoa"],"originalAuth":{"authors":["Ochoa"]}},"details":{"infraspecies":{"genus":"Solanum","species":"tuberosum","infraspecies":[{"value":"wila-koyu","rank":"f.","authorship":{"verbatim":"Ochoa","normalized":"Ochoa","authors":["Ochoa"],"originalAuth":{"authors":["Ochoa"]}}}]}},"words":[{"verbatim":"Solanum","normalized":"Solanum","wordType":"GENUS","start":0,"end":7},{"verbatim":"tuberosum","normalized":"tuberosum","wordType":"SPECIES","start":8,"end":17},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":18,"end":20},{"verbatim":"wila-k'oyu","normalized":"wila-koyu","wordType":"INFRASPECIES","start":21,"end":31},{"verbatim":"Ochoa","normalized":"Ochoa","wordType":"AUTHOR_WORD","start":32,"end":37}],"id":"b45b0e75-d1d0-53f2-ab80-f5a99d24a385","parserVersion":"test_version"}
```

Name: Junellia o'donelli Moldenke, 1946

Canonical: Junellia odonelli

Authorship: Moldenke 1946

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Junellia o'donelli Moldenke, 1946","normalized":"Junellia odonelli Moldenke 1946","canonical":{"stemmed":"Junellia odonell","simple":"Junellia odonelli","full":"Junellia odonelli"},"cardinality":2,"authorship":{"verbatim":"Moldenke, 1946","normalized":"Moldenke 1946","year":"1946","authors":["Moldenke"],"originalAuth":{"authors":["Moldenke"],"year":{"year":"1946"}}},"details":{"species":{"genus":"Junellia","species":"odonelli","authorship":{"verbatim":"Moldenke, 1946","normalized":"Moldenke 1946","year":"1946","authors":["Moldenke"],"originalAuth":{"authors":["Moldenke"],"year":{"year":"1946"}}}}},"words":[{"verbatim":"Junellia","normalized":"Junellia","wordType":"GENUS","start":0,"end":8},{"verbatim":"o'donelli","normalized":"odonelli","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Moldenke","normalized":"Moldenke","wordType":"AUTHOR_WORD","start":19,"end":27},{"verbatim":"1946","normalized":"1946","wordType":"YEAR","start":29,"end":33}],"id":"e39a2d98-6ab2-5fb3-9aae-c48aa86c6026","parserVersion":"test_version"}
```

Name: Trophon d'orbignyi Carcelles, 1946

Canonical: Trophon dorbignyi

Authorship: Carcelles 1946

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Trophon d'orbignyi Carcelles, 1946","normalized":"Trophon dorbignyi Carcelles 1946","canonical":{"stemmed":"Trophon dorbigny","simple":"Trophon dorbignyi","full":"Trophon dorbignyi"},"cardinality":2,"authorship":{"verbatim":"Carcelles, 1946","normalized":"Carcelles 1946","year":"1946","authors":["Carcelles"],"originalAuth":{"authors":["Carcelles"],"year":{"year":"1946"}}},"details":{"species":{"genus":"Trophon","species":"dorbignyi","authorship":{"verbatim":"Carcelles, 1946","normalized":"Carcelles 1946","year":"1946","authors":["Carcelles"],"originalAuth":{"authors":["Carcelles"],"year":{"year":"1946"}}}}},"words":[{"verbatim":"Trophon","normalized":"Trophon","wordType":"GENUS","start":0,"end":7},{"verbatim":"d'orbignyi","normalized":"dorbignyi","wordType":"SPECIES","start":8,"end":18},{"verbatim":"Carcelles","normalized":"Carcelles","wordType":"AUTHOR_WORD","start":19,"end":28},{"verbatim":"1946","normalized":"1946","wordType":"YEAR","start":30,"end":34}],"id":"935d4414-05d4-5c16-be30-466f6144b666","parserVersion":"test_version"}
```

Name: Phrynosoma m’callii

Canonical: Phrynosoma mcallii

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"},{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Phrynosoma m’callii","normalized":"Phrynosoma mcallii","canonical":{"stemmed":"Phrynosoma mcall","simple":"Phrynosoma mcallii","full":"Phrynosoma mcallii"},"cardinality":2,"details":{"species":{"genus":"Phrynosoma","species":"mcallii"}},"words":[{"verbatim":"Phrynosoma","normalized":"Phrynosoma","wordType":"GENUS","start":0,"end":10},{"verbatim":"m’callii","normalized":"mcallii","wordType":"SPECIES","start":11,"end":19}],"id":"7907df5c-50f2-532c-a8fe-e5b75f924f73","parserVersion":"test_version"}
```

Name: Arca m'coyi Tenison-Woods, 1878

Canonical: Arca mcoyi

Authorship: Tenison-Woods 1878

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Arca m'coyi Tenison-Woods, 1878","normalized":"Arca mcoyi Tenison-Woods 1878","canonical":{"stemmed":"Arca mcoy","simple":"Arca mcoyi","full":"Arca mcoyi"},"cardinality":2,"authorship":{"verbatim":"Tenison-Woods, 1878","normalized":"Tenison-Woods 1878","year":"1878","authors":["Tenison-Woods"],"originalAuth":{"authors":["Tenison-Woods"],"year":{"year":"1878"}}},"details":{"species":{"genus":"Arca","species":"mcoyi","authorship":{"verbatim":"Tenison-Woods, 1878","normalized":"Tenison-Woods 1878","year":"1878","authors":["Tenison-Woods"],"originalAuth":{"authors":["Tenison-Woods"],"year":{"year":"1878"}}}}},"words":[{"verbatim":"Arca","normalized":"Arca","wordType":"GENUS","start":0,"end":4},{"verbatim":"m'coyi","normalized":"mcoyi","wordType":"SPECIES","start":5,"end":11},{"verbatim":"Tenison-Woods","normalized":"Tenison-Woods","wordType":"AUTHOR_WORD","start":12,"end":25},{"verbatim":"1878","normalized":"1878","wordType":"YEAR","start":27,"end":31}],"id":"fa855178-bdde-5ebf-b6b1-c1a1aa60bffa","parserVersion":"test_version"}
```

Name: Nucula m'andrewii Hanley, 1860

Canonical: Nucula mandrewii

Authorship: Hanley 1860

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Nucula m'andrewii Hanley, 1860","normalized":"Nucula mandrewii Hanley 1860","canonical":{"stemmed":"Nucula mandrew","simple":"Nucula mandrewii","full":"Nucula mandrewii"},"cardinality":2,"authorship":{"verbatim":"Hanley, 1860","normalized":"Hanley 1860","year":"1860","authors":["Hanley"],"originalAuth":{"authors":["Hanley"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Nucula","species":"mandrewii","authorship":{"verbatim":"Hanley, 1860","normalized":"Hanley 1860","year":"1860","authors":["Hanley"],"originalAuth":{"authors":["Hanley"],"year":{"year":"1860"}}}}},"words":[{"verbatim":"Nucula","normalized":"Nucula","wordType":"GENUS","start":0,"end":6},{"verbatim":"m'andrewii","normalized":"mandrewii","wordType":"SPECIES","start":7,"end":17},{"verbatim":"Hanley","normalized":"Hanley","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"1860","normalized":"1860","wordType":"YEAR","start":26,"end":30}],"id":"8bbc3b0e-149d-5ede-9f12-b516b085da9d","parserVersion":"test_version"}
```

Name: Eristalis l'herminierii Macquart

Canonical: Eristalis lherminierii

Authorship: Macquart

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Eristalis l'herminierii Macquart","normalized":"Eristalis lherminierii Macquart","canonical":{"stemmed":"Eristalis lherminier","simple":"Eristalis lherminierii","full":"Eristalis lherminierii"},"cardinality":2,"authorship":{"verbatim":"Macquart","normalized":"Macquart","authors":["Macquart"],"originalAuth":{"authors":["Macquart"]}},"details":{"species":{"genus":"Eristalis","species":"lherminierii","authorship":{"verbatim":"Macquart","normalized":"Macquart","authors":["Macquart"],"originalAuth":{"authors":["Macquart"]}}}},"words":[{"verbatim":"Eristalis","normalized":"Eristalis","wordType":"GENUS","start":0,"end":9},{"verbatim":"l'herminierii","normalized":"lherminierii","wordType":"SPECIES","start":10,"end":23},{"verbatim":"Macquart","normalized":"Macquart","wordType":"AUTHOR_WORD","start":24,"end":32}],"id":"f7ccb013-ad48-5424-9c26-01657275de9a","parserVersion":"test_version"}
```

Name: Odynerus o'neili Cameron

Canonical: Odynerus oneili

Authorship: Cameron

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Odynerus o'neili Cameron","normalized":"Odynerus oneili Cameron","canonical":{"stemmed":"Odynerus oneil","simple":"Odynerus oneili","full":"Odynerus oneili"},"cardinality":2,"authorship":{"verbatim":"Cameron","normalized":"Cameron","authors":["Cameron"],"originalAuth":{"authors":["Cameron"]}},"details":{"species":{"genus":"Odynerus","species":"oneili","authorship":{"verbatim":"Cameron","normalized":"Cameron","authors":["Cameron"],"originalAuth":{"authors":["Cameron"]}}}},"words":[{"verbatim":"Odynerus","normalized":"Odynerus","wordType":"GENUS","start":0,"end":8},{"verbatim":"o'neili","normalized":"oneili","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Cameron","normalized":"Cameron","wordType":"AUTHOR_WORD","start":17,"end":24}],"id":"39218b39-39f9-5f0d-917a-d5e57301d91c","parserVersion":"test_version"}
```

Name: Serjania meridionalis Cambess. var. o'donelli F.A. Barkley

Canonical: Serjania meridionalis var. odonelli

Authorship: F. A. Barkley

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Serjania meridionalis Cambess. var. o'donelli F.A. Barkley","normalized":"Serjania meridionalis Cambess. var. odonelli F. A. Barkley","canonical":{"stemmed":"Serjania meridional odonell","simple":"Serjania meridionalis odonelli","full":"Serjania meridionalis var. odonelli"},"cardinality":3,"authorship":{"verbatim":"F.A. Barkley","normalized":"F. A. Barkley","authors":["F. A. Barkley"],"originalAuth":{"authors":["F. A. Barkley"]}},"details":{"infraspecies":{"genus":"Serjania","species":"meridionalis","authorship":{"verbatim":"Cambess.","normalized":"Cambess.","authors":["Cambess."],"originalAuth":{"authors":["Cambess."]}},"infraspecies":[{"value":"odonelli","rank":"var.","authorship":{"verbatim":"F.A. Barkley","normalized":"F. A. Barkley","authors":["F. A. Barkley"],"originalAuth":{"authors":["F. A. Barkley"]}}}]}},"words":[{"verbatim":"Serjania","normalized":"Serjania","wordType":"GENUS","start":0,"end":8},{"verbatim":"meridionalis","normalized":"meridionalis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Cambess.","normalized":"Cambess.","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":31,"end":35},{"verbatim":"o'donelli","normalized":"odonelli","wordType":"INFRASPECIES","start":36,"end":45},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":46,"end":48},{"verbatim":"A.","normalized":"A.","wordType":"AUTHOR_WORD","start":48,"end":50},{"verbatim":"Barkley","normalized":"Barkley","wordType":"AUTHOR_WORD","start":51,"end":58}],"id":"019a8f2c-279d-5211-9bfb-5f288795ed73","parserVersion":"test_version"}
```

### Authors with an apostrophe

Name: Galega officinalis (L.) L´Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.

Canonical: Galega officinalis subsp. mackayana var. petiolata

Authorship: (È. Neé) Brüch.

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Galega officinalis (L.) L´Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","normalized":"Galega officinalis (L.) L'Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","canonical":{"stemmed":"Galega officinal mackayan petiolat","simple":"Galega officinalis mackayana petiolata","full":"Galega officinalis subsp. mackayana var. petiolata"},"cardinality":4,"authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}},"details":{"infraspecies":{"genus":"Galega","species":"officinalis","authorship":{"verbatim":"(L.) L´Hèr.","normalized":"(L.) L'Hèr.","authors":["L.","L'Hèr."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Hèr."]}},"infraspecies":[{"value":"mackayana","rank":"subsp.","authorship":{"verbatim":"(O'Flannagan) Mc Inley","normalized":"(O'Flannagan) Mc Inley","authors":["O'Flannagan","Mc Inley"],"originalAuth":{"authors":["O'Flannagan"]},"combinationAuth":{"authors":["Mc Inley"]}}},{"value":"petiolata","rank":"var.","authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}}}]}},"words":[{"verbatim":"Galega","normalized":"Galega","wordType":"GENUS","start":0,"end":6},{"verbatim":"officinalis","normalized":"officinalis","wordType":"SPECIES","start":7,"end":18},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":20,"end":22},{"verbatim":"L´Hèr.","normalized":"L'Hèr.","wordType":"AUTHOR_WORD","start":24,"end":30},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":31,"end":37},{"verbatim":"mackayana","normalized":"mackayana","wordType":"INFRASPECIES","start":38,"end":47},{"verbatim":"O'Flannagan","normalized":"O'Flannagan","wordType":"AUTHOR_WORD","start":49,"end":60},{"verbatim":"Mc","normalized":"Mc","wordType":"AUTHOR_WORD","start":62,"end":64},{"verbatim":"Inley","normalized":"Inley","wordType":"AUTHOR_WORD","start":65,"end":70},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":71,"end":75},{"verbatim":"petiolata","normalized":"petiolata","wordType":"INFRASPECIES","start":76,"end":85},{"verbatim":"È.","normalized":"È.","wordType":"AUTHOR_WORD","start":87,"end":89},{"verbatim":"Neé","normalized":"Neé","wordType":"AUTHOR_WORD","start":90,"end":93},{"verbatim":"Brüch.","normalized":"Brüch.","wordType":"AUTHOR_WORD","start":95,"end":101}],"id":"9555468f-987c-5bc5-bfa2-2581f7c5d41c","parserVersion":"test_version"}
```

Name: Galega officinalis (L.) L`Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.

Canonical: Galega officinalis subsp. mackayana var. petiolata

Authorship: (È. Neé) Brüch.

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Galega officinalis (L.) L`Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","normalized":"Galega officinalis (L.) L'Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","canonical":{"stemmed":"Galega officinal mackayan petiolat","simple":"Galega officinalis mackayana petiolata","full":"Galega officinalis subsp. mackayana var. petiolata"},"cardinality":4,"authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}},"details":{"infraspecies":{"genus":"Galega","species":"officinalis","authorship":{"verbatim":"(L.) L`Hèr.","normalized":"(L.) L'Hèr.","authors":["L.","L'Hèr."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Hèr."]}},"infraspecies":[{"value":"mackayana","rank":"subsp.","authorship":{"verbatim":"(O'Flannagan) Mc Inley","normalized":"(O'Flannagan) Mc Inley","authors":["O'Flannagan","Mc Inley"],"originalAuth":{"authors":["O'Flannagan"]},"combinationAuth":{"authors":["Mc Inley"]}}},{"value":"petiolata","rank":"var.","authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}}}]}},"words":[{"verbatim":"Galega","normalized":"Galega","wordType":"GENUS","start":0,"end":6},{"verbatim":"officinalis","normalized":"officinalis","wordType":"SPECIES","start":7,"end":18},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":20,"end":22},{"verbatim":"L`Hèr.","normalized":"L'Hèr.","wordType":"AUTHOR_WORD","start":24,"end":30},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":31,"end":37},{"verbatim":"mackayana","normalized":"mackayana","wordType":"INFRASPECIES","start":38,"end":47},{"verbatim":"O'Flannagan","normalized":"O'Flannagan","wordType":"AUTHOR_WORD","start":49,"end":60},{"verbatim":"Mc","normalized":"Mc","wordType":"AUTHOR_WORD","start":62,"end":64},{"verbatim":"Inley","normalized":"Inley","wordType":"AUTHOR_WORD","start":65,"end":70},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":71,"end":75},{"verbatim":"petiolata","normalized":"petiolata","wordType":"INFRASPECIES","start":76,"end":85},{"verbatim":"È.","normalized":"È.","wordType":"AUTHOR_WORD","start":87,"end":89},{"verbatim":"Neé","normalized":"Neé","wordType":"AUTHOR_WORD","start":90,"end":93},{"verbatim":"Brüch.","normalized":"Brüch.","wordType":"AUTHOR_WORD","start":95,"end":101}],"id":"af46c9cc-a3be-507e-9690-349f0303fcd7","parserVersion":"test_version"}
```

Name: Galega officinalis (L.) L'Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.

Canonical: Galega officinalis subsp. mackayana var. petiolata

Authorship: (È. Neé) Brüch.

```json
{"parsed":true,"quality":1,"verbatim":"Galega officinalis (L.) L'Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","normalized":"Galega officinalis (L.) L'Hèr. subsp. mackayana (O'Flannagan) Mc Inley var. petiolata (È. Neé) Brüch.","canonical":{"stemmed":"Galega officinal mackayan petiolat","simple":"Galega officinalis mackayana petiolata","full":"Galega officinalis subsp. mackayana var. petiolata"},"cardinality":4,"authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}},"details":{"infraspecies":{"genus":"Galega","species":"officinalis","authorship":{"verbatim":"(L.) L'Hèr.","normalized":"(L.) L'Hèr.","authors":["L.","L'Hèr."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Hèr."]}},"infraspecies":[{"value":"mackayana","rank":"subsp.","authorship":{"verbatim":"(O'Flannagan) Mc Inley","normalized":"(O'Flannagan) Mc Inley","authors":["O'Flannagan","Mc Inley"],"originalAuth":{"authors":["O'Flannagan"]},"combinationAuth":{"authors":["Mc Inley"]}}},{"value":"petiolata","rank":"var.","authorship":{"verbatim":"(È. Neé) Brüch.","normalized":"(È. Neé) Brüch.","authors":["È. Neé","Brüch."],"originalAuth":{"authors":["È. Neé"]},"combinationAuth":{"authors":["Brüch."]}}}]}},"words":[{"verbatim":"Galega","normalized":"Galega","wordType":"GENUS","start":0,"end":6},{"verbatim":"officinalis","normalized":"officinalis","wordType":"SPECIES","start":7,"end":18},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":20,"end":22},{"verbatim":"L'Hèr.","normalized":"L'Hèr.","wordType":"AUTHOR_WORD","start":24,"end":30},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":31,"end":37},{"verbatim":"mackayana","normalized":"mackayana","wordType":"INFRASPECIES","start":38,"end":47},{"verbatim":"O'Flannagan","normalized":"O'Flannagan","wordType":"AUTHOR_WORD","start":49,"end":60},{"verbatim":"Mc","normalized":"Mc","wordType":"AUTHOR_WORD","start":62,"end":64},{"verbatim":"Inley","normalized":"Inley","wordType":"AUTHOR_WORD","start":65,"end":70},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":71,"end":75},{"verbatim":"petiolata","normalized":"petiolata","wordType":"INFRASPECIES","start":76,"end":85},{"verbatim":"È.","normalized":"È.","wordType":"AUTHOR_WORD","start":87,"end":89},{"verbatim":"Neé","normalized":"Neé","wordType":"AUTHOR_WORD","start":90,"end":93},{"verbatim":"Brüch.","normalized":"Brüch.","wordType":"AUTHOR_WORD","start":95,"end":101}],"id":"9d131412-69c9-52e2-a154-dbbfff9e5494","parserVersion":"test_version"}
```

### Digraph unicode characters

Name: Crisia romanica Zágoršek Silye & Szabó 2008

Canonical: Crisia romanica

Authorship:  Zágoršek Silye & Szabó 2008

```json
{"parsed":true,"quality":1,"verbatim":"Crisia romanica Zágoršek Silye \u0026 Szabó 2008","normalized":"Crisia romanica Zágoršek Silye \u0026 Szabó 2008","canonical":{"stemmed":"Crisia romanic","simple":"Crisia romanica","full":"Crisia romanica"},"cardinality":2,"authorship":{"verbatim":"Zágoršek Silye \u0026 Szabó 2008","normalized":"Zágoršek Silye \u0026 Szabó 2008","year":"2008","authors":["Zágoršek Silye","Szabó"],"originalAuth":{"authors":["Zágoršek Silye","Szabó"],"year":{"year":"2008"}}},"details":{"species":{"genus":"Crisia","species":"romanica","authorship":{"verbatim":"Zágoršek Silye \u0026 Szabó 2008","normalized":"Zágoršek Silye \u0026 Szabó 2008","year":"2008","authors":["Zágoršek Silye","Szabó"],"originalAuth":{"authors":["Zágoršek Silye","Szabó"],"year":{"year":"2008"}}}}},"words":[{"verbatim":"Crisia","normalized":"Crisia","wordType":"GENUS","start":0,"end":6},{"verbatim":"romanica","normalized":"romanica","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Zágoršek","normalized":"Zágoršek","wordType":"AUTHOR_WORD","start":16,"end":24},{"verbatim":"Silye","normalized":"Silye","wordType":"AUTHOR_WORD","start":25,"end":30},{"verbatim":"Szabó","normalized":"Szabó","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"2008","normalized":"2008","wordType":"YEAR","start":39,"end":43}],"id":"f95faed8-6aa4-53d2-bd9d-e9463d5eca60","parserVersion":"test_version"}
```

Name: Æschopalæa grisella Pascoe, 1864

Canonical: Aeschopalaea grisella

Authorship: Pascoe 1864

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Æschopalæa grisella Pascoe, 1864","normalized":"Aeschopalaea grisella Pascoe 1864","canonical":{"stemmed":"Aeschopalaea grisell","simple":"Aeschopalaea grisella","full":"Aeschopalaea grisella"},"cardinality":2,"authorship":{"verbatim":"Pascoe, 1864","normalized":"Pascoe 1864","year":"1864","authors":["Pascoe"],"originalAuth":{"authors":["Pascoe"],"year":{"year":"1864"}}},"details":{"species":{"genus":"Aeschopalaea","species":"grisella","authorship":{"verbatim":"Pascoe, 1864","normalized":"Pascoe 1864","year":"1864","authors":["Pascoe"],"originalAuth":{"authors":["Pascoe"],"year":{"year":"1864"}}}}},"words":[{"verbatim":"Æschopalæa","normalized":"Aeschopalaea","wordType":"GENUS","start":0,"end":10},{"verbatim":"grisella","normalized":"grisella","wordType":"SPECIES","start":11,"end":19},{"verbatim":"Pascoe","normalized":"Pascoe","wordType":"AUTHOR_WORD","start":20,"end":26},{"verbatim":"1864","normalized":"1864","wordType":"YEAR","start":28,"end":32}],"id":"82afddf5-4bac-5858-a6a3-93b270b844e8","parserVersion":"test_version"}
```

Name: Læptura laetifica Dow, 1913

Canonical: Laeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Læptura laetifica Dow, 1913","normalized":"Laeptura laetifica Dow 1913","canonical":{"stemmed":"Laeptura laetific","simple":"Laeptura laetifica","full":"Laeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Laeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Læptura","normalized":"Laeptura","wordType":"GENUS","start":0,"end":7},{"verbatim":"laetifica","normalized":"laetifica","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dow","normalized":"Dow","wordType":"AUTHOR_WORD","start":18,"end":21},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":23,"end":27}],"id":"dc1da297-0a85-583d-9a72-d888ddb37ae7","parserVersion":"test_version"}
```

Name: Leptura lætifica Dow, 1913

Canonical: Leptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leptura lætifica Dow, 1913","normalized":"Leptura laetifica Dow 1913","canonical":{"stemmed":"Leptura laetific","simple":"Leptura laetifica","full":"Leptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Leptura","normalized":"Leptura","wordType":"GENUS","start":0,"end":7},{"verbatim":"lætifica","normalized":"laetifica","wordType":"SPECIES","start":8,"end":16},{"verbatim":"Dow","normalized":"Dow","wordType":"AUTHOR_WORD","start":17,"end":20},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":22,"end":26}],"id":"0067abce-1fa8-5911-8176-011065a113a6","parserVersion":"test_version"}
```

Name: Leptura leætifica Dow, 1913

Canonical: Leptura leaetifica

Authorship: Dow 1913

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leptura leætifica Dow, 1913","normalized":"Leptura leaetifica Dow 1913","canonical":{"stemmed":"Leptura leaetific","simple":"Leptura leaetifica","full":"Leptura leaetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leptura","species":"leaetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Leptura","normalized":"Leptura","wordType":"GENUS","start":0,"end":7},{"verbatim":"leætifica","normalized":"leaetifica","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dow","normalized":"Dow","wordType":"AUTHOR_WORD","start":18,"end":21},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":23,"end":27}],"id":"06e6f378-8a12-500a-bab1-27e8b9c6b0cb","parserVersion":"test_version"}
```

Name: Leæptura laetifica Dow, 1913

Canonical: Leaeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leæptura laetifica Dow, 1913","normalized":"Leaeptura laetifica Dow 1913","canonical":{"stemmed":"Leaeptura laetific","simple":"Leaeptura laetifica","full":"Leaeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leaeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Leæptura","normalized":"Leaeptura","wordType":"GENUS","start":0,"end":8},{"verbatim":"laetifica","normalized":"laetifica","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Dow","normalized":"Dow","wordType":"AUTHOR_WORD","start":19,"end":22},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":24,"end":28}],"id":"18311671-6006-5382-b3b9-d9e959fa61c1","parserVersion":"test_version"}
```

Name: Leœptura laetifica Dow, 1913

Canonical: Leoeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leœptura laetifica Dow, 1913","normalized":"Leoeptura laetifica Dow 1913","canonical":{"stemmed":"Leoeptura laetific","simple":"Leoeptura laetifica","full":"Leoeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leoeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Leœptura","normalized":"Leoeptura","wordType":"GENUS","start":0,"end":8},{"verbatim":"laetifica","normalized":"laetifica","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Dow","normalized":"Dow","wordType":"AUTHOR_WORD","start":19,"end":22},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":24,"end":28}],"id":"c31a86ea-3f68-52b4-a746-5ca921816357","parserVersion":"test_version"}
```

Name: Ærenea cognata Lacordaire, 1872

Canonical: Aerenea cognata

Authorship: Lacordaire 1872

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Ærenea cognata Lacordaire, 1872","normalized":"Aerenea cognata Lacordaire 1872","canonical":{"stemmed":"Aerenea cognat","simple":"Aerenea cognata","full":"Aerenea cognata"},"cardinality":2,"authorship":{"verbatim":"Lacordaire, 1872","normalized":"Lacordaire 1872","year":"1872","authors":["Lacordaire"],"originalAuth":{"authors":["Lacordaire"],"year":{"year":"1872"}}},"details":{"species":{"genus":"Aerenea","species":"cognata","authorship":{"verbatim":"Lacordaire, 1872","normalized":"Lacordaire 1872","year":"1872","authors":["Lacordaire"],"originalAuth":{"authors":["Lacordaire"],"year":{"year":"1872"}}}}},"words":[{"verbatim":"Ærenea","normalized":"Aerenea","wordType":"GENUS","start":0,"end":6},{"verbatim":"cognata","normalized":"cognata","wordType":"SPECIES","start":7,"end":14},{"verbatim":"Lacordaire","normalized":"Lacordaire","wordType":"AUTHOR_WORD","start":15,"end":25},{"verbatim":"1872","normalized":"1872","wordType":"YEAR","start":27,"end":31}],"id":"e7f394a9-59f3-5a9c-b375-d6949e232694","parserVersion":"test_version"}
```

Name: Œdicnemus capensis

Canonical: Oedicnemus capensis

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Œdicnemus capensis","normalized":"Oedicnemus capensis","canonical":{"stemmed":"Oedicnemus capens","simple":"Oedicnemus capensis","full":"Oedicnemus capensis"},"cardinality":2,"details":{"species":{"genus":"Oedicnemus","species":"capensis"}},"words":[{"verbatim":"Œdicnemus","normalized":"Oedicnemus","wordType":"GENUS","start":0,"end":9},{"verbatim":"capensis","normalized":"capensis","wordType":"SPECIES","start":10,"end":18}],"id":"33dcf668-48f3-5504-87c1-fe6646a51189","parserVersion":"test_version"}
```

Name: Œnanthe œnanthe

Canonical: Oenanthe oenanthe

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Œnanthe œnanthe","normalized":"Oenanthe oenanthe","canonical":{"stemmed":"Oenanthe oenanth","simple":"Oenanthe oenanthe","full":"Oenanthe oenanthe"},"cardinality":2,"details":{"species":{"genus":"Oenanthe","species":"oenanthe"}},"words":[{"verbatim":"Œnanthe","normalized":"Oenanthe","wordType":"GENUS","start":0,"end":7},{"verbatim":"œnanthe","normalized":"oenanthe","wordType":"SPECIES","start":8,"end":15}],"id":"3e4ce8df-36d0-5529-9725-8336fa694c9a","parserVersion":"test_version"}
```

Name: Hördeum vulgare cœrulescens

Canonical: Hoerdeum vulgare coerulescens

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hördeum vulgare cœrulescens","normalized":"Hoerdeum vulgare coerulescens","canonical":{"stemmed":"Hoerdeum uulgar coerulescens","simple":"Hoerdeum vulgare coerulescens","full":"Hoerdeum vulgare coerulescens"},"cardinality":3,"details":{"infraspecies":{"genus":"Hoerdeum","species":"vulgare","infraspecies":[{"value":"coerulescens"}]}},"words":[{"verbatim":"Hördeum","normalized":"Hoerdeum","wordType":"GENUS","start":0,"end":7},{"verbatim":"vulgare","normalized":"vulgare","wordType":"SPECIES","start":8,"end":15},{"verbatim":"cœrulescens","normalized":"coerulescens","wordType":"INFRASPECIES","start":16,"end":27}],"id":"44916bbf-7112-5604-b691-e425447974d4","parserVersion":"test_version"}
```

Name: Hordeum vulgare cœrulescens Metzger

Canonical: Hordeum vulgare coerulescens

Authorship: Metzger

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hordeum vulgare cœrulescens Metzger","normalized":"Hordeum vulgare coerulescens Metzger","canonical":{"stemmed":"Hordeum uulgar coerulescens","simple":"Hordeum vulgare coerulescens","full":"Hordeum vulgare coerulescens"},"cardinality":3,"authorship":{"verbatim":"Metzger","normalized":"Metzger","authors":["Metzger"],"originalAuth":{"authors":["Metzger"]}},"details":{"infraspecies":{"genus":"Hordeum","species":"vulgare","infraspecies":[{"value":"coerulescens","authorship":{"verbatim":"Metzger","normalized":"Metzger","authors":["Metzger"],"originalAuth":{"authors":["Metzger"]}}}]}},"words":[{"verbatim":"Hordeum","normalized":"Hordeum","wordType":"GENUS","start":0,"end":7},{"verbatim":"vulgare","normalized":"vulgare","wordType":"SPECIES","start":8,"end":15},{"verbatim":"cœrulescens","normalized":"coerulescens","wordType":"INFRASPECIES","start":16,"end":27},{"verbatim":"Metzger","normalized":"Metzger","wordType":"AUTHOR_WORD","start":28,"end":35}],"id":"3029cf1f-da59-5955-86af-40c0b57bd59d","parserVersion":"test_version"}
```

Name: Hordeum vulgare f. cœrulescens

Canonical: Hordeum vulgare f. coerulescens

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hordeum vulgare f. cœrulescens","normalized":"Hordeum vulgare f. coerulescens","canonical":{"stemmed":"Hordeum uulgar coerulescens","simple":"Hordeum vulgare coerulescens","full":"Hordeum vulgare f. coerulescens"},"cardinality":3,"details":{"infraspecies":{"genus":"Hordeum","species":"vulgare","infraspecies":[{"value":"coerulescens","rank":"f."}]}},"words":[{"verbatim":"Hordeum","normalized":"Hordeum","wordType":"GENUS","start":0,"end":7},{"verbatim":"vulgare","normalized":"vulgare","wordType":"SPECIES","start":8,"end":15},{"verbatim":"f.","normalized":"f.","wordType":"RANK","start":16,"end":18},{"verbatim":"cœrulescens","normalized":"coerulescens","wordType":"INFRASPECIES","start":19,"end":30}],"id":"27dd2ab3-8bf9-5f72-90bb-5c94530822f2","parserVersion":"test_version"}
```

### Old style s (ſ)

Name: Musca domeſtica Linnaeus 1758

Canonical: Musca domestica

Authorship: Linnaeus 1758

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Musca domeſtica Linnaeus 1758","normalized":"Musca domestica Linnaeus 1758","canonical":{"stemmed":"Musca domestic","simple":"Musca domestica","full":"Musca domestica"},"cardinality":2,"authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Musca","species":"domestica","authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Musca","normalized":"Musca","wordType":"GENUS","start":0,"end":5},{"verbatim":"domeſtica","normalized":"domestica","wordType":"SPECIES","start":6,"end":15},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":16,"end":24},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":25,"end":29}],"id":"a9f11057-210a-51d0-8402-79d4075607d3","parserVersion":"test_version"}
```

Name: Amphisbæna fuliginoſa Linnaeus 1758

Canonical: Amphisbaena fuliginosa

Authorship: Linnaeus 1758

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Amphisbæna fuliginoſa Linnaeus 1758","normalized":"Amphisbaena fuliginosa Linnaeus 1758","canonical":{"stemmed":"Amphisbaena fuliginos","simple":"Amphisbaena fuliginosa","full":"Amphisbaena fuliginosa"},"cardinality":2,"authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Amphisbaena","species":"fuliginosa","authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Amphisbæna","normalized":"Amphisbaena","wordType":"GENUS","start":0,"end":10},{"verbatim":"fuliginoſa","normalized":"fuliginosa","wordType":"SPECIES","start":11,"end":21},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":31,"end":35}],"id":"d2f6423b-7a8f-5389-a286-c074fb634c5a","parserVersion":"test_version"}
```

Name: Dreyfusia nüßlini

Canonical: Dreyfusia nuesslini

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Dreyfusia nüßlini","normalized":"Dreyfusia nuesslini","canonical":{"stemmed":"Dreyfusia nuesslin","simple":"Dreyfusia nuesslini","full":"Dreyfusia nuesslini"},"cardinality":2,"details":{"species":{"genus":"Dreyfusia","species":"nuesslini"}},"words":[{"verbatim":"Dreyfusia","normalized":"Dreyfusia","wordType":"GENUS","start":0,"end":9},{"verbatim":"nüßlini","normalized":"nuesslini","wordType":"SPECIES","start":10,"end":17}],"id":"27679e50-c41b-5a3d-b619-d378d503be8c","parserVersion":"test_version"}
```

### Miscellaneous diacritics

Name: Pärdosa

Canonical: Paerdosa

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pärdosa","normalized":"Paerdosa","canonical":{"stemmed":"Paerdosa","simple":"Paerdosa","full":"Paerdosa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Paerdosa"}},"words":[{"verbatim":"Pärdosa","normalized":"Paerdosa","wordType":"UNINOMIAL","start":0,"end":7}],"id":"3f493cea-a62c-5bfc-a9a8-e3305e6936db","parserVersion":"test_version"}
```

Name: Pårdosa

Canonical: Paordosa

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pårdosa","normalized":"Paordosa","canonical":{"stemmed":"Paordosa","simple":"Paordosa","full":"Paordosa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Paordosa"}},"words":[{"verbatim":"Pårdosa","normalized":"Paordosa","wordType":"UNINOMIAL","start":0,"end":7}],"id":"eead0d2e-5f37-503c-add2-e344c341be20","parserVersion":"test_version"}
```

Name: Pardøsa

Canonical: Pardoesa

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pardøsa","normalized":"Pardoesa","canonical":{"stemmed":"Pardoesa","simple":"Pardoesa","full":"Pardoesa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pardoesa"}},"words":[{"verbatim":"Pardøsa","normalized":"Pardoesa","wordType":"UNINOMIAL","start":0,"end":7}],"id":"6922fdef-226d-59fc-9cc6-7b446d7ce37b","parserVersion":"test_version"}
```

Name: Pardösa

Canonical: Pardoesa

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pardösa","normalized":"Pardoesa","canonical":{"stemmed":"Pardoesa","simple":"Pardoesa","full":"Pardoesa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pardoesa"}},"words":[{"verbatim":"Pardösa","normalized":"Pardoesa","wordType":"UNINOMIAL","start":0,"end":7}],"id":"7873dfb8-fc08-50e8-bd23-e94deb9317bc","parserVersion":"test_version"}
```

Name: Rühlella

Canonical: Ruehlella

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Rühlella","normalized":"Ruehlella","canonical":{"stemmed":"Ruehlella","simple":"Ruehlella","full":"Ruehlella"},"cardinality":1,"details":{"uninomial":{"uninomial":"Ruehlella"}},"words":[{"verbatim":"Rühlella","normalized":"Ruehlella","wordType":"UNINOMIAL","start":0,"end":8}],"id":"228b2714-3726-5ae8-b802-59bdbc8d20a6","parserVersion":"test_version"}
```

### Open Nomenclature ('approximate' names)

<!-- Open nomenclature -- cf., aff., sp., etc. -->
Name: Solygia ? distanti

Canonical: Solygia

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Solygia ? distanti","normalized":"Solygia","canonical":{"stemmed":"Solygia","simple":"Solygia","full":"Solygia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Solygia","approximationMarker":"?","ignored":" distanti"}},"words":[{"verbatim":"Solygia","normalized":"Solygia","wordType":"GENUS","start":0,"end":7},{"verbatim":"?","normalized":"?","wordType":"APPROXIMATION_MARKER","start":8,"end":9}],"id":"b9e3508f-1c0e-554c-8642-dd1cfd02631c","parserVersion":"test_version"}
```

<!-- Ambiguity -- can be an unknown author or approx name-->
Name: Buteo borealis ? ventralis

Canonical: Buteo borealis ventralis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Author as a question mark"},{"quality":3,"warning":"Author is too short"},{"quality":2,"warning":"Author is unknown"}],"verbatim":"Buteo borealis ? ventralis","normalized":"Buteo borealis anon. ventralis","canonical":{"stemmed":"Buteo boreal uentral","simple":"Buteo borealis ventralis","full":"Buteo borealis ventralis"},"cardinality":3,"details":{"infraspecies":{"genus":"Buteo","species":"borealis","authorship":{"verbatim":"?","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"infraspecies":[{"value":"ventralis"}]}},"words":[{"verbatim":"Buteo","normalized":"Buteo","wordType":"GENUS","start":0,"end":5},{"verbatim":"borealis","normalized":"borealis","wordType":"SPECIES","start":6,"end":14},{"verbatim":"?","normalized":"anon.","wordType":"AUTHOR_WORD","start":15,"end":16},{"verbatim":"ventralis","normalized":"ventralis","wordType":"INFRASPECIES","start":17,"end":26}],"id":"d26a4791-4858-5239-8a57-c88957d40919","parserVersion":"test_version"}
```

Name: Euxoa nr. idahoensis sp. 1clay

Canonical: Euxoa

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Euxoa nr. idahoensis sp. 1clay","normalized":"Euxoa","canonical":{"stemmed":"Euxoa","simple":"Euxoa","full":"Euxoa"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Euxoa","approximationMarker":"nr.","ignored":" idahoensis sp. 1clay"}},"words":[{"verbatim":"Euxoa","normalized":"Euxoa","wordType":"GENUS","start":0,"end":5},{"verbatim":"nr.","normalized":"nr.","wordType":"APPROXIMATION_MARKER","start":6,"end":9}],"id":"02a664be-422a-56cb-b431-99aecf793721","parserVersion":"test_version"}
```

Name: Acarinina aff. pentacamerata

Canonical: Acarinina

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acarinina aff. pentacamerata","normalized":"Acarinina","canonical":{"stemmed":"Acarinina","simple":"Acarinina","full":"Acarinina"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acarinina","approximationMarker":"aff.","ignored":" pentacamerata"}},"words":[{"verbatim":"Acarinina","normalized":"Acarinina","wordType":"GENUS","start":0,"end":9},{"verbatim":"aff.","normalized":"aff.","wordType":"APPROXIMATION_MARKER","start":10,"end":14}],"id":"c4ab66ee-79a2-5100-8b87-20e60cf2a358","parserVersion":"test_version"}
```

Name: Acarinina aff pentacamerata

Canonical: Acarinina

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acarinina aff pentacamerata","normalized":"Acarinina","canonical":{"stemmed":"Acarinina","simple":"Acarinina","full":"Acarinina"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acarinina","approximationMarker":"aff","ignored":" pentacamerata"}},"words":[{"verbatim":"Acarinina","normalized":"Acarinina","wordType":"GENUS","start":0,"end":9},{"verbatim":"aff","normalized":"aff","wordType":"APPROXIMATION_MARKER","start":10,"end":13}],"id":"06a32183-0aa7-5a00-9753-46db1141daa4","parserVersion":"test_version"}
```

Name: Sphingomonas sp. 37

Canonical: Sphingomonas

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Sphingomonas sp. 37","normalized":"Sphingomonas","canonical":{"stemmed":"Sphingomonas","simple":"Sphingomonas","full":"Sphingomonas"},"cardinality":0,"bacteria":"yes","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Sphingomonas","approximationMarker":"sp.","ignored":" 37"}},"words":[{"verbatim":"Sphingomonas","normalized":"Sphingomonas","wordType":"GENUS","start":0,"end":12},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":13,"end":16}],"id":"1daffd3a-f4de-58d9-91e3-ae4d08a50ce0","parserVersion":"test_version"}
```

Name: Thryothorus leucotis spp. bogotensis

Canonical: Thryothorus leucotis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Thryothorus leucotis spp. bogotensis","normalized":"Thryothorus leucotis","canonical":{"stemmed":"Thryothorus leucot","simple":"Thryothorus leucotis","full":"Thryothorus leucotis"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Thryothorus","species":"leucotis","approximationMarker":"spp.","ignored":" bogotensis"}},"words":[{"verbatim":"Thryothorus","normalized":"Thryothorus","wordType":"GENUS","start":0,"end":11},{"verbatim":"leucotis","normalized":"leucotis","wordType":"SPECIES","start":12,"end":20},{"verbatim":"spp.","normalized":"spp.","wordType":"APPROXIMATION_MARKER","start":21,"end":25}],"id":"d2cb7212-ff62-5e31-9ab9-31214a9782d5","parserVersion":"test_version"}
```

Name: Endoxyla sp. GM-, 2003

Canonical: Endoxyla

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Endoxyla sp. GM-, 2003","normalized":"Endoxyla","canonical":{"stemmed":"Endoxyla","simple":"Endoxyla","full":"Endoxyla"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Endoxyla","approximationMarker":"sp.","ignored":" GM-, 2003"}},"words":[{"verbatim":"Endoxyla","normalized":"Endoxyla","wordType":"GENUS","start":0,"end":8},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":9,"end":12}],"id":"8a80bfee-947d-5602-9958-a2338ff46a4d","parserVersion":"test_version"}
```

Name: X Aegilotrichum sp.

Canonical: × Aegilotrichum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Aegilotrichum sp.","normalized":"× Aegilotrichum","canonical":{"stemmed":"Aegilotrichum","simple":"Aegilotrichum","full":"× Aegilotrichum"},"cardinality":0,"hybrid":"NAMED_HYBRID","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Aegilotrichum","approximationMarker":"sp."}},"words":[{"verbatim":"X","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Aegilotrichum","normalized":"Aegilotrichum","wordType":"GENUS","start":2,"end":15},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":16,"end":19}],"id":"308357ff-7f86-53b9-955b-88a52ef7623a","parserVersion":"test_version"}
```

Name: Liopropoma sp.2 Not applicable

Canonical: Liopropoma

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"},{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Liopropoma sp.2 Not applicable","normalized":"Liopropoma","canonical":{"stemmed":"Liopropoma","simple":"Liopropoma","full":"Liopropoma"},"cardinality":0,"surrogate":"APPROXIMATION","tail":" Not applicable","details":{"approximation":{"genus":"Liopropoma","approximationMarker":"sp.","ignored":"2"}},"words":[{"verbatim":"Liopropoma","normalized":"Liopropoma","wordType":"GENUS","start":0,"end":10},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":11,"end":14}],"id":"fb3779a4-57a0-5628-8c4e-e341ca4f952d","parserVersion":"test_version"}
```

Name: Lacanobia sp. nr. subjuncta Bold:Aab, 0925

Canonical: Lacanobia

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Lacanobia sp. nr. subjuncta Bold:Aab, 0925","normalized":"Lacanobia","canonical":{"stemmed":"Lacanobia","simple":"Lacanobia","full":"Lacanobia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Lacanobia","approximationMarker":"sp. nr.","ignored":" subjuncta Bold:Aab, 0925"}},"words":[{"verbatim":"Lacanobia","normalized":"Lacanobia","wordType":"GENUS","start":0,"end":9},{"verbatim":"sp. nr.","normalized":"sp. nr.","wordType":"APPROXIMATION_MARKER","start":10,"end":17}],"id":"05b25429-cb9e-54a1-8e1a-bac9a26d5f46","parserVersion":"test_version"}
```

Name: Lacanobia nr. subjuncta Bold:Aab, 0925

Canonical: Lacanobia

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Lacanobia nr. subjuncta Bold:Aab, 0925","normalized":"Lacanobia","canonical":{"stemmed":"Lacanobia","simple":"Lacanobia","full":"Lacanobia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Lacanobia","approximationMarker":"nr.","ignored":" subjuncta Bold:Aab, 0925"}},"words":[{"verbatim":"Lacanobia","normalized":"Lacanobia","wordType":"GENUS","start":0,"end":9},{"verbatim":"nr.","normalized":"nr.","wordType":"APPROXIMATION_MARKER","start":10,"end":13}],"id":"31763a26-a69b-5af8-8703-5da372bdf895","parserVersion":"test_version"}
```

<!--
 TODO wrong name result
Placodium chrysoleucum cf. chrysoleucum (Sm.) anon.
{"quality":2,"parsed":true,"verbatim":"Placodium chrysoleucum cf. chrysoleucum (Sm.) anon.","surrogate":true,"qualityWarnings":[[2,"Author is unknown"]],"normalized":"Placodium cf. chrysoleucum chrysoleucum (Sm.) anon.","canonicalName":{"value":"Placodium chrysoleucum chrysoleucum","valueRanked":"Placodium chrysoleucum chrysoleucum"},"virus":false,"positions":[["genus",0,9],["specificEpithet",10,22],["annotationIdentification",23,26],["infraspecificEpithet",27,39],["authorWord",41,44],["authorWord",46,51]],"nameStringId":"e0b84689-70e3-508a-8040-8859dc6084c0","parserVersion":"test_version","hybrid":false,"details":[{"genus":{"value":"Placodium"},"specificEpithet":{"value":"chrysoleucum"},"infraspecificEpithets":[{"value":"chrysoleucum","authorship":{"value":"(Sm.) anon.","basionymAuthorship":{"authors":["Sm."]},"combinationAuthorship":{"authors":["anon."]}}}],"annotationIdentification":"cf."}],"bacteria":false}
-->

Name: Abturia cf. alabamensis (Morton )

Canonical: Abturia alabamensis

Authorship: (Morton)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Abturia cf. alabamensis (Morton )","normalized":"Abturia cf. alabamensis (Morton)","canonical":{"stemmed":"Abturia alabamens","simple":"Abturia alabamensis","full":"Abturia alabamensis"},"cardinality":2,"authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"surrogate":"COMPARISON","details":{"comparison":{"genus":"Abturia","species":"alabamensis","authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"comparisonMarker":"cf."}},"words":[{"verbatim":"Abturia","normalized":"Abturia","wordType":"GENUS","start":0,"end":7},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":8,"end":11},{"verbatim":"alabamensis","normalized":"alabamensis","wordType":"SPECIES","start":12,"end":23},{"verbatim":"Morton","normalized":"Morton","wordType":"AUTHOR_WORD","start":25,"end":31}],"id":"5fd4ce59-98d3-50af-9e28-918adc47d264","parserVersion":"test_version"}
```

Name: Abturia cf alabamensis (Morton )

Canonical: Abturia alabamensis

Authorship: (Morton)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Abturia cf alabamensis (Morton )","normalized":"Abturia cf. alabamensis (Morton)","canonical":{"stemmed":"Abturia alabamens","simple":"Abturia alabamensis","full":"Abturia alabamensis"},"cardinality":2,"authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"surrogate":"COMPARISON","details":{"comparison":{"genus":"Abturia","species":"alabamensis","authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"comparisonMarker":"cf."}},"words":[{"verbatim":"Abturia","normalized":"Abturia","wordType":"GENUS","start":0,"end":7},{"verbatim":"cf","normalized":"cf.","wordType":"COMPARISON_MARKER","start":8,"end":10},{"verbatim":"alabamensis","normalized":"alabamensis","wordType":"SPECIES","start":11,"end":22},{"verbatim":"Morton","normalized":"Morton","wordType":"AUTHOR_WORD","start":24,"end":30}],"id":"423cd26d-c6fd-54fb-937b-f98ba8056fc0","parserVersion":"test_version"}
```

<!--TODO Larus occidentalis cf. wymani|{}-->

Name: Calidris cf. cooperi

Canonical: Calidris cooperi

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Calidris cf. cooperi","normalized":"Calidris cf. cooperi","canonical":{"stemmed":"Calidris cooper","simple":"Calidris cooperi","full":"Calidris cooperi"},"cardinality":2,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Calidris","species":"cooperi","comparisonMarker":"cf."}},"words":[{"verbatim":"Calidris","normalized":"Calidris","wordType":"GENUS","start":0,"end":8},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":9,"end":12},{"verbatim":"cooperi","normalized":"cooperi","wordType":"SPECIES","start":13,"end":20}],"id":"bb19b56e-462f-5daf-a1aa-d4ead082f321","parserVersion":"test_version"}
```

<!--TODO merge comparison with species, infraspecies nodes instead of its own node-->
Name: Aesculus cf. × hybrida

Canonical: Aesculus × hybrida

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aesculus cf. × hybrida","normalized":"Aesculus × hybrida","canonical":{"stemmed":"Aesculus hybrid","simple":"Aesculus hybrida","full":"Aesculus × hybrida"},"cardinality":2,"hybrid":"NAMED_HYBRID","surrogate":"COMPARISON","details":{"species":{"genus":"Aesculus","species":"hybrida"}},"words":[{"verbatim":"Aesculus","normalized":"Aesculus","wordType":"GENUS","start":0,"end":8},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":9,"end":12},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":13,"end":14},{"verbatim":"hybrida","normalized":"hybrida","wordType":"SPECIES","start":15,"end":22}],"id":"6e255814-1c53-54f0-8536-fee957312e9a","parserVersion":"test_version"}
```

<!-- TODO missing subgenus info -->
Name: Daphnia (Daphnia) x krausi Flossner 1993

Canonical: Daphnia × krausi

Authorship: Flossner 1993

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Daphnia (Daphnia) x krausi Flossner 1993","normalized":"Daphnia × krausi Flossner 1993","canonical":{"stemmed":"Daphnia kraus","simple":"Daphnia krausi","full":"Daphnia × krausi"},"cardinality":2,"authorship":{"verbatim":"Flossner 1993","normalized":"Flossner 1993","year":"1993","authors":["Flossner"],"originalAuth":{"authors":["Flossner"],"year":{"year":"1993"}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Daphnia","species":"krausi Flossner 1993","authorship":{"verbatim":"Flossner 1993","normalized":"Flossner 1993","year":"1993","authors":["Flossner"],"originalAuth":{"authors":["Flossner"],"year":{"year":"1993"}}}}},"words":[{"verbatim":"Daphnia","normalized":"Daphnia","wordType":"GENUS","start":0,"end":7},{"verbatim":"x","normalized":"×","wordType":"HYBRID_CHAR","start":18,"end":19},{"verbatim":"krausi","normalized":"krausi","wordType":"SPECIES","start":20,"end":26},{"verbatim":"Flossner","normalized":"Flossner","wordType":"AUTHOR_WORD","start":27,"end":35},{"verbatim":"1993","normalized":"1993","wordType":"YEAR","start":36,"end":40}],"id":"b509d1f1-ce1d-56a1-a15e-2aa9430dce0e","parserVersion":"test_version"}
```

<!--TODO incorrect interpretation-->
Name: Barbus cf macrotaenia × toppini

Canonical: Barbus macrotaenia × Barbus toppini

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Incomplete hybrid formula"},{"quality":4,"warning":"Name comparison"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Barbus cf macrotaenia × toppini","normalized":"Barbus cf. macrotaenia × Barbus toppini","canonical":{"stemmed":"Barbus macrotaen × Barbus toppin","simple":"Barbus macrotaenia × Barbus toppini","full":"Barbus macrotaenia × Barbus toppini"},"cardinality":0,"hybrid":"HYBRID_FORMULA","surrogate":"COMPARISON","details":{"hybridFormula":[{"comparison":{"genus":"Barbus","species":"macrotaenia","comparisonMarker":"cf."}},{"species":{"genus":"Barbus","species":"toppini"}}]},"words":[{"verbatim":"Barbus","normalized":"Barbus","wordType":"GENUS","start":0,"end":6},{"verbatim":"cf","normalized":"cf.","wordType":"COMPARISON_MARKER","start":7,"end":9},{"verbatim":"macrotaenia","normalized":"macrotaenia","wordType":"SPECIES","start":10,"end":21},{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":22,"end":23},{"verbatim":"toppini","normalized":"toppini","wordType":"SPECIES","start":24,"end":31}],"id":"37b0b404-d5d9-5699-bbb2-8c3d9bf543a3","parserVersion":"test_version"}
```

Name: Gemmula cf. cosmoi NP-2008

Canonical: Gemmula cosmoi

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"},{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Gemmula cf. cosmoi NP-2008","normalized":"Gemmula cf. cosmoi","canonical":{"stemmed":"Gemmula cosmo","simple":"Gemmula cosmoi","full":"Gemmula cosmoi"},"cardinality":2,"surrogate":"COMPARISON","tail":" NP-2008","details":{"comparison":{"genus":"Gemmula","species":"cosmoi","comparisonMarker":"cf."}},"words":[{"verbatim":"Gemmula","normalized":"Gemmula","wordType":"GENUS","start":0,"end":7},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":8,"end":11},{"verbatim":"cosmoi","normalized":"cosmoi","wordType":"SPECIES","start":12,"end":18}],"id":"87a593b3-2383-5f1b-8772-85e0a4a31b79","parserVersion":"test_version"}
```

### Surrogate Name-Strings

Name: Coleoptera sp. BOLD:AAV0432

Canonical: Coleoptera

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Coleoptera sp. BOLD:AAV0432","normalized":"Coleoptera","canonical":{"stemmed":"Coleoptera","simple":"Coleoptera","full":"Coleoptera"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Coleoptera","approximationMarker":"sp.","ignored":" BOLD:AAV0432"}},"words":[{"verbatim":"Coleoptera","normalized":"Coleoptera","wordType":"GENUS","start":0,"end":10},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":11,"end":14}],"id":"65b09adc-12a0-5fbb-a885-75200eacb98a","parserVersion":"test_version"}
```

Name: Coleoptera Bold:AAV0432

Canonical: Coleoptera

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Coleoptera Bold:AAV0432","normalized":"Coleoptera","canonical":{"stemmed":"Coleoptera","simple":"Coleoptera","full":"Coleoptera"},"cardinality":0,"surrogate":"BOLD_SURROGATE","tail":" Bold:AAV0432","details":{"uninomial":{"uninomial":"Coleoptera"}},"words":[{"verbatim":"Coleoptera","normalized":"Coleoptera","wordType":"UNINOMIAL","start":0,"end":10}],"id":"9b3865ee-dcf6-5861-9910-58d9f3eafbb1","parserVersion":"test_version"}
```

### Virus-like "normal" names

Name: Ceylonesmus vector Chamberlin, 1941

Canonical: Ceylonesmus vector

Authorship: Chamberlin 1941

```json
{"parsed":true,"quality":1,"verbatim":"Ceylonesmus vector Chamberlin, 1941","normalized":"Ceylonesmus vector Chamberlin 1941","canonical":{"stemmed":"Ceylonesmus uector","simple":"Ceylonesmus vector","full":"Ceylonesmus vector"},"cardinality":2,"authorship":{"verbatim":"Chamberlin, 1941","normalized":"Chamberlin 1941","year":"1941","authors":["Chamberlin"],"originalAuth":{"authors":["Chamberlin"],"year":{"year":"1941"}}},"details":{"species":{"genus":"Ceylonesmus","species":"vector","authorship":{"verbatim":"Chamberlin, 1941","normalized":"Chamberlin 1941","year":"1941","authors":["Chamberlin"],"originalAuth":{"authors":["Chamberlin"],"year":{"year":"1941"}}}}},"words":[{"verbatim":"Ceylonesmus","normalized":"Ceylonesmus","wordType":"GENUS","start":0,"end":11},{"verbatim":"vector","normalized":"vector","wordType":"SPECIES","start":12,"end":18},{"verbatim":"Chamberlin","normalized":"Chamberlin","wordType":"AUTHOR_WORD","start":19,"end":29},{"verbatim":"1941","normalized":"1941","wordType":"YEAR","start":31,"end":35}],"id":"00b874b9-c9ac-5b8a-9821-0a641ca26ca0","parserVersion":"test_version"}
```

### Viruses, plasmids, prions etc.

Name: Arv1virus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Arv1virus","cardinality":0,"virus":true,"id":"25c7c012-6600-5073-8e8f-81fbcf841a66","parserVersion":"test_version"}
```

Name: Turtle herpesviruses

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Turtle herpesviruses","cardinality":0,"virus":true,"id":"44dc4404-0bb8-5eaa-b401-1609d98d3b30","parserVersion":"test_version"}
```

Name: Cre expression vector

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Cre expression vector","cardinality":0,"virus":true,"id":"9a282683-c49b-52dc-817f-0281d5b4b831","parserVersion":"test_version"}
```

Name: Cyanophage

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Cyanophage","cardinality":0,"virus":true,"id":"050da5da-716e-5282-97f0-0ea9e375bbf0","parserVersion":"test_version"}
```

Name: Drosophila sturtevanti rhabdovirus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Drosophila sturtevanti rhabdovirus","cardinality":0,"virus":true,"id":"d3510f21-1d57-50e6-98bd-2252259b7052","parserVersion":"test_version"}
```

Name: Hydra expression vector

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Hydra expression vector","cardinality":0,"virus":true,"id":"b22ca1ca-3186-5bc6-9f1a-57ef8c117f25","parserVersion":"test_version"}
```

Name: Gateway destination plasmid

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Gateway destination plasmid","cardinality":0,"id":"21946de0-1c80-543f-ab96-97b81f8d1516","parserVersion":"test_version"}
```

Name: Abutilon mosaic virus [X15983] [X15984] Abutilon mosaic virus ICTV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Abutilon mosaic virus [X15983] [X15984] Abutilon mosaic virus ICTV","cardinality":0,"virus":true,"id":"879da2ea-836c-5ad2-b837-81594a1a208d","parserVersion":"test_version"}
```

Name: Omphalotus sp. Ictv Garcia, 18224

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Omphalotus sp. Ictv Garcia, 18224","cardinality":0,"virus":true,"id":"771a4266-44e3-56d9-9961-9e8a1f1b3936","parserVersion":"test_version"}
```

Name: Acute bee paralysis virus [AF150629] Acute bee paralysis virus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Acute bee paralysis virus [AF150629] Acute bee paralysis virus","cardinality":0,"virus":true,"id":"584822dc-f68f-5abf-aeef-0265172195bf","parserVersion":"test_version"}
```

Name: Adeno-associated virus - 3

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Adeno-associated virus - 3","cardinality":0,"virus":true,"id":"5b16c811-0518-5073-a0be-b59f5faa09fb","parserVersion":"test_version"}
```

Name: ?M1-like Viruses Methanobrevibacter phage PG

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"?M1-like Viruses Methanobrevibacter phage PG","cardinality":0,"virus":true,"id":"b33d05e9-f2a6-5d1b-97e5-3ae061dcd036","parserVersion":"test_version"}
```

Name: Aeromonas phage 65

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Aeromonas phage 65","cardinality":0,"virus":true,"id":"2aef2420-ba68-5887-821f-0ec6eca86660","parserVersion":"test_version"}
```

Name: Bacillus phage SPß [AF020713] Bacillus phage SPb ICTV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Bacillus phage SPß [AF020713] Bacillus phage SPb ICTV","cardinality":0,"virus":true,"id":"ad2b6943-6a54-576d-85e9-e1f8f6aa95db","parserVersion":"test_version"}
```

Name: Apple scar skin viroid

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Apple scar skin viroid","cardinality":0,"virus":true,"id":"7ade78b4-f576-5103-b4a8-4fb9e68845cd","parserVersion":"test_version"}
```

Name: Australian grapevine viroid [X17101] Australian grapevine viroid ICTV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Australian grapevine viroid [X17101] Australian grapevine viroid ICTV","cardinality":0,"virus":true,"id":"381b6868-5d9e-54ec-bae8-84fcc9a3e80c","parserVersion":"test_version"}
```

Name: Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease","cardinality":0,"virus":true,"id":"06193aa6-f2ec-5134-8117-89102448a13e","parserVersion":"test_version"}
```

Name: Phi h-like viruses

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Phi h-like viruses","cardinality":0,"virus":true,"id":"474acd56-6be4-56fc-9045-48a3d570ac97","parserVersion":"test_version"}
```

Name: Viroids

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Viroids","cardinality":0,"virus":true,"id":"641d47bf-c7c4-5218-8e2e-8756ad808653","parserVersion":"test_version"}
```

Name: Fungal prions

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Fungal prions","cardinality":0,"virus":true,"id":"ec273e2d-cdde-5fcb-84dc-a6adf2e309ce","parserVersion":"test_version"}
```

Name: Human rhinovirus A11

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Human rhinovirus A11","cardinality":0,"virus":true,"id":"ba205a7c-1c63-51c7-8f4d-d47665f56c33","parserVersion":"test_version"}
```

Name: Kobuvirus korean black goat/South Korea/2010

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Kobuvirus korean black goat/South Korea/2010","cardinality":0,"virus":true,"id":"4871667d-e362-5f76-a218-6c1bcc090ba9","parserVersion":"test_version"}
```

Name: Australian bat lyssavirus human/AUS/1998

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Australian bat lyssavirus human/AUS/1998","cardinality":0,"virus":true,"id":"5e4fdc2a-3fb3-5776-b94d-04b9f0c6fcbb","parserVersion":"test_version"}
```

Name: Gossypium mustilinum symptomless alphasatellite

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Gossypium mustilinum symptomless alphasatellite","cardinality":0,"virus":true,"id":"d8b1e803-34ba-537b-874b-48521afb92a5","parserVersion":"test_version"}
```

Name: Okra leaf curl Mali alphasatellites-Cameroon

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Okra leaf curl Mali alphasatellites-Cameroon","cardinality":0,"virus":true,"id":"034731b5-3de7-5d48-bf3b-f89272699a45","parserVersion":"test_version"}
```

Name: Bemisia betasatellite LW-2014

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Bemisia betasatellite LW-2014","cardinality":0,"virus":true,"id":"21d06e45-a312-5844-88f7-3eb0b73d1efc","parserVersion":"test_version"}
```

Name: Tomato leaf curl Bangladesh betasatellites [India/Patna/Chilli/2008]

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Tomato leaf curl Bangladesh betasatellites [India/Patna/Chilli/2008]","cardinality":0,"virus":true,"id":"c5def37b-c5d9-57e4-822a-0436629f5d99","parserVersion":"test_version"}
```

Name: Intracisternal A-particles

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Intracisternal A-particles","cardinality":0,"virus":true,"id":"4f16a692-534b-5ec5-87f4-58fe76a0ed9d","parserVersion":"test_version"}
```

Name: Saccharomyces cerevisiae killer particle M1

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Saccharomyces cerevisiae killer particle M1","cardinality":0,"virus":true,"id":"879050a7-5085-5679-85e4-fe47308843dd","parserVersion":"test_version"}
```

Name: Uranotaenia sapphirina NPV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Uranotaenia sapphirina NPV","cardinality":0,"virus":true,"id":"83886b77-a81a-52ba-9b0e-5743b4242b97","parserVersion":"test_version"}
```

Name: Uranotaenia sapphirina Npv

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Uranotaenia sapphirina Npv","cardinality":0,"virus":true,"id":"917cfcbc-3a38-5f59-affc-56c87f04a7ec","parserVersion":"test_version"}
```

Name: Spodoptera exigua nuclear polyhedrosis virus SeMNPV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Spodoptera exigua nuclear polyhedrosis virus SeMNPV","cardinality":0,"virus":true,"id":"a0356512-17eb-51ab-92b3-21d92393b84c","parserVersion":"test_version"}
```

Name: Spodoptera frugiperda MNPV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Spodoptera frugiperda MNPV","cardinality":0,"virus":true,"id":"5a694933-6187-54bb-ae35-77ed3384b69d","parserVersion":"test_version"}
```

Name: Rachiplusia ou MNPV (strain R1)

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Rachiplusia ou MNPV (strain R1)","cardinality":0,"virus":true,"id":"ca77e2a5-fa26-5c7f-bf68-a449c32ea95e","parserVersion":"test_version"}
```

Name: Orgyia pseudotsugata nuclear polyhedrosis virus OpMNPV

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Orgyia pseudotsugata nuclear polyhedrosis virus OpMNPV","cardinality":0,"virus":true,"id":"f3b4269c-a97f-5ff7-bb4a-56d982b3707c","parserVersion":"test_version"}
```

Name: Mamestra configurata NPV-A

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Mamestra configurata NPV-A","cardinality":0,"virus":true,"id":"59160819-f61d-5360-85c5-78b6140a05ca","parserVersion":"test_version"}
```

Name: Helicoverpa armigera SNPV NNg1

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Helicoverpa armigera SNPV NNg1","cardinality":0,"virus":true,"id":"933f0a27-1fd8-5066-90ee-df1ed8148c9c","parserVersion":"test_version"}
```

Name: Zamilon virophage

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Zamilon virophage","cardinality":0,"virus":true,"id":"661132c0-7012-5405-bfc7-31e9a4b3946c","parserVersion":"test_version"}
```

Name: Sputnik virophage 3

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Sputnik virophage 3","cardinality":0,"virus":true,"id":"b206bb35-01bf-59a7-8dad-bc8f99ca0a2a","parserVersion":"test_version"}
```

Name: Bacteriophage PH75

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Bacteriophage PH75","cardinality":0,"virus":true,"id":"605f428e-a4a3-57a2-9dfa-a6a3d99b801d","parserVersion":"test_version"}
```

Name: Escherichia coli bacteriophage

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Escherichia coli bacteriophage","cardinality":0,"virus":true,"id":"c01315c2-e1cc-58c2-b113-2d756985d64b","parserVersion":"test_version"}
```

Name: Betasatellites

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Betasatellites","cardinality":0,"virus":true,"id":"1a6aa729-5fc5-5fbd-9299-efb9a6198310","parserVersion":"test_version"}
```

Name: Satellite Nucleic Acids (Subviral DNA-ssDNA)

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Satellite Nucleic Acids (Subviral DNA-ssDNA)","cardinality":0,"virus":true,"id":"1a769ed9-62cd-54b9-9c94-36d99117b89f","parserVersion":"test_version"}
```

### Name-strings with RNA

Name: ssRNA

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"ssRNA","cardinality":0,"id":"10d5f30c-e51b-54ed-be43-c0ac1656a88a","parserVersion":"test_version"}
```

Name: Alpha proteobacterium RNA12

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Alpha proteobacterium RNA12","cardinality":0,"id":"c2826f30-f6f3-543f-80cf-646adf374a59","parserVersion":"test_version"}
```

Name: Ustilaginoidea virens RNA virus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Ustilaginoidea virens RNA virus","cardinality":0,"virus":true,"id":"61fff10f-7f16-5f42-b642-ba0195abccb8","parserVersion":"test_version"}
```

Name: Candida albicans RNA_CTR0-3

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Candida albicans RNA_CTR0-3","cardinality":0,"id":"0182d44b-5d8b-501d-8f5c-4ef44dff8db4","parserVersion":"test_version"}
```

Name: Carabus satyrus satyrus KURNAKOV, 1962

Canonical: Carabus satyrus satyrus

Authorship: Kurnakov 1962

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"}],"verbatim":"Carabus satyrus satyrus KURNAKOV, 1962","normalized":"Carabus satyrus satyrus Kurnakov 1962","canonical":{"stemmed":"Carabus satyr satyr","simple":"Carabus satyrus satyrus","full":"Carabus satyrus satyrus"},"cardinality":3,"authorship":{"verbatim":"KURNAKOV, 1962","normalized":"Kurnakov 1962","year":"1962","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1962"}}},"details":{"infraspecies":{"genus":"Carabus","species":"satyrus","infraspecies":[{"value":"satyrus","authorship":{"verbatim":"KURNAKOV, 1962","normalized":"Kurnakov 1962","year":"1962","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1962"}}}}]}},"words":[{"verbatim":"Carabus","normalized":"Carabus","wordType":"GENUS","start":0,"end":7},{"verbatim":"satyrus","normalized":"satyrus","wordType":"SPECIES","start":8,"end":15},{"verbatim":"satyrus","normalized":"satyrus","wordType":"INFRASPECIES","start":16,"end":23},{"verbatim":"KURNAKOV","normalized":"Kurnakov","wordType":"AUTHOR_WORD","start":24,"end":32},{"verbatim":"1962","normalized":"1962","wordType":"YEAR","start":34,"end":38}],"id":"81654954-0f47-5715-acb1-1cd8d2c49e9a","parserVersion":"test_version"}
```


### Epithet prioni is not a prion

Name: Fakus prioni

Canonical: Fakus prioni

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Fakus prioni","normalized":"Fakus prioni","canonical":{"stemmed":"Fakus prion","simple":"Fakus prioni","full":"Fakus prioni"},"cardinality":2,"details":{"species":{"genus":"Fakus","species":"prioni"}},"words":[{"verbatim":"Fakus","normalized":"Fakus","wordType":"GENUS","start":0,"end":5},{"verbatim":"prioni","normalized":"prioni","wordType":"SPECIES","start":6,"end":12}],"id":"f2561b5b-37ed-592d-9c12-4ef96d09f554","parserVersion":"test_version"}
```

### Names with "satellite" as a substring

Name: Crassatellites fulvida

Canonical: Crassatellites fulvida

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Crassatellites fulvida","normalized":"Crassatellites fulvida","canonical":{"stemmed":"Crassatellites fuluid","simple":"Crassatellites fulvida","full":"Crassatellites fulvida"},"cardinality":2,"details":{"species":{"genus":"Crassatellites","species":"fulvida"}},"words":[{"verbatim":"Crassatellites","normalized":"Crassatellites","wordType":"GENUS","start":0,"end":14},{"verbatim":"fulvida","normalized":"fulvida","wordType":"SPECIES","start":15,"end":22}],"id":"089171ac-f672-5973-950a-9419651e6b0e","parserVersion":"test_version"}
```

### Bacterial genus

Name: Salmonella werahensis (Castellani) Hauduroy and Ehringer in Hauduroy 1937

Canonical: Salmonella werahensis

Authorship: (Castellani) Hauduroy & Ehringer ex Hauduroy 1937

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Salmonella werahensis (Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"Salmonella werahensis (Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","canonical":{"stemmed":"Salmonella werahens","simple":"Salmonella werahensis","full":"Salmonella werahensis"},"cardinality":2,"authorship":{"verbatim":"(Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"(Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","authors":["Castellani","Hauduroy","Ehringer"],"originalAuth":{"authors":["Castellani"]},"combinationAuth":{"authors":["Hauduroy","Ehringer"],"exAuthors":{"authors":["Hauduroy"],"year":{"year":"1937"}}}},"bacteria":"yes","details":{"species":{"genus":"Salmonella","species":"werahensis","authorship":{"verbatim":"(Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"(Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","authors":["Castellani","Hauduroy","Ehringer"],"originalAuth":{"authors":["Castellani"]},"combinationAuth":{"authors":["Hauduroy","Ehringer"],"exAuthors":{"authors":["Hauduroy"],"year":{"year":"1937"}}}}}},"words":[{"verbatim":"Salmonella","normalized":"Salmonella","wordType":"GENUS","start":0,"end":10},{"verbatim":"werahensis","normalized":"werahensis","wordType":"SPECIES","start":11,"end":21},{"verbatim":"Castellani","normalized":"Castellani","wordType":"AUTHOR_WORD","start":23,"end":33},{"verbatim":"Hauduroy","normalized":"Hauduroy","wordType":"AUTHOR_WORD","start":35,"end":43},{"verbatim":"Ehringer","normalized":"Ehringer","wordType":"AUTHOR_WORD","start":48,"end":56},{"verbatim":"Hauduroy","normalized":"Hauduroy","wordType":"AUTHOR_WORD","start":60,"end":68},{"verbatim":"1937","normalized":"1937","wordType":"YEAR","start":69,"end":73}],"id":"bb6e2a9f-6813-5b00-9a3f-e12a085e515e","parserVersion":"test_version"}
```

### Bacteria genus homonym

Name: Actinomyces cardiffensis

Canonical: Actinomyces cardiffensis

Authorship:

```json
{"parsed":true,"quality":1,"qualityWarnings":[{"quality":1,"warning":"The genus is a homonym of a bacterial genus"}],"verbatim":"Actinomyces cardiffensis","normalized":"Actinomyces cardiffensis","canonical":{"stemmed":"Actinomyces cardiffens","simple":"Actinomyces cardiffensis","full":"Actinomyces cardiffensis"},"cardinality":2,"bacteria":"maybe","details":{"species":{"genus":"Actinomyces","species":"cardiffensis"}},"words":[{"verbatim":"Actinomyces","normalized":"Actinomyces","wordType":"GENUS","start":0,"end":11},{"verbatim":"cardiffensis","normalized":"cardiffensis","wordType":"SPECIES","start":12,"end":24}],"id":"fc1def53-81ba-5d2f-9f4c-0d9ac591cd13","parserVersion":"test_version"}
```

### Bacteria with pathovar rank

Name: Xanthomonas axonopodis pv. phaseoli

Canonical: Xanthomonas axonopodis pv. phaseoli

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Xanthomonas axonopodis pv. phaseoli","normalized":"Xanthomonas axonopodis pv. phaseoli","canonical":{"stemmed":"Xanthomonas axonopod phaseol","simple":"Xanthomonas axonopodis phaseoli","full":"Xanthomonas axonopodis pv. phaseoli"},"cardinality":3,"bacteria":"yes","details":{"infraspecies":{"genus":"Xanthomonas","species":"axonopodis","infraspecies":[{"value":"phaseoli","rank":"pv."}]}},"words":[{"verbatim":"Xanthomonas","normalized":"Xanthomonas","wordType":"GENUS","start":0,"end":11},{"verbatim":"axonopodis","normalized":"axonopodis","wordType":"SPECIES","start":12,"end":22},{"verbatim":"pv.","normalized":"pv.","wordType":"RANK","start":23,"end":26},{"verbatim":"phaseoli","normalized":"phaseoli","wordType":"INFRASPECIES","start":27,"end":35}],"id":"ea35594e-41c7-5706-b3b8-bb1b94d11a77","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pathovar. phaseoli

Canonical: Xanthomonas axonopodis pathovar. phaseoli

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Xanthomonas axonopodis pathovar. phaseoli","normalized":"Xanthomonas axonopodis pathovar. phaseoli","canonical":{"stemmed":"Xanthomonas axonopod phaseol","simple":"Xanthomonas axonopodis phaseoli","full":"Xanthomonas axonopodis pathovar. phaseoli"},"cardinality":3,"bacteria":"yes","details":{"infraspecies":{"genus":"Xanthomonas","species":"axonopodis","infraspecies":[{"value":"phaseoli","rank":"pathovar."}]}},"words":[{"verbatim":"Xanthomonas","normalized":"Xanthomonas","wordType":"GENUS","start":0,"end":11},{"verbatim":"axonopodis","normalized":"axonopodis","wordType":"SPECIES","start":12,"end":22},{"verbatim":"pathovar.","normalized":"pathovar.","wordType":"RANK","start":23,"end":32},{"verbatim":"phaseoli","normalized":"phaseoli","wordType":"INFRASPECIES","start":33,"end":41}],"id":"816ce2bc-4cdc-59ab-8900-e4414e8d2125","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pathovar.

Canonical: Xanthomonas axonopodis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Xanthomonas axonopodis pathovar.","normalized":"Xanthomonas axonopodis","canonical":{"stemmed":"Xanthomonas axonopod","simple":"Xanthomonas axonopodis","full":"Xanthomonas axonopodis"},"cardinality":2,"bacteria":"yes","tail":" pathovar.","details":{"species":{"genus":"Xanthomonas","species":"axonopodis"}},"words":[{"verbatim":"Xanthomonas","normalized":"Xanthomonas","wordType":"GENUS","start":0,"end":11},{"verbatim":"axonopodis","normalized":"axonopodis","wordType":"SPECIES","start":12,"end":22}],"id":"851a86de-df67-5fba-b3f7-73937a5edbce","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pv.

Canonical: Xanthomonas axonopodis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Xanthomonas axonopodis pv.","normalized":"Xanthomonas axonopodis","canonical":{"stemmed":"Xanthomonas axonopod","simple":"Xanthomonas axonopodis","full":"Xanthomonas axonopodis"},"cardinality":2,"bacteria":"yes","tail":" pv.","details":{"species":{"genus":"Xanthomonas","species":"axonopodis"}},"words":[{"verbatim":"Xanthomonas","normalized":"Xanthomonas","wordType":"GENUS","start":0,"end":11},{"verbatim":"axonopodis","normalized":"axonopodis","wordType":"SPECIES","start":12,"end":22}],"id":"0c0ce6dd-e5ea-5c17-8be3-c381ff662f12","parserVersion":"test_version"}
```

### "Stray" ex is not parsed as species

Name: Pelargonium cucullatum ssp. cucullatum (L.) L'Her. ex [Soland.]

Canonical: Pelargonium cucullatum subsp. cucullatum

Authorship: (L.) L'Her.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pelargonium cucullatum ssp. cucullatum (L.) L'Her. ex [Soland.]","normalized":"Pelargonium cucullatum subsp. cucullatum (L.) L'Her.","canonical":{"stemmed":"Pelargonium cucullat cucullat","simple":"Pelargonium cucullatum cucullatum","full":"Pelargonium cucullatum subsp. cucullatum"},"cardinality":3,"authorship":{"verbatim":"(L.) L'Her.","normalized":"(L.) L'Her.","authors":["L.","L'Her."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Her."]}},"tail":" ex [Soland.]","details":{"infraspecies":{"genus":"Pelargonium","species":"cucullatum","infraspecies":[{"value":"cucullatum","rank":"subsp.","authorship":{"verbatim":"(L.) L'Her.","normalized":"(L.) L'Her.","authors":["L.","L'Her."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Her."]}}}]}},"words":[{"verbatim":"Pelargonium","normalized":"Pelargonium","wordType":"GENUS","start":0,"end":11},{"verbatim":"cucullatum","normalized":"cucullatum","wordType":"SPECIES","start":12,"end":22},{"verbatim":"ssp.","normalized":"subsp.","wordType":"RANK","start":23,"end":27},{"verbatim":"cucullatum","normalized":"cucullatum","wordType":"INFRASPECIES","start":28,"end":38},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":40,"end":42},{"verbatim":"L'Her.","normalized":"L'Her.","wordType":"AUTHOR_WORD","start":44,"end":50}],"id":"83811b74-a581-5801-aa49-d4eab6775fdb","parserVersion":"test_version"}
```

<!-- not dealing with ex. gr for now -->
Name: Acastella ex gr. rouaulti

Canonical: Acastella

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acastella ex gr. rouaulti","normalized":"Acastella","canonical":{"stemmed":"Acastella","simple":"Acastella","full":"Acastella"},"cardinality":1,"tail":" ex gr. rouaulti","details":{"uninomial":{"uninomial":"Acastella"}},"words":[{"verbatim":"Acastella","normalized":"Acastella","wordType":"UNINOMIAL","start":0,"end":9}],"id":"c1864b52-848a-5de7-8f2d-a3cfe2025c40","parserVersion":"test_version"}
```

### Authorship in upper case

Name: Lecanora strobilinoides GIRALT & GÓMEZ-BOLEA

Canonical: Lecanora strobilinoides

Authorship: Giralt & Gómez-Bolea

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"}],"verbatim":"Lecanora strobilinoides GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Lecanora strobilinoides Giralt \u0026 Gómez-Bolea","canonical":{"stemmed":"Lecanora strobilinoid","simple":"Lecanora strobilinoides","full":"Lecanora strobilinoides"},"cardinality":2,"authorship":{"verbatim":"GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Giralt \u0026 Gómez-Bolea","authors":["Giralt","Gómez-Bolea"],"originalAuth":{"authors":["Giralt","Gómez-Bolea"]}},"details":{"species":{"genus":"Lecanora","species":"strobilinoides","authorship":{"verbatim":"GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Giralt \u0026 Gómez-Bolea","authors":["Giralt","Gómez-Bolea"],"originalAuth":{"authors":["Giralt","Gómez-Bolea"]}}}},"words":[{"verbatim":"Lecanora","normalized":"Lecanora","wordType":"GENUS","start":0,"end":8},{"verbatim":"strobilinoides","normalized":"strobilinoides","wordType":"SPECIES","start":9,"end":23},{"verbatim":"GIRALT","normalized":"Giralt","wordType":"AUTHOR_WORD","start":24,"end":30},{"verbatim":"GÓMEZ-BOLEA","normalized":"Gómez-Bolea","wordType":"AUTHOR_WORD","start":33,"end":44}],"id":"f2bfaa25-c25f-5a31-90c6-a19bd4dc23f4","parserVersion":"test_version"}
```

### Numbers and letters separated with '-' are not parsed as authors

Name: Astatotilapia cf. bloyeti OS-2017

Canonical: Astatotilapia bloyeti

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"},{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Astatotilapia cf. bloyeti OS-2017","normalized":"Astatotilapia cf. bloyeti","canonical":{"stemmed":"Astatotilapia bloyet","simple":"Astatotilapia bloyeti","full":"Astatotilapia bloyeti"},"cardinality":2,"surrogate":"COMPARISON","tail":" OS-2017","details":{"comparison":{"genus":"Astatotilapia","species":"bloyeti","comparisonMarker":"cf."}},"words":[{"verbatim":"Astatotilapia","normalized":"Astatotilapia","wordType":"GENUS","start":0,"end":13},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":14,"end":17},{"verbatim":"bloyeti","normalized":"bloyeti","wordType":"SPECIES","start":18,"end":25}],"id":"c841aa1d-78ea-5b6a-93fc-e18c54164144","parserVersion":"test_version"}
```

### Double parenthesis
Name: Eichornia crassipes ( (Martius) ) Solms-Laub.

Canonical: Eichornia crassipes

Authorship: (Martius) Solms-Laub.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship in double parentheses"}],"verbatim":"Eichornia crassipes ( (Martius) ) Solms-Laub.","normalized":"Eichornia crassipes (Martius) Solms-Laub.","canonical":{"stemmed":"Eichornia crassip","simple":"Eichornia crassipes","full":"Eichornia crassipes"},"cardinality":2,"authorship":{"verbatim":"( (Martius) ) Solms-Laub.","normalized":"(Martius) Solms-Laub.","authors":["Martius","Solms-Laub."],"originalAuth":{"authors":["Martius"]},"combinationAuth":{"authors":["Solms-Laub."]}},"details":{"species":{"genus":"Eichornia","species":"crassipes","authorship":{"verbatim":"( (Martius) ) Solms-Laub.","normalized":"(Martius) Solms-Laub.","authors":["Martius","Solms-Laub."],"originalAuth":{"authors":["Martius"]},"combinationAuth":{"authors":["Solms-Laub."]}}}},"words":[{"verbatim":"Eichornia","normalized":"Eichornia","wordType":"GENUS","start":0,"end":9},{"verbatim":"crassipes","normalized":"crassipes","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Martius","normalized":"Martius","wordType":"AUTHOR_WORD","start":23,"end":30},{"verbatim":"Solms-Laub.","normalized":"Solms-Laub.","wordType":"AUTHOR_WORD","start":34,"end":45}],"id":"95b90189-29d1-51ca-a1fa-0fb1c19a1fa1","parserVersion":"test_version"}
```

### Numbers at the start/middle of names

Name: Nesomyrmex madecassus_01m

Canonical: Nesomyrmex

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nesomyrmex madecassus_01m","normalized":"Nesomyrmex","canonical":{"stemmed":"Nesomyrmex","simple":"Nesomyrmex","full":"Nesomyrmex"},"cardinality":1,"tail":" madecassus_01m","details":{"uninomial":{"uninomial":"Nesomyrmex"}},"words":[{"verbatim":"Nesomyrmex","normalized":"Nesomyrmex","wordType":"UNINOMIAL","start":0,"end":10}],"id":"30dd0028-1ad4-5f65-ba5e-3df4963825d2","parserVersion":"test_version"}
```

Name: Hypochrys0des

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Hypochrys0des","cardinality":0,"id":"859c6279-20ea-5e60-9b7d-0c5283e06377","parserVersion":"test_version"}
```

Name: Hypochrys0des Leraut 1981

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Hypochrys0des Leraut 1981","cardinality":0,"id":"c053bbbf-de6c-5b22-a0f9-0803093b9b2d","parserVersion":"test_version"}
```

Name: Phyllodoce mucosa 0ersted, 1843

Canonical: Phyllodoce mucosa

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Phyllodoce mucosa 0ersted, 1843","normalized":"Phyllodoce mucosa","canonical":{"stemmed":"Phyllodoce mucos","simple":"Phyllodoce mucosa","full":"Phyllodoce mucosa"},"cardinality":2,"tail":" 0ersted, 1843","details":{"species":{"genus":"Phyllodoce","species":"mucosa"}},"words":[{"verbatim":"Phyllodoce","normalized":"Phyllodoce","wordType":"GENUS","start":0,"end":10},{"verbatim":"mucosa","normalized":"mucosa","wordType":"SPECIES","start":11,"end":17}],"id":"52695b7b-ebef-5624-9ccf-f9d07cd8133c","parserVersion":"test_version"}
```

Name: Attelabus 0l.

Canonical: Attelabus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Attelabus 0l.","normalized":"Attelabus","canonical":{"stemmed":"Attelabus","simple":"Attelabus","full":"Attelabus"},"cardinality":1,"tail":" 0l.","details":{"uninomial":{"uninomial":"Attelabus"}},"words":[{"verbatim":"Attelabus","normalized":"Attelabus","wordType":"UNINOMIAL","start":0,"end":9}],"id":"b9edee54-a7ae-525a-a319-ffeed18cf88a","parserVersion":"test_version"}
```

Name: Acrobothrium 0lsson 1872

Canonical: Acrobothrium

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acrobothrium 0lsson 1872","normalized":"Acrobothrium","canonical":{"stemmed":"Acrobothrium","simple":"Acrobothrium","full":"Acrobothrium"},"cardinality":1,"tail":" 0lsson 1872","details":{"uninomial":{"uninomial":"Acrobothrium"}},"words":[{"verbatim":"Acrobothrium","normalized":"Acrobothrium","wordType":"UNINOMIAL","start":0,"end":12}],"id":"2edfbcca-af28-5498-a762-663e5d5b9f73","parserVersion":"test_version"}
```

Name: Staphylinus haemrrhoidalis 0l. nec Gmel

Canonical: Staphylinus haemrrhoidalis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Staphylinus haemrrhoidalis 0l. nec Gmel","normalized":"Staphylinus haemrrhoidalis","canonical":{"stemmed":"Staphylinus haemrrhoidal","simple":"Staphylinus haemrrhoidalis","full":"Staphylinus haemrrhoidalis"},"cardinality":2,"tail":" 0l. nec Gmel","details":{"species":{"genus":"Staphylinus","species":"haemrrhoidalis"}},"words":[{"verbatim":"Staphylinus","normalized":"Staphylinus","wordType":"GENUS","start":0,"end":11},{"verbatim":"haemrrhoidalis","normalized":"haemrrhoidalis","wordType":"SPECIES","start":12,"end":26}],"id":"3ef602da-08a5-5acf-8f8a-9c515373ccda","parserVersion":"test_version"}
```

Name: Ea92virus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Ea92virus","cardinality":0,"virus":true,"id":"2465682c-cd5c-5408-859b-8bcc5489125f","parserVersion":"test_version"}
```

### Year without authorship

<!--TODO: collect year information-->
Name: Acarospora cratericola 1929

Canonical: Acarospora cratericola

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola 1929","normalized":"Acarospora cratericola","canonical":{"stemmed":"Acarospora cratericol","simple":"Acarospora cratericola","full":"Acarospora cratericola"},"cardinality":2,"tail":" 1929","details":{"species":{"genus":"Acarospora","species":"cratericola"}},"words":[{"verbatim":"Acarospora","normalized":"Acarospora","wordType":"GENUS","start":0,"end":10},{"verbatim":"cratericola","normalized":"cratericola","wordType":"SPECIES","start":11,"end":22}],"id":"11335046-cf05-5571-84bb-f9c8a4b2d8de","parserVersion":"test_version"}
```

Name: Goggia gemmula 1996

Canonical: Goggia gemmula

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Goggia gemmula 1996","normalized":"Goggia gemmula","canonical":{"stemmed":"Goggia gemmul","simple":"Goggia gemmula","full":"Goggia gemmula"},"cardinality":2,"tail":" 1996","details":{"species":{"genus":"Goggia","species":"gemmula"}},"words":[{"verbatim":"Goggia","normalized":"Goggia","wordType":"GENUS","start":0,"end":6},{"verbatim":"gemmula","normalized":"gemmula","wordType":"SPECIES","start":7,"end":14}],"id":"707ab43c-41bd-56bc-b2aa-96db4913ad35","parserVersion":"test_version"}
```

### Year range

Name: Eurodryas orientalis Herrich-Schäffer 1845-1847

Canonical: Eurodryas orientalis

Authorship: Herrich-Schäffer (1845)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Eurodryas orientalis Herrich-Schäffer 1845-1847","normalized":"Eurodryas orientalis Herrich-Schäffer (1845)","canonical":{"stemmed":"Eurodryas oriental","simple":"Eurodryas orientalis","full":"Eurodryas orientalis"},"cardinality":2,"authorship":{"verbatim":"Herrich-Schäffer 1845-1847","normalized":"Herrich-Schäffer (1845)","year":"(1845)","authors":["Herrich-Schäffer"],"originalAuth":{"authors":["Herrich-Schäffer"],"year":{"year":"1845","isApproximate":true}}},"details":{"species":{"genus":"Eurodryas","species":"orientalis","authorship":{"verbatim":"Herrich-Schäffer 1845-1847","normalized":"Herrich-Schäffer (1845)","year":"(1845)","authors":["Herrich-Schäffer"],"originalAuth":{"authors":["Herrich-Schäffer"],"year":{"year":"1845","isApproximate":true}}}}},"words":[{"verbatim":"Eurodryas","normalized":"Eurodryas","wordType":"GENUS","start":0,"end":9},{"verbatim":"orientalis","normalized":"orientalis","wordType":"SPECIES","start":10,"end":20},{"verbatim":"Herrich-Schäffer","normalized":"Herrich-Schäffer","wordType":"AUTHOR_WORD","start":21,"end":37},{"verbatim":"1845","normalized":"1845","wordType":"APPROXIMATE_YEAR","start":38,"end":42}],"id":"5fbca057-cd1e-5334-b6d3-496559b31818","parserVersion":"test_version"}
```

Name: Tridentella tangeroae Bruce, 1987-92

Canonical: Tridentella tangeroae

Authorship: Bruce (1987)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Tridentella tangeroae Bruce, 1987-92","normalized":"Tridentella tangeroae Bruce (1987)","canonical":{"stemmed":"Tridentella tangero","simple":"Tridentella tangeroae","full":"Tridentella tangeroae"},"cardinality":2,"authorship":{"verbatim":"Bruce, 1987-92","normalized":"Bruce (1987)","year":"(1987)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"1987","isApproximate":true}}},"details":{"species":{"genus":"Tridentella","species":"tangeroae","authorship":{"verbatim":"Bruce, 1987-92","normalized":"Bruce (1987)","year":"(1987)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"1987","isApproximate":true}}}}},"words":[{"verbatim":"Tridentella","normalized":"Tridentella","wordType":"GENUS","start":0,"end":11},{"verbatim":"tangeroae","normalized":"tangeroae","wordType":"SPECIES","start":12,"end":21},{"verbatim":"Bruce","normalized":"Bruce","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"1987","normalized":"1987","wordType":"APPROXIMATE_YEAR","start":29,"end":33}],"id":"6c943756-7f67-51ee-9c06-8f9016538be6","parserVersion":"test_version"}
```

Name: Macroplectra unicolor Moore, 1858/59

Canonical: Macroplectra unicolor

Authorship: Moore (1858)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Macroplectra unicolor Moore, 1858/59","normalized":"Macroplectra unicolor Moore (1858)","canonical":{"stemmed":"Macroplectra unicolor","simple":"Macroplectra unicolor","full":"Macroplectra unicolor"},"cardinality":2,"authorship":{"verbatim":"Moore, 1858/59","normalized":"Moore (1858)","year":"(1858)","authors":["Moore"],"originalAuth":{"authors":["Moore"],"year":{"year":"1858","isApproximate":true}}},"details":{"species":{"genus":"Macroplectra","species":"unicolor","authorship":{"verbatim":"Moore, 1858/59","normalized":"Moore (1858)","year":"(1858)","authors":["Moore"],"originalAuth":{"authors":["Moore"],"year":{"year":"1858","isApproximate":true}}}}},"words":[{"verbatim":"Macroplectra","normalized":"Macroplectra","wordType":"GENUS","start":0,"end":12},{"verbatim":"unicolor","normalized":"unicolor","wordType":"SPECIES","start":13,"end":21},{"verbatim":"Moore","normalized":"Moore","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"1858","normalized":"1858","wordType":"APPROXIMATE_YEAR","start":29,"end":33}],"id":"d6fc4a96-793c-58ce-9926-ec40281062b2","parserVersion":"test_version"}
```

Name: Seryda basirei Druce, 1891/901

Canonical: Seryda basirei

Authorship: Druce (1891)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Seryda basirei Druce, 1891/901","normalized":"Seryda basirei Druce (1891)","canonical":{"stemmed":"Seryda basire","simple":"Seryda basirei","full":"Seryda basirei"},"cardinality":2,"authorship":{"verbatim":"Druce, 1891/901","normalized":"Druce (1891)","year":"(1891)","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1891","isApproximate":true}}},"details":{"species":{"genus":"Seryda","species":"basirei","authorship":{"verbatim":"Druce, 1891/901","normalized":"Druce (1891)","year":"(1891)","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1891","isApproximate":true}}}}},"words":[{"verbatim":"Seryda","normalized":"Seryda","wordType":"GENUS","start":0,"end":6},{"verbatim":"basirei","normalized":"basirei","wordType":"SPECIES","start":7,"end":14},{"verbatim":"Druce","normalized":"Druce","wordType":"AUTHOR_WORD","start":15,"end":20},{"verbatim":"1891","normalized":"1891","wordType":"APPROXIMATE_YEAR","start":22,"end":26}],"id":"574ff67d-f220-5c14-9634-fcadc3794891","parserVersion":"test_version"}
```

### Year with page number

Name: Recilia truncatus Dash & Viraktamath, 1998a: 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a: 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Recilia","normalized":"Recilia","wordType":"GENUS","start":0,"end":7},{"verbatim":"truncatus","normalized":"truncatus","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dash","normalized":"Dash","wordType":"AUTHOR_WORD","start":18,"end":22},{"verbatim":"Viraktamath","normalized":"Viraktamath","wordType":"AUTHOR_WORD","start":25,"end":36},{"verbatim":"1998a","normalized":"1998","wordType":"YEAR","start":38,"end":43}],"id":"227ada89-45e5-56a9-83ad-47bee641e373","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998: 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998: 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Recilia","normalized":"Recilia","wordType":"GENUS","start":0,"end":7},{"verbatim":"truncatus","normalized":"truncatus","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dash","normalized":"Dash","wordType":"AUTHOR_WORD","start":18,"end":22},{"verbatim":"Viraktamath","normalized":"Viraktamath","wordType":"AUTHOR_WORD","start":25,"end":36},{"verbatim":"1998","normalized":"1998","wordType":"YEAR","start":38,"end":42}],"id":"47a39cf1-7be1-5937-b8fa-03a1696c1de6","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998a:29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a:29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a:29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a:29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Recilia","normalized":"Recilia","wordType":"GENUS","start":0,"end":7},{"verbatim":"truncatus","normalized":"truncatus","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dash","normalized":"Dash","wordType":"AUTHOR_WORD","start":18,"end":22},{"verbatim":"Viraktamath","normalized":"Viraktamath","wordType":"AUTHOR_WORD","start":25,"end":36},{"verbatim":"1998a","normalized":"1998","wordType":"YEAR","start":38,"end":43}],"id":"68b51644-5fef-5d5f-819d-f5bf8c9e6051","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998a : 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a : 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a : 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a : 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Recilia","normalized":"Recilia","wordType":"GENUS","start":0,"end":7},{"verbatim":"truncatus","normalized":"truncatus","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Dash","normalized":"Dash","wordType":"AUTHOR_WORD","start":18,"end":22},{"verbatim":"Viraktamath","normalized":"Viraktamath","wordType":"AUTHOR_WORD","start":25,"end":36},{"verbatim":"1998a","normalized":"1998","wordType":"YEAR","start":38,"end":43}],"id":"08507e4f-412c-59c9-b1f2-906dd4b27aa8","parserVersion":"test_version"}
```

### Year in square brackets

Name: Anthoscopus Cabanis [1851]

Canonical: Anthoscopus

Authorship: Cabanis (1851)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"}],"verbatim":"Anthoscopus Cabanis [1851]","normalized":"Anthoscopus Cabanis (1851)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}}}},"words":[{"verbatim":"Anthoscopus","normalized":"Anthoscopus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Cabanis","normalized":"Cabanis","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"1851","normalized":"1851","wordType":"APPROXIMATE_YEAR","start":21,"end":25}],"id":"8d86299b-3028-5be2-b2f6-6e4897f4c748","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [185?]

Canonical: Anthoscopus

Authorship: Cabanis (185?)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [185?]","normalized":"Anthoscopus Cabanis (185?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [185?]","normalized":"Cabanis (185?)","year":"(185?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"185?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [185?]","normalized":"Cabanis (185?)","year":"(185?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"185?","isApproximate":true}}}}},"words":[{"verbatim":"Anthoscopus","normalized":"Anthoscopus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Cabanis","normalized":"Cabanis","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"185?","normalized":"185?","wordType":"APPROXIMATE_YEAR","start":21,"end":25}],"id":"3434c072-d015-5f54-ad32-45b01de7fd08","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851?]

Canonical: Anthoscopus

Authorship: Cabanis (1851?)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [1851?]","normalized":"Anthoscopus Cabanis (1851?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}}}},"words":[{"verbatim":"Anthoscopus","normalized":"Anthoscopus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Cabanis","normalized":"Cabanis","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"1851?","normalized":"1851?","wordType":"APPROXIMATE_YEAR","start":21,"end":26}],"id":"6b12b541-b58b-5f11-ba66-bb314b53813f","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851]

Canonical: Anthoscopus

Authorship: Cabanis (1851)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"}],"verbatim":"Anthoscopus Cabanis [1851]","normalized":"Anthoscopus Cabanis (1851)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}}}},"words":[{"verbatim":"Anthoscopus","normalized":"Anthoscopus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Cabanis","normalized":"Cabanis","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"1851","normalized":"1851","wordType":"APPROXIMATE_YEAR","start":21,"end":25}],"id":"8d86299b-3028-5be2-b2f6-6e4897f4c748","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851?]

Canonical: Anthoscopus

Authorship: Cabanis (1851?)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [1851?]","normalized":"Anthoscopus Cabanis (1851?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}}}},"words":[{"verbatim":"Anthoscopus","normalized":"Anthoscopus","wordType":"UNINOMIAL","start":0,"end":11},{"verbatim":"Cabanis","normalized":"Cabanis","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"1851?","normalized":"1851?","wordType":"APPROXIMATE_YEAR","start":21,"end":26}],"id":"6b12b541-b58b-5f11-ba66-bb314b53813f","parserVersion":"test_version"}
```

Name: Trismegistia monodii Ando, 1973 [1974]

Canonical: Trismegistia monodii

Authorship: Ando 1973

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Trismegistia monodii Ando, 1973 [1974]","normalized":"Trismegistia monodii Ando 1973","canonical":{"stemmed":"Trismegistia monod","simple":"Trismegistia monodii","full":"Trismegistia monodii"},"cardinality":2,"authorship":{"verbatim":"Ando, 1973","normalized":"Ando 1973","year":"1973","authors":["Ando"],"originalAuth":{"authors":["Ando"],"year":{"year":"1973"}}},"tail":" [1974]","details":{"species":{"genus":"Trismegistia","species":"monodii","authorship":{"verbatim":"Ando, 1973","normalized":"Ando 1973","year":"1973","authors":["Ando"],"originalAuth":{"authors":["Ando"],"year":{"year":"1973"}}}}},"words":[{"verbatim":"Trismegistia","normalized":"Trismegistia","wordType":"GENUS","start":0,"end":12},{"verbatim":"monodii","normalized":"monodii","wordType":"SPECIES","start":13,"end":20},{"verbatim":"Ando","normalized":"Ando","wordType":"AUTHOR_WORD","start":21,"end":25},{"verbatim":"1973","normalized":"1973","wordType":"YEAR","start":27,"end":31}],"id":"f396d2d0-b14e-537f-ae8f-c383310f813e","parserVersion":"test_version"}
```

Name: Zygaena witti Wiegel [1973]

Canonical: Zygaena witti

Authorship: Wiegel (1973)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brackets"}],"verbatim":"Zygaena witti Wiegel [1973]","normalized":"Zygaena witti Wiegel (1973)","canonical":{"stemmed":"Zygaena witt","simple":"Zygaena witti","full":"Zygaena witti"},"cardinality":2,"authorship":{"verbatim":"Wiegel [1973]","normalized":"Wiegel (1973)","year":"(1973)","authors":["Wiegel"],"originalAuth":{"authors":["Wiegel"],"year":{"year":"1973","isApproximate":true}}},"details":{"species":{"genus":"Zygaena","species":"witti","authorship":{"verbatim":"Wiegel [1973]","normalized":"Wiegel (1973)","year":"(1973)","authors":["Wiegel"],"originalAuth":{"authors":["Wiegel"],"year":{"year":"1973","isApproximate":true}}}}},"words":[{"verbatim":"Zygaena","normalized":"Zygaena","wordType":"GENUS","start":0,"end":7},{"verbatim":"witti","normalized":"witti","wordType":"SPECIES","start":8,"end":13},{"verbatim":"Wiegel","normalized":"Wiegel","wordType":"AUTHOR_WORD","start":14,"end":20},{"verbatim":"1973","normalized":"1973","wordType":"APPROXIMATE_YEAR","start":22,"end":26}],"id":"76eef612-f125-54f9-b241-6b3a9be0a6c6","parserVersion":"test_version"}
```

Name: Deyeuxia coarctata Kunth, 1815 [1816]

Canonical: Deyeuxia coarctata

Authorship: Kunth 1815

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Deyeuxia coarctata Kunth, 1815 [1816]","normalized":"Deyeuxia coarctata Kunth 1815","canonical":{"stemmed":"Deyeuxia coarctat","simple":"Deyeuxia coarctata","full":"Deyeuxia coarctata"},"cardinality":2,"authorship":{"verbatim":"Kunth, 1815","normalized":"Kunth 1815","year":"1815","authors":["Kunth"],"originalAuth":{"authors":["Kunth"],"year":{"year":"1815"}}},"tail":" [1816]","details":{"species":{"genus":"Deyeuxia","species":"coarctata","authorship":{"verbatim":"Kunth, 1815","normalized":"Kunth 1815","year":"1815","authors":["Kunth"],"originalAuth":{"authors":["Kunth"],"year":{"year":"1815"}}}}},"words":[{"verbatim":"Deyeuxia","normalized":"Deyeuxia","wordType":"GENUS","start":0,"end":8},{"verbatim":"coarctata","normalized":"coarctata","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Kunth","normalized":"Kunth","wordType":"AUTHOR_WORD","start":19,"end":24},{"verbatim":"1815","normalized":"1815","wordType":"YEAR","start":26,"end":30}],"id":"2f479365-40be-5181-b194-8a24fc743f73","parserVersion":"test_version"}
```

### Names with broken conversion between encodings

Name: Macrotes cordovaria Guen�e 1857

Canonical: Macrotes cordovaria

Authorship: Guen�e 1857

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Incorrect conversion to UTF-8"}],"verbatim":"Macrotes cordovaria Guen�e 1857","normalized":"Macrotes cordovaria Guen�e 1857","canonical":{"stemmed":"Macrotes cordouar","simple":"Macrotes cordovaria","full":"Macrotes cordovaria"},"cardinality":2,"authorship":{"verbatim":"Guen�e 1857","normalized":"Guen�e 1857","year":"1857","authors":["Guen�e"],"originalAuth":{"authors":["Guen�e"],"year":{"year":"1857"}}},"details":{"species":{"genus":"Macrotes","species":"cordovaria","authorship":{"verbatim":"Guen�e 1857","normalized":"Guen�e 1857","year":"1857","authors":["Guen�e"],"originalAuth":{"authors":["Guen�e"],"year":{"year":"1857"}}}}},"words":[{"verbatim":"Macrotes","normalized":"Macrotes","wordType":"GENUS","start":0,"end":8},{"verbatim":"cordovaria","normalized":"cordovaria","wordType":"SPECIES","start":9,"end":19},{"verbatim":"Guen�e","normalized":"Guen�e","wordType":"AUTHOR_WORD","start":20,"end":26},{"verbatim":"1857","normalized":"1857","wordType":"YEAR","start":27,"end":31}],"id":"9217d59c-d1e7-5c79-af65-f52623446c15","parserVersion":"test_version"}
```

Name: Fusinus eucos�nius

Canonical: Fusinus eucos�nius

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Incorrect conversion to UTF-8"}],"verbatim":"Fusinus eucos�nius","normalized":"Fusinus eucos�nius","canonical":{"stemmed":"Fusinus eucos�n","simple":"Fusinus eucos�nius","full":"Fusinus eucos�nius"},"cardinality":2,"details":{"species":{"genus":"Fusinus","species":"eucos�nius"}},"words":[{"verbatim":"Fusinus","normalized":"Fusinus","wordType":"GENUS","start":0,"end":7},{"verbatim":"eucos�nius","normalized":"eucos�nius","wordType":"SPECIES","start":8,"end":18}],"id":"157cf8c1-0b0d-5b81-a3a9-f02bdc1413a5","parserVersion":"test_version"}
```

### UTF-8 0xA0 character (NO_BREAK_SPACE)

Name: Byssochlamys fulva Olliver & G. Smith

Canonical: Byssochlamys fulva

Authorship: Olliver & G. Smith

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Byssochlamys fulva Olliver \u0026 G. Smith","normalized":"Byssochlamys fulva Olliver \u0026 G. Smith","canonical":{"stemmed":"Byssochlamys fulu","simple":"Byssochlamys fulva","full":"Byssochlamys fulva"},"cardinality":2,"authorship":{"verbatim":"Olliver \u0026 G. Smith","normalized":"Olliver \u0026 G. Smith","authors":["Olliver","G. Smith"],"originalAuth":{"authors":["Olliver","G. Smith"]}},"details":{"species":{"genus":"Byssochlamys","species":"fulva","authorship":{"verbatim":"Olliver \u0026 G. Smith","normalized":"Olliver \u0026 G. Smith","authors":["Olliver","G. Smith"],"originalAuth":{"authors":["Olliver","G. Smith"]}}}},"words":[{"verbatim":"Byssochlamys","normalized":"Byssochlamys","wordType":"GENUS","start":0,"end":12},{"verbatim":"fulva","normalized":"fulva","wordType":"SPECIES","start":13,"end":18},{"verbatim":"Olliver","normalized":"Olliver","wordType":"AUTHOR_WORD","start":19,"end":26},{"verbatim":"G.","normalized":"G.","wordType":"AUTHOR_WORD","start":29,"end":31},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":32,"end":37}],"id":"83523455-cfe4-5ff9-bc54-841f026576b7","parserVersion":"test_version"}
```

### UTF-8 0x3000 character (IDEOGRAPHIC_SPACE)

Name: Kinosternidae　Agassiz, 1857

Canonical: Kinosternidae

Authorship: Agassiz 1857

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Kinosternidae　Agassiz, 1857","normalized":"Kinosternidae Agassiz 1857","canonical":{"stemmed":"Kinosternidae","simple":"Kinosternidae","full":"Kinosternidae"},"cardinality":1,"authorship":{"verbatim":"Agassiz, 1857","normalized":"Agassiz 1857","year":"1857","authors":["Agassiz"],"originalAuth":{"authors":["Agassiz"],"year":{"year":"1857"}}},"details":{"uninomial":{"uninomial":"Kinosternidae","authorship":{"verbatim":"Agassiz, 1857","normalized":"Agassiz 1857","year":"1857","authors":["Agassiz"],"originalAuth":{"authors":["Agassiz"],"year":{"year":"1857"}}}}},"words":[{"verbatim":"Kinosternidae","normalized":"Kinosternidae","wordType":"UNINOMIAL","start":0,"end":13},{"verbatim":"Agassiz","normalized":"Agassiz","wordType":"AUTHOR_WORD","start":14,"end":21},{"verbatim":"1857","normalized":"1857","wordType":"YEAR","start":23,"end":27}],"id":"7e74b6b8-5242-5802-9238-320192f4eaa4","parserVersion":"test_version"}
```

### Punctuation in the end

Name: Melanius:

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Melanius:","cardinality":0,"id":"0a761224-66db-55b4-b6f0-85de52534125","parserVersion":"test_version"}
```

Name: Negalasa fumalis Barnes & McDunnough 1913. Next sentence

Canonical: Negalasa fumalis

Authorship: Barnes & McDunnough 1913

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis Barnes \u0026 McDunnough 1913. Next sentence","normalized":"Negalasa fumalis Barnes \u0026 McDunnough 1913","canonical":{"stemmed":"Negalasa fumal","simple":"Negalasa fumalis","full":"Negalasa fumalis"},"cardinality":2,"authorship":{"verbatim":"Barnes \u0026 McDunnough 1913.","normalized":"Barnes \u0026 McDunnough 1913","year":"1913","authors":["Barnes","McDunnough"],"originalAuth":{"authors":["Barnes","McDunnough"],"year":{"year":"1913"}}},"tail":" Next sentence","details":{"species":{"genus":"Negalasa","species":"fumalis","authorship":{"verbatim":"Barnes \u0026 McDunnough 1913.","normalized":"Barnes \u0026 McDunnough 1913","year":"1913","authors":["Barnes","McDunnough"],"originalAuth":{"authors":["Barnes","McDunnough"],"year":{"year":"1913"}}}}},"words":[{"verbatim":"Negalasa","normalized":"Negalasa","wordType":"GENUS","start":0,"end":8},{"verbatim":"fumalis","normalized":"fumalis","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Barnes","normalized":"Barnes","wordType":"AUTHOR_WORD","start":17,"end":23},{"verbatim":"McDunnough","normalized":"McDunnough","wordType":"AUTHOR_WORD","start":26,"end":36},{"verbatim":"1913","normalized":"1913","wordType":"YEAR","start":37,"end":41}],"id":"45b7343f-d42a-52d5-b0a4-25956d46427b","parserVersion":"test_version"}
```

Name: Negalasa fumalis. Next sentence

Canonical: Negalasa

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis. Next sentence","normalized":"Negalasa","canonical":{"stemmed":"Negalasa","simple":"Negalasa","full":"Negalasa"},"cardinality":1,"tail":" fumalis. Next sentence","details":{"uninomial":{"uninomial":"Negalasa"}},"words":[{"verbatim":"Negalasa","normalized":"Negalasa","wordType":"UNINOMIAL","start":0,"end":8}],"id":"ce740482-fa87-5d84-b335-1c063fd18de1","parserVersion":"test_version"}
```

Name: Negalasa fumalis, continuation of a sentence

Canonical: Negalasa

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis, continuation of a sentence","normalized":"Negalasa","canonical":{"stemmed":"Negalasa","simple":"Negalasa","full":"Negalasa"},"cardinality":1,"tail":" fumalis, continuation of a sentence","details":{"uninomial":{"uninomial":"Negalasa"}},"words":[{"verbatim":"Negalasa","normalized":"Negalasa","wordType":"UNINOMIAL","start":0,"end":8}],"id":"7862a3d9-ba4d-5f53-a106-ea048e558f1a","parserVersion":"test_version"}
```

Name: Negalasa fumalis Barnes; something else

Canonical: Negalasa fumalis

Authorship: Barnes

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis Barnes; something else","normalized":"Negalasa fumalis Barnes","canonical":{"stemmed":"Negalasa fumal","simple":"Negalasa fumalis","full":"Negalasa fumalis"},"cardinality":2,"authorship":{"verbatim":"Barnes","normalized":"Barnes","authors":["Barnes"],"originalAuth":{"authors":["Barnes"]}},"tail":"; something else","details":{"species":{"genus":"Negalasa","species":"fumalis","authorship":{"verbatim":"Barnes","normalized":"Barnes","authors":["Barnes"],"originalAuth":{"authors":["Barnes"]}}}},"words":[{"verbatim":"Negalasa","normalized":"Negalasa","wordType":"GENUS","start":0,"end":8},{"verbatim":"fumalis","normalized":"fumalis","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Barnes","normalized":"Barnes","wordType":"AUTHOR_WORD","start":17,"end":23}],"id":"6359dac4-1a88-5b41-86d3-9c01aaee4a2e","parserVersion":"test_version"}
```

Name: Negaprion brevirostris Negaprion brevirostris, the rest of the sentence

Canonical: Negaprion brevirostris

Authorship: Negaprion

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negaprion brevirostris Negaprion brevirostris, the rest of the sentence","normalized":"Negaprion brevirostris Negaprion","canonical":{"stemmed":"Negaprion breuirostr","simple":"Negaprion brevirostris","full":"Negaprion brevirostris"},"cardinality":2,"authorship":{"verbatim":"Negaprion","normalized":"Negaprion","authors":["Negaprion"],"originalAuth":{"authors":["Negaprion"]}},"tail":" brevirostris, the rest of the sentence","details":{"species":{"genus":"Negaprion","species":"brevirostris","authorship":{"verbatim":"Negaprion","normalized":"Negaprion","authors":["Negaprion"],"originalAuth":{"authors":["Negaprion"]}}}},"words":[{"verbatim":"Negaprion","normalized":"Negaprion","wordType":"GENUS","start":0,"end":9},{"verbatim":"brevirostris","normalized":"brevirostris","wordType":"SPECIES","start":10,"end":22},{"verbatim":"Negaprion","normalized":"Negaprion","wordType":"AUTHOR_WORD","start":23,"end":32}],"id":"619b95fa-017d-5b9b-b800-64ebd5ed433b","parserVersion":"test_version"}
```

Name: Negaprion fronto (Jordan and Gilbert, 1882):

Canonical: Negaprion fronto

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negaprion fronto (Jordan and Gilbert, 1882):","normalized":"Negaprion fronto","canonical":{"stemmed":"Negaprion front","simple":"Negaprion fronto","full":"Negaprion fronto"},"cardinality":2,"tail":" (Jordan and Gilbert, 1882):","details":{"species":{"genus":"Negaprion","species":"fronto"}},"words":[{"verbatim":"Negaprion","normalized":"Negaprion","wordType":"GENUS","start":0,"end":9},{"verbatim":"fronto","normalized":"fronto","wordType":"SPECIES","start":10,"end":16}],"id":"4bb6a543-d757-5fa5-ae8b-a5ac95722e1d","parserVersion":"test_version"}
```

### Names with 'ex' as sp. epithet

<!-- not dealing with this misspelling...-->
Name: Acanthochiton ex quisitus

Canonical: Acanthochiton

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acanthochiton ex quisitus","normalized":"Acanthochiton","canonical":{"stemmed":"Acanthochiton","simple":"Acanthochiton","full":"Acanthochiton"},"cardinality":1,"tail":" ex quisitus","details":{"uninomial":{"uninomial":"Acanthochiton"}},"words":[{"verbatim":"Acanthochiton","normalized":"Acanthochiton","wordType":"UNINOMIAL","start":0,"end":13}],"id":"00392ae2-1bd9-5a14-bea9-9d26f1107892","parserVersion":"test_version"}
```

### Names with Spanish 'y' instead of '&'

Name: Caloptenopsis crassiusculus (Martínez y Fernández-Castillo, 1896)

Canonical: Caloptenopsis crassiusculus

Authorship: (Martínez & Fernández-Castillo 1896)

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Caloptenopsis crassiusculus (Martínez y Fernández-Castillo, 1896)","normalized":"Caloptenopsis crassiusculus (Martínez \u0026 Fernández-Castillo 1896)","canonical":{"stemmed":"Caloptenopsis crassiuscul","simple":"Caloptenopsis crassiusculus","full":"Caloptenopsis crassiusculus"},"cardinality":2,"authorship":{"verbatim":"(Martínez y Fernández-Castillo, 1896)","normalized":"(Martínez \u0026 Fernández-Castillo 1896)","year":"1896","authors":["Martínez","Fernández-Castillo"],"originalAuth":{"authors":["Martínez","Fernández-Castillo"],"year":{"year":"1896"}}},"details":{"species":{"genus":"Caloptenopsis","species":"crassiusculus","authorship":{"verbatim":"(Martínez y Fernández-Castillo, 1896)","normalized":"(Martínez \u0026 Fernández-Castillo 1896)","year":"1896","authors":["Martínez","Fernández-Castillo"],"originalAuth":{"authors":["Martínez","Fernández-Castillo"],"year":{"year":"1896"}}}}},"words":[{"verbatim":"Caloptenopsis","normalized":"Caloptenopsis","wordType":"GENUS","start":0,"end":13},{"verbatim":"crassiusculus","normalized":"crassiusculus","wordType":"SPECIES","start":14,"end":27},{"verbatim":"Martínez","normalized":"Martínez","wordType":"AUTHOR_WORD","start":29,"end":37},{"verbatim":"Fernández-Castillo","normalized":"Fernández-Castillo","wordType":"AUTHOR_WORD","start":40,"end":58},{"verbatim":"1896","normalized":"1896","wordType":"YEAR","start":60,"end":64}],"id":"0080ce8d-aba5-512d-8e33-8ee3914e386a","parserVersion":"test_version"}
```

Name: Dicranum saxatile Lagasca y Segura, García & Clemente y Rubio, 1802

Canonical: Dicranum saxatile

Authorship: Lagasca, Segura, García, Clemente & Rubio 1802

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Dicranum saxatile Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Dicranum saxatile Lagasca, Segura, García, Clemente \u0026 Rubio 1802","canonical":{"stemmed":"Dicranum saxatil","simple":"Dicranum saxatile","full":"Dicranum saxatile"},"cardinality":2,"authorship":{"verbatim":"Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Lagasca, Segura, García, Clemente \u0026 Rubio 1802","year":"1802","authors":["Lagasca","Segura","García","Clemente","Rubio"],"originalAuth":{"authors":["Lagasca","Segura","García","Clemente","Rubio"],"year":{"year":"1802"}}},"details":{"species":{"genus":"Dicranum","species":"saxatile","authorship":{"verbatim":"Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Lagasca, Segura, García, Clemente \u0026 Rubio 1802","year":"1802","authors":["Lagasca","Segura","García","Clemente","Rubio"],"originalAuth":{"authors":["Lagasca","Segura","García","Clemente","Rubio"],"year":{"year":"1802"}}}}},"words":[{"verbatim":"Dicranum","normalized":"Dicranum","wordType":"GENUS","start":0,"end":8},{"verbatim":"saxatile","normalized":"saxatile","wordType":"SPECIES","start":9,"end":17},{"verbatim":"Lagasca","normalized":"Lagasca","wordType":"AUTHOR_WORD","start":18,"end":25},{"verbatim":"Segura","normalized":"Segura","wordType":"AUTHOR_WORD","start":28,"end":34},{"verbatim":"García","normalized":"García","wordType":"AUTHOR_WORD","start":36,"end":42},{"verbatim":"Clemente","normalized":"Clemente","wordType":"AUTHOR_WORD","start":45,"end":53},{"verbatim":"Rubio","normalized":"Rubio","wordType":"AUTHOR_WORD","start":56,"end":61},{"verbatim":"1802","normalized":"1802","wordType":"YEAR","start":63,"end":67}],"id":"39054306-2722-5119-a040-f8671b5b31a0","parserVersion":"test_version"}
```

Name: Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger & Coronado 1967

Canonical: Carabus hendrichsi

Authorship: Bolvar, Pieltain, Rotger & Coronado 1967

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Carabus (Tanaocarabus) hendrichsi Bolvar, Pieltain, Rotger \u0026 Coronado 1967","canonical":{"stemmed":"Carabus hendrichs","simple":"Carabus hendrichsi","full":"Carabus hendrichsi"},"cardinality":2,"authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado"],"year":{"year":"1967"}}},"details":{"species":{"genus":"Carabus","subgenus":"Tanaocarabus","species":"hendrichsi","authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado"],"year":{"year":"1967"}}}}},"words":[{"verbatim":"Carabus","normalized":"Carabus","wordType":"GENUS","start":0,"end":7},{"verbatim":"Tanaocarabus","normalized":"Tanaocarabus","wordType":"INFRA_GENUS","start":9,"end":21},{"verbatim":"hendrichsi","normalized":"hendrichsi","wordType":"SPECIES","start":23,"end":33},{"verbatim":"Bolvar","normalized":"Bolvar","wordType":"AUTHOR_WORD","start":34,"end":40},{"verbatim":"Pieltain","normalized":"Pieltain","wordType":"AUTHOR_WORD","start":43,"end":51},{"verbatim":"Rotger","normalized":"Rotger","wordType":"AUTHOR_WORD","start":53,"end":59},{"verbatim":"Coronado","normalized":"Coronado","wordType":"AUTHOR_WORD","start":62,"end":70},{"verbatim":"1967","normalized":"1967","wordType":"YEAR","start":71,"end":75}],"id":"519c0687-2303-5b8c-a69f-68e2bd055b5e","parserVersion":"test_version"}
```

### Normalize atypical dashes

Name: Passalus (Pertinax) gaboi Jiménez‑Ferbans & Reyes‑Castillo, 2022

Canonical: Passalus gaboi

Authorship: Jiménez-Ferbans & Reyes-Castillo 2022

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Atypical hyphen character"}],"verbatim":"Passalus (Pertinax) gaboi Jiménez‑Ferbans \u0026 Reyes‑Castillo, 2022","normalized":"Passalus (Pertinax) gaboi Jiménez-Ferbans \u0026 Reyes-Castillo 2022","canonical":{"stemmed":"Passalus gabo","simple":"Passalus gaboi","full":"Passalus gaboi"},"cardinality":2,"authorship":{"verbatim":"Jiménez‑Ferbans \u0026 Reyes‑Castillo, 2022","normalized":"Jiménez-Ferbans \u0026 Reyes-Castillo 2022","year":"2022","authors":["Jiménez-Ferbans","Reyes-Castillo"],"originalAuth":{"authors":["Jiménez-Ferbans","Reyes-Castillo"],"year":{"year":"2022"}}},"details":{"species":{"genus":"Passalus","subgenus":"Pertinax","species":"gaboi","authorship":{"verbatim":"Jiménez‑Ferbans \u0026 Reyes‑Castillo, 2022","normalized":"Jiménez-Ferbans \u0026 Reyes-Castillo 2022","year":"2022","authors":["Jiménez-Ferbans","Reyes-Castillo"],"originalAuth":{"authors":["Jiménez-Ferbans","Reyes-Castillo"],"year":{"year":"2022"}}}}},"words":[{"verbatim":"Passalus","normalized":"Passalus","wordType":"GENUS","start":0,"end":8},{"verbatim":"Pertinax","normalized":"Pertinax","wordType":"INFRA_GENUS","start":10,"end":18},{"verbatim":"gaboi","normalized":"gaboi","wordType":"SPECIES","start":20,"end":25},{"verbatim":"Jiménez‑Ferbans","normalized":"Jiménez-Ferbans","wordType":"AUTHOR_WORD","start":26,"end":41},{"verbatim":"Reyes‑Castillo","normalized":"Reyes-Castillo","wordType":"AUTHOR_WORD","start":44,"end":58},{"verbatim":"2022","normalized":"2022","wordType":"YEAR","start":60,"end":64}],"id":"4cf1b94a-b80f-5666-92d0-5f7fc2076ce8","parserVersion":"test_version"}
```

### Discard apostrophes at the start and end of words

Name: Labeotropheus trewavasae 'albino

Canonical: Labeotropheus trewavasae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Labeotropheus trewavasae 'albino","normalized":"Labeotropheus trewavasae","canonical":{"stemmed":"Labeotropheus trewauas","simple":"Labeotropheus trewavasae","full":"Labeotropheus trewavasae"},"cardinality":2,"tail":" 'albino","details":{"species":{"genus":"Labeotropheus","species":"trewavasae"}},"words":[{"verbatim":"Labeotropheus","normalized":"Labeotropheus","wordType":"GENUS","start":0,"end":13},{"verbatim":"trewavasae","normalized":"trewavasae","wordType":"SPECIES","start":14,"end":24}],"id":"0cb9e0ae-1201-5023-8d20-689d60a3e20c","parserVersion":"test_version"}
```

Name: Labeotropheus trewavasae albino'

Canonical: Labeotropheus trewavasae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Labeotropheus trewavasae albino'","normalized":"Labeotropheus trewavasae","canonical":{"stemmed":"Labeotropheus trewauas","simple":"Labeotropheus trewavasae","full":"Labeotropheus trewavasae"},"cardinality":2,"tail":" albino'","details":{"species":{"genus":"Labeotropheus","species":"trewavasae"}},"words":[{"verbatim":"Labeotropheus","normalized":"Labeotropheus","wordType":"GENUS","start":0,"end":13},{"verbatim":"trewavasae","normalized":"trewavasae","wordType":"SPECIES","start":14,"end":24}],"id":"f190cdee-14f0-5174-947d-476dab6baeff","parserVersion":"test_version"}
```

Name: Phedimus takesimensis (Nakai) 't Hart

Canonical: Phedimus takesimensis

Authorship: (Nakai) 't Hart

```json
{"parsed":true,"quality":1,"verbatim":"Phedimus takesimensis (Nakai) 't Hart","normalized":"Phedimus takesimensis (Nakai) 't Hart","canonical":{"stemmed":"Phedimus takesimens","simple":"Phedimus takesimensis","full":"Phedimus takesimensis"},"cardinality":2,"authorship":{"verbatim":"(Nakai) 't Hart","normalized":"(Nakai) 't Hart","authors":["Nakai","'t Hart"],"originalAuth":{"authors":["Nakai"]},"combinationAuth":{"authors":["'t Hart"]}},"details":{"species":{"genus":"Phedimus","species":"takesimensis","authorship":{"verbatim":"(Nakai) 't Hart","normalized":"(Nakai) 't Hart","authors":["Nakai","'t Hart"],"originalAuth":{"authors":["Nakai"]},"combinationAuth":{"authors":["'t Hart"]}}}},"words":[{"verbatim":"Phedimus","normalized":"Phedimus","wordType":"GENUS","start":0,"end":8},{"verbatim":"takesimensis","normalized":"takesimensis","wordType":"SPECIES","start":9,"end":21},{"verbatim":"Nakai","normalized":"Nakai","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"'t","normalized":"'t","wordType":"AUTHOR_WORD","start":30,"end":32},{"verbatim":"Hart","normalized":"Hart","wordType":"AUTHOR_WORD","start":33,"end":37}],"id":"14379aa4-1eb9-5ef7-b355-7e3ef3c1fe5e","parserVersion":"test_version"}
```

### Discard apostrophe with dash (rare, needs further investigation)

<!-- correctly parsed -->
Name: Solanum juzepczukii janck'o-ckaisalla

Canonical: Solanum juzepczukii jancko-ckaisalla

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Solanum juzepczukii janck'o-ckaisalla","normalized":"Solanum juzepczukii jancko-ckaisalla","canonical":{"stemmed":"Solanum iuzepczuk iancko-ckaisall","simple":"Solanum juzepczukii jancko-ckaisalla","full":"Solanum juzepczukii jancko-ckaisalla"},"cardinality":3,"details":{"infraspecies":{"genus":"Solanum","species":"juzepczukii","infraspecies":[{"value":"jancko-ckaisalla"}]}},"words":[{"verbatim":"Solanum","normalized":"Solanum","wordType":"GENUS","start":0,"end":7},{"verbatim":"juzepczukii","normalized":"juzepczukii","wordType":"SPECIES","start":8,"end":19},{"verbatim":"janck'o-ckaisalla","normalized":"jancko-ckaisalla","wordType":"INFRASPECIES","start":20,"end":37}],"id":"9ec56934-e986-5392-a531-55d97e5e9dd1","parserVersion":"test_version"}
```

### Possible canonical

Name: Morea (Morea) burtius 2342343242 23424322342 23424234

Canonical: Morea burtius

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Morea (Morea) burtius 2342343242 23424322342 23424234","normalized":"Morea (Morea) burtius","canonical":{"stemmed":"Morea burt","simple":"Morea burtius","full":"Morea burtius"},"cardinality":2,"tail":" 2342343242 23424322342 23424234","details":{"species":{"genus":"Morea","subgenus":"Morea","species":"burtius"}},"words":[{"verbatim":"Morea","normalized":"Morea","wordType":"GENUS","start":0,"end":5},{"verbatim":"Morea","normalized":"Morea","wordType":"INFRA_GENUS","start":7,"end":12},{"verbatim":"burtius","normalized":"burtius","wordType":"SPECIES","start":14,"end":21}],"id":"03f59808-c30e-55da-bea5-27aa035feb5d","parserVersion":"test_version"}
```

Name: Verpericola megasoma ""Dall" Pils.

Canonical: Verpericola megasoma

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Cultivar epithet"}],"verbatim":"Verpericola megasoma \"\"Dall\" Pils.","normalized":"Verpericola megasoma","canonical":{"stemmed":"Verpericola megasom","simple":"Verpericola megasoma","full":"Verpericola megasoma"},"cardinality":2,"tail":" Pils.","details":{"species":{"genus":"Verpericola","species":"megasoma","cultivar":"‘\"Dall’"}},"words":[{"verbatim":"Verpericola","normalized":"Verpericola","wordType":"GENUS","start":0,"end":11},{"verbatim":"megasoma","normalized":"megasoma","wordType":"SPECIES","start":12,"end":20},{"verbatim":"\"Dall","normalized":"‘\"Dall’","wordType":"CULTIVAR","start":22,"end":27}],"id":"cebb60d9-fc8e-5fa0-874a-ae21819b242b","parserVersion":"test_version"}
```

Name: Verpericola megasoma "Dall" Pils.

Canonical: Verpericola megasoma

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Cultivar epithet"}],"verbatim":"Verpericola megasoma \"Dall\" Pils.","normalized":"Verpericola megasoma","canonical":{"stemmed":"Verpericola megasom","simple":"Verpericola megasoma","full":"Verpericola megasoma"},"cardinality":2,"tail":" Pils.","details":{"species":{"genus":"Verpericola","species":"megasoma","cultivar":"‘Dall’"}},"words":[{"verbatim":"Verpericola","normalized":"Verpericola","wordType":"GENUS","start":0,"end":11},{"verbatim":"megasoma","normalized":"megasoma","wordType":"SPECIES","start":12,"end":20},{"verbatim":"Dall","normalized":"‘Dall’","wordType":"CULTIVAR","start":22,"end":26}],"id":"02011460-ba94-5162-98c9-4064a700c7f8","parserVersion":"test_version"}
```



Name: Moraea spathulata ( (L. f. Klatt

Canonical: Moraea spathulata

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Moraea spathulata ( (L. f. Klatt","normalized":"Moraea spathulata","canonical":{"stemmed":"Moraea spathulat","simple":"Moraea spathulata","full":"Moraea spathulata"},"cardinality":2,"tail":" ( (L. f. Klatt","details":{"species":{"genus":"Moraea","species":"spathulata"}},"words":[{"verbatim":"Moraea","normalized":"Moraea","wordType":"GENUS","start":0,"end":6},{"verbatim":"spathulata","normalized":"spathulata","wordType":"SPECIES","start":7,"end":17}],"id":"21cb8638-ff53-534f-b816-1e15ecbb818b","parserVersion":"test_version"}
```

Name: Stewartia micrantha (Chun) Sealy, Bot. Mag. 176: t. 510. 1967.

Canonical: Stewartia micrantha

Authorship: (Chun) Sealy & Bot. Mag.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Stewartia micrantha (Chun) Sealy, Bot. Mag. 176: t. 510. 1967.","normalized":"Stewartia micrantha (Chun) Sealy \u0026 Bot. Mag.","canonical":{"stemmed":"Stewartia micranth","simple":"Stewartia micrantha","full":"Stewartia micrantha"},"cardinality":2,"authorship":{"verbatim":"(Chun) Sealy, Bot. Mag.","normalized":"(Chun) Sealy \u0026 Bot. Mag.","authors":["Chun","Sealy","Bot. Mag."],"originalAuth":{"authors":["Chun"]},"combinationAuth":{"authors":["Sealy","Bot. Mag."]}},"tail":" 176: t. 510. 1967.","details":{"species":{"genus":"Stewartia","species":"micrantha","authorship":{"verbatim":"(Chun) Sealy, Bot. Mag.","normalized":"(Chun) Sealy \u0026 Bot. Mag.","authors":["Chun","Sealy","Bot. Mag."],"originalAuth":{"authors":["Chun"]},"combinationAuth":{"authors":["Sealy","Bot. Mag."]}}}},"words":[{"verbatim":"Stewartia","normalized":"Stewartia","wordType":"GENUS","start":0,"end":9},{"verbatim":"micrantha","normalized":"micrantha","wordType":"SPECIES","start":10,"end":19},{"verbatim":"Chun","normalized":"Chun","wordType":"AUTHOR_WORD","start":21,"end":25},{"verbatim":"Sealy","normalized":"Sealy","wordType":"AUTHOR_WORD","start":27,"end":32},{"verbatim":"Bot.","normalized":"Bot.","wordType":"AUTHOR_WORD","start":34,"end":38},{"verbatim":"Mag.","normalized":"Mag.","wordType":"AUTHOR_WORD","start":39,"end":43}],"id":"7a4ffc19-61a9-551b-bea2-ebb0f5fe9c5a","parserVersion":"test_version"}
```

Name: Pyrobaculum neutrophilum V24Sta

Canonical: Pyrobaculum neutrophilum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pyrobaculum neutrophilum V24Sta","normalized":"Pyrobaculum neutrophilum","canonical":{"stemmed":"Pyrobaculum neutrophil","simple":"Pyrobaculum neutrophilum","full":"Pyrobaculum neutrophilum"},"cardinality":2,"tail":" V24Sta","details":{"species":{"genus":"Pyrobaculum","species":"neutrophilum"}},"words":[{"verbatim":"Pyrobaculum","normalized":"Pyrobaculum","wordType":"GENUS","start":0,"end":11},{"verbatim":"neutrophilum","normalized":"neutrophilum","wordType":"SPECIES","start":12,"end":24}],"id":"6d0be585-ec54-5662-9d30-1d369ecf2a64","parserVersion":"test_version"}
```

Name: Rana aurora Baird and Girard, 1852; H.B. Shaffer et al., 2004

Canonical: Rana aurora

Authorship: Baird & Girard 1852

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Rana aurora Baird and Girard, 1852; H.B. Shaffer et al., 2004","normalized":"Rana aurora Baird \u0026 Girard 1852","canonical":{"stemmed":"Rana auror","simple":"Rana aurora","full":"Rana aurora"},"cardinality":2,"authorship":{"verbatim":"Baird and Girard, 1852","normalized":"Baird \u0026 Girard 1852","year":"1852","authors":["Baird","Girard"],"originalAuth":{"authors":["Baird","Girard"],"year":{"year":"1852"}}},"tail":"; H.B. Shaffer et al., 2004","details":{"species":{"genus":"Rana","species":"aurora","authorship":{"verbatim":"Baird and Girard, 1852","normalized":"Baird \u0026 Girard 1852","year":"1852","authors":["Baird","Girard"],"originalAuth":{"authors":["Baird","Girard"],"year":{"year":"1852"}}}}},"words":[{"verbatim":"Rana","normalized":"Rana","wordType":"GENUS","start":0,"end":4},{"verbatim":"aurora","normalized":"aurora","wordType":"SPECIES","start":5,"end":11},{"verbatim":"Baird","normalized":"Baird","wordType":"AUTHOR_WORD","start":12,"end":17},{"verbatim":"Girard","normalized":"Girard","wordType":"AUTHOR_WORD","start":22,"end":28},{"verbatim":"1852","normalized":"1852","wordType":"YEAR","start":30,"end":34}],"id":"f0fa6cd1-8018-5fec-92ad-1bda9ac929ca","parserVersion":"test_version"}
```

Name: Agropyron pectiniforme var. karabaljikji ined.?

Canonical: Agropyron pectiniforme var. karabaljikji

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Agropyron pectiniforme var. karabaljikji ined.?","normalized":"Agropyron pectiniforme var. karabaljikji","canonical":{"stemmed":"Agropyron pectiniform karabaliik","simple":"Agropyron pectiniforme karabaljikji","full":"Agropyron pectiniforme var. karabaljikji"},"cardinality":3,"tail":" ined.?","details":{"infraspecies":{"genus":"Agropyron","species":"pectiniforme","infraspecies":[{"value":"karabaljikji","rank":"var."}]}},"words":[{"verbatim":"Agropyron","normalized":"Agropyron","wordType":"GENUS","start":0,"end":9},{"verbatim":"pectiniforme","normalized":"pectiniforme","wordType":"SPECIES","start":10,"end":22},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":23,"end":27},{"verbatim":"karabaljikji","normalized":"karabaljikji","wordType":"INFRASPECIES","start":28,"end":40}],"id":"e951b7d4-0009-54df-9de6-efbb392dc8d6","parserVersion":"test_version"}
```

Name: Staphylococcus hyicus chromogenes Devriese et al. 1978 (Approved Lists 1980).

Canonical: Staphylococcus hyicus chromogenes

Authorship: Devriese et al. 1978

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Staphylococcus hyicus chromogenes Devriese et al. 1978 (Approved Lists 1980).","normalized":"Staphylococcus hyicus chromogenes Devriese et al. 1978","canonical":{"stemmed":"Staphylococcus hyic chromogen","simple":"Staphylococcus hyicus chromogenes","full":"Staphylococcus hyicus chromogenes"},"cardinality":3,"authorship":{"verbatim":"Devriese et al. 1978","normalized":"Devriese et al. 1978","year":"1978","authors":["Devriese et al."],"originalAuth":{"authors":["Devriese et al."],"year":{"year":"1978"}}},"bacteria":"yes","tail":" (Approved Lists 1980).","details":{"infraspecies":{"genus":"Staphylococcus","species":"hyicus","infraspecies":[{"value":"chromogenes","authorship":{"verbatim":"Devriese et al. 1978","normalized":"Devriese et al. 1978","year":"1978","authors":["Devriese et al."],"originalAuth":{"authors":["Devriese et al."],"year":{"year":"1978"}}}}]}},"words":[{"verbatim":"Staphylococcus","normalized":"Staphylococcus","wordType":"GENUS","start":0,"end":14},{"verbatim":"hyicus","normalized":"hyicus","wordType":"SPECIES","start":15,"end":21},{"verbatim":"chromogenes","normalized":"chromogenes","wordType":"INFRASPECIES","start":22,"end":33},{"verbatim":"Devriese","normalized":"Devriese","wordType":"AUTHOR_WORD","start":34,"end":42},{"verbatim":"et al.","normalized":"et al.","wordType":"AUTHOR_WORD","start":43,"end":49},{"verbatim":"1978","normalized":"1978","wordType":"YEAR","start":50,"end":54}],"id":"ec17eb44-742c-5325-aca6-e33a0888ef0d","parserVersion":"test_version"}
```

### Treating `& al.` as `et al.`

Name: Adonis cyllenea Boiss. & al.

Canonical: Adonis cyllenea

Authorship: Boiss. et al.

```json
{"parsed":true,"quality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al.","normalized":"Adonis cyllenea Boiss. et al.","canonical":{"stemmed":"Adonis cyllene","simple":"Adonis cyllenea","full":"Adonis cyllenea"},"cardinality":2,"authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"details":{"species":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}}}},"words":[{"verbatim":"Adonis","normalized":"Adonis","wordType":"GENUS","start":0,"end":6},{"verbatim":"cyllenea","normalized":"cyllenea","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"\u0026 al.","normalized":"et al.","wordType":"AUTHOR_WORD","start":23,"end":28}],"id":"a7c2cb28-2ec2-55b5-88a2-6cfd633cbd00","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al

Canonical: Adonis cyllenea

Authorship: Boiss. et al.

```json
{"parsed":true,"quality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al","normalized":"Adonis cyllenea Boiss. et al.","canonical":{"stemmed":"Adonis cyllene","simple":"Adonis cyllenea","full":"Adonis cyllenea"},"cardinality":2,"authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"details":{"species":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}}}},"words":[{"verbatim":"Adonis","normalized":"Adonis","wordType":"GENUS","start":0,"end":6},{"verbatim":"cyllenea","normalized":"cyllenea","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"\u0026 al","normalized":"et al.","wordType":"AUTHOR_WORD","start":23,"end":27}],"id":"85e122ea-f581-5d4b-a29f-b87c48d0a716","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al. var. paryadrica Boiss.

Canonical: Adonis cyllenea var. paryadrica

Authorship: Boiss.

```json
{"parsed":true,"quality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al. var. paryadrica Boiss.","normalized":"Adonis cyllenea Boiss. et al. var. paryadrica Boiss.","canonical":{"stemmed":"Adonis cyllene paryadric","simple":"Adonis cyllenea paryadrica","full":"Adonis cyllenea var. paryadrica"},"cardinality":3,"authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}},"details":{"infraspecies":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"infraspecies":[{"value":"paryadrica","rank":"var.","authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}}}]}},"words":[{"verbatim":"Adonis","normalized":"Adonis","wordType":"GENUS","start":0,"end":6},{"verbatim":"cyllenea","normalized":"cyllenea","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"\u0026 al.","normalized":"et al.","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":29,"end":33},{"verbatim":"paryadrica","normalized":"paryadrica","wordType":"INFRASPECIES","start":34,"end":44},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":45,"end":51}],"id":"6bc790ae-210d-518e-9e20-2d4d517a08ef","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al var. paryadrica Boiss.

Canonical: Adonis cyllenea var. paryadrica

Authorship: Boiss.

```json
{"parsed":true,"quality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al var. paryadrica Boiss.","normalized":"Adonis cyllenea Boiss. et al. var. paryadrica Boiss.","canonical":{"stemmed":"Adonis cyllene paryadric","simple":"Adonis cyllenea paryadrica","full":"Adonis cyllenea var. paryadrica"},"cardinality":3,"authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}},"details":{"infraspecies":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"infraspecies":[{"value":"paryadrica","rank":"var.","authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}}}]}},"words":[{"verbatim":"Adonis","normalized":"Adonis","wordType":"GENUS","start":0,"end":6},{"verbatim":"cyllenea","normalized":"cyllenea","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":16,"end":22},{"verbatim":"\u0026 al","normalized":"et al.","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":28,"end":32},{"verbatim":"paryadrica","normalized":"paryadrica","wordType":"INFRASPECIES","start":33,"end":43},{"verbatim":"Boiss.","normalized":"Boiss.","wordType":"AUTHOR_WORD","start":44,"end":50}],"id":"eb7aee15-e462-5189-8335-a3a323be6907","parserVersion":"test_version"}
```

Name: Adetus fuscoapicalis Souza f. et al. 2001

Canonical: Adetus fuscoapicalis

Authorship: Souza fil. et al. 2001

```json
{"parsed":true,"quality":1,"verbatim":"Adetus fuscoapicalis Souza f. et al. 2001","normalized":"Adetus fuscoapicalis Souza fil. et al. 2001","canonical":{"stemmed":"Adetus fuscoapical","simple":"Adetus fuscoapicalis","full":"Adetus fuscoapicalis"},"cardinality":2,"authorship":{"verbatim":"Souza f. et al. 2001","normalized":"Souza fil. et al. 2001","year":"2001","authors":["Souza fil. et al."],"originalAuth":{"authors":["Souza fil. et al."],"year":{"year":"2001"}}},"details":{"species":{"genus":"Adetus","species":"fuscoapicalis","authorship":{"verbatim":"Souza f. et al. 2001","normalized":"Souza fil. et al. 2001","year":"2001","authors":["Souza fil. et al."],"originalAuth":{"authors":["Souza fil. et al."],"year":{"year":"2001"}}}}},"words":[{"verbatim":"Adetus","normalized":"Adetus","wordType":"GENUS","start":0,"end":6},{"verbatim":"fuscoapicalis","normalized":"fuscoapicalis","wordType":"SPECIES","start":7,"end":20},{"verbatim":"Souza","normalized":"Souza","wordType":"AUTHOR_WORD","start":21,"end":26},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":27,"end":29},{"verbatim":"et al.","normalized":"et al.","wordType":"AUTHOR_WORD","start":30,"end":36},{"verbatim":"2001","normalized":"2001","wordType":"YEAR","start":37,"end":41}],"id":"08b8a86b-2f1d-5739-81f1-a5703c124130","parserVersion":"test_version"}
```

Name: Sterigmostemon rhodanthum Rech. f. et al. in Rech. f.

Canonical: Sterigmostemon rhodanthum

Authorship: Rech. fil. et al. ex Rech. fil.

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Sterigmostemon rhodanthum Rech. f. et al. in Rech. f.","normalized":"Sterigmostemon rhodanthum Rech. fil. et al. ex Rech. fil.","canonical":{"stemmed":"Sterigmostemon rhodanth","simple":"Sterigmostemon rhodanthum","full":"Sterigmostemon rhodanthum"},"cardinality":2,"authorship":{"verbatim":"Rech. f. et al. in Rech. f.","normalized":"Rech. fil. et al. ex Rech. fil.","authors":["Rech. fil. et al.","Rech. fil."],"originalAuth":{"authors":["Rech. fil. et al."],"exAuthors":{"authors":["Rech. fil."]}}},"details":{"species":{"genus":"Sterigmostemon","species":"rhodanthum","authorship":{"verbatim":"Rech. f. et al. in Rech. f.","normalized":"Rech. fil. et al. ex Rech. fil.","authors":["Rech. fil. et al.","Rech. fil."],"originalAuth":{"authors":["Rech. fil. et al."],"exAuthors":{"authors":["Rech. fil."]}}}}},"words":[{"verbatim":"Sterigmostemon","normalized":"Sterigmostemon","wordType":"GENUS","start":0,"end":14},{"verbatim":"rhodanthum","normalized":"rhodanthum","wordType":"SPECIES","start":15,"end":25},{"verbatim":"Rech.","normalized":"Rech.","wordType":"AUTHOR_WORD","start":26,"end":31},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":32,"end":34},{"verbatim":"et al.","normalized":"et al.","wordType":"AUTHOR_WORD","start":35,"end":41},{"verbatim":"Rech.","normalized":"Rech.","wordType":"AUTHOR_WORD","start":45,"end":50},{"verbatim":"f.","normalized":"fil.","wordType":"AUTHOR_WORD_FILIUS","start":51,"end":53}],"id":"7352ecfa-8253-574c-8b37-c0586ae48f5d","parserVersion":"test_version"}
```

### Authors do not start with apostrophe

Name: Nereidavus kulkovi 'Kulkov

Canonical: Nereidavus kulkovi

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nereidavus kulkovi 'Kulkov","normalized":"Nereidavus kulkovi","canonical":{"stemmed":"Nereidavus kulkou","simple":"Nereidavus kulkovi","full":"Nereidavus kulkovi"},"cardinality":2,"tail":" 'Kulkov","details":{"species":{"genus":"Nereidavus","species":"kulkovi"}},"words":[{"verbatim":"Nereidavus","normalized":"Nereidavus","wordType":"GENUS","start":0,"end":10},{"verbatim":"kulkovi","normalized":"kulkovi","wordType":"SPECIES","start":11,"end":18}],"id":"6a4999cd-95cc-509d-8e0a-26a0dfcef67d","parserVersion":"test_version"}
```

### Epithets do not start or end with a dash

Name: Abryna -petri Paiva, 1860

Canonical: Abryna

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abryna -petri Paiva, 1860","normalized":"Abryna","canonical":{"stemmed":"Abryna","simple":"Abryna","full":"Abryna"},"cardinality":1,"tail":" -petri Paiva, 1860","details":{"uninomial":{"uninomial":"Abryna"}},"words":[{"verbatim":"Abryna","normalized":"Abryna","wordType":"UNINOMIAL","start":0,"end":6}],"id":"6ccc6217-9084-5b31-81f7-6b4cd7963f65","parserVersion":"test_version"}
```

Name: Abryna petri- Paiva, 1860

Canonical: Abryna

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abryna petri- Paiva, 1860","normalized":"Abryna","canonical":{"stemmed":"Abryna","simple":"Abryna","full":"Abryna"},"cardinality":1,"tail":" petri- Paiva, 1860","details":{"uninomial":{"uninomial":"Abryna"}},"words":[{"verbatim":"Abryna","normalized":"Abryna","wordType":"UNINOMIAL","start":0,"end":6}],"id":"b1e37ace-3ca8-5274-bd93-7333aa3e5223","parserVersion":"test_version"}
```

### Names that contain "of"

Name: Musca capraria Trustees of the British Museum (Natural History), 1939

Canonical: Musca capraria

Authorship: Trustees

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Musca capraria Trustees of the British Museum (Natural History), 1939","normalized":"Musca capraria Trustees","canonical":{"stemmed":"Musca caprar","simple":"Musca capraria","full":"Musca capraria"},"cardinality":2,"authorship":{"verbatim":"Trustees","normalized":"Trustees","authors":["Trustees"],"originalAuth":{"authors":["Trustees"]}},"tail":" of the British Museum (Natural History), 1939","details":{"species":{"genus":"Musca","species":"capraria","authorship":{"verbatim":"Trustees","normalized":"Trustees","authors":["Trustees"],"originalAuth":{"authors":["Trustees"]}}}},"words":[{"verbatim":"Musca","normalized":"Musca","wordType":"GENUS","start":0,"end":5},{"verbatim":"capraria","normalized":"capraria","wordType":"SPECIES","start":6,"end":14},{"verbatim":"Trustees","normalized":"Trustees","wordType":"AUTHOR_WORD","start":15,"end":23}],"id":"aa70cf4b-14bb-57a3-9fe1-0a9a544a16da","parserVersion":"test_version"}
```

Name: Nassellarid genera of uncertain affinities

Canonical: Nassellarid genera

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nassellarid genera of uncertain affinities","normalized":"Nassellarid genera","canonical":{"stemmed":"Nassellarid gener","simple":"Nassellarid genera","full":"Nassellarid genera"},"cardinality":2,"tail":" of uncertain affinities","details":{"species":{"genus":"Nassellarid","species":"genera"}},"words":[{"verbatim":"Nassellarid","normalized":"Nassellarid","wordType":"GENUS","start":0,"end":11},{"verbatim":"genera","normalized":"genera","wordType":"SPECIES","start":12,"end":18}],"id":"ca46eccc-6b42-5faf-be0f-aad069d3e3dd","parserVersion":"test_version"}
```

Name: Natica of nidus

Canonical: Natica

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Natica of nidus","normalized":"Natica","canonical":{"stemmed":"Natica","simple":"Natica","full":"Natica"},"cardinality":1,"tail":" of nidus","details":{"uninomial":{"uninomial":"Natica"}},"words":[{"verbatim":"Natica","normalized":"Natica","wordType":"UNINOMIAL","start":0,"end":6}],"id":"6a049500-f407-56e7-80b4-41ab91f64b8c","parserVersion":"test_version"}
```

Name: Neritina chemmoi Reeve var of cornea Linn

Canonical: Neritina chemmoi

Authorship: Reeve

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Neritina chemmoi Reeve var of cornea Linn","normalized":"Neritina chemmoi Reeve","canonical":{"stemmed":"Neritina chemmo","simple":"Neritina chemmoi","full":"Neritina chemmoi"},"cardinality":2,"authorship":{"verbatim":"Reeve","normalized":"Reeve","authors":["Reeve"],"originalAuth":{"authors":["Reeve"]}},"tail":" var of cornea Linn","details":{"species":{"genus":"Neritina","species":"chemmoi","authorship":{"verbatim":"Reeve","normalized":"Reeve","authors":["Reeve"],"originalAuth":{"authors":["Reeve"]}}}},"words":[{"verbatim":"Neritina","normalized":"Neritina","wordType":"GENUS","start":0,"end":8},{"verbatim":"chemmoi","normalized":"chemmoi","wordType":"SPECIES","start":9,"end":16},{"verbatim":"Reeve","normalized":"Reeve","wordType":"AUTHOR_WORD","start":17,"end":22}],"id":"d6cbded0-dc9b-5da2-8fb9-8d8b124cc5b4","parserVersion":"test_version"}
```

### Cultivars

Name: Sarracenia flava 'Maxima'

Canonical: Sarracenia flava

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Cultivar epithet"}],"verbatim":"Sarracenia flava 'Maxima'","normalized":"Sarracenia flava","canonical":{"stemmed":"Sarracenia flau","simple":"Sarracenia flava","full":"Sarracenia flava"},"cardinality":2,"details":{"species":{"genus":"Sarracenia","species":"flava","cultivar":"‘Maxima’"}},"words":[{"verbatim":"Sarracenia","normalized":"Sarracenia","wordType":"GENUS","start":0,"end":10},{"verbatim":"flava","normalized":"flava","wordType":"SPECIES","start":11,"end":16},{"verbatim":"Maxima","normalized":"‘Maxima’","wordType":"CULTIVAR","start":18,"end":24}],"id":"39178008-65ee-5de3-af88-63ffdd67e00b","parserVersion":"test_version"}
```

### "Open taxonomy" with ranks unfinished

Name: Alyxia reinwardti var

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti var","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" var","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"words":[{"verbatim":"Alyxia","normalized":"Alyxia","wordType":"GENUS","start":0,"end":6},{"verbatim":"reinwardti","normalized":"reinwardti","wordType":"SPECIES","start":7,"end":17}],"id":"2f0ee2be-8d37-5e43-9eed-776c17f47e93","parserVersion":"test_version"}
```

Name: Alyxia reinwardti var.

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti var.","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" var.","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"words":[{"verbatim":"Alyxia","normalized":"Alyxia","wordType":"GENUS","start":0,"end":6},{"verbatim":"reinwardti","normalized":"reinwardti","wordType":"SPECIES","start":7,"end":17}],"id":"aed34708-82ed-52e4-876f-d4468af73fc3","parserVersion":"test_version"}
```

Name: Alyxia reinwardti ssp

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti ssp","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" ssp","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"words":[{"verbatim":"Alyxia","normalized":"Alyxia","wordType":"GENUS","start":0,"end":6},{"verbatim":"reinwardti","normalized":"reinwardti","wordType":"SPECIES","start":7,"end":17}],"id":"760486d1-93ed-55c5-ade1-ba2c5b2aa900","parserVersion":"test_version"}
```

Name: Alyxia reinwardti ssp.

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti ssp.","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" ssp.","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"words":[{"verbatim":"Alyxia","normalized":"Alyxia","wordType":"GENUS","start":0,"end":6},{"verbatim":"reinwardti","normalized":"reinwardti","wordType":"SPECIES","start":7,"end":17}],"id":"72b5072a-d952-54f8-aea1-5b5bd3c65c45","parserVersion":"test_version"}
```

Name: Alaria spp

Canonical: Alaria

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Alaria spp","normalized":"Alaria","canonical":{"stemmed":"Alaria","simple":"Alaria","full":"Alaria"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Alaria","approximationMarker":"spp"}},"words":[{"verbatim":"Alaria","normalized":"Alaria","wordType":"GENUS","start":0,"end":6},{"verbatim":"spp","normalized":"spp","wordType":"APPROXIMATION_MARKER","start":7,"end":10}],"id":"5b31e830-ccf6-5918-94c5-75c4db7ef302","parserVersion":"test_version"}
```

Name: Alaria spp.

Canonical: Alaria

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Alaria spp.","normalized":"Alaria","canonical":{"stemmed":"Alaria","simple":"Alaria","full":"Alaria"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Alaria","approximationMarker":"spp."}},"words":[{"verbatim":"Alaria","normalized":"Alaria","wordType":"GENUS","start":0,"end":6},{"verbatim":"spp.","normalized":"spp.","wordType":"APPROXIMATION_MARKER","start":7,"end":11}],"id":"d1cd4f1a-f511-5d5a-8f41-64911995fdec","parserVersion":"test_version"}
```

Name: Xenodon sp

Canonical: Xenodon

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Xenodon sp","normalized":"Xenodon","canonical":{"stemmed":"Xenodon","simple":"Xenodon","full":"Xenodon"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Xenodon","approximationMarker":"sp"}},"words":[{"verbatim":"Xenodon","normalized":"Xenodon","wordType":"GENUS","start":0,"end":7},{"verbatim":"sp","normalized":"sp","wordType":"APPROXIMATION_MARKER","start":8,"end":10}],"id":"7b0cb348-7fe9-5248-b396-b0336225ba2a","parserVersion":"test_version"}
```

Name: Xenodon sp.

Canonical: Xenodon

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Xenodon sp.","normalized":"Xenodon","canonical":{"stemmed":"Xenodon","simple":"Xenodon","full":"Xenodon"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Xenodon","approximationMarker":"sp."}},"words":[{"verbatim":"Xenodon","normalized":"Xenodon","wordType":"GENUS","start":0,"end":7},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":8,"end":11}],"id":"77b6718f-a26e-5ddf-a4cf-119e972cd015","parserVersion":"test_version"}
```

Name: Formicidae cf.

Canonical: Formicidae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Formicidae cf.","normalized":"Formicidae cf.","canonical":{"stemmed":"Formicidae","simple":"Formicidae","full":"Formicidae"},"cardinality":1,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Formicidae","comparisonMarker":"cf."}},"words":[{"verbatim":"Formicidae","normalized":"Formicidae","wordType":"GENUS","start":0,"end":10},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":11,"end":14}],"id":"61f9ebc4-346e-5857-ab45-38808ff1c960","parserVersion":"test_version"}
```

Name: Formicidae cf

Canonical: Formicidae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Formicidae cf","normalized":"Formicidae cf.","canonical":{"stemmed":"Formicidae","simple":"Formicidae","full":"Formicidae"},"cardinality":1,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Formicidae","comparisonMarker":"cf."}},"words":[{"verbatim":"Formicidae","normalized":"Formicidae","wordType":"GENUS","start":0,"end":10},{"verbatim":"cf","normalized":"cf.","wordType":"COMPARISON_MARKER","start":11,"end":13}],"id":"90473425-7ce1-5ec6-8160-737646816ea7","parserVersion":"test_version"}
```

Name: Arctostaphylos preglauca cf.

Canonical: Arctostaphylos preglauca

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Arctostaphylos preglauca cf.","normalized":"Arctostaphylos preglauca cf.","canonical":{"stemmed":"Arctostaphylos preglauc","simple":"Arctostaphylos preglauca","full":"Arctostaphylos preglauca"},"cardinality":2,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Arctostaphylos","species":"preglauca","comparisonMarker":"cf."}},"words":[{"verbatim":"Arctostaphylos","normalized":"Arctostaphylos","wordType":"GENUS","start":0,"end":14},{"verbatim":"preglauca","normalized":"preglauca","wordType":"SPECIES","start":15,"end":24},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":25,"end":28}],"id":"246b43d4-9786-5157-8d35-b81a470e6379","parserVersion":"test_version"}
```

Name: Albinaria brevicollis cf. sica Fuchs & Kaufel 1936

Canonical: Albinaria brevicollis sica

Authorship: Fuchs & Kaufel 1936

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Albinaria brevicollis cf. sica Fuchs \u0026 Kaufel 1936","normalized":"Albinaria brevicollis cf. sica Fuchs \u0026 Kaufel 1936","canonical":{"stemmed":"Albinaria breuicoll sic","simple":"Albinaria brevicollis sica","full":"Albinaria brevicollis sica"},"cardinality":3,"authorship":{"verbatim":"Fuchs \u0026 Kaufel 1936","normalized":"Fuchs \u0026 Kaufel 1936","year":"1936","authors":["Fuchs","Kaufel"],"originalAuth":{"authors":["Fuchs","Kaufel"],"year":{"year":"1936"}}},"surrogate":"COMPARISON","details":{"comparison":{"genus":"Albinaria","species":"brevicollis","infraspecies":{"value":"sica","authorship":{"verbatim":"Fuchs \u0026 Kaufel 1936","normalized":"Fuchs \u0026 Kaufel 1936","year":"1936","authors":["Fuchs","Kaufel"],"originalAuth":{"authors":["Fuchs","Kaufel"],"year":{"year":"1936"}}}},"comparisonMarker":"cf."}},"words":[{"verbatim":"Albinaria","normalized":"Albinaria","wordType":"GENUS","start":0,"end":9},{"verbatim":"brevicollis","normalized":"brevicollis","wordType":"SPECIES","start":10,"end":21},{"verbatim":"cf.","normalized":"cf.","wordType":"COMPARISON_MARKER","start":22,"end":25},{"verbatim":"sica","normalized":"sica","wordType":"INFRASPECIES","start":26,"end":30},{"verbatim":"Fuchs","normalized":"Fuchs","wordType":"AUTHOR_WORD","start":31,"end":36},{"verbatim":"Kaufel","normalized":"Kaufel","wordType":"AUTHOR_WORD","start":39,"end":45},{"verbatim":"1936","normalized":"1936","wordType":"YEAR","start":46,"end":50}],"id":"cc77e528-f730-563f-ba5c-5696ec456b69","parserVersion":"test_version"}
```

<!-- we do not support this -->

Name: Albinaria cf brevicollis sica Fuchs & Kaufel 1936

Canonical: Albinaria brevicollis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"},{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Albinaria cf brevicollis sica Fuchs \u0026 Kaufel 1936","normalized":"Albinaria cf. brevicollis","canonical":{"stemmed":"Albinaria breuicoll","simple":"Albinaria brevicollis","full":"Albinaria brevicollis"},"cardinality":2,"surrogate":"COMPARISON","tail":" sica Fuchs \u0026 Kaufel 1936","details":{"comparison":{"genus":"Albinaria","species":"brevicollis","comparisonMarker":"cf."}},"words":[{"verbatim":"Albinaria","normalized":"Albinaria","wordType":"GENUS","start":0,"end":9},{"verbatim":"cf","normalized":"cf.","wordType":"COMPARISON_MARKER","start":10,"end":12},{"verbatim":"brevicollis","normalized":"brevicollis","wordType":"SPECIES","start":13,"end":24}],"id":"8e2beae0-6a8e-54da-ac16-53de069fb3f0","parserVersion":"test_version"}
```

Name: Albinaria brevicollis cf

Canonical: Albinaria brevicollis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Albinaria brevicollis cf","normalized":"Albinaria brevicollis cf.","canonical":{"stemmed":"Albinaria breuicoll","simple":"Albinaria brevicollis","full":"Albinaria brevicollis"},"cardinality":2,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Albinaria","species":"brevicollis","comparisonMarker":"cf."}},"words":[{"verbatim":"Albinaria","normalized":"Albinaria","wordType":"GENUS","start":0,"end":9},{"verbatim":"brevicollis","normalized":"brevicollis","wordType":"SPECIES","start":10,"end":21},{"verbatim":"cf","normalized":"cf.","wordType":"COMPARISON_MARKER","start":22,"end":24}],"id":"591f1263-acfb-58f0-bcae-07a0e0977adf","parserVersion":"test_version"}
```

Name: Acastoides spp.

Canonical: Acastoides

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acastoides spp.","normalized":"Acastoides","canonical":{"stemmed":"Acastoides","simple":"Acastoides","full":"Acastoides"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acastoides","approximationMarker":"spp."}},"words":[{"verbatim":"Acastoides","normalized":"Acastoides","wordType":"GENUS","start":0,"end":10},{"verbatim":"spp.","normalized":"spp.","wordType":"APPROXIMATION_MARKER","start":11,"end":15}],"id":"9853f0a4-6324-5a7d-8108-e910578e612b","parserVersion":"test_version"}
```

### Ignoring serovar/serotype

Name: Aggregatibacter actinomycetemcomitans serotype d str. SA508

Canonical: Aggregatibacter actinomycetemcomitans

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aggregatibacter actinomycetemcomitans serotype d str. SA508","normalized":"Aggregatibacter actinomycetemcomitans","canonical":{"stemmed":"Aggregatibacter actinomycetemcomitans","simple":"Aggregatibacter actinomycetemcomitans","full":"Aggregatibacter actinomycetemcomitans"},"cardinality":2,"bacteria":"yes","tail":" serotype d str. SA508","details":{"species":{"genus":"Aggregatibacter","species":"actinomycetemcomitans"}},"words":[{"verbatim":"Aggregatibacter","normalized":"Aggregatibacter","wordType":"GENUS","start":0,"end":15},{"verbatim":"actinomycetemcomitans","normalized":"actinomycetemcomitans","wordType":"SPECIES","start":16,"end":37}],"id":"6f5d556a-6225-5412-8aa6-bebca2d9bfd5","parserVersion":"test_version"}
```

Name: Bacterium sp. (serotype) aboney Dräger 1951

Canonical: Bacterium

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Bacterium sp. (serotype) aboney Dräger 1951","normalized":"Bacterium","canonical":{"stemmed":"Bacterium","simple":"Bacterium","full":"Bacterium"},"cardinality":0,"bacteria":"yes","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Bacterium","approximationMarker":"sp.","ignored":" (serotype) aboney Dräger 1951"}},"words":[{"verbatim":"Bacterium","normalized":"Bacterium","wordType":"GENUS","start":0,"end":9},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":10,"end":13}],"id":"abe2f30e-d76a-5bdd-be47-a01c6572561a","parserVersion":"test_version"}
```

Name: Streptococcus pyogenes (serotype M18)

Canonical: Streptococcus pyogenes

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Streptococcus pyogenes (serotype M18)","normalized":"Streptococcus pyogenes","canonical":{"stemmed":"Streptococcus pyogen","simple":"Streptococcus pyogenes","full":"Streptococcus pyogenes"},"cardinality":2,"bacteria":"yes","tail":" (serotype M18)","details":{"species":{"genus":"Streptococcus","species":"pyogenes"}},"words":[{"verbatim":"Streptococcus","normalized":"Streptococcus","wordType":"GENUS","start":0,"end":13},{"verbatim":"pyogenes","normalized":"pyogenes","wordType":"SPECIES","start":14,"end":22}],"id":"cd677118-8336-56de-bfa6-fd849c6f7679","parserVersion":"test_version"}
```

Name: Actinobacillus pleuropneumoniae serovar 2 strain S1536

Canonical: Actinobacillus pleuropneumoniae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Actinobacillus pleuropneumoniae serovar 2 strain S1536","normalized":"Actinobacillus pleuropneumoniae","canonical":{"stemmed":"Actinobacillus pleuropneumoni","simple":"Actinobacillus pleuropneumoniae","full":"Actinobacillus pleuropneumoniae"},"cardinality":2,"bacteria":"yes","tail":" serovar 2 strain S1536","details":{"species":{"genus":"Actinobacillus","species":"pleuropneumoniae"}},"words":[{"verbatim":"Actinobacillus","normalized":"Actinobacillus","wordType":"GENUS","start":0,"end":14},{"verbatim":"pleuropneumoniae","normalized":"pleuropneumoniae","wordType":"SPECIES","start":15,"end":31}],"id":"fc0e4082-e830-5082-959c-02b69ea08f82","parserVersion":"test_version"}
```

Name: Leptospira interrogans serovar Fugis

Canonical: Leptospira interrogans

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Leptospira interrogans serovar Fugis","normalized":"Leptospira interrogans","canonical":{"stemmed":"Leptospira interrogans","simple":"Leptospira interrogans","full":"Leptospira interrogans"},"cardinality":2,"bacteria":"yes","tail":" serovar Fugis","details":{"species":{"genus":"Leptospira","species":"interrogans"}},"words":[{"verbatim":"Leptospira","normalized":"Leptospira","wordType":"GENUS","start":0,"end":10},{"verbatim":"interrogans","normalized":"interrogans","wordType":"SPECIES","start":11,"end":22}],"id":"026a23f1-dea7-5c57-8958-1efbe712a363","parserVersion":"test_version"}
```

### Ignoring sensu sec

Name: Senecio legionensis sensu Samp., non Lange

Canonical: Senecio legionensis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Senecio legionensis sensu Samp., non Lange","normalized":"Senecio legionensis","canonical":{"stemmed":"Senecio legionens","simple":"Senecio legionensis","full":"Senecio legionensis"},"cardinality":2,"tail":" sensu Samp., non Lange","details":{"species":{"genus":"Senecio","species":"legionensis"}},"words":[{"verbatim":"Senecio","normalized":"Senecio","wordType":"GENUS","start":0,"end":7},{"verbatim":"legionensis","normalized":"legionensis","wordType":"SPECIES","start":8,"end":19}],"id":"948d73b7-499b-5060-ace4-dd061f2f4373","parserVersion":"test_version"}
```

Name: Pseudomonas methanica (Söhngen 1906) sensu. Dworkin and Foster 1956

Canonical: Pseudomonas methanica

Authorship: (Söhngen 1906)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pseudomonas methanica (Söhngen 1906) sensu. Dworkin and Foster 1956","normalized":"Pseudomonas methanica (Söhngen 1906)","canonical":{"stemmed":"Pseudomonas methanic","simple":"Pseudomonas methanica","full":"Pseudomonas methanica"},"cardinality":2,"authorship":{"verbatim":"(Söhngen 1906)","normalized":"(Söhngen 1906)","year":"1906","authors":["Söhngen"],"originalAuth":{"authors":["Söhngen"],"year":{"year":"1906"}}},"bacteria":"yes","tail":" sensu. Dworkin and Foster 1956","details":{"species":{"genus":"Pseudomonas","species":"methanica","authorship":{"verbatim":"(Söhngen 1906)","normalized":"(Söhngen 1906)","year":"1906","authors":["Söhngen"],"originalAuth":{"authors":["Söhngen"],"year":{"year":"1906"}}}}},"words":[{"verbatim":"Pseudomonas","normalized":"Pseudomonas","wordType":"GENUS","start":0,"end":11},{"verbatim":"methanica","normalized":"methanica","wordType":"SPECIES","start":12,"end":21},{"verbatim":"Söhngen","normalized":"Söhngen","wordType":"AUTHOR_WORD","start":23,"end":30},{"verbatim":"1906","normalized":"1906","wordType":"YEAR","start":31,"end":35}],"id":"f4261966-4f80-52c1-a3ff-8eaece507964","parserVersion":"test_version"}
```

Name: Abarema scutifera sensu auct., non (Blanco)Kosterm.

Canonical: Abarema scutifera

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema scutifera sensu auct., non (Blanco)Kosterm.","normalized":"Abarema scutifera","canonical":{"stemmed":"Abarema scutifer","simple":"Abarema scutifera","full":"Abarema scutifera"},"cardinality":2,"tail":" sensu auct., non (Blanco)Kosterm.","details":{"species":{"genus":"Abarema","species":"scutifera"}},"words":[{"verbatim":"Abarema","normalized":"Abarema","wordType":"GENUS","start":0,"end":7},{"verbatim":"scutifera","normalized":"scutifera","wordType":"SPECIES","start":8,"end":17}],"id":"59f4b32d-3f8c-569f-bc81-3fe49d708c88","parserVersion":"test_version"}
```

Name: Puya acris Auct.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris Auct.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" Auct.","details":{"species":{"genus":"Puya","species":"acris"}},"words":[{"verbatim":"Puya","normalized":"Puya","wordType":"GENUS","start":0,"end":4},{"verbatim":"acris","normalized":"acris","wordType":"SPECIES","start":5,"end":10}],"id":"926ec12b-a597-5842-92f2-4b0ae4989df1","parserVersion":"test_version"}
```

Name: Puya acris Auct non L.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris Auct non L.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" Auct non L.","details":{"species":{"genus":"Puya","species":"acris"}},"words":[{"verbatim":"Puya","normalized":"Puya","wordType":"GENUS","start":0,"end":4},{"verbatim":"acris","normalized":"acris","wordType":"SPECIES","start":5,"end":10}],"id":"6c11df68-9e9d-5e97-b0f0-3609e4f18121","parserVersion":"test_version"}
```

Name: Galium tricorne Stokes, pro parte

Canonical: Galium tricorne

Authorship: Stokes

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Galium tricorne Stokes, pro parte","normalized":"Galium tricorne Stokes","canonical":{"stemmed":"Galium tricorn","simple":"Galium tricorne","full":"Galium tricorne"},"cardinality":2,"authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}},"tail":", pro parte","details":{"species":{"genus":"Galium","species":"tricorne","authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}}}},"words":[{"verbatim":"Galium","normalized":"Galium","wordType":"GENUS","start":0,"end":6},{"verbatim":"tricorne","normalized":"tricorne","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Stokes","normalized":"Stokes","wordType":"AUTHOR_WORD","start":16,"end":22}],"id":"c4d3da85-86b7-5ca9-925b-6e09ffad3a30","parserVersion":"test_version"}
```

Name: Galium tricorne Stokes,pro parte

Canonical: Galium tricorne

Authorship: Stokes

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Galium tricorne Stokes,pro parte","normalized":"Galium tricorne Stokes","canonical":{"stemmed":"Galium tricorn","simple":"Galium tricorne","full":"Galium tricorne"},"cardinality":2,"authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}},"tail":",pro parte","details":{"species":{"genus":"Galium","species":"tricorne","authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}}}},"words":[{"verbatim":"Galium","normalized":"Galium","wordType":"GENUS","start":0,"end":6},{"verbatim":"tricorne","normalized":"tricorne","wordType":"SPECIES","start":7,"end":15},{"verbatim":"Stokes","normalized":"Stokes","wordType":"AUTHOR_WORD","start":16,"end":22}],"id":"7166cbd9-2b0f-5537-9ac9-98157b60a395","parserVersion":"test_version"}
```

Name: Senecio jacquinianus sec. Rchb.

Canonical: Senecio jacquinianus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Senecio jacquinianus sec. Rchb.","normalized":"Senecio jacquinianus","canonical":{"stemmed":"Senecio iacquinian","simple":"Senecio jacquinianus","full":"Senecio jacquinianus"},"cardinality":2,"tail":" sec. Rchb.","details":{"species":{"genus":"Senecio","species":"jacquinianus"}},"words":[{"verbatim":"Senecio","normalized":"Senecio","wordType":"GENUS","start":0,"end":7},{"verbatim":"jacquinianus","normalized":"jacquinianus","wordType":"SPECIES","start":8,"end":20}],"id":"e8ad283f-afa8-5fd2-ae8f-bbedf2fb0bb7","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum s.l. (Schultes) Boiss.

Canonical: Acantholimon ulicinum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acantholimon ulicinum s.l. (Schultes) Boiss.","normalized":"Acantholimon ulicinum","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"tail":" s.l. (Schultes) Boiss.","details":{"species":{"genus":"Acantholimon","species":"ulicinum"}},"words":[{"verbatim":"Acantholimon","normalized":"Acantholimon","wordType":"GENUS","start":0,"end":12},{"verbatim":"ulicinum","normalized":"ulicinum","wordType":"SPECIES","start":13,"end":21}],"id":"cf4b7aa4-b78f-5b79-86c3-9416de24c918","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum s. l. (Schultes) Boiss.

Canonical: Acantholimon ulicinum

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acantholimon ulicinum s. l. (Schultes) Boiss.","normalized":"Acantholimon ulicinum","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"tail":" s. l. (Schultes) Boiss.","details":{"species":{"genus":"Acantholimon","species":"ulicinum"}},"words":[{"verbatim":"Acantholimon","normalized":"Acantholimon","wordType":"GENUS","start":0,"end":12},{"verbatim":"ulicinum","normalized":"ulicinum","wordType":"SPECIES","start":13,"end":21}],"id":"3a0b0412-f076-5714-8537-62761718ca7c","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum S. L. Schultes

Canonical: Acantholimon ulicinum

Authorship: S. L. Schultes

```json
{"parsed":true,"quality":1,"verbatim":"Acantholimon ulicinum S. L. Schultes","normalized":"Acantholimon ulicinum S. L. Schultes","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"authorship":{"verbatim":"S. L. Schultes","normalized":"S. L. Schultes","authors":["S. L. Schultes"],"originalAuth":{"authors":["S. L. Schultes"]}},"details":{"species":{"genus":"Acantholimon","species":"ulicinum","authorship":{"verbatim":"S. L. Schultes","normalized":"S. L. Schultes","authors":["S. L. Schultes"],"originalAuth":{"authors":["S. L. Schultes"]}}}},"words":[{"verbatim":"Acantholimon","normalized":"Acantholimon","wordType":"GENUS","start":0,"end":12},{"verbatim":"ulicinum","normalized":"ulicinum","wordType":"SPECIES","start":13,"end":21},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Schultes","normalized":"Schultes","wordType":"AUTHOR_WORD","start":28,"end":36}],"id":"702f97e0-792b-5ed4-b2d5-d813544c4139","parserVersion":"test_version"}
```

Name: Amitostigma formosana (S.S.Ying) S.S.Ying

Canonical: Amitostigma formosana

Authorship: (S. S. Ying) S. S. Ying

```json
{"parsed":true,"quality":1,"verbatim":"Amitostigma formosana (S.S.Ying) S.S.Ying","normalized":"Amitostigma formosana (S. S. Ying) S. S. Ying","canonical":{"stemmed":"Amitostigma formosan","simple":"Amitostigma formosana","full":"Amitostigma formosana"},"cardinality":2,"authorship":{"verbatim":"(S.S.Ying) S.S.Ying","normalized":"(S. S. Ying) S. S. Ying","authors":["S. S. Ying"],"originalAuth":{"authors":["S. S. Ying"]},"combinationAuth":{"authors":["S. S. Ying"]}},"details":{"species":{"genus":"Amitostigma","species":"formosana","authorship":{"verbatim":"(S.S.Ying) S.S.Ying","normalized":"(S. S. Ying) S. S. Ying","authors":["S. S. Ying"],"originalAuth":{"authors":["S. S. Ying"]},"combinationAuth":{"authors":["S. S. Ying"]}}}},"words":[{"verbatim":"Amitostigma","normalized":"Amitostigma","wordType":"GENUS","start":0,"end":11},{"verbatim":"formosana","normalized":"formosana","wordType":"SPECIES","start":12,"end":21},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Ying","normalized":"Ying","wordType":"AUTHOR_WORD","start":27,"end":31},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"S.","normalized":"S.","wordType":"AUTHOR_WORD","start":35,"end":37},{"verbatim":"Ying","normalized":"Ying","wordType":"AUTHOR_WORD","start":37,"end":41}],"id":"fcd831ea-57b6-5151-81e4-86e1c42f4695","parserVersion":"test_version"}
```

Name: Amaurorhinus bewichianus (Wollaston,1860) (s.str.)

Canonical: Amaurorhinus bewichianus

Authorship: (Wollaston 1860)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Amaurorhinus bewichianus (Wollaston,1860) (s.str.)","normalized":"Amaurorhinus bewichianus (Wollaston 1860)","canonical":{"stemmed":"Amaurorhinus bewichian","simple":"Amaurorhinus bewichianus","full":"Amaurorhinus bewichianus"},"cardinality":2,"authorship":{"verbatim":"(Wollaston,1860)","normalized":"(Wollaston 1860)","year":"1860","authors":["Wollaston"],"originalAuth":{"authors":["Wollaston"],"year":{"year":"1860"}}},"tail":" (s.str.)","details":{"species":{"genus":"Amaurorhinus","species":"bewichianus","authorship":{"verbatim":"(Wollaston,1860)","normalized":"(Wollaston 1860)","year":"1860","authors":["Wollaston"],"originalAuth":{"authors":["Wollaston"],"year":{"year":"1860"}}}}},"words":[{"verbatim":"Amaurorhinus","normalized":"Amaurorhinus","wordType":"GENUS","start":0,"end":12},{"verbatim":"bewichianus","normalized":"bewichianus","wordType":"SPECIES","start":13,"end":24},{"verbatim":"Wollaston","normalized":"Wollaston","wordType":"AUTHOR_WORD","start":26,"end":35},{"verbatim":"1860","normalized":"1860","wordType":"YEAR","start":36,"end":40}],"id":"b76e9160-d301-5696-bb87-499328996a7d","parserVersion":"test_version"}
```

Name: Ammodramus caudacutus (s.s.) diversus

Canonical: Ammodramus caudacutus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ammodramus caudacutus (s.s.) diversus","normalized":"Ammodramus caudacutus","canonical":{"stemmed":"Ammodramus caudacut","simple":"Ammodramus caudacutus","full":"Ammodramus caudacutus"},"cardinality":2,"tail":" (s.s.) diversus","details":{"species":{"genus":"Ammodramus","species":"caudacutus"}},"words":[{"verbatim":"Ammodramus","normalized":"Ammodramus","wordType":"GENUS","start":0,"end":10},{"verbatim":"caudacutus","normalized":"caudacutus","wordType":"SPECIES","start":11,"end":21}],"id":"2fb79b29-1579-5604-97bd-530c90c245cd","parserVersion":"test_version"}
```

Name: Arenaria serpyllifolia L. s.str.

Canonical: Arenaria serpyllifolia

Authorship: L.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arenaria serpyllifolia L. s.str.","normalized":"Arenaria serpyllifolia L.","canonical":{"stemmed":"Arenaria serpyllifol","simple":"Arenaria serpyllifolia","full":"Arenaria serpyllifolia"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" s.str.","details":{"species":{"genus":"Arenaria","species":"serpyllifolia","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"words":[{"verbatim":"Arenaria","normalized":"Arenaria","wordType":"GENUS","start":0,"end":8},{"verbatim":"serpyllifolia","normalized":"serpyllifolia","wordType":"SPECIES","start":9,"end":22},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":23,"end":25}],"id":"8a350298-0dfc-5ad0-9a10-60902587f335","parserVersion":"test_version"}
```

Name: Asplenium trichomanes L. s.lat. - Asplen trich

Canonical: Asplenium trichomanes

Authorship: L.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium trichomanes L. s.lat. - Asplen trich","normalized":"Asplenium trichomanes L.","canonical":{"stemmed":"Asplenium trichoman","simple":"Asplenium trichomanes","full":"Asplenium trichomanes"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" s.lat. - Asplen trich","details":{"species":{"genus":"Asplenium","species":"trichomanes","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"trichomanes","normalized":"trichomanes","wordType":"SPECIES","start":10,"end":21},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24}],"id":"1687d870-6bea-5573-80ef-4e55eca3199f","parserVersion":"test_version"}
```

Name: Asplenium anisophyllum Kunze, s.l.

Canonical: Asplenium anisophyllum

Authorship: Kunze

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium anisophyllum Kunze, s.l.","normalized":"Asplenium anisophyllum Kunze","canonical":{"stemmed":"Asplenium anisophyll","simple":"Asplenium anisophyllum","full":"Asplenium anisophyllum"},"cardinality":2,"authorship":{"verbatim":"Kunze","normalized":"Kunze","authors":["Kunze"],"originalAuth":{"authors":["Kunze"]}},"tail":", s.l.","details":{"species":{"genus":"Asplenium","species":"anisophyllum","authorship":{"verbatim":"Kunze","normalized":"Kunze","authors":["Kunze"],"originalAuth":{"authors":["Kunze"]}}}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"anisophyllum","normalized":"anisophyllum","wordType":"SPECIES","start":10,"end":22},{"verbatim":"Kunze","normalized":"Kunze","wordType":"AUTHOR_WORD","start":23,"end":28}],"id":"a0d7a55a-ffad-5243-905e-048177b440df","parserVersion":"test_version"}
```

Name: Abramis Cuvier 1816 sec. Dybowski 1862

Canonical: Abramis

Authorship: Cuvier 1816

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abramis Cuvier 1816 sec. Dybowski 1862","normalized":"Abramis Cuvier 1816","canonical":{"stemmed":"Abramis","simple":"Abramis","full":"Abramis"},"cardinality":1,"authorship":{"verbatim":"Cuvier 1816","normalized":"Cuvier 1816","year":"1816","authors":["Cuvier"],"originalAuth":{"authors":["Cuvier"],"year":{"year":"1816"}}},"tail":" sec. Dybowski 1862","details":{"uninomial":{"uninomial":"Abramis","authorship":{"verbatim":"Cuvier 1816","normalized":"Cuvier 1816","year":"1816","authors":["Cuvier"],"originalAuth":{"authors":["Cuvier"],"year":{"year":"1816"}}}}},"words":[{"verbatim":"Abramis","normalized":"Abramis","wordType":"UNINOMIAL","start":0,"end":7},{"verbatim":"Cuvier","normalized":"Cuvier","wordType":"AUTHOR_WORD","start":8,"end":14},{"verbatim":"1816","normalized":"1816","wordType":"YEAR","start":15,"end":19}],"id":"1fddff95-f470-5c36-8bc5-4436fe727bda","parserVersion":"test_version"}
```

Name: Abramis brama subsp. bergi Grib & Vernidub 1935 sec Eschmeyer 2004

Canonical: Abramis brama subsp. bergi

Authorship: Grib & Vernidub 1935

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abramis brama subsp. bergi Grib \u0026 Vernidub 1935 sec Eschmeyer 2004","normalized":"Abramis brama subsp. bergi Grib \u0026 Vernidub 1935","canonical":{"stemmed":"Abramis bram berg","simple":"Abramis brama bergi","full":"Abramis brama subsp. bergi"},"cardinality":3,"authorship":{"verbatim":"Grib \u0026 Vernidub 1935","normalized":"Grib \u0026 Vernidub 1935","year":"1935","authors":["Grib","Vernidub"],"originalAuth":{"authors":["Grib","Vernidub"],"year":{"year":"1935"}}},"tail":" sec Eschmeyer 2004","details":{"infraspecies":{"genus":"Abramis","species":"brama","infraspecies":[{"value":"bergi","rank":"subsp.","authorship":{"verbatim":"Grib \u0026 Vernidub 1935","normalized":"Grib \u0026 Vernidub 1935","year":"1935","authors":["Grib","Vernidub"],"originalAuth":{"authors":["Grib","Vernidub"],"year":{"year":"1935"}}}}]}},"words":[{"verbatim":"Abramis","normalized":"Abramis","wordType":"GENUS","start":0,"end":7},{"verbatim":"brama","normalized":"brama","wordType":"SPECIES","start":8,"end":13},{"verbatim":"subsp.","normalized":"subsp.","wordType":"RANK","start":14,"end":20},{"verbatim":"bergi","normalized":"bergi","wordType":"INFRASPECIES","start":21,"end":26},{"verbatim":"Grib","normalized":"Grib","wordType":"AUTHOR_WORD","start":27,"end":31},{"verbatim":"Vernidub","normalized":"Vernidub","wordType":"AUTHOR_WORD","start":34,"end":42},{"verbatim":"1935","normalized":"1935","wordType":"YEAR","start":43,"end":47}],"id":"5ac5f7fd-0a42-5133-961e-df94a54fb75f","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., P. P.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., P. P.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", P. P.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"words":[{"verbatim":"Abarema","normalized":"Abarema","wordType":"GENUS","start":0,"end":7},{"verbatim":"clypearia","normalized":"clypearia","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Jack","normalized":"Jack","wordType":"AUTHOR_WORD","start":19,"end":23},{"verbatim":"Kosterm.","normalized":"Kosterm.","wordType":"AUTHOR_WORD","start":25,"end":33}],"id":"2e18b789-865b-55dc-831b-f1fdd6bf740d","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., p.p.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., p.p.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", p.p.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"words":[{"verbatim":"Abarema","normalized":"Abarema","wordType":"GENUS","start":0,"end":7},{"verbatim":"clypearia","normalized":"clypearia","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Jack","normalized":"Jack","wordType":"AUTHOR_WORD","start":19,"end":23},{"verbatim":"Kosterm.","normalized":"Kosterm.","wordType":"AUTHOR_WORD","start":25,"end":33}],"id":"bc9b0feb-8a33-5f35-97a9-8ee93220fff8","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., p. p.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., p. p.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", p. p.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"words":[{"verbatim":"Abarema","normalized":"Abarema","wordType":"GENUS","start":0,"end":7},{"verbatim":"clypearia","normalized":"clypearia","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Jack","normalized":"Jack","wordType":"AUTHOR_WORD","start":19,"end":23},{"verbatim":"Kosterm.","normalized":"Kosterm.","wordType":"AUTHOR_WORD","start":25,"end":33}],"id":"1fae34cb-12f4-5600-9589-672199934719","parserVersion":"test_version"}
```

Name: Indigofera phyllogramme var. aphylla R.Vig., p.p.B

Canonical: Indigofera phyllogramme var. aphylla

Authorship: R. Vig.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Indigofera phyllogramme var. aphylla R.Vig., p.p.B","normalized":"Indigofera phyllogramme var. aphylla R. Vig.","canonical":{"stemmed":"Indigofera phyllogramm aphyll","simple":"Indigofera phyllogramme aphylla","full":"Indigofera phyllogramme var. aphylla"},"cardinality":3,"authorship":{"verbatim":"R.Vig.","normalized":"R. Vig.","authors":["R. Vig."],"originalAuth":{"authors":["R. Vig."]}},"tail":", p.p.B","details":{"infraspecies":{"genus":"Indigofera","species":"phyllogramme","infraspecies":[{"value":"aphylla","rank":"var.","authorship":{"verbatim":"R.Vig.","normalized":"R. Vig.","authors":["R. Vig."],"originalAuth":{"authors":["R. Vig."]}}}]}},"words":[{"verbatim":"Indigofera","normalized":"Indigofera","wordType":"GENUS","start":0,"end":10},{"verbatim":"phyllogramme","normalized":"phyllogramme","wordType":"SPECIES","start":11,"end":23},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":24,"end":28},{"verbatim":"aphylla","normalized":"aphylla","wordType":"INFRASPECIES","start":29,"end":36},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"Vig.","normalized":"Vig.","wordType":"AUTHOR_WORD","start":39,"end":43}],"id":"04bb878e-4442-5b7c-86d7-a41f2f6aefd3","parserVersion":"test_version"}
```

### Ignore terminal annotations

Name: Abida secale margaridae I.M.Fake Ms

Canonical: Abida secale margaridae

Authorship: I. M. Fake

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abida secale margaridae I.M.Fake Ms","normalized":"Abida secale margaridae I. M. Fake","canonical":{"stemmed":"Abida secal margarid","simple":"Abida secale margaridae","full":"Abida secale margaridae"},"cardinality":3,"authorship":{"verbatim":"I.M.Fake","normalized":"I. M. Fake","authors":["I. M. Fake"],"originalAuth":{"authors":["I. M. Fake"]}},"tail":" Ms","details":{"infraspecies":{"genus":"Abida","species":"secale","infraspecies":[{"value":"margaridae","authorship":{"verbatim":"I.M.Fake","normalized":"I. M. Fake","authors":["I. M. Fake"],"originalAuth":{"authors":["I. M. Fake"]}}}]}},"words":[{"verbatim":"Abida","normalized":"Abida","wordType":"GENUS","start":0,"end":5},{"verbatim":"secale","normalized":"secale","wordType":"SPECIES","start":6,"end":12},{"verbatim":"margaridae","normalized":"margaridae","wordType":"INFRASPECIES","start":13,"end":23},{"verbatim":"I.","normalized":"I.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"Fake","normalized":"Fake","wordType":"AUTHOR_WORD","start":28,"end":32}],"id":"a1409474-7c90-54c9-9161-7b003c9dffcb","parserVersion":"test_version"}
```

Name: Abida secale margaridae I.M.Fake ms

Canonical: Abida secale margaridae

Authorship: I. M. Fake

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abida secale margaridae I.M.Fake ms","normalized":"Abida secale margaridae I. M. Fake","canonical":{"stemmed":"Abida secal margarid","simple":"Abida secale margaridae","full":"Abida secale margaridae"},"cardinality":3,"authorship":{"verbatim":"I.M.Fake","normalized":"I. M. Fake","authors":["I. M. Fake"],"originalAuth":{"authors":["I. M. Fake"]}},"tail":" ms","details":{"infraspecies":{"genus":"Abida","species":"secale","infraspecies":[{"value":"margaridae","authorship":{"verbatim":"I.M.Fake","normalized":"I. M. Fake","authors":["I. M. Fake"],"originalAuth":{"authors":["I. M. Fake"]}}}]}},"words":[{"verbatim":"Abida","normalized":"Abida","wordType":"GENUS","start":0,"end":5},{"verbatim":"secale","normalized":"secale","wordType":"SPECIES","start":6,"end":12},{"verbatim":"margaridae","normalized":"margaridae","wordType":"INFRASPECIES","start":13,"end":23},{"verbatim":"I.","normalized":"I.","wordType":"AUTHOR_WORD","start":24,"end":26},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":26,"end":28},{"verbatim":"Fake","normalized":"Fake","wordType":"AUTHOR_WORD","start":28,"end":32}],"id":"cfa8d6e1-3913-512b-8e4f-163419c662bc","parserVersion":"test_version"}
```

### Unparseable hort. annotations

Name: Asplenium mayi ht.May; Gard.

Canonical: Asplenium mayi

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium mayi ht.May; Gard.","normalized":"Asplenium mayi","canonical":{"stemmed":"Asplenium may","simple":"Asplenium mayi","full":"Asplenium mayi"},"cardinality":2,"tail":" ht.May; Gard.","details":{"species":{"genus":"Asplenium","species":"mayi"}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"mayi","normalized":"mayi","wordType":"SPECIES","start":10,"end":14}],"id":"74446da2-14ce-5951-95c6-054d29417131","parserVersion":"test_version"}
```

Name: Asplenium mayii ht.May; Gard.

Canonical: Asplenium mayii

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium mayii ht.May; Gard.","normalized":"Asplenium mayii","canonical":{"stemmed":"Asplenium may","simple":"Asplenium mayii","full":"Asplenium mayii"},"cardinality":2,"tail":" ht.May; Gard.","details":{"species":{"genus":"Asplenium","species":"mayii"}},"words":[{"verbatim":"Asplenium","normalized":"Asplenium","wordType":"GENUS","start":0,"end":9},{"verbatim":"mayii","normalized":"mayii","wordType":"SPECIES","start":10,"end":15}],"id":"00764ac3-b9eb-56bf-9856-6de62459646e","parserVersion":"test_version"}
```

Name: Davallia decora ht.Bull.; Gard.Chr.

Canonical: Davallia decora

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Davallia decora ht.Bull.; Gard.Chr.","normalized":"Davallia decora","canonical":{"stemmed":"Davallia decor","simple":"Davallia decora","full":"Davallia decora"},"cardinality":2,"tail":" ht.Bull.; Gard.Chr.","details":{"species":{"genus":"Davallia","species":"decora"}},"words":[{"verbatim":"Davallia","normalized":"Davallia","wordType":"GENUS","start":0,"end":8},{"verbatim":"decora","normalized":"decora","wordType":"SPECIES","start":9,"end":15}],"id":"2e6032e9-1a08-5149-8339-5361c84c4a2d","parserVersion":"test_version"}
```

Name: Gymnogramma alstoni ht.Birkenh.; Gard.

Canonical: Gymnogramma alstoni

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Gymnogramma alstoni ht.Birkenh.; Gard.","normalized":"Gymnogramma alstoni","canonical":{"stemmed":"Gymnogramma alston","simple":"Gymnogramma alstoni","full":"Gymnogramma alstoni"},"cardinality":2,"tail":" ht.Birkenh.; Gard.","details":{"species":{"genus":"Gymnogramma","species":"alstoni"}},"words":[{"verbatim":"Gymnogramma","normalized":"Gymnogramma","wordType":"GENUS","start":0,"end":11},{"verbatim":"alstoni","normalized":"alstoni","wordType":"SPECIES","start":12,"end":19}],"id":"77b0759a-2b8f-51ef-8a40-df9268c72cf1","parserVersion":"test_version"}
```

Name: Gymnogramma sprengeriana ht.Wiener Ill.

Canonical: Gymnogramma sprengeriana

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Gymnogramma sprengeriana ht.Wiener Ill.","normalized":"Gymnogramma sprengeriana","canonical":{"stemmed":"Gymnogramma sprengerian","simple":"Gymnogramma sprengeriana","full":"Gymnogramma sprengeriana"},"cardinality":2,"tail":" ht.Wiener Ill.","details":{"species":{"genus":"Gymnogramma","species":"sprengeriana"}},"words":[{"verbatim":"Gymnogramma","normalized":"Gymnogramma","wordType":"GENUS","start":0,"end":11},{"verbatim":"sprengeriana","normalized":"sprengeriana","wordType":"SPECIES","start":12,"end":24}],"id":"4e5517fa-4b2c-55f6-8471-76c26ed9983a","parserVersion":"test_version"}
```

### Removing nomenclatural annotations

Name: Amphiprora pseudoduplex (Osada & Kobayasi, 1990) comb. nov.

Canonical: Amphiprora pseudoduplex

Authorship: (Osada & Kobayasi 1990)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Amphiprora pseudoduplex (Osada \u0026 Kobayasi, 1990) comb. nov.","normalized":"Amphiprora pseudoduplex (Osada \u0026 Kobayasi 1990)","canonical":{"stemmed":"Amphiprora pseudoduplex","simple":"Amphiprora pseudoduplex","full":"Amphiprora pseudoduplex"},"cardinality":2,"authorship":{"verbatim":"(Osada \u0026 Kobayasi, 1990)","normalized":"(Osada \u0026 Kobayasi 1990)","year":"1990","authors":["Osada","Kobayasi"],"originalAuth":{"authors":["Osada","Kobayasi"],"year":{"year":"1990"}}},"tail":" comb. nov.","details":{"species":{"genus":"Amphiprora","species":"pseudoduplex","authorship":{"verbatim":"(Osada \u0026 Kobayasi, 1990)","normalized":"(Osada \u0026 Kobayasi 1990)","year":"1990","authors":["Osada","Kobayasi"],"originalAuth":{"authors":["Osada","Kobayasi"],"year":{"year":"1990"}}}}},"words":[{"verbatim":"Amphiprora","normalized":"Amphiprora","wordType":"GENUS","start":0,"end":10},{"verbatim":"pseudoduplex","normalized":"pseudoduplex","wordType":"SPECIES","start":11,"end":23},{"verbatim":"Osada","normalized":"Osada","wordType":"AUTHOR_WORD","start":25,"end":30},{"verbatim":"Kobayasi","normalized":"Kobayasi","wordType":"AUTHOR_WORD","start":33,"end":41},{"verbatim":"1990","normalized":"1990","wordType":"YEAR","start":43,"end":47}],"id":"06b58578-d00c-5c90-b77a-bc2325694b51","parserVersion":"test_version"}
```

Name: Methanosarcina barkeri str. fusaro

Canonical: Methanosarcina barkeri

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Methanosarcina barkeri str. fusaro","normalized":"Methanosarcina barkeri","canonical":{"stemmed":"Methanosarcina barker","simple":"Methanosarcina barkeri","full":"Methanosarcina barkeri"},"cardinality":2,"tail":" str. fusaro","details":{"species":{"genus":"Methanosarcina","species":"barkeri"}},"words":[{"verbatim":"Methanosarcina","normalized":"Methanosarcina","wordType":"GENUS","start":0,"end":14},{"verbatim":"barkeri","normalized":"barkeri","wordType":"SPECIES","start":15,"end":22}],"id":"b1d6747d-6aa3-5b7a-a8ed-7ca53c4b19ac","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Nyl.) R.C. Harris comb. nov.

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl.) R. C. Harris

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arthopyrenia hyalospora (Nyl.) R.C. Harris comb. nov.","normalized":"Arthopyrenia hyalospora (Nyl.) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl.) R.C. Harris","normalized":"(Nyl.) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["R. C. Harris"]}},"tail":" comb. nov.","details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl.) R.C. Harris","normalized":"(Nyl.) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["R. C. Harris"]}}}},"words":[{"verbatim":"Arthopyrenia","normalized":"Arthopyrenia","wordType":"GENUS","start":0,"end":12},{"verbatim":"hyalospora","normalized":"hyalospora","wordType":"SPECIES","start":13,"end":23},{"verbatim":"Nyl.","normalized":"Nyl.","wordType":"AUTHOR_WORD","start":25,"end":29},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":31,"end":33},{"verbatim":"C.","normalized":"C.","wordType":"AUTHOR_WORD","start":33,"end":35},{"verbatim":"Harris","normalized":"Harris","wordType":"AUTHOR_WORD","start":36,"end":42}],"id":"2dcef387-edc3-55a1-9cfc-ee95200bff08","parserVersion":"test_version"}
```

Name: Acanthophis lancasteri WELLS & WELLINGTON (nomen nudum)

Canonical: Acanthophis lancasteri

Authorship: Wells & Wellington

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Author in upper case"}],"verbatim":"Acanthophis lancasteri WELLS \u0026 WELLINGTON (nomen nudum)","normalized":"Acanthophis lancasteri Wells \u0026 Wellington","canonical":{"stemmed":"Acanthophis lancaster","simple":"Acanthophis lancasteri","full":"Acanthophis lancasteri"},"cardinality":2,"authorship":{"verbatim":"WELLS \u0026 WELLINGTON","normalized":"Wells \u0026 Wellington","authors":["Wells","Wellington"],"originalAuth":{"authors":["Wells","Wellington"]}},"tail":" (nomen nudum)","details":{"species":{"genus":"Acanthophis","species":"lancasteri","authorship":{"verbatim":"WELLS \u0026 WELLINGTON","normalized":"Wells \u0026 Wellington","authors":["Wells","Wellington"],"originalAuth":{"authors":["Wells","Wellington"]}}}},"words":[{"verbatim":"Acanthophis","normalized":"Acanthophis","wordType":"GENUS","start":0,"end":11},{"verbatim":"lancasteri","normalized":"lancasteri","wordType":"SPECIES","start":12,"end":22},{"verbatim":"WELLS","normalized":"Wells","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"WELLINGTON","normalized":"Wellington","wordType":"AUTHOR_WORD","start":31,"end":41}],"id":"aa527c3b-972e-56e9-9b8b-0c61c497422d","parserVersion":"test_version"}
```

Name: Acontias lineatus WAGLER 1830: 196 (nomen nudum)

Canonical: Acontias lineatus

Authorship: Wagler 1830

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Acontias lineatus WAGLER 1830: 196 (nomen nudum)","normalized":"Acontias lineatus Wagler 1830","canonical":{"stemmed":"Acontias lineat","simple":"Acontias lineatus","full":"Acontias lineatus"},"cardinality":2,"authorship":{"verbatim":"WAGLER 1830: 196","normalized":"Wagler 1830","year":"1830","authors":["Wagler"],"originalAuth":{"authors":["Wagler"],"year":{"year":"1830"}}},"tail":" (nomen nudum)","details":{"species":{"genus":"Acontias","species":"lineatus","authorship":{"verbatim":"WAGLER 1830: 196","normalized":"Wagler 1830","year":"1830","authors":["Wagler"],"originalAuth":{"authors":["Wagler"],"year":{"year":"1830"}}}}},"words":[{"verbatim":"Acontias","normalized":"Acontias","wordType":"GENUS","start":0,"end":8},{"verbatim":"lineatus","normalized":"lineatus","wordType":"SPECIES","start":9,"end":17},{"verbatim":"WAGLER","normalized":"Wagler","wordType":"AUTHOR_WORD","start":18,"end":24},{"verbatim":"1830","normalized":"1830","wordType":"YEAR","start":25,"end":29}],"id":"16afe3dd-7724-5dc0-817c-f6d138d27174","parserVersion":"test_version"}
```

Name: Akeratidae Nomen Nudum

Canonical: Akeratidae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Akeratidae Nomen Nudum","normalized":"Akeratidae","canonical":{"stemmed":"Akeratidae","simple":"Akeratidae","full":"Akeratidae"},"cardinality":1,"tail":" Nomen Nudum","details":{"uninomial":{"uninomial":"Akeratidae"}},"words":[{"verbatim":"Akeratidae","normalized":"Akeratidae","wordType":"UNINOMIAL","start":0,"end":10}],"id":"6bd60fba-9b78-5e4e-b904-dda976085fc7","parserVersion":"test_version"}
```

Name: Aster exilis Ell., nomen dubium

Canonical: Aster exilis

Authorship: Ell.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aster exilis Ell., nomen dubium","normalized":"Aster exilis Ell.","canonical":{"stemmed":"Aster exil","simple":"Aster exilis","full":"Aster exilis"},"cardinality":2,"authorship":{"verbatim":"Ell.","normalized":"Ell.","authors":["Ell."],"originalAuth":{"authors":["Ell."]}},"tail":", nomen dubium","details":{"species":{"genus":"Aster","species":"exilis","authorship":{"verbatim":"Ell.","normalized":"Ell.","authors":["Ell."],"originalAuth":{"authors":["Ell."]}}}},"words":[{"verbatim":"Aster","normalized":"Aster","wordType":"GENUS","start":0,"end":5},{"verbatim":"exilis","normalized":"exilis","wordType":"SPECIES","start":6,"end":12},{"verbatim":"Ell.","normalized":"Ell.","wordType":"AUTHOR_WORD","start":13,"end":17}],"id":"00884bdf-ca19-5c07-8e48-e1adef987844","parserVersion":"test_version"}
```

Name: Abutilon avicennae Gaertn., nom. illeg.

Canonical: Abutilon avicennae

Authorship: Gaertn.

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abutilon avicennae Gaertn., nom. illeg.","normalized":"Abutilon avicennae Gaertn.","canonical":{"stemmed":"Abutilon auicenn","simple":"Abutilon avicennae","full":"Abutilon avicennae"},"cardinality":2,"authorship":{"verbatim":"Gaertn.","normalized":"Gaertn.","authors":["Gaertn."],"originalAuth":{"authors":["Gaertn."]}},"tail":", nom. illeg.","details":{"species":{"genus":"Abutilon","species":"avicennae","authorship":{"verbatim":"Gaertn.","normalized":"Gaertn.","authors":["Gaertn."],"originalAuth":{"authors":["Gaertn."]}}}},"words":[{"verbatim":"Abutilon","normalized":"Abutilon","wordType":"GENUS","start":0,"end":8},{"verbatim":"avicennae","normalized":"avicennae","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Gaertn.","normalized":"Gaertn.","wordType":"AUTHOR_WORD","start":19,"end":26}],"id":"366d9605-0686-5072-b025-6c7b3695f086","parserVersion":"test_version"}
```

Name: Achillea bonarota nom. in herb.

Canonical: Achillea bonarota

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Achillea bonarota nom. in herb.","normalized":"Achillea bonarota","canonical":{"stemmed":"Achillea bonarot","simple":"Achillea bonarota","full":"Achillea bonarota"},"cardinality":2,"tail":" nom. in herb.","details":{"species":{"genus":"Achillea","species":"bonarota"}},"words":[{"verbatim":"Achillea","normalized":"Achillea","wordType":"GENUS","start":0,"end":8},{"verbatim":"bonarota","normalized":"bonarota","wordType":"SPECIES","start":9,"end":17}],"id":"cae8ac71-b3c4-52f7-94cb-31e639081e0d","parserVersion":"test_version"}
```

Name: Aconitum napellus var. formosum (Rchb.) W. D. J. Koch (nom. ambig.)

Canonical: Aconitum napellus var. formosum

Authorship: (Rchb.) W. D. J. Koch

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aconitum napellus var. formosum (Rchb.) W. D. J. Koch (nom. ambig.)","normalized":"Aconitum napellus var. formosum (Rchb.) W. D. J. Koch","canonical":{"stemmed":"Aconitum napell formos","simple":"Aconitum napellus formosum","full":"Aconitum napellus var. formosum"},"cardinality":3,"authorship":{"verbatim":"(Rchb.) W. D. J. Koch","normalized":"(Rchb.) W. D. J. Koch","authors":["Rchb.","W. D. J. Koch"],"originalAuth":{"authors":["Rchb."]},"combinationAuth":{"authors":["W. D. J. Koch"]}},"tail":" (nom. ambig.)","details":{"infraspecies":{"genus":"Aconitum","species":"napellus","infraspecies":[{"value":"formosum","rank":"var.","authorship":{"verbatim":"(Rchb.) W. D. J. Koch","normalized":"(Rchb.) W. D. J. Koch","authors":["Rchb.","W. D. J. Koch"],"originalAuth":{"authors":["Rchb."]},"combinationAuth":{"authors":["W. D. J. Koch"]}}}]}},"words":[{"verbatim":"Aconitum","normalized":"Aconitum","wordType":"GENUS","start":0,"end":8},{"verbatim":"napellus","normalized":"napellus","wordType":"SPECIES","start":9,"end":17},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":18,"end":22},{"verbatim":"formosum","normalized":"formosum","wordType":"INFRASPECIES","start":23,"end":31},{"verbatim":"Rchb.","normalized":"Rchb.","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":40,"end":42},{"verbatim":"D.","normalized":"D.","wordType":"AUTHOR_WORD","start":43,"end":45},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":46,"end":48},{"verbatim":"Koch","normalized":"Koch","wordType":"AUTHOR_WORD","start":49,"end":53}],"id":"9f79b2b3-cfd1-541a-9898-b60829134b11","parserVersion":"test_version"}
```

Name: Aesculus canadensis Hort. ex Lavallée

Canonical: Aesculus canadensis

Authorship: Hort. ex Lavallée

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required (ICZN only)"}],"verbatim":"Aesculus canadensis Hort. ex Lavallée","normalized":"Aesculus canadensis Hort. ex Lavallée","canonical":{"stemmed":"Aesculus canadens","simple":"Aesculus canadensis","full":"Aesculus canadensis"},"cardinality":2,"authorship":{"verbatim":"Hort. ex Lavallée","normalized":"Hort. ex Lavallée","authors":["Hort.","Lavallée"],"originalAuth":{"authors":["Hort."],"exAuthors":{"authors":["Lavallée"]}}},"details":{"species":{"genus":"Aesculus","species":"canadensis","authorship":{"verbatim":"Hort. ex Lavallée","normalized":"Hort. ex Lavallée","authors":["Hort.","Lavallée"],"originalAuth":{"authors":["Hort."],"exAuthors":{"authors":["Lavallée"]}}}}},"words":[{"verbatim":"Aesculus","normalized":"Aesculus","wordType":"GENUS","start":0,"end":8},{"verbatim":"canadensis","normalized":"canadensis","wordType":"SPECIES","start":9,"end":19},{"verbatim":"Hort.","normalized":"Hort.","wordType":"AUTHOR_WORD","start":20,"end":25},{"verbatim":"Lavallée","normalized":"Lavallée","wordType":"AUTHOR_WORD","start":29,"end":37}],"id":"a1c7935f-26c2-5388-a1e2-b5a9508d70ef","parserVersion":"test_version"}
```

Name: × Dialaeliopsis hort.

Canonical: × Dialaeliopsis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"× Dialaeliopsis hort.","normalized":"× Dialaeliopsis","canonical":{"stemmed":"Dialaeliopsis","simple":"Dialaeliopsis","full":"× Dialaeliopsis"},"cardinality":1,"hybrid":"NAMED_HYBRID","tail":" hort.","details":{"uninomial":{"uninomial":"Dialaeliopsis"}},"words":[{"verbatim":"×","normalized":"×","wordType":"HYBRID_CHAR","start":0,"end":1},{"verbatim":"Dialaeliopsis","normalized":"Dialaeliopsis","wordType":"UNINOMIAL","start":2,"end":15}],"id":"5e0197df-26c1-55bc-a5c0-64376c599fa5","parserVersion":"test_version"}
```

### Misc annotations

Name: Feldmannia species

Canonical: Feldmannia

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Feldmannia species","normalized":"Feldmannia","canonical":{"stemmed":"Feldmannia","simple":"Feldmannia","full":"Feldmannia"},"cardinality":1,"tail":" species","details":{"uninomial":{"uninomial":"Feldmannia"}},"words":[{"verbatim":"Feldmannia","normalized":"Feldmannia","wordType":"UNINOMIAL","start":0,"end":10}],"id":"55474a4d-2fc1-5417-8fac-06485167c33e","parserVersion":"test_version"}
```

Name: Periglypta G. Paulay, MS

Canonical: Periglypta

Authorship: G. Paulay

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Periglypta G. Paulay, MS","normalized":"Periglypta G. Paulay","canonical":{"stemmed":"Periglypta","simple":"Periglypta","full":"Periglypta"},"cardinality":1,"authorship":{"verbatim":"G. Paulay","normalized":"G. Paulay","authors":["G. Paulay"],"originalAuth":{"authors":["G. Paulay"]}},"tail":", MS","details":{"uninomial":{"uninomial":"Periglypta","authorship":{"verbatim":"G. Paulay","normalized":"G. Paulay","authors":["G. Paulay"],"originalAuth":{"authors":["G. Paulay"]}}}},"words":[{"verbatim":"Periglypta","normalized":"Periglypta","wordType":"UNINOMIAL","start":0,"end":10},{"verbatim":"G.","normalized":"G.","wordType":"AUTHOR_WORD","start":11,"end":13},{"verbatim":"Paulay","normalized":"Paulay","wordType":"AUTHOR_WORD","start":14,"end":20}],"id":"6da4ccdf-99c9-5cef-ae1c-a2d332a9c476","parserVersion":"test_version"}
```

Name: Teredo not found

Canonical: Teredo

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Teredo not found","normalized":"Teredo","canonical":{"stemmed":"Teredo","simple":"Teredo","full":"Teredo"},"cardinality":1,"tail":" not found","details":{"uninomial":{"uninomial":"Teredo"}},"words":[{"verbatim":"Teredo","normalized":"Teredo","wordType":"UNINOMIAL","start":0,"end":6}],"id":"81d633f5-1f21-53ca-bbd0-92e436f440d1","parserVersion":"test_version"}
```
Name: Velutina haliotoides (Linnaeus, 1758), sensu Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758), sensu Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":", sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Velutina","normalized":"Velutina","wordType":"GENUS","start":0,"end":8},{"verbatim":"haliotoides","normalized":"haliotoides","wordType":"SPECIES","start":9,"end":20},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":32,"end":36}],"id":"5efd63de-f4ec-55f1-bd5b-494988e58f9b","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 group

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 group","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" group","details":{"infraspecies":{"genus":"Acarospora","species":"cratericola","infraspecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"words":[{"verbatim":"Acarospora","normalized":"Acarospora","wordType":"GENUS","start":0,"end":10},{"verbatim":"cratericola","normalized":"cratericola","wordType":"SPECIES","start":11,"end":22},{"verbatim":"cratericola","normalized":"cratericola","wordType":"INFRASPECIES","start":23,"end":34},{"verbatim":"Shenk","normalized":"Shenk","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"1974","normalized":"1974","wordType":"YEAR","start":41,"end":45}],"id":"0f466e31-7e23-5320-ac7e-4c1026bc8af6","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 species group

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 species group","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" species group","details":{"infraspecies":{"genus":"Acarospora","species":"cratericola","infraspecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"words":[{"verbatim":"Acarospora","normalized":"Acarospora","wordType":"GENUS","start":0,"end":10},{"verbatim":"cratericola","normalized":"cratericola","wordType":"SPECIES","start":11,"end":22},{"verbatim":"cratericola","normalized":"cratericola","wordType":"INFRASPECIES","start":23,"end":34},{"verbatim":"Shenk","normalized":"Shenk","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"1974","normalized":"1974","wordType":"YEAR","start":41,"end":45}],"id":"a7684260-ed99-5d55-9a35-fd97b67e8933","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 species complex

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 species complex","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" species complex","details":{"infraspecies":{"genus":"Acarospora","species":"cratericola","infraspecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"words":[{"verbatim":"Acarospora","normalized":"Acarospora","wordType":"GENUS","start":0,"end":10},{"verbatim":"cratericola","normalized":"cratericola","wordType":"SPECIES","start":11,"end":22},{"verbatim":"cratericola","normalized":"cratericola","wordType":"INFRASPECIES","start":23,"end":34},{"verbatim":"Shenk","normalized":"Shenk","wordType":"AUTHOR_WORD","start":35,"end":40},{"verbatim":"1974","normalized":"1974","wordType":"YEAR","start":41,"end":45}],"id":"d227da04-7c89-50f7-8cf1-de09bc5aa903","parserVersion":"test_version"}
```

Name: Parus caeruleus species complex

Canonical: Parus caeruleus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Parus caeruleus species complex","normalized":"Parus caeruleus","canonical":{"stemmed":"Parus caerule","simple":"Parus caeruleus","full":"Parus caeruleus"},"cardinality":2,"tail":" species complex","details":{"species":{"genus":"Parus","species":"caeruleus"}},"words":[{"verbatim":"Parus","normalized":"Parus","wordType":"GENUS","start":0,"end":5},{"verbatim":"caeruleus","normalized":"caeruleus","wordType":"SPECIES","start":6,"end":15}],"id":"f3752c09-242f-501c-8c8c-0feaf86c4693","parserVersion":"test_version"}
```

Name: Crenarchaeote enrichment culture clone OREC-B1022

Canonical: Crenarchaeote

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Crenarchaeote enrichment culture clone OREC-B1022","normalized":"Crenarchaeote","canonical":{"stemmed":"Crenarchaeote","simple":"Crenarchaeote","full":"Crenarchaeote"},"cardinality":1,"tail":" enrichment culture clone OREC-B1022","details":{"uninomial":{"uninomial":"Crenarchaeote"}},"words":[{"verbatim":"Crenarchaeote","normalized":"Crenarchaeote","wordType":"UNINOMIAL","start":0,"end":13}],"id":"f16c9aa3-f749-5025-b9cb-2dcfc6d7629b","parserVersion":"test_version"}
```

Name: Diodora dorsata  CF

Canonical: Diodora dorsata

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Diodora dorsata  CF","normalized":"Diodora dorsata","canonical":{"stemmed":"Diodora dorsat","simple":"Diodora dorsata","full":"Diodora dorsata"},"cardinality":2,"tail":"  CF","details":{"species":{"genus":"Diodora","species":"dorsata"}},"words":[{"verbatim":"Diodora","normalized":"Diodora","wordType":"GENUS","start":0,"end":7},{"verbatim":"dorsata","normalized":"dorsata","wordType":"SPECIES","start":8,"end":15}],"id":"d3991dd5-f6c2-54aa-94e1-419fb560e703","parserVersion":"test_version"}
```

Name: Dasysyrphus intrudens complex sp. BBDCQ003-10

Canonical: Dasysyrphus intrudens

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Dasysyrphus intrudens complex sp. BBDCQ003-10","normalized":"Dasysyrphus intrudens","canonical":{"stemmed":"Dasysyrphus intrudens","simple":"Dasysyrphus intrudens","full":"Dasysyrphus intrudens"},"cardinality":2,"tail":" complex sp. BBDCQ003-10","details":{"species":{"genus":"Dasysyrphus","species":"intrudens"}},"words":[{"verbatim":"Dasysyrphus","normalized":"Dasysyrphus","wordType":"GENUS","start":0,"end":11},{"verbatim":"intrudens","normalized":"intrudens","wordType":"SPECIES","start":12,"end":21}],"id":"5c436c20-40bc-5969-a788-72e3b87451b2","parserVersion":"test_version"}
```

### Horticultural annotation

Name: Lachenalia tricolor var. nelsonii (ht.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lachenalia tricolor var. nelsonii (ht.) Baker","normalized":"Lachenalia tricolor var. nelsonii","canonical":{"stemmed":"Lachenalia tricolor nelson","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"tail":" (ht.) Baker","details":{"infraspecies":{"genus":"Lachenalia","species":"tricolor","infraspecies":[{"value":"nelsonii","rank":"var."}]}},"words":[{"verbatim":"Lachenalia","normalized":"Lachenalia","wordType":"GENUS","start":0,"end":10},{"verbatim":"tricolor","normalized":"tricolor","wordType":"SPECIES","start":11,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"nelsonii","normalized":"nelsonii","wordType":"INFRASPECIES","start":25,"end":33}],"id":"0f7ce439-6b8d-53db-9ea3-82628f25b9bd","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (hort.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lachenalia tricolor var. nelsonii (hort.) Baker","normalized":"Lachenalia tricolor var. nelsonii","canonical":{"stemmed":"Lachenalia tricolor nelson","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"tail":" (hort.) Baker","details":{"infraspecies":{"genus":"Lachenalia","species":"tricolor","infraspecies":[{"value":"nelsonii","rank":"var."}]}},"words":[{"verbatim":"Lachenalia","normalized":"Lachenalia","wordType":"GENUS","start":0,"end":10},{"verbatim":"tricolor","normalized":"tricolor","wordType":"SPECIES","start":11,"end":19},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":20,"end":24},{"verbatim":"nelsonii","normalized":"nelsonii","wordType":"INFRASPECIES","start":25,"end":33}],"id":"cc118b05-14ff-5a42-8780-802f60eba565","parserVersion":"test_version"}
```

Name: Puya acris ht.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris ht.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" ht.","details":{"species":{"genus":"Puya","species":"acris"}},"words":[{"verbatim":"Puya","normalized":"Puya","wordType":"GENUS","start":0,"end":4},{"verbatim":"acris","normalized":"acris","wordType":"SPECIES","start":5,"end":10}],"id":"83c98b8e-f373-57df-92bf-5a39a56d9909","parserVersion":"test_version"}
```

Name: Puya acris hort.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris hort.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" hort.","details":{"species":{"genus":"Puya","species":"acris"}},"words":[{"verbatim":"Puya","normalized":"Puya","wordType":"GENUS","start":0,"end":4},{"verbatim":"acris","normalized":"acris","wordType":"SPECIES","start":5,"end":10}],"id":"78228a5e-dcd3-58f9-bf21-b452c378f6ee","parserVersion":"test_version"}
```

### Names with "mihi"

Name: Characium obovatum mihi. var. longipes mihi

Canonical: Characium obovatum var. longipes

Authorship:

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ignored annotation `mihi`"}],"verbatim":"Characium obovatum mihi. var. longipes mihi","normalized":"Characium obovatum var. longipes","canonical":{"stemmed":"Characium obouat longip","simple":"Characium obovatum longipes","full":"Characium obovatum var. longipes"},"cardinality":3,"details":{"infraspecies":{"genus":"Characium","species":"obovatum","infraspecies":[{"value":"longipes","rank":"var."}]}},"words":[{"verbatim":"Characium","normalized":"Characium","wordType":"GENUS","start":0,"end":9},{"verbatim":"obovatum","normalized":"obovatum","wordType":"SPECIES","start":10,"end":18},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":25,"end":29},{"verbatim":"longipes","normalized":"longipes","wordType":"INFRASPECIES","start":30,"end":38}],"id":"39baca43-fcb1-5b13-8458-0729fb5f22dd","parserVersion":"test_version"}
```

Name: Regulus modestus mihi. Gould 1837

Canonical: Regulus modestus

Authorship: Gould 1837

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"Ignored annotation `mihi`"}],"verbatim":"Regulus modestus mihi. Gould 1837","normalized":"Regulus modestus Gould 1837","canonical":{"stemmed":"Regulus modest","simple":"Regulus modestus","full":"Regulus modestus"},"cardinality":2,"authorship":{"verbatim":"Gould 1837","normalized":"Gould 1837","year":"1837","authors":["Gould"],"originalAuth":{"authors":["Gould"],"year":{"year":"1837"}}},"details":{"species":{"genus":"Regulus","species":"modestus","authorship":{"verbatim":"Gould 1837","normalized":"Gould 1837","year":"1837","authors":["Gould"],"originalAuth":{"authors":["Gould"],"year":{"year":"1837"}}}}},"words":[{"verbatim":"Regulus","normalized":"Regulus","wordType":"GENUS","start":0,"end":7},{"verbatim":"modestus","normalized":"modestus","wordType":"SPECIES","start":8,"end":16},{"verbatim":"Gould","normalized":"Gould","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"1837","normalized":"1837","wordType":"YEAR","start":29,"end":33}],"id":"4cb15cc3-9327-552f-9afb-2af349a874a5","parserVersion":"test_version"}
```

### Exceptions with "mihi"

Name: Eucyclops serrulatus mihi Dussart, Graf & Husson, 1966

Canonical: Eucyclops serrulatus mihi

Authorship: Dussart, Graf & Husson 1966

```json
{"parsed":true,"quality":1,"verbatim":"Eucyclops serrulatus mihi Dussart, Graf \u0026 Husson, 1966","normalized":"Eucyclops serrulatus mihi Dussart, Graf \u0026 Husson 1966","canonical":{"stemmed":"Eucyclops serrulat mih","simple":"Eucyclops serrulatus mihi","full":"Eucyclops serrulatus mihi"},"cardinality":3,"authorship":{"verbatim":"Dussart, Graf \u0026 Husson, 1966","normalized":"Dussart, Graf \u0026 Husson 1966","year":"1966","authors":["Dussart","Graf","Husson"],"originalAuth":{"authors":["Dussart","Graf","Husson"],"year":{"year":"1966"}}},"details":{"infraspecies":{"genus":"Eucyclops","species":"serrulatus","infraspecies":[{"value":"kihi","authorship":{"verbatim":"Dussart, Graf \u0026 Husson, 1966","normalized":"Dussart, Graf \u0026 Husson 1966","year":"1966","authors":["Dussart","Graf","Husson"],"originalAuth":{"authors":["Dussart","Graf","Husson"],"year":{"year":"1966"}}}}]}},"words":[{"verbatim":"Eucyclops","normalized":"Eucyclops","wordType":"GENUS","start":0,"end":9},{"verbatim":"serrulatus","normalized":"serrulatus","wordType":"SPECIES","start":10,"end":20},{"verbatim":"mihi","normalized":"mihi","wordType":"INFRASPECIES","start":21,"end":25},{"verbatim":"Dussart","normalized":"Dussart","wordType":"AUTHOR_WORD","start":26,"end":33},{"verbatim":"Graf","normalized":"Graf","wordType":"AUTHOR_WORD","start":35,"end":39},{"verbatim":"Husson","normalized":"Husson","wordType":"AUTHOR_WORD","start":42,"end":48},{"verbatim":"1966","normalized":"1966","wordType":"YEAR","start":50,"end":54}],"id":"fd1806d9-761c-5b8c-9ee4-94299b9dd289","parserVersion":"test_version"}
```

### Exceptions from ranks (rank-line epithets)

Name: Selenops ab Logunov & Jäger, 2015

Canonical: Selenops ab

Authorship: Logunov & Jäger 2015

```json
{"parsed":true,"quality":1,"verbatim":"Selenops ab Logunov \u0026 Jäger, 2015","normalized":"Selenops ab Logunov \u0026 Jäger 2015","canonical":{"stemmed":"Selenops ab","simple":"Selenops ab","full":"Selenops ab"},"cardinality":2,"authorship":{"verbatim":"Logunov \u0026 Jäger, 2015","normalized":"Logunov \u0026 Jäger 2015","year":"2015","authors":["Logunov","Jäger"],"originalAuth":{"authors":["Logunov","Jäger"],"year":{"year":"2015"}}},"details":{"species":{"genus":"Selenops","species":"ab","authorship":{"verbatim":"Logunov \u0026 Jäger, 2015","normalized":"Logunov \u0026 Jäger 2015","year":"2015","authors":["Logunov","Jäger"],"originalAuth":{"authors":["Logunov","Jäger"],"year":{"year":"2015"}}}}},"words":[{"verbatim":"Selenops","normalized":"Selenops","wordType":"GENUS","start":0,"end":8},{"verbatim":"ab","normalized":"ab","wordType":"SPECIES","start":9,"end":11},{"verbatim":"Logunov","normalized":"Logunov","wordType":"AUTHOR_WORD","start":12,"end":19},{"verbatim":"Jäger","normalized":"Jäger","wordType":"AUTHOR_WORD","start":22,"end":27},{"verbatim":"2015","normalized":"2015","wordType":"YEAR","start":29,"end":33}],"id":"03859723-914e-5c1b-89ee-93e43fa98b6a","parserVersion":"test_version"}
```

Name: Helophorus (Lihelophorus) ser Zaitzev, 1908

Canonical: Helophorus ser

Authorship: Zaitzev 1908

```json
{"parsed":true,"quality":1,"verbatim":"Helophorus (Lihelophorus) ser Zaitzev, 1908","normalized":"Helophorus (Lihelophorus) ser Zaitzev 1908","canonical":{"stemmed":"Helophorus ser","simple":"Helophorus ser","full":"Helophorus ser"},"cardinality":2,"authorship":{"verbatim":"Zaitzev, 1908","normalized":"Zaitzev 1908","year":"1908","authors":["Zaitzev"],"originalAuth":{"authors":["Zaitzev"],"year":{"year":"1908"}}},"details":{"species":{"genus":"Helophorus","subgenus":"Lihelophorus","species":"ser","authorship":{"verbatim":"Zaitzev, 1908","normalized":"Zaitzev 1908","year":"1908","authors":["Zaitzev"],"originalAuth":{"authors":["Zaitzev"],"year":{"year":"1908"}}}}},"words":[{"verbatim":"Helophorus","normalized":"Helophorus","wordType":"GENUS","start":0,"end":10},{"verbatim":"Lihelophorus","normalized":"Lihelophorus","wordType":"INFRA_GENUS","start":12,"end":24},{"verbatim":"ser","normalized":"ser","wordType":"SPECIES","start":26,"end":29},{"verbatim":"Zaitzev","normalized":"Zaitzev","wordType":"AUTHOR_WORD","start":30,"end":37},{"verbatim":"1908","normalized":"1908","wordType":"YEAR","start":39,"end":43}],"id":"50392bf7-88e2-51fe-83d4-642dc0e2a887","parserVersion":"test_version"}
```

Name: Serina subser Gredler, 1898

Canonical: Serina subser

Authorship: Gredler 1898

```json
{"parsed":true,"quality":1,"verbatim":"Serina subser Gredler, 1898","normalized":"Serina subser Gredler 1898","canonical":{"stemmed":"Serina subser","simple":"Serina subser","full":"Serina subser"},"cardinality":2,"authorship":{"verbatim":"Gredler, 1898","normalized":"Gredler 1898","year":"1898","authors":["Gredler"],"originalAuth":{"authors":["Gredler"],"year":{"year":"1898"}}},"details":{"species":{"genus":"Serina","species":"subser","authorship":{"verbatim":"Gredler, 1898","normalized":"Gredler 1898","year":"1898","authors":["Gredler"],"originalAuth":{"authors":["Gredler"],"year":{"year":"1898"}}}}},"words":[{"verbatim":"Serina","normalized":"Serina","wordType":"GENUS","start":0,"end":6},{"verbatim":"subser","normalized":"subser","wordType":"SPECIES","start":7,"end":13},{"verbatim":"Gredler","normalized":"Gredler","wordType":"AUTHOR_WORD","start":14,"end":21},{"verbatim":"1898","normalized":"1898","wordType":"YEAR","start":23,"end":27}],"id":"e769d367-a02c-5079-9ae2-472c54123412","parserVersion":"test_version"}
```

Name: Serina ser Gredler, 1898

Canonical: Serina ser

Authorship: Gredler 1898

```json
{"parsed":true,"quality":1,"verbatim":"Serina ser Gredler, 1898","normalized":"Serina ser Gredler 1898","canonical":{"stemmed":"Serina ser","simple":"Serina ser","full":"Serina ser"},"cardinality":2,"authorship":{"verbatim":"Gredler, 1898","normalized":"Gredler 1898","year":"1898","authors":["Gredler"],"originalAuth":{"authors":["Gredler"],"year":{"year":"1898"}}},"details":{"species":{"genus":"Serina","species":"ser","authorship":{"verbatim":"Gredler, 1898","normalized":"Gredler 1898","year":"1898","authors":["Gredler"],"originalAuth":{"authors":["Gredler"],"year":{"year":"1898"}}}}},"words":[{"verbatim":"Serina","normalized":"Serina","wordType":"GENUS","start":0,"end":6},{"verbatim":"ser","normalized":"ser","wordType":"SPECIES","start":7,"end":10},{"verbatim":"Gredler","normalized":"Gredler","wordType":"AUTHOR_WORD","start":11,"end":18},{"verbatim":"1898","normalized":"1898","wordType":"YEAR","start":20,"end":24}],"id":"a271dad8-c530-5016-b8d0-881b4863dc6a","parserVersion":"test_version"}
```

### Exceptions from author prefixes (prefix-like epithets)

Name: Campylosphaera dela (M.N.Bramlette & F.R.Sullivan) W.W.Hay & H.Mohler

Canonical: Campylosphaera dela

Authorship: (M. N. Bramlette & F. R. Sullivan) W. W. Hay & H. Mohler

```json
{"parsed":true,"quality":1,"verbatim":"Campylosphaera dela (M.N.Bramlette \u0026 F.R.Sullivan) W.W.Hay \u0026 H.Mohler","normalized":"Campylosphaera dela (M. N. Bramlette \u0026 F. R. Sullivan) W. W. Hay \u0026 H. Mohler","canonical":{"stemmed":"Campylosphaera del","simple":"Campylosphaera dela","full":"Campylosphaera dela"},"cardinality":2,"authorship":{"verbatim":"(M.N.Bramlette \u0026 F.R.Sullivan) W.W.Hay \u0026 H.Mohler","normalized":"(M. N. Bramlette \u0026 F. R. Sullivan) W. W. Hay \u0026 H. Mohler","authors":["M. N. Bramlette","F. R. Sullivan","W. W. Hay","H. Mohler"],"originalAuth":{"authors":["M. N. Bramlette","F. R. Sullivan"]},"combinationAuth":{"authors":["W. W. Hay","H. Mohler"]}},"details":{"species":{"genus":"Campylosphaera","species":"dela","authorship":{"verbatim":"(M.N.Bramlette \u0026 F.R.Sullivan) W.W.Hay \u0026 H.Mohler","normalized":"(M. N. Bramlette \u0026 F. R. Sullivan) W. W. Hay \u0026 H. Mohler","authors":["M. N. Bramlette","F. R. Sullivan","W. W. Hay","H. Mohler"],"originalAuth":{"authors":["M. N. Bramlette","F. R. Sullivan"]},"combinationAuth":{"authors":["W. W. Hay","H. Mohler"]}}}},"words":[{"verbatim":"Campylosphaera","normalized":"Campylosphaera","wordType":"GENUS","start":0,"end":14},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":15,"end":19},{"verbatim":"M.","normalized":"M.","wordType":"AUTHOR_WORD","start":21,"end":23},{"verbatim":"N.","normalized":"N.","wordType":"AUTHOR_WORD","start":23,"end":25},{"verbatim":"Bramlette","normalized":"Bramlette","wordType":"AUTHOR_WORD","start":25,"end":34},{"verbatim":"F.","normalized":"F.","wordType":"AUTHOR_WORD","start":37,"end":39},{"verbatim":"R.","normalized":"R.","wordType":"AUTHOR_WORD","start":39,"end":41},{"verbatim":"Sullivan","normalized":"Sullivan","wordType":"AUTHOR_WORD","start":41,"end":49},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":51,"end":53},{"verbatim":"W.","normalized":"W.","wordType":"AUTHOR_WORD","start":53,"end":55},{"verbatim":"Hay","normalized":"Hay","wordType":"AUTHOR_WORD","start":55,"end":58},{"verbatim":"H.","normalized":"H.","wordType":"AUTHOR_WORD","start":61,"end":63},{"verbatim":"Mohler","normalized":"Mohler","wordType":"AUTHOR_WORD","start":63,"end":69}],"id":"3746bbe1-c63b-56ba-9591-d18767dd18a6","parserVersion":"test_version"}
```

Name: Antaplaga dela Druce, 1904

Canonical: Antaplaga dela

Authorship: Druce 1904

```json
{"parsed":true,"quality":1,"verbatim":"Antaplaga dela Druce, 1904","normalized":"Antaplaga dela Druce 1904","canonical":{"stemmed":"Antaplaga del","simple":"Antaplaga dela","full":"Antaplaga dela"},"cardinality":2,"authorship":{"verbatim":"Druce, 1904","normalized":"Druce 1904","year":"1904","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1904"}}},"details":{"species":{"genus":"Antaplaga","species":"dela","authorship":{"verbatim":"Druce, 1904","normalized":"Druce 1904","year":"1904","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1904"}}}}},"words":[{"verbatim":"Antaplaga","normalized":"Antaplaga","wordType":"GENUS","start":0,"end":9},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Druce","normalized":"Druce","wordType":"AUTHOR_WORD","start":15,"end":20},{"verbatim":"1904","normalized":"1904","wordType":"YEAR","start":22,"end":26}],"id":"061c7413-a5eb-50d0-ab38-ebc27502441b","parserVersion":"test_version"}
```

Name: Baeolidia dela (Er. Marcus & Ev. Marcus, 1960)

Canonical: Baeolidia dela

Authorship: (Er. Marcus & Ev. Marcus 1960)

```json
{"parsed":true,"quality":1,"verbatim":"Baeolidia dela (Er. Marcus \u0026 Ev. Marcus, 1960)","normalized":"Baeolidia dela (Er. Marcus \u0026 Ev. Marcus 1960)","canonical":{"stemmed":"Baeolidia del","simple":"Baeolidia dela","full":"Baeolidia dela"},"cardinality":2,"authorship":{"verbatim":"(Er. Marcus \u0026 Ev. Marcus, 1960)","normalized":"(Er. Marcus \u0026 Ev. Marcus 1960)","year":"1960","authors":["Er. Marcus","Ev. Marcus"],"originalAuth":{"authors":["Er. Marcus","Ev. Marcus"],"year":{"year":"1960"}}},"details":{"species":{"genus":"Baeolidia","species":"dela","authorship":{"verbatim":"(Er. Marcus \u0026 Ev. Marcus, 1960)","normalized":"(Er. Marcus \u0026 Ev. Marcus 1960)","year":"1960","authors":["Er. Marcus","Ev. Marcus"],"originalAuth":{"authors":["Er. Marcus","Ev. Marcus"],"year":{"year":"1960"}}}}},"words":[{"verbatim":"Baeolidia","normalized":"Baeolidia","wordType":"GENUS","start":0,"end":9},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Er.","normalized":"Er.","wordType":"AUTHOR_WORD","start":16,"end":19},{"verbatim":"Marcus","normalized":"Marcus","wordType":"AUTHOR_WORD","start":20,"end":26},{"verbatim":"Ev.","normalized":"Ev.","wordType":"AUTHOR_WORD","start":29,"end":32},{"verbatim":"Marcus","normalized":"Marcus","wordType":"AUTHOR_WORD","start":33,"end":39},{"verbatim":"1960","normalized":"1960","wordType":"YEAR","start":41,"end":45}],"id":"72c7698c-901d-5b68-924c-4ec42a658bb9","parserVersion":"test_version"}
```
Name: Dicentria dela Druce, 1894

Canonical: Dicentria dela

Authorship: Druce 1894

```json
{"parsed":true,"quality":1,"verbatim":"Dicentria dela Druce, 1894","normalized":"Dicentria dela Druce 1894","canonical":{"stemmed":"Dicentria del","simple":"Dicentria dela","full":"Dicentria dela"},"cardinality":2,"authorship":{"verbatim":"Druce, 1894","normalized":"Druce 1894","year":"1894","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1894"}}},"details":{"species":{"genus":"Dicentria","species":"dela","authorship":{"verbatim":"Druce, 1894","normalized":"Druce 1894","year":"1894","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1894"}}}}},"words":[{"verbatim":"Dicentria","normalized":"Dicentria","wordType":"GENUS","start":0,"end":9},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Druce","normalized":"Druce","wordType":"AUTHOR_WORD","start":15,"end":20},{"verbatim":"1894","normalized":"1894","wordType":"YEAR","start":22,"end":26}],"id":"207f6ad0-14c0-529e-9808-fa33b47a82ac","parserVersion":"test_version"}
```

Name: Eulaira dela Chamberlin & Ivie, 1933

Canonical: Eulaira dela

Authorship: Chamberlin & Ivie 1933

```json
{"parsed":true,"quality":1,"verbatim":"Eulaira dela Chamberlin \u0026 Ivie, 1933","normalized":"Eulaira dela Chamberlin \u0026 Ivie 1933","canonical":{"stemmed":"Eulaira del","simple":"Eulaira dela","full":"Eulaira dela"},"cardinality":2,"authorship":{"verbatim":"Chamberlin \u0026 Ivie, 1933","normalized":"Chamberlin \u0026 Ivie 1933","year":"1933","authors":["Chamberlin","Ivie"],"originalAuth":{"authors":["Chamberlin","Ivie"],"year":{"year":"1933"}}},"details":{"species":{"genus":"Eulaira","species":"dela","authorship":{"verbatim":"Chamberlin \u0026 Ivie, 1933","normalized":"Chamberlin \u0026 Ivie 1933","year":"1933","authors":["Chamberlin","Ivie"],"originalAuth":{"authors":["Chamberlin","Ivie"],"year":{"year":"1933"}}}}},"words":[{"verbatim":"Eulaira","normalized":"Eulaira","wordType":"GENUS","start":0,"end":7},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":8,"end":12},{"verbatim":"Chamberlin","normalized":"Chamberlin","wordType":"AUTHOR_WORD","start":13,"end":23},{"verbatim":"Ivie","normalized":"Ivie","wordType":"AUTHOR_WORD","start":26,"end":30},{"verbatim":"1933","normalized":"1933","wordType":"YEAR","start":32,"end":36}],"id":"160ef9a2-f484-57bf-b4ac-5721fb0e4cfc","parserVersion":"test_version"}
```

Name: Paralvinella dela Detinova, 1988

Canonical: Paralvinella dela

Authorship: Detinova 1988

```json
{"parsed":true,"quality":1,"verbatim":"Paralvinella dela Detinova, 1988","normalized":"Paralvinella dela Detinova 1988","canonical":{"stemmed":"Paralvinella del","simple":"Paralvinella dela","full":"Paralvinella dela"},"cardinality":2,"authorship":{"verbatim":"Detinova, 1988","normalized":"Detinova 1988","year":"1988","authors":["Detinova"],"originalAuth":{"authors":["Detinova"],"year":{"year":"1988"}}},"details":{"species":{"genus":"Paralvinella","species":"dela","authorship":{"verbatim":"Detinova, 1988","normalized":"Detinova 1988","year":"1988","authors":["Detinova"],"originalAuth":{"authors":["Detinova"],"year":{"year":"1988"}}}}},"words":[{"verbatim":"Paralvinella","normalized":"Paralvinella","wordType":"GENUS","start":0,"end":12},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":13,"end":17},{"verbatim":"Detinova","normalized":"Detinova","wordType":"AUTHOR_WORD","start":18,"end":26},{"verbatim":"1988","normalized":"1988","wordType":"YEAR","start":28,"end":32}],"id":"111a303b-327b-542f-8a6c-9c0aae97e142","parserVersion":"test_version"}
```

Name: Scoparia dela Clarke, 1965

Canonical: Scoparia dela

Authorship: Clarke 1965

```json
{"parsed":true,"quality":1,"verbatim":"Scoparia dela Clarke, 1965","normalized":"Scoparia dela Clarke 1965","canonical":{"stemmed":"Scoparia del","simple":"Scoparia dela","full":"Scoparia dela"},"cardinality":2,"authorship":{"verbatim":"Clarke, 1965","normalized":"Clarke 1965","year":"1965","authors":["Clarke"],"originalAuth":{"authors":["Clarke"],"year":{"year":"1965"}}},"details":{"species":{"genus":"Scoparia","species":"dela","authorship":{"verbatim":"Clarke, 1965","normalized":"Clarke 1965","year":"1965","authors":["Clarke"],"originalAuth":{"authors":["Clarke"],"year":{"year":"1965"}}}}},"words":[{"verbatim":"Scoparia","normalized":"Scoparia","wordType":"GENUS","start":0,"end":8},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":9,"end":13},{"verbatim":"Clarke","normalized":"Clarke","wordType":"AUTHOR_WORD","start":14,"end":20},{"verbatim":"1965","normalized":"1965","wordType":"YEAR","start":22,"end":26}],"id":"b2efca41-c18e-58f8-9300-0f542037e6a2","parserVersion":"test_version"}
```
Name: Tortolena dela Chamberlin & Ivie, 1941

Canonical: Tortolena dela

Authorship: Chamberlin & Ivie 1941

```json
{"parsed":true,"quality":1,"verbatim":"Tortolena dela Chamberlin \u0026 Ivie, 1941","normalized":"Tortolena dela Chamberlin \u0026 Ivie 1941","canonical":{"stemmed":"Tortolena del","simple":"Tortolena dela","full":"Tortolena dela"},"cardinality":2,"authorship":{"verbatim":"Chamberlin \u0026 Ivie, 1941","normalized":"Chamberlin \u0026 Ivie 1941","year":"1941","authors":["Chamberlin","Ivie"],"originalAuth":{"authors":["Chamberlin","Ivie"],"year":{"year":"1941"}}},"details":{"species":{"genus":"Tortolena","species":"dela","authorship":{"verbatim":"Chamberlin \u0026 Ivie, 1941","normalized":"Chamberlin \u0026 Ivie 1941","year":"1941","authors":["Chamberlin","Ivie"],"originalAuth":{"authors":["Chamberlin","Ivie"],"year":{"year":"1941"}}}}},"words":[{"verbatim":"Tortolena","normalized":"Tortolena","wordType":"GENUS","start":0,"end":9},{"verbatim":"dela","normalized":"dela","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Chamberlin","normalized":"Chamberlin","wordType":"AUTHOR_WORD","start":15,"end":25},{"verbatim":"Ivie","normalized":"Ivie","wordType":"AUTHOR_WORD","start":28,"end":32},{"verbatim":"1941","normalized":"1941","wordType":"YEAR","start":34,"end":38}],"id":"760c81cc-b336-58a3-888d-5fd85196e287","parserVersion":"test_version"}
```

Name: Semiothisa da Dyar, 1916

Canonical: Semiothisa da

Authorship: Dyar 1916

```json
{"parsed":true,"quality":1,"verbatim":"Semiothisa da Dyar, 1916","normalized":"Semiothisa da Dyar 1916","canonical":{"stemmed":"Semiothisa da","simple":"Semiothisa da","full":"Semiothisa da"},"cardinality":2,"authorship":{"verbatim":"Dyar, 1916","normalized":"Dyar 1916","year":"1916","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1916"}}},"details":{"species":{"genus":"Semiothisa","species":"da","authorship":{"verbatim":"Dyar, 1916","normalized":"Dyar 1916","year":"1916","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1916"}}}}},"words":[{"verbatim":"Semiothisa","normalized":"Semiothisa","wordType":"GENUS","start":0,"end":10},{"verbatim":"da","normalized":"da","wordType":"SPECIES","start":11,"end":13},{"verbatim":"Dyar","normalized":"Dyar","wordType":"AUTHOR_WORD","start":14,"end":18},{"verbatim":"1916","normalized":"1916","wordType":"YEAR","start":20,"end":24}],"id":"441392b1-1e13-5eb1-afbe-393f3bdfd8c0","parserVersion":"test_version"}
```

Name: Gnathopleustes den (J.L. Barnard, 1969)

Canonical: Gnathopleustes den

Authorship: (J. L. Barnard 1969)

```json
{"parsed":true,"quality":1,"verbatim":"Gnathopleustes den (J.L. Barnard, 1969)","normalized":"Gnathopleustes den (J. L. Barnard 1969)","canonical":{"stemmed":"Gnathopleustes den","simple":"Gnathopleustes den","full":"Gnathopleustes den"},"cardinality":2,"authorship":{"verbatim":"(J.L. Barnard, 1969)","normalized":"(J. L. Barnard 1969)","year":"1969","authors":["J. L. Barnard"],"originalAuth":{"authors":["J. L. Barnard"],"year":{"year":"1969"}}},"details":{"species":{"genus":"Gnathopleustes","species":"den","authorship":{"verbatim":"(J.L. Barnard, 1969)","normalized":"(J. L. Barnard 1969)","year":"1969","authors":["J. L. Barnard"],"originalAuth":{"authors":["J. L. Barnard"],"year":{"year":"1969"}}}}},"words":[{"verbatim":"Gnathopleustes","normalized":"Gnathopleustes","wordType":"GENUS","start":0,"end":14},{"verbatim":"den","normalized":"den","wordType":"SPECIES","start":15,"end":18},{"verbatim":"J.","normalized":"J.","wordType":"AUTHOR_WORD","start":20,"end":22},{"verbatim":"L.","normalized":"L.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"Barnard","normalized":"Barnard","wordType":"AUTHOR_WORD","start":25,"end":32},{"verbatim":"1969","normalized":"1969","wordType":"YEAR","start":34,"end":38}],"id":"d70496f0-2bd6-50c9-9d87-3e6ae48e6793","parserVersion":"test_version"}
```

Name: Agnetina den Cao, T.K.T. & Bae, 2006

Canonical: Agnetina den

Authorship: Cao, T. K. T. & Bae 2006

```json
{"parsed":true,"quality":1,"verbatim":"Agnetina den Cao, T.K.T. \u0026 Bae, 2006","normalized":"Agnetina den Cao, T. K. T. \u0026 Bae 2006","canonical":{"stemmed":"Agnetina den","simple":"Agnetina den","full":"Agnetina den"},"cardinality":2,"authorship":{"verbatim":"Cao, T.K.T. \u0026 Bae, 2006","normalized":"Cao, T. K. T. \u0026 Bae 2006","year":"2006","authors":["Cao","T. K. T.","Bae"],"originalAuth":{"authors":["Cao","T. K. T.","Bae"],"year":{"year":"2006"}}},"details":{"species":{"genus":"Agnetina","species":"den","authorship":{"verbatim":"Cao, T.K.T. \u0026 Bae, 2006","normalized":"Cao, T. K. T. \u0026 Bae 2006","year":"2006","authors":["Cao","T. K. T.","Bae"],"originalAuth":{"authors":["Cao","T. K. T.","Bae"],"year":{"year":"2006"}}}}},"words":[{"verbatim":"Agnetina","normalized":"Agnetina","wordType":"GENUS","start":0,"end":8},{"verbatim":"den","normalized":"den","wordType":"SPECIES","start":9,"end":12},{"verbatim":"Cao","normalized":"Cao","wordType":"AUTHOR_WORD","start":13,"end":16},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":18,"end":20},{"verbatim":"K.","normalized":"K.","wordType":"AUTHOR_WORD","start":20,"end":22},{"verbatim":"T.","normalized":"T.","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"Bae","normalized":"Bae","wordType":"AUTHOR_WORD","start":27,"end":30},{"verbatim":"2006","normalized":"2006","wordType":"YEAR","start":32,"end":36}],"id":"db92136c-7dc9-5d31-ac57-83ba03f05294","parserVersion":"test_version"}
```
Name: Desmoxytes des Srisonchai, Enghoff & Panha, 2016

Canonical: Desmoxytes des

Authorship: Srisonchai, Enghoff & Panha 2016

```json
{"parsed":true,"quality":1,"verbatim":"Desmoxytes des Srisonchai, Enghoff \u0026 Panha, 2016","normalized":"Desmoxytes des Srisonchai, Enghoff \u0026 Panha 2016","canonical":{"stemmed":"Desmoxytes des","simple":"Desmoxytes des","full":"Desmoxytes des"},"cardinality":2,"authorship":{"verbatim":"Srisonchai, Enghoff \u0026 Panha, 2016","normalized":"Srisonchai, Enghoff \u0026 Panha 2016","year":"2016","authors":["Srisonchai","Enghoff","Panha"],"originalAuth":{"authors":["Srisonchai","Enghoff","Panha"],"year":{"year":"2016"}}},"details":{"species":{"genus":"Desmoxytes","species":"des","authorship":{"verbatim":"Srisonchai, Enghoff \u0026 Panha, 2016","normalized":"Srisonchai, Enghoff \u0026 Panha 2016","year":"2016","authors":["Srisonchai","Enghoff","Panha"],"originalAuth":{"authors":["Srisonchai","Enghoff","Panha"],"year":{"year":"2016"}}}}},"words":[{"verbatim":"Desmoxytes","normalized":"Desmoxytes","wordType":"GENUS","start":0,"end":10},{"verbatim":"des","normalized":"des","wordType":"SPECIES","start":11,"end":14},{"verbatim":"Srisonchai","normalized":"Srisonchai","wordType":"AUTHOR_WORD","start":15,"end":25},{"verbatim":"Enghoff","normalized":"Enghoff","wordType":"AUTHOR_WORD","start":27,"end":34},{"verbatim":"Panha","normalized":"Panha","wordType":"AUTHOR_WORD","start":37,"end":42},{"verbatim":"2016","normalized":"2016","wordType":"YEAR","start":44,"end":48}],"id":"6cbf87ea-fb64-5cf9-b6b8-6f73bbe568b7","parserVersion":"test_version"}
```
Name: Meteorus dos Zitani, 1998

Canonical: Meteorus dos

Authorship: Zitani 1998

```json
{"parsed":true,"quality":1,"verbatim":"Meteorus dos Zitani, 1998","normalized":"Meteorus dos Zitani 1998","canonical":{"stemmed":"Meteorus dos","simple":"Meteorus dos","full":"Meteorus dos"},"cardinality":2,"authorship":{"verbatim":"Zitani, 1998","normalized":"Zitani 1998","year":"1998","authors":["Zitani"],"originalAuth":{"authors":["Zitani"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Meteorus","species":"dos","authorship":{"verbatim":"Zitani, 1998","normalized":"Zitani 1998","year":"1998","authors":["Zitani"],"originalAuth":{"authors":["Zitani"],"year":{"year":"1998"}}}}},"words":[{"verbatim":"Meteorus","normalized":"Meteorus","wordType":"GENUS","start":0,"end":8},{"verbatim":"dos","normalized":"dos","wordType":"SPECIES","start":9,"end":12},{"verbatim":"Zitani","normalized":"Zitani","wordType":"AUTHOR_WORD","start":13,"end":19},{"verbatim":"1998","normalized":"1998","wordType":"YEAR","start":21,"end":25}],"id":"8c93aded-0398-5495-bd0d-948928f982f1","parserVersion":"test_version"}
```
Name: Stenoecia dos Freyer, 1838

Canonical: Stenoecia dos

Authorship: Freyer 1838

```json
{"parsed":true,"quality":1,"verbatim":"Stenoecia dos Freyer, 1838","normalized":"Stenoecia dos Freyer 1838","canonical":{"stemmed":"Stenoecia dos","simple":"Stenoecia dos","full":"Stenoecia dos"},"cardinality":2,"authorship":{"verbatim":"Freyer, 1838","normalized":"Freyer 1838","year":"1838","authors":["Freyer"],"originalAuth":{"authors":["Freyer"],"year":{"year":"1838"}}},"details":{"species":{"genus":"Stenoecia","species":"dos","authorship":{"verbatim":"Freyer, 1838","normalized":"Freyer 1838","year":"1838","authors":["Freyer"],"originalAuth":{"authors":["Freyer"],"year":{"year":"1838"}}}}},"words":[{"verbatim":"Stenoecia","normalized":"Stenoecia","wordType":"GENUS","start":0,"end":9},{"verbatim":"dos","normalized":"dos","wordType":"SPECIES","start":10,"end":13},{"verbatim":"Freyer","normalized":"Freyer","wordType":"AUTHOR_WORD","start":14,"end":20},{"verbatim":"1838","normalized":"1838","wordType":"YEAR","start":22,"end":26}],"id":"ba1d6f79-a3b8-585f-a2fc-7fd6cfe42f19","parserVersion":"test_version"}
```

Name: Sympycnus du Curran, 1929

Canonical: Sympycnus du

Authorship: Curran 1929

```json
{"parsed":true,"quality":1,"verbatim":"Sympycnus du Curran, 1929","normalized":"Sympycnus du Curran 1929","canonical":{"stemmed":"Sympycnus du","simple":"Sympycnus du","full":"Sympycnus du"},"cardinality":2,"authorship":{"verbatim":"Curran, 1929","normalized":"Curran 1929","year":"1929","authors":["Curran"],"originalAuth":{"authors":["Curran"],"year":{"year":"1929"}}},"details":{"species":{"genus":"Sympycnus","species":"du","authorship":{"verbatim":"Curran, 1929","normalized":"Curran 1929","year":"1929","authors":["Curran"],"originalAuth":{"authors":["Curran"],"year":{"year":"1929"}}}}},"words":[{"verbatim":"Sympycnus","normalized":"Sympycnus","wordType":"GENUS","start":0,"end":9},{"verbatim":"du","normalized":"du","wordType":"SPECIES","start":10,"end":12},{"verbatim":"Curran","normalized":"Curran","wordType":"AUTHOR_WORD","start":13,"end":19},{"verbatim":"1929","normalized":"1929","wordType":"YEAR","start":21,"end":25}],"id":"abca8348-1ec7-5b02-b54e-55d4313bb383","parserVersion":"test_version"}
```

Name: Bolitoglossa la Campbell, Smith, Streicher, Acevedo & Brodie, 2010

Canonical: Bolitoglossa la

Authorship: Campbell, Smith, Streicher, Acevedo & Brodie 2010

```json
{"parsed":true,"quality":1,"verbatim":"Bolitoglossa la Campbell, Smith, Streicher, Acevedo \u0026 Brodie, 2010","normalized":"Bolitoglossa la Campbell, Smith, Streicher, Acevedo \u0026 Brodie 2010","canonical":{"stemmed":"Bolitoglossa la","simple":"Bolitoglossa la","full":"Bolitoglossa la"},"cardinality":2,"authorship":{"verbatim":"Campbell, Smith, Streicher, Acevedo \u0026 Brodie, 2010","normalized":"Campbell, Smith, Streicher, Acevedo \u0026 Brodie 2010","year":"2010","authors":["Campbell","Smith","Streicher","Acevedo","Brodie"],"originalAuth":{"authors":["Campbell","Smith","Streicher","Acevedo","Brodie"],"year":{"year":"2010"}}},"details":{"species":{"genus":"Bolitoglossa","species":"la","authorship":{"verbatim":"Campbell, Smith, Streicher, Acevedo \u0026 Brodie, 2010","normalized":"Campbell, Smith, Streicher, Acevedo \u0026 Brodie 2010","year":"2010","authors":["Campbell","Smith","Streicher","Acevedo","Brodie"],"originalAuth":{"authors":["Campbell","Smith","Streicher","Acevedo","Brodie"],"year":{"year":"2010"}}}}},"words":[{"verbatim":"Bolitoglossa","normalized":"Bolitoglossa","wordType":"GENUS","start":0,"end":12},{"verbatim":"la","normalized":"la","wordType":"SPECIES","start":13,"end":15},{"verbatim":"Campbell","normalized":"Campbell","wordType":"AUTHOR_WORD","start":16,"end":24},{"verbatim":"Smith","normalized":"Smith","wordType":"AUTHOR_WORD","start":26,"end":31},{"verbatim":"Streicher","normalized":"Streicher","wordType":"AUTHOR_WORD","start":33,"end":42},{"verbatim":"Acevedo","normalized":"Acevedo","wordType":"AUTHOR_WORD","start":44,"end":51},{"verbatim":"Brodie","normalized":"Brodie","wordType":"AUTHOR_WORD","start":54,"end":60},{"verbatim":"2010","normalized":"2010","wordType":"YEAR","start":62,"end":66}],"id":"accc157e-e26e-513d-9d40-d99ca4f2286f","parserVersion":"test_version"}
```

Name: Leptonetela la Wang & Li, 2017

Canonical: Leptonetela la

Authorship: Wang & Li 2017

```json
{"parsed":true,"quality":1,"verbatim":"Leptonetela la Wang \u0026 Li, 2017","normalized":"Leptonetela la Wang \u0026 Li 2017","canonical":{"stemmed":"Leptonetela la","simple":"Leptonetela la","full":"Leptonetela la"},"cardinality":2,"authorship":{"verbatim":"Wang \u0026 Li, 2017","normalized":"Wang \u0026 Li 2017","year":"2017","authors":["Wang","Li"],"originalAuth":{"authors":["Wang","Li"],"year":{"year":"2017"}}},"details":{"species":{"genus":"Leptonetela","species":"la","authorship":{"verbatim":"Wang \u0026 Li, 2017","normalized":"Wang \u0026 Li 2017","year":"2017","authors":["Wang","Li"],"originalAuth":{"authors":["Wang","Li"],"year":{"year":"2017"}}}}},"words":[{"verbatim":"Leptonetela","normalized":"Leptonetela","wordType":"GENUS","start":0,"end":11},{"verbatim":"la","normalized":"la","wordType":"SPECIES","start":12,"end":14},{"verbatim":"Wang","normalized":"Wang","wordType":"AUTHOR_WORD","start":15,"end":19},{"verbatim":"Li","normalized":"Li","wordType":"AUTHOR_WORD","start":22,"end":24},{"verbatim":"2017","normalized":"2017","wordType":"YEAR","start":26,"end":30}],"id":"c21c3b32-c3cb-5004-9e4b-c0812e378c0c","parserVersion":"test_version"}
```

Name: Nocaracris van Ünal, 2016

Canonical: Nocaracris van

Authorship: Ünal 2016

```json
{"parsed":true,"quality":1,"verbatim":"Nocaracris van Ünal, 2016","normalized":"Nocaracris van Ünal 2016","canonical":{"stemmed":"Nocaracris uan","simple":"Nocaracris van","full":"Nocaracris van"},"cardinality":2,"authorship":{"verbatim":"Ünal, 2016","normalized":"Ünal 2016","year":"2016","authors":["Ünal"],"originalAuth":{"authors":["Ünal"],"year":{"year":"2016"}}},"details":{"species":{"genus":"Nocaracris","species":"van","authorship":{"verbatim":"Ünal, 2016","normalized":"Ünal 2016","year":"2016","authors":["Ünal"],"originalAuth":{"authors":["Ünal"],"year":{"year":"2016"}}}}},"words":[{"verbatim":"Nocaracris","normalized":"Nocaracris","wordType":"GENUS","start":0,"end":10},{"verbatim":"van","normalized":"van","wordType":"SPECIES","start":11,"end":14},{"verbatim":"Ünal","normalized":"Ünal","wordType":"AUTHOR_WORD","start":15,"end":19},{"verbatim":"2016","normalized":"2016","wordType":"YEAR","start":21,"end":25}],"id":"da0de28d-1241-56dc-8a6a-d2f42824d283","parserVersion":"test_version"}
```

Name: Zodarion van Bosmans, 2009

Canonical: Zodarion van

Authorship: Bosmans 2009

```json
{"parsed":true,"quality":1,"verbatim":"Zodarion van Bosmans, 2009","normalized":"Zodarion van Bosmans 2009","canonical":{"stemmed":"Zodarion uan","simple":"Zodarion van","full":"Zodarion van"},"cardinality":2,"authorship":{"verbatim":"Bosmans, 2009","normalized":"Bosmans 2009","year":"2009","authors":["Bosmans"],"originalAuth":{"authors":["Bosmans"],"year":{"year":"2009"}}},"details":{"species":{"genus":"Zodarion","species":"van","authorship":{"verbatim":"Bosmans, 2009","normalized":"Bosmans 2009","year":"2009","authors":["Bosmans"],"originalAuth":{"authors":["Bosmans"],"year":{"year":"2009"}}}}},"words":[{"verbatim":"Zodarion","normalized":"Zodarion","wordType":"GENUS","start":0,"end":8},{"verbatim":"van","normalized":"van","wordType":"SPECIES","start":9,"end":12},{"verbatim":"Bosmans","normalized":"Bosmans","wordType":"AUTHOR_WORD","start":13,"end":20},{"verbatim":"2009","normalized":"2009","wordType":"YEAR","start":22,"end":26}],"id":"f9e66138-6fdd-524a-bf7e-0ea56f4f0484","parserVersion":"test_version"}
```

Name: Malamatidia zu Jäger & Dankittipakul, 2010

Canonical: Malamatidia zu

Authorship: Jäger & Dankittipakul 2010

```json
{"parsed":true,"quality":1,"verbatim":"Malamatidia zu Jäger \u0026 Dankittipakul, 2010","normalized":"Malamatidia zu Jäger \u0026 Dankittipakul 2010","canonical":{"stemmed":"Malamatidia zu","simple":"Malamatidia zu","full":"Malamatidia zu"},"cardinality":2,"authorship":{"verbatim":"Jäger \u0026 Dankittipakul, 2010","normalized":"Jäger \u0026 Dankittipakul 2010","year":"2010","authors":["Jäger","Dankittipakul"],"originalAuth":{"authors":["Jäger","Dankittipakul"],"year":{"year":"2010"}}},"details":{"species":{"genus":"Malamatidia","species":"zu","authorship":{"verbatim":"Jäger \u0026 Dankittipakul, 2010","normalized":"Jäger \u0026 Dankittipakul 2010","year":"2010","authors":["Jäger","Dankittipakul"],"originalAuth":{"authors":["Jäger","Dankittipakul"],"year":{"year":"2010"}}}}},"words":[{"verbatim":"Malamatidia","normalized":"Malamatidia","wordType":"GENUS","start":0,"end":11},{"verbatim":"zu","normalized":"zu","wordType":"SPECIES","start":12,"end":14},{"verbatim":"Jäger","normalized":"Jäger","wordType":"AUTHOR_WORD","start":15,"end":20},{"verbatim":"Dankittipakul","normalized":"Dankittipakul","wordType":"AUTHOR_WORD","start":23,"end":36},{"verbatim":"2010","normalized":"2010","wordType":"YEAR","start":38,"end":42}],"id":"0b16fdbe-1edd-57ea-aea6-248237d7bacb","parserVersion":"test_version"}
```

### Exceptions from author suffixes (suffix-like epithets)

Name: Ruteloryctes bis Dechambre, 2006

Canonical: Ruteloryctes bis

Authorship: Dechambre 2006

```json
{"parsed":true,"quality":1,"verbatim":"Ruteloryctes bis Dechambre, 2006","normalized":"Ruteloryctes bis Dechambre 2006","canonical":{"stemmed":"Ruteloryctes bis","simple":"Ruteloryctes bis","full":"Ruteloryctes bis"},"cardinality":2,"authorship":{"verbatim":"Dechambre, 2006","normalized":"Dechambre 2006","year":"2006","authors":["Dechambre"],"originalAuth":{"authors":["Dechambre"],"year":{"year":"2006"}}},"details":{"species":{"genus":"Ruteloryctes","species":"bis","authorship":{"verbatim":"Dechambre, 2006","normalized":"Dechambre 2006","year":"2006","authors":["Dechambre"],"originalAuth":{"authors":["Dechambre"],"year":{"year":"2006"}}}}},"words":[{"verbatim":"Ruteloryctes","normalized":"Ruteloryctes","wordType":"GENUS","start":0,"end":12},{"verbatim":"bis","normalized":"bis","wordType":"SPECIES","start":13,"end":16},{"verbatim":"Dechambre","normalized":"Dechambre","wordType":"AUTHOR_WORD","start":17,"end":26},{"verbatim":"2006","normalized":"2006","wordType":"YEAR","start":28,"end":32}],"id":"ec9442cc-46cf-5451-ab72-e5aca85d26c0","parserVersion":"test_version"}
```

### Not parsed OCR errors to get better precision/recall ratio

Name: Mom.alpium (Osbeck, 1778)

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Mom.alpium (Osbeck, 1778)","cardinality":0,"id":"f1452bcf-b779-5d98-bfc8-56455105e3f5","parserVersion":"test_version"}
```

### No parsing -- Genera abbreviated to 3 letters (too rare)

Name: Gen. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Gen. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma","cardinality":0,"id":"54d27b31-2fbd-56e1-85e1-1438970f8953","parserVersion":"test_version"}
```

Name: Genn. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Genn. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma","cardinality":0,"id":"8edd1515-a4a1-52c5-ad1b-df7f112e68a9","parserVersion":"test_version"}
```

### No parsing -- incertae sedis

Name: Incertae sedis

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Incertae sedis","cardinality":0,"id":"74d54496-7f1c-52f8-81a9-9a9fb3a25ecb","parserVersion":"test_version"}
```

Name: </i>Hipponicidae<i> incertae sedis</i>

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"\u003c/i\u003eHipponicidae\u003ci\u003e incertae sedis\u003c/i\u003e","cardinality":0,"id":"5967f6bf-f4c7-5ea3-a7f3-fe16a7ee88e0","parserVersion":"test_version"}
```

Name: incertae sedis

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"incertae sedis","cardinality":0,"id":"14f6de42-21d9-5e67-89cd-a05ebd974a1b","parserVersion":"test_version"}
```

Name: Inc.   sed.

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Inc.   sed.","cardinality":0,"id":"2e1319c9-a44b-531c-8964-67025bbf3b40","parserVersion":"test_version"}
```

Name: inc.sed.

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"inc.sed.","cardinality":0,"id":"dbb95e14-cebc-56a9-a1d2-a70d4b759e8d","parserVersion":"test_version"}
```

Name: inc.   sed.

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"inc.   sed.","cardinality":0,"id":"f5245bf6-a459-5602-9979-02ba9428cf17","parserVersion":"test_version"}
```

Name: Incertaesedis obscuricornis Fairmaire LMH 1893

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Incertaesedis obscuricornis Fairmaire LMH 1893","cardinality":0,"id":"2601fa55-350f-5591-a549-c558284d6e9e","parserVersion":"test_version"}
```

Name: Uropodoideaincertaesedis

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Uropodoideaincertaesedis","cardinality":0,"id":"3bf556bb-ea7c-536e-8b62-93ba329c559d","parserVersion":"test_version"}
```

### No parsing -- bacterium, Candidatus

Name: Acidobacteria bacterium

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Acidobacteria bacterium","cardinality":0,"id":"c982b4fd-c41a-5987-bcc8-989c4164b9ec","parserVersion":"test_version"}
```

Name: Oscillatoriales cyanobacterium PCC 10608

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Oscillatoriales cyanobacterium PCC 10608","cardinality":0,"id":"80236e5d-1d14-5279-864d-28957b98adf8","parserVersion":"test_version"}
```

Name: Acidimicrobiales bacterium JGI 01_E13

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Acidimicrobiales bacterium JGI 01_E13","cardinality":0,"id":"8b71a29b-4271-5a83-8a92-5dab1d9dc4c3","parserVersion":"test_version"}
```

Name: Acidobacterium ailaaui Myers & King, 2016

Canonical: Acidobacterium ailaaui

Authorship: Myers & King 2016

```json
{"parsed":true,"quality":1,"verbatim":"Acidobacterium ailaaui Myers \u0026 King, 2016","normalized":"Acidobacterium ailaaui Myers \u0026 King 2016","canonical":{"stemmed":"Acidobacterium ailaau","simple":"Acidobacterium ailaaui","full":"Acidobacterium ailaaui"},"cardinality":2,"authorship":{"verbatim":"Myers \u0026 King, 2016","normalized":"Myers \u0026 King 2016","year":"2016","authors":["Myers","King"],"originalAuth":{"authors":["Myers","King"],"year":{"year":"2016"}}},"bacteria":"yes","details":{"species":{"genus":"Acidobacterium","species":"ailaaui","authorship":{"verbatim":"Myers \u0026 King, 2016","normalized":"Myers \u0026 King 2016","year":"2016","authors":["Myers","King"],"originalAuth":{"authors":["Myers","King"],"year":{"year":"2016"}}}}},"words":[{"verbatim":"Acidobacterium","normalized":"Acidobacterium","wordType":"GENUS","start":0,"end":14},{"verbatim":"ailaaui","normalized":"ailaaui","wordType":"SPECIES","start":15,"end":22},{"verbatim":"Myers","normalized":"Myers","wordType":"AUTHOR_WORD","start":23,"end":28},{"verbatim":"King","normalized":"King","wordType":"AUTHOR_WORD","start":31,"end":35},{"verbatim":"2016","normalized":"2016","wordType":"YEAR","start":37,"end":41}],"id":"b9f4555f-d2e0-5d40-acde-2b546a28a7fc","parserVersion":"test_version"}
```

Name: Candidatus Amesbacteria bacterium GW2011_GWC1_46_24

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Candidatus Amesbacteria bacterium GW2011_GWC1_46_24","cardinality":0,"id":"83382178-94bf-5bf3-a8c8-fdbca4af927c","parserVersion":"test_version"}
```

Name: Candidatus

Canonical: Candidatus

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Candidatus","normalized":"Candidatus","canonical":{"stemmed":"Candidatus","simple":"Candidatus","full":"Candidatus"},"cardinality":1,"details":{"uninomial":{"uninomial":"Candidatus"}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"UNINOMIAL","start":0,"end":10}],"id":"fb9138ac-ae7a-58c9-a912-d31d0a4eeed3","parserVersion":"test_version"}
```

Name: Candidatus Puniceispirillum Oh, Kwon, Kang, Kang, Lee, Kim & Cho, 2010

Canonical: Candidatus Puniceispirillum

Authorship: Oh, Kwon, Kang, Kang, Lee, Kim & Cho 2010

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Bacterial `Candidatus` name"}],"verbatim":"Candidatus Puniceispirillum Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho, 2010","normalized":"Candidatus Puniceispirillum Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho 2010","canonical":{"stemmed":"Puniceispirillum","simple":"Puniceispirillum","full":"Candidatus Puniceispirillum"},"cardinality":1,"authorship":{"verbatim":"Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho, 2010","normalized":"Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho 2010","year":"2010","authors":["Oh","Kwon","Kang","Lee","Kim","Cho"],"originalAuth":{"authors":["Oh","Kwon","Kang","Kang","Lee","Kim","Cho"],"year":{"year":"2010"}}},"bacteria":"yes","details":{"uninomial":{"uninomial":"Puniceispirillum","authorship":{"verbatim":"Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho, 2010","normalized":"Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho 2010","year":"2010","authors":["Oh","Kwon","Kang","Lee","Kim","Cho"],"originalAuth":{"authors":["Oh","Kwon","Kang","Kang","Lee","Kim","Cho"],"year":{"year":"2010"}}}}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"CANDIDATUS","start":0,"end":10},{"verbatim":"Puniceispirillum","normalized":"Puniceispirillum","wordType":"UNINOMIAL","start":11,"end":27},{"verbatim":"Oh","normalized":"Oh","wordType":"AUTHOR_WORD","start":28,"end":30},{"verbatim":"Kwon","normalized":"Kwon","wordType":"AUTHOR_WORD","start":32,"end":36},{"verbatim":"Kang","normalized":"Kang","wordType":"AUTHOR_WORD","start":38,"end":42},{"verbatim":"Kang","normalized":"Kang","wordType":"AUTHOR_WORD","start":44,"end":48},{"verbatim":"Lee","normalized":"Lee","wordType":"AUTHOR_WORD","start":50,"end":53},{"verbatim":"Kim","normalized":"Kim","wordType":"AUTHOR_WORD","start":55,"end":58},{"verbatim":"Cho","normalized":"Cho","wordType":"AUTHOR_WORD","start":61,"end":64},{"verbatim":"2010","normalized":"2010","wordType":"YEAR","start":66,"end":70}],"id":"82fde2e2-8e50-5fd0-8ffe-96f34f85505b","parserVersion":"test_version"}

```

Name: Candidatus Halobonum

Canonical: Candidatus Halobonum

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Bacterial `Candidatus` name"}],"verbatim":"Candidatus Halobonum","normalized":"Candidatus Halobonum","canonical":{"stemmed":"Halobonum","simple":"Halobonum","full":"Candidatus Halobonum"},"cardinality":1,"bacteria":"yes","details":{"uninomial":{"uninomial":"Halobonum"}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"CANDIDATUS","start":0,"end":10},{"verbatim":"Halobonum","normalized":"Halobonum","wordType":"UNINOMIAL","start":11,"end":20}],"id":"289152c0-1042-5cac-a649-44314b25c857","parserVersion":"test_version"}
```

Name: Candidatus Endomicrobium sp. MdDo-005

Canonical: Candidatus Endomicrobium

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"},{"quality":2,"warning":"Bacterial `Candidatus` name"}],"verbatim":"Candidatus Endomicrobium sp. MdDo-005","normalized":"Candidatus Endomicrobium","canonical":{"stemmed":"Endomicrobium","simple":"Endomicrobium","full":"Candidatus Endomicrobium"},"cardinality":0,"bacteria":"yes","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Endomicrobium","approximationMarker":"sp.","ignored":" MdDo-005"}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"CANDIDATUS","start":0,"end":10},{"verbatim":"Endomicrobium","normalized":"Endomicrobium","wordType":"GENUS","start":11,"end":24},{"verbatim":"sp.","normalized":"sp.","wordType":"APPROXIMATION_MARKER","start":25,"end":28}],"id":"f9231593-37a4-5e11-b3e8-3963f90b37e8","parserVersion":"test_version"}
```

Name: Candidatus Abawacabacteria bacterium

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Candidatus Abawacabacteria bacterium","cardinality":0,"id":"33ac7170-8bed-5051-9dd8-c6aac30a95cd","parserVersion":"test_version"}
```

Name: Candidatus Accumulibacter phosphatis clade IIA str. UW-1

Canonical: Candidatus Accumulibacter phosphatis

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Bacterial `Candidatus` name"}],"verbatim":"Candidatus Accumulibacter phosphatis clade IIA str. UW-1","normalized":"Candidatus Accumulibacter phosphatis","canonical":{"stemmed":"Accumulibacter phosphat","simple":"Accumulibacter phosphatis","full":"Candidatus Accumulibacter phosphatis"},"cardinality":2,"bacteria":"yes","tail":" clade IIA str. UW-1","details":{"species":{"genus":"Accumulibacter","species":"phosphatis"}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"CANDIDATUS","start":0,"end":10},{"verbatim":"Accumulibacter","normalized":"Accumulibacter","wordType":"GENUS","start":11,"end":25},{"verbatim":"phosphatis","normalized":"phosphatis","wordType":"SPECIES","start":26,"end":36}],"id":"0c1f98d9-0c9a-5750-8e44-3e4156f04825","parserVersion":"test_version"}
```

Name: Candidatus Anammoxoglobus environmental samples

Canonical: Candidatus Anammoxoglobus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Bacterial `Candidatus` name"}],"verbatim":"Candidatus Anammoxoglobus environmental samples","normalized":"Candidatus Anammoxoglobus","canonical":{"stemmed":"Anammoxoglobus","simple":"Anammoxoglobus","full":"Candidatus Anammoxoglobus"},"cardinality":1,"bacteria":"yes","tail":" environmental samples","details":{"uninomial":{"uninomial":"Anammoxoglobus"}},"words":[{"verbatim":"Candidatus","normalized":"Candidatus","wordType":"CANDIDATUS","start":0,"end":10},{"verbatim":"Anammoxoglobus","normalized":"Anammoxoglobus","wordType":"UNINOMIAL","start":11,"end":25}],"id":"c2c440df-a095-59bc-b2b7-ed79460af6a3","parserVersion":"test_version"}
```
### No parsing -- 'Not', 'None', 'Unidentified'  phrases

Name: None recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"None recorded","cardinality":0,"id":"54d66439-b10d-50dc-a659-c9bce413ed5d","parserVersion":"test_version"}
```

Name: NONE recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"NONE recorded","cardinality":0,"id":"cedc6de2-aed6-58dc-904f-a14348588f8a","parserVersion":"test_version"}
```

Name: NoNe recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"NoNe recorded","cardinality":0,"id":"39682f61-d0d0-5dc0-bf57-b73ffb97b3ef","parserVersion":"test_version"}
```

Name: None

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"None","cardinality":0,"id":"8cf8696e-6ca6-5ec7-b441-e04a37ea751c","parserVersion":"test_version"}
```

Name: unidentified recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"unidentified recorded","cardinality":0,"id":"4c391bc1-d3f6-5e33-80df-262cbfb09dfe","parserVersion":"test_version"}
```

Name: UniDentiFied recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"UniDentiFied recorded","cardinality":0,"id":"57b55b46-c874-59ae-b3d8-2888d8a3bc1c","parserVersion":"test_version"}
```

Name: not recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"not recorded","cardinality":0,"id":"830df5b1-ef3b-5240-8ecf-4fd74c2fff72","parserVersion":"test_version"}
```

Name: NOT recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"NOT recorded","cardinality":0,"id":"52b51d9e-29db-561c-84ac-cd1592c762c1","parserVersion":"test_version"}
```

Name: Not recorded

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Not recorded","cardinality":0,"id":"025b92f4-2b2c-5593-a02b-66f121b0a42b","parserVersion":"test_version"}
```

Name: Not assigned

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Not assigned","cardinality":0,"id":"19bffdbe-f1c7-5d39-b7b6-3dc96a317c4b","parserVersion":"test_version"}
```

Name: Notassigned

Canonical: Notassigned

Authorship:

```json
{"parsed":true,"quality":1,"verbatim":"Notassigned","normalized":"Notassigned","canonical":{"stemmed":"Notassigned","simple":"Notassigned","full":"Notassigned"},"cardinality":1,"details":{"uninomial":{"uninomial":"Notassigned"}},"words":[{"verbatim":"Notassigned","normalized":"Notassigned","wordType":"UNINOMIAL","start":0,"end":11}],"id":"8c07b58a-be4e-5c31-871b-cffe36b9860a","parserVersion":"test_version"}
```

Name: Unnamed clade

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Unnamed clade","cardinality":0,"id":"d510b662-0a4d-5678-a1a7-c58b20d25fa0","parserVersion":"test_version"}
```

Name: Unamed clade

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Unamed clade","cardinality":0,"id":"be6943d3-fa83-5e5d-9515-7cc339473d4d","parserVersion":"test_version"}
```

### No parsing -- genus with apostrophe

Name: Abbott's moray eel

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Abbott's moray eel","cardinality":0,"id":"6a870e4b-5cc5-5226-ac5d-b769521b640f","parserVersion":"test_version"}
```

Name: Chambers' twinpod

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Chambers' twinpod","cardinality":0,"id":"f109486d-9809-5196-b135-75f4cf9d7ef6","parserVersion":"test_version"}
```

Name: Columnea × Alladin's

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Columnea × Alladin's","cardinality":0,"id":"bc01a624-d49e-588d-b49d-253ac7e12939","parserVersion":"test_version"}
```

Name: Hawai'i silversword

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Hawai'i silversword","cardinality":0,"id":"f4ba0445-a5f2-525c-97ce-9316fe16e3cd","parserVersion":"test_version"}
```

### No parsing -- CamelCase 'genus' word

Name: PomaTomus

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"PomaTomus","cardinality":0,"id":"106ff909-e787-52b2-9139-25d0eb7d161e","parserVersion":"test_version"}
```

Name: DizygopUwa stosei

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"DizygopUwa stosei","cardinality":0,"id":"46511ef9-02d8-5f24-8364-b72df3e1494d","parserVersion":"test_version"}
```

Name: Oxytox[idae] Lindermann

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Oxytox[idae] Lindermann","cardinality":0,"id":"39a37760-d9f9-54d6-b49b-f6830e59f34e","parserVersion":"test_version"}
```

Name: ScarabaeinGCsp.

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"ScarabaeinGCsp.","cardinality":0,"id":"c84b775e-cc80-588f-b7bb-0094bab2c6a2","parserVersion":"test_version"}
```

### No parsing -- phytoplasma

Name: Alfalfa witches'-broom phytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Alfalfa witches'-broom phytoplasma","cardinality":0,"id":"b31676ed-c1ed-522c-8380-19a27af11e0d","parserVersion":"test_version"}
```

Name: Allium ampeloprasumphytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Allium ampeloprasumphytoplasma","cardinality":0,"id":"f84e58c5-8e49-5b2d-a4d0-4f1e538c8c7c","parserVersion":"test_version"}
```

Name: Alstroemeria sp. phytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Alstroemeria sp. phytoplasma","cardinality":0,"id":"5348845f-c94a-5c7e-bba1-307e4c07a42d","parserVersion":"test_version"}
```

### No parsing symbiont

Name: Alvinella pompejana symbiont

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Alvinella pompejana symbiont","cardinality":0,"id":"9bbe5639-7f50-5abd-a2df-78674e0cf583","parserVersion":"test_version"}
```

Name: Acyrthosiphon kondoi endosymbiont

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Acyrthosiphon kondoi endosymbiont","cardinality":0,"id":"fc3764e6-3154-5e7f-8925-fccaca0dc8f0","parserVersion":"test_version"}
```

Name: Burkholderia sp. (Gigaspora margarita endosymbiont)

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Burkholderia sp. (Gigaspora margarita endosymbiont)","cardinality":0,"id":"f6d482de-0884-54e0-95a4-d52bc26870a0","parserVersion":"test_version"}
```

Name: Dictyochloropsis symbiontica Tschermak-Woess

Canonical: Dictyochloropsis symbiontica

Authorship: Tschermak-Woess

```json
{"parsed":true,"quality":1,"verbatim":"Dictyochloropsis symbiontica Tschermak-Woess","normalized":"Dictyochloropsis symbiontica Tschermak-Woess","canonical":{"stemmed":"Dictyochloropsis symbiontic","simple":"Dictyochloropsis symbiontica","full":"Dictyochloropsis symbiontica"},"cardinality":2,"authorship":{"verbatim":"Tschermak-Woess","normalized":"Tschermak-Woess","authors":["Tschermak-Woess"],"originalAuth":{"authors":["Tschermak-Woess"]}},"details":{"species":{"genus":"Dictyochloropsis","species":"symbiontica","authorship":{"verbatim":"Tschermak-Woess","normalized":"Tschermak-Woess","authors":["Tschermak-Woess"],"originalAuth":{"authors":["Tschermak-Woess"]}}}},"words":[{"verbatim":"Dictyochloropsis","normalized":"Dictyochloropsis","wordType":"GENUS","start":0,"end":16},{"verbatim":"symbiontica","normalized":"symbiontica","wordType":"SPECIES","start":17,"end":28},{"verbatim":"Tschermak-Woess","normalized":"Tschermak-Woess","wordType":"AUTHOR_WORD","start":29,"end":44}],"id":"a8c3f410-92d4-5cd5-a659-b3e078c21947","parserVersion":"test_version"}
```

Name: Dylakosoma symbionticum var. valens Skuja

Canonical: Dylakosoma symbionticum var. valens

Authorship: Skuja

```json
{"parsed":true,"quality":1,"verbatim":"Dylakosoma symbionticum var. valens Skuja","normalized":"Dylakosoma symbionticum var. valens Skuja","canonical":{"stemmed":"Dylakosoma symbiontic ualens","simple":"Dylakosoma symbionticum valens","full":"Dylakosoma symbionticum var. valens"},"cardinality":3,"authorship":{"verbatim":"Skuja","normalized":"Skuja","authors":["Skuja"],"originalAuth":{"authors":["Skuja"]}},"details":{"infraspecies":{"genus":"Dylakosoma","species":"symbionticum","infraspecies":[{"value":"valens","rank":"var.","authorship":{"verbatim":"Skuja","normalized":"Skuja","authors":["Skuja"],"originalAuth":{"authors":["Skuja"]}}}]}},"words":[{"verbatim":"Dylakosoma","normalized":"Dylakosoma","wordType":"GENUS","start":0,"end":10},{"verbatim":"symbionticum","normalized":"symbionticum","wordType":"SPECIES","start":11,"end":23},{"verbatim":"var.","normalized":"var.","wordType":"RANK","start":24,"end":28},{"verbatim":"valens","normalized":"valens","wordType":"INFRASPECIES","start":29,"end":35},{"verbatim":"Skuja","normalized":"Skuja","wordType":"AUTHOR_WORD","start":36,"end":41}],"id":"845f2025-acea-54b5-a974-3c9ad13065a9","parserVersion":"test_version"}
```

Name: Wolbachia endosymbiont of Leptogenys gracilis

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"Wolbachia endosymbiont of Leptogenys gracilis","cardinality":0,"id":"ed4bbf5e-068a-518a-8eb3-42ead52b941b","parserVersion":"test_version"}
```

### Names with spec., nov spec

Name: Lampona spec Platnick, 2000

Canonical: Lampona spec

Authorship: Platnick 2000

```json
{"parsed":true,"quality":1,"verbatim":"Lampona spec Platnick, 2000","normalized":"Lampona spec Platnick 2000","canonical":{"stemmed":"Lampona spec","simple":"Lampona spec","full":"Lampona spec"},"cardinality":2,"authorship":{"verbatim":"Platnick, 2000","normalized":"Platnick 2000","year":"2000","authors":["Platnick"],"originalAuth":{"authors":["Platnick"],"year":{"year":"2000"}}},"details":{"species":{"genus":"Lampona","species":"spec","authorship":{"verbatim":"Platnick, 2000","normalized":"Platnick 2000","year":"2000","authors":["Platnick"],"originalAuth":{"authors":["Platnick"],"year":{"year":"2000"}}}}},"words":[{"verbatim":"Lampona","normalized":"Lampona","wordType":"GENUS","start":0,"end":7},{"verbatim":"spec","normalized":"spec","wordType":"SPECIES","start":8,"end":12},{"verbatim":"Platnick","normalized":"Platnick","wordType":"AUTHOR_WORD","start":13,"end":21},{"verbatim":"2000","normalized":"2000","wordType":"YEAR","start":23,"end":27}],"id":"d05d7916-4868-57f6-a97b-c46886f29cd8","parserVersion":"test_version"}
```

Name: Gobiosoma spec (Ginsburg, 1939)

Canonical: Gobiosoma spec

Authorship: (Ginsburg 1939)

```json
{"parsed":true,"quality":1,"verbatim":"Gobiosoma spec (Ginsburg, 1939)","normalized":"Gobiosoma spec (Ginsburg 1939)","canonical":{"stemmed":"Gobiosoma spec","simple":"Gobiosoma spec","full":"Gobiosoma spec"},"cardinality":2,"authorship":{"verbatim":"(Ginsburg, 1939)","normalized":"(Ginsburg 1939)","year":"1939","authors":["Ginsburg"],"originalAuth":{"authors":["Ginsburg"],"year":{"year":"1939"}}},"details":{"species":{"genus":"Gobiosoma","species":"spec","authorship":{"verbatim":"(Ginsburg, 1939)","normalized":"(Ginsburg 1939)","year":"1939","authors":["Ginsburg"],"originalAuth":{"authors":["Ginsburg"],"year":{"year":"1939"}}}}},"words":[{"verbatim":"Gobiosoma","normalized":"Gobiosoma","wordType":"GENUS","start":0,"end":9},{"verbatim":"spec","normalized":"spec","wordType":"SPECIES","start":10,"end":14},{"verbatim":"Ginsburg","normalized":"Ginsburg","wordType":"AUTHOR_WORD","start":16,"end":24},{"verbatim":"1939","normalized":"1939","wordType":"YEAR","start":26,"end":30}],"id":"eb47c188-86fd-54c4-a058-48a980f9419f","parserVersion":"test_version"}
```

Name: Globigerina spec

Canonical: Globigerina

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Globigerina spec","normalized":"Globigerina","canonical":{"stemmed":"Globigerina","simple":"Globigerina","full":"Globigerina"},"cardinality":1,"tail":" spec","details":{"uninomial":{"uninomial":"Globigerina"}},"words":[{"verbatim":"Globigerina","normalized":"Globigerina","wordType":"UNINOMIAL","start":0,"end":11}],"id":"4f8f7189-42a0-59e2-8d6f-67c3889673d9","parserVersion":"test_version"}
```

Name: Eunotia genuflexa Norpel-Schempp nov spec

Canonical: Eunotia genuflexa

Authorship: Norpel-Schempp

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Eunotia genuflexa Norpel-Schempp nov spec","normalized":"Eunotia genuflexa Norpel-Schempp","canonical":{"stemmed":"Eunotia genuflex","simple":"Eunotia genuflexa","full":"Eunotia genuflexa"},"cardinality":2,"authorship":{"verbatim":"Norpel-Schempp","normalized":"Norpel-Schempp","authors":["Norpel-Schempp"],"originalAuth":{"authors":["Norpel-Schempp"]}},"tail":" nov spec","details":{"species":{"genus":"Eunotia","species":"genuflexa","authorship":{"verbatim":"Norpel-Schempp","normalized":"Norpel-Schempp","authors":["Norpel-Schempp"],"originalAuth":{"authors":["Norpel-Schempp"]}}}},"words":[{"verbatim":"Eunotia","normalized":"Eunotia","wordType":"GENUS","start":0,"end":7},{"verbatim":"genuflexa","normalized":"genuflexa","wordType":"SPECIES","start":8,"end":17},{"verbatim":"Norpel-Schempp","normalized":"Norpel-Schempp","wordType":"AUTHOR_WORD","start":18,"end":32}],"id":"4cc2a699-d38d-5337-8a44-ecc0f79ef138","parserVersion":"test_version"}
```

Name: Ctenotus spec.

Canonical: Ctenotus

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ctenotus spec.","normalized":"Ctenotus","canonical":{"stemmed":"Ctenotus","simple":"Ctenotus","full":"Ctenotus"},"cardinality":1,"tail":" spec.","details":{"uninomial":{"uninomial":"Ctenotus"}},"words":[{"verbatim":"Ctenotus","normalized":"Ctenotus","wordType":"UNINOMIAL","start":0,"end":8}],"id":"991b9ee5-2f56-56e7-a29b-86c47a4901bb","parserVersion":"test_version"}
```

Name: Byrsophlebidae spec. 2

Canonical: Byrsophlebidae

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Byrsophlebidae spec. 2","normalized":"Byrsophlebidae","canonical":{"stemmed":"Byrsophlebidae","simple":"Byrsophlebidae","full":"Byrsophlebidae"},"cardinality":1,"tail":" spec. 2","details":{"uninomial":{"uninomial":"Byrsophlebidae"}},"words":[{"verbatim":"Byrsophlebidae","normalized":"Byrsophlebidae","wordType":"UNINOMIAL","start":0,"end":14}],"id":"3b07753b-71e2-5602-9a6e-bf91e672d834","parserVersion":"test_version"}
```

Name: Naviculadicta witkowskii LB & Metzeltin nov spec

Canonical: Naviculadicta witkowskii

Authorship: LB & Metzeltin

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Naviculadicta witkowskii LB \u0026 Metzeltin nov spec","normalized":"Naviculadicta witkowskii LB \u0026 Metzeltin","canonical":{"stemmed":"Naviculadicta witkowsk","simple":"Naviculadicta witkowskii","full":"Naviculadicta witkowskii"},"cardinality":2,"authorship":{"verbatim":"LB \u0026 Metzeltin","normalized":"LB \u0026 Metzeltin","authors":["LB","Metzeltin"],"originalAuth":{"authors":["LB","Metzeltin"]}},"tail":" nov spec","details":{"species":{"genus":"Naviculadicta","species":"witkowskii","authorship":{"verbatim":"LB \u0026 Metzeltin","normalized":"LB \u0026 Metzeltin","authors":["LB","Metzeltin"],"originalAuth":{"authors":["LB","Metzeltin"]}}}},"words":[{"verbatim":"Naviculadicta","normalized":"Naviculadicta","wordType":"GENUS","start":0,"end":13},{"verbatim":"witkowskii","normalized":"witkowskii","wordType":"SPECIES","start":14,"end":24},{"verbatim":"LB","normalized":"LB","wordType":"AUTHOR_WORD","start":25,"end":27},{"verbatim":"Metzeltin","normalized":"Metzeltin","wordType":"AUTHOR_WORD","start":30,"end":39}],"id":"c4dd80b7-984b-51f8-a4ec-573b4b32358b","parserVersion":"test_version"}
```

### HTML tags and entities

Name: Velutina haliotoides (Linnaeus, 1758) <i>sensu</i> Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758) \u003ci\u003esensu\u003c/i\u003e Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":" sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Velutina","normalized":"Velutina","wordType":"GENUS","start":0,"end":8},{"verbatim":"haliotoides","normalized":"haliotoides","wordType":"SPECIES","start":9,"end":20},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":32,"end":36}],"id":"189c94f6-96aa-52bb-b019-103a2103ce21","parserVersion":"test_version"}
```

Name: Velutina haliotoides (Linnaeus, 1758), <i>sensu</i> Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758), \u003ci\u003esensu\u003c/i\u003e Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":", sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Velutina","normalized":"Velutina","wordType":"GENUS","start":0,"end":8},{"verbatim":"haliotoides","normalized":"haliotoides","wordType":"SPECIES","start":9,"end":20},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":22,"end":30},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":32,"end":36}],"id":"b8d77a78-2698-5050-9c7a-638f615bd357","parserVersion":"test_version"}
```

Name: <i>Velutina halioides</i> (Linnaeus, 1758)

Canonical: Velutina halioides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"\u003ci\u003eVelutina halioides\u003c/i\u003e (Linnaeus, 1758)","normalized":"Velutina halioides (Linnaeus 1758)","canonical":{"stemmed":"Velutina halioid","simple":"Velutina halioides","full":"Velutina halioides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Velutina","species":"halioides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"words":[{"verbatim":"Velutina","normalized":"Velutina","wordType":"GENUS","start":0,"end":8},{"verbatim":"halioides","normalized":"halioides","wordType":"SPECIES","start":9,"end":18},{"verbatim":"Linnaeus","normalized":"Linnaeus","wordType":"AUTHOR_WORD","start":20,"end":28},{"verbatim":"1758","normalized":"1758","wordType":"YEAR","start":30,"end":34}],"id":"653bbe42-aef4-5847-add4-8c7f8a4d1f9b","parserVersion":"test_version"}
```

Name: Quadrella steyermarkii (Standl.) Iltis &amp; Cornejo

Canonical: Quadrella steyermarkii

Authorship: (Standl.) Iltis & Cornejo

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Quadrella steyermarkii (Standl.) Iltis \u0026amp; Cornejo","normalized":"Quadrella steyermarkii (Standl.) Iltis \u0026 Cornejo","canonical":{"stemmed":"Quadrella steyermark","simple":"Quadrella steyermarkii","full":"Quadrella steyermarkii"},"cardinality":2,"authorship":{"verbatim":"(Standl.) Iltis \u0026 Cornejo","normalized":"(Standl.) Iltis \u0026 Cornejo","authors":["Standl.","Iltis","Cornejo"],"originalAuth":{"authors":["Standl."]},"combinationAuth":{"authors":["Iltis","Cornejo"]}},"details":{"species":{"genus":"Quadrella","species":"steyermarkii","authorship":{"verbatim":"(Standl.) Iltis \u0026 Cornejo","normalized":"(Standl.) Iltis \u0026 Cornejo","authors":["Standl.","Iltis","Cornejo"],"originalAuth":{"authors":["Standl."]},"combinationAuth":{"authors":["Iltis","Cornejo"]}}}},"words":[{"verbatim":"Quadrella","normalized":"Quadrella","wordType":"GENUS","start":0,"end":9},{"verbatim":"steyermarkii","normalized":"steyermarkii","wordType":"SPECIES","start":10,"end":22},{"verbatim":"Standl.","normalized":"Standl.","wordType":"AUTHOR_WORD","start":24,"end":31},{"verbatim":"Iltis","normalized":"Iltis","wordType":"AUTHOR_WORD","start":33,"end":38},{"verbatim":"Cornejo","normalized":"Cornejo","wordType":"AUTHOR_WORD","start":41,"end":48}],"id":"fbd1b4fe-f8ed-5390-9cb1-e0f798691b1e","parserVersion":"test_version"}
```

Name: Torymus bangalorensis (Mani &amp; Kurian, 1953)

Canonical: Torymus bangalorensis

Authorship: (Mani & Kurian 1953)

```json
{"parsed":true,"quality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Torymus bangalorensis (Mani \u0026amp; Kurian, 1953)","normalized":"Torymus bangalorensis (Mani \u0026 Kurian 1953)","canonical":{"stemmed":"Torymus bangalorens","simple":"Torymus bangalorensis","full":"Torymus bangalorensis"},"cardinality":2,"authorship":{"verbatim":"(Mani \u0026 Kurian, 1953)","normalized":"(Mani \u0026 Kurian 1953)","year":"1953","authors":["Mani","Kurian"],"originalAuth":{"authors":["Mani","Kurian"],"year":{"year":"1953"}}},"details":{"species":{"genus":"Torymus","species":"bangalorensis","authorship":{"verbatim":"(Mani \u0026 Kurian, 1953)","normalized":"(Mani \u0026 Kurian 1953)","year":"1953","authors":["Mani","Kurian"],"originalAuth":{"authors":["Mani","Kurian"],"year":{"year":"1953"}}}}},"words":[{"verbatim":"Torymus","normalized":"Torymus","wordType":"GENUS","start":0,"end":7},{"verbatim":"bangalorensis","normalized":"bangalorensis","wordType":"SPECIES","start":8,"end":21},{"verbatim":"Mani","normalized":"Mani","wordType":"AUTHOR_WORD","start":23,"end":27},{"verbatim":"Kurian","normalized":"Kurian","wordType":"AUTHOR_WORD","start":30,"end":36},{"verbatim":"1953","normalized":"1953","wordType":"YEAR","start":38,"end":42}],"id":"8131ebda-dce6-5aaf-97ae-2370fe8e77d7","parserVersion":"test_version"}
```

### Underscores instead of spaces

Name: Oxalis_barrelieri

Canonical: Oxalis barrelieri

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Oxalis_barrelieri","normalized":"Oxalis barrelieri","canonical":{"stemmed":"Oxalis barrelier","simple":"Oxalis barrelieri","full":"Oxalis barrelieri"},"cardinality":2,"details":{"species":{"genus":"Oxalis","species":"barrelieri"}},"words":[{"verbatim":"Oxalis","normalized":"Oxalis","wordType":"GENUS","start":0,"end":6},{"verbatim":"barrelieri","normalized":"barrelieri","wordType":"SPECIES","start":7,"end":17}],"id":"ad546700-9cae-50d3-9eaf-6adcbbb67bae","parserVersion":"test_version"}
```

Name:   Oxalis_barrelieri ined.?

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"  Oxalis_barrelieri ined.?","cardinality":0,"id":"c065444b-dbdd-5f29-96f9-629f49469abd","parserVersion":"test_version"}
```

Name: Pseudocercospora__dendrobii

Canonical: Pseudocercospora dendrobii

Authorship:

```json
{"parsed":true,"quality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Pseudocercospora__dendrobii","normalized":"Pseudocercospora dendrobii","canonical":{"stemmed":"Pseudocercospora dendrob","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii"}},"words":[{"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","wordType":"GENUS","start":0,"end":16},{"verbatim":"dendrobii","normalized":"dendrobii","wordType":"SPECIES","start":18,"end":27}],"id":"ae8a4688-2b2a-5974-81bf-1962838a9cbe","parserVersion":"test_version"}
```

Name:   Oxalis_barrelieri

Canonical:

Authorship:

```json
{"parsed":false,"quality":0,"verbatim":"  Oxalis_barrelieri","cardinality":0,"id":"1c4bb48b-d134-54c8-bac1-6771d1f4c9c6","parserVersion":"test_version"}
```

Name: Oxalis barrelieri XXZ_21243

Canonical: Oxalis barrelieri

Authorship:

```json
{"parsed":true,"quality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Oxalis barrelieri XXZ_21243","normalized":"Oxalis barrelieri","canonical":{"stemmed":"Oxalis barrelier","simple":"Oxalis barrelieri","full":"Oxalis barrelieri"},"cardinality":2,"tail":" XXZ_21243","details":{"species":{"genus":"Oxalis","species":"barrelieri"}},"words":[{"verbatim":"Oxalis","normalized":"Oxalis","wordType":"GENUS","start":0,"end":6},{"verbatim":"barrelieri","normalized":"barrelieri","wordType":"SPECIES","start":7,"end":17}],"id":"8a722b76-cf2f-51d1-b60e-7f9236ddd189","parserVersion":"test_version"}
```
