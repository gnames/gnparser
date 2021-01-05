# Global Names Parser Test


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
  * [Binomials with several authours](#binomials-with-several-authours)
  * [Binomials with several authors and a year](#binomials-with-several-authors-and-a-year)
  * [Binomials with basionym and combination authors](#binomials-with-basionym-and-combination-authors)
  * [Infraspecies without rank (ICZN)](#infraspecies-without-rank-iczn)
  * [Legacy ICZN names with rank](#legacy-iczn-names-with-rank)
  * [Infraspecies with rank (ICN)](#infraspecies-with-rank-icn)
  * [Infraspecies multiple (ICN)](#infraspecies-multiple-icn)
  * [Infraspecies with greek letters (ICN)](#infraspecies-with-greek-letters-icn)
  * [Hybrids with notho- ranks](#hybrids-with-notho--ranks)
  * [Named hybrids](#named-hybrids)
  * [Hybrid formulae](#hybrid-formulae)
  * [Genus with hyphen (allowed by ICN)](#genus-with-hyphen-allowed-by-icn)
  * [Misspeled name](#misspeled-name)
  * [A 'basionym' author in parenthesis (basionym is an ICN term)](#a-basionym-author-in-parenthesis-basionym-is-an-icn-term)
  * [Infrageneric epithets (ICZN)](#infrageneric-epithets-iczn)
  * [Names with multiple dashes in specific epithet](#names-with-multiple-dashes-in-specific-epithet)
  * [Genus with question mark](#genus-with-question-mark)
  * [Epithets starting with authors' prefixes (de, di, la, von etc.)](#epithets-starting-with-authors-prefixes-de-di-la-von-etc)
  * [Authorship missing one parenthesis](#authorship-missing-one-parenthesis)
  * [Unknown authorship](#unknown-authorship)
  * [Treating apud (with)](#treating-apud-with)
  * [Names with ex authors (we follow ICZN convention)](#names-with-ex-authors-we-follow-iczn-convention)
  * [Empty spaces](#empty-spaces)
  * [Names with a dash](#names-with-a-dash)
  * [Authorship with filius (son of)](#authorship-with-filius-son-of)
  * [Names with emend (rectified by) authorship](#names-with-emend-rectified-by-authorship)
  * ["Tail" annotations](#tail-annotations)
  * [Abbreviated words after a name](#abbreviated-words-after-a-name)
  * [Epithets starting with numeric value (not allowed anymore)](#epithets-starting-with-numeric-value-not-allowed-anymore)
  * [Non-ASCII UTF-8 characters in a name](#non-ascii-utf-8-characters-in-a-name)
  * [Epithets with an apostrophe](#epithets-with-an-apostrophe)
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
  * [Authoship in upper case](#authoship-in-upper-case)
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
  * [Names with unparsed "tail" at the end](#names-with-unparsed-tail-at-the-end)
  * [Discard apostrophes at the start and end of words](#discard-apostrophes-at-the-start-and-end-of-words)
  * [Discard apostrophe with dash (rare, needs further investigation)](#discard-apostrophe-with-dash-rare-needs-further-investigation)
  * [Possible canonical](#possible-canonical)
  * [Treating `& al.` as `et al.`](#treating--al-as-et-al)
  * [Authors do not start with apostrophe](#authors-do-not-start-with-apostrophe)
  * [Epithets do not start or end with a dash](#epithets-do-not-start-or-end-with-a-dash)
  * [names that contain "of"](#names-that-contain-of)
  * [Names that contain "cv" (cultivar)](#names-that-contain-cv-cultivar)
  * ["Open taxonomy" with ranks unfinished](#open-taxonomy-with-ranks-unfinished)
  * [Ignoring sensu sec](#ignoring-sensu-sec)
  * [Unparseable hort. annotations](#unparseable-hort-annotations)
  * [Removing nomenclatural annotations](#removing-nomenclatural-annotations)
  * [Misc annotations](#misc-annotations)
  * [Horticultural annotation](#horticultural-annotation)
  * [Not parsed OCR errors to get better precision/recall ratio](#not-parsed-ocr-errors-to-get-better-precisionrecall-ratio)
  * [No parsing -- Genera abbreviated to 3 letters (too rare)](#no-parsing----genera-abbreviated-to-3-letters-too-rare)
  * [No parsing -- incertae sedis](#no-parsing----incertae-sedis)
  * [No parsing -- bacterium, Candidatus](#no-parsing----bacterium-candidatus)
  * [No parsing -- 'Not', 'None', 'Unidentified'  phrases](#no-parsing----not-none-unidentified--phrases)
  * [No parsing -- genus with apostrophe](#no-parsing----genus-with-apostrophe)
  * [No parsing -- CamelCase 'genus' word](#no-parsing----camelcase-genus-word)
  * [No parsing -- phytoplasma](#no-parsing----phytoplasma)
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
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora","normalized":"Pseudocercospora","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pseudocercospora"}},"pos":[{"wordType":"uninomial","start":0,"end":16}],"id":"9c1167ca-79e7-53de-b4c3-fcdb68410527","parserVersion":"test_version"}
```

### Uninomials with authorship

Name: Pseudocercospora Speg.

Canonical: Pseudocercospora

Authorship: Speg.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora Speg.","normalized":"Pseudocercospora Speg.","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg.","normalized":"Speg.","authors":["Speg."],"originalAuth":{"authors":["Speg."]}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg.","normalized":"Speg.","authors":["Speg."],"originalAuth":{"authors":["Speg."]}}}},"pos":[{"wordType":"uninomial","start":0,"end":16},{"wordType":"authorWord","start":17,"end":22}],"id":"ccc7780b-c68b-53c6-9166-6b2d4902923e","parserVersion":"test_version"}
```

Name: Döringina Ihering 1929 (synonym)

Canonical: Doeringina

Authorship: Ihering 1929

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Döringina Ihering 1929 (synonym)","normalized":"Doeringina Ihering 1929","canonical":{"stemmed":"Doeringina","simple":"Doeringina","full":"Doeringina"},"cardinality":1,"authorship":{"verbatim":"Ihering 1929","normalized":"Ihering 1929","year":"1929","authors":["Ihering"],"originalAuth":{"authors":["Ihering"],"year":{"year":"1929"}}},"tail":" (synonym)","details":{"uninomial":{"uninomial":"Doeringina","authorship":{"verbatim":"Ihering 1929","normalized":"Ihering 1929","year":"1929","authors":["Ihering"],"originalAuth":{"authors":["Ihering"],"year":{"year":"1929"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":9},{"wordType":"authorWord","start":10,"end":17},{"wordType":"year","start":18,"end":22}],"id":"95eb9081-5fe5-5497-be3d-ef0ce65a472c","parserVersion":"test_version"}
```

Name: Pseudocercospora Speg., Francis Jack.-Drake.

Canonical: Pseudocercospora

Authorship: Speg. & Francis Jack.-Drake.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora Speg., Francis Jack.-Drake.","normalized":"Pseudocercospora Speg. \u0026 Francis Jack.-Drake.","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg., Francis Jack.-Drake.","normalized":"Speg. \u0026 Francis Jack.-Drake.","authors":["Speg.","Francis Jack.-Drake."],"originalAuth":{"authors":["Speg.","Francis Jack.-Drake."]}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg., Francis Jack.-Drake.","normalized":"Speg. \u0026 Francis Jack.-Drake.","authors":["Speg.","Francis Jack.-Drake."],"originalAuth":{"authors":["Speg.","Francis Jack.-Drake."]}}}},"pos":[{"wordType":"uninomial","start":0,"end":16},{"wordType":"authorWord","start":17,"end":22},{"wordType":"authorWord","start":24,"end":31},{"wordType":"authorWord","start":32,"end":44}],"id":"25b015c7-a099-5bf6-91a9-cc8fde31f388","parserVersion":"test_version"}
```

Name: Aaaba de Laubenfels, 1936

Canonical: Aaaba

Authorship: de Laubenfels 1936

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aaaba de Laubenfels, 1936","normalized":"Aaaba de Laubenfels 1936","canonical":{"stemmed":"Aaaba","simple":"Aaaba","full":"Aaaba"},"cardinality":1,"authorship":{"verbatim":"de Laubenfels, 1936","normalized":"de Laubenfels 1936","year":"1936","authors":["de Laubenfels"],"originalAuth":{"authors":["de Laubenfels"],"year":{"year":"1936"}}},"details":{"uninomial":{"uninomial":"Aaaba","authorship":{"verbatim":"de Laubenfels, 1936","normalized":"de Laubenfels 1936","year":"1936","authors":["de Laubenfels"],"originalAuth":{"authors":["de Laubenfels"],"year":{"year":"1936"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":5},{"wordType":"authorWord","start":6,"end":8},{"wordType":"authorWord","start":9,"end":19},{"wordType":"year","start":21,"end":25}],"id":"abead069-293d-5299-badd-c10c0f5545fb","parserVersion":"test_version"}
```

Name: Abbottia F. von Mueller, 1875

Canonical: Abbottia

Authorship: F. von Mueller 1875

```json
{"parsed":true,"parseQuality":1,"verbatim":"Abbottia F. von Mueller, 1875","normalized":"Abbottia F. von Mueller 1875","canonical":{"stemmed":"Abbottia","simple":"Abbottia","full":"Abbottia"},"cardinality":1,"authorship":{"verbatim":"F. von Mueller, 1875","normalized":"F. von Mueller 1875","year":"1875","authors":["F. von Mueller"],"originalAuth":{"authors":["F. von Mueller"],"year":{"year":"1875"}}},"details":{"uninomial":{"uninomial":"Abbottia","authorship":{"verbatim":"F. von Mueller, 1875","normalized":"F. von Mueller 1875","year":"1875","authors":["F. von Mueller"],"originalAuth":{"authors":["F. von Mueller"],"year":{"year":"1875"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":11},{"wordType":"authorWord","start":12,"end":15},{"wordType":"authorWord","start":16,"end":23},{"wordType":"year","start":25,"end":29}],"id":"34738de5-0112-56f0-85f2-0f4e815161b5","parserVersion":"test_version"}
```

Name: Abella von Heyden, 1826

Canonical: Abella

Authorship: von Heyden 1826

```json
{"parsed":true,"parseQuality":1,"verbatim":"Abella von Heyden, 1826","normalized":"Abella von Heyden 1826","canonical":{"stemmed":"Abella","simple":"Abella","full":"Abella"},"cardinality":1,"authorship":{"verbatim":"von Heyden, 1826","normalized":"von Heyden 1826","year":"1826","authors":["von Heyden"],"originalAuth":{"authors":["von Heyden"],"year":{"year":"1826"}}},"details":{"uninomial":{"uninomial":"Abella","authorship":{"verbatim":"von Heyden, 1826","normalized":"von Heyden 1826","year":"1826","authors":["von Heyden"],"originalAuth":{"authors":["von Heyden"],"year":{"year":"1826"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":6},{"wordType":"authorWord","start":7,"end":10},{"wordType":"authorWord","start":11,"end":17},{"wordType":"year","start":19,"end":23}],"id":"7dc5b624-1232-5072-bc4c-8eebde6c48b2","parserVersion":"test_version"}
```

Name: Micropleura v Linstow 1906

Canonical: Micropleura

Authorship: v Linstow 1906

```json
{"parsed":true,"parseQuality":1,"verbatim":"Micropleura v Linstow 1906","normalized":"Micropleura v Linstow 1906","canonical":{"stemmed":"Micropleura","simple":"Micropleura","full":"Micropleura"},"cardinality":1,"authorship":{"verbatim":"v Linstow 1906","normalized":"v Linstow 1906","year":"1906","authors":["v Linstow"],"originalAuth":{"authors":["v Linstow"],"year":{"year":"1906"}}},"details":{"uninomial":{"uninomial":"Micropleura","authorship":{"verbatim":"v Linstow 1906","normalized":"v Linstow 1906","year":"1906","authors":["v Linstow"],"originalAuth":{"authors":["v Linstow"],"year":{"year":"1906"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":13},{"wordType":"authorWord","start":14,"end":21},{"wordType":"year","start":22,"end":26}],"id":"94f99223-2631-52a9-9497-a29452387980","parserVersion":"test_version"}
```

Name: Pseudocercospora Speg. 1910

Canonical: Pseudocercospora

Authorship: Speg. 1910

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora Speg. 1910","normalized":"Pseudocercospora Speg. 1910","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Speg. 1910","normalized":"Speg. 1910","year":"1910","authors":["Speg."],"originalAuth":{"authors":["Speg."],"year":{"year":"1910"}}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Speg. 1910","normalized":"Speg. 1910","year":"1910","authors":["Speg."],"originalAuth":{"authors":["Speg."],"year":{"year":"1910"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":16},{"wordType":"authorWord","start":17,"end":22},{"wordType":"year","start":23,"end":27}],"id":"eac97817-869a-5400-8b1e-0a125876189d","parserVersion":"test_version"}
```

Name: Pseudocercospora Spegazzini, 1910

Canonical: Pseudocercospora

Authorship: Spegazzini 1910

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora Spegazzini, 1910","normalized":"Pseudocercospora Spegazzini 1910","canonical":{"stemmed":"Pseudocercospora","simple":"Pseudocercospora","full":"Pseudocercospora"},"cardinality":1,"authorship":{"verbatim":"Spegazzini, 1910","normalized":"Spegazzini 1910","year":"1910","authors":["Spegazzini"],"originalAuth":{"authors":["Spegazzini"],"year":{"year":"1910"}}},"details":{"uninomial":{"uninomial":"Pseudocercospora","authorship":{"verbatim":"Spegazzini, 1910","normalized":"Spegazzini 1910","year":"1910","authors":["Spegazzini"],"originalAuth":{"authors":["Spegazzini"],"year":{"year":"1910"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":16},{"wordType":"authorWord","start":17,"end":27},{"wordType":"year","start":29,"end":33}],"id":"6cc2922a-1f1d-5a40-90a7-b155fd16b233","parserVersion":"test_version"}
```

Name: Rhynchonellidae d'Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rhynchonellidae d'Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d'Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d'Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":25},{"wordType":"year","start":26,"end":30}],"id":"f3b90050-32f2-5009-ae9d-705fc58e45c4","parserVersion":"test_version"}
```

Name: Rhynchonellidae d‘Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Rhynchonellidae d‘Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d‘Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d‘Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":25},{"wordType":"year","start":26,"end":30}],"id":"8a72add4-b276-5a92-ad30-a4c8bc03598a","parserVersion":"test_version"}
```

Name: Rhynchonellidae d’Orbigny 1847

Canonical: Rhynchonellidae

Authorship: d'Orbigny 1847

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Rhynchonellidae d’Orbigny 1847","normalized":"Rhynchonellidae d'Orbigny 1847","canonical":{"stemmed":"Rhynchonellidae","simple":"Rhynchonellidae","full":"Rhynchonellidae"},"cardinality":1,"authorship":{"verbatim":"d’Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}},"details":{"uninomial":{"uninomial":"Rhynchonellidae","authorship":{"verbatim":"d’Orbigny 1847","normalized":"d'Orbigny 1847","year":"1847","authors":["d'Orbigny"],"originalAuth":{"authors":["d'Orbigny"],"year":{"year":"1847"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":25},{"wordType":"year","start":26,"end":30}],"id":"cc9b39b8-b4d0-5e8e-9ffe-866454d3e49a","parserVersion":"test_version"}
```

Name: Ataladoris Iredale & O'Donoghue 1923

Canonical: Ataladoris

Authorship: Iredale & O'Donoghue 1923

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ataladoris Iredale \u0026 O'Donoghue 1923","normalized":"Ataladoris Iredale \u0026 O'Donoghue 1923","canonical":{"stemmed":"Ataladoris","simple":"Ataladoris","full":"Ataladoris"},"cardinality":1,"authorship":{"verbatim":"Iredale \u0026 O'Donoghue 1923","normalized":"Iredale \u0026 O'Donoghue 1923","year":"1923","authors":["Iredale","O'Donoghue"],"originalAuth":{"authors":["Iredale","O'Donoghue"],"year":{"year":"1923"}}},"details":{"uninomial":{"uninomial":"Ataladoris","authorship":{"verbatim":"Iredale \u0026 O'Donoghue 1923","normalized":"Iredale \u0026 O'Donoghue 1923","year":"1923","authors":["Iredale","O'Donoghue"],"originalAuth":{"authors":["Iredale","O'Donoghue"],"year":{"year":"1923"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":10},{"wordType":"authorWord","start":11,"end":18},{"wordType":"authorWord","start":21,"end":31},{"wordType":"year","start":32,"end":36}],"id":"dbb90380-0552-5237-82ef-8a8b07e42049","parserVersion":"test_version"}
```

Name: Anteplana le Renard 1995

Canonical: Anteplana

Authorship: le Renard 1995

```json
{"parsed":true,"parseQuality":1,"verbatim":"Anteplana le Renard 1995","normalized":"Anteplana le Renard 1995","canonical":{"stemmed":"Anteplana","simple":"Anteplana","full":"Anteplana"},"cardinality":1,"authorship":{"verbatim":"le Renard 1995","normalized":"le Renard 1995","year":"1995","authors":["le Renard"],"originalAuth":{"authors":["le Renard"],"year":{"year":"1995"}}},"details":{"uninomial":{"uninomial":"Anteplana","authorship":{"verbatim":"le Renard 1995","normalized":"le Renard 1995","year":"1995","authors":["le Renard"],"originalAuth":{"authors":["le Renard"],"year":{"year":"1995"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":9},{"wordType":"authorWord","start":10,"end":12},{"wordType":"authorWord","start":13,"end":19},{"wordType":"year","start":20,"end":24}],"id":"6920744c-27e9-546f-96d9-c8859544ef78","parserVersion":"test_version"}
```

Name: Candinia le Renard, Sabelli & Taviani 1996

Canonical: Candinia

Authorship: le Renard, Sabelli & Taviani 1996

```json
{"parsed":true,"parseQuality":1,"verbatim":"Candinia le Renard, Sabelli \u0026 Taviani 1996","normalized":"Candinia le Renard, Sabelli \u0026 Taviani 1996","canonical":{"stemmed":"Candinia","simple":"Candinia","full":"Candinia"},"cardinality":1,"authorship":{"verbatim":"le Renard, Sabelli \u0026 Taviani 1996","normalized":"le Renard, Sabelli \u0026 Taviani 1996","year":"1996","authors":["le Renard","Sabelli","Taviani"],"originalAuth":{"authors":["le Renard","Sabelli","Taviani"],"year":{"year":"1996"}}},"details":{"uninomial":{"uninomial":"Candinia","authorship":{"verbatim":"le Renard, Sabelli \u0026 Taviani 1996","normalized":"le Renard, Sabelli \u0026 Taviani 1996","year":"1996","authors":["le Renard","Sabelli","Taviani"],"originalAuth":{"authors":["le Renard","Sabelli","Taviani"],"year":{"year":"1996"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":11},{"wordType":"authorWord","start":12,"end":18},{"wordType":"authorWord","start":20,"end":27},{"wordType":"authorWord","start":30,"end":37},{"wordType":"year","start":38,"end":42}],"id":"2a92b7b1-4da8-5571-98de-9cd225526081","parserVersion":"test_version"}
```

Name: Polypodium le Sourdianum Fourn.

Canonical: Polypodium

Authorship: le Sourdianum Fourn.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Polypodium le Sourdianum Fourn.","normalized":"Polypodium le Sourdianum Fourn.","canonical":{"stemmed":"Polypodium","simple":"Polypodium","full":"Polypodium"},"cardinality":1,"authorship":{"verbatim":"le Sourdianum Fourn.","normalized":"le Sourdianum Fourn.","authors":["le Sourdianum Fourn."],"originalAuth":{"authors":["le Sourdianum Fourn."]}},"details":{"uninomial":{"uninomial":"Polypodium","authorship":{"verbatim":"le Sourdianum Fourn.","normalized":"le Sourdianum Fourn.","authors":["le Sourdianum Fourn."],"originalAuth":{"authors":["le Sourdianum Fourn."]}}}},"pos":[{"wordType":"uninomial","start":0,"end":10},{"wordType":"authorWord","start":11,"end":13},{"wordType":"authorWord","start":14,"end":24},{"wordType":"authorWord","start":25,"end":31}],"id":"ea72f0d9-2f8a-5ba0-95c7-986075eda321","parserVersion":"test_version"}
```

### Two-letter genus names (legacy genera, not allowed anymore)

Name: Ca Dyar 1914

Canonical: Ca

Authorship: Dyar 1914

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ca Dyar 1914","normalized":"Ca Dyar 1914","canonical":{"stemmed":"Ca","simple":"Ca","full":"Ca"},"cardinality":1,"authorship":{"verbatim":"Dyar 1914","normalized":"Dyar 1914","year":"1914","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1914"}}},"details":{"uninomial":{"uninomial":"Ca","authorship":{"verbatim":"Dyar 1914","normalized":"Dyar 1914","year":"1914","authors":["Dyar"],"originalAuth":{"authors":["Dyar"],"year":{"year":"1914"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":7},{"wordType":"year","start":8,"end":12}],"id":"ccb4663f-3d9a-5447-ab28-13e453738075","parserVersion":"test_version"}
```

Name: Ea Distant 1911

Canonical: Ea

Authorship: Distant 1911

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ea Distant 1911","normalized":"Ea Distant 1911","canonical":{"stemmed":"Ea","simple":"Ea","full":"Ea"},"cardinality":1,"authorship":{"verbatim":"Distant 1911","normalized":"Distant 1911","year":"1911","authors":["Distant"],"originalAuth":{"authors":["Distant"],"year":{"year":"1911"}}},"details":{"uninomial":{"uninomial":"Ea","authorship":{"verbatim":"Distant 1911","normalized":"Distant 1911","year":"1911","authors":["Distant"],"originalAuth":{"authors":["Distant"],"year":{"year":"1911"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10},{"wordType":"year","start":11,"end":15}],"id":"c5a5643f-452f-5c51-91eb-42789ed6f3a4","parserVersion":"test_version"}
```

Name: Ge Nicéville 1895

Canonical: Ge

Authorship: Nicéville 1895

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ge Nicéville 1895","normalized":"Ge Nicéville 1895","canonical":{"stemmed":"Ge","simple":"Ge","full":"Ge"},"cardinality":1,"authorship":{"verbatim":"Nicéville 1895","normalized":"Nicéville 1895","year":"1895","authors":["Nicéville"],"originalAuth":{"authors":["Nicéville"],"year":{"year":"1895"}}},"details":{"uninomial":{"uninomial":"Ge","authorship":{"verbatim":"Nicéville 1895","normalized":"Nicéville 1895","year":"1895","authors":["Nicéville"],"originalAuth":{"authors":["Nicéville"],"year":{"year":"1895"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":12},{"wordType":"year","start":13,"end":17}],"id":"ba4f0f90-1df5-5054-a17b-15938a942d88","parserVersion":"test_version"}
```

Name: Ia Thomas 1902

Canonical: Ia

Authorship: Thomas 1902

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ia Thomas 1902","normalized":"Ia Thomas 1902","canonical":{"stemmed":"Ia","simple":"Ia","full":"Ia"},"cardinality":1,"authorship":{"verbatim":"Thomas 1902","normalized":"Thomas 1902","year":"1902","authors":["Thomas"],"originalAuth":{"authors":["Thomas"],"year":{"year":"1902"}}},"details":{"uninomial":{"uninomial":"Ia","authorship":{"verbatim":"Thomas 1902","normalized":"Thomas 1902","year":"1902","authors":["Thomas"],"originalAuth":{"authors":["Thomas"],"year":{"year":"1902"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":9},{"wordType":"year","start":10,"end":14}],"id":"9826997c-1d52-5de2-8b7b-facdc9fb73f2","parserVersion":"test_version"}
```

Name: Io Lea 1831

Canonical: Io

Authorship: Lea 1831

```json
{"parsed":true,"parseQuality":1,"verbatim":"Io Lea 1831","normalized":"Io Lea 1831","canonical":{"stemmed":"Io","simple":"Io","full":"Io"},"cardinality":1,"authorship":{"verbatim":"Lea 1831","normalized":"Lea 1831","year":"1831","authors":["Lea"],"originalAuth":{"authors":["Lea"],"year":{"year":"1831"}}},"details":{"uninomial":{"uninomial":"Io","authorship":{"verbatim":"Lea 1831","normalized":"Lea 1831","year":"1831","authors":["Lea"],"originalAuth":{"authors":["Lea"],"year":{"year":"1831"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":6},{"wordType":"year","start":7,"end":11}],"id":"3cc533a5-4f2c-5aec-ba30-85a27548aa95","parserVersion":"test_version"}
```

Name: Io Blanchard 1852

Canonical: Io

Authorship: Blanchard 1852

```json
{"parsed":true,"parseQuality":1,"verbatim":"Io Blanchard 1852","normalized":"Io Blanchard 1852","canonical":{"stemmed":"Io","simple":"Io","full":"Io"},"cardinality":1,"authorship":{"verbatim":"Blanchard 1852","normalized":"Blanchard 1852","year":"1852","authors":["Blanchard"],"originalAuth":{"authors":["Blanchard"],"year":{"year":"1852"}}},"details":{"uninomial":{"uninomial":"Io","authorship":{"verbatim":"Blanchard 1852","normalized":"Blanchard 1852","year":"1852","authors":["Blanchard"],"originalAuth":{"authors":["Blanchard"],"year":{"year":"1852"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":12},{"wordType":"year","start":13,"end":17}],"id":"4de7e503-a5a5-5309-bc6c-cbaf90a9199b","parserVersion":"test_version"}
```

Name: Ix Bergroth 1916

Canonical: Ix

Authorship: Bergroth 1916

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ix Bergroth 1916","normalized":"Ix Bergroth 1916","canonical":{"stemmed":"Ix","simple":"Ix","full":"Ix"},"cardinality":1,"authorship":{"verbatim":"Bergroth 1916","normalized":"Bergroth 1916","year":"1916","authors":["Bergroth"],"originalAuth":{"authors":["Bergroth"],"year":{"year":"1916"}}},"details":{"uninomial":{"uninomial":"Ix","authorship":{"verbatim":"Bergroth 1916","normalized":"Bergroth 1916","year":"1916","authors":["Bergroth"],"originalAuth":{"authors":["Bergroth"],"year":{"year":"1916"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":11},{"wordType":"year","start":12,"end":16}],"id":"981228e8-45fe-5b7b-ab78-4793cae51602","parserVersion":"test_version"}
```

Name: Lo Seale 1906

Canonical: Lo

Authorship: Seale 1906

```json
{"parsed":true,"parseQuality":1,"verbatim":"Lo Seale 1906","normalized":"Lo Seale 1906","canonical":{"stemmed":"Lo","simple":"Lo","full":"Lo"},"cardinality":1,"authorship":{"verbatim":"Seale 1906","normalized":"Seale 1906","year":"1906","authors":["Seale"],"originalAuth":{"authors":["Seale"],"year":{"year":"1906"}}},"details":{"uninomial":{"uninomial":"Lo","authorship":{"verbatim":"Seale 1906","normalized":"Seale 1906","year":"1906","authors":["Seale"],"originalAuth":{"authors":["Seale"],"year":{"year":"1906"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":8},{"wordType":"year","start":9,"end":13}],"id":"8d9cb022-3458-5473-aa5a-91da319d5d78","parserVersion":"test_version"}
```

Name: Oa Girault 1929

Canonical: Oa

Authorship: Girault 1929

```json
{"parsed":true,"parseQuality":1,"verbatim":"Oa Girault 1929","normalized":"Oa Girault 1929","canonical":{"stemmed":"Oa","simple":"Oa","full":"Oa"},"cardinality":1,"authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}},"details":{"uninomial":{"uninomial":"Oa","authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10},{"wordType":"year","start":11,"end":15}],"id":"14647a9c-70c8-55a8-b2a7-1fc47c39732b","parserVersion":"test_version"}
```

Name: Ra Whitley 1931

Canonical: Ra

Authorship: Whitley 1931

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ra Whitley 1931","normalized":"Ra Whitley 1931","canonical":{"stemmed":"Ra","simple":"Ra","full":"Ra"},"cardinality":1,"authorship":{"verbatim":"Whitley 1931","normalized":"Whitley 1931","year":"1931","authors":["Whitley"],"originalAuth":{"authors":["Whitley"],"year":{"year":"1931"}}},"details":{"uninomial":{"uninomial":"Ra","authorship":{"verbatim":"Whitley 1931","normalized":"Whitley 1931","year":"1931","authors":["Whitley"],"originalAuth":{"authors":["Whitley"],"year":{"year":"1931"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10},{"wordType":"year","start":11,"end":15}],"id":"72b5b436-6381-5939-b8d1-7f04bb2a82bb","parserVersion":"test_version"}
```

Name: Ty Bory de St. Vincent 1827

Canonical: Ty

Authorship: Bory de St. Vincent 1827

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ty Bory de St. Vincent 1827","normalized":"Ty Bory de St. Vincent 1827","canonical":{"stemmed":"Ty","simple":"Ty","full":"Ty"},"cardinality":1,"authorship":{"verbatim":"Bory de St. Vincent 1827","normalized":"Bory de St. Vincent 1827","year":"1827","authors":["Bory de St. Vincent"],"originalAuth":{"authors":["Bory de St. Vincent"],"year":{"year":"1827"}}},"details":{"uninomial":{"uninomial":"Ty","authorship":{"verbatim":"Bory de St. Vincent 1827","normalized":"Bory de St. Vincent 1827","year":"1827","authors":["Bory de St. Vincent"],"originalAuth":{"authors":["Bory de St. Vincent"],"year":{"year":"1827"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":7},{"wordType":"authorWord","start":8,"end":10},{"wordType":"authorWord","start":11,"end":14},{"wordType":"authorWord","start":15,"end":22},{"wordType":"year","start":23,"end":27}],"id":"1d05b120-8f75-58ab-bdf7-c181fdf1bc3c","parserVersion":"test_version"}
```

Name: Ua Girault 1929

Canonical: Ua

Authorship: Girault 1929

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ua Girault 1929","normalized":"Ua Girault 1929","canonical":{"stemmed":"Ua","simple":"Ua","full":"Ua"},"cardinality":1,"authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}},"details":{"uninomial":{"uninomial":"Ua","authorship":{"verbatim":"Girault 1929","normalized":"Girault 1929","year":"1929","authors":["Girault"],"originalAuth":{"authors":["Girault"],"year":{"year":"1929"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10},{"wordType":"year","start":11,"end":15}],"id":"aee3fe77-1797-5172-82f1-5ee233108c15","parserVersion":"test_version"}
```

Name: Aa Baker 1940

Canonical: Aa

Authorship: Baker 1940

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aa Baker 1940","normalized":"Aa Baker 1940","canonical":{"stemmed":"Aa","simple":"Aa","full":"Aa"},"cardinality":1,"authorship":{"verbatim":"Baker 1940","normalized":"Baker 1940","year":"1940","authors":["Baker"],"originalAuth":{"authors":["Baker"],"year":{"year":"1940"}}},"details":{"uninomial":{"uninomial":"Aa","authorship":{"verbatim":"Baker 1940","normalized":"Baker 1940","year":"1940","authors":["Baker"],"originalAuth":{"authors":["Baker"],"year":{"year":"1940"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":8},{"wordType":"year","start":9,"end":13}],"id":"101d126d-c14a-5043-a1d8-72bc6a9f4dcf","parserVersion":"test_version"}
```

Name: Ja Uéno 1955

Canonical: Ja

Authorship: Uéno 1955

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ja Uéno 1955","normalized":"Ja Uéno 1955","canonical":{"stemmed":"Ja","simple":"Ja","full":"Ja"},"cardinality":1,"authorship":{"verbatim":"Uéno 1955","normalized":"Uéno 1955","year":"1955","authors":["Uéno"],"originalAuth":{"authors":["Uéno"],"year":{"year":"1955"}}},"details":{"uninomial":{"uninomial":"Ja","authorship":{"verbatim":"Uéno 1955","normalized":"Uéno 1955","year":"1955","authors":["Uéno"],"originalAuth":{"authors":["Uéno"],"year":{"year":"1955"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":7},{"wordType":"year","start":8,"end":12}],"id":"45f6eba8-1063-590d-bc4a-9f9ffdef4a10","parserVersion":"test_version"}
```

Name: Zu Walters & Fitch 1960

Canonical: Zu

Authorship: Walters & Fitch 1960

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zu Walters \u0026 Fitch 1960","normalized":"Zu Walters \u0026 Fitch 1960","canonical":{"stemmed":"Zu","simple":"Zu","full":"Zu"},"cardinality":1,"authorship":{"verbatim":"Walters \u0026 Fitch 1960","normalized":"Walters \u0026 Fitch 1960","year":"1960","authors":["Walters","Fitch"],"originalAuth":{"authors":["Walters","Fitch"],"year":{"year":"1960"}}},"details":{"uninomial":{"uninomial":"Zu","authorship":{"verbatim":"Walters \u0026 Fitch 1960","normalized":"Walters \u0026 Fitch 1960","year":"1960","authors":["Walters","Fitch"],"originalAuth":{"authors":["Walters","Fitch"],"year":{"year":"1960"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10},{"wordType":"authorWord","start":13,"end":18},{"wordType":"year","start":19,"end":23}],"id":"c8724802-7dfb-5743-9988-a5f11b4c57b5","parserVersion":"test_version"}
```

Name: La Bleszynski 1966

Canonical: La

Authorship: Bleszynski 1966

```json
{"parsed":true,"parseQuality":1,"verbatim":"La Bleszynski 1966","normalized":"La Bleszynski 1966","canonical":{"stemmed":"La","simple":"La","full":"La"},"cardinality":1,"authorship":{"verbatim":"Bleszynski 1966","normalized":"Bleszynski 1966","year":"1966","authors":["Bleszynski"],"originalAuth":{"authors":["Bleszynski"],"year":{"year":"1966"}}},"details":{"uninomial":{"uninomial":"La","authorship":{"verbatim":"Bleszynski 1966","normalized":"Bleszynski 1966","year":"1966","authors":["Bleszynski"],"originalAuth":{"authors":["Bleszynski"],"year":{"year":"1966"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":13},{"wordType":"year","start":14,"end":18}],"id":"002f2de4-3661-5c8f-9175-cc1d1a9d6467","parserVersion":"test_version"}
```

Name: Qu Durkoop

Canonical: Qu

Authorship: Durkoop

```json
{"parsed":true,"parseQuality":1,"verbatim":"Qu Durkoop","normalized":"Qu Durkoop","canonical":{"stemmed":"Qu","simple":"Qu","full":"Qu"},"cardinality":1,"authorship":{"verbatim":"Durkoop","normalized":"Durkoop","authors":["Durkoop"],"originalAuth":{"authors":["Durkoop"]}},"details":{"uninomial":{"uninomial":"Qu","authorship":{"verbatim":"Durkoop","normalized":"Durkoop","authors":["Durkoop"],"originalAuth":{"authors":["Durkoop"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":10}],"id":"b4d879fa-028f-5b03-ad38-cc3a0765779a","parserVersion":"test_version"}
```

Name: As Slipinski 1982

Canonical: As

Authorship: Slipinski 1982

```json
{"parsed":true,"parseQuality":1,"verbatim":"As Slipinski 1982","normalized":"As Slipinski 1982","canonical":{"stemmed":"As","simple":"As","full":"As"},"cardinality":1,"authorship":{"verbatim":"Slipinski 1982","normalized":"Slipinski 1982","year":"1982","authors":["Slipinski"],"originalAuth":{"authors":["Slipinski"],"year":{"year":"1982"}}},"details":{"uninomial":{"uninomial":"As","authorship":{"verbatim":"Slipinski 1982","normalized":"Slipinski 1982","year":"1982","authors":["Slipinski"],"originalAuth":{"authors":["Slipinski"],"year":{"year":"1982"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":12},{"wordType":"year","start":13,"end":17}],"id":"55237f82-2126-5579-a8c6-385c0eb7ed8e","parserVersion":"test_version"}
```

Name: Ba Solem 1983

Canonical: Ba

Authorship: Solem 1983

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ba Solem 1983","normalized":"Ba Solem 1983","canonical":{"stemmed":"Ba","simple":"Ba","full":"Ba"},"cardinality":1,"authorship":{"verbatim":"Solem 1983","normalized":"Solem 1983","year":"1983","authors":["Solem"],"originalAuth":{"authors":["Solem"],"year":{"year":"1983"}}},"details":{"uninomial":{"uninomial":"Ba","authorship":{"verbatim":"Solem 1983","normalized":"Solem 1983","year":"1983","authors":["Solem"],"originalAuth":{"authors":["Solem"],"year":{"year":"1983"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":2},{"wordType":"authorWord","start":3,"end":8},{"wordType":"year","start":9,"end":13}],"id":"452f1a8e-711a-5b9c-906c-f475015229dd","parserVersion":"test_version"}
```

### Combination of two uninomials

Name: Poaceae subtrib. Scolochloinae Soreng

Canonical: Poaceae subtrib. Scolochloinae

Authorship: Soreng

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Poaceae subtrib. Scolochloinae Soreng","normalized":"Poaceae subtrib. Scolochloinae Soreng","canonical":{"stemmed":"Scolochloinae","simple":"Scolochloinae","full":"Poaceae subtrib. Scolochloinae"},"cardinality":1,"authorship":{"verbatim":"Soreng","normalized":"Soreng","authors":["Soreng"],"originalAuth":{"authors":["Soreng"]}},"details":{"uninomial":{"uninomial":"Scolochloinae","rank":"subtrib.","parent":"Poaceae","authorship":{"verbatim":"Soreng","normalized":"Soreng","authors":["Soreng"],"originalAuth":{"authors":["Soreng"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":7},{"wordType":"rank","start":8,"end":16},{"wordType":"uninomial","start":17,"end":30},{"wordType":"authorWord","start":31,"end":37}],"id":"d10510a7-ad50-587a-8411-e03d30d44214","parserVersion":"test_version"}
```

Name: Zygophyllaceae subfam. Tribuloideae D.M.Porter

Canonical: Zygophyllaceae subfam. Tribuloideae

Authorship: D. M. Porter

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Zygophyllaceae subfam. Tribuloideae D.M.Porter","normalized":"Zygophyllaceae subfam. Tribuloideae D. M. Porter","canonical":{"stemmed":"Tribuloideae","simple":"Tribuloideae","full":"Zygophyllaceae subfam. Tribuloideae"},"cardinality":1,"authorship":{"verbatim":"D.M.Porter","normalized":"D. M. Porter","authors":["D. M. Porter"],"originalAuth":{"authors":["D. M. Porter"]}},"details":{"uninomial":{"uninomial":"Tribuloideae","rank":"subfam.","parent":"Zygophyllaceae","authorship":{"verbatim":"D.M.Porter","normalized":"D. M. Porter","authors":["D. M. Porter"],"originalAuth":{"authors":["D. M. Porter"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":14},{"wordType":"rank","start":15,"end":22},{"wordType":"uninomial","start":23,"end":35},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":40,"end":46}],"id":"c60c1ff6-8e9d-5817-b49c-5845a5eaa9f5","parserVersion":"test_version"}
```

Name: Cordia (Adans.) Kuntze sect. Salimori

Canonical: Cordia sect. Salimori

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Cordia (Adans.) Kuntze sect. Salimori","normalized":"Cordia sect. Salimori","canonical":{"stemmed":"Salimori","simple":"Salimori","full":"Cordia sect. Salimori"},"cardinality":1,"details":{"uninomial":{"uninomial":"Salimori","rank":"sect.","parent":"Cordia"}},"pos":[{"wordType":"uninomial","start":0,"end":6},{"wordType":"authorWord","start":8,"end":14},{"wordType":"authorWord","start":16,"end":22},{"wordType":"rank","start":23,"end":28},{"wordType":"uninomial","start":29,"end":37}],"id":"48d5dbbe-50ff-50ae-a1f8-1cf4b3e2144b","parserVersion":"test_version"}
```

Name: Cordia sect. Salimori (Adans.) Kuntz

Canonical: Cordia sect. Salimori

Authorship: (Adans.) Kuntz

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Cordia sect. Salimori (Adans.) Kuntz","normalized":"Cordia sect. Salimori (Adans.) Kuntz","canonical":{"stemmed":"Salimori","simple":"Salimori","full":"Cordia sect. Salimori"},"cardinality":1,"authorship":{"verbatim":"(Adans.) Kuntz","normalized":"(Adans.) Kuntz","authors":["Adans.","Kuntz"],"originalAuth":{"authors":["Adans."]},"combinationAuth":{"authors":["Kuntz"]}},"details":{"uninomial":{"uninomial":"Salimori","rank":"sect.","parent":"Cordia","authorship":{"verbatim":"(Adans.) Kuntz","normalized":"(Adans.) Kuntz","authors":["Adans.","Kuntz"],"originalAuth":{"authors":["Adans."]},"combinationAuth":{"authors":["Kuntz"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":6},{"wordType":"rank","start":7,"end":12},{"wordType":"uninomial","start":13,"end":21},{"wordType":"authorWord","start":23,"end":29},{"wordType":"authorWord","start":31,"end":36}],"id":"337ef30d-f5da-5194-8bca-5354b262a05c","parserVersion":"test_version"}
```

Name: Poaceae supertrib. Arundinarodae L.Liu

Canonical: Poaceae supertrib. Arundinarodae

Authorship: L. Liu

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Poaceae supertrib. Arundinarodae L.Liu","normalized":"Poaceae supertrib. Arundinarodae L. Liu","canonical":{"stemmed":"Arundinarodae","simple":"Arundinarodae","full":"Poaceae supertrib. Arundinarodae"},"cardinality":1,"authorship":{"verbatim":"L.Liu","normalized":"L. Liu","authors":["L. Liu"],"originalAuth":{"authors":["L. Liu"]}},"details":{"uninomial":{"uninomial":"Arundinarodae","rank":"supertrib.","parent":"Poaceae","authorship":{"verbatim":"L.Liu","normalized":"L. Liu","authors":["L. Liu"],"originalAuth":{"authors":["L. Liu"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":7},{"wordType":"rank","start":8,"end":18},{"wordType":"uninomial","start":19,"end":32},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":35,"end":38}],"id":"c589a60b-1273-5b0b-93ea-25919d86647d","parserVersion":"test_version"}
```

Name: Alchemilla subsect. Sericeae A.Plocek

Canonical: Alchemilla subsect. Sericeae

Authorship: A. Plocek

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Alchemilla subsect. Sericeae A.Plocek","normalized":"Alchemilla subsect. Sericeae A. Plocek","canonical":{"stemmed":"Sericeae","simple":"Sericeae","full":"Alchemilla subsect. Sericeae"},"cardinality":1,"authorship":{"verbatim":"A.Plocek","normalized":"A. Plocek","authors":["A. Plocek"],"originalAuth":{"authors":["A. Plocek"]}},"details":{"uninomial":{"uninomial":"Sericeae","rank":"subsect.","parent":"Alchemilla","authorship":{"verbatim":"A.Plocek","normalized":"A. Plocek","authors":["A. Plocek"],"originalAuth":{"authors":["A. Plocek"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":10},{"wordType":"rank","start":11,"end":19},{"wordType":"uninomial","start":20,"end":28},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":31,"end":37}],"id":"bedd1b9c-91dd-5ad9-9cd6-0504b85aae30","parserVersion":"test_version"}
```

Name: Hymenophyllum subgen. Hymenoglossum (Presl) R.M.Tryon & A.Tryon

Canonical: Hymenophyllum subgen. Hymenoglossum

Authorship: (Presl) R. M. Tryon & A. Tryon

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Hymenophyllum subgen. Hymenoglossum (Presl) R.M.Tryon \u0026 A.Tryon","normalized":"Hymenophyllum subgen. Hymenoglossum (Presl) R. M. Tryon \u0026 A. Tryon","canonical":{"stemmed":"Hymenoglossum","simple":"Hymenoglossum","full":"Hymenophyllum subgen. Hymenoglossum"},"cardinality":1,"authorship":{"verbatim":"(Presl) R.M.Tryon \u0026 A.Tryon","normalized":"(Presl) R. M. Tryon \u0026 A. Tryon","authors":["Presl","R. M. Tryon","A. Tryon"],"originalAuth":{"authors":["Presl"]},"combinationAuth":{"authors":["R. M. Tryon","A. Tryon"]}},"details":{"uninomial":{"uninomial":"Hymenoglossum","rank":"subgen.","parent":"Hymenophyllum","authorship":{"verbatim":"(Presl) R.M.Tryon \u0026 A.Tryon","normalized":"(Presl) R. M. Tryon \u0026 A. Tryon","authors":["Presl","R. M. Tryon","A. Tryon"],"originalAuth":{"authors":["Presl"]},"combinationAuth":{"authors":["R. M. Tryon","A. Tryon"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":13},{"wordType":"rank","start":14,"end":21},{"wordType":"uninomial","start":22,"end":35},{"wordType":"authorWord","start":37,"end":42},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":46,"end":48},{"wordType":"authorWord","start":48,"end":53},{"wordType":"authorWord","start":56,"end":58},{"wordType":"authorWord","start":58,"end":63}],"id":"22ea4710-3a2a-5526-a42e-7c7ff508ee79","parserVersion":"test_version"}
```

Name: Pereskia subg. Maihuenia Philippi ex F.A.C.Weber, 1898

Canonical: Pereskia subgen. Maihuenia

Authorship: Philippi ex F. A. C. Weber 1898

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Pereskia subg. Maihuenia Philippi ex F.A.C.Weber, 1898","normalized":"Pereskia subgen. Maihuenia Philippi ex F. A. C. Weber 1898","canonical":{"stemmed":"Maihuenia","simple":"Maihuenia","full":"Pereskia subgen. Maihuenia"},"cardinality":1,"authorship":{"verbatim":"Philippi ex F.A.C.Weber, 1898","normalized":"Philippi ex F. A. C. Weber 1898","authors":["Philippi"],"originalAuth":{"authors":["Philippi"],"exAuthors":{"authors":["F. A. C. Weber"],"year":{"year":"1898"}}}},"details":{"uninomial":{"uninomial":"Maihuenia","rank":"subgen.","parent":"Pereskia","authorship":{"verbatim":"Philippi ex F.A.C.Weber, 1898","normalized":"Philippi ex F. A. C. Weber 1898","authors":["Philippi"],"originalAuth":{"authors":["Philippi"],"exAuthors":{"authors":["F. A. C. Weber"],"year":{"year":"1898"}}}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"rank","start":9,"end":14},{"wordType":"uninomial","start":15,"end":24},{"wordType":"authorWord","start":25,"end":33},{"wordType":"authorWord","start":37,"end":39},{"wordType":"authorWord","start":39,"end":41},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":43,"end":48},{"wordType":"year","start":50,"end":54}],"id":"344bd8c1-a4d2-5120-a738-0903aafad63d","parserVersion":"test_version"}
```

Name: Aconitum ser. Tangutica W.T. Wang

Canonical: Aconitum ser. Tangutica

Authorship: W. T. Wang

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aconitum ser. Tangutica W.T. Wang","normalized":"Aconitum ser. Tangutica W. T. Wang","canonical":{"stemmed":"Tangutica","simple":"Tangutica","full":"Aconitum ser. Tangutica"},"cardinality":1,"authorship":{"verbatim":"W.T. Wang","normalized":"W. T. Wang","authors":["W. T. Wang"],"originalAuth":{"authors":["W. T. Wang"]}},"details":{"uninomial":{"uninomial":"Tangutica","rank":"ser.","parent":"Aconitum","authorship":{"verbatim":"W.T. Wang","normalized":"W. T. Wang","authors":["W. T. Wang"],"originalAuth":{"authors":["W. T. Wang"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"rank","start":9,"end":13},{"wordType":"uninomial","start":14,"end":23},{"wordType":"authorWord","start":24,"end":26},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":29,"end":33}],"id":"8f5d7bd0-90a1-556d-a8ef-1a440b157c34","parserVersion":"test_version"}
```

Name: Calathus (Lindrothius) KURNAKOV 1961

Canonical: Calathus subgen. Lindrothius

Authorship: Kurnakov 1961

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Calathus (Lindrothius) KURNAKOV 1961","normalized":"Calathus subgen. Lindrothius Kurnakov 1961","canonical":{"stemmed":"Lindrothius","simple":"Lindrothius","full":"Calathus subgen. Lindrothius"},"cardinality":1,"authorship":{"verbatim":"KURNAKOV 1961","normalized":"Kurnakov 1961","year":"1961","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1961"}}},"details":{"uninomial":{"uninomial":"Lindrothius","rank":"subgen.","parent":"Calathus","authorship":{"verbatim":"KURNAKOV 1961","normalized":"Kurnakov 1961","year":"1961","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1961"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"uninomial","start":10,"end":21},{"wordType":"authorWord","start":23,"end":31},{"wordType":"year","start":32,"end":36}],"id":"aa113505-61a1-58fe-92f3-8fd511dcfd61","parserVersion":"test_version"}
```

Name: Eucalyptus subser. Regulares Brooker

Canonical: Eucalyptus subser. Regulares

Authorship: Brooker

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Eucalyptus subser. Regulares Brooker","normalized":"Eucalyptus subser. Regulares Brooker","canonical":{"stemmed":"Regulares","simple":"Regulares","full":"Eucalyptus subser. Regulares"},"cardinality":1,"authorship":{"verbatim":"Brooker","normalized":"Brooker","authors":["Brooker"],"originalAuth":{"authors":["Brooker"]}},"details":{"uninomial":{"uninomial":"Regulares","rank":"subser.","parent":"Eucalyptus","authorship":{"verbatim":"Brooker","normalized":"Brooker","authors":["Brooker"],"originalAuth":{"authors":["Brooker"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":10},{"wordType":"rank","start":11,"end":18},{"wordType":"uninomial","start":19,"end":28},{"wordType":"authorWord","start":29,"end":36}],"id":"783aa15c-f54f-5233-b792-16774a21a34d","parserVersion":"test_version"}
```

Name: Aaleniella (Danocythere)

Canonical: Aaleniella subgen. Danocythere

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aaleniella (Danocythere)","normalized":"Aaleniella subgen. Danocythere","canonical":{"stemmed":"Danocythere","simple":"Danocythere","full":"Aaleniella subgen. Danocythere"},"cardinality":1,"details":{"uninomial":{"uninomial":"Danocythere","rank":"subgen.","parent":"Aaleniella"}},"pos":[{"wordType":"uninomial","start":0,"end":10},{"wordType":"uninomial","start":12,"end":23}],"id":"8b7eddb1-b9a4-5cca-8fa8-25527e25d8df","parserVersion":"test_version"}
```

### ICN names that look like combined uninomials for ICZN

Name: Clathrotropis (Bentham) Harms in Dalla Torre & Harms, 1901

Canonical: Clathrotropis

Authorship: (Bentham) Harms ex Dalla Torre & Harms 1901

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Clathrotropis (Bentham) Harms in Dalla Torre \u0026 Harms, 1901","normalized":"Clathrotropis (Bentham) Harms ex Dalla Torre \u0026 Harms 1901","canonical":{"stemmed":"Clathrotropis","simple":"Clathrotropis","full":"Clathrotropis"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Bentham) Harms ex Dalla Torre \u0026 Harms 1901","authors":["Bentham","Harms"],"originalAuth":{"authors":["Bentham"]},"combinationAuth":{"authors":["Harms"],"exAuthors":{"authors":["Dalla Torre","Harms"],"year":{"year":"1901"}}}},"details":{"uninomial":{"uninomial":"Clathrotropis","authorship":{"verbatim":"","normalized":"(Bentham) Harms ex Dalla Torre \u0026 Harms 1901","authors":["Bentham","Harms"],"originalAuth":{"authors":["Bentham"]},"combinationAuth":{"authors":["Harms"],"exAuthors":{"authors":["Dalla Torre","Harms"],"year":{"year":"1901"}}}}}},"pos":[{"wordType":"uninomial","start":0,"end":13},{"wordType":"authorWord","start":15,"end":22},{"wordType":"authorWord","start":24,"end":29},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWord","start":39,"end":44},{"wordType":"authorWord","start":47,"end":52},{"wordType":"year","start":54,"end":58}],"id":"6b730cea-e81b-53ba-a511-caaa233b9b84","parserVersion":"test_version"}
```

Name: Humiriastrum (Urban) Cuatrecasas, 1961

Canonical: Humiriastrum

Authorship: (Urban) Cuatrecasas 1961

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Humiriastrum (Urban) Cuatrecasas, 1961","normalized":"Humiriastrum (Urban) Cuatrecasas 1961","canonical":{"stemmed":"Humiriastrum","simple":"Humiriastrum","full":"Humiriastrum"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Urban) Cuatrecasas 1961","authors":["Urban","Cuatrecasas"],"originalAuth":{"authors":["Urban"]},"combinationAuth":{"authors":["Cuatrecasas"],"year":{"year":"1961"}}},"details":{"uninomial":{"uninomial":"Humiriastrum","authorship":{"verbatim":"","normalized":"(Urban) Cuatrecasas 1961","authors":["Urban","Cuatrecasas"],"originalAuth":{"authors":["Urban"]},"combinationAuth":{"authors":["Cuatrecasas"],"year":{"year":"1961"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":12},{"wordType":"authorWord","start":14,"end":19},{"wordType":"authorWord","start":21,"end":32},{"wordType":"year","start":34,"end":38}],"id":"98f8aa31-1cc3-59c2-a4f2-ebf18e0929ab","parserVersion":"test_version"}
```

Name: Pampocactus (Doweld) Doweld

Canonical: Pampocactus

Authorship: (Doweld) Doweld

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Pampocactus (Doweld) Doweld","normalized":"Pampocactus (Doweld) Doweld","canonical":{"stemmed":"Pampocactus","simple":"Pampocactus","full":"Pampocactus"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Doweld) Doweld","authors":["Doweld","Doweld"],"originalAuth":{"authors":["Doweld"]},"combinationAuth":{"authors":["Doweld"]}},"details":{"uninomial":{"uninomial":"Pampocactus","authorship":{"verbatim":"","normalized":"(Doweld) Doweld","authors":["Doweld","Doweld"],"originalAuth":{"authors":["Doweld"]},"combinationAuth":{"authors":["Doweld"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":13,"end":19},{"wordType":"authorWord","start":21,"end":27}],"id":"82494c70-6400-51a3-b786-2a8a747f8305","parserVersion":"test_version"}
```

Name: Pampocactus (Doweld)

Canonical: Pampocactus

Authorship: (Doweld)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Pampocactus (Doweld)","normalized":"Pampocactus (Doweld)","canonical":{"stemmed":"Pampocactus","simple":"Pampocactus","full":"Pampocactus"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Doweld)","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]}},"details":{"uninomial":{"uninomial":"Pampocactus","authorship":{"verbatim":"","normalized":"(Doweld)","authors":["Doweld"],"originalAuth":{"authors":["Doweld"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":13,"end":19}],"id":"3ed64c9a-ec8a-52c9-a913-eae09b6c71b9","parserVersion":"test_version"}
```

Name: Drepanolejeunea (Spruce) (Steph.)

Canonical: Drepanolejeunea

Authorship: (Spruce)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Drepanolejeunea (Spruce) (Steph.)","normalized":"Drepanolejeunea (Spruce)","canonical":{"stemmed":"Drepanolejeunea","simple":"Drepanolejeunea","full":"Drepanolejeunea"},"cardinality":1,"authorship":{"verbatim":"","normalized":"(Spruce)","authors":["Spruce"],"originalAuth":{"authors":["Spruce"]}},"tail":"(Steph.)","details":{"uninomial":{"uninomial":"Drepanolejeunea","authorship":{"verbatim":"","normalized":"(Spruce)","authors":["Spruce"],"originalAuth":{"authors":["Spruce"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":17,"end":23}],"id":"19265c95-0a2b-5e8a-b2c4-478716e9c9ec","parserVersion":"test_version"}
```


### Binomials without authorship

Name: Notopholia corrusca

Canonical: Notopholia corrusca

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Notopholia corrusca","normalized":"Notopholia corrusca","canonical":{"stemmed":"Notopholia corrusc","simple":"Notopholia corrusca","full":"Notopholia corrusca"},"cardinality":2,"details":{"species":{"genus":"Notopholia","species":"corrusca"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19}],"id":"755cef9c-65e4-598d-abf5-4d4a91be9845","parserVersion":"test_version"}
```

Name: Cyathicula scelobelonium

Canonical: Cyathicula scelobelonium

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cyathicula scelobelonium","normalized":"Cyathicula scelobelonium","canonical":{"stemmed":"Cyathicula scelobeloni","simple":"Cyathicula scelobelonium","full":"Cyathicula scelobelonium"},"cardinality":2,"details":{"species":{"genus":"Cyathicula","species":"scelobelonium"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":24}],"id":"21047543-b5ef-5426-b2b4-bc19f3498407","parserVersion":"test_version"}
```

Name: Pseudocercospora     dendrobii

Canonical: Pseudocercospora dendrobii

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Pseudocercospora     dendrobii","normalized":"Pseudocercospora dendrobii","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii"}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":21,"end":30}],"id":"5b320aa4-d417-5eda-be2d-83632e0d3624","parserVersion":"test_version"}
```

Name: Cucurbita pepo

Canonical: Cucurbita pepo

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cucurbita pepo","normalized":"Cucurbita pepo","canonical":{"stemmed":"Cucurbita pep","simple":"Cucurbita pepo","full":"Cucurbita pepo"},"cardinality":2,"details":{"species":{"genus":"Cucurbita","species":"pepo"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":14}],"id":"022e85ce-a786-5478-9799-ac2e0f2cc726","parserVersion":"test_version"}
```

Name: Hirsutëlla mâle

Canonical: Hirsutella male

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hirsutëlla mâle","normalized":"Hirsutella male","canonical":{"stemmed":"Hirsutella mal","simple":"Hirsutella male","full":"Hirsutella male"},"cardinality":2,"details":{"species":{"genus":"Hirsutella","species":"male"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":15}],"id":"62cc5704-b486-5aba-882c-dc29f5282179","parserVersion":"test_version"}
```

Name: Aëtosaurus ferratus

Canonical: Aetosaurus ferratus

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Aëtosaurus ferratus","normalized":"Aetosaurus ferratus","canonical":{"stemmed":"Aetosaurus ferrat","simple":"Aetosaurus ferratus","full":"Aetosaurus ferratus"},"cardinality":2,"details":{"species":{"genus":"Aetosaurus","species":"ferratus"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19}],"id":"9d95ffa0-0203-541f-854a-77ca7ff187fa","parserVersion":"test_version"}
```

Name: Remera cvancarai

Canonical: Remera cvancarai

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Remera cvancarai","normalized":"Remera cvancarai","canonical":{"stemmed":"Remera cuancara","simple":"Remera cvancarai","full":"Remera cvancarai"},"cardinality":2,"details":{"species":{"genus":"Remera","species":"cvancarai"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":16}],"id":"d5d77ab3-2648-5409-a6c7-e3e20d75c38b","parserVersion":"test_version"}
```

### Binomials with authorship

Name: Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012

Canonical: Nototriton matama

Authorship: Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños & Wake 2012

```json
{"parsed":true,"parseQuality":1,"verbatim":"Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Nototriton matama Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","canonical":{"stemmed":"Nototriton matam","simple":"Nototriton matama","full":"Nototriton matama"},"cardinality":2,"authorship":{"verbatim":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","year":"2012","authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"originalAuth":{"authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"year":{"year":"2012"}}},"details":{"species":{"genus":"Nototriton","species":"matama","authorship":{"verbatim":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños, and Wake, 2012","normalized":"Boza-Oviedo, Rovito, Chaves, García-Rodríguez, Artavia, Bolaños \u0026 Wake 2012","year":"2012","authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"originalAuth":{"authors":["Boza-Oviedo","Rovito","Chaves","García-Rodríguez","Artavia","Bolaños","Wake"],"year":{"year":"2012"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":17},{"wordType":"authorWord","start":18,"end":29},{"wordType":"authorWord","start":31,"end":37},{"wordType":"authorWord","start":39,"end":45},{"wordType":"authorWord","start":47,"end":63},{"wordType":"authorWord","start":65,"end":72},{"wordType":"authorWord","start":74,"end":81},{"wordType":"authorWord","start":87,"end":91},{"wordType":"year","start":93,"end":97}],"id":"49503e24-3297-57c6-bc6e-c1a68a338fd3","parserVersion":"test_version"}
```

Name: Architectonica offlexa Iredale, 1931

Canonical: Architectonica offlexa

Authorship: Iredale 1931

```json
{"parsed":true,"parseQuality":1,"verbatim":"Architectonica offlexa Iredale, 1931","normalized":"Architectonica offlexa Iredale 1931","canonical":{"stemmed":"Architectonica offlex","simple":"Architectonica offlexa","full":"Architectonica offlexa"},"cardinality":2,"authorship":{"verbatim":"Iredale, 1931","normalized":"Iredale 1931","year":"1931","authors":["Iredale"],"originalAuth":{"authors":["Iredale"],"year":{"year":"1931"}}},"details":{"species":{"genus":"Architectonica","species":"offlexa","authorship":{"verbatim":"Iredale, 1931","normalized":"Iredale 1931","year":"1931","authors":["Iredale"],"originalAuth":{"authors":["Iredale"],"year":{"year":"1931"}}}}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22},{"wordType":"authorWord","start":23,"end":30},{"wordType":"year","start":32,"end":36}],"id":"d8088d2a-6d20-5ef6-9ec8-68753e2e6da0","parserVersion":"test_version"}
```

Name: Maracanda amoena Mc'Lach

Canonical: Maracanda amoena

Authorship: Mc'Lach

```json
{"parsed":true,"parseQuality":1,"verbatim":"Maracanda amoena Mc'Lach","normalized":"Maracanda amoena Mc'Lach","canonical":{"stemmed":"Maracanda amoen","simple":"Maracanda amoena","full":"Maracanda amoena"},"cardinality":2,"authorship":{"verbatim":"Mc'Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}},"details":{"species":{"genus":"Maracanda","species":"amoena","authorship":{"verbatim":"Mc'Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":16},{"wordType":"authorWord","start":17,"end":24}],"id":"b561edfc-29e8-5e8d-8849-60899356be0d","parserVersion":"test_version"}
```

Name: Maracanda amoena Mc’Lach

Canonical: Maracanda amoena

Authorship: Mc'Lach

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"}],"verbatim":"Maracanda amoena Mc’Lach","normalized":"Maracanda amoena Mc'Lach","canonical":{"stemmed":"Maracanda amoen","simple":"Maracanda amoena","full":"Maracanda amoena"},"cardinality":2,"authorship":{"verbatim":"Mc’Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}},"details":{"species":{"genus":"Maracanda","species":"amoena","authorship":{"verbatim":"Mc’Lach","normalized":"Mc'Lach","authors":["Mc'Lach"],"originalAuth":{"authors":["Mc'Lach"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":16},{"wordType":"authorWord","start":17,"end":24}],"id":"98ddd2f7-2f78-5970-adac-677273dc3caf","parserVersion":"test_version"}
```

Name: Tridentella tangeroae Bruce, 198?

Canonical: Tridentella tangeroae

Authorship: Bruce (198?)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with question mark"}],"verbatim":"Tridentella tangeroae Bruce, 198?","normalized":"Tridentella tangeroae Bruce (198?)","canonical":{"stemmed":"Tridentella tangero","simple":"Tridentella tangeroae","full":"Tridentella tangeroae"},"cardinality":2,"authorship":{"verbatim":"Bruce, 198?","normalized":"Bruce (198?)","year":"(198?)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"198?","isApproximate":true}}},"details":{"species":{"genus":"Tridentella","species":"tangeroae","authorship":{"verbatim":"Bruce, 198?","normalized":"Bruce (198?)","year":"(198?)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"198?","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":22,"end":27},{"wordType":"approximateYear","start":29,"end":33}],"id":"179d63c9-bad4-5e61-bf2e-7261b4aa5066","parserVersion":"test_version"}
```

Name: Zanthopsis bispinosa M'Coy, 1849

Canonical: Zanthopsis bispinosa

Authorship: M'Coy 1849

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zanthopsis bispinosa M'Coy, 1849","normalized":"Zanthopsis bispinosa M'Coy 1849","canonical":{"stemmed":"Zanthopsis bispinos","simple":"Zanthopsis bispinosa","full":"Zanthopsis bispinosa"},"cardinality":2,"authorship":{"verbatim":"M'Coy, 1849","normalized":"M'Coy 1849","year":"1849","authors":["M'Coy"],"originalAuth":{"authors":["M'Coy"],"year":{"year":"1849"}}},"details":{"species":{"genus":"Zanthopsis","species":"bispinosa","authorship":{"verbatim":"M'Coy, 1849","normalized":"M'Coy 1849","year":"1849","authors":["M'Coy"],"originalAuth":{"authors":["M'Coy"],"year":{"year":"1849"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":20},{"wordType":"authorWord","start":21,"end":26},{"wordType":"year","start":28,"end":32}],"id":"88b58b88-d8fd-55d9-a9c4-ddd11459820e","parserVersion":"test_version"}
```

Name: Scilla rupestris v.d. Merwe

Canonical: Scilla rupestris

Authorship: v.d. Merwe

```json
{"parsed":true,"parseQuality":1,"verbatim":"Scilla rupestris v.d. Merwe","normalized":"Scilla rupestris v.d. Merwe","canonical":{"stemmed":"Scilla rupestr","simple":"Scilla rupestris","full":"Scilla rupestris"},"cardinality":2,"authorship":{"verbatim":"v.d. Merwe","normalized":"v.d. Merwe","authors":["v.d. Merwe"],"originalAuth":{"authors":["v.d. Merwe"]}},"details":{"species":{"genus":"Scilla","species":"rupestris","authorship":{"verbatim":"v.d. Merwe","normalized":"v.d. Merwe","authors":["v.d. Merwe"],"originalAuth":{"authors":["v.d. Merwe"]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":16},{"wordType":"authorWord","start":17,"end":21},{"wordType":"authorWord","start":22,"end":27}],"id":"72ec3a37-8a80-5a82-97dd-b6a67a52d209","parserVersion":"test_version"}
```

Name: Bembix bidentata v.d.L.

Canonical: Bembix bidentata

Authorship: v.d. L.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Bembix bidentata v.d.L.","normalized":"Bembix bidentata v.d. L.","canonical":{"stemmed":"Bembix bidentat","simple":"Bembix bidentata","full":"Bembix bidentata"},"cardinality":2,"authorship":{"verbatim":"v.d.L.","normalized":"v.d. L.","authors":["v.d. L."],"originalAuth":{"authors":["v.d. L."]}},"details":{"species":{"genus":"Bembix","species":"bidentata","authorship":{"verbatim":"v.d.L.","normalized":"v.d. L.","authors":["v.d. L."],"originalAuth":{"authors":["v.d. L."]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":16},{"wordType":"authorWord","start":17,"end":21},{"wordType":"authorWord","start":21,"end":23}],"id":"6f226f43-dfa0-5d61-8a3f-200b2277fcf2","parserVersion":"test_version"}
```

Name: Pompilus cinctellus v. d. L.

Canonical: Pompilus cinctellus

Authorship: v. d. L.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pompilus cinctellus v. d. L.","normalized":"Pompilus cinctellus v. d. L.","canonical":{"stemmed":"Pompilus cinctell","simple":"Pompilus cinctellus","full":"Pompilus cinctellus"},"cardinality":2,"authorship":{"verbatim":"v. d. L.","normalized":"v. d. L.","authors":["v. d. L."],"originalAuth":{"authors":["v. d. L."]}},"details":{"species":{"genus":"Pompilus","species":"cinctellus","authorship":{"verbatim":"v. d. L.","normalized":"v. d. L.","authors":["v. d. L."],"originalAuth":{"authors":["v. d. L."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"authorWord","start":20,"end":25},{"wordType":"authorWord","start":26,"end":28}],"id":"8954c0f2-eab4-561d-9f94-6cebd4f8024d","parserVersion":"test_version"}
```

Name: Setaphis viridis v. d.G.

Canonical: Setaphis viridis

Authorship: v. d. G.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Setaphis viridis v. d.G.","normalized":"Setaphis viridis v. d. G.","canonical":{"stemmed":"Setaphis uirid","simple":"Setaphis viridis","full":"Setaphis viridis"},"cardinality":2,"authorship":{"verbatim":"v. d.G.","normalized":"v. d. G.","authors":["v. d. G."],"originalAuth":{"authors":["v. d. G."]}},"details":{"species":{"genus":"Setaphis","species":"viridis","authorship":{"verbatim":"v. d.G.","normalized":"v. d. G.","authors":["v. d. G."],"originalAuth":{"authors":["v. d. G."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":17,"end":22},{"wordType":"authorWord","start":22,"end":24}],"id":"19792117-31fc-52d7-9990-e89b67c459d3","parserVersion":"test_version"}
```

Name: Coleophora mendica Baldizzone & v. d.Wolf 2000

Canonical: Coleophora mendica

Authorship: Baldizzone & v. d. Wolf 2000

```json
{"parsed":true,"parseQuality":1,"verbatim":"Coleophora mendica Baldizzone \u0026 v. d.Wolf 2000","normalized":"Coleophora mendica Baldizzone \u0026 v. d. Wolf 2000","canonical":{"stemmed":"Coleophora mendic","simple":"Coleophora mendica","full":"Coleophora mendica"},"cardinality":2,"authorship":{"verbatim":"Baldizzone \u0026 v. d.Wolf 2000","normalized":"Baldizzone \u0026 v. d. Wolf 2000","year":"2000","authors":["Baldizzone","v. d. Wolf"],"originalAuth":{"authors":["Baldizzone","v. d. Wolf"],"year":{"year":"2000"}}},"details":{"species":{"genus":"Coleophora","species":"mendica","authorship":{"verbatim":"Baldizzone \u0026 v. d.Wolf 2000","normalized":"Baldizzone \u0026 v. d. Wolf 2000","year":"2000","authors":["Baldizzone","v. d. Wolf"],"originalAuth":{"authors":["Baldizzone","v. d. Wolf"],"year":{"year":"2000"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":18},{"wordType":"authorWord","start":19,"end":29},{"wordType":"authorWord","start":32,"end":37},{"wordType":"authorWord","start":37,"end":41},{"wordType":"year","start":42,"end":46}],"id":"982affab-249b-5858-8ea1-ba226378c233","parserVersion":"test_version"}
```

Name: Psoronaias semigranosa von dem Busch in Philippi, 1845

Canonical: Psoronaias semigranosa

Authorship: von dem Busch ex Philippi 1845

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Psoronaias semigranosa von dem Busch in Philippi, 1845","normalized":"Psoronaias semigranosa von dem Busch ex Philippi 1845","canonical":{"stemmed":"Psoronaias semigranos","simple":"Psoronaias semigranosa","full":"Psoronaias semigranosa"},"cardinality":2,"authorship":{"verbatim":"von dem Busch in Philippi, 1845","normalized":"von dem Busch ex Philippi 1845","authors":["von dem Busch"],"originalAuth":{"authors":["von dem Busch"],"exAuthors":{"authors":["Philippi"],"year":{"year":"1845"}}}},"details":{"species":{"genus":"Psoronaias","species":"semigranosa","authorship":{"verbatim":"von dem Busch in Philippi, 1845","normalized":"von dem Busch ex Philippi 1845","authors":["von dem Busch"],"originalAuth":{"authors":["von dem Busch"],"exAuthors":{"authors":["Philippi"],"year":{"year":"1845"}}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"authorWord","start":23,"end":30},{"wordType":"authorWord","start":31,"end":36},{"wordType":"authorWord","start":40,"end":48},{"wordType":"year","start":50,"end":54}],"id":"948809ee-be49-598d-a755-fded9ba496c5","parserVersion":"test_version"}
```

Name: Phora sororcula v d Wulp 1871

Canonical: Phora sororcula

Authorship: v d Wulp 1871

```json
{"parsed":true,"parseQuality":1,"verbatim":"Phora sororcula v d Wulp 1871","normalized":"Phora sororcula v d Wulp 1871","canonical":{"stemmed":"Phora sororcul","simple":"Phora sororcula","full":"Phora sororcula"},"cardinality":2,"authorship":{"verbatim":"v d Wulp 1871","normalized":"v d Wulp 1871","year":"1871","authors":["v d Wulp"],"originalAuth":{"authors":["v d Wulp"],"year":{"year":"1871"}}},"details":{"species":{"genus":"Phora","species":"sororcula","authorship":{"verbatim":"v d Wulp 1871","normalized":"v d Wulp 1871","year":"1871","authors":["v d Wulp"],"originalAuth":{"authors":["v d Wulp"],"year":{"year":"1871"}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":15},{"wordType":"authorWord","start":16,"end":19},{"wordType":"authorWord","start":20,"end":24},{"wordType":"year","start":25,"end":29}],"id":"dad2ef8b-4f74-5de5-844b-29b6ee09ce68","parserVersion":"test_version"}
```

Name: Aeolothrips andalusiacus zur Strassen 1973

Canonical: Aeolothrips andalusiacus

Authorship: zur Strassen 1973

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aeolothrips andalusiacus zur Strassen 1973","normalized":"Aeolothrips andalusiacus zur Strassen 1973","canonical":{"stemmed":"Aeolothrips andalusiac","simple":"Aeolothrips andalusiacus","full":"Aeolothrips andalusiacus"},"cardinality":2,"authorship":{"verbatim":"zur Strassen 1973","normalized":"zur Strassen 1973","year":"1973","authors":["zur Strassen"],"originalAuth":{"authors":["zur Strassen"],"year":{"year":"1973"}}},"details":{"species":{"genus":"Aeolothrips","species":"andalusiacus","authorship":{"verbatim":"zur Strassen 1973","normalized":"zur Strassen 1973","year":"1973","authors":["zur Strassen"],"originalAuth":{"authors":["zur Strassen"],"year":{"year":"1973"}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":24},{"wordType":"authorWord","start":25,"end":28},{"wordType":"authorWord","start":29,"end":37},{"wordType":"year","start":38,"end":42}],"id":"1e99cbcb-7fc9-5454-a40b-4786d3e35751","parserVersion":"test_version"}
```

Name: Orthosia kindermannii Fischer v. Roslerstamm, 1837

Canonical: Orthosia kindermannii

Authorship: Fischer v. Roslerstamm 1837

```json
{"parsed":true,"parseQuality":1,"verbatim":"Orthosia kindermannii Fischer v. Roslerstamm, 1837","normalized":"Orthosia kindermannii Fischer v. Roslerstamm 1837","canonical":{"stemmed":"Orthosia kindermanni","simple":"Orthosia kindermannii","full":"Orthosia kindermannii"},"cardinality":2,"authorship":{"verbatim":"Fischer v. Roslerstamm, 1837","normalized":"Fischer v. Roslerstamm 1837","year":"1837","authors":["Fischer v. Roslerstamm"],"originalAuth":{"authors":["Fischer v. Roslerstamm"],"year":{"year":"1837"}}},"details":{"species":{"genus":"Orthosia","species":"kindermannii","authorship":{"verbatim":"Fischer v. Roslerstamm, 1837","normalized":"Fischer v. Roslerstamm 1837","year":"1837","authors":["Fischer v. Roslerstamm"],"originalAuth":{"authors":["Fischer v. Roslerstamm"],"year":{"year":"1837"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":29},{"wordType":"authorWord","start":30,"end":32},{"wordType":"authorWord","start":33,"end":44},{"wordType":"year","start":46,"end":50}],"id":"53abecc3-4083-5cdc-966c-09648fe9383d","parserVersion":"test_version"}
```

Name: Nereidavus kulkovi Kul'kov in Kul'kov & Obut, 1973

Canonical: Nereidavus kulkovi

Authorship: Kul'kov ex Kul'kov & Obut 1973

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Nereidavus kulkovi Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Nereidavus kulkovi Kul'kov ex Kul'kov \u0026 Obut 1973","canonical":{"stemmed":"Nereidavus kulkou","simple":"Nereidavus kulkovi","full":"Nereidavus kulkovi"},"cardinality":2,"authorship":{"verbatim":"Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Kul'kov ex Kul'kov \u0026 Obut 1973","authors":["Kul'kov"],"originalAuth":{"authors":["Kul'kov"],"exAuthors":{"authors":["Kul'kov","Obut"],"year":{"year":"1973"}}}},"details":{"species":{"genus":"Nereidavus","species":"kulkovi","authorship":{"verbatim":"Kul'kov in Kul'kov \u0026 Obut, 1973","normalized":"Kul'kov ex Kul'kov \u0026 Obut 1973","authors":["Kul'kov"],"originalAuth":{"authors":["Kul'kov"],"exAuthors":{"authors":["Kul'kov","Obut"],"year":{"year":"1973"}}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":18},{"wordType":"authorWord","start":19,"end":26},{"wordType":"authorWord","start":30,"end":37},{"wordType":"authorWord","start":40,"end":44},{"wordType":"year","start":46,"end":50}],"id":"4aa8305f-884f-5515-9bdc-f586e037028c","parserVersion":"test_version"}
```

Name: Xylaria potentillae A S. Xu

Canonical: Xylaria potentillae

Authorship: A S. Xu

```json
{"parsed":true,"parseQuality":1,"verbatim":"Xylaria potentillae A S. Xu","normalized":"Xylaria potentillae A S. Xu","canonical":{"stemmed":"Xylaria potentill","simple":"Xylaria potentillae","full":"Xylaria potentillae"},"cardinality":2,"authorship":{"verbatim":"A S. Xu","normalized":"A S. Xu","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"]}},"details":{"species":{"genus":"Xylaria","species":"potentillae","authorship":{"verbatim":"A S. Xu","normalized":"A S. Xu","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":19},{"wordType":"authorWord","start":20,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWord","start":25,"end":27}],"id":"6bc4bb61-e0b9-5c22-a9b6-46c45757f2c2","parserVersion":"test_version"}
```

Name: Pseudocyrtopora el Hajjaji 1987

Canonical: Pseudocyrtopora

Authorship: el Hajjaji 1987

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocyrtopora el Hajjaji 1987","normalized":"Pseudocyrtopora el Hajjaji 1987","canonical":{"stemmed":"Pseudocyrtopora","simple":"Pseudocyrtopora","full":"Pseudocyrtopora"},"cardinality":1,"authorship":{"verbatim":"el Hajjaji 1987","normalized":"el Hajjaji 1987","year":"1987","authors":["el Hajjaji"],"originalAuth":{"authors":["el Hajjaji"],"year":{"year":"1987"}}},"details":{"uninomial":{"uninomial":"Pseudocyrtopora","authorship":{"verbatim":"el Hajjaji 1987","normalized":"el Hajjaji 1987","year":"1987","authors":["el Hajjaji"],"originalAuth":{"authors":["el Hajjaji"],"year":{"year":"1987"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":19,"end":26},{"wordType":"year","start":27,"end":31}],"id":"61db186c-cbf4-5949-9fd1-79efe7157873","parserVersion":"test_version"}
```

Name: Geositta poeciloptera (zu Wied-Neuwied, 1830)

Canonical: Geositta poeciloptera

Authorship: (zu Wied-Neuwied 1830)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Geositta poeciloptera (zu Wied-Neuwied, 1830)","normalized":"Geositta poeciloptera (zu Wied-Neuwied 1830)","canonical":{"stemmed":"Geositta poecilopter","simple":"Geositta poeciloptera","full":"Geositta poeciloptera"},"cardinality":2,"authorship":{"verbatim":"(zu Wied-Neuwied, 1830)","normalized":"(zu Wied-Neuwied 1830)","year":"1830","authors":["zu Wied-Neuwied"],"originalAuth":{"authors":["zu Wied-Neuwied"],"year":{"year":"1830"}}},"details":{"species":{"genus":"Geositta","species":"poeciloptera","authorship":{"verbatim":"(zu Wied-Neuwied, 1830)","normalized":"(zu Wied-Neuwied 1830)","year":"1830","authors":["zu Wied-Neuwied"],"originalAuth":{"authors":["zu Wied-Neuwied"],"year":{"year":"1830"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":23,"end":25},{"wordType":"authorWord","start":26,"end":38},{"wordType":"year","start":40,"end":44}],"id":"c2abf205-a19a-5bf1-9a95-668101143dd8","parserVersion":"test_version"}
```

Name: Abacetus laevicollis de Chaudoir, 1869

Canonical: Abacetus laevicollis

Authorship: de Chaudoir 1869

```json
{"parsed":true,"parseQuality":1,"verbatim":"Abacetus laevicollis de Chaudoir, 1869","normalized":"Abacetus laevicollis de Chaudoir 1869","canonical":{"stemmed":"Abacetus laeuicoll","simple":"Abacetus laevicollis","full":"Abacetus laevicollis"},"cardinality":2,"authorship":{"verbatim":"de Chaudoir, 1869","normalized":"de Chaudoir 1869","year":"1869","authors":["de Chaudoir"],"originalAuth":{"authors":["de Chaudoir"],"year":{"year":"1869"}}},"details":{"species":{"genus":"Abacetus","species":"laevicollis","authorship":{"verbatim":"de Chaudoir, 1869","normalized":"de Chaudoir 1869","year":"1869","authors":["de Chaudoir"],"originalAuth":{"authors":["de Chaudoir"],"year":{"year":"1869"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":21,"end":23},{"wordType":"authorWord","start":24,"end":32},{"wordType":"year","start":34,"end":38}],"id":"8d81b939-695f-5a38-86c7-0f6efd1cacf3","parserVersion":"test_version"}
```

Name: Gastrosericus eremorum von Beaumont 1955

Canonical: Gastrosericus eremorum

Authorship: von Beaumont 1955

```json
{"parsed":true,"parseQuality":1,"verbatim":"Gastrosericus eremorum von Beaumont 1955","normalized":"Gastrosericus eremorum von Beaumont 1955","canonical":{"stemmed":"Gastrosericus eremor","simple":"Gastrosericus eremorum","full":"Gastrosericus eremorum"},"cardinality":2,"authorship":{"verbatim":"von Beaumont 1955","normalized":"von Beaumont 1955","year":"1955","authors":["von Beaumont"],"originalAuth":{"authors":["von Beaumont"],"year":{"year":"1955"}}},"details":{"species":{"genus":"Gastrosericus","species":"eremorum","authorship":{"verbatim":"von Beaumont 1955","normalized":"von Beaumont 1955","year":"1955","authors":["von Beaumont"],"originalAuth":{"authors":["von Beaumont"],"year":{"year":"1955"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":22},{"wordType":"authorWord","start":23,"end":26},{"wordType":"authorWord","start":27,"end":35},{"wordType":"year","start":36,"end":40}],"id":"98df7228-03ef-511c-9f2d-7f91e10c2af5","parserVersion":"test_version"}
```

Name: Agaricus squamula Berk. & M.A. Curtis 1860

Canonical: Agaricus squamula

Authorship: Berk. & M. A. Curtis 1860

```json
{"parsed":true,"parseQuality":1,"verbatim":"Agaricus squamula Berk. \u0026 M.A. Curtis 1860","normalized":"Agaricus squamula Berk. \u0026 M. A. Curtis 1860","canonical":{"stemmed":"Agaricus squamul","simple":"Agaricus squamula","full":"Agaricus squamula"},"cardinality":2,"authorship":{"verbatim":"Berk. \u0026 M.A. Curtis 1860","normalized":"Berk. \u0026 M. A. Curtis 1860","year":"1860","authors":["Berk.","M. A. Curtis"],"originalAuth":{"authors":["Berk.","M. A. Curtis"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Agaricus","species":"squamula","authorship":{"verbatim":"Berk. \u0026 M.A. Curtis 1860","normalized":"Berk. \u0026 M. A. Curtis 1860","year":"1860","authors":["Berk.","M. A. Curtis"],"originalAuth":{"authors":["Berk.","M. A. Curtis"],"year":{"year":"1860"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":23},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":28,"end":30},{"wordType":"authorWord","start":31,"end":37},{"wordType":"year","start":38,"end":42}],"id":"153b8745-887a-56ba-ad4a-69c10b0ad513","parserVersion":"test_version"}
```

Name: Peltula coriacea Büdel, Henssen & Wessels 1986

Canonical: Peltula coriacea

Authorship: Büdel, Henssen & Wessels 1986

```json
{"parsed":true,"parseQuality":1,"verbatim":"Peltula coriacea Büdel, Henssen \u0026 Wessels 1986","normalized":"Peltula coriacea Büdel, Henssen \u0026 Wessels 1986","canonical":{"stemmed":"Peltula coriace","simple":"Peltula coriacea","full":"Peltula coriacea"},"cardinality":2,"authorship":{"verbatim":"Büdel, Henssen \u0026 Wessels 1986","normalized":"Büdel, Henssen \u0026 Wessels 1986","year":"1986","authors":["Büdel","Henssen","Wessels"],"originalAuth":{"authors":["Büdel","Henssen","Wessels"],"year":{"year":"1986"}}},"details":{"species":{"genus":"Peltula","species":"coriacea","authorship":{"verbatim":"Büdel, Henssen \u0026 Wessels 1986","normalized":"Büdel, Henssen \u0026 Wessels 1986","year":"1986","authors":["Büdel","Henssen","Wessels"],"originalAuth":{"authors":["Büdel","Henssen","Wessels"],"year":{"year":"1986"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":16},{"wordType":"authorWord","start":17,"end":22},{"wordType":"authorWord","start":24,"end":31},{"wordType":"authorWord","start":34,"end":41},{"wordType":"year","start":42,"end":46}],"id":"081f5751-4042-597e-bccc-788754ce0248","parserVersion":"test_version"}
```

Name: Tuber liui A S. Xu 1999

Canonical: Tuber liui

Authorship: A S. Xu 1999

```json
{"parsed":true,"parseQuality":1,"verbatim":"Tuber liui A S. Xu 1999","normalized":"Tuber liui A S. Xu 1999","canonical":{"stemmed":"Tuber liu","simple":"Tuber liui","full":"Tuber liui"},"cardinality":2,"authorship":{"verbatim":"A S. Xu 1999","normalized":"A S. Xu 1999","year":"1999","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"],"year":{"year":"1999"}}},"details":{"species":{"genus":"Tuber","species":"liui","authorship":{"verbatim":"A S. Xu 1999","normalized":"A S. Xu 1999","year":"1999","authors":["A S. Xu"],"originalAuth":{"authors":["A S. Xu"],"year":{"year":"1999"}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":10},{"wordType":"authorWord","start":11,"end":12},{"wordType":"authorWord","start":13,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"year","start":19,"end":23}],"id":"4c79eb26-ae4c-5f4a-b5c5-07722ef1fa4f","parserVersion":"test_version"}
```

Name: Lecanora wetmorei Śliwa 2004

Canonical: Lecanora wetmorei

Authorship: Śliwa 2004

```json
{"parsed":true,"parseQuality":1,"verbatim":"Lecanora wetmorei Śliwa 2004","normalized":"Lecanora wetmorei Śliwa 2004","canonical":{"stemmed":"Lecanora wetmore","simple":"Lecanora wetmorei","full":"Lecanora wetmorei"},"cardinality":2,"authorship":{"verbatim":"Śliwa 2004","normalized":"Śliwa 2004","year":"2004","authors":["Śliwa"],"originalAuth":{"authors":["Śliwa"],"year":{"year":"2004"}}},"details":{"species":{"genus":"Lecanora","species":"wetmorei","authorship":{"verbatim":"Śliwa 2004","normalized":"Śliwa 2004","year":"2004","authors":["Śliwa"],"originalAuth":{"authors":["Śliwa"],"year":{"year":"2004"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":23},{"wordType":"year","start":24,"end":28}],"id":"50e874e9-f807-5446-a416-ca459475b1db","parserVersion":"test_version"}
```

Name: Vachonobisium troglophilum Vitali-di Castri, 1963

Canonical: Vachonobisium troglophilum

Authorship: Vitali-di Castri 1963

```json
{"parsed":true,"parseQuality":1,"verbatim":"Vachonobisium troglophilum Vitali-di Castri, 1963","normalized":"Vachonobisium troglophilum Vitali-di Castri 1963","canonical":{"stemmed":"Vachonobisium troglophil","simple":"Vachonobisium troglophilum","full":"Vachonobisium troglophilum"},"cardinality":2,"authorship":{"verbatim":"Vitali-di Castri, 1963","normalized":"Vitali-di Castri 1963","year":"1963","authors":["Vitali-di Castri"],"originalAuth":{"authors":["Vitali-di Castri"],"year":{"year":"1963"}}},"details":{"species":{"genus":"Vachonobisium","species":"troglophilum","authorship":{"verbatim":"Vitali-di Castri, 1963","normalized":"Vitali-di Castri 1963","year":"1963","authors":["Vitali-di Castri"],"originalAuth":{"authors":["Vitali-di Castri"],"year":{"year":"1963"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":26},{"wordType":"authorWord","start":27,"end":36},{"wordType":"authorWord","start":37,"end":43},{"wordType":"year","start":45,"end":49}],"id":"97424f96-2408-53b6-a6bf-a26613eec14c","parserVersion":"test_version"}
```

Name: Hyalesthes angustula Horvßth, 1909

Canonical: Hyalesthes angustula

Authorship: Horvßth 1909

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hyalesthes angustula Horvßth, 1909","normalized":"Hyalesthes angustula Horvßth 1909","canonical":{"stemmed":"Hyalesthes angustul","simple":"Hyalesthes angustula","full":"Hyalesthes angustula"},"cardinality":2,"authorship":{"verbatim":"Horvßth, 1909","normalized":"Horvßth 1909","year":"1909","authors":["Horvßth"],"originalAuth":{"authors":["Horvßth"],"year":{"year":"1909"}}},"details":{"species":{"genus":"Hyalesthes","species":"angustula","authorship":{"verbatim":"Horvßth, 1909","normalized":"Horvßth 1909","year":"1909","authors":["Horvßth"],"originalAuth":{"authors":["Horvßth"],"year":{"year":"1909"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":20},{"wordType":"authorWord","start":21,"end":28},{"wordType":"year","start":30,"end":34}],"id":"02058420-6623-5c22-b5ae-bc6a576f72fe","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl (1935h)

Canonical: Platypus bicaudatulus

Authorship: Schedl (1935)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Platypus bicaudatulus Schedl (1935h)","normalized":"Platypus bicaudatulus Schedl (1935)","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl (1935h)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl (1935h)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"approximateYear","start":30,"end":35}],"id":"5bf2e3f3-46dc-5138-a912-0e0ab2fdb22d","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl (1935)

Canonical: Platypus bicaudatulus

Authorship: Schedl (1935)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Platypus bicaudatulus Schedl (1935)","normalized":"Platypus bicaudatulus Schedl (1935)","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl (1935)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl (1935)","normalized":"Schedl (1935)","year":"(1935)","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"approximateYear","start":30,"end":34}],"id":"c13ffa95-76e8-5ad1-aec6-311d65dc4dc0","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl 1935

```json
{"parsed":true,"parseQuality":1,"verbatim":"Platypus bicaudatulus Schedl 1935","normalized":"Platypus bicaudatulus Schedl 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl 1935","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl 1935","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"year","start":29,"end":33}],"id":"d192a4f8-424f-5eba-affb-9855b153ff53","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl, 1935h

Canonical: Platypus bicaudatulus

Authorship: Schedl 1935

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"}],"verbatim":"Platypus bicaudatulus Schedl, 1935h","normalized":"Platypus bicaudatulus Schedl 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl, 1935h","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl, 1935h","normalized":"Schedl 1935","year":"1935","authors":["Schedl"],"originalAuth":{"authors":["Schedl"],"year":{"year":"1935"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"year","start":30,"end":35}],"id":"2f3b49aa-7d42-557b-9949-41df0e6059e8","parserVersion":"test_version"}
```

Name: Rotalina cultrata d'Orb. 1840

Canonical: Rotalina cultrata

Authorship: d'Orb. 1840

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rotalina cultrata d'Orb. 1840","normalized":"Rotalina cultrata d'Orb. 1840","canonical":{"stemmed":"Rotalina cultrat","simple":"Rotalina cultrata","full":"Rotalina cultrata"},"cardinality":2,"authorship":{"verbatim":"d'Orb. 1840","normalized":"d'Orb. 1840","year":"1840","authors":["d'Orb."],"originalAuth":{"authors":["d'Orb."],"year":{"year":"1840"}}},"details":{"species":{"genus":"Rotalina","species":"cultrata","authorship":{"verbatim":"d'Orb. 1840","normalized":"d'Orb. 1840","year":"1840","authors":["d'Orb."],"originalAuth":{"authors":["d'Orb."],"year":{"year":"1840"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":24},{"wordType":"year","start":25,"end":29}],"id":"085048a9-a6b8-525e-95ad-ae715b8c00ca","parserVersion":"test_version"}
```

Name: Stylosanthes guianensis (Aubl.) Sw. var. robusta L.'t Mannetje

Canonical: Stylosanthes guianensis var. robusta

Authorship: L. 't Mannetje

```json
{"parsed":true,"parseQuality":1,"verbatim":"Stylosanthes guianensis (Aubl.) Sw. var. robusta L.'t Mannetje","normalized":"Stylosanthes guianensis (Aubl.) Sw. var. robusta L. 't Mannetje","canonical":{"stemmed":"Stylosanthes guianens robust","simple":"Stylosanthes guianensis robusta","full":"Stylosanthes guianensis var. robusta"},"cardinality":3,"authorship":{"verbatim":"L.'t Mannetje","normalized":"L. 't Mannetje","authors":["L. 't Mannetje"],"originalAuth":{"authors":["L. 't Mannetje"]}},"details":{"infraSpecies":{"genus":"Stylosanthes","species":"guianensis","authorship":{"verbatim":"(Aubl.) Sw.","normalized":"(Aubl.) Sw.","authors":["Aubl.","Sw."],"originalAuth":{"authors":["Aubl."]},"combinationAuth":{"authors":["Sw."]}},"infraSpecies":[{"value":"robusta","rank":"var.","authorship":{"verbatim":"L.'t Mannetje","normalized":"L. 't Mannetje","authors":["L. 't Mannetje"],"originalAuth":{"authors":["L. 't Mannetje"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":30},{"wordType":"authorWord","start":32,"end":35},{"wordType":"rank","start":36,"end":40},{"wordType":"infraspecificEpithet","start":41,"end":48},{"wordType":"authorWord","start":49,"end":51},{"wordType":"authorWord","start":51,"end":53},{"wordType":"authorWord","start":54,"end":62}],"id":"fa16f59c-69a2-50cc-a4f6-bf4e8891eb9a","parserVersion":"test_version"}
```

Name: Doxander vittatus entropi (Man in 't Veld & Visser, 1993)

Canonical: Doxander vittatus entropi

Authorship: (Man ex 't Veld & Visser 1993)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Doxander vittatus entropi (Man in 't Veld \u0026 Visser, 1993)","normalized":"Doxander vittatus entropi (Man ex 't Veld \u0026 Visser 1993)","canonical":{"stemmed":"Doxander uittat entrop","simple":"Doxander vittatus entropi","full":"Doxander vittatus entropi"},"cardinality":3,"authorship":{"verbatim":"(Man in 't Veld \u0026 Visser, 1993)","normalized":"(Man ex 't Veld \u0026 Visser 1993)","authors":["Man"],"originalAuth":{"authors":["Man"],"exAuthors":{"authors":["'t Veld","Visser"],"year":{"year":"1993"}}}},"details":{"infraSpecies":{"genus":"Doxander","species":"vittatus","infraSpecies":[{"value":"entropi","authorship":{"verbatim":"(Man in 't Veld \u0026 Visser, 1993)","normalized":"(Man ex 't Veld \u0026 Visser 1993)","authors":["Man"],"originalAuth":{"authors":["Man"],"exAuthors":{"authors":["'t Veld","Visser"],"year":{"year":"1993"}}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"infraspecificEpithet","start":18,"end":25},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":34,"end":36},{"wordType":"authorWord","start":37,"end":41},{"wordType":"authorWord","start":44,"end":50},{"wordType":"year","start":52,"end":56}],"id":"1b3da2cb-82db-511d-86f5-4421966e3b65","parserVersion":"test_version"}
```

Name: Elaeagnus triflora Roxb. var. brevilimbatus E.'t Hart

Canonical: Elaeagnus triflora var. brevilimbatus

Authorship: E. 't Hart

```json
{"parsed":true,"parseQuality":1,"verbatim":"Elaeagnus triflora Roxb. var. brevilimbatus E.'t Hart","normalized":"Elaeagnus triflora Roxb. var. brevilimbatus E. 't Hart","canonical":{"stemmed":"Elaeagnus triflor breuilimbat","simple":"Elaeagnus triflora brevilimbatus","full":"Elaeagnus triflora var. brevilimbatus"},"cardinality":3,"authorship":{"verbatim":"E.'t Hart","normalized":"E. 't Hart","authors":["E. 't Hart"],"originalAuth":{"authors":["E. 't Hart"]}},"details":{"infraSpecies":{"genus":"Elaeagnus","species":"triflora","authorship":{"verbatim":"Roxb.","normalized":"Roxb.","authors":["Roxb."],"originalAuth":{"authors":["Roxb."]}},"infraSpecies":[{"value":"brevilimbatus","rank":"var.","authorship":{"verbatim":"E.'t Hart","normalized":"E. 't Hart","authors":["E. 't Hart"],"originalAuth":{"authors":["E. 't Hart"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":18},{"wordType":"authorWord","start":19,"end":24},{"wordType":"rank","start":25,"end":29},{"wordType":"infraspecificEpithet","start":30,"end":43},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":46,"end":48},{"wordType":"authorWord","start":49,"end":53}],"id":"e3b3f47c-856a-5c21-bfa7-ac8c89453232","parserVersion":"test_version"}
```

Name: Laevistrombus guidoi (Man in't Veld & De Turck, 1998)

Canonical: Laevistrombus guidoi

Authorship: (Man in't Veld & De Turck 1998)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Laevistrombus guidoi (Man in't Veld \u0026 De Turck, 1998)","normalized":"Laevistrombus guidoi (Man in't Veld \u0026 De Turck 1998)","canonical":{"stemmed":"Laevistrombus guido","simple":"Laevistrombus guidoi","full":"Laevistrombus guidoi"},"cardinality":2,"authorship":{"verbatim":"(Man in't Veld \u0026 De Turck, 1998)","normalized":"(Man in't Veld \u0026 De Turck 1998)","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Laevistrombus","species":"guidoi","authorship":{"verbatim":"(Man in't Veld \u0026 De Turck, 1998)","normalized":"(Man in't Veld \u0026 De Turck 1998)","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":20},{"wordType":"authorWord","start":22,"end":25},{"wordType":"authorWord","start":26,"end":30},{"wordType":"authorWord","start":31,"end":35},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":41,"end":46},{"wordType":"year","start":48,"end":52}],"id":"e3ff94a0-92d0-5894-8599-f288e92077c8","parserVersion":"test_version"}
```

Name: Strombus guidoi Man in't Veld & De Turck, 1998

Canonical: Strombus guidoi

Authorship: Man in't Veld & De Turck 1998

```json
{"parsed":true,"parseQuality":1,"verbatim":"Strombus guidoi Man in't Veld \u0026 De Turck, 1998","normalized":"Strombus guidoi Man in't Veld \u0026 De Turck 1998","canonical":{"stemmed":"Strombus guido","simple":"Strombus guidoi","full":"Strombus guidoi"},"cardinality":2,"authorship":{"verbatim":"Man in't Veld \u0026 De Turck, 1998","normalized":"Man in't Veld \u0026 De Turck 1998","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Strombus","species":"guidoi","authorship":{"verbatim":"Man in't Veld \u0026 De Turck, 1998","normalized":"Man in't Veld \u0026 De Turck 1998","year":"1998","authors":["Man in't Veld","De Turck"],"originalAuth":{"authors":["Man in't Veld","De Turck"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"authorWord","start":16,"end":19},{"wordType":"authorWord","start":20,"end":24},{"wordType":"authorWord","start":25,"end":29},{"wordType":"authorWord","start":32,"end":34},{"wordType":"authorWord","start":35,"end":40},{"wordType":"year","start":42,"end":46}],"id":"100d3b6e-62d3-51ad-baf6-60408babc574","parserVersion":"test_version"}
```

Name: Strombus vittatus entropi Man in't Veld & Visser, 1993

Canonical: Strombus vittatus entropi

Authorship: Man in't Veld & Visser 1993

```json
{"parsed":true,"parseQuality":1,"verbatim":"Strombus vittatus entropi Man in't Veld \u0026 Visser, 1993","normalized":"Strombus vittatus entropi Man in't Veld \u0026 Visser 1993","canonical":{"stemmed":"Strombus uittat entrop","simple":"Strombus vittatus entropi","full":"Strombus vittatus entropi"},"cardinality":3,"authorship":{"verbatim":"Man in't Veld \u0026 Visser, 1993","normalized":"Man in't Veld \u0026 Visser 1993","year":"1993","authors":["Man in't Veld","Visser"],"originalAuth":{"authors":["Man in't Veld","Visser"],"year":{"year":"1993"}}},"details":{"infraSpecies":{"genus":"Strombus","species":"vittatus","infraSpecies":[{"value":"entropi","authorship":{"verbatim":"Man in't Veld \u0026 Visser, 1993","normalized":"Man in't Veld \u0026 Visser 1993","year":"1993","authors":["Man in't Veld","Visser"],"originalAuth":{"authors":["Man in't Veld","Visser"],"year":{"year":"1993"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"infraspecificEpithet","start":18,"end":25},{"wordType":"authorWord","start":26,"end":29},{"wordType":"authorWord","start":30,"end":34},{"wordType":"authorWord","start":35,"end":39},{"wordType":"authorWord","start":42,"end":48},{"wordType":"year","start":50,"end":54}],"id":"c74691e3-0f71-576b-81ea-6173bdae9817","parserVersion":"test_version"}
```

Name: Velutina haliotoides (Linnaeus, 1758),

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758),","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":",","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":22,"end":30},{"wordType":"year","start":32,"end":36}],"id":"59093ba7-64a1-53c4-9795-12de7ff9e718","parserVersion":"test_version"}
```

Name: Hennediella microphylla (R.Br.bis) Paris

Canonical: Hennediella microphylla

Authorship: (R. Br. bis) Paris

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hennediella microphylla (R.Br.bis) Paris","normalized":"Hennediella microphylla (R. Br. bis) Paris","canonical":{"stemmed":"Hennediella microphyll","simple":"Hennediella microphylla","full":"Hennediella microphylla"},"cardinality":2,"authorship":{"verbatim":"(R.Br.bis) Paris","normalized":"(R. Br. bis) Paris","authors":["R. Br. bis","Paris"],"originalAuth":{"authors":["R. Br. bis"]},"combinationAuth":{"authors":["Paris"]}},"details":{"species":{"genus":"Hennediella","species":"microphylla","authorship":{"verbatim":"(R.Br.bis) Paris","normalized":"(R. Br. bis) Paris","authors":["R. Br. bis","Paris"],"originalAuth":{"authors":["R. Br. bis"]},"combinationAuth":{"authors":["Paris"]}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":23},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":30,"end":33},{"wordType":"authorWord","start":35,"end":40}],"id":"e8cc6d9d-6e6c-53a1-99a9-59f636009ed0","parserVersion":"test_version"}
```

### Binomials with an abbreviated genus

Name: M. alpium

Canonical: M. alpium

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"}],"verbatim":"M. alpium","normalized":"M. alpium","canonical":{"stemmed":"M. alpi","simple":"M. alpium","full":"M. alpium"},"cardinality":2,"details":{"species":{"genus":"M.","species":"alpium"}},"pos":[{"wordType":"genus","start":0,"end":2},{"wordType":"specificEpithet","start":3,"end":9}],"id":"9001ffb5-eac2-5bb4-8f78-d7b7e3e02bd8","parserVersion":"test_version"}
```

Name: Mo. alpium (Osbeck, 1778)

Canonical: Mo. alpium

Authorship: (Osbeck 1778)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"}],"verbatim":"Mo. alpium (Osbeck, 1778)","normalized":"Mo. alpium (Osbeck 1778)","canonical":{"stemmed":"Mo. alpi","simple":"Mo. alpium","full":"Mo. alpium"},"cardinality":2,"authorship":{"verbatim":"(Osbeck, 1778)","normalized":"(Osbeck 1778)","year":"1778","authors":["Osbeck"],"originalAuth":{"authors":["Osbeck"],"year":{"year":"1778"}}},"details":{"species":{"genus":"Mo.","species":"alpium","authorship":{"verbatim":"(Osbeck, 1778)","normalized":"(Osbeck 1778)","year":"1778","authors":["Osbeck"],"originalAuth":{"authors":["Osbeck"],"year":{"year":"1778"}}}}},"pos":[{"wordType":"genus","start":0,"end":3},{"wordType":"specificEpithet","start":4,"end":10},{"wordType":"authorWord","start":12,"end":18},{"wordType":"year","start":20,"end":24}],"id":"1e9437b7-bf45-5b12-8da0-8966c6ea1c5c","parserVersion":"test_version"}
```

### Binomials with several authours

Name: Nemcia epacridoides (Meissner)Crisp

Canonical: Nemcia epacridoides

Authorship: (Meissner) Crisp

```json
{"parsed":true,"parseQuality":1,"verbatim":"Nemcia epacridoides (Meissner)Crisp","normalized":"Nemcia epacridoides (Meissner) Crisp","canonical":{"stemmed":"Nemcia epacridoid","simple":"Nemcia epacridoides","full":"Nemcia epacridoides"},"cardinality":2,"authorship":{"verbatim":"(Meissner)Crisp","normalized":"(Meissner) Crisp","authors":["Meissner","Crisp"],"originalAuth":{"authors":["Meissner"]},"combinationAuth":{"authors":["Crisp"]}},"details":{"species":{"genus":"Nemcia","species":"epacridoides","authorship":{"verbatim":"(Meissner)Crisp","normalized":"(Meissner) Crisp","authors":["Meissner","Crisp"],"originalAuth":{"authors":["Meissner"]},"combinationAuth":{"authors":["Crisp"]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":19},{"wordType":"authorWord","start":21,"end":29},{"wordType":"authorWord","start":30,"end":35}],"id":"6ea9d43f-33c1-5bed-b9a9-edb164966eb6","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh & W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora dendrobii Goh \u0026 W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh \u0026 W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh \u0026 W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":35,"end":37},{"wordType":"authorWord","start":38,"end":43},{"wordType":"year","start":44,"end":48}],"id":"988fd6ba-0221-5b62-a041-fb81addc4465","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh and W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora dendrobii Goh and W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh and W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh and W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":35,"end":37},{"wordType":"authorWord","start":37,"end":39},{"wordType":"authorWord","start":40,"end":45},{"wordType":"year","start":46,"end":50}],"id":"4d701dca-8774-5a5e-9378-11f60c0e735c","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii Goh et W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh & W. H. Hsieh 1990

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora dendrobii Goh et W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh \u0026 W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh et W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh et W.H. Hsieh 1990","normalized":"Goh \u0026 W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":34,"end":36},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":39,"end":44},{"wordType":"year","start":45,"end":49}],"id":"13175b62-b95b-53b7-8d88-1be6fca794ec","parserVersion":"test_version"}
```

Name: Schottera nicaeënsis (J.V. Lamouroux ex Duby) Guiry & Hollenberg

Canonical: Schottera nicaeensis

Authorship: (J. V. Lamouroux ex Duby) Guiry & Hollenberg

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Schottera nicaeënsis (J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"Schottera nicaeensis (J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","canonical":{"stemmed":"Schottera nicaeens","simple":"Schottera nicaeensis","full":"Schottera nicaeensis"},"cardinality":2,"authorship":{"verbatim":"(J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"(J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","authors":["J. V. Lamouroux","Guiry","Hollenberg"],"originalAuth":{"authors":["J. V. Lamouroux"],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Guiry","Hollenberg"]}},"details":{"species":{"genus":"Schottera","species":"nicaeensis","authorship":{"verbatim":"(J.V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","normalized":"(J. V. Lamouroux ex Duby) Guiry \u0026 Hollenberg","authors":["J. V. Lamouroux","Guiry","Hollenberg"],"originalAuth":{"authors":["J. V. Lamouroux"],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Guiry","Hollenberg"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWord","start":24,"end":26},{"wordType":"authorWord","start":27,"end":36},{"wordType":"authorWord","start":40,"end":44},{"wordType":"authorWord","start":46,"end":51},{"wordType":"authorWord","start":54,"end":64}],"id":"ffeb3703-63e5-5ff3-b296-582c0c3a3373","parserVersion":"test_version"}
```

### Binomials with several authors and a year

Name: Cladoniicola staurospora Diederich, van den Boom & Aptroot 2001

Canonical: Cladoniicola staurospora

Authorship: Diederich, van den Boom & Aptroot 2001

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cladoniicola staurospora Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Cladoniicola staurospora Diederich, van den Boom \u0026 Aptroot 2001","canonical":{"stemmed":"Cladoniicola staurospor","simple":"Cladoniicola staurospora","full":"Cladoniicola staurospora"},"cardinality":2,"authorship":{"verbatim":"Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Diederich, van den Boom \u0026 Aptroot 2001","year":"2001","authors":["Diederich","van den Boom","Aptroot"],"originalAuth":{"authors":["Diederich","van den Boom","Aptroot"],"year":{"year":"2001"}}},"details":{"species":{"genus":"Cladoniicola","species":"staurospora","authorship":{"verbatim":"Diederich, van den Boom \u0026 Aptroot 2001","normalized":"Diederich, van den Boom \u0026 Aptroot 2001","year":"2001","authors":["Diederich","van den Boom","Aptroot"],"originalAuth":{"authors":["Diederich","van den Boom","Aptroot"],"year":{"year":"2001"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":24},{"wordType":"authorWord","start":25,"end":34},{"wordType":"authorWord","start":36,"end":39},{"wordType":"authorWord","start":40,"end":43},{"wordType":"authorWord","start":44,"end":48},{"wordType":"authorWord","start":51,"end":58},{"wordType":"year","start":59,"end":63}],"id":"e59e3b01-311d-5dda-88e7-7e821440f5ee","parserVersion":"test_version"}
```

Name: Stagonospora polyspora M.T. Lucas & Sousa da Câmara 1934

Canonical: Stagonospora polyspora

Authorship: M. T. Lucas & Sousa da Câmara 1934

```json
{"parsed":true,"parseQuality":1,"verbatim":"Stagonospora polyspora M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"Stagonospora polyspora M. T. Lucas \u0026 Sousa da Câmara 1934","canonical":{"stemmed":"Stagonospora polyspor","simple":"Stagonospora polyspora","full":"Stagonospora polyspora"},"cardinality":2,"authorship":{"verbatim":"M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}},"details":{"species":{"genus":"Stagonospora","species":"polyspora","authorship":{"verbatim":"M.T. Lucas \u0026 Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22},{"wordType":"authorWord","start":23,"end":25},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":28,"end":33},{"wordType":"authorWord","start":36,"end":41},{"wordType":"authorWord","start":42,"end":44},{"wordType":"authorWord","start":45,"end":51},{"wordType":"year","start":52,"end":56}],"id":"f03d53d7-2db1-591f-8727-6b77c0af2e0c","parserVersion":"test_version"}
```

Name: Stagonospora polyspora M.T. Lucas et Sousa da Câmara 1934

Canonical: Stagonospora polyspora

Authorship: M. T. Lucas & Sousa da Câmara 1934

```json
{"parsed":true,"parseQuality":1,"verbatim":"Stagonospora polyspora M.T. Lucas et Sousa da Câmara 1934","normalized":"Stagonospora polyspora M. T. Lucas \u0026 Sousa da Câmara 1934","canonical":{"stemmed":"Stagonospora polyspor","simple":"Stagonospora polyspora","full":"Stagonospora polyspora"},"cardinality":2,"authorship":{"verbatim":"M.T. Lucas et Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}},"details":{"species":{"genus":"Stagonospora","species":"polyspora","authorship":{"verbatim":"M.T. Lucas et Sousa da Câmara 1934","normalized":"M. T. Lucas \u0026 Sousa da Câmara 1934","year":"1934","authors":["M. T. Lucas","Sousa da Câmara"],"originalAuth":{"authors":["M. T. Lucas","Sousa da Câmara"],"year":{"year":"1934"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22},{"wordType":"authorWord","start":23,"end":25},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":28,"end":33},{"wordType":"authorWord","start":37,"end":42},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":52},{"wordType":"year","start":53,"end":57}],"id":"a8a48393-0ca9-5916-83e3-fb32b7b0c422","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii U. Braun & Crous 2003

Canonical: Pseudocercospora dendrobii

Authorship: U. Braun & Crous 2003

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora dendrobii U. Braun \u0026 Crous 2003","normalized":"Pseudocercospora dendrobii U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"U. Braun \u0026 Crous 2003","normalized":"U. Braun \u0026 Crous 2003","year":"2003","authors":["U. Braun","Crous"],"originalAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"U. Braun \u0026 Crous 2003","normalized":"U. Braun \u0026 Crous 2003","year":"2003","authors":["U. Braun","Crous"],"originalAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":29},{"wordType":"authorWord","start":30,"end":35},{"wordType":"authorWord","start":38,"end":43},{"wordType":"year","start":44,"end":48}],"id":"afd958fc-82a5-5551-951b-a725a49d3df0","parserVersion":"test_version"}
```

Name: Abaxisotima acuminata (Wang, Yuwen & Xiangwei Liu 1996)

Canonical: Abaxisotima acuminata

Authorship: (Wang, Yuwen & Xiangwei Liu 1996)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Abaxisotima acuminata (Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"Abaxisotima acuminata (Wang, Yuwen \u0026 Xiangwei Liu 1996)","canonical":{"stemmed":"Abaxisotima acuminat","simple":"Abaxisotima acuminata","full":"Abaxisotima acuminata"},"cardinality":2,"authorship":{"verbatim":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","year":"1996","authors":["Wang","Yuwen","Xiangwei Liu"],"originalAuth":{"authors":["Wang","Yuwen","Xiangwei Liu"],"year":{"year":"1996"}}},"details":{"species":{"genus":"Abaxisotima","species":"acuminata","authorship":{"verbatim":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","normalized":"(Wang, Yuwen \u0026 Xiangwei Liu 1996)","year":"1996","authors":["Wang","Yuwen","Xiangwei Liu"],"originalAuth":{"authors":["Wang","Yuwen","Xiangwei Liu"],"year":{"year":"1996"}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":23,"end":27},{"wordType":"authorWord","start":29,"end":34},{"wordType":"authorWord","start":37,"end":45},{"wordType":"authorWord","start":46,"end":49},{"wordType":"year","start":50,"end":54}],"id":"5eecff7d-181c-508c-832d-df4619b8b027","parserVersion":"test_version"}
```

Name: Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi & L. Tang, 2009

Canonical: Aboilomimus sichuanensis ornatus

Authorship: Liu, Xiang-wei, M. Zhou, W Bi & L. Tang 2009

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Aboilomimus sichuanensis ornatus Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","canonical":{"stemmed":"Aboilomimus sichuanens ornat","simple":"Aboilomimus sichuanensis ornatus","full":"Aboilomimus sichuanensis ornatus"},"cardinality":3,"authorship":{"verbatim":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","year":"2009","authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"originalAuth":{"authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"year":{"year":"2009"}}},"details":{"infraSpecies":{"genus":"Aboilomimus","species":"sichuanensis","infraSpecies":[{"value":"ornatus","authorship":{"verbatim":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang, 2009","normalized":"Liu, Xiang-wei, M. Zhou, W Bi \u0026 L. Tang 2009","year":"2009","authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"originalAuth":{"authors":["Liu","Xiang-wei","M. Zhou","W Bi","L. Tang"],"year":{"year":"2009"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":32},{"wordType":"authorWord","start":33,"end":36},{"wordType":"authorWord","start":38,"end":47},{"wordType":"authorWord","start":49,"end":51},{"wordType":"authorWord","start":52,"end":56},{"wordType":"authorWord","start":58,"end":59},{"wordType":"authorWord","start":60,"end":62},{"wordType":"authorWord","start":65,"end":67},{"wordType":"authorWord","start":68,"end":72},{"wordType":"year","start":74,"end":78}],"id":"25ac4ba8-6595-5ab3-8463-f99f738bf4e4","parserVersion":"test_version"}
```

### Binomials with basionym and combination authors

Name: Yarrowia lipolytica var. lipolytica (Wick., Kurtzman & E.A. Herrm.) Van der Walt & Arx 1981

Canonical: Yarrowia lipolytica var. lipolytica

Authorship: (Wick., Kurtzman & E. A. Herrm.) Van der Walt & Arx 1981

```json
{"parsed":true,"parseQuality":1,"verbatim":"Yarrowia lipolytica var. lipolytica (Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"Yarrowia lipolytica var. lipolytica (Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","canonical":{"stemmed":"Yarrowia lipolytic lipolytic","simple":"Yarrowia lipolytica lipolytica","full":"Yarrowia lipolytica var. lipolytica"},"cardinality":3,"authorship":{"verbatim":"(Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"(Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","authors":["Wick.","Kurtzman","E. A. Herrm.","Van der Walt","Arx"],"originalAuth":{"authors":["Wick.","Kurtzman","E. A. Herrm."]},"combinationAuth":{"authors":["Van der Walt","Arx"],"year":{"year":"1981"}}},"details":{"infraSpecies":{"genus":"Yarrowia","species":"lipolytica","infraSpecies":[{"value":"lipolytica","rank":"var.","authorship":{"verbatim":"(Wick., Kurtzman \u0026 E.A. Herrm.) Van der Walt \u0026 Arx 1981","normalized":"(Wick., Kurtzman \u0026 E. A. Herrm.) Van der Walt \u0026 Arx 1981","authors":["Wick.","Kurtzman","E. A. Herrm.","Van der Walt","Arx"],"originalAuth":{"authors":["Wick.","Kurtzman","E. A. Herrm."]},"combinationAuth":{"authors":["Van der Walt","Arx"],"year":{"year":"1981"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":35},{"wordType":"authorWord","start":37,"end":42},{"wordType":"authorWord","start":44,"end":52},{"wordType":"authorWord","start":55,"end":57},{"wordType":"authorWord","start":57,"end":59},{"wordType":"authorWord","start":60,"end":66},{"wordType":"authorWord","start":68,"end":71},{"wordType":"authorWord","start":72,"end":75},{"wordType":"authorWord","start":76,"end":80},{"wordType":"authorWord","start":83,"end":86},{"wordType":"year","start":87,"end":91}],"id":"e649d828-0ae9-5b5b-b079-1485c9bbf872","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett)U. Braun & Crous     2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett) U. Braun & Crous 2003

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett) U. Braun \u0026 Crous 2003","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"]},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett) U. Braun \u0026 Crous 2003","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"]},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":29},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":36,"end":43},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":47,"end":52},{"wordType":"authorWord","start":55,"end":60},{"wordType":"year","start":65,"end":69}],"id":"3c52bc21-3ac9-5be4-9d5f-1f84fe9d3325","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett, 1873)U. Braun & Crous     2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett 1873) U. Braun & Crous 2003

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett 1873) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett, 1873)U. Braun \u0026 Crous     2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":29},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":36,"end":43},{"wordType":"year","start":45,"end":49},{"wordType":"authorWord","start":50,"end":52},{"wordType":"authorWord","start":53,"end":58},{"wordType":"authorWord","start":61,"end":66},{"wordType":"year","start":71,"end":75}],"id":"8e5dd168-d7f1-51e4-989c-cedb253d572c","parserVersion":"test_version"}
```

Name: Pseudocercospora dendrobii(H.C.     Burnett 1873)U. Braun & Crous ,    2003

Canonical: Pseudocercospora dendrobii

Authorship: (H. C. Burnett 1873) U. Braun & Crous 2003

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Pseudocercospora dendrobii(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"Pseudocercospora dendrobii (H. C. Burnett 1873) U. Braun \u0026 Crous 2003","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"(H.C.     Burnett 1873)U. Braun \u0026 Crous ,    2003","normalized":"(H. C. Burnett 1873) U. Braun \u0026 Crous 2003","year":"1873","authors":["H. C. Burnett","U. Braun","Crous"],"originalAuth":{"authors":["H. C. Burnett"],"year":{"year":"1873"}},"combinationAuth":{"authors":["U. Braun","Crous"],"year":{"year":"2003"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":29},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":36,"end":43},{"wordType":"year","start":44,"end":48},{"wordType":"authorWord","start":49,"end":51},{"wordType":"authorWord","start":52,"end":57},{"wordType":"authorWord","start":60,"end":65},{"wordType":"year","start":71,"end":75}],"id":"a35b47c6-6716-5750-ab81-a19aed44143b","parserVersion":"test_version"}
```

Name: Sedella pumila (Benth.) Britton & Rose

Canonical: Sedella pumila

Authorship: (Benth.) Britton & Rose

```json
{"parsed":true,"parseQuality":1,"verbatim":"Sedella pumila (Benth.) Britton \u0026 Rose","normalized":"Sedella pumila (Benth.) Britton \u0026 Rose","canonical":{"stemmed":"Sedella pumil","simple":"Sedella pumila","full":"Sedella pumila"},"cardinality":2,"authorship":{"verbatim":"(Benth.) Britton \u0026 Rose","normalized":"(Benth.) Britton \u0026 Rose","authors":["Benth.","Britton","Rose"],"originalAuth":{"authors":["Benth."]},"combinationAuth":{"authors":["Britton","Rose"]}},"details":{"species":{"genus":"Sedella","species":"pumila","authorship":{"verbatim":"(Benth.) Britton \u0026 Rose","normalized":"(Benth.) Britton \u0026 Rose","authors":["Benth.","Britton","Rose"],"originalAuth":{"authors":["Benth."]},"combinationAuth":{"authors":["Britton","Rose"]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":14},{"wordType":"authorWord","start":16,"end":22},{"wordType":"authorWord","start":24,"end":31},{"wordType":"authorWord","start":34,"end":38}],"id":"393cedba-6ff1-5e5c-83f0-21e32f031ab7","parserVersion":"test_version"}
```

Name: Impatiens nomenyae Eb.Fisch. & Raheliv.

Canonical: Impatiens nomenyae

Authorship: Eb. Fisch. & Raheliv.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Impatiens nomenyae Eb.Fisch. \u0026 Raheliv.","normalized":"Impatiens nomenyae Eb. Fisch. \u0026 Raheliv.","canonical":{"stemmed":"Impatiens nomeny","simple":"Impatiens nomenyae","full":"Impatiens nomenyae"},"cardinality":2,"authorship":{"verbatim":"Eb.Fisch. \u0026 Raheliv.","normalized":"Eb. Fisch. \u0026 Raheliv.","authors":["Eb. Fisch.","Raheliv."],"originalAuth":{"authors":["Eb. Fisch.","Raheliv."]}},"details":{"species":{"genus":"Impatiens","species":"nomenyae","authorship":{"verbatim":"Eb.Fisch. \u0026 Raheliv.","normalized":"Eb. Fisch. \u0026 Raheliv.","authors":["Eb. Fisch.","Raheliv."],"originalAuth":{"authors":["Eb. Fisch.","Raheliv."]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":18},{"wordType":"authorWord","start":19,"end":22},{"wordType":"authorWord","start":22,"end":28},{"wordType":"authorWord","start":31,"end":39}],"id":"6452d4ac-738b-5773-8d69-50232e2842a1","parserVersion":"test_version"}
```

Name: Armeria carpetana ssp. carpetana H. del Villar

Canonical: Armeria carpetana subsp. carpetana

Authorship: H. del Villar

```json
{"parsed":true,"parseQuality":1,"verbatim":"Armeria carpetana ssp. carpetana H. del Villar","normalized":"Armeria carpetana subsp. carpetana H. del Villar","canonical":{"stemmed":"Armeria carpetan carpetan","simple":"Armeria carpetana carpetana","full":"Armeria carpetana subsp. carpetana"},"cardinality":3,"authorship":{"verbatim":"H. del Villar","normalized":"H. del Villar","authors":["H. del Villar"],"originalAuth":{"authors":["H. del Villar"]}},"details":{"infraSpecies":{"genus":"Armeria","species":"carpetana","infraSpecies":[{"value":"carpetana","rank":"subsp.","authorship":{"verbatim":"H. del Villar","normalized":"H. del Villar","authors":["H. del Villar"],"originalAuth":{"authors":["H. del Villar"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"rank","start":18,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":32},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":36,"end":39},{"wordType":"authorWord","start":40,"end":46}],"id":"4b16116e-549d-56bf-959a-ff11edb25021","parserVersion":"test_version"}
```

### Infraspecies without rank (ICZN)

Name: Peristernia nassatula forskali Tapparone-Canefri 1875

Canonical: Peristernia nassatula forskali

Authorship: Tapparone-Canefri 1875

```json
{"parsed":true,"parseQuality":1,"verbatim":"Peristernia nassatula forskali Tapparone-Canefri 1875","normalized":"Peristernia nassatula forskali Tapparone-Canefri 1875","canonical":{"stemmed":"Peristernia nassatul forskal","simple":"Peristernia nassatula forskali","full":"Peristernia nassatula forskali"},"cardinality":3,"authorship":{"verbatim":"Tapparone-Canefri 1875","normalized":"Tapparone-Canefri 1875","year":"1875","authors":["Tapparone-Canefri"],"originalAuth":{"authors":["Tapparone-Canefri"],"year":{"year":"1875"}}},"details":{"infraSpecies":{"genus":"Peristernia","species":"nassatula","infraSpecies":[{"value":"forskali","authorship":{"verbatim":"Tapparone-Canefri 1875","normalized":"Tapparone-Canefri 1875","year":"1875","authors":["Tapparone-Canefri"],"originalAuth":{"authors":["Tapparone-Canefri"],"year":{"year":"1875"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"infraspecificEpithet","start":22,"end":30},{"wordType":"authorWord","start":31,"end":48},{"wordType":"year","start":49,"end":53}],"id":"5aa39b53-32ee-5e9f-aa29-c268a9662fd7","parserVersion":"test_version"}
```

Name: Cypraeovula (Luponia) amphithales perdentata

Canonical: Cypraeovula amphithales perdentata

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cypraeovula (Luponia) amphithales perdentata","normalized":"Cypraeovula (Luponia) amphithales perdentata","canonical":{"stemmed":"Cypraeovula amphithal perdentat","simple":"Cypraeovula amphithales perdentata","full":"Cypraeovula amphithales perdentata"},"cardinality":3,"details":{"infraSpecies":{"genus":"Cypraeovula","subGenus":"Luponia","species":"amphithales","infraSpecies":[{"value":"perdentata"}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"infragenericEpithet","start":13,"end":20},{"wordType":"specificEpithet","start":22,"end":33},{"wordType":"infraspecificEpithet","start":34,"end":44}],"id":"d05be4e3-a0e3-5af4-9104-7922df1bcb47","parserVersion":"test_version"}
```

Name: Triticum repens vulgäre

Canonical: Triticum repens vulgaere

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Triticum repens vulgäre","normalized":"Triticum repens vulgaere","canonical":{"stemmed":"Triticum repens uulgaer","simple":"Triticum repens vulgaere","full":"Triticum repens vulgaere"},"cardinality":3,"details":{"infraSpecies":{"genus":"Triticum","species":"repens","infraSpecies":[{"value":"vulgaere"}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"infraspecificEpithet","start":16,"end":23}],"id":"5fb6ae9c-d7be-5d81-88b8-3c96d4c48a74","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Batsch) K. A. Harrison 1961

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961","normalized":"Hydnellum scrobiculatum zonatum (Batsch) K. A. Harrison 1961","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Batsch) K. A. Harrison 1961","normalized":"(Batsch) K. A. Harrison 1961","authors":["Batsch","K. A. Harrison"],"originalAuth":{"authors":["Batsch"]},"combinationAuth":{"authors":["K. A. Harrison"],"year":{"year":"1961"}}},"details":{"infraSpecies":{"genus":"Hydnellum","species":"scrobiculatum","infraSpecies":[{"value":"zonatum","authorship":{"verbatim":"(Batsch) K. A. Harrison 1961","normalized":"(Batsch) K. A. Harrison 1961","authors":["Batsch","K. A. Harrison"],"originalAuth":{"authors":["Batsch"]},"combinationAuth":{"authors":["K. A. Harrison"],"year":{"year":"1961"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":23},{"wordType":"infraspecificEpithet","start":24,"end":31},{"wordType":"authorWord","start":33,"end":39},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":47,"end":55},{"wordType":"year","start":56,"end":60}],"id":"8368c11a-7c1b-5e82-bdad-a4887bfa81d2","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hydnellum scrobiculatum zonatum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum scrobiculatum zonatum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraSpecies":{"genus":"Hydnellum","species":"scrobiculatum","infraSpecies":[{"value":"zonatum","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":23},{"wordType":"infraspecificEpithet","start":24,"end":31},{"wordType":"authorWord","start":33,"end":39},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":44,"end":48},{"wordType":"authorWord","start":51,"end":53},{"wordType":"authorWord","start":53,"end":55},{"wordType":"authorWord","start":56,"end":62},{"wordType":"year","start":63,"end":67}],"id":"fa3448c6-168e-575f-a6eb-c5adc6f3e89d","parserVersion":"test_version"}
```

Name: Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum zonatum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum (Hydnellum) scrobiculatum zonatum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraSpecies":{"genus":"Hydnellum","subGenus":"Hydnellum","species":"scrobiculatum","infraSpecies":[{"value":"zonatum","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"infragenericEpithet","start":11,"end":20},{"wordType":"specificEpithet","start":22,"end":35},{"wordType":"infraspecificEpithet","start":36,"end":43},{"wordType":"authorWord","start":45,"end":51},{"wordType":"authorWord","start":53,"end":55},{"wordType":"authorWord","start":56,"end":60},{"wordType":"authorWord","start":63,"end":65},{"wordType":"authorWord","start":65,"end":67},{"wordType":"authorWord","start":68,"end":74},{"wordType":"year","start":75,"end":79}],"id":"14e5eb1f-82a3-598c-9ada-3a9a20ab54cc","parserVersion":"test_version"}
```

Name: Hydnellum scrobiculatum zonatum

Canonical: Hydnellum scrobiculatum zonatum

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hydnellum scrobiculatum zonatum","normalized":"Hydnellum scrobiculatum zonatum","canonical":{"stemmed":"Hydnellum scrobiculat zonat","simple":"Hydnellum scrobiculatum zonatum","full":"Hydnellum scrobiculatum zonatum"},"cardinality":3,"details":{"infraSpecies":{"genus":"Hydnellum","species":"scrobiculatum","infraSpecies":[{"value":"zonatum"}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":23},{"wordType":"infraspecificEpithet","start":24,"end":31}],"id":"22af845f-773e-502e-be46-ac73ae5960be","parserVersion":"test_version"}
```

Name: Mus musculus hortulanus

Canonical: Mus musculus hortulanus

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Mus musculus hortulanus","normalized":"Mus musculus hortulanus","canonical":{"stemmed":"Mus muscul hortulan","simple":"Mus musculus hortulanus","full":"Mus musculus hortulanus"},"cardinality":3,"details":{"infraSpecies":{"genus":"Mus","species":"musculus","infraSpecies":[{"value":"hortulanus"}]}},"pos":[{"wordType":"genus","start":0,"end":3},{"wordType":"specificEpithet","start":4,"end":12},{"wordType":"infraspecificEpithet","start":13,"end":23}],"id":"5fd9a4aa-9fa8-5200-909a-6c9ec8a9a088","parserVersion":"test_version"}
```

Name: Ortygospiza atricollis mülleri

Canonical: Ortygospiza atricollis muelleri

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Ortygospiza atricollis mülleri","normalized":"Ortygospiza atricollis muelleri","canonical":{"stemmed":"Ortygospiza atricoll mueller","simple":"Ortygospiza atricollis muelleri","full":"Ortygospiza atricollis muelleri"},"cardinality":3,"details":{"infraSpecies":{"genus":"Ortygospiza","species":"atricollis","infraSpecies":[{"value":"muelleri"}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":30}],"id":"1ee6bf1d-90d8-5c4b-98c1-2646c301d07c","parserVersion":"test_version"}
```

Name: Cortinarius angulatus B gracilescens Fr. 1838

Canonical: Cortinarius angulatus gracilescens

Authorship: Fr. 1838

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Author is too short"}],"verbatim":"Cortinarius angulatus B gracilescens Fr. 1838","normalized":"Cortinarius angulatus B gracilescens Fr. 1838","canonical":{"stemmed":"Cortinarius angulat gracilescens","simple":"Cortinarius angulatus gracilescens","full":"Cortinarius angulatus gracilescens"},"cardinality":3,"authorship":{"verbatim":"Fr. 1838","normalized":"Fr. 1838","year":"1838","authors":["Fr."],"originalAuth":{"authors":["Fr."],"year":{"year":"1838"}}},"details":{"infraSpecies":{"genus":"Cortinarius","species":"angulatus","authorship":{"verbatim":"B","normalized":"B","authors":["B"],"originalAuth":{"authors":["B"]}},"infraSpecies":[{"value":"gracilescens","authorship":{"verbatim":"Fr. 1838","normalized":"Fr. 1838","year":"1838","authors":["Fr."],"originalAuth":{"authors":["Fr."],"year":{"year":"1838"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":22,"end":23},{"wordType":"infraspecificEpithet","start":24,"end":36},{"wordType":"authorWord","start":37,"end":40},{"wordType":"year","start":41,"end":45}],"id":"3fb101ad-d05e-5648-993b-bfbb8c76166e","parserVersion":"test_version"}
```

Name: Caulerpa fastigiata confervoides P. L. Crouan & H. M. Crouan ex Weber-van Bosse

Canonical: Caulerpa fastigiata confervoides

Authorship: P. L. Crouan & H. M. Crouan ex Weber-van Bosse

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Caulerpa fastigiata confervoides P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"Caulerpa fastigiata confervoides P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","canonical":{"stemmed":"Caulerpa fastigiat conferuoid","simple":"Caulerpa fastigiata confervoides","full":"Caulerpa fastigiata confervoides"},"cardinality":3,"authorship":{"verbatim":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","authors":["P. L. Crouan","H. M. Crouan"],"originalAuth":{"authors":["P. L. Crouan","H. M. Crouan"],"exAuthors":{"authors":["Weber-van Bosse"]}}},"details":{"infraSpecies":{"genus":"Caulerpa","species":"fastigiata","infraSpecies":[{"value":"confervoides","authorship":{"verbatim":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","normalized":"P. L. Crouan \u0026 H. M. Crouan ex Weber-van Bosse","authors":["P. L. Crouan","H. M. Crouan"],"originalAuth":{"authors":["P. L. Crouan","H. M. Crouan"],"exAuthors":{"authors":["Weber-van Bosse"]}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"infraspecificEpithet","start":20,"end":32},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":39,"end":45},{"wordType":"authorWord","start":48,"end":50},{"wordType":"authorWord","start":51,"end":53},{"wordType":"authorWord","start":54,"end":60},{"wordType":"authorWord","start":64,"end":73},{"wordType":"authorWord","start":74,"end":79}],"id":"8934dbda-1fd2-52c4-af76-8f80e5f02791","parserVersion":"test_version"}
```

### Legacy ICZN names with rank

Name: Acipenser gueldenstaedti colchicus natio danubicus Movchan, 1967

Canonical: Acipenser gueldenstaedti colchicus natio danubicus

Authorship: Movchan 1967

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Acipenser gueldenstaedti colchicus natio danubicus Movchan, 1967","normalized":"Acipenser gueldenstaedti colchicus natio danubicus Movchan 1967","canonical":{"stemmed":"Acipenser gueldenstaedt colchic danubic","simple":"Acipenser gueldenstaedti colchicus danubicus","full":"Acipenser gueldenstaedti colchicus natio danubicus"},"cardinality":4,"authorship":{"verbatim":"Movchan, 1967","normalized":"Movchan 1967","year":"1967","authors":["Movchan"],"originalAuth":{"authors":["Movchan"],"year":{"year":"1967"}}},"details":{"infraSpecies":{"genus":"Acipenser","species":"gueldenstaedti","infraSpecies":[{"value":"colchicus"},{"value":"danubicus","rank":"natio","authorship":{"verbatim":"Movchan, 1967","normalized":"Movchan 1967","year":"1967","authors":["Movchan"],"originalAuth":{"authors":["Movchan"],"year":{"year":"1967"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":34},{"wordType":"rank","start":35,"end":40},{"wordType":"infraspecificEpithet","start":41,"end":50},{"wordType":"authorWord","start":51,"end":58},{"wordType":"year","start":60,"end":64}],"id":"d572e7a6-bcbd-59ef-bc60-1e5d659fd51c","parserVersion":"test_version"}
```

### Infraspecies with rank (ICN)

Name: Crematogaster impressa st. brazzai Santschi 1937

Canonical: Crematogaster impressa st. brazzai

Authorship: Santschi 1937

```json
{"parsed":true,"parseQuality":1,"verbatim":"Crematogaster impressa st. brazzai Santschi 1937","normalized":"Crematogaster impressa st. brazzai Santschi 1937","canonical":{"stemmed":"Crematogaster impress brazza","simple":"Crematogaster impressa brazzai","full":"Crematogaster impressa st. brazzai"},"cardinality":3,"authorship":{"verbatim":"Santschi 1937","normalized":"Santschi 1937","year":"1937","authors":["Santschi"],"originalAuth":{"authors":["Santschi"],"year":{"year":"1937"}}},"details":{"infraSpecies":{"genus":"Crematogaster","species":"impressa","infraSpecies":[{"value":"brazzai","rank":"st.","authorship":{"verbatim":"Santschi 1937","normalized":"Santschi 1937","year":"1937","authors":["Santschi"],"originalAuth":{"authors":["Santschi"],"year":{"year":"1937"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":22},{"wordType":"rank","start":23,"end":26},{"wordType":"infraspecificEpithet","start":27,"end":34},{"wordType":"authorWord","start":35,"end":43},{"wordType":"year","start":44,"end":48}],"id":"853d0cff-b499-5d38-ae49-75b558f9ddf0","parserVersion":"test_version"}
```

<!-- badly formed name, we do not deal with it for now -->
Name: Cibotium st.-johnii Krajina

Canonical: Cibotium

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Cibotium st.-johnii Krajina","normalized":"Cibotium","canonical":{"stemmed":"Cibotium","simple":"Cibotium","full":"Cibotium"},"cardinality":1,"tail":" st.-johnii Krajina","details":{"uninomial":{"uninomial":"Cibotium"}},"pos":[{"wordType":"uninomial","start":0,"end":8}],"id":"6b34256d-6c3b-5870-a781-77eeac49b6c4","parserVersion":"test_version"}
```

Name: Camponotus conspicuus st. zonatus

Canonical: Camponotus conspicuus st. zonatus

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Camponotus conspicuus st. zonatus","normalized":"Camponotus conspicuus st. zonatus","canonical":{"stemmed":"Camponotus conspicu zonat","simple":"Camponotus conspicuus zonatus","full":"Camponotus conspicuus st. zonatus"},"cardinality":3,"details":{"infraSpecies":{"genus":"Camponotus","species":"conspicuus","infraSpecies":[{"value":"zonatus","rank":"st."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"rank","start":22,"end":25},{"wordType":"infraspecificEpithet","start":26,"end":33}],"id":"67364c72-53e0-54d3-9795-f04fd1938d75","parserVersion":"test_version"}
```

Name: Fagus sylvatica subsp. orientalis (Lipsky) Greuter & Burdet

Canonical: Fagus sylvatica subsp. orientalis

Authorship: (Lipsky) Greuter & Burdet

```json
{"parsed":true,"parseQuality":1,"verbatim":"Fagus sylvatica subsp. orientalis (Lipsky) Greuter \u0026 Burdet","normalized":"Fagus sylvatica subsp. orientalis (Lipsky) Greuter \u0026 Burdet","canonical":{"stemmed":"Fagus syluatic oriental","simple":"Fagus sylvatica orientalis","full":"Fagus sylvatica subsp. orientalis"},"cardinality":3,"authorship":{"verbatim":"(Lipsky) Greuter \u0026 Burdet","normalized":"(Lipsky) Greuter \u0026 Burdet","authors":["Lipsky","Greuter","Burdet"],"originalAuth":{"authors":["Lipsky"]},"combinationAuth":{"authors":["Greuter","Burdet"]}},"details":{"infraSpecies":{"genus":"Fagus","species":"sylvatica","infraSpecies":[{"value":"orientalis","rank":"subsp.","authorship":{"verbatim":"(Lipsky) Greuter \u0026 Burdet","normalized":"(Lipsky) Greuter \u0026 Burdet","authors":["Lipsky","Greuter","Burdet"],"originalAuth":{"authors":["Lipsky"]},"combinationAuth":{"authors":["Greuter","Burdet"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":15},{"wordType":"rank","start":16,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":33},{"wordType":"authorWord","start":35,"end":41},{"wordType":"authorWord","start":43,"end":50},{"wordType":"authorWord","start":53,"end":59}],"id":"f0bff1a3-0923-58d1-807f-c5da5b85531e","parserVersion":"test_version"}
```

Name: Tillandsia utriculata subspec. utriculata

Canonical: Tillandsia utriculata subsp. utriculata

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Tillandsia utriculata subspec. utriculata","normalized":"Tillandsia utriculata subsp. utriculata","canonical":{"stemmed":"Tillandsia utriculat utriculat","simple":"Tillandsia utriculata utriculata","full":"Tillandsia utriculata subsp. utriculata"},"cardinality":3,"details":{"infraSpecies":{"genus":"Tillandsia","species":"utriculata","infraSpecies":[{"value":"utriculata","rank":"subsp."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"rank","start":22,"end":30},{"wordType":"infraspecificEpithet","start":31,"end":41}],"id":"fa612e5d-f697-5227-a5a0-fdb4a1aafe7a","parserVersion":"test_version"}
```

Name: Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.

Canonical: Prunus mexicana var. reticulata

Authorship: (Sarg.) Sarg.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.","normalized":"Prunus mexicana S. Watson var. reticulata (Sarg.) Sarg.","canonical":{"stemmed":"Prunus mexican reticulat","simple":"Prunus mexicana reticulata","full":"Prunus mexicana var. reticulata"},"cardinality":3,"authorship":{"verbatim":"(Sarg.) Sarg.","normalized":"(Sarg.) Sarg.","authors":["Sarg.","Sarg."],"originalAuth":{"authors":["Sarg."]},"combinationAuth":{"authors":["Sarg."]}},"details":{"infraSpecies":{"genus":"Prunus","species":"mexicana","authorship":{"verbatim":"S. Watson","normalized":"S. Watson","authors":["S. Watson"],"originalAuth":{"authors":["S. Watson"]}},"infraSpecies":[{"value":"reticulata","rank":"var.","authorship":{"verbatim":"(Sarg.) Sarg.","normalized":"(Sarg.) Sarg.","authors":["Sarg.","Sarg."],"originalAuth":{"authors":["Sarg."]},"combinationAuth":{"authors":["Sarg."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":19,"end":25},{"wordType":"rank","start":26,"end":30},{"wordType":"infraspecificEpithet","start":31,"end":41},{"wordType":"authorWord","start":43,"end":48},{"wordType":"authorWord","start":50,"end":55}],"id":"5ba1cc96-ab40-51b3-951d-f91b5bff1da8","parserVersion":"test_version"}
```

Name: Potamogeton iilinoensis var. ventanicola

Canonical: Potamogeton iilinoensis var. ventanicola

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Potamogeton iilinoensis var. ventanicola","normalized":"Potamogeton iilinoensis var. ventanicola","canonical":{"stemmed":"Potamogeton iilinoens uentanicol","simple":"Potamogeton iilinoensis ventanicola","full":"Potamogeton iilinoensis var. ventanicola"},"cardinality":3,"details":{"infraSpecies":{"genus":"Potamogeton","species":"iilinoensis","infraSpecies":[{"value":"ventanicola","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":23},{"wordType":"rank","start":24,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":40}],"id":"edf418ec-98b3-52fb-a8de-26808b61c50f","parserVersion":"test_version"}
```

Name: Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien

Canonical: Potamogeton iilinoensis var. ventanicola

Authorship: (Hicken) Horn af Rantzien

```json
{"parsed":true,"parseQuality":1,"verbatim":"Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien","normalized":"Potamogeton iilinoensis var. ventanicola (Hicken) Horn af Rantzien","canonical":{"stemmed":"Potamogeton iilinoens uentanicol","simple":"Potamogeton iilinoensis ventanicola","full":"Potamogeton iilinoensis var. ventanicola"},"cardinality":3,"authorship":{"verbatim":"(Hicken) Horn af Rantzien","normalized":"(Hicken) Horn af Rantzien","authors":["Hicken","Horn af Rantzien"],"originalAuth":{"authors":["Hicken"]},"combinationAuth":{"authors":["Horn af Rantzien"]}},"details":{"infraSpecies":{"genus":"Potamogeton","species":"iilinoensis","infraSpecies":[{"value":"ventanicola","rank":"var.","authorship":{"verbatim":"(Hicken) Horn af Rantzien","normalized":"(Hicken) Horn af Rantzien","authors":["Hicken","Horn af Rantzien"],"originalAuth":{"authors":["Hicken"]},"combinationAuth":{"authors":["Horn af Rantzien"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":23},{"wordType":"rank","start":24,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":40},{"wordType":"authorWord","start":42,"end":48},{"wordType":"authorWord","start":50,"end":54},{"wordType":"authorWord","start":55,"end":57},{"wordType":"authorWord","start":58,"end":66}],"id":"e7888abd-4365-5d74-8d5f-a69c8196328e","parserVersion":"test_version"}
```

Name: Triticum repens var. vulgäre

Canonical: Triticum repens var. vulgaere

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Triticum repens var. vulgäre","normalized":"Triticum repens var. vulgaere","canonical":{"stemmed":"Triticum repens uulgaer","simple":"Triticum repens vulgaere","full":"Triticum repens var. vulgaere"},"cardinality":3,"details":{"infraSpecies":{"genus":"Triticum","species":"repens","infraSpecies":[{"value":"vulgaere","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"rank","start":16,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":28}],"id":"3421b13b-aaa9-5234-bc1d-9d3fe7a6b19e","parserVersion":"test_version"}
```

Name: Aus bus Linn. var. bus

Canonical: Aus bus var. bus

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aus bus Linn. var. bus","normalized":"Aus bus Linn. var. bus","canonical":{"stemmed":"Aus bus bus","simple":"Aus bus bus","full":"Aus bus var. bus"},"cardinality":3,"details":{"infraSpecies":{"genus":"Aus","species":"bus","authorship":{"verbatim":"Linn.","normalized":"Linn.","authors":["Linn."],"originalAuth":{"authors":["Linn."]}},"infraSpecies":[{"value":"bus","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":3},{"wordType":"specificEpithet","start":4,"end":7},{"wordType":"authorWord","start":8,"end":13},{"wordType":"rank","start":14,"end":18},{"wordType":"infraspecificEpithet","start":19,"end":22}],"id":"2a6e45e2-5737-514b-8055-06f8a878dd36","parserVersion":"test_version"}
```

Name: Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987

Canonical: Agalinis purpurea var. borealis

Authorship: (Berg.) Peterson 1987

```json
{"parsed":true,"parseQuality":1,"verbatim":"Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987","normalized":"Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987","canonical":{"stemmed":"Agalinis purpure boreal","simple":"Agalinis purpurea borealis","full":"Agalinis purpurea var. borealis"},"cardinality":3,"authorship":{"verbatim":"(Berg.) Peterson 1987","normalized":"(Berg.) Peterson 1987","authors":["Berg.","Peterson"],"originalAuth":{"authors":["Berg."]},"combinationAuth":{"authors":["Peterson"],"year":{"year":"1987"}}},"details":{"infraSpecies":{"genus":"Agalinis","species":"purpurea","authorship":{"verbatim":"(L.) Briton","normalized":"(L.) Briton","authors":["L.","Briton"],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Briton"]}},"infraSpecies":[{"value":"borealis","rank":"var.","authorship":{"verbatim":"(Berg.) Peterson 1987","normalized":"(Berg.) Peterson 1987","authors":["Berg.","Peterson"],"originalAuth":{"authors":["Berg."]},"combinationAuth":{"authors":["Peterson"],"year":{"year":"1987"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":19,"end":21},{"wordType":"authorWord","start":23,"end":29},{"wordType":"rank","start":30,"end":34},{"wordType":"infraspecificEpithet","start":35,"end":43},{"wordType":"authorWord","start":45,"end":50},{"wordType":"authorWord","start":52,"end":60},{"wordType":"year","start":61,"end":65}],"id":"769863cd-7c9d-5d4a-bf5c-fb6903a96431","parserVersion":"test_version"}
```

Name: Callideriphus flavicollis morph. reductus Fuchs 1961

Canonical: Callideriphus flavicollis morph. reductus

Authorship: Fuchs 1961

```json
{"parsed":true,"parseQuality":1,"verbatim":"Callideriphus flavicollis morph. reductus Fuchs 1961","normalized":"Callideriphus flavicollis morph. reductus Fuchs 1961","canonical":{"stemmed":"Callideriphus flauicoll reduct","simple":"Callideriphus flavicollis reductus","full":"Callideriphus flavicollis morph. reductus"},"cardinality":3,"authorship":{"verbatim":"Fuchs 1961","normalized":"Fuchs 1961","year":"1961","authors":["Fuchs"],"originalAuth":{"authors":["Fuchs"],"year":{"year":"1961"}}},"details":{"infraSpecies":{"genus":"Callideriphus","species":"flavicollis","infraSpecies":[{"value":"reductus","rank":"morph.","authorship":{"verbatim":"Fuchs 1961","normalized":"Fuchs 1961","year":"1961","authors":["Fuchs"],"originalAuth":{"authors":["Fuchs"],"year":{"year":"1961"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":25},{"wordType":"rank","start":26,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":41},{"wordType":"authorWord","start":42,"end":47},{"wordType":"year","start":48,"end":52}],"id":"2b01f892-dbb3-5776-870a-c6cb8f09f2bc","parserVersion":"test_version"}
```

Name: Caulerpa cupressoides forma nuda

Canonical: Caulerpa cupressoides f. nuda

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Caulerpa cupressoides forma nuda","normalized":"Caulerpa cupressoides f. nuda","canonical":{"stemmed":"Caulerpa cupressoid nud","simple":"Caulerpa cupressoides nuda","full":"Caulerpa cupressoides f. nuda"},"cardinality":3,"details":{"infraSpecies":{"genus":"Caulerpa","species":"cupressoides","infraSpecies":[{"value":"nuda","rank":"f."}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"rank","start":22,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":32}],"id":"805ee92d-001e-5f05-abad-446f683860cb","parserVersion":"test_version"}
```

Name: Chlorocyperus glaber form. fasciculariforme (Lojac.) Soó

Canonical: Chlorocyperus glaber f. fasciculariforme

Authorship: (Lojac.) Soó

```json
{"parsed":true,"parseQuality":1,"verbatim":"Chlorocyperus glaber form. fasciculariforme (Lojac.) Soó","normalized":"Chlorocyperus glaber f. fasciculariforme (Lojac.) Soó","canonical":{"stemmed":"Chlorocyperus glaber fasciculariform","simple":"Chlorocyperus glaber fasciculariforme","full":"Chlorocyperus glaber f. fasciculariforme"},"cardinality":3,"authorship":{"verbatim":"(Lojac.) Soó","normalized":"(Lojac.) Soó","authors":["Lojac.","Soó"],"originalAuth":{"authors":["Lojac."]},"combinationAuth":{"authors":["Soó"]}},"details":{"infraSpecies":{"genus":"Chlorocyperus","species":"glaber","infraSpecies":[{"value":"fasciculariforme","rank":"f.","authorship":{"verbatim":"(Lojac.) Soó","normalized":"(Lojac.) Soó","authors":["Lojac.","Soó"],"originalAuth":{"authors":["Lojac."]},"combinationAuth":{"authors":["Soó"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":20},{"wordType":"rank","start":21,"end":26},{"wordType":"infraspecificEpithet","start":27,"end":43},{"wordType":"authorWord","start":45,"end":51},{"wordType":"authorWord","start":53,"end":56}],"id":"beee0dba-bef6-5550-954f-c978af09310a","parserVersion":"test_version"}
```

Name: Sphaerotheca    fuliginea    f.     dahliae    Movss.     1967

Canonical: Sphaerotheca fuliginea f. dahliae

Authorship: Movss. 1967

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Sphaerotheca    fuliginea    f.     dahliae    Movss.     1967","normalized":"Sphaerotheca fuliginea f. dahliae Movss. 1967","canonical":{"stemmed":"Sphaerotheca fuligine dahli","simple":"Sphaerotheca fuliginea dahliae","full":"Sphaerotheca fuliginea f. dahliae"},"cardinality":3,"authorship":{"verbatim":"Movss.     1967","normalized":"Movss. 1967","year":"1967","authors":["Movss."],"originalAuth":{"authors":["Movss."],"year":{"year":"1967"}}},"details":{"infraSpecies":{"genus":"Sphaerotheca","species":"fuliginea","infraSpecies":[{"value":"dahliae","rank":"f.","authorship":{"verbatim":"Movss.     1967","normalized":"Movss. 1967","year":"1967","authors":["Movss."],"originalAuth":{"authors":["Movss."],"year":{"year":"1967"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":16,"end":25},{"wordType":"rank","start":29,"end":31},{"wordType":"infraspecificEpithet","start":36,"end":43},{"wordType":"authorWord","start":47,"end":53},{"wordType":"year","start":58,"end":62}],"id":"bbd48fd4-ceee-5c66-ae42-f7fa43a8ea97","parserVersion":"test_version"}
```

Name: Allophylus amazonicus var amazonicus

Canonical: Allophylus amazonicus var. amazonicus

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Allophylus amazonicus var amazonicus","normalized":"Allophylus amazonicus var. amazonicus","canonical":{"stemmed":"Allophylus amazonic amazonic","simple":"Allophylus amazonicus amazonicus","full":"Allophylus amazonicus var. amazonicus"},"cardinality":3,"details":{"infraSpecies":{"genus":"Allophylus","species":"amazonicus","infraSpecies":[{"value":"amazonicus","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"rank","start":22,"end":25},{"wordType":"infraspecificEpithet","start":26,"end":36}],"id":"4e5c108c-b089-5198-9088-dd58d74d951f","parserVersion":"test_version"}
```

Name: Yarrowia lipolytica variety lipolytic

Canonical: Yarrowia lipolytica var. lipolytic

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Yarrowia lipolytica variety lipolytic","normalized":"Yarrowia lipolytica var. lipolytic","canonical":{"stemmed":"Yarrowia lipolytic lipolytic","simple":"Yarrowia lipolytica lipolytic","full":"Yarrowia lipolytica var. lipolytic"},"cardinality":3,"details":{"infraSpecies":{"genus":"Yarrowia","species":"lipolytica","infraSpecies":[{"value":"lipolytic","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"rank","start":20,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":37}],"id":"5ecc8759-e1c3-5632-a863-7664625fc58d","parserVersion":"test_version"}
```

Name: Prunus armeniaca convar. budae (Pénzes) Soó

Canonical: Prunus armeniaca convar. budae

Authorship: (Pénzes) Soó

```json
{"parsed":true,"parseQuality":1,"verbatim":"Prunus armeniaca convar. budae (Pénzes) Soó","normalized":"Prunus armeniaca convar. budae (Pénzes) Soó","canonical":{"stemmed":"Prunus armeniac bud","simple":"Prunus armeniaca budae","full":"Prunus armeniaca convar. budae"},"cardinality":3,"authorship":{"verbatim":"(Pénzes) Soó","normalized":"(Pénzes) Soó","authors":["Pénzes","Soó"],"originalAuth":{"authors":["Pénzes"]},"combinationAuth":{"authors":["Soó"]}},"details":{"infraSpecies":{"genus":"Prunus","species":"armeniaca","infraSpecies":[{"value":"budae","rank":"convar.","authorship":{"verbatim":"(Pénzes) Soó","normalized":"(Pénzes) Soó","authors":["Pénzes","Soó"],"originalAuth":{"authors":["Pénzes"]},"combinationAuth":{"authors":["Soó"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":16},{"wordType":"rank","start":17,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":30},{"wordType":"authorWord","start":32,"end":38},{"wordType":"authorWord","start":40,"end":43}],"id":"c2133c2d-0486-54cb-a8cb-d355d458e19f","parserVersion":"test_version"}
```

Name: Polypodium pectinatum (L.) f. typica Rosenst.

Canonical: Polypodium pectinatum f. typica

Authorship: Rosenst.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Polypodium pectinatum (L.) f. typica Rosenst.","normalized":"Polypodium pectinatum (L.) f. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum f. typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraSpecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"(L.)","normalized":"(L.)","authors":["L."],"originalAuth":{"authors":["L."]}},"infraSpecies":[{"value":"typica","rank":"f.","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":23,"end":25},{"wordType":"rank","start":27,"end":29},{"wordType":"infraspecificEpithet","start":30,"end":36},{"wordType":"authorWord","start":37,"end":45}],"id":"b74dfd6b-c2d5-5e21-a807-f138667f0370","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f. typica Rosenst.

Canonical: Polypodium pectinatum typica

Authorship: Rosenst.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium pectinatum L. f. typica Rosenst.","normalized":"Polypodium pectinatum L. fil. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraSpecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L. f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"infraSpecies":[{"value":"typica","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWordFilius","start":25,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":34},{"wordType":"authorWord","start":35,"end":43}],"id":"68a2dccb-8b41-5a4f-92aa-06ae377b1503","parserVersion":"test_version"}
```

Name: Rubus fruticosus agamosp. chloocladus (W.C.R. Watson) A. & D. Löve

Canonical: Rubus fruticosus agamosp. chloocladus

Authorship: (W. C. R. Watson) A. & D. Löve

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rubus fruticosus agamosp. chloocladus (W.C.R. Watson) A. \u0026 D. Löve","normalized":"Rubus fruticosus agamosp. chloocladus (W. C. R. Watson) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos chlooclad","simple":"Rubus fruticosus chloocladus","full":"Rubus fruticosus agamosp. chloocladus"},"cardinality":3,"authorship":{"verbatim":"(W.C.R. Watson) A. \u0026 D. Löve","normalized":"(W. C. R. Watson) A. \u0026 D. Löve","authors":["W. C. R. Watson","A.","D. Löve"],"originalAuth":{"authors":["W. C. R. Watson"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraSpecies":{"genus":"Rubus","species":"fruticosus","infraSpecies":[{"value":"chloocladus","rank":"agamosp.","authorship":{"verbatim":"(W.C.R. Watson) A. \u0026 D. Löve","normalized":"(W. C. R. Watson) A. \u0026 D. Löve","authors":["W. C. R. Watson","A.","D. Löve"],"originalAuth":{"authors":["W. C. R. Watson"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":16},{"wordType":"rank","start":17,"end":25},{"wordType":"infraspecificEpithet","start":26,"end":37},{"wordType":"authorWord","start":39,"end":41},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":52},{"wordType":"authorWord","start":54,"end":56},{"wordType":"authorWord","start":59,"end":61},{"wordType":"authorWord","start":62,"end":66}],"id":"c6a80c28-12ab-550e-8255-3b96032ef98c","parserVersion":"test_version"}
```

Name: Rubus fruticosus L. agamossp. discolor (Weihe & Nees) A. & D. Löve

Canonical: Rubus fruticosus agamossp. discolor

Authorship: (Weihe & Nees) A. & D. Löve

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rubus fruticosus L. agamossp. discolor (Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"Rubus fruticosus L. agamossp. discolor (Weihe \u0026 Nees) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos discolor","simple":"Rubus fruticosus discolor","full":"Rubus fruticosus agamossp. discolor"},"cardinality":3,"authorship":{"verbatim":"(Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"(Weihe \u0026 Nees) A. \u0026 D. Löve","authors":["Weihe","Nees","A.","D. Löve"],"originalAuth":{"authors":["Weihe","Nees"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraSpecies":{"genus":"Rubus","species":"fruticosus","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraSpecies":[{"value":"discolor","rank":"agamossp.","authorship":{"verbatim":"(Weihe \u0026 Nees) A. \u0026 D. Löve","normalized":"(Weihe \u0026 Nees) A. \u0026 D. Löve","authors":["Weihe","Nees","A.","D. Löve"],"originalAuth":{"authors":["Weihe","Nees"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":16},{"wordType":"authorWord","start":17,"end":19},{"wordType":"rank","start":20,"end":29},{"wordType":"infraspecificEpithet","start":30,"end":38},{"wordType":"authorWord","start":40,"end":45},{"wordType":"authorWord","start":48,"end":52},{"wordType":"authorWord","start":54,"end":56},{"wordType":"authorWord","start":59,"end":61},{"wordType":"authorWord","start":62,"end":66}],"id":"a4265faa-5096-575b-914c-cd9cea4bbb7d","parserVersion":"test_version"}
```

Name: Rubus fruticosus agamovar. graecensis (W.Maurer) A. & D. Löve

Canonical: Rubus fruticosus agamovar. graecensis

Authorship: (W. Maurer) A. & D. Löve

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rubus fruticosus agamovar. graecensis (W.Maurer) A. \u0026 D. Löve","normalized":"Rubus fruticosus agamovar. graecensis (W. Maurer) A. \u0026 D. Löve","canonical":{"stemmed":"Rubus fruticos graecens","simple":"Rubus fruticosus graecensis","full":"Rubus fruticosus agamovar. graecensis"},"cardinality":3,"authorship":{"verbatim":"(W.Maurer) A. \u0026 D. Löve","normalized":"(W. Maurer) A. \u0026 D. Löve","authors":["W. Maurer","A.","D. Löve"],"originalAuth":{"authors":["W. Maurer"]},"combinationAuth":{"authors":["A.","D. Löve"]}},"details":{"infraSpecies":{"genus":"Rubus","species":"fruticosus","infraSpecies":[{"value":"graecensis","rank":"agamovar.","authorship":{"verbatim":"(W.Maurer) A. \u0026 D. Löve","normalized":"(W. Maurer) A. \u0026 D. Löve","authors":["W. Maurer","A.","D. Löve"],"originalAuth":{"authors":["W. Maurer"]},"combinationAuth":{"authors":["A.","D. Löve"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":16},{"wordType":"rank","start":17,"end":26},{"wordType":"infraspecificEpithet","start":27,"end":37},{"wordType":"authorWord","start":39,"end":41},{"wordType":"authorWord","start":41,"end":47},{"wordType":"authorWord","start":49,"end":51},{"wordType":"authorWord","start":54,"end":56},{"wordType":"authorWord","start":57,"end":61}],"id":"9e3158af-63bd-5c94-91d1-f795342709d6","parserVersion":"test_version"}
```

<!-- TODO: the following phrasing can be ambiguous.
Does f mean forma or filius? Currently capturing it as filius -->
Name: Polypodium pectinatum L.f. typica Rosenst.

Canonical: Polypodium pectinatum typica

Authorship: Rosenst.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium pectinatum L.f. typica Rosenst.","normalized":"Polypodium pectinatum L. fil. typica Rosenst.","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum typica"},"cardinality":3,"authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}},"details":{"infraSpecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L.f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"infraSpecies":[{"value":"typica","authorship":{"verbatim":"Rosenst.","normalized":"Rosenst.","authors":["Rosenst."],"originalAuth":{"authors":["Rosenst."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWordFilius","start":24,"end":26},{"wordType":"infraspecificEpithet","start":27,"end":33},{"wordType":"authorWord","start":34,"end":42}],"id":"ea87b733-cae3-5a0f-a74d-3d921dcdbeb6","parserVersion":"test_version"}
```

Name: Polypodium lineare C.Chr. f. caudatoattenuatum Takeda

Canonical: Polypodium lineare caudatoattenuatum

Authorship: Takeda

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Polypodium lineare C.Chr. f. caudatoattenuatum Takeda","normalized":"Polypodium lineare C. Chr. fil. caudatoattenuatum Takeda","canonical":{"stemmed":"Polypodium linear caudatoattenuat","simple":"Polypodium lineare caudatoattenuatum","full":"Polypodium lineare caudatoattenuatum"},"cardinality":3,"authorship":{"verbatim":"Takeda","normalized":"Takeda","authors":["Takeda"],"originalAuth":{"authors":["Takeda"]}},"details":{"infraSpecies":{"genus":"Polypodium","species":"lineare","authorship":{"verbatim":"C.Chr. f.","normalized":"C. Chr. fil.","authors":["C. Chr. fil."],"originalAuth":{"authors":["C. Chr. fil."]}},"infraSpecies":[{"value":"caudatoattenuatum","authorship":{"verbatim":"Takeda","normalized":"Takeda","authors":["Takeda"],"originalAuth":{"authors":["Takeda"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":18},{"wordType":"authorWord","start":19,"end":21},{"wordType":"authorWord","start":21,"end":25},{"wordType":"authorWordFilius","start":26,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":46},{"wordType":"authorWord","start":47,"end":53}],"id":"18cfd931-1ccd-5ea2-823a-71ba9604c783","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. f. albiflorum T.Yamaz.

Canonical: Rhododendron weyrichii albiflorum

Authorship: T. Yamaz.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Rhododendron weyrichii Maxim. f. albiflorum T.Yamaz.","normalized":"Rhododendron weyrichii Maxim. fil. albiflorum T. Yamaz.","canonical":{"stemmed":"Rhododendron weyrichi albiflor","simple":"Rhododendron weyrichii albiflorum","full":"Rhododendron weyrichii albiflorum"},"cardinality":3,"authorship":{"verbatim":"T.Yamaz.","normalized":"T. Yamaz.","authors":["T. Yamaz."],"originalAuth":{"authors":["T. Yamaz."]}},"details":{"infraSpecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim. f.","normalized":"Maxim. fil.","authors":["Maxim. fil."],"originalAuth":{"authors":["Maxim. fil."]}},"infraSpecies":[{"value":"albiflorum","authorship":{"verbatim":"T.Yamaz.","normalized":"T. Yamaz.","authors":["T. Yamaz."],"originalAuth":{"authors":["T. Yamaz."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22},{"wordType":"authorWord","start":23,"end":29},{"wordType":"authorWordFilius","start":30,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":43},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":46,"end":52}],"id":"e515f1c8-3b95-5930-bcd1-09176727f0b7","parserVersion":"test_version"}
```

Name: Armeria maaritima (Mill.) Willd. fma. originaria Bern.

Canonical: Armeria maaritima f. originaria

Authorship: Bern.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Armeria maaritima (Mill.) Willd. fma. originaria Bern.","normalized":"Armeria maaritima (Mill.) Willd. f. originaria Bern.","canonical":{"stemmed":"Armeria maaritim originar","simple":"Armeria maaritima originaria","full":"Armeria maaritima f. originaria"},"cardinality":3,"authorship":{"verbatim":"Bern.","normalized":"Bern.","authors":["Bern."],"originalAuth":{"authors":["Bern."]}},"details":{"infraSpecies":{"genus":"Armeria","species":"maaritima","authorship":{"verbatim":"(Mill.) Willd.","normalized":"(Mill.) Willd.","authors":["Mill.","Willd."],"originalAuth":{"authors":["Mill."]},"combinationAuth":{"authors":["Willd."]}},"infraSpecies":[{"value":"originaria","rank":"f.","authorship":{"verbatim":"Bern.","normalized":"Bern.","authors":["Bern."],"originalAuth":{"authors":["Bern."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":19,"end":24},{"wordType":"authorWord","start":26,"end":32},{"wordType":"rank","start":33,"end":37},{"wordType":"infraspecificEpithet","start":38,"end":48},{"wordType":"authorWord","start":49,"end":54}],"id":"00d88bea-f076-5911-a450-fcfac1fe98bc","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. albiflorum T.Yamaz. f. fakeepithet

Canonical: Rhododendron weyrichii albiflorum fakeepithet

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguous f. (filius or forma)"}],"verbatim":"Rhododendron weyrichii Maxim. albiflorum T.Yamaz. f. fakeepithet","normalized":"Rhododendron weyrichii Maxim. albiflorum T. Yamaz. fil. fakeepithet","canonical":{"stemmed":"Rhododendron weyrichi albiflor fakeepithet","simple":"Rhododendron weyrichii albiflorum fakeepithet","full":"Rhododendron weyrichii albiflorum fakeepithet"},"cardinality":4,"details":{"infraSpecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim.","normalized":"Maxim.","authors":["Maxim."],"originalAuth":{"authors":["Maxim."]}},"infraSpecies":[{"value":"albiflorum","authorship":{"verbatim":"T.Yamaz. f.","normalized":"T. Yamaz. fil.","authors":["T. Yamaz. fil."],"originalAuth":{"authors":["T. Yamaz. fil."]}}},{"value":"fakeepithet"}]}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22},{"wordType":"authorWord","start":23,"end":29},{"wordType":"infraspecificEpithet","start":30,"end":40},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":43,"end":49},{"wordType":"authorWordFilius","start":50,"end":52},{"wordType":"infraspecificEpithet","start":53,"end":64}],"id":"ad0e299f-cd2c-52f3-9cab-49c70c5814f8","parserVersion":"test_version"}
```

Name: Rhododendron weyrichii Maxim. albiflorum (T.Yamaz. f.) fakeepithet

Canonical: Rhododendron weyrichii albiflorum fakeepithet

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rhododendron weyrichii Maxim. albiflorum (T.Yamaz. f.) fakeepithet","normalized":"Rhododendron weyrichii Maxim. albiflorum (T. Yamaz. fil.) fakeepithet","canonical":{"stemmed":"Rhododendron weyrichi albiflor fakeepithet","simple":"Rhododendron weyrichii albiflorum fakeepithet","full":"Rhododendron weyrichii albiflorum fakeepithet"},"cardinality":4,"details":{"infraSpecies":{"genus":"Rhododendron","species":"weyrichii","authorship":{"verbatim":"Maxim.","normalized":"Maxim.","authors":["Maxim."],"originalAuth":{"authors":["Maxim."]}},"infraSpecies":[{"value":"albiflorum","authorship":{"verbatim":"(T.Yamaz. f.)","normalized":"(T. Yamaz. fil.)","authors":["T. Yamaz. fil."],"originalAuth":{"authors":["T. Yamaz. fil."]}}},{"value":"fakeepithet"}]}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22},{"wordType":"authorWord","start":23,"end":29},{"wordType":"infraspecificEpithet","start":30,"end":40},{"wordType":"authorWord","start":42,"end":44},{"wordType":"authorWord","start":44,"end":50},{"wordType":"authorWordFilius","start":51,"end":53},{"wordType":"infraspecificEpithet","start":55,"end":66}],"id":"2a7d1bab-b208-5654-9406-f7afc696b00b","parserVersion":"test_version"}
```

Name: Cotoneaster (Pyracantha) rogersiana var.aurantiaca

Canonical: Cotoneaster rogersiana var. aurantiaca

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cotoneaster (Pyracantha) rogersiana var.aurantiaca","normalized":"Cotoneaster (Pyracantha) rogersiana var. aurantiaca","canonical":{"stemmed":"Cotoneaster rogersian aurantiac","simple":"Cotoneaster rogersiana aurantiaca","full":"Cotoneaster rogersiana var. aurantiaca"},"cardinality":3,"details":{"infraSpecies":{"genus":"Cotoneaster","subGenus":"Pyracantha","species":"rogersiana","infraSpecies":[{"value":"aurantiaca","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"infragenericEpithet","start":13,"end":23},{"wordType":"specificEpithet","start":25,"end":35},{"wordType":"rank","start":36,"end":40},{"wordType":"infraspecificEpithet","start":40,"end":50}],"id":"86716b35-27ce-5d21-ab18-e8bb0c5d80be","parserVersion":"test_version"}
```

Name: Poa annua fo varia

Canonical: Poa annua f. varia

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Poa annua fo varia","normalized":"Poa annua f. varia","canonical":{"stemmed":"Poa annu uar","simple":"Poa annua varia","full":"Poa annua f. varia"},"cardinality":3,"details":{"infraSpecies":{"genus":"Poa","species":"annua","infraSpecies":[{"value":"varia","rank":"f."}]}},"pos":[{"wordType":"genus","start":0,"end":3},{"wordType":"specificEpithet","start":4,"end":9},{"wordType":"rank","start":10,"end":12},{"wordType":"infraspecificEpithet","start":13,"end":18}],"id":"32838647-3c46-509b-a81b-62d24940845f","parserVersion":"test_version"}
```

Name: Physarum globuliferum forma. flavum Leontyev & Dudka

Canonical: Physarum globuliferum f. flavum

Authorship: Leontyev & Dudka

```json
{"parsed":true,"parseQuality":1,"verbatim":"Physarum globuliferum forma. flavum Leontyev \u0026 Dudka","normalized":"Physarum globuliferum f. flavum Leontyev \u0026 Dudka","canonical":{"stemmed":"Physarum globulifer flau","simple":"Physarum globuliferum flavum","full":"Physarum globuliferum f. flavum"},"cardinality":3,"authorship":{"verbatim":"Leontyev \u0026 Dudka","normalized":"Leontyev \u0026 Dudka","authors":["Leontyev","Dudka"],"originalAuth":{"authors":["Leontyev","Dudka"]}},"details":{"infraSpecies":{"genus":"Physarum","species":"globuliferum","infraSpecies":[{"value":"flavum","rank":"f.","authorship":{"verbatim":"Leontyev \u0026 Dudka","normalized":"Leontyev \u0026 Dudka","authors":["Leontyev","Dudka"],"originalAuth":{"authors":["Leontyev","Dudka"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"rank","start":22,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":35},{"wordType":"authorWord","start":36,"end":44},{"wordType":"authorWord","start":47,"end":52}],"id":"bbcecb18-4484-528b-a8b9-93e1634d31b5","parserVersion":"test_version"}
```

Name: Homalanthus nutans (Mull.Arg.) Benth. & Hook. f. ex Drake

Canonical: Homalanthus nutans

Authorship: (Mull. Arg.) Benth. & Hook. fil. ex Drake

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Homalanthus nutans (Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"Homalanthus nutans (Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","canonical":{"stemmed":"Homalanthus nutans","simple":"Homalanthus nutans","full":"Homalanthus nutans"},"cardinality":2,"authorship":{"verbatim":"(Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"(Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","authors":["Mull. Arg.","Benth.","Hook. fil."],"originalAuth":{"authors":["Mull. Arg."]},"combinationAuth":{"authors":["Benth.","Hook. fil."],"exAuthors":{"authors":["Drake"]}}},"details":{"species":{"genus":"Homalanthus","species":"nutans","authorship":{"verbatim":"(Mull.Arg.) Benth. \u0026 Hook. f. ex Drake","normalized":"(Mull. Arg.) Benth. \u0026 Hook. fil. ex Drake","authors":["Mull. Arg.","Benth.","Hook. fil."],"originalAuth":{"authors":["Mull. Arg."]},"combinationAuth":{"authors":["Benth.","Hook. fil."],"exAuthors":{"authors":["Drake"]}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":18},{"wordType":"authorWord","start":20,"end":25},{"wordType":"authorWord","start":25,"end":29},{"wordType":"authorWord","start":31,"end":37},{"wordType":"authorWord","start":40,"end":45},{"wordType":"authorWordFilius","start":46,"end":48},{"wordType":"authorWord","start":52,"end":57}],"id":"83c06d35-e323-5750-84fb-f8c184fd1ee4","parserVersion":"test_version"}
```

Name: Calicium furfuraceum * furfuraceum (L.) Pers. 1797

Canonical: Calicium furfuraceum * furfuraceum

Authorship: (L.) Pers. 1797

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Calicium furfuraceum * furfuraceum (L.) Pers. 1797","normalized":"Calicium furfuraceum * furfuraceum (L.) Pers. 1797","canonical":{"stemmed":"Calicium furfurace furfurace","simple":"Calicium furfuraceum furfuraceum","full":"Calicium furfuraceum * furfuraceum"},"cardinality":3,"authorship":{"verbatim":"(L.) Pers. 1797","normalized":"(L.) Pers. 1797","authors":["L.","Pers."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Pers."],"year":{"year":"1797"}}},"details":{"infraSpecies":{"genus":"Calicium","species":"furfuraceum","infraSpecies":[{"value":"furfuraceum","rank":"*","authorship":{"verbatim":"(L.) Pers. 1797","normalized":"(L.) Pers. 1797","authors":["L.","Pers."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Pers."],"year":{"year":"1797"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"rank","start":21,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":34},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":40,"end":45},{"wordType":"year","start":46,"end":50}],"id":"6c5da8ae-cc50-5ce3-835d-d42e16aa0757","parserVersion":"test_version"}
```

Name: Polyrhachis orsyllus nat musculus Forel 1901

Canonical: Polyrhachis orsyllus nat musculus

Authorship: Forel 1901

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Uncommon rank"}],"verbatim":"Polyrhachis orsyllus nat musculus Forel 1901","normalized":"Polyrhachis orsyllus nat musculus Forel 1901","canonical":{"stemmed":"Polyrhachis orsyll muscul","simple":"Polyrhachis orsyllus musculus","full":"Polyrhachis orsyllus nat musculus"},"cardinality":3,"authorship":{"verbatim":"Forel 1901","normalized":"Forel 1901","year":"1901","authors":["Forel"],"originalAuth":{"authors":["Forel"],"year":{"year":"1901"}}},"details":{"infraSpecies":{"genus":"Polyrhachis","species":"orsyllus","infraSpecies":[{"value":"musculus","rank":"nat","authorship":{"verbatim":"Forel 1901","normalized":"Forel 1901","year":"1901","authors":["Forel"],"originalAuth":{"authors":["Forel"],"year":{"year":"1901"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":20},{"wordType":"rank","start":21,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33},{"wordType":"authorWord","start":34,"end":39},{"wordType":"year","start":40,"end":44}],"id":"3392132e-3dba-5b7e-a7c9-e4a68954c8b2","parserVersion":"test_version"}
```

Name: Acidalia remutaria ab. n. undularia

Canonical: Acidalia remutaria ab. n. undularia

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Acidalia remutaria ab. n. undularia","normalized":"Acidalia remutaria ab. n. undularia","canonical":{"stemmed":"Acidalia remutar undular","simple":"Acidalia remutaria undularia","full":"Acidalia remutaria ab. n. undularia"},"cardinality":3,"details":{"infraSpecies":{"genus":"Acidalia","species":"remutaria","infraSpecies":[{"value":"undularia","rank":"ab. n."}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"rank","start":19,"end":25},{"wordType":"infraspecificEpithet","start":26,"end":35}],"id":"ac834e3e-b861-5fbf-9cf9-197ad3effb99","parserVersion":"test_version"}
```

Name: Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius, 1912

Canonical: Acmaeops bivittata ab. fusciceps

Authorship: Aurivillius 1912

```json
{"parsed":true,"parseQuality":1,"verbatim":"Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius, 1912","normalized":"Acmaeops (Pseudodinoptera) bivittata ab. fusciceps Aurivillius 1912","canonical":{"stemmed":"Acmaeops biuittat fusciceps","simple":"Acmaeops bivittata fusciceps","full":"Acmaeops bivittata ab. fusciceps"},"cardinality":3,"authorship":{"verbatim":"Aurivillius, 1912","normalized":"Aurivillius 1912","year":"1912","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1912"}}},"details":{"infraSpecies":{"genus":"Acmaeops","subGenus":"Pseudodinoptera","species":"bivittata","infraSpecies":[{"value":"fusciceps","rank":"ab.","authorship":{"verbatim":"Aurivillius, 1912","normalized":"Aurivillius 1912","year":"1912","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1912"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"infragenericEpithet","start":10,"end":25},{"wordType":"specificEpithet","start":27,"end":36},{"wordType":"rank","start":37,"end":40},{"wordType":"infraspecificEpithet","start":41,"end":50},{"wordType":"authorWord","start":51,"end":62},{"wordType":"year","start":64,"end":68}],"id":"3f3dfc38-f660-56d6-a4f8-568f84a6878a","parserVersion":"test_version"}
```

### Infraspecies multiple (ICN)

Name: Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall & D.E. Stuntz 1972

Canonical: Hydnellum scrobiculatum var. zonatum f. parvum

Authorship: (Banker) D. Hall & D. E. Stuntz 1972

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"Hydnellum scrobiculatum var. zonatum f. parvum (Banker) D. Hall \u0026 D. E. Stuntz 1972","canonical":{"stemmed":"Hydnellum scrobiculat zonat paru","simple":"Hydnellum scrobiculatum zonatum parvum","full":"Hydnellum scrobiculatum var. zonatum f. parvum"},"cardinality":4,"authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}},"details":{"infraSpecies":{"genus":"Hydnellum","species":"scrobiculatum","infraSpecies":[{"value":"zonatum","rank":"var."},{"value":"parvum","rank":"f.","authorship":{"verbatim":"(Banker) D. Hall \u0026 D.E. Stuntz 1972","normalized":"(Banker) D. Hall \u0026 D. E. Stuntz 1972","authors":["Banker","D. Hall","D. E. Stuntz"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall","D. E. Stuntz"],"year":{"year":"1972"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":23},{"wordType":"rank","start":24,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":36},{"wordType":"rank","start":37,"end":39},{"wordType":"infraspecificEpithet","start":40,"end":46},{"wordType":"authorWord","start":48,"end":54},{"wordType":"authorWord","start":56,"end":58},{"wordType":"authorWord","start":59,"end":63},{"wordType":"authorWord","start":66,"end":68},{"wordType":"authorWord","start":68,"end":70},{"wordType":"authorWord","start":71,"end":77},{"wordType":"year","start":78,"end":82}],"id":"805654ed-0115-5f3e-af92-5808f215afbf","parserVersion":"test_version"}
```

Name: Senecio fuchsii C.C.Gmel. subsp. fuchsii var. expansus (Boiss. & Heldr.) Hayek

Canonical: Senecio fuchsii subsp. fuchsii var. expansus

Authorship: (Boiss. & Heldr.) Hayek

```json
{"parsed":true,"parseQuality":1,"verbatim":"Senecio fuchsii C.C.Gmel. subsp. fuchsii var. expansus (Boiss. \u0026 Heldr.) Hayek","normalized":"Senecio fuchsii C. C. Gmel. subsp. fuchsii var. expansus (Boiss. \u0026 Heldr.) Hayek","canonical":{"stemmed":"Senecio fuchsi fuchsi expans","simple":"Senecio fuchsii fuchsii expansus","full":"Senecio fuchsii subsp. fuchsii var. expansus"},"cardinality":4,"authorship":{"verbatim":"(Boiss. \u0026 Heldr.) Hayek","normalized":"(Boiss. \u0026 Heldr.) Hayek","authors":["Boiss.","Heldr.","Hayek"],"originalAuth":{"authors":["Boiss.","Heldr."]},"combinationAuth":{"authors":["Hayek"]}},"details":{"infraSpecies":{"genus":"Senecio","species":"fuchsii","authorship":{"verbatim":"C.C.Gmel.","normalized":"C. C. Gmel.","authors":["C. C. Gmel."],"originalAuth":{"authors":["C. C. Gmel."]}},"infraSpecies":[{"value":"fuchsii","rank":"subsp."},{"value":"expansus","rank":"var.","authorship":{"verbatim":"(Boiss. \u0026 Heldr.) Hayek","normalized":"(Boiss. \u0026 Heldr.) Hayek","authors":["Boiss.","Heldr.","Hayek"],"originalAuth":{"authors":["Boiss.","Heldr."]},"combinationAuth":{"authors":["Hayek"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":18,"end":20},{"wordType":"authorWord","start":20,"end":25},{"wordType":"rank","start":26,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":40},{"wordType":"rank","start":41,"end":45},{"wordType":"infraspecificEpithet","start":46,"end":54},{"wordType":"authorWord","start":56,"end":62},{"wordType":"authorWord","start":65,"end":71},{"wordType":"authorWord","start":73,"end":78}],"id":"93ed1df3-5016-56e7-8aa8-3a01df49a11a","parserVersion":"test_version"}
```

Name: Senecio fuchsii C.C.Gmel. subsp. fuchsii var. fuchsii

Canonical: Senecio fuchsii subsp. fuchsii var. fuchsii

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Senecio fuchsii C.C.Gmel. subsp. fuchsii var. fuchsii","normalized":"Senecio fuchsii C. C. Gmel. subsp. fuchsii var. fuchsii","canonical":{"stemmed":"Senecio fuchsi fuchsi fuchsi","simple":"Senecio fuchsii fuchsii fuchsii","full":"Senecio fuchsii subsp. fuchsii var. fuchsii"},"cardinality":4,"details":{"infraSpecies":{"genus":"Senecio","species":"fuchsii","authorship":{"verbatim":"C.C.Gmel.","normalized":"C. C. Gmel.","authors":["C. C. Gmel."],"originalAuth":{"authors":["C. C. Gmel."]}},"infraSpecies":[{"value":"fuchsii","rank":"subsp."},{"value":"fuchsii","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":18,"end":20},{"wordType":"authorWord","start":20,"end":25},{"wordType":"rank","start":26,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":40},{"wordType":"rank","start":41,"end":45},{"wordType":"infraspecificEpithet","start":46,"end":53}],"id":"481c3fc6-6f0c-55fa-b119-64d78d0bde03","parserVersion":"test_version"}
```

Name: Euastrum divergens var. rhodesiense f. coronulum A.M. Scott & Prescott

Canonical: Euastrum divergens var. rhodesiense f. coronulum

Authorship: A. M. Scott & Prescott

```json
{"parsed":true,"parseQuality":1,"verbatim":"Euastrum divergens var. rhodesiense f. coronulum A.M. Scott \u0026 Prescott","normalized":"Euastrum divergens var. rhodesiense f. coronulum A. M. Scott \u0026 Prescott","canonical":{"stemmed":"Euastrum diuergens rhodesiens coronul","simple":"Euastrum divergens rhodesiense coronulum","full":"Euastrum divergens var. rhodesiense f. coronulum"},"cardinality":4,"authorship":{"verbatim":"A.M. Scott \u0026 Prescott","normalized":"A. M. Scott \u0026 Prescott","authors":["A. M. Scott","Prescott"],"originalAuth":{"authors":["A. M. Scott","Prescott"]}},"details":{"infraSpecies":{"genus":"Euastrum","species":"divergens","infraSpecies":[{"value":"rhodesiense","rank":"var."},{"value":"coronulum","rank":"f.","authorship":{"verbatim":"A.M. Scott \u0026 Prescott","normalized":"A. M. Scott \u0026 Prescott","authors":["A. M. Scott","Prescott"],"originalAuth":{"authors":["A. M. Scott","Prescott"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"rank","start":19,"end":23},{"wordType":"infraspecificEpithet","start":24,"end":35},{"wordType":"rank","start":36,"end":38},{"wordType":"infraspecificEpithet","start":39,"end":48},{"wordType":"authorWord","start":49,"end":51},{"wordType":"authorWord","start":51,"end":53},{"wordType":"authorWord","start":54,"end":59},{"wordType":"authorWord","start":62,"end":70}],"id":"3e5a8eed-9f34-5f2b-95b5-1a45740e4306","parserVersion":"test_version"}
```

### Infraspecies with greek letters (ICN)

Name: Aristotelia fruticosa var. δ. microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var. δ. microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraSpecies":{"genus":"Aristotelia","species":"fruticosa","infraSpecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"rank","start":22,"end":26},{"wordType":"infraspecificEpithet","start":30,"end":41},{"wordType":"authorWord","start":42,"end":47},{"wordType":"authorWordFilius","start":47,"end":49}],"id":"34378b1d-27ef-5a38-a3ad-b2da249bc9d4","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var. δ microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var. δ microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraSpecies":{"genus":"Aristotelia","species":"fruticosa","infraSpecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"rank","start":22,"end":26},{"wordType":"infraspecificEpithet","start":29,"end":40},{"wordType":"authorWord","start":41,"end":46},{"wordType":"authorWordFilius","start":46,"end":48}],"id":"d31a653a-8686-5bf4-b657-6164f494e6b4","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var.δ.microphylla Hook.f.

Canonical: Aristotelia fruticosa var. microphylla

Authorship: Hook. fil.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Deprecated Greek letter enumeration in rank"}],"verbatim":"Aristotelia fruticosa var.δ.microphylla Hook.f.","normalized":"Aristotelia fruticosa var. microphylla Hook. fil.","canonical":{"stemmed":"Aristotelia fruticos microphyll","simple":"Aristotelia fruticosa microphylla","full":"Aristotelia fruticosa var. microphylla"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraSpecies":{"genus":"Aristotelia","species":"fruticosa","infraSpecies":[{"value":"microphylla","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"rank","start":22,"end":26},{"wordType":"infraspecificEpithet","start":28,"end":39},{"wordType":"authorWord","start":40,"end":45},{"wordType":"authorWordFilius","start":45,"end":47}],"id":"c2f051e5-c1a2-52f8-a02f-70510030faa1","parserVersion":"test_version"}
```

Name: Aristotelia fruticosa var. δmicrophylla Hook.f.

Canonical: Aristotelia fruticosa

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aristotelia fruticosa var. δmicrophylla Hook.f.","normalized":"Aristotelia fruticosa","canonical":{"stemmed":"Aristotelia fruticos","simple":"Aristotelia fruticosa","full":"Aristotelia fruticosa"},"cardinality":2,"tail":" var. δmicrophylla Hook.f.","details":{"species":{"genus":"Aristotelia","species":"fruticosa"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21}],"id":"f7749c21-82a6-5c42-ab58-7b3d5a824e96","parserVersion":"test_version"}
```

### Hybrids with notho- ranks

Name: Crataegus curvisepala nvar. naviculiformis T. Petauer

Canonical: Crataegus curvisepala nvar. naviculiformis

Authorship: T. Petauer

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Crataegus curvisepala nvar. naviculiformis T. Petauer","normalized":"Crataegus curvisepala nvar. naviculiformis T. Petauer","canonical":{"stemmed":"Crataegus curuisepal nauiculiform","simple":"Crataegus curvisepala naviculiformis","full":"Crataegus curvisepala nvar. naviculiformis"},"cardinality":3,"authorship":{"verbatim":"T. Petauer","normalized":"T. Petauer","authors":["T. Petauer"],"originalAuth":{"authors":["T. Petauer"]}},"hybrid":"NOTHO_HYBRID","details":{"infraSpecies":{"genus":"Crataegus","species":"curvisepala","infraSpecies":[{"value":"naviculiformis","rank":"nvar.","authorship":{"verbatim":"T. Petauer","normalized":"T. Petauer","authors":["T. Petauer"],"originalAuth":{"authors":["T. Petauer"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":21},{"wordType":"rank","start":22,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":42},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":53}],"id":"f3e2ccac-4844-57a7-8903-4e3b6a0d0851","parserVersion":"test_version"}
```

Name: Aconitum W. Mucher nothosect. Acopellus

Canonical: Aconitum nothosect. Acopellus

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aconitum W. Mucher nothosect. Acopellus","normalized":"Aconitum nothosect. Acopellus","canonical":{"stemmed":"Acopellus","simple":"Acopellus","full":"Aconitum nothosect. Acopellus"},"cardinality":1,"hybrid":"NOTHO_HYBRID","details":{"uninomial":{"uninomial":"Acopellus","rank":"nothosect.","parent":"Aconitum"}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":11},{"wordType":"authorWord","start":12,"end":18},{"wordType":"rank","start":19,"end":29},{"wordType":"uninomial","start":30,"end":39}],"id":"815f38e4-2425-551d-b054-4949a457d6a6","parserVersion":"test_version"}
```

Name: Aconitum W. Mucher nothoser. Acotoxicum

Canonical: Aconitum nothoser. Acotoxicum

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Aconitum W. Mucher nothoser. Acotoxicum","normalized":"Aconitum nothoser. Acotoxicum","canonical":{"stemmed":"Acotoxicum","simple":"Acotoxicum","full":"Aconitum nothoser. Acotoxicum"},"cardinality":1,"hybrid":"NOTHO_HYBRID","details":{"uninomial":{"uninomial":"Acotoxicum","rank":"nothoser.","parent":"Aconitum"}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":11},{"wordType":"authorWord","start":12,"end":18},{"wordType":"rank","start":19,"end":28},{"wordType":"uninomial","start":29,"end":39}],"id":"6fd8d3d4-bdb6-5fc6-a94d-966af669c7e9","parserVersion":"test_version"}
```

Name: Abies masjoannis nothof. mesoides

Canonical: Abies masjoannis nothof. mesoides

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Abies masjoannis nothof. mesoides","normalized":"Abies masjoannis nothof. mesoides","canonical":{"stemmed":"Abies masioann mesoid","simple":"Abies masjoannis mesoides","full":"Abies masjoannis nothof. mesoides"},"cardinality":3,"hybrid":"NOTHO_HYBRID","details":{"infraSpecies":{"genus":"Abies","species":"masjoannis","infraSpecies":[{"value":"mesoides","rank":"nothof."}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":16},{"wordType":"rank","start":17,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33}],"id":"5be2cd2f-c81f-5d81-8eaf-54bd231f5230","parserVersion":"test_version"}
```

Name: Aconitum berdaui nothosubsp. walasii (Mitka) Mitka

Canonical: Aconitum berdaui nothosubsp. walasii

Authorship: (Mitka) Mitka

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum berdaui nothosubsp. walasii (Mitka) Mitka","normalized":"Aconitum berdaui nothosubsp. walasii (Mitka) Mitka","canonical":{"stemmed":"Aconitum berdau walasi","simple":"Aconitum berdaui walasii","full":"Aconitum berdaui nothosubsp. walasii"},"cardinality":3,"authorship":{"verbatim":"(Mitka) Mitka","normalized":"(Mitka) Mitka","authors":["Mitka","Mitka"],"originalAuth":{"authors":["Mitka"]},"combinationAuth":{"authors":["Mitka"]}},"hybrid":"NOTHO_HYBRID","details":{"infraSpecies":{"genus":"Aconitum","species":"berdaui","infraSpecies":[{"value":"walasii","rank":"nothosubsp.","authorship":{"verbatim":"(Mitka) Mitka","normalized":"(Mitka) Mitka","authors":["Mitka","Mitka"],"originalAuth":{"authors":["Mitka"]},"combinationAuth":{"authors":["Mitka"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"rank","start":17,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":36},{"wordType":"authorWord","start":38,"end":43},{"wordType":"authorWord","start":45,"end":50}],"id":"ba2f82ac-9312-5595-928a-2ba07aebb04f","parserVersion":"test_version"}
```

Name: Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu

Canonical: Aconitum tauricum nothossp. hayekianum

Authorship: (Gáyer) Grintescu

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu","normalized":"Aconitum tauricum nothossp. hayekianum (Gáyer) Grintescu","canonical":{"stemmed":"Aconitum tauric hayekian","simple":"Aconitum tauricum hayekianum","full":"Aconitum tauricum nothossp. hayekianum"},"cardinality":3,"authorship":{"verbatim":"(Gáyer) Grintescu","normalized":"(Gáyer) Grintescu","authors":["Gáyer","Grintescu"],"originalAuth":{"authors":["Gáyer"]},"combinationAuth":{"authors":["Grintescu"]}},"hybrid":"NOTHO_HYBRID","details":{"infraSpecies":{"genus":"Aconitum","species":"tauricum","infraSpecies":[{"value":"hayekianum","rank":"nothossp.","authorship":{"verbatim":"(Gáyer) Grintescu","normalized":"(Gáyer) Grintescu","authors":["Gáyer","Grintescu"],"originalAuth":{"authors":["Gáyer"]},"combinationAuth":{"authors":["Grintescu"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"rank","start":18,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":38},{"wordType":"authorWord","start":40,"end":45},{"wordType":"authorWord","start":47,"end":56}],"id":"c02c80bf-11b1-59f9-9fed-6627fb954dd8","parserVersion":"test_version"}
```

Name: Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares

Canonical: Aeonium holospathulatum nothovar. sanchezii

Authorship: (Bañares) Bañares

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares","normalized":"Aeonium holospathulatum nothovar. sanchezii (Bañares) Bañares","canonical":{"stemmed":"Aeonium holospathulat sanchezi","simple":"Aeonium holospathulatum sanchezii","full":"Aeonium holospathulatum nothovar. sanchezii"},"cardinality":3,"authorship":{"verbatim":"(Bañares) Bañares","normalized":"(Bañares) Bañares","authors":["Bañares","Bañares"],"originalAuth":{"authors":["Bañares"]},"combinationAuth":{"authors":["Bañares"]}},"hybrid":"NOTHO_HYBRID","details":{"infraSpecies":{"genus":"Aeonium","species":"holospathulatum","infraSpecies":[{"value":"sanchezii","rank":"nothovar.","authorship":{"verbatim":"(Bañares) Bañares","normalized":"(Bañares) Bañares","authors":["Bañares","Bañares"],"originalAuth":{"authors":["Bañares"]},"combinationAuth":{"authors":["Bañares"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":23},{"wordType":"rank","start":24,"end":33},{"wordType":"infraspecificEpithet","start":34,"end":43},{"wordType":"authorWord","start":45,"end":52},{"wordType":"authorWord","start":54,"end":61}],"id":"fc173db1-3977-5cad-a96b-472165bb0bbd","parserVersion":"test_version"}
```

Name: Amaranthus ×ozanonii (Contré) Lambinon nothosubsp. ralletii

Canonical: Amaranthus × ozanonii nothosubsp. ralletii

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Amaranthus ×ozanonii (Contré) Lambinon nothosubsp. ralletii","normalized":"Amaranthus × ozanonii (Contré) Lambinon nothosubsp. ralletii","canonical":{"stemmed":"Amaranthus ozanoni ralleti","simple":"Amaranthus ozanonii ralletii","full":"Amaranthus × ozanonii nothosubsp. ralletii"},"cardinality":3,"hybrid":"NAMED_HYBRID","details":{"infraSpecies":{"genus":"Amaranthus","species":"ozanonii (Contré) Lambinon","authorship":{"verbatim":"(Contré) Lambinon","normalized":"(Contré) Lambinon","authors":["Contré","Lambinon"],"originalAuth":{"authors":["Contré"]},"combinationAuth":{"authors":["Lambinon"]}},"infraSpecies":[{"value":"ralletii","rank":"nothosubsp."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"hybridChar","start":11,"end":12},{"wordType":"specificEpithet","start":12,"end":20},{"wordType":"authorWord","start":22,"end":28},{"wordType":"authorWord","start":30,"end":38},{"wordType":"rank","start":39,"end":50},{"wordType":"infraspecificEpithet","start":51,"end":59}],"id":"678535c6-c679-5716-a874-1cf92bca3ce9","parserVersion":"test_version"}
```

Name: Aconitum ×teppneri Mucher ex Starm. nothosubsp. goetzii

Canonical: Aconitum × teppneri nothosubsp. goetzii

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aconitum ×teppneri Mucher ex Starm. nothosubsp. goetzii","normalized":"Aconitum × teppneri Mucher ex Starm. nothosubsp. goetzii","canonical":{"stemmed":"Aconitum teppner goetzi","simple":"Aconitum teppneri goetzii","full":"Aconitum × teppneri nothosubsp. goetzii"},"cardinality":3,"hybrid":"NAMED_HYBRID","details":{"infraSpecies":{"genus":"Aconitum","species":"teppneri Mucher ex Starm.","authorship":{"verbatim":"Mucher ex Starm.","normalized":"Mucher ex Starm.","authors":["Mucher"],"originalAuth":{"authors":["Mucher"],"exAuthors":{"authors":["Starm."]}}},"infraSpecies":[{"value":"goetzii","rank":"nothosubsp."}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"hybridChar","start":9,"end":10},{"wordType":"specificEpithet","start":10,"end":18},{"wordType":"authorWord","start":19,"end":25},{"wordType":"authorWord","start":29,"end":35},{"wordType":"rank","start":36,"end":47},{"wordType":"infraspecificEpithet","start":48,"end":55}],"id":"2387941b-9e4f-5fb5-a440-74934cc66c4f","parserVersion":"test_version"}
```

Name: Aeonium × proliferum Bañares nothovar. glabrifolium Bañares

Canonical: Aeonium × proliferum nothovar. glabrifolium

Authorship: Bañares

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aeonium × proliferum Bañares nothovar. glabrifolium Bañares","normalized":"Aeonium × proliferum Bañares nothovar. glabrifolium Bañares","canonical":{"stemmed":"Aeonium prolifer glabrifoli","simple":"Aeonium proliferum glabrifolium","full":"Aeonium × proliferum nothovar. glabrifolium"},"cardinality":3,"authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}},"hybrid":"NAMED_HYBRID","details":{"infraSpecies":{"genus":"Aeonium","species":"proliferum Bañares","authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}},"infraSpecies":[{"value":"glabrifolium","rank":"nothovar.","authorship":{"verbatim":"Bañares","normalized":"Bañares","authors":["Bañares"],"originalAuth":{"authors":["Bañares"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"hybridChar","start":8,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"authorWord","start":21,"end":28},{"wordType":"rank","start":29,"end":38},{"wordType":"infraspecificEpithet","start":39,"end":51},{"wordType":"authorWord","start":52,"end":59}],"id":"dc38d07a-f949-5c72-9463-d36a4ae96bea","parserVersion":"test_version"}
```

<!-- Very rare people make this mistake. We do not cover it yet.
Agropyron x pseudorepens notho morph. vulpinum (Rydb.) Bowden, 1965
-->

Name: Biscogniauxia nothofagi Whalley, Læssøe & Kile 1990

Canonical: Biscogniauxia nothofagi

Authorship: Whalley, Læssøe & Kile 1990

```json
{"parsed":true,"parseQuality":1,"verbatim":"Biscogniauxia nothofagi Whalley, Læssøe \u0026 Kile 1990","normalized":"Biscogniauxia nothofagi Whalley, Læssøe \u0026 Kile 1990","canonical":{"stemmed":"Biscogniauxia nothofag","simple":"Biscogniauxia nothofagi","full":"Biscogniauxia nothofagi"},"cardinality":2,"authorship":{"verbatim":"Whalley, Læssøe \u0026 Kile 1990","normalized":"Whalley, Læssøe \u0026 Kile 1990","year":"1990","authors":["Whalley","Læssøe","Kile"],"originalAuth":{"authors":["Whalley","Læssøe","Kile"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Biscogniauxia","species":"nothofagi","authorship":{"verbatim":"Whalley, Læssøe \u0026 Kile 1990","normalized":"Whalley, Læssøe \u0026 Kile 1990","year":"1990","authors":["Whalley","Læssøe","Kile"],"originalAuth":{"authors":["Whalley","Læssøe","Kile"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":23},{"wordType":"authorWord","start":24,"end":31},{"wordType":"authorWord","start":33,"end":39},{"wordType":"authorWord","start":42,"end":46},{"wordType":"year","start":47,"end":51}],"id":"1f8935ad-5ae2-507e-96aa-f0bb1d22245e","parserVersion":"test_version"}
```

### Named hybrids
Name: ×Agropogon P. Fourn. 1934

Canonical: × Agropogon

Authorship: P. Fourn. 1934

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Agropogon P. Fourn. 1934","normalized":"× Agropogon P. Fourn. 1934","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P. Fourn. 1934","normalized":"P. Fourn. 1934","year":"1934","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."],"year":{"year":"1934"}}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P. Fourn. 1934","normalized":"P. Fourn. 1934","year":"1934","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."],"year":{"year":"1934"}}}}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":1,"end":10},{"wordType":"authorWord","start":11,"end":13},{"wordType":"authorWord","start":14,"end":20},{"wordType":"year","start":21,"end":25}],"id":"f2bb2ddc-003e-5fc0-83b1-038dca1deb52","parserVersion":"test_version"}
```

Name: xAgropogon P. Fourn.

Canonical: × Agropogon

Authorship: P. Fourn.

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"xAgropogon P. Fourn.","normalized":"× Agropogon P. Fourn.","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P. Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P. Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}}}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":1,"end":10},{"wordType":"authorWord","start":11,"end":13},{"wordType":"authorWord","start":14,"end":20}],"id":"b36871e3-e412-5b4f-a859-eb09fcf83a8e","parserVersion":"test_version"}
```

Name: XAgropogon P.Fourn.

Canonical: × Agropogon

Authorship: P. Fourn.

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"XAgropogon P.Fourn.","normalized":"× Agropogon P. Fourn.","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"authorship":{"verbatim":"P.Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}},"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon","authorship":{"verbatim":"P.Fourn.","normalized":"P. Fourn.","authors":["P. Fourn."],"originalAuth":{"authors":["P. Fourn."]}}}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":1,"end":10},{"wordType":"authorWord","start":11,"end":13},{"wordType":"authorWord","start":13,"end":19}],"id":"f6257985-ad38-5c29-94e2-bb305cab893a","parserVersion":"test_version"}
```

Name: × Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"× Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":2,"end":11}],"id":"b1858609-4fff-5a00-8d2b-0cb354100b10","parserVersion":"test_version"}
```

Name: x Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"x Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":2,"end":11}],"id":"79c27436-a61a-59cd-acf0-51425556e26f","parserVersion":"test_version"}
```

Name: X Agropogon

Canonical: × Agropogon

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Agropogon","normalized":"× Agropogon","canonical":{"stemmed":"Agropogon","simple":"Agropogon","full":"× Agropogon"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agropogon"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":2,"end":11}],"id":"37eac7d5-a258-503b-ae3f-206739be74fa","parserVersion":"test_version"}
```

Name: X Cupressocyparis leylandii

Canonical: × Cupressocyparis leylandii

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Cupressocyparis leylandii","normalized":"× Cupressocyparis leylandii","canonical":{"stemmed":"Cupressocyparis leylandi","simple":"Cupressocyparis leylandii","full":"× Cupressocyparis leylandii"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Cupressocyparis","species":"leylandii"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":2,"end":17},{"wordType":"specificEpithet","start":18,"end":27}],"id":"a6ebd2cf-a021-50fe-b158-8be16844079d","parserVersion":"test_version"}
```

Name: ×Heucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Heucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":1,"end":12},{"wordType":"specificEpithet","start":13,"end":25}],"id":"6aab4b31-89fb-5a41-97ee-2024becc9169","parserVersion":"test_version"}
```

Name: xHeucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"xHeucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":1,"end":12},{"wordType":"specificEpithet","start":13,"end":25}],"id":"726d4f33-a175-5449-aea2-0e3c26dc7a0b","parserVersion":"test_version"}
```

Name: x Heucherella tiarelloides

Canonical: × Heucherella tiarelloides

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"x Heucherella tiarelloides","normalized":"× Heucherella tiarelloides","canonical":{"stemmed":"Heucherella tiarelloid","simple":"Heucherella tiarelloides","full":"× Heucherella tiarelloides"},"cardinality":2,"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Heucherella","species":"tiarelloides"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":2,"end":13},{"wordType":"specificEpithet","start":14,"end":26}],"id":"da549587-a768-51b6-af26-1bb3c1977b31","parserVersion":"test_version"}
```

Name: XAgroelymus Lapage sect. Agroelinelymus

Canonical: × Agroelymus sect. Agroelinelymus

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"XAgroelymus Lapage sect. Agroelinelymus","normalized":"× Agroelymus sect. Agroelinelymus","canonical":{"stemmed":"Agroelinelymus","simple":"Agroelinelymus","full":"× Agroelymus sect. Agroelinelymus"},"cardinality":1,"hybrid":"NAMED_HYBRID","details":{"uninomial":{"uninomial":"Agroelinelymus","rank":"sect.","parent":"Agroelymus"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":1,"end":11},{"wordType":"authorWord","start":12,"end":18},{"wordType":"rank","start":19,"end":24},{"wordType":"uninomial","start":25,"end":39}],"id":"419d1a5d-64b9-5e0d-87f4-624b19ddab0f","parserVersion":"test_version"}
```

Name: ×Agropogon littoralis (Sm.) C. E. Hubb. 1946

Canonical: × Agropogon littoralis

Authorship: (Sm.) C. E. Hubb. 1946

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"×Agropogon littoralis (Sm.) C. E. Hubb. 1946","normalized":"× Agropogon littoralis (Sm.) C. E. Hubb. 1946","canonical":{"stemmed":"Agropogon littoral","simple":"Agropogon littoralis","full":"× Agropogon littoralis"},"cardinality":2,"authorship":{"verbatim":"(Sm.) C. E. Hubb. 1946","normalized":"(Sm.) C. E. Hubb. 1946","authors":["Sm.","C. E. Hubb."],"originalAuth":{"authors":["Sm."]},"combinationAuth":{"authors":["C. E. Hubb."],"year":{"year":"1946"}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Agropogon","species":"littoralis","authorship":{"verbatim":"(Sm.) C. E. Hubb. 1946","normalized":"(Sm.) C. E. Hubb. 1946","authors":["Sm.","C. E. Hubb."],"originalAuth":{"authors":["Sm."]},"combinationAuth":{"authors":["C. E. Hubb."],"year":{"year":"1946"}}}}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":1,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":23,"end":26},{"wordType":"authorWord","start":28,"end":30},{"wordType":"authorWord","start":31,"end":33},{"wordType":"authorWord","start":34,"end":39},{"wordType":"year","start":40,"end":44}],"id":"66beda81-d796-5d60-be9f-b3188ef730dc","parserVersion":"test_version"}
```

Name: Asplenium X inexpectatum (E.L. Braun 1940) Morton (1956)

Canonical: Asplenium × inexpectatum

Authorship: (E. L. Braun 1940) Morton (1956)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Asplenium X inexpectatum (E.L. Braun 1940) Morton (1956)","normalized":"Asplenium × inexpectatum (E. L. Braun 1940) Morton (1956)","canonical":{"stemmed":"Asplenium inexpectat","simple":"Asplenium inexpectatum","full":"Asplenium × inexpectatum"},"cardinality":2,"authorship":{"verbatim":"(E.L. Braun 1940) Morton (1956)","normalized":"(E. L. Braun 1940) Morton (1956)","year":"1940","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1940"}},"combinationAuth":{"authors":["Morton"],"year":{"year":"1956","isApproximate":true}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Asplenium","species":"inexpectatum (E. L. Braun 1940) Morton (1956)","authorship":{"verbatim":"(E.L. Braun 1940) Morton (1956)","normalized":"(E. L. Braun 1940) Morton (1956)","year":"1940","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1940"}},"combinationAuth":{"authors":["Morton"],"year":{"year":"1956","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"hybridChar","start":10,"end":11},{"wordType":"specificEpithet","start":12,"end":24},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":28,"end":30},{"wordType":"authorWord","start":31,"end":36},{"wordType":"year","start":37,"end":41},{"wordType":"authorWord","start":43,"end":49},{"wordType":"approximateYear","start":51,"end":55}],"id":"d37e04e4-90bc-5031-b91c-dbb61113bcfa","parserVersion":"test_version"}
```

Name: Salix ×capreola Andersson (1867)

Canonical: Salix × capreola

Authorship: Andersson (1867)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Hybrid char not separated by space"},{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Year with parentheses"}],"verbatim":"Salix ×capreola Andersson (1867)","normalized":"Salix × capreola Andersson (1867)","canonical":{"stemmed":"Salix capreol","simple":"Salix capreola","full":"Salix × capreola"},"cardinality":2,"authorship":{"verbatim":"Andersson (1867)","normalized":"Andersson (1867)","year":"(1867)","authors":["Andersson"],"originalAuth":{"authors":["Andersson"],"year":{"year":"1867","isApproximate":true}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Salix","species":"capreola Andersson (1867)","authorship":{"verbatim":"Andersson (1867)","normalized":"Andersson (1867)","year":"(1867)","authors":["Andersson"],"originalAuth":{"authors":["Andersson"],"year":{"year":"1867","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"hybridChar","start":6,"end":7},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":25},{"wordType":"approximateYear","start":27,"end":31}],"id":"9965be0c-0db2-506a-97f7-e709ef950ef7","parserVersion":"test_version"}
```

Name: Polypodium  x vulgare nothosubsp. mantoniae (Rothm.) Schidlay

Canonical: Polypodium × vulgare nothosubsp. mantoniae

Authorship: (Rothm.) Schidlay

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"Polypodium  x vulgare nothosubsp. mantoniae (Rothm.) Schidlay","normalized":"Polypodium × vulgare nothosubsp. mantoniae (Rothm.) Schidlay","canonical":{"stemmed":"Polypodium uulgar mantoni","simple":"Polypodium vulgare mantoniae","full":"Polypodium × vulgare nothosubsp. mantoniae"},"cardinality":3,"authorship":{"verbatim":"(Rothm.) Schidlay","normalized":"(Rothm.) Schidlay","authors":["Rothm.","Schidlay"],"originalAuth":{"authors":["Rothm."]},"combinationAuth":{"authors":["Schidlay"]}},"hybrid":"NAMED_HYBRID","details":{"infraSpecies":{"genus":"Polypodium","species":"vulgare","infraSpecies":[{"value":"mantoniae","rank":"nothosubsp.","authorship":{"verbatim":"(Rothm.) Schidlay","normalized":"(Rothm.) Schidlay","authors":["Rothm.","Schidlay"],"originalAuth":{"authors":["Rothm."]},"combinationAuth":{"authors":["Schidlay"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"hybridChar","start":12,"end":13},{"wordType":"specificEpithet","start":14,"end":21},{"wordType":"rank","start":22,"end":33},{"wordType":"infraspecificEpithet","start":34,"end":43},{"wordType":"authorWord","start":45,"end":51},{"wordType":"authorWord","start":53,"end":61}],"id":"8666c370-8843-5324-a7f3-754ca778d618","parserVersion":"test_version"}
```

Name: Salix x capreola Andersson

Canonical: Salix × capreola

Authorship: Andersson

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Salix x capreola Andersson","normalized":"Salix × capreola Andersson","canonical":{"stemmed":"Salix capreol","simple":"Salix capreola","full":"Salix × capreola"},"cardinality":2,"authorship":{"verbatim":"Andersson","normalized":"Andersson","authors":["Andersson"],"originalAuth":{"authors":["Andersson"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Salix","species":"capreola Andersson","authorship":{"verbatim":"Andersson","normalized":"Andersson","authors":["Andersson"],"originalAuth":{"authors":["Andersson"]}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"hybridChar","start":6,"end":7},{"wordType":"specificEpithet","start":8,"end":16},{"wordType":"authorWord","start":17,"end":26}],"id":"5780473c-18ac-5386-9c3a-f74bbe426624","parserVersion":"test_version"}
```

### Hybrid formulae

Name: Stanhopea tigrina Bateman ex Lindl. x S. ecornuta Lem.

Canonical: Stanhopea tigrina × Stanhopea ecornuta

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Stanhopea tigrina Bateman ex Lindl. x S. ecornuta Lem.","normalized":"Stanhopea tigrina Bateman ex Lindl. × Stanhopea ecornuta Lem.","canonical":{"stemmed":"Stanhopea tigrin × Stanhope ecornut","simple":"Stanhopea tigrina × Stanhopea ecornuta","full":"Stanhopea tigrina × Stanhopea ecornuta"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Stanhopea","species":"tigrina","authorship":{"verbatim":"Bateman ex Lindl.","normalized":"Bateman ex Lindl.","authors":["Bateman"],"originalAuth":{"authors":["Bateman"],"exAuthors":{"authors":["Lindl."]}}}}},{"species":{"genus":"Stanhopea","species":"ecornuta","authorship":{"verbatim":"Lem.","normalized":"Lem.","authors":["Lem."],"originalAuth":{"authors":["Lem."]}}}}]},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":17},{"wordType":"authorWord","start":18,"end":25},{"wordType":"authorWord","start":29,"end":35},{"wordType":"hybridChar","start":36,"end":37},{"wordType":"genus","start":38,"end":40},{"wordType":"specificEpithet","start":41,"end":49},{"wordType":"authorWord","start":50,"end":54}],"id":"80c0a17d-3422-515c-88bc-3a927438df88","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora X Hydnellum scrobiculatum

Canonical: Arthopyrenia hyalospora × Hydnellum scrobiculatum

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Arthopyrenia hyalospora X Hydnellum scrobiculatum","normalized":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","canonical":{"stemmed":"Arthopyrenia hyalospor × Hydnell scrobiculat","simple":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","full":"Arthopyrenia hyalospora × Hydnellum scrobiculatum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}},{"species":{"genus":"Hydnellum","species":"scrobiculatum"}}]},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"hybridChar","start":24,"end":25},{"wordType":"genus","start":26,"end":35},{"wordType":"specificEpithet","start":36,"end":49}],"id":"e78d9299-9fd4-55d2-aeb4-b2864f5bff45","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Banker) D. Hall X Hydnellum scrobiculatum D.E. Stuntz

Canonical: Arthopyrenia hyalospora × Hydnellum scrobiculatum

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Arthopyrenia hyalospora (Banker) D. Hall X Hydnellum scrobiculatum D.E. Stuntz","normalized":"Arthopyrenia hyalospora (Banker) D. Hall × Hydnellum scrobiculatum D. E. Stuntz","canonical":{"stemmed":"Arthopyrenia hyalospor × Hydnell scrobiculat","simple":"Arthopyrenia hyalospora × Hydnellum scrobiculatum","full":"Arthopyrenia hyalospora × Hydnellum scrobiculatum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Banker) D. Hall","normalized":"(Banker) D. Hall","authors":["Banker","D. Hall"],"originalAuth":{"authors":["Banker"]},"combinationAuth":{"authors":["D. Hall"]}}}},{"species":{"genus":"Hydnellum","species":"scrobiculatum","authorship":{"verbatim":"D.E. Stuntz","normalized":"D. E. Stuntz","authors":["D. E. Stuntz"],"originalAuth":{"authors":["D. E. Stuntz"]}}}}]},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":31},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":36,"end":40},{"wordType":"hybridChar","start":41,"end":42},{"wordType":"genus","start":43,"end":52},{"wordType":"specificEpithet","start":53,"end":66},{"wordType":"authorWord","start":67,"end":69},{"wordType":"authorWord","start":69,"end":71},{"wordType":"authorWord","start":72,"end":78}],"id":"a13ac2e0-5eec-569c-af9c-dd8163dbbd72","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora x

Canonical: Arthopyrenia hyalospora ×

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"},{"quality":2,"warning":"Probably incomplete hybrid formula"}],"verbatim":"Arthopyrenia hyalospora x","normalized":"Arthopyrenia hyalospora ×","canonical":{"stemmed":"Arthopyrenia hyalospor ×","simple":"Arthopyrenia hyalospora ×","full":"Arthopyrenia hyalospora ×"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}}]},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"hybridChar","start":24,"end":25}],"id":"c056b89e-789b-5c28-89e7-e820ea0baebf","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora × ?

Canonical: Arthopyrenia hyalospora ×

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Hybrid formula"},{"quality":2,"warning":"Probably incomplete hybrid formula"}],"verbatim":"Arthopyrenia hyalospora × ?","normalized":"Arthopyrenia hyalospora ×","canonical":{"stemmed":"Arthopyrenia hyalospor ×","simple":"Arthopyrenia hyalospora ×","full":"Arthopyrenia hyalospora ×"},"cardinality":0,"hybrid":"HYBRID_FORMULA","tail":" ?","details":{"hybridFormula":[{"species":{"genus":"Arthopyrenia","species":"hyalospora"}}]},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"hybridChar","start":24,"end":25}],"id":"638cc013-3821-55c2-b9d3-b2ea3de33ecf","parserVersion":"test_version"}
```

Name: Agrostis L. × Polypogon Desf.

Canonical: Agrostis × Polypogon

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Agrostis L. × Polypogon Desf.","normalized":"Agrostis L. × Polypogon Desf.","canonical":{"stemmed":"Agrostis × Polypogon","simple":"Agrostis × Polypogon","full":"Agrostis × Polypogon"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"uninomial":{"uninomial":"Agrostis","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"uninomial":{"uninomial":"Polypogon","authorship":{"verbatim":"Desf.","normalized":"Desf.","authors":["Desf."],"originalAuth":{"authors":["Desf."]}}}}]},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":11},{"wordType":"hybridChar","start":12,"end":13},{"wordType":"uninomial","start":14,"end":23},{"wordType":"authorWord","start":24,"end":29}],"id":"e914b63f-f19a-5437-ad19-85bfc98a0de2","parserVersion":"test_version"}
```

Name: Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.

Canonical: Agrostis stolonifera × Polypogon monspeliensis

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.","normalized":"Agrostis stolonifera L. × Polypogon monspeliensis (L.) Desf.","canonical":{"stemmed":"Agrostis stolonifer × Polypogon monspeliens","simple":"Agrostis stolonifera × Polypogon monspeliensis","full":"Agrostis stolonifera × Polypogon monspeliensis"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Agrostis","species":"stolonifera","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"species":{"genus":"Polypogon","species":"monspeliensis","authorship":{"verbatim":"(L.) Desf.","normalized":"(L.) Desf.","authors":["L.","Desf."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Desf."]}}}}]},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":21,"end":23},{"wordType":"hybridChar","start":24,"end":25},{"wordType":"genus","start":26,"end":35},{"wordType":"specificEpithet","start":36,"end":49},{"wordType":"authorWord","start":51,"end":53},{"wordType":"authorWord","start":55,"end":60}],"id":"a2aeb842-18c5-54b4-a4d9-c78bd0445c10","parserVersion":"test_version"}
```

Name: Coeloglossum viride (L.) Hartman x Dactylorhiza majalis (Rchb. f.) P.F. Hunt & Summerhayes ssp. praetermissa (Druce) D.M. Moore & Soó

Canonical: Coeloglossum viride × Dactylorhiza majalis subsp. praetermissa

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Coeloglossum viride (L.) Hartman x Dactylorhiza majalis (Rchb. f.) P.F. Hunt \u0026 Summerhayes ssp. praetermissa (Druce) D.M. Moore \u0026 Soó","normalized":"Coeloglossum viride (L.) Hartman × Dactylorhiza majalis (Rchb. fil.) P. F. Hunt \u0026 Summerhayes subsp. praetermissa (Druce) D. M. Moore \u0026 Soó","canonical":{"stemmed":"Coeloglossum uirid × Dactylorhiz maial praetermiss","simple":"Coeloglossum viride × Dactylorhiza majalis praetermissa","full":"Coeloglossum viride × Dactylorhiza majalis subsp. praetermissa"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Coeloglossum","species":"viride","authorship":{"verbatim":"(L.) Hartman","normalized":"(L.) Hartman","authors":["L.","Hartman"],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["Hartman"]}}}},{"infraSpecies":{"genus":"Dactylorhiza","species":"majalis","authorship":{"verbatim":"(Rchb. f.) P.F. Hunt \u0026 Summerhayes","normalized":"(Rchb. fil.) P. F. Hunt \u0026 Summerhayes","authors":["Rchb. fil.","P. F. Hunt","Summerhayes"],"originalAuth":{"authors":["Rchb. fil."]},"combinationAuth":{"authors":["P. F. Hunt","Summerhayes"]}},"infraSpecies":[{"value":"praetermissa","rank":"subsp.","authorship":{"verbatim":"(Druce) D.M. Moore \u0026 Soó","normalized":"(Druce) D. M. Moore \u0026 Soó","authors":["Druce","D. M. Moore","Soó"],"originalAuth":{"authors":["Druce"]},"combinationAuth":{"authors":["D. M. Moore","Soó"]}}}]}}]},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":19},{"wordType":"authorWord","start":21,"end":23},{"wordType":"authorWord","start":25,"end":32},{"wordType":"hybridChar","start":33,"end":34},{"wordType":"genus","start":35,"end":47},{"wordType":"specificEpithet","start":48,"end":55},{"wordType":"authorWord","start":57,"end":62},{"wordType":"authorWordFilius","start":63,"end":65},{"wordType":"authorWord","start":67,"end":69},{"wordType":"authorWord","start":69,"end":71},{"wordType":"authorWord","start":72,"end":76},{"wordType":"authorWord","start":79,"end":90},{"wordType":"rank","start":91,"end":95},{"wordType":"infraspecificEpithet","start":96,"end":108},{"wordType":"authorWord","start":110,"end":115},{"wordType":"authorWord","start":117,"end":119},{"wordType":"authorWord","start":119,"end":121},{"wordType":"authorWord","start":122,"end":127},{"wordType":"authorWord","start":130,"end":133}],"id":"76fc857a-442a-590e-98c6-174aeb199e68","parserVersion":"test_version"}
```

Name: Salix aurita L. × S. caprea L.

Canonical: Salix aurita × Salix caprea

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Salix aurita L. × S. caprea L.","normalized":"Salix aurita L. × Salix caprea L.","canonical":{"stemmed":"Salix aurit × Salix capre","simple":"Salix aurita × Salix caprea","full":"Salix aurita × Salix caprea"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Salix","species":"aurita","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},{"species":{"genus":"Salix","species":"caprea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}}]},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":12},{"wordType":"authorWord","start":13,"end":15},{"wordType":"hybridChar","start":16,"end":17},{"wordType":"genus","start":18,"end":20},{"wordType":"specificEpithet","start":21,"end":27},{"wordType":"authorWord","start":28,"end":30}],"id":"a8de3172-b5e8-55c0-b495-b13b7af462d4","parserVersion":"test_version"}
```

Name: Asplenium rhizophyllum X A. ruta-muraria E.L. Braun 1939

Canonical: Asplenium rhizophyllum × Asplenium ruta-muraria

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Asplenium rhizophyllum X A. ruta-muraria E.L. Braun 1939","normalized":"Asplenium rhizophyllum × Asplenium ruta-muraria E. L. Braun 1939","canonical":{"stemmed":"Asplenium rhizophyll × Aspleni ruta-murar","simple":"Asplenium rhizophyllum × Asplenium ruta-muraria","full":"Asplenium rhizophyllum × Asplenium ruta-muraria"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Asplenium","species":"rhizophyllum"}},{"species":{"genus":"Asplenium","species":"ruta-muraria","authorship":{"verbatim":"E.L. Braun 1939","normalized":"E. L. Braun 1939","year":"1939","authors":["E. L. Braun"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1939"}}}}}]},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"hybridChar","start":23,"end":24},{"wordType":"genus","start":25,"end":27},{"wordType":"specificEpithet","start":28,"end":40},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":51},{"wordType":"year","start":52,"end":56}],"id":"1fa2c609-ce9b-5eea-a1b2-187d36b695cb","parserVersion":"test_version"}
```

Name: Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939

Canonical: Asplenium rhizophyllum × Asplenium ruta-muraria

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Incomplete hybrid formula"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939","normalized":"Asplenium rhizophyllum DC. × Asplenium ruta-muraria E. L. Braun 1939","canonical":{"stemmed":"Asplenium rhizophyll × Aspleni ruta-murar","simple":"Asplenium rhizophyllum × Asplenium ruta-muraria","full":"Asplenium rhizophyllum × Asplenium ruta-muraria"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Asplenium","species":"rhizophyllum","authorship":{"verbatim":"DC.","normalized":"DC.","authors":["DC."],"originalAuth":{"authors":["DC."]}}}},{"species":{"genus":"Asplenium","species":"ruta-muraria","authorship":{"verbatim":"E.L. Braun 1939","normalized":"E. L. Braun 1939","year":"1939","authors":["E. L. Braun"],"originalAuth":{"authors":["E. L. Braun"],"year":{"year":"1939"}}}}}]},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"authorWord","start":23,"end":26},{"wordType":"hybridChar","start":27,"end":28},{"wordType":"specificEpithet","start":29,"end":41},{"wordType":"authorWord","start":42,"end":44},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":47,"end":52},{"wordType":"year","start":53,"end":57}],"id":"dcb8fb0f-8207-5c67-b02b-81c8e03001b2","parserVersion":"test_version"}
```

<!--
TODO Mentha aquatica L. × M. arvensis L. × M. spicata L.|''
TODO Polypodium vulgare subsp. prionodes (Asch.) Rothm. × subsp. vulgare|''
-->

Name: Tilletia caries (Bjerk.) Tul. × T. foetida (Wallr.) Liro.

Canonical: Tilletia caries × Tilletia foetida

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Tilletia caries (Bjerk.) Tul. × T. foetida (Wallr.) Liro.","normalized":"Tilletia caries (Bjerk.) Tul. × Tilletia foetida (Wallr.) Liro.","canonical":{"stemmed":"Tilletia cari × Tillet foetid","simple":"Tilletia caries × Tilletia foetida","full":"Tilletia caries × Tilletia foetida"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Tilletia","species":"caries","authorship":{"verbatim":"(Bjerk.) Tul.","normalized":"(Bjerk.) Tul.","authors":["Bjerk.","Tul."],"originalAuth":{"authors":["Bjerk."]},"combinationAuth":{"authors":["Tul."]}}}},{"species":{"genus":"Tilletia","species":"foetida","authorship":{"verbatim":"(Wallr.) Liro.","normalized":"(Wallr.) Liro.","authors":["Wallr.","Liro."],"originalAuth":{"authors":["Wallr."]},"combinationAuth":{"authors":["Liro."]}}}}]},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"authorWord","start":17,"end":23},{"wordType":"authorWord","start":25,"end":29},{"wordType":"hybridChar","start":30,"end":31},{"wordType":"genus","start":32,"end":34},{"wordType":"specificEpithet","start":35,"end":42},{"wordType":"authorWord","start":44,"end":50},{"wordType":"authorWord","start":52,"end":57}],"id":"65d2072c-86e1-5205-a188-0d554dccd0e7","parserVersion":"test_version"}
```

Name: Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × B. oleracea L. subsp. capitata (L.) var. costata DC.

Canonical: Brassica oleracea subsp. capitata convar. fruticosa × Brassica oleracea subsp. capitata var. costata

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × B. oleracea L. subsp. capitata (L.) var. costata DC.","normalized":"Brassica oleracea L. subsp. capitata (L.) DC. convar. fruticosa (Metzg.) Alef. × Brassica oleracea L. subsp. capitata (L.) var. costata DC.","canonical":{"stemmed":"Brassica olerace capitat fruticos × Brassic olerace capitat costat","simple":"Brassica oleracea capitata fruticosa × Brassica oleracea capitata costata","full":"Brassica oleracea subsp. capitata convar. fruticosa × Brassica oleracea subsp. capitata var. costata"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"infraSpecies":{"genus":"Brassica","species":"oleracea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraSpecies":[{"value":"capitata","rank":"subsp.","authorship":{"verbatim":"(L.) DC.","normalized":"(L.) DC.","authors":["L.","DC."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["DC."]}}},{"value":"fruticosa","rank":"convar.","authorship":{"verbatim":"(Metzg.) Alef.","normalized":"(Metzg.) Alef.","authors":["Metzg.","Alef."],"originalAuth":{"authors":["Metzg."]},"combinationAuth":{"authors":["Alef."]}}}]}},{"infraSpecies":{"genus":"Brassica","species":"oleracea","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"infraSpecies":[{"value":"capitata","rank":"subsp.","authorship":{"verbatim":"(L.)","normalized":"(L.)","authors":["L."],"originalAuth":{"authors":["L."]}}},{"value":"costata","rank":"var.","authorship":{"verbatim":"DC.","normalized":"DC.","authors":["DC."],"originalAuth":{"authors":["DC."]}}}]}}]},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":20},{"wordType":"rank","start":21,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":36},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":42,"end":45},{"wordType":"rank","start":46,"end":53},{"wordType":"infraspecificEpithet","start":54,"end":63},{"wordType":"authorWord","start":65,"end":71},{"wordType":"authorWord","start":73,"end":78},{"wordType":"hybridChar","start":79,"end":80},{"wordType":"genus","start":81,"end":83},{"wordType":"specificEpithet","start":84,"end":92},{"wordType":"authorWord","start":93,"end":95},{"wordType":"rank","start":96,"end":102},{"wordType":"infraspecificEpithet","start":103,"end":111},{"wordType":"authorWord","start":113,"end":115},{"wordType":"rank","start":117,"end":121},{"wordType":"infraspecificEpithet","start":122,"end":129},{"wordType":"authorWord","start":130,"end":133}],"id":"2e0f4d35-ccd2-5d4a-ab42-956932ea8fb0","parserVersion":"test_version"}
```

Name: Ambystoma laterale × A. texanum × A. tigrinum

Canonical: Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Abbreviated uninomial word"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Ambystoma laterale × A. texanum × A. tigrinum","normalized":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum","canonical":{"stemmed":"Ambystoma lateral × Ambystom texan × Ambystom tigrin","simple":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum","full":"Ambystoma laterale × Ambystoma texanum × Ambystoma tigrinum"},"cardinality":0,"hybrid":"HYBRID_FORMULA","details":{"hybridFormula":[{"species":{"genus":"Ambystoma","species":"laterale"}},{"species":{"genus":"Ambystoma","species":"texanum"}},{"species":{"genus":"Ambystoma","species":"tigrinum"}}]},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":18},{"wordType":"hybridChar","start":19,"end":20},{"wordType":"genus","start":21,"end":23},{"wordType":"specificEpithet","start":24,"end":31},{"wordType":"hybridChar","start":32,"end":33},{"wordType":"genus","start":34,"end":36},{"wordType":"specificEpithet","start":37,"end":45}],"id":"ae91df82-158b-5307-83eb-f448044acec5","parserVersion":"test_version"}
```

<!-- NOTE: handle 'X' in author name correctly -->
Name: Pseudocercospora broussonetiae (Chupp & Linder) X.J. Liu & Y.L. Guo 1989

Canonical: Pseudocercospora broussonetiae

Authorship: (Chupp & Linder) X. J. Liu & Y. L. Guo 1989

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora broussonetiae (Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"Pseudocercospora broussonetiae (Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","canonical":{"stemmed":"Pseudocercospora broussoneti","simple":"Pseudocercospora broussonetiae","full":"Pseudocercospora broussonetiae"},"cardinality":2,"authorship":{"verbatim":"(Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"(Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","authors":["Chupp","Linder","X. J. Liu","Y. L. Guo"],"originalAuth":{"authors":["Chupp","Linder"]},"combinationAuth":{"authors":["X. J. Liu","Y. L. Guo"],"year":{"year":"1989"}}},"details":{"species":{"genus":"Pseudocercospora","species":"broussonetiae","authorship":{"verbatim":"(Chupp \u0026 Linder) X.J. Liu \u0026 Y.L. Guo 1989","normalized":"(Chupp \u0026 Linder) X. J. Liu \u0026 Y. L. Guo 1989","authors":["Chupp","Linder","X. J. Liu","Y. L. Guo"],"originalAuth":{"authors":["Chupp","Linder"]},"combinationAuth":{"authors":["X. J. Liu","Y. L. Guo"],"year":{"year":"1989"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":30},{"wordType":"authorWord","start":32,"end":37},{"wordType":"authorWord","start":40,"end":46},{"wordType":"authorWord","start":48,"end":50},{"wordType":"authorWord","start":50,"end":52},{"wordType":"authorWord","start":53,"end":56},{"wordType":"authorWord","start":59,"end":61},{"wordType":"authorWord","start":61,"end":63},{"wordType":"authorWord","start":64,"end":67},{"wordType":"year","start":68,"end":72}],"id":"64f92545-9139-5e53-9ba5-c5c9edb51be5","parserVersion":"test_version"}
```

### Genus with hyphen (allowed by ICN)

Name: Saxo-Fridericia R. H. Schomb.

Canonical: Saxo-fridericia

Authorship: R. H. Schomb.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Saxo-Fridericia R. H. Schomb.","normalized":"Saxo-fridericia R. H. Schomb.","canonical":{"stemmed":"Saxo-fridericia","simple":"Saxo-fridericia","full":"Saxo-fridericia"},"cardinality":1,"authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}},"details":{"uninomial":{"uninomial":"Saxo-fridericia","authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":19,"end":21},{"wordType":"authorWord","start":22,"end":29}],"id":"f11d6164-5f08-5bb3-8432-5f07d1ee3bd4","parserVersion":"test_version"}
```

Name: Saxo-fridericia R. H. Schomb.

Canonical: Saxo-fridericia

Authorship: R. H. Schomb.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Saxo-fridericia R. H. Schomb.","normalized":"Saxo-fridericia R. H. Schomb.","canonical":{"stemmed":"Saxo-fridericia","simple":"Saxo-fridericia","full":"Saxo-fridericia"},"cardinality":1,"authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}},"details":{"uninomial":{"uninomial":"Saxo-fridericia","authorship":{"verbatim":"R. H. Schomb.","normalized":"R. H. Schomb.","authors":["R. H. Schomb."],"originalAuth":{"authors":["R. H. Schomb."]}}}},"pos":[{"wordType":"uninomial","start":0,"end":15},{"wordType":"authorWord","start":16,"end":18},{"wordType":"authorWord","start":19,"end":21},{"wordType":"authorWord","start":22,"end":29}],"id":"9eac48bf-fbb1-57a3-b171-0b3bfda9757f","parserVersion":"test_version"}
```

Name: Uva-ursi cinerea (Howell) A. Heller

Canonical: Uva-ursi cinerea

Authorship: (Howell) A. Heller

```json
{"parsed":true,"parseQuality":1,"verbatim":"Uva-ursi cinerea (Howell) A. Heller","normalized":"Uva-ursi cinerea (Howell) A. Heller","canonical":{"stemmed":"Uva-ursi cinere","simple":"Uva-ursi cinerea","full":"Uva-ursi cinerea"},"cardinality":2,"authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}},"details":{"species":{"genus":"Uva-ursi","species":"cinerea","authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":18,"end":24},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":29,"end":35}],"id":"1f0bc087-ceec-5326-9fa1-2ce3b369bd7d","parserVersion":"test_version"}
```

Name: Uva-Ursi cinerea (Howell) A. Heller

Canonical: Uva-ursi cinerea

Authorship: (Howell) A. Heller

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Uva-Ursi cinerea (Howell) A. Heller","normalized":"Uva-ursi cinerea (Howell) A. Heller","canonical":{"stemmed":"Uva-ursi cinere","simple":"Uva-ursi cinerea","full":"Uva-ursi cinerea"},"cardinality":2,"authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}},"details":{"species":{"genus":"Uva-ursi","species":"cinerea","authorship":{"verbatim":"(Howell) A. Heller","normalized":"(Howell) A. Heller","authors":["Howell","A. Heller"],"originalAuth":{"authors":["Howell"]},"combinationAuth":{"authors":["A. Heller"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":18,"end":24},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":29,"end":35}],"id":"c89977a6-b948-5d3f-b4f2-d25b4d0b6ea0","parserVersion":"test_version"}
```

### Misspeled name

Name: Ambrysus-Stål, 1862

Canonical: Ambrysus-stål

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Non-standard characters in canonical"},{"quality":2,"warning":"Apparent genus with capital character after hyphen"}],"verbatim":"Ambrysus-Stål, 1862","normalized":"Ambrysus-stål","canonical":{"stemmed":"Ambrysus-stål","simple":"Ambrysus-stål","full":"Ambrysus-stål"},"cardinality":1,"tail":", 1862","details":{"uninomial":{"uninomial":"Ambrysus-stål"}},"pos":[{"wordType":"uninomial","start":0,"end":13}],"id":"ab9e69c4-9418-5f86-ad51-3bfc87f76016","parserVersion":"test_version"}
```

### A 'basionym' author in parenthesis (basionym is an ICN term)

Name: Zophosis persis (Chatanay, 1914)

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zophosis persis (Chatanay, 1914)","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay, 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay, 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"authorWord","start":17,"end":25},{"wordType":"year","start":27,"end":31}],"id":"b70a2324-4f36-5fef-80b3-5f6ab9c7788d","parserVersion":"test_version"}
```

Name: Zophosis persis (Chatanay 1914)

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zophosis persis (Chatanay 1914)","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay 1914)","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"authorWord","start":17,"end":25},{"wordType":"year","start":26,"end":30}],"id":"c6c42947-16b5-5c1c-a889-51392d82a03b","parserVersion":"test_version"}
```

Name: Zophosis persis (Chatanay), 1914

Canonical: Zophosis persis

Authorship: (Chatanay 1914)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Misplaced basionym year"}],"verbatim":"Zophosis persis (Chatanay), 1914","normalized":"Zophosis persis (Chatanay 1914)","canonical":{"stemmed":"Zophosis pers","simple":"Zophosis persis","full":"Zophosis persis"},"cardinality":2,"authorship":{"verbatim":"(Chatanay), 1914","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}},"details":{"species":{"genus":"Zophosis","species":"persis","authorship":{"verbatim":"(Chatanay), 1914","normalized":"(Chatanay 1914)","year":"1914","authors":["Chatanay"],"originalAuth":{"authors":["Chatanay"],"year":{"year":"1914"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15},{"wordType":"authorWord","start":17,"end":25},{"wordType":"year","start":28,"end":32}],"id":"3f9b079c-510a-5c0c-9df6-f1660e1b005f","parserVersion":"test_version"}
```

Name: Zophosis quadrilineata (Oliv. )

Canonical: Zophosis quadrilineata

Authorship: (Oliv.)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zophosis quadrilineata (Oliv. )","normalized":"Zophosis quadrilineata (Oliv.)","canonical":{"stemmed":"Zophosis quadrilineat","simple":"Zophosis quadrilineata","full":"Zophosis quadrilineata"},"cardinality":2,"authorship":{"verbatim":"(Oliv. )","normalized":"(Oliv.)","authors":["Oliv."],"originalAuth":{"authors":["Oliv."]}},"details":{"species":{"genus":"Zophosis","species":"quadrilineata","authorship":{"verbatim":"(Oliv. )","normalized":"(Oliv.)","authors":["Oliv."],"originalAuth":{"authors":["Oliv."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":22},{"wordType":"authorWord","start":24,"end":29}],"id":"4d327524-3514-5faf-85fa-e461cbf6c99e","parserVersion":"test_version"}
```

Name: Zophosis quadrilineata (Olivier 1795)

Canonical: Zophosis quadrilineata

Authorship: (Olivier 1795)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Zophosis quadrilineata (Olivier 1795)","normalized":"Zophosis quadrilineata (Olivier 1795)","canonical":{"stemmed":"Zophosis quadrilineat","simple":"Zophosis quadrilineata","full":"Zophosis quadrilineata"},"cardinality":2,"authorship":{"verbatim":"(Olivier 1795)","normalized":"(Olivier 1795)","year":"1795","authors":["Olivier"],"originalAuth":{"authors":["Olivier"],"year":{"year":"1795"}}},"details":{"species":{"genus":"Zophosis","species":"quadrilineata","authorship":{"verbatim":"(Olivier 1795)","normalized":"(Olivier 1795)","year":"1795","authors":["Olivier"],"originalAuth":{"authors":["Olivier"],"year":{"year":"1795"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":22},{"wordType":"authorWord","start":24,"end":31},{"wordType":"year","start":32,"end":36}],"id":"837cbd42-87a0-573f-9dbf-d089503028ad","parserVersion":"test_version"}
```

### Infrageneric epithets (ICZN)

Name: Hegeter (Hegeter) tenuipunctatus Brullé, 1838

Canonical: Hegeter tenuipunctatus

Authorship: Brullé 1838

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hegeter (Hegeter) tenuipunctatus Brullé, 1838","normalized":"Hegeter (Hegeter) tenuipunctatus Brullé 1838","canonical":{"stemmed":"Hegeter tenuipunctat","simple":"Hegeter tenuipunctatus","full":"Hegeter tenuipunctatus"},"cardinality":2,"authorship":{"verbatim":"Brullé, 1838","normalized":"Brullé 1838","year":"1838","authors":["Brullé"],"originalAuth":{"authors":["Brullé"],"year":{"year":"1838"}}},"details":{"species":{"genus":"Hegeter","subGenus":"Hegeter","species":"tenuipunctatus","authorship":{"verbatim":"Brullé, 1838","normalized":"Brullé 1838","year":"1838","authors":["Brullé"],"originalAuth":{"authors":["Brullé"],"year":{"year":"1838"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"infragenericEpithet","start":9,"end":16},{"wordType":"specificEpithet","start":18,"end":32},{"wordType":"authorWord","start":33,"end":39},{"wordType":"year","start":41,"end":45}],"id":"a5d28cfb-77a8-509c-a7c6-aa598a7cd3d9","parserVersion":"test_version"}
```

Name: Hegeter (Hegeter) intercedens Lindberg H 1950

Canonical: Hegeter intercedens

Authorship: Lindberg H 1950

```json
{"parsed":true,"parseQuality":1,"verbatim":"Hegeter (Hegeter) intercedens Lindberg H 1950","normalized":"Hegeter (Hegeter) intercedens Lindberg H 1950","canonical":{"stemmed":"Hegeter intercedens","simple":"Hegeter intercedens","full":"Hegeter intercedens"},"cardinality":2,"authorship":{"verbatim":"Lindberg H 1950","normalized":"Lindberg H 1950","year":"1950","authors":["Lindberg H"],"originalAuth":{"authors":["Lindberg H"],"year":{"year":"1950"}}},"details":{"species":{"genus":"Hegeter","subGenus":"Hegeter","species":"intercedens","authorship":{"verbatim":"Lindberg H 1950","normalized":"Lindberg H 1950","year":"1950","authors":["Lindberg H"],"originalAuth":{"authors":["Lindberg H"],"year":{"year":"1950"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"infragenericEpithet","start":9,"end":16},{"wordType":"specificEpithet","start":18,"end":29},{"wordType":"authorWord","start":30,"end":38},{"wordType":"authorWord","start":39,"end":40},{"wordType":"year","start":41,"end":45}],"id":"2486503e-b9fb-547f-a310-944a50d1bce8","parserVersion":"test_version"}
```

<!--
Brachytrypus (B.) grandidieri
-->

Name: Cyprideis (Cyprideis) thessalonike amasyaensis

Canonical: Cyprideis thessalonike amasyaensis

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cyprideis (Cyprideis) thessalonike amasyaensis","normalized":"Cyprideis (Cyprideis) thessalonike amasyaensis","canonical":{"stemmed":"Cyprideis thessalonik amasyaens","simple":"Cyprideis thessalonike amasyaensis","full":"Cyprideis thessalonike amasyaensis"},"cardinality":3,"details":{"infraSpecies":{"genus":"Cyprideis","subGenus":"Cyprideis","species":"thessalonike","infraSpecies":[{"value":"amasyaensis"}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"infragenericEpithet","start":11,"end":20},{"wordType":"specificEpithet","start":22,"end":34},{"wordType":"infraspecificEpithet","start":35,"end":46}],"id":"19945ce1-52ee-5416-af46-0d6f0803b44e","parserVersion":"test_version"}
```

Name: Acanthoderes (acanthoderes) satanas Aurivillius, 1923

Canonical: Acanthoderes satanas

Authorship: Aurivillius 1923

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ambiguity: subgenus or superspecies found"}],"verbatim":"Acanthoderes (acanthoderes) satanas Aurivillius, 1923","normalized":"Acanthoderes satanas Aurivillius 1923","canonical":{"stemmed":"Acanthoderes satan","simple":"Acanthoderes satanas","full":"Acanthoderes satanas"},"cardinality":2,"authorship":{"verbatim":"Aurivillius, 1923","normalized":"Aurivillius 1923","year":"1923","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1923"}}},"details":{"species":{"genus":"Acanthoderes","species":"satanas","authorship":{"verbatim":"Aurivillius, 1923","normalized":"Aurivillius 1923","year":"1923","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"],"year":{"year":"1923"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":28,"end":35},{"wordType":"authorWord","start":36,"end":47},{"wordType":"year","start":49,"end":53}],"id":"f1082b19-d13f-54a2-95a9-6e342f2a9e6b","parserVersion":"test_version"}
```

<!-- A fake name to illustrate botaincal author instead of subgenus -->
Name: Acanthoderes (Abramov) satanas Aurivillius

Canonical: Acanthoderes satanas

Authorship: Aurivillius

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Possible ICN author instead of subgenus"}],"verbatim":"Acanthoderes (Abramov) satanas Aurivillius","normalized":"Acanthoderes satanas Aurivillius","canonical":{"stemmed":"Acanthoderes satan","simple":"Acanthoderes satanas","full":"Acanthoderes satanas"},"cardinality":2,"authorship":{"verbatim":"Aurivillius","normalized":"Aurivillius","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"]}},"details":{"species":{"genus":"Acanthoderes","species":"satanas","authorship":{"verbatim":"Aurivillius","normalized":"Aurivillius","authors":["Aurivillius"],"originalAuth":{"authors":["Aurivillius"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":23,"end":30},{"wordType":"authorWord","start":31,"end":42}],"id":"8eb2a9be-eb11-537e-8488-eacdb6e2b9e7","parserVersion":"test_version"}
```

### Names with multiple dashes in specific epithet

There are less than 100 of names like this, and only one in CoL with 3 dashes

Name: Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.

Canonical: Athyrium boreo-occidentali-indobharaticola-birianum

Authorship: Fraser-Jenk.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.","normalized":"Athyrium boreo-occidentali-indobharaticola-birianum Fraser-Jenk.","canonical":{"stemmed":"Athyrium boreo-occidentali-indobharaticola-birian","simple":"Athyrium boreo-occidentali-indobharaticola-birianum","full":"Athyrium boreo-occidentali-indobharaticola-birianum"},"cardinality":2,"authorship":{"verbatim":"Fraser-Jenk.","normalized":"Fraser-Jenk.","authors":["Fraser-Jenk."],"originalAuth":{"authors":["Fraser-Jenk."]}},"details":{"species":{"genus":"Athyrium","species":"boreo-occidentali-indobharaticola-birianum","authorship":{"verbatim":"Fraser-Jenk.","normalized":"Fraser-Jenk.","authors":["Fraser-Jenk."],"originalAuth":{"authors":["Fraser-Jenk."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":51},{"wordType":"authorWord","start":52,"end":64}],"id":"6b979652-191f-5d93-ae23-614768ee0be4","parserVersion":"test_version"}
```

Name: Puccinia band-i-amirii Durrieu, 1975

Canonical: Puccinia band-i-amirii

Authorship: Durrieu 1975

```json
{"parsed":true,"parseQuality":1,"verbatim":"Puccinia band-i-amirii Durrieu, 1975","normalized":"Puccinia band-i-amirii Durrieu 1975","canonical":{"stemmed":"Puccinia band-i-amiri","simple":"Puccinia band-i-amirii","full":"Puccinia band-i-amirii"},"cardinality":2,"authorship":{"verbatim":"Durrieu, 1975","normalized":"Durrieu 1975","year":"1975","authors":["Durrieu"],"originalAuth":{"authors":["Durrieu"],"year":{"year":"1975"}}},"details":{"species":{"genus":"Puccinia","species":"band-i-amirii","authorship":{"verbatim":"Durrieu, 1975","normalized":"Durrieu 1975","year":"1975","authors":["Durrieu"],"originalAuth":{"authors":["Durrieu"],"year":{"year":"1975"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":22},{"wordType":"authorWord","start":23,"end":30},{"wordType":"year","start":32,"end":36}],"id":"9733e3df-0b03-5e1e-93f9-5931a4e85f12","parserVersion":"test_version"}
```

### Genus with question mark

Name: Ferganoconcha? oblonga

Canonical: Ferganoconcha oblonga

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Uninomial word with question mark"}],"verbatim":"Ferganoconcha? oblonga","normalized":"Ferganoconcha oblonga","canonical":{"stemmed":"Ferganoconcha oblong","simple":"Ferganoconcha oblonga","full":"Ferganoconcha oblonga"},"cardinality":2,"details":{"species":{"genus":"Ferganoconcha","species":"oblonga"}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22}],"id":"487912fd-85c3-556a-a1b1-8fe802e9ccb1","parserVersion":"test_version"}
```

### Epithets starting with authors' prefixes (de, di, la, von etc.)

<-- There is a danger that such epithets will be interpreted as authors-->

Name: Aspicilia desertorum desertorum

Canonical: Aspicilia desertorum desertorum

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aspicilia desertorum desertorum","normalized":"Aspicilia desertorum desertorum","canonical":{"stemmed":"Aspicilia desertor desertor","simple":"Aspicilia desertorum desertorum","full":"Aspicilia desertorum desertorum"},"cardinality":3,"details":{"infraSpecies":{"genus":"Aspicilia","species":"desertorum","infraSpecies":[{"value":"desertorum"}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31}],"id":"06de3555-3226-5e05-930e-6706044c1f7a","parserVersion":"test_version"}
```

Name: Theope thestias discus

Canonical: Theope thestias discus

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Theope thestias discus","normalized":"Theope thestias discus","canonical":{"stemmed":"Theope thesti disc","simple":"Theope thestias discus","full":"Theope thestias discus"},"cardinality":3,"details":{"infraSpecies":{"genus":"Theope","species":"thestias","infraSpecies":[{"value":"discus"}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"infraspecificEpithet","start":16,"end":22}],"id":"a254509a-11e4-52f3-bd57-2271d9e1d99b","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus (Dejean, 1831)

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ocydromus dalmatinus dalmatinus (Dejean, 1831)","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"(Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraSpecies":{"genus":"Ocydromus","species":"dalmatinus","infraSpecies":[{"value":"dalmatinus","authorship":{"verbatim":"(Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31},{"wordType":"authorWord","start":33,"end":39},{"wordType":"year","start":41,"end":45}],"id":"5701cc12-ec23-5015-b426-3d065c94ea0a","parserVersion":"test_version"}
```

Name: Rhipidia gracilirama lassula

Canonical: Rhipidia gracilirama lassula

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Rhipidia gracilirama lassula","normalized":"Rhipidia gracilirama lassula","canonical":{"stemmed":"Rhipidia graciliram lassul","simple":"Rhipidia gracilirama lassula","full":"Rhipidia gracilirama lassula"},"cardinality":3,"details":{"infraSpecies":{"genus":"Rhipidia","species":"gracilirama","infraSpecies":[{"value":"lassula"}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":28}],"id":"0b40c395-7466-5879-9b16-9a31d38d21a0","parserVersion":"test_version"}
```

### Authorship missing one parenthesis

Name: Ocydromus dalmatinus dalmatinus Dejean, 1831)

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus Dejean, 1831)","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraSpecies":{"genus":"Ocydromus","species":"dalmatinus","infraSpecies":[{"value":"dalmatinus","authorship":{"verbatim":"Dejean, 1831)","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31},{"wordType":"authorWord","start":32,"end":38},{"wordType":"year","start":40,"end":44}],"id":"5de70fe3-959a-5555-afdb-3ab85b91f1d7","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus Dejean, 1831 )

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus Dejean, 1831 )","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831)","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"Dejean, 1831 )","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}},"details":{"infraSpecies":{"genus":"Ocydromus","species":"dalmatinus","infraSpecies":[{"value":"dalmatinus","authorship":{"verbatim":"Dejean, 1831 )","normalized":"(Dejean 1831)","year":"1831","authors":["Dejean"],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31},{"wordType":"authorWord","start":32,"end":38},{"wordType":"year","start":40,"end":44}],"id":"88dcc885-3360-5234-9620-371c2ebb636c","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus ( Dejean, 1831 Mill.

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831) Mill.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus ( Dejean, 1831 Mill.","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831) Mill.","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"( Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}},"details":{"infraSpecies":{"genus":"Ocydromus","species":"dalmatinus","infraSpecies":[{"value":"dalmatinus","authorship":{"verbatim":"( Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31},{"wordType":"authorWord","start":34,"end":40},{"wordType":"year","start":42,"end":46},{"wordType":"authorWord","start":47,"end":52}],"id":"0e8758a1-2567-543b-bafd-c8f9c81e2f08","parserVersion":"test_version"}
```

Name: Ocydromus dalmatinus dalmatinus (Dejean, 1831 Mill.

Canonical: Ocydromus dalmatinus dalmatinus

Authorship: (Dejean 1831) Mill.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship is missing one parenthesis"}],"verbatim":"Ocydromus dalmatinus dalmatinus (Dejean, 1831 Mill.","normalized":"Ocydromus dalmatinus dalmatinus (Dejean 1831) Mill.","canonical":{"stemmed":"Ocydromus dalmatin dalmatin","simple":"Ocydromus dalmatinus dalmatinus","full":"Ocydromus dalmatinus dalmatinus"},"cardinality":3,"authorship":{"verbatim":"(Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}},"details":{"infraSpecies":{"genus":"Ocydromus","species":"dalmatinus","infraSpecies":[{"value":"dalmatinus","authorship":{"verbatim":"(Dejean, 1831 Mill.","normalized":"(Dejean 1831) Mill.","year":"1831","authors":["Dejean","Mill."],"originalAuth":{"authors":["Dejean"],"year":{"year":"1831"}},"combinationAuth":{"authors":["Mill."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":31},{"wordType":"authorWord","start":33,"end":39},{"wordType":"year","start":41,"end":45},{"wordType":"authorWord","start":46,"end":51}],"id":"b3c856b3-16a7-5dfc-abfd-3bba539b634f","parserVersion":"test_version"}
```

### Unknown authorship

Name: Saccharomyces drosophilae anon.

Canonical: Saccharomyces drosophilae

Authorship: anon.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Saccharomyces drosophilae anon.","normalized":"Saccharomyces drosophilae anon.","canonical":{"stemmed":"Saccharomyces drosophil","simple":"Saccharomyces drosophilae","full":"Saccharomyces drosophilae"},"cardinality":2,"authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Saccharomyces","species":"drosophilae","authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":25},{"wordType":"authorWord","start":26,"end":31}],"id":"45e537d2-6833-5429-a58c-178fe37fc3f5","parserVersion":"test_version"}
```

Name: Physalospora rubiginosa (Fr.) anon.

Canonical: Physalospora rubiginosa

Authorship: (Fr.) anon.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Physalospora rubiginosa (Fr.) anon.","normalized":"Physalospora rubiginosa (Fr.) anon.","canonical":{"stemmed":"Physalospora rubiginos","simple":"Physalospora rubiginosa","full":"Physalospora rubiginosa"},"cardinality":2,"authorship":{"verbatim":"(Fr.) anon.","normalized":"(Fr.) anon.","authors":["Fr.","anon."],"originalAuth":{"authors":["Fr."]},"combinationAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Physalospora","species":"rubiginosa","authorship":{"verbatim":"(Fr.) anon.","normalized":"(Fr.) anon.","authors":["Fr.","anon."],"originalAuth":{"authors":["Fr."]},"combinationAuth":{"authors":["anon."]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":28},{"wordType":"authorWord","start":30,"end":35}],"id":"85151e19-ab25-5ba5-8a19-47a5859c41bb","parserVersion":"test_version"}
```

Name: Tragacantha leporina (?) Kuntze

Canonical: Tragacantha leporina

Authorship: (anon.) Kuntze

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Author as a question mark"},{"quality":3,"warning":"Author is too short"},{"quality":2,"warning":"Author is unknown"}],"verbatim":"Tragacantha leporina (?) Kuntze","normalized":"Tragacantha leporina (anon.) Kuntze","canonical":{"stemmed":"Tragacantha leporin","simple":"Tragacantha leporina","full":"Tragacantha leporina"},"cardinality":2,"authorship":{"verbatim":"(?) Kuntze","normalized":"(anon.) Kuntze","authors":["anon.","Kuntze"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Kuntze"]}},"details":{"species":{"genus":"Tragacantha","species":"leporina","authorship":{"verbatim":"(?) Kuntze","normalized":"(anon.) Kuntze","authors":["anon.","Kuntze"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Kuntze"]}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":20},{"wordType":"authorWord","start":22,"end":23},{"wordType":"authorWord","start":25,"end":31}],"id":"af91bdc5-b6d3-5841-9a85-174c0afe6c1b","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (auct.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship: (anon.) Baker

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Lachenalia tricolor var. nelsonii (auct.) Baker","normalized":"Lachenalia tricolor var. nelsonii (anon.) Baker","canonical":{"stemmed":"Lachenalia tricolor nelsoni","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"authorship":{"verbatim":"(auct.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}},"details":{"infraSpecies":{"genus":"Lachenalia","species":"tricolor","infraSpecies":[{"value":"nelsonii","rank":"var.","authorship":{"verbatim":"(auct.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33},{"wordType":"authorWord","start":35,"end":40},{"wordType":"authorWord","start":42,"end":47}],"id":"f8d5d993-3d39-550f-bb7b-68f5b6e906df","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (anon.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship: (anon.) Baker

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Lachenalia tricolor var. nelsonii (anon.) Baker","normalized":"Lachenalia tricolor var. nelsonii (anon.) Baker","canonical":{"stemmed":"Lachenalia tricolor nelsoni","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"authorship":{"verbatim":"(anon.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}},"details":{"infraSpecies":{"genus":"Lachenalia","species":"tricolor","infraSpecies":[{"value":"nelsonii","rank":"var.","authorship":{"verbatim":"(anon.) Baker","normalized":"(anon.) Baker","authors":["anon.","Baker"],"originalAuth":{"authors":["anon."]},"combinationAuth":{"authors":["Baker"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33},{"wordType":"authorWord","start":35,"end":40},{"wordType":"authorWord","start":42,"end":47}],"id":"4cc8e603-13fb-551f-a637-04378f3321c2","parserVersion":"test_version"}
```

Name: Puya acris anon.

Canonical: Puya acris

Authorship: anon.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author is unknown"}],"verbatim":"Puya acris anon.","normalized":"Puya acris anon.","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"details":{"species":{"genus":"Puya","species":"acris","authorship":{"verbatim":"anon.","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}}}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":10},{"wordType":"authorWord","start":11,"end":16}],"id":"2b5243d3-e8a7-5e6c-a2c1-beb2ee5c3020","parserVersion":"test_version"}
```

### Treating apud (with)

Name: Pseudocercospora dendrobii Goh apud W.H. Hsieh 1990

Canonical: Pseudocercospora dendrobii

Authorship: Goh apud W. H. Hsieh 1990

```json
{"parsed":true,"parseQuality":1,"verbatim":"Pseudocercospora dendrobii Goh apud W.H. Hsieh 1990","normalized":"Pseudocercospora dendrobii Goh apud W. H. Hsieh 1990","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"authorship":{"verbatim":"Goh apud W.H. Hsieh 1990","normalized":"Goh apud W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}},"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii","authorship":{"verbatim":"Goh apud W.H. Hsieh 1990","normalized":"Goh apud W. H. Hsieh 1990","year":"1990","authors":["Goh","W. H. Hsieh"],"originalAuth":{"authors":["Goh","W. H. Hsieh"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":17,"end":26},{"wordType":"authorWord","start":27,"end":30},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":41,"end":46},{"wordType":"year","start":47,"end":51}],"id":"4dee6fc8-3be1-520c-9937-5a7342a17241","parserVersion":"test_version"}
```

### Names with ex authors (we follow ICZN convention)

Name: Arthopyrenia hyalospora (Nyl. ex Banker) R.C. Harris

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl. ex Banker) R. C. Harris

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Arthopyrenia hyalospora (Nyl. ex Banker) R.C. Harris","normalized":"Arthopyrenia hyalospora (Nyl. ex Banker) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl. ex Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl. ex Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":29},{"wordType":"authorWord","start":33,"end":39},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":52}],"id":"ab3998af-53dc-53fd-af8b-fab94dacbcbc","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Nyl. ex. Banker) R.C. Harris

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl. ex Banker) R. C. Harris

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Arthopyrenia hyalospora (Nyl. ex. Banker) R.C. Harris","normalized":"Arthopyrenia hyalospora (Nyl. ex Banker) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl. ex. Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl. ex. Banker) R.C. Harris","normalized":"(Nyl. ex Banker) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}},"combinationAuth":{"authors":["R. C. Harris"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":29},{"wordType":"authorWord","start":34,"end":40},{"wordType":"authorWord","start":42,"end":44},{"wordType":"authorWord","start":44,"end":46},{"wordType":"authorWord","start":47,"end":53}],"id":"166fd290-17f5-5b9f-8f72-86830a9bd152","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora Nyl. ex Banker

Canonical: Arthopyrenia hyalospora

Authorship: Nyl. ex Banker

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Arthopyrenia hyalospora Nyl. ex Banker","normalized":"Arthopyrenia hyalospora Nyl. ex Banker","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"Nyl. ex Banker","normalized":"Nyl. ex Banker","authors":["Nyl."],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"Nyl. ex Banker","normalized":"Nyl. ex Banker","authors":["Nyl."],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":24,"end":28},{"wordType":"authorWord","start":32,"end":38}],"id":"7744aea4-d071-593a-82bc-059788724d81","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora Nyl. ex. Banker

Canonical: Arthopyrenia hyalospora

Authorship: Nyl. ex Banker

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Arthopyrenia hyalospora Nyl. ex. Banker","normalized":"Arthopyrenia hyalospora Nyl. ex Banker","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"Nyl. ex. Banker","normalized":"Nyl. ex Banker","authors":["Nyl."],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}},"details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"Nyl. ex. Banker","normalized":"Nyl. ex Banker","authors":["Nyl."],"originalAuth":{"authors":["Nyl."],"exAuthors":{"authors":["Banker"]}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":24,"end":28},{"wordType":"authorWord","start":33,"end":39}],"id":"e9097ad7-7bb6-57a2-bad4-52822e5fd655","parserVersion":"test_version"}
```

Name: Glomopsis lonicerae Peck ex C.J. Gould 1945

Canonical: Glomopsis lonicerae

Authorship: Peck ex C. J. Gould 1945

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Glomopsis lonicerae Peck ex C.J. Gould 1945","normalized":"Glomopsis lonicerae Peck ex C. J. Gould 1945","canonical":{"stemmed":"Glomopsis lonicer","simple":"Glomopsis lonicerae","full":"Glomopsis lonicerae"},"cardinality":2,"authorship":{"verbatim":"Peck ex C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","authors":["Peck"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}},"details":{"species":{"genus":"Glomopsis","species":"lonicerae","authorship":{"verbatim":"Peck ex C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","authors":["Peck"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":19},{"wordType":"authorWord","start":20,"end":24},{"wordType":"authorWord","start":28,"end":30},{"wordType":"authorWord","start":30,"end":32},{"wordType":"authorWord","start":33,"end":38},{"wordType":"year","start":39,"end":43}],"id":"422687ca-7f4b-5720-8d99-88695f765530","parserVersion":"test_version"}
```

Name: Glomopsis lonicerae Peck ex. C.J. Gould 1945

Canonical: Glomopsis lonicerae

Authorship: Peck ex C. J. Gould 1945

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Glomopsis lonicerae Peck ex. C.J. Gould 1945","normalized":"Glomopsis lonicerae Peck ex C. J. Gould 1945","canonical":{"stemmed":"Glomopsis lonicer","simple":"Glomopsis lonicerae","full":"Glomopsis lonicerae"},"cardinality":2,"authorship":{"verbatim":"Peck ex. C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","authors":["Peck"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}},"details":{"species":{"genus":"Glomopsis","species":"lonicerae","authorship":{"verbatim":"Peck ex. C.J. Gould 1945","normalized":"Peck ex C. J. Gould 1945","authors":["Peck"],"originalAuth":{"authors":["Peck"],"exAuthors":{"authors":["C. J. Gould"],"year":{"year":"1945"}}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":19},{"wordType":"authorWord","start":20,"end":24},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":31,"end":33},{"wordType":"authorWord","start":34,"end":39},{"wordType":"year","start":40,"end":44}],"id":"a9cdd33f-990c-59b6-abc2-de9698d2f085","parserVersion":"test_version"}
```

Name: Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979

Canonical: Acanthobasidium delicatum

Authorship: (Wakef.) Oberw. ex Jülich 1979

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","normalized":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","canonical":{"stemmed":"Acanthobasidium delicat","simple":"Acanthobasidium delicatum","full":"Acanthobasidium delicatum"},"cardinality":2,"authorship":{"verbatim":"(Wakef.) Oberw. ex Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw."],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}},"details":{"species":{"genus":"Acanthobasidium","species":"delicatum","authorship":{"verbatim":"(Wakef.) Oberw. ex Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw."],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}}}},"pos":[{"wordType":"genus","start":0,"end":15},{"wordType":"specificEpithet","start":16,"end":25},{"wordType":"authorWord","start":27,"end":33},{"wordType":"authorWord","start":35,"end":41},{"wordType":"authorWord","start":45,"end":51},{"wordType":"year","start":52,"end":56}],"id":"ed0841f3-d063-5341-a1b6-feafe6ffb70d","parserVersion":"test_version"}
```

Name: Acanthobasidium delicatum (Wakef.) Oberw. ex. Jülich 1979

Canonical: Acanthobasidium delicatum

Authorship: (Wakef.) Oberw. ex Jülich 1979

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Acanthobasidium delicatum (Wakef.) Oberw. ex. Jülich 1979","normalized":"Acanthobasidium delicatum (Wakef.) Oberw. ex Jülich 1979","canonical":{"stemmed":"Acanthobasidium delicat","simple":"Acanthobasidium delicatum","full":"Acanthobasidium delicatum"},"cardinality":2,"authorship":{"verbatim":"(Wakef.) Oberw. ex. Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw."],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}},"details":{"species":{"genus":"Acanthobasidium","species":"delicatum","authorship":{"verbatim":"(Wakef.) Oberw. ex. Jülich 1979","normalized":"(Wakef.) Oberw. ex Jülich 1979","authors":["Wakef.","Oberw."],"originalAuth":{"authors":["Wakef."]},"combinationAuth":{"authors":["Oberw."],"exAuthors":{"authors":["Jülich"],"year":{"year":"1979"}}}}}},"pos":[{"wordType":"genus","start":0,"end":15},{"wordType":"specificEpithet","start":16,"end":25},{"wordType":"authorWord","start":27,"end":33},{"wordType":"authorWord","start":35,"end":41},{"wordType":"authorWord","start":46,"end":52},{"wordType":"year","start":53,"end":57}],"id":"96adf61e-3316-5a08-afac-2b7cd0430eee","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. ex Duby) Johanson ex Oudem. 1897

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","canonical":{"stemmed":"Mycosphaerella eryngi","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Johanson"],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}},"details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. ex Duby) Johanson ex Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Johanson"],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}}}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22},{"wordType":"authorWord","start":24,"end":27},{"wordType":"authorWord","start":31,"end":35},{"wordType":"authorWord","start":37,"end":45},{"wordType":"authorWord","start":49,"end":55},{"wordType":"year","start":56,"end":60}],"id":"8ca3d249-fe7d-5a10-af03-f21c413e3503","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. ex. Duby) Johanson ex. Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. ex Duby) Johanson ex Oudem. 1897

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`ex` ends with a period"},{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Mycosphaerella eryngii (Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. ex Duby) Johanson ex Oudem. 1897","canonical":{"stemmed":"Mycosphaerella eryngi","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Johanson"],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}},"details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. ex. Duby) Johanson ex. Oudem. 1897","normalized":"(Fr. ex Duby) Johanson ex Oudem. 1897","authors":["Fr.","Johanson"],"originalAuth":{"authors":["Fr."],"exAuthors":{"authors":["Duby"]}},"combinationAuth":{"authors":["Johanson"],"exAuthors":{"authors":["Oudem."],"year":{"year":"1897"}}}}}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22},{"wordType":"authorWord","start":24,"end":27},{"wordType":"authorWord","start":32,"end":36},{"wordType":"authorWord","start":38,"end":46},{"wordType":"authorWord","start":51,"end":57},{"wordType":"year","start":58,"end":62}],"id":"201b50d3-507b-56d1-99b4-50ab9120bca9","parserVersion":"test_version"}
```

Name: Mycosphaerella eryngii (Fr. Duby) ex Oudem. 1897

Canonical: Mycosphaerella eryngii

Authorship: (Fr. Duby)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Mycosphaerella eryngii (Fr. Duby) ex Oudem. 1897","normalized":"Mycosphaerella eryngii (Fr. Duby)","canonical":{"stemmed":"Mycosphaerella eryngi","simple":"Mycosphaerella eryngii","full":"Mycosphaerella eryngii"},"cardinality":2,"authorship":{"verbatim":"(Fr. Duby)","normalized":"(Fr. Duby)","authors":["Fr. Duby"],"originalAuth":{"authors":["Fr. Duby"]}},"tail":" ex Oudem. 1897","details":{"species":{"genus":"Mycosphaerella","species":"eryngii","authorship":{"verbatim":"(Fr. Duby)","normalized":"(Fr. Duby)","authors":["Fr. Duby"],"originalAuth":{"authors":["Fr. Duby"]}}}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22},{"wordType":"authorWord","start":24,"end":27},{"wordType":"authorWord","start":28,"end":32}],"id":"e5a49f2e-c7a2-5ebf-9349-8a36a410ec77","parserVersion":"test_version"}
```

### Empty spaces
Name:     Asplenium       X inexpectatum(E. L. Braun ex Friesner      )Morton

Canonical: Asplenium × inexpectatum

Authorship: (E. L. Braun ex Friesner) Morton

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Named hybrid"},{"quality":2,"warning":"Multiple adjacent space characters"}],"verbatim":"    Asplenium       X inexpectatum(E. L. Braun ex Friesner      )Morton","normalized":"Asplenium × inexpectatum (E. L. Braun ex Friesner) Morton","canonical":{"stemmed":"Asplenium inexpectat","simple":"Asplenium inexpectatum","full":"Asplenium × inexpectatum"},"cardinality":2,"authorship":{"verbatim":"(E. L. Braun ex Friesner      )Morton","normalized":"(E. L. Braun ex Friesner) Morton","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"exAuthors":{"authors":["Friesner"]}},"combinationAuth":{"authors":["Morton"]}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Asplenium","species":"inexpectatum (E. L. Braun ex Friesner) Morton","authorship":{"verbatim":"(E. L. Braun ex Friesner      )Morton","normalized":"(E. L. Braun ex Friesner) Morton","authors":["E. L. Braun","Morton"],"originalAuth":{"authors":["E. L. Braun"],"exAuthors":{"authors":["Friesner"]}},"combinationAuth":{"authors":["Morton"]}}}},"pos":[{"wordType":"genus","start":4,"end":13},{"wordType":"hybridChar","start":20,"end":21},{"wordType":"specificEpithet","start":22,"end":34},{"wordType":"authorWord","start":35,"end":37},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":41,"end":46},{"wordType":"authorWord","start":50,"end":58},{"wordType":"authorWord","start":65,"end":71}],"id":"a2c7a7ee-51c9-5f3a-8117-bffd799b39f4","parserVersion":"test_version"}
```

### Names with a dash

Name: Drosophila obscura-x Burla, 1951

Canonical: Drosophila obscura-x

Authorship: Burla 1951

```json
{"parsed":true,"parseQuality":1,"verbatim":"Drosophila obscura-x Burla, 1951","normalized":"Drosophila obscura-x Burla 1951","canonical":{"stemmed":"Drosophila obscura-x","simple":"Drosophila obscura-x","full":"Drosophila obscura-x"},"cardinality":2,"authorship":{"verbatim":"Burla, 1951","normalized":"Burla 1951","year":"1951","authors":["Burla"],"originalAuth":{"authors":["Burla"],"year":{"year":"1951"}}},"details":{"species":{"genus":"Drosophila","species":"obscura-x","authorship":{"verbatim":"Burla, 1951","normalized":"Burla 1951","year":"1951","authors":["Burla"],"originalAuth":{"authors":["Burla"],"year":{"year":"1951"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":20},{"wordType":"authorWord","start":21,"end":26},{"wordType":"year","start":28,"end":32}],"id":"778f9878-8e47-5c7a-a464-33805b6bf173","parserVersion":"test_version"}
```

Name: Sanogasta x-signata (Keyserling,1891)

Canonical: Sanogasta x-signata

Authorship: (Keyserling 1891)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Sanogasta x-signata (Keyserling,1891)","normalized":"Sanogasta x-signata (Keyserling 1891)","canonical":{"stemmed":"Sanogasta x-signat","simple":"Sanogasta x-signata","full":"Sanogasta x-signata"},"cardinality":2,"authorship":{"verbatim":"(Keyserling,1891)","normalized":"(Keyserling 1891)","year":"1891","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1891"}}},"details":{"species":{"genus":"Sanogasta","species":"x-signata","authorship":{"verbatim":"(Keyserling,1891)","normalized":"(Keyserling 1891)","year":"1891","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1891"}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":19},{"wordType":"authorWord","start":21,"end":31},{"wordType":"year","start":32,"end":36}],"id":"ffe6799d-387a-53d8-8fdd-be73cdc681b8","parserVersion":"test_version"}
```

Name: Aedes w-albus (Theobald, 1905)

Canonical: Aedes w-albus

Authorship: (Theobald 1905)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Aedes w-albus (Theobald, 1905)","normalized":"Aedes w-albus (Theobald 1905)","canonical":{"stemmed":"Aedes w-alb","simple":"Aedes w-albus","full":"Aedes w-albus"},"cardinality":2,"authorship":{"verbatim":"(Theobald, 1905)","normalized":"(Theobald 1905)","year":"1905","authors":["Theobald"],"originalAuth":{"authors":["Theobald"],"year":{"year":"1905"}}},"details":{"species":{"genus":"Aedes","species":"w-albus","authorship":{"verbatim":"(Theobald, 1905)","normalized":"(Theobald 1905)","year":"1905","authors":["Theobald"],"originalAuth":{"authors":["Theobald"],"year":{"year":"1905"}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":13},{"wordType":"authorWord","start":15,"end":23},{"wordType":"year","start":25,"end":29}],"id":"7b0dd259-10ae-5b47-95ca-2685d4c323ce","parserVersion":"test_version"}
```

Name: Abryna regis-petri Paiva, 1860

Canonical: Abryna regis-petri

Authorship: Paiva 1860

```json
{"parsed":true,"parseQuality":1,"verbatim":"Abryna regis-petri Paiva, 1860","normalized":"Abryna regis-petri Paiva 1860","canonical":{"stemmed":"Abryna regis-petr","simple":"Abryna regis-petri","full":"Abryna regis-petri"},"cardinality":2,"authorship":{"verbatim":"Paiva, 1860","normalized":"Paiva 1860","year":"1860","authors":["Paiva"],"originalAuth":{"authors":["Paiva"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Abryna","species":"regis-petri","authorship":{"verbatim":"Paiva, 1860","normalized":"Paiva 1860","year":"1860","authors":["Paiva"],"originalAuth":{"authors":["Paiva"],"year":{"year":"1860"}}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":18},{"wordType":"authorWord","start":19,"end":24},{"wordType":"year","start":26,"end":30}],"id":"27ad601d-bb92-515b-9c45-1faa55cdf7f3","parserVersion":"test_version"}
```

<!--
Abryna- regis|{"name_string_id":"9ff9c1fa-068e-5296-8c39-66e1c58f0660","parsed":false,"parser_version":"test_version","verbatim":"Abryna- regis","normalized":null,"canonical":null,"hybrid":false,"virus":false}
Abryna regis- Paiva, 1860|{"name_string_id":"473b8b63-8d5c-521f-9a68-7aecd5b9a62c","parsed":false,"parser_version":"test_version","verbatim":"Abryna regis- Paiva, 1860","normalized":null,"canonical":null,"hybrid":false,"virus":false}
-->

Name: Solms-laubachia orbiculata Y.C. Lan & T.Y. Cheo

Canonical: Solms-laubachia orbiculata

Authorship: Y. C. Lan & T. Y. Cheo

```json
{"parsed":true,"parseQuality":1,"verbatim":"Solms-laubachia orbiculata Y.C. Lan \u0026 T.Y. Cheo","normalized":"Solms-laubachia orbiculata Y. C. Lan \u0026 T. Y. Cheo","canonical":{"stemmed":"Solms-laubachia orbiculat","simple":"Solms-laubachia orbiculata","full":"Solms-laubachia orbiculata"},"cardinality":2,"authorship":{"verbatim":"Y.C. Lan \u0026 T.Y. Cheo","normalized":"Y. C. Lan \u0026 T. Y. Cheo","authors":["Y. C. Lan","T. Y. Cheo"],"originalAuth":{"authors":["Y. C. Lan","T. Y. Cheo"]}},"details":{"species":{"genus":"Solms-laubachia","species":"orbiculata","authorship":{"verbatim":"Y.C. Lan \u0026 T.Y. Cheo","normalized":"Y. C. Lan \u0026 T. Y. Cheo","authors":["Y. C. Lan","T. Y. Cheo"],"originalAuth":{"authors":["Y. C. Lan","T. Y. Cheo"]}}}},"pos":[{"wordType":"genus","start":0,"end":15},{"wordType":"specificEpithet","start":16,"end":26},{"wordType":"authorWord","start":27,"end":29},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":32,"end":35},{"wordType":"authorWord","start":38,"end":40},{"wordType":"authorWord","start":40,"end":42},{"wordType":"authorWord","start":43,"end":47}],"id":"4dce39e2-ffd7-5a1b-bd1a-2bc12049be90","parserVersion":"test_version"}
```

### Authorship with filius (son of)

Name: Oxytropis minjanensis Rech. f.

Canonical: Oxytropis minjanensis

Authorship: Rech. fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Oxytropis minjanensis Rech. f.","normalized":"Oxytropis minjanensis Rech. fil.","canonical":{"stemmed":"Oxytropis minianens","simple":"Oxytropis minjanensis","full":"Oxytropis minjanensis"},"cardinality":2,"authorship":{"verbatim":"Rech. f.","normalized":"Rech. fil.","authors":["Rech. fil."],"originalAuth":{"authors":["Rech. fil."]}},"details":{"species":{"genus":"Oxytropis","species":"minjanensis","authorship":{"verbatim":"Rech. f.","normalized":"Rech. fil.","authors":["Rech. fil."],"originalAuth":{"authors":["Rech. fil."]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":21},{"wordType":"authorWord","start":22,"end":27},{"wordType":"authorWordFilius","start":28,"end":30}],"id":"6027cbc2-fa15-510b-ab3e-e1fa44cbd551","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl f. 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl fil. 1935

```json
{"parsed":true,"parseQuality":1,"verbatim":"Platypus bicaudatulus Schedl f. 1935","normalized":"Platypus bicaudatulus Schedl fil. 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl f. 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl f. 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"authorWordFilius","start":29,"end":31},{"wordType":"year","start":32,"end":36}],"id":"05799df9-471e-5c68-92fe-4edcc0a69d29","parserVersion":"test_version"}
```

Name: Platypus bicaudatulus Schedl filius 1935

Canonical: Platypus bicaudatulus

Authorship: Schedl fil. 1935

```json
{"parsed":true,"parseQuality":1,"verbatim":"Platypus bicaudatulus Schedl filius 1935","normalized":"Platypus bicaudatulus Schedl fil. 1935","canonical":{"stemmed":"Platypus bicaudatul","simple":"Platypus bicaudatulus","full":"Platypus bicaudatulus"},"cardinality":2,"authorship":{"verbatim":"Schedl filius 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}},"details":{"species":{"genus":"Platypus","species":"bicaudatulus","authorship":{"verbatim":"Schedl filius 1935","normalized":"Schedl fil. 1935","year":"1935","authors":["Schedl fil."],"originalAuth":{"authors":["Schedl fil."],"year":{"year":"1935"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":28},{"wordType":"authorWordFilius","start":29,"end":35},{"wordType":"year","start":36,"end":40}],"id":"2b6cd51f-aa0f-58fd-88fa-2e261cedacbb","parserVersion":"test_version"}
```

Name: Fimbristylis ovata (Burm. f.) J. Kern

Canonical: Fimbristylis ovata

Authorship: (Burm. fil.) J. Kern

```json
{"parsed":true,"parseQuality":1,"verbatim":"Fimbristylis ovata (Burm. f.) J. Kern","normalized":"Fimbristylis ovata (Burm. fil.) J. Kern","canonical":{"stemmed":"Fimbristylis ouat","simple":"Fimbristylis ovata","full":"Fimbristylis ovata"},"cardinality":2,"authorship":{"verbatim":"(Burm. f.) J. Kern","normalized":"(Burm. fil.) J. Kern","authors":["Burm. fil.","J. Kern"],"originalAuth":{"authors":["Burm. fil."]},"combinationAuth":{"authors":["J. Kern"]}},"details":{"species":{"genus":"Fimbristylis","species":"ovata","authorship":{"verbatim":"(Burm. f.) J. Kern","normalized":"(Burm. fil.) J. Kern","authors":["Burm. fil.","J. Kern"],"originalAuth":{"authors":["Burm. fil."]},"combinationAuth":{"authors":["J. Kern"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":18},{"wordType":"authorWord","start":20,"end":25},{"wordType":"authorWordFilius","start":26,"end":28},{"wordType":"authorWord","start":30,"end":32},{"wordType":"authorWord","start":33,"end":37}],"id":"01207e0b-8de4-5a4e-99fc-e60b581c0d1c","parserVersion":"test_version"}
```

Name: Carex chordorrhiza Ehrh. ex L. f.

Canonical: Carex chordorrhiza

Authorship: Ehrh. ex L. fil.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Carex chordorrhiza Ehrh. ex L. f.","normalized":"Carex chordorrhiza Ehrh. ex L. fil.","canonical":{"stemmed":"Carex chordorrhiz","simple":"Carex chordorrhiza","full":"Carex chordorrhiza"},"cardinality":2,"authorship":{"verbatim":"Ehrh. ex L. f.","normalized":"Ehrh. ex L. fil.","authors":["Ehrh."],"originalAuth":{"authors":["Ehrh."],"exAuthors":{"authors":["L. fil."]}}},"details":{"species":{"genus":"Carex","species":"chordorrhiza","authorship":{"verbatim":"Ehrh. ex L. f.","normalized":"Ehrh. ex L. fil.","authors":["Ehrh."],"originalAuth":{"authors":["Ehrh."],"exAuthors":{"authors":["L. fil."]}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":18},{"wordType":"authorWord","start":19,"end":24},{"wordType":"authorWord","start":28,"end":30},{"wordType":"authorWordFilius","start":31,"end":33}],"id":"b972d277-3714-5549-9103-869675f490bd","parserVersion":"test_version"}
```

Name: Amelanchier arborea var. arborea (Michx. f.) Fernald

Canonical: Amelanchier arborea var. arborea

Authorship: (Michx. fil.) Fernald

```json
{"parsed":true,"parseQuality":1,"verbatim":"Amelanchier arborea var. arborea (Michx. f.) Fernald","normalized":"Amelanchier arborea var. arborea (Michx. fil.) Fernald","canonical":{"stemmed":"Amelanchier arbore arbore","simple":"Amelanchier arborea arborea","full":"Amelanchier arborea var. arborea"},"cardinality":3,"authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}},"details":{"infraSpecies":{"genus":"Amelanchier","species":"arborea","infraSpecies":[{"value":"arborea","rank":"var.","authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":32},{"wordType":"authorWord","start":34,"end":40},{"wordType":"authorWordFilius","start":41,"end":43},{"wordType":"authorWord","start":45,"end":52}],"id":"1644869c-3e0c-5e7e-a709-a86dee11b917","parserVersion":"test_version"}
```

Name: Cerastium arvense var. fuegianum Hook. f.

Canonical: Cerastium arvense var. fuegianum

Authorship: Hook. fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cerastium arvense var. fuegianum Hook. f.","normalized":"Cerastium arvense var. fuegianum Hook. fil.","canonical":{"stemmed":"Cerastium aruens fuegian","simple":"Cerastium arvense fuegianum","full":"Cerastium arvense var. fuegianum"},"cardinality":3,"authorship":{"verbatim":"Hook. f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraSpecies":{"genus":"Cerastium","species":"arvense","infraSpecies":[{"value":"fuegianum","rank":"var.","authorship":{"verbatim":"Hook. f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":17},{"wordType":"rank","start":18,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":32},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWordFilius","start":39,"end":41}],"id":"f9fb925a-777f-5a2c-892d-bdf11528dbfc","parserVersion":"test_version"}
```

Name: Cerastium arvense var. fuegianum Hook.f.

Canonical: Cerastium arvense var. fuegianum

Authorship: Hook. fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cerastium arvense var. fuegianum Hook.f.","normalized":"Cerastium arvense var. fuegianum Hook. fil.","canonical":{"stemmed":"Cerastium aruens fuegian","simple":"Cerastium arvense fuegianum","full":"Cerastium arvense var. fuegianum"},"cardinality":3,"authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}},"details":{"infraSpecies":{"genus":"Cerastium","species":"arvense","infraSpecies":[{"value":"fuegianum","rank":"var.","authorship":{"verbatim":"Hook.f.","normalized":"Hook. fil.","authors":["Hook. fil."],"originalAuth":{"authors":["Hook. fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":17},{"wordType":"rank","start":18,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":32},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWordFilius","start":38,"end":40}],"id":"35ea20fb-b794-572f-ba90-36c1463e1927","parserVersion":"test_version"}
```

Name: Cerastium arvense ssp. velutinum var. velutinum (Raf.) Britton f.

Canonical: Cerastium arvense subsp. velutinum var. velutinum

Authorship: (Raf.) Britton fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Cerastium arvense ssp. velutinum var. velutinum (Raf.) Britton f.","normalized":"Cerastium arvense subsp. velutinum var. velutinum (Raf.) Britton fil.","canonical":{"stemmed":"Cerastium aruens uelutin uelutin","simple":"Cerastium arvense velutinum velutinum","full":"Cerastium arvense subsp. velutinum var. velutinum"},"cardinality":4,"authorship":{"verbatim":"(Raf.) Britton f.","normalized":"(Raf.) Britton fil.","authors":["Raf.","Britton fil."],"originalAuth":{"authors":["Raf."]},"combinationAuth":{"authors":["Britton fil."]}},"details":{"infraSpecies":{"genus":"Cerastium","species":"arvense","infraSpecies":[{"value":"velutinum","rank":"subsp."},{"value":"velutinum","rank":"var.","authorship":{"verbatim":"(Raf.) Britton f.","normalized":"(Raf.) Britton fil.","authors":["Raf.","Britton fil."],"originalAuth":{"authors":["Raf."]},"combinationAuth":{"authors":["Britton fil."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":17},{"wordType":"rank","start":18,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":32},{"wordType":"rank","start":33,"end":37},{"wordType":"infraspecificEpithet","start":38,"end":47},{"wordType":"authorWord","start":49,"end":53},{"wordType":"authorWord","start":55,"end":62},{"wordType":"authorWordFilius","start":63,"end":65}],"id":"c7841295-3aa3-5c40-8adf-88d177f74cbe","parserVersion":"test_version"}
```

Name: Jacquemontia spiciflora (Choisy) Hall. fil.

Canonical: Jacquemontia spiciflora

Authorship: (Choisy) Hall. fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Jacquemontia spiciflora (Choisy) Hall. fil.","normalized":"Jacquemontia spiciflora (Choisy) Hall. fil.","canonical":{"stemmed":"Jacquemontia spiciflor","simple":"Jacquemontia spiciflora","full":"Jacquemontia spiciflora"},"cardinality":2,"authorship":{"verbatim":"(Choisy) Hall. fil.","normalized":"(Choisy) Hall. fil.","authors":["Choisy","Hall. fil."],"originalAuth":{"authors":["Choisy"]},"combinationAuth":{"authors":["Hall. fil."]}},"details":{"species":{"genus":"Jacquemontia","species":"spiciflora","authorship":{"verbatim":"(Choisy) Hall. fil.","normalized":"(Choisy) Hall. fil.","authors":["Choisy","Hall. fil."],"originalAuth":{"authors":["Choisy"]},"combinationAuth":{"authors":["Hall. fil."]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":31},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWordFilius","start":39,"end":43}],"id":"14a98945-4e97-5c13-a0b9-97741641a6a4","parserVersion":"test_version"}
```

Name: Amelanchier arborea f. hirsuta (Michx. f.) Fernald

Canonical: Amelanchier arborea f. hirsuta

Authorship: (Michx. fil.) Fernald

```json
{"parsed":true,"parseQuality":1,"verbatim":"Amelanchier arborea f. hirsuta (Michx. f.) Fernald","normalized":"Amelanchier arborea f. hirsuta (Michx. fil.) Fernald","canonical":{"stemmed":"Amelanchier arbore hirsut","simple":"Amelanchier arborea hirsuta","full":"Amelanchier arborea f. hirsuta"},"cardinality":3,"authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}},"details":{"infraSpecies":{"genus":"Amelanchier","species":"arborea","infraSpecies":[{"value":"hirsuta","rank":"f.","authorship":{"verbatim":"(Michx. f.) Fernald","normalized":"(Michx. fil.) Fernald","authors":["Michx. fil.","Fernald"],"originalAuth":{"authors":["Michx. fil."]},"combinationAuth":{"authors":["Fernald"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":19},{"wordType":"rank","start":20,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":30},{"wordType":"authorWord","start":32,"end":38},{"wordType":"authorWordFilius","start":39,"end":41},{"wordType":"authorWord","start":43,"end":50}],"id":"f5786fa9-2b40-5ee4-8786-ffe86ed02ab5","parserVersion":"test_version"}
```

Name: Betula pendula fo. dalecarlica (L. f.) C.K. Schneid.

Canonical: Betula pendula f. dalecarlica

Authorship: (L. fil.) C. K. Schneid.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Betula pendula fo. dalecarlica (L. f.) C.K. Schneid.","normalized":"Betula pendula f. dalecarlica (L. fil.) C. K. Schneid.","canonical":{"stemmed":"Betula pendul dalecarlic","simple":"Betula pendula dalecarlica","full":"Betula pendula f. dalecarlica"},"cardinality":3,"authorship":{"verbatim":"(L. f.) C.K. Schneid.","normalized":"(L. fil.) C. K. Schneid.","authors":["L. fil.","C. K. Schneid."],"originalAuth":{"authors":["L. fil."]},"combinationAuth":{"authors":["C. K. Schneid."]}},"details":{"infraSpecies":{"genus":"Betula","species":"pendula","infraSpecies":[{"value":"dalecarlica","rank":"f.","authorship":{"verbatim":"(L. f.) C.K. Schneid.","normalized":"(L. fil.) C. K. Schneid.","authors":["L. fil.","C. K. Schneid."],"originalAuth":{"authors":["L. fil."]},"combinationAuth":{"authors":["C. K. Schneid."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":14},{"wordType":"rank","start":15,"end":18},{"wordType":"infraspecificEpithet","start":19,"end":30},{"wordType":"authorWord","start":32,"end":34},{"wordType":"authorWordFilius","start":35,"end":37},{"wordType":"authorWord","start":39,"end":41},{"wordType":"authorWord","start":41,"end":43},{"wordType":"authorWord","start":44,"end":52}],"id":"4c4ee33c-9738-5542-b22f-2326996aa6f7","parserVersion":"test_version"}
```

Name: Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.

Canonical: Racomitrium canescens f. ericoides

Authorship: (F. Weber ex Brid.) Mönk.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","normalized":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","canonical":{"stemmed":"Racomitrium canescens ericoid","simple":"Racomitrium canescens ericoides","full":"Racomitrium canescens f. ericoides"},"cardinality":3,"authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}},"details":{"infraSpecies":{"genus":"Racomitrium","species":"canescens","infraSpecies":[{"value":"ericoides","rank":"f.","authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"rank","start":22,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":34},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":39,"end":44},{"wordType":"authorWord","start":48,"end":53},{"wordType":"authorWord","start":55,"end":60}],"id":"45a001f1-749f-5803-bd92-93c6d524e9db","parserVersion":"test_version"}
```

Name: Racomitrium canescens forma ericoides (F. Weber ex Brid.) Mönk.

Canonical: Racomitrium canescens f. ericoides

Authorship: (F. Weber ex Brid.) Mönk.

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Racomitrium canescens forma ericoides (F. Weber ex Brid.) Mönk.","normalized":"Racomitrium canescens f. ericoides (F. Weber ex Brid.) Mönk.","canonical":{"stemmed":"Racomitrium canescens ericoid","simple":"Racomitrium canescens ericoides","full":"Racomitrium canescens f. ericoides"},"cardinality":3,"authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}},"details":{"infraSpecies":{"genus":"Racomitrium","species":"canescens","infraSpecies":[{"value":"ericoides","rank":"f.","authorship":{"verbatim":"(F. Weber ex Brid.) Mönk.","normalized":"(F. Weber ex Brid.) Mönk.","authors":["F. Weber","Mönk."],"originalAuth":{"authors":["F. Weber"],"exAuthors":{"authors":["Brid."]}},"combinationAuth":{"authors":["Mönk."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"rank","start":22,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":37},{"wordType":"authorWord","start":39,"end":41},{"wordType":"authorWord","start":42,"end":47},{"wordType":"authorWord","start":51,"end":56},{"wordType":"authorWord","start":58,"end":63}],"id":"8a58ed91-9a71-5278-9bd1-b8e82188e938","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f., Rosenst.

Canonical: Polypodium pectinatum

Authorship: L. fil. & Rosenst.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Polypodium pectinatum L. f., Rosenst.","normalized":"Polypodium pectinatum L. fil. \u0026 Rosenst.","canonical":{"stemmed":"Polypodium pectinat","simple":"Polypodium pectinatum","full":"Polypodium pectinatum"},"cardinality":2,"authorship":{"verbatim":"L. f., Rosenst.","normalized":"L. fil. \u0026 Rosenst.","authors":["L. fil.","Rosenst."],"originalAuth":{"authors":["L. fil.","Rosenst."]}},"details":{"species":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L. f., Rosenst.","normalized":"L. fil. \u0026 Rosenst.","authors":["L. fil.","Rosenst."],"originalAuth":{"authors":["L. fil.","Rosenst."]}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWordFilius","start":25,"end":27},{"wordType":"authorWord","start":29,"end":37}],"id":"bac3cf47-358a-51e2-83a6-6577d0f362af","parserVersion":"test_version"}
```

Name: Polypodium pectinatum L. f.

Canonical: Polypodium pectinatum

Authorship: L. fil.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Polypodium pectinatum L. f.","normalized":"Polypodium pectinatum L. fil.","canonical":{"stemmed":"Polypodium pectinat","simple":"Polypodium pectinatum","full":"Polypodium pectinatum"},"cardinality":2,"authorship":{"verbatim":"L. f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"details":{"species":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"L. f.","normalized":"L. fil.","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWordFilius","start":25,"end":27}],"id":"e4c2c98c-79c9-5ee1-865a-300a0c0287ef","parserVersion":"test_version"}
```

Name: Polypodium pectinatum (L. f.) typica Rosent

Canonical: Polypodium pectinatum typica

Authorship: Rosent

```json
{"parsed":true,"parseQuality":1,"verbatim":"Polypodium pectinatum (L. f.) typica Rosent","normalized":"Polypodium pectinatum (L. fil.) typica Rosent","canonical":{"stemmed":"Polypodium pectinat typic","simple":"Polypodium pectinatum typica","full":"Polypodium pectinatum typica"},"cardinality":3,"authorship":{"verbatim":"Rosent","normalized":"Rosent","authors":["Rosent"],"originalAuth":{"authors":["Rosent"]}},"details":{"infraSpecies":{"genus":"Polypodium","species":"pectinatum","authorship":{"verbatim":"(L. f.)","normalized":"(L. fil.)","authors":["L. fil."],"originalAuth":{"authors":["L. fil."]}},"infraSpecies":[{"value":"typica","authorship":{"verbatim":"Rosent","normalized":"Rosent","authors":["Rosent"],"originalAuth":{"authors":["Rosent"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":23,"end":25},{"wordType":"authorWordFilius","start":26,"end":28},{"wordType":"infraspecificEpithet","start":30,"end":36},{"wordType":"authorWord","start":37,"end":43}],"id":"b345d921-7466-50bb-812c-850b1f368c57","parserVersion":"test_version"}
```

### Names with emend (rectified by) authorship

Name: Chlorobium phaeobacteroides Pfennig, 1968 emend. Imhoff, 2003

Canonical: Chlorobium phaeobacteroides

Authorship: Pfennig 1968 emend. Imhoff 2003

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Emend authors are not required"}],"verbatim":"Chlorobium phaeobacteroides Pfennig, 1968 emend. Imhoff, 2003","normalized":"Chlorobium phaeobacteroides Pfennig 1968 emend. Imhoff 2003","canonical":{"stemmed":"Chlorobium phaeobacteroid","simple":"Chlorobium phaeobacteroides","full":"Chlorobium phaeobacteroides"},"cardinality":2,"authorship":{"verbatim":"Pfennig, 1968 emend. Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}},"bacteria":"yes","details":{"species":{"genus":"Chlorobium","species":"phaeobacteroides","authorship":{"verbatim":"Pfennig, 1968 emend. Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":27},{"wordType":"authorWord","start":28,"end":35},{"wordType":"year","start":37,"end":41},{"wordType":"authorWord","start":49,"end":55},{"wordType":"year","start":57,"end":61}],"id":"4513701d-e56b-54d6-84a7-941bf4b62e69","parserVersion":"test_version"}
```

Name: Chlorobium phaeobacteroides Pfennig, 1968 emend Imhoff, 2003

Canonical: Chlorobium phaeobacteroides

Authorship: Pfennig 1968 emend. Imhoff 2003

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"`emend` without a period"},{"quality":2,"warning":"Emend authors are not required"}],"verbatim":"Chlorobium phaeobacteroides Pfennig, 1968 emend Imhoff, 2003","normalized":"Chlorobium phaeobacteroides Pfennig 1968 emend. Imhoff 2003","canonical":{"stemmed":"Chlorobium phaeobacteroid","simple":"Chlorobium phaeobacteroides","full":"Chlorobium phaeobacteroides"},"cardinality":2,"authorship":{"verbatim":"Pfennig, 1968 emend Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}},"bacteria":"yes","details":{"species":{"genus":"Chlorobium","species":"phaeobacteroides","authorship":{"verbatim":"Pfennig, 1968 emend Imhoff, 2003","normalized":"Pfennig 1968 emend. Imhoff 2003","year":"1968","authors":["Pfennig"],"originalAuth":{"authors":["Pfennig"],"year":{"year":"1968"},"emendAuthors":{"authors":["Imhoff"],"year":{"year":"2003"}}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":27},{"wordType":"authorWord","start":28,"end":35},{"wordType":"year","start":37,"end":41},{"wordType":"authorWord","start":48,"end":54},{"wordType":"year","start":56,"end":60}],"id":"3cbaceda-83c2-5e36-b170-4f13837782dc","parserVersion":"test_version"}
```

### "Tail" annotations

Name: Dryopteris X separabilis Small (pro sp.)

Canonical: Dryopteris × separabilis

Authorship: Small

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Dryopteris X separabilis Small (pro sp.)","normalized":"Dryopteris × separabilis Small","canonical":{"stemmed":"Dryopteris separabil","simple":"Dryopteris separabilis","full":"Dryopteris × separabilis"},"cardinality":2,"authorship":{"verbatim":"Small","normalized":"Small","authors":["Small"],"originalAuth":{"authors":["Small"]}},"hybrid":"NAMED_HYBRID","tail":" (pro sp.)","details":{"species":{"genus":"Dryopteris","species":"separabilis Small","authorship":{"verbatim":"Small","normalized":"Small","authors":["Small"],"originalAuth":{"authors":["Small"]}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"hybridChar","start":11,"end":12},{"wordType":"specificEpithet","start":13,"end":24},{"wordType":"authorWord","start":25,"end":30}],"id":"34bf83d8-0466-51c4-b95d-70e583ba1c9f","parserVersion":"test_version"}
```

### Abbreviated words after a name

Name: Graphis scripta L. a.b pulverulenta

Canonical: Graphis scripta

Authorship: L.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Graphis scripta L. a.b pulverulenta","normalized":"Graphis scripta L.","canonical":{"stemmed":"Graphis script","simple":"Graphis scripta","full":"Graphis scripta"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" a.b pulverulenta","details":{"species":{"genus":"Graphis","species":"scripta","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"authorWord","start":16,"end":18}],"id":"ecb4751f-7d9e-5868-8ef7-c96f6ef07f2d","parserVersion":"test_version"}
```

Name: Cetraria iberica a.crespo & barreno

Canonical: Cetraria iberica

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Cetraria iberica a.crespo \u0026 barreno","normalized":"Cetraria iberica","canonical":{"stemmed":"Cetraria iberic","simple":"Cetraria iberica","full":"Cetraria iberica"},"cardinality":2,"tail":" a.crespo \u0026 barreno","details":{"species":{"genus":"Cetraria","species":"iberica"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16}],"id":"233626eb-645c-5ca0-bb8b-6f410a078a85","parserVersion":"test_version"}
```

Name: Lecanora achariana a.l.sm.

Canonical: Lecanora achariana

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lecanora achariana a.l.sm.","normalized":"Lecanora achariana","canonical":{"stemmed":"Lecanora acharian","simple":"Lecanora achariana","full":"Lecanora achariana"},"cardinality":2,"tail":" a.l.sm.","details":{"species":{"genus":"Lecanora","species":"achariana"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18}],"id":"4393f813-14e9-5a26-aab0-bf7686463c6a","parserVersion":"test_version"}
```

Name: Arthrosporum populorum a.massal.

Canonical: Arthrosporum populorum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arthrosporum populorum a.massal.","normalized":"Arthrosporum populorum","canonical":{"stemmed":"Arthrosporum populor","simple":"Arthrosporum populorum","full":"Arthrosporum populorum"},"cardinality":2,"tail":" a.massal.","details":{"species":{"genus":"Arthrosporum","species":"populorum"}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":22}],"id":"88db792d-7061-512d-9275-b7fe81493665","parserVersion":"test_version"}
```

Name: Eletica laeviceps ab.lateapicalis Pic

Canonical: Eletica laeviceps

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Eletica laeviceps ab.lateapicalis Pic","normalized":"Eletica laeviceps","canonical":{"stemmed":"Eletica laeuiceps","simple":"Eletica laeviceps","full":"Eletica laeviceps"},"cardinality":2,"tail":" ab.lateapicalis Pic","details":{"species":{"genus":"Eletica","species":"laeviceps"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17}],"id":"12389c9a-7aaf-56d1-8b8a-dffd4b74c58f","parserVersion":"test_version"}
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
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"},{"quality":2,"warning":"Author in upper case"}],"verbatim":"Acanthoderes 4-gibbus RILEY Charles Valentine, 1880","normalized":"Acanthoderes quadrigibbus Riley Charles Valentine 1880","canonical":{"stemmed":"Acanthoderes quadrigibb","simple":"Acanthoderes quadrigibbus","full":"Acanthoderes quadrigibbus"},"cardinality":2,"authorship":{"verbatim":"RILEY Charles Valentine, 1880","normalized":"Riley Charles Valentine 1880","year":"1880","authors":["Riley Charles Valentine"],"originalAuth":{"authors":["Riley Charles Valentine"],"year":{"year":"1880"}}},"details":{"species":{"genus":"Acanthoderes","species":"quadrigibbus","authorship":{"verbatim":"RILEY Charles Valentine, 1880","normalized":"Riley Charles Valentine 1880","year":"1880","authors":["Riley Charles Valentine"],"originalAuth":{"authors":["Riley Charles Valentine"],"year":{"year":"1880"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":21},{"wordType":"authorWord","start":22,"end":27},{"wordType":"authorWord","start":28,"end":35},{"wordType":"authorWord","start":36,"end":45},{"wordType":"year","start":47,"end":51}],"id":"90bb5882-b093-586d-881a-aeabc55f248b","parserVersion":"test_version"}
```

Name: Acrosoma 12-spinosa Keyserling, 1892

Canonical: Acrosoma duodecimspinosa

Authorship: Keyserling 1892

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Acrosoma 12-spinosa Keyserling, 1892","normalized":"Acrosoma duodecimspinosa Keyserling 1892","canonical":{"stemmed":"Acrosoma duodecimspinos","simple":"Acrosoma duodecimspinosa","full":"Acrosoma duodecimspinosa"},"cardinality":2,"authorship":{"verbatim":"Keyserling, 1892","normalized":"Keyserling 1892","year":"1892","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1892"}}},"details":{"species":{"genus":"Acrosoma","species":"duodecimspinosa","authorship":{"verbatim":"Keyserling, 1892","normalized":"Keyserling 1892","year":"1892","authors":["Keyserling"],"originalAuth":{"authors":["Keyserling"],"year":{"year":"1892"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"authorWord","start":20,"end":30},{"wordType":"year","start":32,"end":36}],"id":"d789c68a-4e40-59d8-a763-3ebadac6fdeb","parserVersion":"test_version"}
```

Name: Canuleius 24-spinosus Redtenbacher, 1906

Canonical: Canuleius vigintiquatuorspinosus

Authorship: Redtenbacher 1906

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Canuleius 24-spinosus Redtenbacher, 1906","normalized":"Canuleius vigintiquatuorspinosus Redtenbacher 1906","canonical":{"stemmed":"Canuleius uigintiquatuorspinos","simple":"Canuleius vigintiquatuorspinosus","full":"Canuleius vigintiquatuorspinosus"},"cardinality":2,"authorship":{"verbatim":"Redtenbacher, 1906","normalized":"Redtenbacher 1906","year":"1906","authors":["Redtenbacher"],"originalAuth":{"authors":["Redtenbacher"],"year":{"year":"1906"}}},"details":{"species":{"genus":"Canuleius","species":"vigintiquatuorspinosus","authorship":{"verbatim":"Redtenbacher, 1906","normalized":"Redtenbacher 1906","year":"1906","authors":["Redtenbacher"],"originalAuth":{"authors":["Redtenbacher"],"year":{"year":"1906"}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":21},{"wordType":"authorWord","start":22,"end":34},{"wordType":"year","start":36,"end":40}],"id":"6dbf79a3-89dd-55ee-aa7d-6394c226cb02","parserVersion":"test_version"}
```

<!-- numeric prefix cannot be more than 2 digits long -->
Name: Canuleius 777-spinosus Redtenbacher, 1906

Canonical: Canuleius

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Canuleius 777-spinosus Redtenbacher, 1906","normalized":"Canuleius","canonical":{"stemmed":"Canuleius","simple":"Canuleius","full":"Canuleius"},"cardinality":1,"tail":" 777-spinosus Redtenbacher, 1906","details":{"uninomial":{"uninomial":"Canuleius"}},"pos":[{"wordType":"uninomial","start":0,"end":9}],"id":"40a1b1cd-0437-5ed8-82bf-8bea169cb8b1","parserVersion":"test_version"}
```

Name: Rhynchophorus 13punctatus Herbst, J.F.W., 1795

Canonical: Rhynchophorus tredecimpunctatus

Authorship: Herbst & J. F. W. 1795

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Rhynchophorus 13punctatus Herbst, J.F.W., 1795","normalized":"Rhynchophorus tredecimpunctatus Herbst \u0026 J. F. W. 1795","canonical":{"stemmed":"Rhynchophorus tredecimpunctat","simple":"Rhynchophorus tredecimpunctatus","full":"Rhynchophorus tredecimpunctatus"},"cardinality":2,"authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}},"details":{"species":{"genus":"Rhynchophorus","species":"tredecimpunctatus","authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":25},{"wordType":"authorWord","start":26,"end":32},{"wordType":"authorWord","start":34,"end":36},{"wordType":"authorWord","start":36,"end":38},{"wordType":"authorWord","start":38,"end":40},{"wordType":"year","start":42,"end":46}],"id":"8724e04d-a1a0-5b5e-9c0e-1c0f586507d6","parserVersion":"test_version"}
```

Name: Rhynchophorus 13.punctatus Herbst, J.F.W., 1795

Canonical: Rhynchophorus tredecimpunctatus

Authorship: Herbst & J. F. W. 1795

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Numeric prefix"}],"verbatim":"Rhynchophorus 13.punctatus Herbst, J.F.W., 1795","normalized":"Rhynchophorus tredecimpunctatus Herbst \u0026 J. F. W. 1795","canonical":{"stemmed":"Rhynchophorus tredecimpunctat","simple":"Rhynchophorus tredecimpunctatus","full":"Rhynchophorus tredecimpunctatus"},"cardinality":2,"authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}},"details":{"species":{"genus":"Rhynchophorus","species":"tredecimpunctatus","authorship":{"verbatim":"Herbst, J.F.W., 1795","normalized":"Herbst \u0026 J. F. W. 1795","year":"1795","authors":["Herbst","J. F. W."],"originalAuth":{"authors":["Herbst","J. F. W."],"year":{"year":"1795"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":26},{"wordType":"authorWord","start":27,"end":33},{"wordType":"authorWord","start":35,"end":37},{"wordType":"authorWord","start":37,"end":39},{"wordType":"authorWord","start":39,"end":41},{"wordType":"year","start":43,"end":47}],"id":"590b3805-23bc-5a94-a7ca-ea89dcfb5ed1","parserVersion":"test_version"}
```

### Non-ASCII UTF-8 characters in a name

Name: Pleurotus ëous (Berk.) Sacc. 1887

Canonical: Pleurotus eous

Authorship: (Berk.) Sacc. 1887

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pleurotus ëous (Berk.) Sacc. 1887","normalized":"Pleurotus eous (Berk.) Sacc. 1887","canonical":{"stemmed":"Pleurotus eo","simple":"Pleurotus eous","full":"Pleurotus eous"},"cardinality":2,"authorship":{"verbatim":"(Berk.) Sacc. 1887","normalized":"(Berk.) Sacc. 1887","authors":["Berk.","Sacc."],"originalAuth":{"authors":["Berk."]},"combinationAuth":{"authors":["Sacc."],"year":{"year":"1887"}}},"details":{"species":{"genus":"Pleurotus","species":"eous","authorship":{"verbatim":"(Berk.) Sacc. 1887","normalized":"(Berk.) Sacc. 1887","authors":["Berk.","Sacc."],"originalAuth":{"authors":["Berk."]},"combinationAuth":{"authors":["Sacc."],"year":{"year":"1887"}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":14},{"wordType":"authorWord","start":16,"end":21},{"wordType":"authorWord","start":23,"end":28},{"wordType":"year","start":29,"end":33}],"id":"fe8c9a43-3480-5598-891d-e2a864781d13","parserVersion":"test_version"}
```

Name: Sténométope laevissimus Bibron 1855

Canonical: Stenometope laevissimus

Authorship: Bibron 1855

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Sténométope laevissimus Bibron 1855","normalized":"Stenometope laevissimus Bibron 1855","canonical":{"stemmed":"Stenometope laeuissim","simple":"Stenometope laevissimus","full":"Stenometope laevissimus"},"cardinality":2,"authorship":{"verbatim":"Bibron 1855","normalized":"Bibron 1855","year":"1855","authors":["Bibron"],"originalAuth":{"authors":["Bibron"],"year":{"year":"1855"}}},"details":{"species":{"genus":"Stenometope","species":"laevissimus","authorship":{"verbatim":"Bibron 1855","normalized":"Bibron 1855","year":"1855","authors":["Bibron"],"originalAuth":{"authors":["Bibron"],"year":{"year":"1855"}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":23},{"wordType":"authorWord","start":24,"end":30},{"wordType":"year","start":31,"end":35}],"id":"363ea9fc-ac47-50e5-ae4b-1bfb104a8e34","parserVersion":"test_version"}
```

Name: Choriozopella trägårdhi Lawrence, 1947

Canonical: Choriozopella traegaordhi

Authorship: Lawrence 1947

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Choriozopella trägårdhi Lawrence, 1947","normalized":"Choriozopella traegaordhi Lawrence 1947","canonical":{"stemmed":"Choriozopella traegaordh","simple":"Choriozopella traegaordhi","full":"Choriozopella traegaordhi"},"cardinality":2,"authorship":{"verbatim":"Lawrence, 1947","normalized":"Lawrence 1947","year":"1947","authors":["Lawrence"],"originalAuth":{"authors":["Lawrence"],"year":{"year":"1947"}}},"details":{"species":{"genus":"Choriozopella","species":"traegaordhi","authorship":{"verbatim":"Lawrence, 1947","normalized":"Lawrence 1947","year":"1947","authors":["Lawrence"],"originalAuth":{"authors":["Lawrence"],"year":{"year":"1947"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":23},{"wordType":"authorWord","start":24,"end":32},{"wordType":"year","start":34,"end":38}],"id":"3d02292a-3526-5364-96c9-f73738b9d2fa","parserVersion":"test_version"}
```

Name: Isoëtes asplundii H. P. Fuchs

Canonical: Isoetes asplundii

Authorship: H. P. Fuchs

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Isoëtes asplundii H. P. Fuchs","normalized":"Isoetes asplundii H. P. Fuchs","canonical":{"stemmed":"Isoetes asplundi","simple":"Isoetes asplundii","full":"Isoetes asplundii"},"cardinality":2,"authorship":{"verbatim":"H. P. Fuchs","normalized":"H. P. Fuchs","authors":["H. P. Fuchs"],"originalAuth":{"authors":["H. P. Fuchs"]}},"details":{"species":{"genus":"Isoetes","species":"asplundii","authorship":{"verbatim":"H. P. Fuchs","normalized":"H. P. Fuchs","authors":["H. P. Fuchs"],"originalAuth":{"authors":["H. P. Fuchs"]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":20},{"wordType":"authorWord","start":21,"end":23},{"wordType":"authorWord","start":24,"end":29}],"id":"8d713775-782a-5083-92a1-ddaf4af9d785","parserVersion":"test_version"}
```

Name: Cerambyx thomæ GMELIN J. F., 1790

Canonical: Cerambyx thomae

Authorship: Gmelin J. F. 1790

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Cerambyx thomæ GMELIN J. F., 1790","normalized":"Cerambyx thomae Gmelin J. F. 1790","canonical":{"stemmed":"Cerambyx thom","simple":"Cerambyx thomae","full":"Cerambyx thomae"},"cardinality":2,"authorship":{"verbatim":"GMELIN J. F., 1790","normalized":"Gmelin J. F. 1790","year":"1790","authors":["Gmelin J. F."],"originalAuth":{"authors":["Gmelin J. F."],"year":{"year":"1790"}}},"details":{"species":{"genus":"Cerambyx","species":"thomae","authorship":{"verbatim":"GMELIN J. F., 1790","normalized":"Gmelin J. F. 1790","year":"1790","authors":["Gmelin J. F."],"originalAuth":{"authors":["Gmelin J. F."],"year":{"year":"1790"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":14},{"wordType":"authorWord","start":15,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWord","start":25,"end":27},{"wordType":"year","start":29,"end":33}],"id":"f9689237-693f-5d6e-b62e-b1622214863e","parserVersion":"test_version"}
```

Name: Campethera cailliautii fülleborni

Canonical: Campethera cailliautii fuelleborni

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Campethera cailliautii fülleborni","normalized":"Campethera cailliautii fuelleborni","canonical":{"stemmed":"Campethera cailliauti fuelleborn","simple":"Campethera cailliautii fuelleborni","full":"Campethera cailliautii fuelleborni"},"cardinality":3,"details":{"infraSpecies":{"genus":"Campethera","species":"cailliautii","infraSpecies":[{"value":"fuelleborni"}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":33}],"id":"6a47e9c0-908f-5141-be93-76490af08606","parserVersion":"test_version"}
```

Name: Östrupia Heiden ex Hustedt, 1935

Canonical: Oestrupia

Authorship: Heiden ex Hustedt 1935

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"},{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Östrupia Heiden ex Hustedt, 1935","normalized":"Oestrupia Heiden ex Hustedt 1935","canonical":{"stemmed":"Oestrupia","simple":"Oestrupia","full":"Oestrupia"},"cardinality":1,"authorship":{"verbatim":"Heiden ex Hustedt, 1935","normalized":"Heiden ex Hustedt 1935","authors":["Heiden"],"originalAuth":{"authors":["Heiden"],"exAuthors":{"authors":["Hustedt"],"year":{"year":"1935"}}}},"details":{"uninomial":{"uninomial":"Oestrupia","authorship":{"verbatim":"Heiden ex Hustedt, 1935","normalized":"Heiden ex Hustedt 1935","authors":["Heiden"],"originalAuth":{"authors":["Heiden"],"exAuthors":{"authors":["Hustedt"],"year":{"year":"1935"}}}}}},"pos":[{"wordType":"uninomial","start":0,"end":8},{"wordType":"authorWord","start":9,"end":15},{"wordType":"authorWord","start":19,"end":26},{"wordType":"year","start":28,"end":32}],"id":"940aba5b-2334-5846-98ba-ce29c7305734","parserVersion":"test_version"}
```

### Epithets with an apostrophe

Name: Junellia o'donelli Moldenke, 1946

Canonical: Junellia odonelli

Authorship: Moldenke 1946

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Junellia o'donelli Moldenke, 1946","normalized":"Junellia odonelli Moldenke 1946","canonical":{"stemmed":"Junellia odonell","simple":"Junellia odonelli","full":"Junellia odonelli"},"cardinality":2,"authorship":{"verbatim":"Moldenke, 1946","normalized":"Moldenke 1946","year":"1946","authors":["Moldenke"],"originalAuth":{"authors":["Moldenke"],"year":{"year":"1946"}}},"details":{"species":{"genus":"Junellia","species":"odonelli","authorship":{"verbatim":"Moldenke, 1946","normalized":"Moldenke 1946","year":"1946","authors":["Moldenke"],"originalAuth":{"authors":["Moldenke"],"year":{"year":"1946"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":19,"end":27},{"wordType":"year","start":29,"end":33}],"id":"e39a2d98-6ab2-5fb3-9aae-c48aa86c6026","parserVersion":"test_version"}
```

Name: Trophon d'orbignyi Carcelles, 1946

Canonical: Trophon dorbignyi

Authorship: Carcelles 1946

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Trophon d'orbignyi Carcelles, 1946","normalized":"Trophon dorbignyi Carcelles 1946","canonical":{"stemmed":"Trophon dorbigny","simple":"Trophon dorbignyi","full":"Trophon dorbignyi"},"cardinality":2,"authorship":{"verbatim":"Carcelles, 1946","normalized":"Carcelles 1946","year":"1946","authors":["Carcelles"],"originalAuth":{"authors":["Carcelles"],"year":{"year":"1946"}}},"details":{"species":{"genus":"Trophon","species":"dorbignyi","authorship":{"verbatim":"Carcelles, 1946","normalized":"Carcelles 1946","year":"1946","authors":["Carcelles"],"originalAuth":{"authors":["Carcelles"],"year":{"year":"1946"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":18},{"wordType":"authorWord","start":19,"end":28},{"wordType":"year","start":30,"end":34}],"id":"935d4414-05d4-5c16-be30-466f6144b666","parserVersion":"test_version"}
```

Name: Phrynosoma m’callii

Canonical: Phrynosoma mcallii

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Not an ASCII apostrophe"},{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Phrynosoma m’callii","normalized":"Phrynosoma mcallii","canonical":{"stemmed":"Phrynosoma mcalli","simple":"Phrynosoma mcallii","full":"Phrynosoma mcallii"},"cardinality":2,"details":{"species":{"genus":"Phrynosoma","species":"mcallii"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19}],"id":"7907df5c-50f2-532c-a8fe-e5b75f924f73","parserVersion":"test_version"}
```

Name: Arca m'coyi Tenison-Woods, 1878

Canonical: Arca mcoyi

Authorship: Tenison-Woods 1878

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Arca m'coyi Tenison-Woods, 1878","normalized":"Arca mcoyi Tenison-Woods 1878","canonical":{"stemmed":"Arca mcoy","simple":"Arca mcoyi","full":"Arca mcoyi"},"cardinality":2,"authorship":{"verbatim":"Tenison-Woods, 1878","normalized":"Tenison-Woods 1878","year":"1878","authors":["Tenison-Woods"],"originalAuth":{"authors":["Tenison-Woods"],"year":{"year":"1878"}}},"details":{"species":{"genus":"Arca","species":"mcoyi","authorship":{"verbatim":"Tenison-Woods, 1878","normalized":"Tenison-Woods 1878","year":"1878","authors":["Tenison-Woods"],"originalAuth":{"authors":["Tenison-Woods"],"year":{"year":"1878"}}}}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":11},{"wordType":"authorWord","start":12,"end":25},{"wordType":"year","start":27,"end":31}],"id":"fa855178-bdde-5ebf-b6b1-c1a1aa60bffa","parserVersion":"test_version"}
```

Name: Nucula m'andrewii Hanley, 1860

Canonical: Nucula mandrewii

Authorship: Hanley 1860

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Nucula m'andrewii Hanley, 1860","normalized":"Nucula mandrewii Hanley 1860","canonical":{"stemmed":"Nucula mandrewi","simple":"Nucula mandrewii","full":"Nucula mandrewii"},"cardinality":2,"authorship":{"verbatim":"Hanley, 1860","normalized":"Hanley 1860","year":"1860","authors":["Hanley"],"originalAuth":{"authors":["Hanley"],"year":{"year":"1860"}}},"details":{"species":{"genus":"Nucula","species":"mandrewii","authorship":{"verbatim":"Hanley, 1860","normalized":"Hanley 1860","year":"1860","authors":["Hanley"],"originalAuth":{"authors":["Hanley"],"year":{"year":"1860"}}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17},{"wordType":"authorWord","start":18,"end":24},{"wordType":"year","start":26,"end":30}],"id":"8bbc3b0e-149d-5ede-9f12-b516b085da9d","parserVersion":"test_version"}
```

Name: Eristalis l'herminierii Macquart

Canonical: Eristalis lherminierii

Authorship: Macquart

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Eristalis l'herminierii Macquart","normalized":"Eristalis lherminierii Macquart","canonical":{"stemmed":"Eristalis lherminieri","simple":"Eristalis lherminierii","full":"Eristalis lherminierii"},"cardinality":2,"authorship":{"verbatim":"Macquart","normalized":"Macquart","authors":["Macquart"],"originalAuth":{"authors":["Macquart"]}},"details":{"species":{"genus":"Eristalis","species":"lherminierii","authorship":{"verbatim":"Macquart","normalized":"Macquart","authors":["Macquart"],"originalAuth":{"authors":["Macquart"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":23},{"wordType":"authorWord","start":24,"end":32}],"id":"f7ccb013-ad48-5424-9c26-01657275de9a","parserVersion":"test_version"}
```

Name: Odynerus o'neili Cameron

Canonical: Odynerus oneili

Authorship: Cameron

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Odynerus o'neili Cameron","normalized":"Odynerus oneili Cameron","canonical":{"stemmed":"Odynerus oneil","simple":"Odynerus oneili","full":"Odynerus oneili"},"cardinality":2,"authorship":{"verbatim":"Cameron","normalized":"Cameron","authors":["Cameron"],"originalAuth":{"authors":["Cameron"]}},"details":{"species":{"genus":"Odynerus","species":"oneili","authorship":{"verbatim":"Cameron","normalized":"Cameron","authors":["Cameron"],"originalAuth":{"authors":["Cameron"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":17,"end":24}],"id":"39218b39-39f9-5f0d-917a-d5e57301d91c","parserVersion":"test_version"}
```

Name: Serjania meridionalis Cambess. var. o'donelli F.A. Barkley

Canonical: Serjania meridionalis var. odonelli

Authorship: F. A. Barkley

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Serjania meridionalis Cambess. var. o'donelli F.A. Barkley","normalized":"Serjania meridionalis Cambess. var. odonelli F. A. Barkley","canonical":{"stemmed":"Serjania meridional odonell","simple":"Serjania meridionalis odonelli","full":"Serjania meridionalis var. odonelli"},"cardinality":3,"authorship":{"verbatim":"F.A. Barkley","normalized":"F. A. Barkley","authors":["F. A. Barkley"],"originalAuth":{"authors":["F. A. Barkley"]}},"details":{"infraSpecies":{"genus":"Serjania","species":"meridionalis","authorship":{"verbatim":"Cambess.","normalized":"Cambess.","authors":["Cambess."],"originalAuth":{"authors":["Cambess."]}},"infraSpecies":[{"value":"odonelli","rank":"var.","authorship":{"verbatim":"F.A. Barkley","normalized":"F. A. Barkley","authors":["F. A. Barkley"],"originalAuth":{"authors":["F. A. Barkley"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":22,"end":30},{"wordType":"rank","start":31,"end":35},{"wordType":"infraspecificEpithet","start":36,"end":45},{"wordType":"authorWord","start":46,"end":48},{"wordType":"authorWord","start":48,"end":50},{"wordType":"authorWord","start":51,"end":58}],"id":"019a8f2c-279d-5211-9bfb-5f288795ed73","parserVersion":"test_version"}
```

### Digraph unicode characters

Name: Æschopalæa grisella Pascoe, 1864

Canonical: Aeschopalaea grisella

Authorship: Pascoe 1864

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Æschopalæa grisella Pascoe, 1864","normalized":"Aeschopalaea grisella Pascoe 1864","canonical":{"stemmed":"Aeschopalaea grisell","simple":"Aeschopalaea grisella","full":"Aeschopalaea grisella"},"cardinality":2,"authorship":{"verbatim":"Pascoe, 1864","normalized":"Pascoe 1864","year":"1864","authors":["Pascoe"],"originalAuth":{"authors":["Pascoe"],"year":{"year":"1864"}}},"details":{"species":{"genus":"Aeschopalaea","species":"grisella","authorship":{"verbatim":"Pascoe, 1864","normalized":"Pascoe 1864","year":"1864","authors":["Pascoe"],"originalAuth":{"authors":["Pascoe"],"year":{"year":"1864"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19},{"wordType":"authorWord","start":20,"end":26},{"wordType":"year","start":28,"end":32}],"id":"82afddf5-4bac-5858-a6a3-93b270b844e8","parserVersion":"test_version"}
```

Name: Læptura laetifica Dow, 1913

Canonical: Laeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Læptura laetifica Dow, 1913","normalized":"Laeptura laetifica Dow 1913","canonical":{"stemmed":"Laeptura laetific","simple":"Laeptura laetifica","full":"Laeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Laeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":21},{"wordType":"year","start":23,"end":27}],"id":"dc1da297-0a85-583d-9a72-d888ddb37ae7","parserVersion":"test_version"}
```

Name: Leptura lætifica Dow, 1913

Canonical: Leptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leptura lætifica Dow, 1913","normalized":"Leptura laetifica Dow 1913","canonical":{"stemmed":"Leptura laetific","simple":"Leptura laetifica","full":"Leptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":16},{"wordType":"authorWord","start":17,"end":20},{"wordType":"year","start":22,"end":26}],"id":"0067abce-1fa8-5911-8176-011065a113a6","parserVersion":"test_version"}
```

Name: Leptura leætifica Dow, 1913

Canonical: Leptura leaetifica

Authorship: Dow 1913

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leptura leætifica Dow, 1913","normalized":"Leptura leaetifica Dow 1913","canonical":{"stemmed":"Leptura leaetific","simple":"Leptura leaetifica","full":"Leptura leaetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leptura","species":"leaetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":21},{"wordType":"year","start":23,"end":27}],"id":"06e6f378-8a12-500a-bab1-27e8b9c6b0cb","parserVersion":"test_version"}
```

Name: Leæptura laetifica Dow, 1913

Canonical: Leaeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leæptura laetifica Dow, 1913","normalized":"Leaeptura laetifica Dow 1913","canonical":{"stemmed":"Leaeptura laetific","simple":"Leaeptura laetifica","full":"Leaeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leaeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":19,"end":22},{"wordType":"year","start":24,"end":28}],"id":"18311671-6006-5382-b3b9-d9e959fa61c1","parserVersion":"test_version"}
```

Name: Leœptura laetifica Dow, 1913

Canonical: Leoeptura laetifica

Authorship: Dow 1913

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Leœptura laetifica Dow, 1913","normalized":"Leoeptura laetifica Dow 1913","canonical":{"stemmed":"Leoeptura laetific","simple":"Leoeptura laetifica","full":"Leoeptura laetifica"},"cardinality":2,"authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}},"details":{"species":{"genus":"Leoeptura","species":"laetifica","authorship":{"verbatim":"Dow, 1913","normalized":"Dow 1913","year":"1913","authors":["Dow"],"originalAuth":{"authors":["Dow"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":19,"end":22},{"wordType":"year","start":24,"end":28}],"id":"c31a86ea-3f68-52b4-a746-5ca921816357","parserVersion":"test_version"}
```

Name: Ærenea cognata Lacordaire, 1872

Canonical: Aerenea cognata

Authorship: Lacordaire 1872

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Ærenea cognata Lacordaire, 1872","normalized":"Aerenea cognata Lacordaire 1872","canonical":{"stemmed":"Aerenea cognat","simple":"Aerenea cognata","full":"Aerenea cognata"},"cardinality":2,"authorship":{"verbatim":"Lacordaire, 1872","normalized":"Lacordaire 1872","year":"1872","authors":["Lacordaire"],"originalAuth":{"authors":["Lacordaire"],"year":{"year":"1872"}}},"details":{"species":{"genus":"Aerenea","species":"cognata","authorship":{"verbatim":"Lacordaire, 1872","normalized":"Lacordaire 1872","year":"1872","authors":["Lacordaire"],"originalAuth":{"authors":["Lacordaire"],"year":{"year":"1872"}}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":14},{"wordType":"authorWord","start":15,"end":25},{"wordType":"year","start":27,"end":31}],"id":"e7f394a9-59f3-5a9c-b375-d6949e232694","parserVersion":"test_version"}
```

Name: Œdicnemus capensis

Canonical: Oedicnemus capensis

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Œdicnemus capensis","normalized":"Oedicnemus capensis","canonical":{"stemmed":"Oedicnemus capens","simple":"Oedicnemus capensis","full":"Oedicnemus capensis"},"cardinality":2,"details":{"species":{"genus":"Oedicnemus","species":"capensis"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":18}],"id":"33dcf668-48f3-5504-87c1-fe6646a51189","parserVersion":"test_version"}
```

Name: Œnanthe œnanthe

Canonical: Oenanthe oenanthe

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Œnanthe œnanthe","normalized":"Oenanthe oenanthe","canonical":{"stemmed":"Oenanthe oenanth","simple":"Oenanthe oenanthe","full":"Oenanthe oenanthe"},"cardinality":2,"details":{"species":{"genus":"Oenanthe","species":"oenanthe"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15}],"id":"3e4ce8df-36d0-5529-9725-8336fa694c9a","parserVersion":"test_version"}
```

Name: Hördeum vulgare cœrulescens

Canonical: Hoerdeum vulgare coerulescens

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hördeum vulgare cœrulescens","normalized":"Hoerdeum vulgare coerulescens","canonical":{"stemmed":"Hoerdeum uulgar coerulescens","simple":"Hoerdeum vulgare coerulescens","full":"Hoerdeum vulgare coerulescens"},"cardinality":3,"details":{"infraSpecies":{"genus":"Hoerdeum","species":"vulgare","infraSpecies":[{"value":"coerulescens"}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"infraspecificEpithet","start":16,"end":27}],"id":"44916bbf-7112-5604-b691-e425447974d4","parserVersion":"test_version"}
```

Name: Hordeum vulgare cœrulescens Metzger

Canonical: Hordeum vulgare coerulescens

Authorship: Metzger

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hordeum vulgare cœrulescens Metzger","normalized":"Hordeum vulgare coerulescens Metzger","canonical":{"stemmed":"Hordeum uulgar coerulescens","simple":"Hordeum vulgare coerulescens","full":"Hordeum vulgare coerulescens"},"cardinality":3,"authorship":{"verbatim":"Metzger","normalized":"Metzger","authors":["Metzger"],"originalAuth":{"authors":["Metzger"]}},"details":{"infraSpecies":{"genus":"Hordeum","species":"vulgare","infraSpecies":[{"value":"coerulescens","authorship":{"verbatim":"Metzger","normalized":"Metzger","authors":["Metzger"],"originalAuth":{"authors":["Metzger"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"infraspecificEpithet","start":16,"end":27},{"wordType":"authorWord","start":28,"end":35}],"id":"3029cf1f-da59-5955-86af-40c0b57bd59d","parserVersion":"test_version"}
```

Name: Hordeum vulgare f. cœrulescens

Canonical: Hordeum vulgare f. coerulescens

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Hordeum vulgare f. cœrulescens","normalized":"Hordeum vulgare f. coerulescens","canonical":{"stemmed":"Hordeum uulgar coerulescens","simple":"Hordeum vulgare coerulescens","full":"Hordeum vulgare f. coerulescens"},"cardinality":3,"details":{"infraSpecies":{"genus":"Hordeum","species":"vulgare","infraSpecies":[{"value":"coerulescens","rank":"f."}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"rank","start":16,"end":18},{"wordType":"infraspecificEpithet","start":19,"end":30}],"id":"27dd2ab3-8bf9-5f72-90bb-5c94530822f2","parserVersion":"test_version"}
```

### Old style s (ſ)

Name: Musca domeſtica Linnaeus 1758

Canonical: Musca domestica

Authorship: Linnaeus 1758

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Musca domeſtica Linnaeus 1758","normalized":"Musca domestica Linnaeus 1758","canonical":{"stemmed":"Musca domestic","simple":"Musca domestica","full":"Musca domestica"},"cardinality":2,"authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Musca","species":"domestica","authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":15},{"wordType":"authorWord","start":16,"end":24},{"wordType":"year","start":25,"end":29}],"id":"a9f11057-210a-51d0-8402-79d4075607d3","parserVersion":"test_version"}
```

Name: Amphisbæna fuliginoſa Linnaeus 1758

Canonical: Amphisbaena fuliginosa

Authorship: Linnaeus 1758

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Amphisbæna fuliginoſa Linnaeus 1758","normalized":"Amphisbaena fuliginosa Linnaeus 1758","canonical":{"stemmed":"Amphisbaena fuliginos","simple":"Amphisbaena fuliginosa","full":"Amphisbaena fuliginosa"},"cardinality":2,"authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Amphisbaena","species":"fuliginosa","authorship":{"verbatim":"Linnaeus 1758","normalized":"Linnaeus 1758","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":22,"end":30},{"wordType":"year","start":31,"end":35}],"id":"d2f6423b-7a8f-5389-a286-c074fb634c5a","parserVersion":"test_version"}
```

Name: Dreyfusia nüßlini

Canonical: Dreyfusia nuesslini

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Dreyfusia nüßlini","normalized":"Dreyfusia nuesslini","canonical":{"stemmed":"Dreyfusia nuesslin","simple":"Dreyfusia nuesslini","full":"Dreyfusia nuesslini"},"cardinality":2,"details":{"species":{"genus":"Dreyfusia","species":"nuesslini"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":17}],"id":"27679e50-c41b-5a3d-b619-d378d503be8c","parserVersion":"test_version"}
```

### Miscellaneous diacritics

Name: Pärdosa

Canonical: Paerdosa

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pärdosa","normalized":"Paerdosa","canonical":{"stemmed":"Paerdosa","simple":"Paerdosa","full":"Paerdosa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Paerdosa"}},"pos":[{"wordType":"uninomial","start":0,"end":7}],"id":"3f493cea-a62c-5bfc-a9a8-e3305e6936db","parserVersion":"test_version"}
```

Name: Pårdosa

Canonical: Paordosa

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pårdosa","normalized":"Paordosa","canonical":{"stemmed":"Paordosa","simple":"Paordosa","full":"Paordosa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Paordosa"}},"pos":[{"wordType":"uninomial","start":0,"end":7}],"id":"eead0d2e-5f37-503c-add2-e344c341be20","parserVersion":"test_version"}
```

Name: Pardøsa

Canonical: Pardoesa

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pardøsa","normalized":"Pardoesa","canonical":{"stemmed":"Pardoesa","simple":"Pardoesa","full":"Pardoesa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pardoesa"}},"pos":[{"wordType":"uninomial","start":0,"end":7}],"id":"6922fdef-226d-59fc-9cc6-7b446d7ce37b","parserVersion":"test_version"}
```

Name: Pardösa

Canonical: Pardoesa

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Pardösa","normalized":"Pardoesa","canonical":{"stemmed":"Pardoesa","simple":"Pardoesa","full":"Pardoesa"},"cardinality":1,"details":{"uninomial":{"uninomial":"Pardoesa"}},"pos":[{"wordType":"uninomial","start":0,"end":7}],"id":"7873dfb8-fc08-50e8-bd23-e94deb9317bc","parserVersion":"test_version"}
```

Name: Rühlella

Canonical: Ruehlella

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard characters in canonical"}],"verbatim":"Rühlella","normalized":"Ruehlella","canonical":{"stemmed":"Ruehlella","simple":"Ruehlella","full":"Ruehlella"},"cardinality":1,"details":{"uninomial":{"uninomial":"Ruehlella"}},"pos":[{"wordType":"uninomial","start":0,"end":8}],"id":"228b2714-3726-5ae8-b802-59bdbc8d20a6","parserVersion":"test_version"}
```

### Open Nomenclature ('approximate' names)

<!-- Open nomenclature -- cf., aff., sp., etc. -->
Name: Solygia ? distanti

Canonical: Solygia

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Solygia ? distanti","normalized":"Solygia","canonical":{"stemmed":"Solygia","simple":"Solygia","full":"Solygia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Solygia","approximationMarker":"?","ignored":" distanti"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"approximationMarker","start":8,"end":9}],"id":"b9e3508f-1c0e-554c-8642-dd1cfd02631c","parserVersion":"test_version"}
```

<!-- Ambiguity -- can be an unknown author or approx name-->
Name: Buteo borealis ? ventralis

Canonical: Buteo borealis ventralis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Author as a question mark"},{"quality":3,"warning":"Author is too short"},{"quality":2,"warning":"Author is unknown"}],"verbatim":"Buteo borealis ? ventralis","normalized":"Buteo borealis anon. ventralis","canonical":{"stemmed":"Buteo boreal uentral","simple":"Buteo borealis ventralis","full":"Buteo borealis ventralis"},"cardinality":3,"details":{"infraSpecies":{"genus":"Buteo","species":"borealis","authorship":{"verbatim":"?","normalized":"anon.","authors":["anon."],"originalAuth":{"authors":["anon."]}},"infraSpecies":[{"value":"ventralis"}]}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":14},{"wordType":"authorWord","start":15,"end":16},{"wordType":"infraspecificEpithet","start":17,"end":26}],"id":"d26a4791-4858-5239-8a57-c88957d40919","parserVersion":"test_version"}
```

Name: Euxoa nr. idahoensis sp. 1clay

Canonical: Euxoa

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Euxoa nr. idahoensis sp. 1clay","normalized":"Euxoa","canonical":{"stemmed":"Euxoa","simple":"Euxoa","full":"Euxoa"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Euxoa","approximationMarker":"nr.","ignored":" idahoensis sp. 1clay"}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"approximationMarker","start":6,"end":9}],"id":"02a664be-422a-56cb-b431-99aecf793721","parserVersion":"test_version"}
```

Name: Acarinina aff. pentacamerata

Canonical: Acarinina

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acarinina aff. pentacamerata","normalized":"Acarinina","canonical":{"stemmed":"Acarinina","simple":"Acarinina","full":"Acarinina"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acarinina","approximationMarker":"aff.","ignored":" pentacamerata"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"approximationMarker","start":10,"end":14}],"id":"c4ab66ee-79a2-5100-8b87-20e60cf2a358","parserVersion":"test_version"}
```

Name: Acarinina aff pentacamerata

Canonical: Acarinina

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acarinina aff pentacamerata","normalized":"Acarinina","canonical":{"stemmed":"Acarinina","simple":"Acarinina","full":"Acarinina"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acarinina","approximationMarker":"aff","ignored":" pentacamerata"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"approximationMarker","start":10,"end":13}],"id":"06a32183-0aa7-5a00-9753-46db1141daa4","parserVersion":"test_version"}
```

Name: Sphingomonas sp. 37

Canonical: Sphingomonas

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Sphingomonas sp. 37","normalized":"Sphingomonas","canonical":{"stemmed":"Sphingomonas","simple":"Sphingomonas","full":"Sphingomonas"},"cardinality":0,"bacteria":"yes","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Sphingomonas","approximationMarker":"sp.","ignored":" 37"}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"approximationMarker","start":13,"end":16}],"id":"1daffd3a-f4de-58d9-91e3-ae4d08a50ce0","parserVersion":"test_version"}
```

Name: Thryothorus leucotis spp. bogotensis

Canonical: Thryothorus leucotis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Thryothorus leucotis spp. bogotensis","normalized":"Thryothorus leucotis","canonical":{"stemmed":"Thryothorus leucot","simple":"Thryothorus leucotis","full":"Thryothorus leucotis"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Thryothorus","species":"leucotis","approximationMarker":"spp.","ignored":" bogotensis"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":20},{"wordType":"approximationMarker","start":21,"end":25}],"id":"d2cb7212-ff62-5e31-9ab9-31214a9782d5","parserVersion":"test_version"}
```

Name: Endoxyla sp. GM-, 2003

Canonical: Endoxyla

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Endoxyla sp. GM-, 2003","normalized":"Endoxyla","canonical":{"stemmed":"Endoxyla","simple":"Endoxyla","full":"Endoxyla"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Endoxyla","approximationMarker":"sp.","ignored":" GM-, 2003"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"approximationMarker","start":9,"end":12}],"id":"8a80bfee-947d-5602-9958-a2338ff46a4d","parserVersion":"test_version"}
```

Name: X Aegilotrichum sp.

Canonical: × Aegilotrichum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"X Aegilotrichum sp.","normalized":"× Aegilotrichum","canonical":{"stemmed":"Aegilotrichum","simple":"Aegilotrichum","full":"× Aegilotrichum"},"cardinality":0,"hybrid":"NAMED_HYBRID","surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Aegilotrichum","approximationMarker":"sp."}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"genus","start":2,"end":15},{"wordType":"approximationMarker","start":16,"end":19}],"id":"308357ff-7f86-53b9-955b-88a52ef7623a","parserVersion":"test_version"}
```

Name: Liopropoma sp.2 Not applicable

Canonical: Liopropoma

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Liopropoma sp.2 Not applicable","normalized":"Liopropoma","canonical":{"stemmed":"Liopropoma","simple":"Liopropoma","full":"Liopropoma"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Liopropoma","approximationMarker":"sp.","ignored":"2 Not applicable"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"approximationMarker","start":11,"end":14}],"id":"fb3779a4-57a0-5628-8c4e-e341ca4f952d","parserVersion":"test_version"}
```

Name: Lacanobia sp. nr. subjuncta Bold:Aab, 0925

Canonical: Lacanobia

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Lacanobia sp. nr. subjuncta Bold:Aab, 0925","normalized":"Lacanobia","canonical":{"stemmed":"Lacanobia","simple":"Lacanobia","full":"Lacanobia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Lacanobia","approximationMarker":"sp. nr.","ignored":" subjuncta Bold:Aab, 0925"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"approximationMarker","start":10,"end":17}],"id":"05b25429-cb9e-54a1-8e1a-bac9a26d5f46","parserVersion":"test_version"}
```

Name: Lacanobia nr. subjuncta Bold:Aab, 0925

Canonical: Lacanobia

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Lacanobia nr. subjuncta Bold:Aab, 0925","normalized":"Lacanobia","canonical":{"stemmed":"Lacanobia","simple":"Lacanobia","full":"Lacanobia"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Lacanobia","approximationMarker":"nr.","ignored":" subjuncta Bold:Aab, 0925"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"approximationMarker","start":10,"end":13}],"id":"31763a26-a69b-5af8-8703-5da372bdf895","parserVersion":"test_version"}
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
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Abturia cf. alabamensis (Morton )","normalized":"Abturia cf. alabamensis (Morton)","canonical":{"stemmed":"Abturia alabamens","simple":"Abturia alabamensis","full":"Abturia alabamensis"},"cardinality":2,"authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"surrogate":"COMPARISON","details":{"comparison":{"genus":"Abturia","species":"alabamensis (Morton)","authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"comparisonMarker":"cf."}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"comparisonMarker","start":8,"end":11},{"wordType":"specificEpithet","start":12,"end":23},{"wordType":"authorWord","start":25,"end":31}],"id":"5fd4ce59-98d3-50af-9e28-918adc47d264","parserVersion":"test_version"}
```

Name: Abturia cf alabamensis (Morton )

Canonical: Abturia alabamensis

Authorship: (Morton)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Abturia cf alabamensis (Morton )","normalized":"Abturia cf alabamensis (Morton)","canonical":{"stemmed":"Abturia alabamens","simple":"Abturia alabamensis","full":"Abturia alabamensis"},"cardinality":2,"authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"surrogate":"COMPARISON","details":{"comparison":{"genus":"Abturia","species":"alabamensis (Morton)","authorship":{"verbatim":"(Morton )","normalized":"(Morton)","authors":["Morton"],"originalAuth":{"authors":["Morton"]}},"comparisonMarker":"cf"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"comparisonMarker","start":8,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"authorWord","start":24,"end":30}],"id":"423cd26d-c6fd-54fb-937b-f98ba8056fc0","parserVersion":"test_version"}
```

<!--TODO Larus occidentalis cf. wymani|{}-->

Name: Calidris cf. cooperi

Canonical: Calidris cooperi

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Calidris cf. cooperi","normalized":"Calidris cf. cooperi","canonical":{"stemmed":"Calidris cooper","simple":"Calidris cooperi","full":"Calidris cooperi"},"cardinality":2,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Calidris","species":"cooperi","comparisonMarker":"cf."}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"comparisonMarker","start":9,"end":12},{"wordType":"specificEpithet","start":13,"end":20}],"id":"bb19b56e-462f-5daf-a1aa-d4ead082f321","parserVersion":"test_version"}
```

<!--TODO merge comparison with species, infraspecies nodes instead of its own node-->
Name: Aesculus cf. × hybrida

Canonical: Aesculus × hybrida

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"Aesculus cf. × hybrida","normalized":"Aesculus × hybrida","canonical":{"stemmed":"Aesculus hybrid","simple":"Aesculus hybrida","full":"Aesculus × hybrida"},"cardinality":2,"hybrid":"NAMED_HYBRID","surrogate":"COMPARISON","details":{"species":{"genus":"Aesculus","species":"hybrida"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"comparisonMarker","start":9,"end":12},{"wordType":"hybridChar","start":13,"end":14},{"wordType":"specificEpithet","start":15,"end":22}],"id":"6e255814-1c53-54f0-8536-fee957312e9a","parserVersion":"test_version"}
```

<!-- TODO missing subgenus info -->
Name: Daphnia (Daphnia) x krausi Flossner 1993

Canonical: Daphnia × krausi

Authorship: Flossner 1993

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Named hybrid"}],"verbatim":"Daphnia (Daphnia) x krausi Flossner 1993","normalized":"Daphnia × krausi Flossner 1993","canonical":{"stemmed":"Daphnia kraus","simple":"Daphnia krausi","full":"Daphnia × krausi"},"cardinality":2,"authorship":{"verbatim":"Flossner 1993","normalized":"Flossner 1993","year":"1993","authors":["Flossner"],"originalAuth":{"authors":["Flossner"],"year":{"year":"1993"}}},"hybrid":"NAMED_HYBRID","details":{"species":{"genus":"Daphnia","species":"krausi Flossner 1993","authorship":{"verbatim":"Flossner 1993","normalized":"Flossner 1993","year":"1993","authors":["Flossner"],"originalAuth":{"authors":["Flossner"],"year":{"year":"1993"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"hybridChar","start":18,"end":19},{"wordType":"specificEpithet","start":20,"end":26},{"wordType":"authorWord","start":27,"end":35},{"wordType":"year","start":36,"end":40}],"id":"b509d1f1-ce1d-56a1-a15e-2aa9430dce0e","parserVersion":"test_version"}
```

<!--TODO incorrect interpretation-->
Name: Barbus cf macrotaenia × toppini

Canonical: Barbus macrotaenia × Barbus toppini

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Incomplete hybrid formula"},{"quality":4,"warning":"Name comparison"},{"quality":2,"warning":"Hybrid formula"}],"verbatim":"Barbus cf macrotaenia × toppini","normalized":"Barbus cf macrotaenia × Barbus toppini","canonical":{"stemmed":"Barbus macrotaen × Barb toppin","simple":"Barbus macrotaenia × Barbus toppini","full":"Barbus macrotaenia × Barbus toppini"},"cardinality":0,"hybrid":"HYBRID_FORMULA","surrogate":"COMPARISON","details":{"hybridFormula":[{"comparison":{"genus":"Barbus","species":"macrotaenia","comparisonMarker":"cf"}},{"species":{"genus":"Barbus","species":"toppini"}}]},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"comparisonMarker","start":7,"end":9},{"wordType":"specificEpithet","start":10,"end":21},{"wordType":"hybridChar","start":22,"end":23},{"wordType":"specificEpithet","start":24,"end":31}],"id":"37b0b404-d5d9-5699-bbb2-8c3d9bf543a3","parserVersion":"test_version"}
```

Name: Gemmula cf. cosmoi NP-2008

Canonical: Gemmula cosmoi

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":4,"warning":"Name comparison"}],"verbatim":"Gemmula cf. cosmoi NP-2008","normalized":"Gemmula cf. cosmoi","canonical":{"stemmed":"Gemmula cosmo","simple":"Gemmula cosmoi","full":"Gemmula cosmoi"},"cardinality":2,"surrogate":"COMPARISON","tail":" NP-2008","details":{"comparison":{"genus":"Gemmula","species":"cosmoi","comparisonMarker":"cf."}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"comparisonMarker","start":8,"end":11},{"wordType":"specificEpithet","start":12,"end":18}],"id":"87a593b3-2383-5f1b-8772-85e0a4a31b79","parserVersion":"test_version"}
```

### Surrogate Name-Strings

Name: Coleoptera sp. BOLD:AAV0432

Canonical: Coleoptera

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Coleoptera sp. BOLD:AAV0432","normalized":"Coleoptera","canonical":{"stemmed":"Coleoptera","simple":"Coleoptera","full":"Coleoptera"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Coleoptera","approximationMarker":"sp.","ignored":" BOLD:AAV0432"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"approximationMarker","start":11,"end":14}],"id":"65b09adc-12a0-5fbb-a885-75200eacb98a","parserVersion":"test_version"}
```

Name: Coleoptera Bold:AAV0432

Canonical: Coleoptera

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Coleoptera Bold:AAV0432","normalized":"Coleoptera","canonical":{"stemmed":"Coleoptera","simple":"Coleoptera","full":"Coleoptera"},"cardinality":0,"surrogate":"BOLD_SURROGATE","tail":" Bold:AAV0432","details":{"uninomial":{"uninomial":"Coleoptera"}},"pos":[{"wordType":"uninomial","start":0,"end":10}],"id":"9b3865ee-dcf6-5861-9910-58d9f3eafbb1","parserVersion":"test_version"}
```

### Virus-like "normal" names

Name: Ceylonesmus vector Chamberlin, 1941

Canonical: Ceylonesmus vector

Authorship: Chamberlin 1941

```json
{"parsed":true,"parseQuality":1,"verbatim":"Ceylonesmus vector Chamberlin, 1941","normalized":"Ceylonesmus vector Chamberlin 1941","canonical":{"stemmed":"Ceylonesmus uector","simple":"Ceylonesmus vector","full":"Ceylonesmus vector"},"cardinality":2,"authorship":{"verbatim":"Chamberlin, 1941","normalized":"Chamberlin 1941","year":"1941","authors":["Chamberlin"],"originalAuth":{"authors":["Chamberlin"],"year":{"year":"1941"}}},"details":{"species":{"genus":"Ceylonesmus","species":"vector","authorship":{"verbatim":"Chamberlin, 1941","normalized":"Chamberlin 1941","year":"1941","authors":["Chamberlin"],"originalAuth":{"authors":["Chamberlin"],"year":{"year":"1941"}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":18},{"wordType":"authorWord","start":19,"end":29},{"wordType":"year","start":31,"end":35}],"id":"00b874b9-c9ac-5b8a-9821-0a641ca26ca0","parserVersion":"test_version"}
```

### Viruses, plasmids, prions etc.

Name: Arv1virus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Arv1virus","cardinality":0,"virus":true,"id":"25c7c012-6600-5073-8e8f-81fbcf841a66","parserVersion":"test_version"}
```

Name: Turtle herpesviruses

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Turtle herpesviruses","cardinality":0,"virus":true,"id":"44dc4404-0bb8-5eaa-b401-1609d98d3b30","parserVersion":"test_version"}
```

Name: Cre expression vector

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Cre expression vector","cardinality":0,"virus":true,"id":"9a282683-c49b-52dc-817f-0281d5b4b831","parserVersion":"test_version"}
```

Name: Drosophila sturtevanti rhabdovirus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Drosophila sturtevanti rhabdovirus","cardinality":0,"virus":true,"id":"d3510f21-1d57-50e6-98bd-2252259b7052","parserVersion":"test_version"}
```

Name: Hydra expression vector

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Hydra expression vector","cardinality":0,"virus":true,"id":"b22ca1ca-3186-5bc6-9f1a-57ef8c117f25","parserVersion":"test_version"}
```

Name: Gateway destination plasmid

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Gateway destination plasmid","cardinality":0,"id":"21946de0-1c80-543f-ab96-97b81f8d1516","parserVersion":"test_version"}
```

Name: Abutilon mosaic virus [X15983] [X15984] Abutilon mosaic virus ICTV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Abutilon mosaic virus [X15983] [X15984] Abutilon mosaic virus ICTV","cardinality":0,"virus":true,"id":"879da2ea-836c-5ad2-b837-81594a1a208d","parserVersion":"test_version"}
```

Name: Omphalotus sp. Ictv Garcia, 18224

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Omphalotus sp. Ictv Garcia, 18224","cardinality":0,"virus":true,"id":"771a4266-44e3-56d9-9961-9e8a1f1b3936","parserVersion":"test_version"}
```

Name: Acute bee paralysis virus [AF150629] Acute bee paralysis virus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Acute bee paralysis virus [AF150629] Acute bee paralysis virus","cardinality":0,"virus":true,"id":"584822dc-f68f-5abf-aeef-0265172195bf","parserVersion":"test_version"}
```

Name: Adeno-associated virus - 3

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Adeno-associated virus - 3","cardinality":0,"virus":true,"id":"5b16c811-0518-5073-a0be-b59f5faa09fb","parserVersion":"test_version"}
```

Name: ?M1-like Viruses Methanobrevibacter phage PG

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"?M1-like Viruses Methanobrevibacter phage PG","cardinality":0,"virus":true,"id":"b33d05e9-f2a6-5d1b-97e5-3ae061dcd036","parserVersion":"test_version"}
```

Name: Aeromonas phage 65

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Aeromonas phage 65","cardinality":0,"virus":true,"id":"2aef2420-ba68-5887-821f-0ec6eca86660","parserVersion":"test_version"}
```

Name: Bacillus phage SPß [AF020713] Bacillus phage SPb ICTV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Bacillus phage SPß [AF020713] Bacillus phage SPb ICTV","cardinality":0,"virus":true,"id":"ad2b6943-6a54-576d-85e9-e1f8f6aa95db","parserVersion":"test_version"}
```

Name: Apple scar skin viroid

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Apple scar skin viroid","cardinality":0,"virus":true,"id":"7ade78b4-f576-5103-b4a8-4fb9e68845cd","parserVersion":"test_version"}
```

Name: Australian grapevine viroid [X17101] Australian grapevine viroid ICTV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Australian grapevine viroid [X17101] Australian grapevine viroid ICTV","cardinality":0,"virus":true,"id":"381b6868-5d9e-54ec-bae8-84fcc9a3e80c","parserVersion":"test_version"}
```

Name: Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease","cardinality":0,"virus":true,"id":"06193aa6-f2ec-5134-8117-89102448a13e","parserVersion":"test_version"}
```

Name: Phi h-like viruses

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Phi h-like viruses","cardinality":0,"virus":true,"id":"474acd56-6be4-56fc-9045-48a3d570ac97","parserVersion":"test_version"}
```

Name: Viroids

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Viroids","cardinality":0,"virus":true,"id":"641d47bf-c7c4-5218-8e2e-8756ad808653","parserVersion":"test_version"}
```

Name: Fungal prions

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Fungal prions","cardinality":0,"virus":true,"id":"ec273e2d-cdde-5fcb-84dc-a6adf2e309ce","parserVersion":"test_version"}
```

Name: Human rhinovirus A11

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Human rhinovirus A11","cardinality":0,"virus":true,"id":"ba205a7c-1c63-51c7-8f4d-d47665f56c33","parserVersion":"test_version"}
```

Name: Kobuvirus korean black goat/South Korea/2010

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Kobuvirus korean black goat/South Korea/2010","cardinality":0,"virus":true,"id":"4871667d-e362-5f76-a218-6c1bcc090ba9","parserVersion":"test_version"}
```

Name: Australian bat lyssavirus human/AUS/1998

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Australian bat lyssavirus human/AUS/1998","cardinality":0,"virus":true,"id":"5e4fdc2a-3fb3-5776-b94d-04b9f0c6fcbb","parserVersion":"test_version"}
```

Name: Gossypium mustilinum symptomless alphasatellite

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Gossypium mustilinum symptomless alphasatellite","cardinality":0,"virus":true,"id":"d8b1e803-34ba-537b-874b-48521afb92a5","parserVersion":"test_version"}
```

Name: Okra leaf curl Mali alphasatellites-Cameroon

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Okra leaf curl Mali alphasatellites-Cameroon","cardinality":0,"virus":true,"id":"034731b5-3de7-5d48-bf3b-f89272699a45","parserVersion":"test_version"}
```

Name: Bemisia betasatellite LW-2014

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Bemisia betasatellite LW-2014","cardinality":0,"virus":true,"id":"21d06e45-a312-5844-88f7-3eb0b73d1efc","parserVersion":"test_version"}
```

Name: Tomato leaf curl Bangladesh betasatellites [India/Patna/Chilli/2008]

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Tomato leaf curl Bangladesh betasatellites [India/Patna/Chilli/2008]","cardinality":0,"virus":true,"id":"c5def37b-c5d9-57e4-822a-0436629f5d99","parserVersion":"test_version"}
```

Name: Intracisternal A-particles

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Intracisternal A-particles","cardinality":0,"virus":true,"id":"4f16a692-534b-5ec5-87f4-58fe76a0ed9d","parserVersion":"test_version"}
```

Name: Saccharomyces cerevisiae killer particle M1

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Saccharomyces cerevisiae killer particle M1","cardinality":0,"virus":true,"id":"879050a7-5085-5679-85e4-fe47308843dd","parserVersion":"test_version"}
```

Name: Uranotaenia sapphirina NPV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Uranotaenia sapphirina NPV","cardinality":0,"virus":true,"id":"83886b77-a81a-52ba-9b0e-5743b4242b97","parserVersion":"test_version"}
```

Name: Uranotaenia sapphirina Npv

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Uranotaenia sapphirina Npv","cardinality":0,"virus":true,"id":"917cfcbc-3a38-5f59-affc-56c87f04a7ec","parserVersion":"test_version"}
```

Name: Spodoptera exigua nuclear polyhedrosis virus SeMNPV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Spodoptera exigua nuclear polyhedrosis virus SeMNPV","cardinality":0,"virus":true,"id":"a0356512-17eb-51ab-92b3-21d92393b84c","parserVersion":"test_version"}
```

Name: Spodoptera frugiperda MNPV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Spodoptera frugiperda MNPV","cardinality":0,"virus":true,"id":"5a694933-6187-54bb-ae35-77ed3384b69d","parserVersion":"test_version"}
```

Name: Rachiplusia ou MNPV (strain R1)

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Rachiplusia ou MNPV (strain R1)","cardinality":0,"virus":true,"id":"ca77e2a5-fa26-5c7f-bf68-a449c32ea95e","parserVersion":"test_version"}
```

Name: Orgyia pseudotsugata nuclear polyhedrosis virus OpMNPV

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Orgyia pseudotsugata nuclear polyhedrosis virus OpMNPV","cardinality":0,"virus":true,"id":"f3b4269c-a97f-5ff7-bb4a-56d982b3707c","parserVersion":"test_version"}
```

Name: Mamestra configurata NPV-A

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Mamestra configurata NPV-A","cardinality":0,"virus":true,"id":"59160819-f61d-5360-85c5-78b6140a05ca","parserVersion":"test_version"}
```

Name: Helicoverpa armigera SNPV NNg1

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Helicoverpa armigera SNPV NNg1","cardinality":0,"virus":true,"id":"933f0a27-1fd8-5066-90ee-df1ed8148c9c","parserVersion":"test_version"}
```

Name: Zamilon virophage

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Zamilon virophage","cardinality":0,"virus":true,"id":"661132c0-7012-5405-bfc7-31e9a4b3946c","parserVersion":"test_version"}
```

Name: Sputnik virophage 3

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Sputnik virophage 3","cardinality":0,"virus":true,"id":"b206bb35-01bf-59a7-8dad-bc8f99ca0a2a","parserVersion":"test_version"}
```

Name: Bacteriophage PH75

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Bacteriophage PH75","cardinality":0,"virus":true,"id":"605f428e-a4a3-57a2-9dfa-a6a3d99b801d","parserVersion":"test_version"}
```

Name: Escherichia coli bacteriophage

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Escherichia coli bacteriophage","cardinality":0,"virus":true,"id":"c01315c2-e1cc-58c2-b113-2d756985d64b","parserVersion":"test_version"}
```

Name: Betasatellites

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Betasatellites","cardinality":0,"virus":true,"id":"1a6aa729-5fc5-5fbd-9299-efb9a6198310","parserVersion":"test_version"}
```

Name: Satellite Nucleic Acids (Subviral DNA-ssDNA)

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Satellite Nucleic Acids (Subviral DNA-ssDNA)","cardinality":0,"virus":true,"id":"1a769ed9-62cd-54b9-9c94-36d99117b89f","parserVersion":"test_version"}
```

### Name-strings with RNA

Name: ssRNA

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"ssRNA","cardinality":0,"id":"10d5f30c-e51b-54ed-be43-c0ac1656a88a","parserVersion":"test_version"}
```

Name: Alpha proteobacterium RNA12

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Alpha proteobacterium RNA12","cardinality":0,"id":"c2826f30-f6f3-543f-80cf-646adf374a59","parserVersion":"test_version"}
```

Name: Ustilaginoidea virens RNA virus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Ustilaginoidea virens RNA virus","cardinality":0,"virus":true,"id":"61fff10f-7f16-5f42-b642-ba0195abccb8","parserVersion":"test_version"}
```

Name: Candida albicans RNA_CTR0-3

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Candida albicans RNA_CTR0-3","cardinality":0,"id":"0182d44b-5d8b-501d-8f5c-4ef44dff8db4","parserVersion":"test_version"}
```

Name: Carabus satyrus satyrus KURNAKOV, 1962

Canonical: Carabus satyrus satyrus

Authorship: Kurnakov 1962

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"}],"verbatim":"Carabus satyrus satyrus KURNAKOV, 1962","normalized":"Carabus satyrus satyrus Kurnakov 1962","canonical":{"stemmed":"Carabus satyr satyr","simple":"Carabus satyrus satyrus","full":"Carabus satyrus satyrus"},"cardinality":3,"authorship":{"verbatim":"KURNAKOV, 1962","normalized":"Kurnakov 1962","year":"1962","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1962"}}},"details":{"infraSpecies":{"genus":"Carabus","species":"satyrus","infraSpecies":[{"value":"satyrus","authorship":{"verbatim":"KURNAKOV, 1962","normalized":"Kurnakov 1962","year":"1962","authors":["Kurnakov"],"originalAuth":{"authors":["Kurnakov"],"year":{"year":"1962"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":15},{"wordType":"infraspecificEpithet","start":16,"end":23},{"wordType":"authorWord","start":24,"end":32},{"wordType":"year","start":34,"end":38}],"id":"81654954-0f47-5715-acb1-1cd8d2c49e9a","parserVersion":"test_version"}
```


### Epithet prioni is not a prion

Name: Fakus prioni

Canonical: Fakus prioni

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Fakus prioni","normalized":"Fakus prioni","canonical":{"stemmed":"Fakus prion","simple":"Fakus prioni","full":"Fakus prioni"},"cardinality":2,"details":{"species":{"genus":"Fakus","species":"prioni"}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":12}],"id":"f2561b5b-37ed-592d-9c12-4ef96d09f554","parserVersion":"test_version"}
```

### Names with "satellite" as a substring

Name: Crassatellites fulvida

Canonical: Crassatellites fulvida

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Crassatellites fulvida","normalized":"Crassatellites fulvida","canonical":{"stemmed":"Crassatellites fuluid","simple":"Crassatellites fulvida","full":"Crassatellites fulvida"},"cardinality":2,"details":{"species":{"genus":"Crassatellites","species":"fulvida"}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22}],"id":"089171ac-f672-5973-950a-9419651e6b0e","parserVersion":"test_version"}
```

### Bacterial genus

Name: Salmonella werahensis (Castellani) Hauduroy and Ehringer in Hauduroy 1937

Canonical: Salmonella werahensis

Authorship: (Castellani) Hauduroy & Ehringer ex Hauduroy 1937

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Salmonella werahensis (Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"Salmonella werahensis (Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","canonical":{"stemmed":"Salmonella werahens","simple":"Salmonella werahensis","full":"Salmonella werahensis"},"cardinality":2,"authorship":{"verbatim":"(Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"(Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","authors":["Castellani","Hauduroy","Ehringer"],"originalAuth":{"authors":["Castellani"]},"combinationAuth":{"authors":["Hauduroy","Ehringer"],"exAuthors":{"authors":["Hauduroy"],"year":{"year":"1937"}}}},"bacteria":"yes","details":{"species":{"genus":"Salmonella","species":"werahensis","authorship":{"verbatim":"(Castellani) Hauduroy and Ehringer in Hauduroy 1937","normalized":"(Castellani) Hauduroy \u0026 Ehringer ex Hauduroy 1937","authors":["Castellani","Hauduroy","Ehringer"],"originalAuth":{"authors":["Castellani"]},"combinationAuth":{"authors":["Hauduroy","Ehringer"],"exAuthors":{"authors":["Hauduroy"],"year":{"year":"1937"}}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21},{"wordType":"authorWord","start":23,"end":33},{"wordType":"authorWord","start":35,"end":43},{"wordType":"authorWord","start":48,"end":56},{"wordType":"authorWord","start":60,"end":68},{"wordType":"year","start":69,"end":73}],"id":"bb6e2a9f-6813-5b00-9a3f-e12a085e515e","parserVersion":"test_version"}
```

### Bacteria genus homonym

Name: Actinomyces cardiffensis

Canonical: Actinomyces cardiffensis

Authorship:

```json
{"parsed":true,"parseQuality":1,"qualityWarnings":[{"quality":1,"warning":"The genus is a homonym of a bacterial genus"}],"verbatim":"Actinomyces cardiffensis","normalized":"Actinomyces cardiffensis","canonical":{"stemmed":"Actinomyces cardiffens","simple":"Actinomyces cardiffensis","full":"Actinomyces cardiffensis"},"cardinality":2,"bacteria":"maybe","details":{"species":{"genus":"Actinomyces","species":"cardiffensis"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":24}],"id":"fc1def53-81ba-5d2f-9f4c-0d9ac591cd13","parserVersion":"test_version"}
```

### Bacteria with pathovar rank

Name: Xanthomonas axonopodis pv. phaseoli

Canonical: Xanthomonas axonopodis pv. phaseoli

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Xanthomonas axonopodis pv. phaseoli","normalized":"Xanthomonas axonopodis pv. phaseoli","canonical":{"stemmed":"Xanthomonas axonopod phaseol","simple":"Xanthomonas axonopodis phaseoli","full":"Xanthomonas axonopodis pv. phaseoli"},"cardinality":3,"bacteria":"yes","details":{"infraSpecies":{"genus":"Xanthomonas","species":"axonopodis","infraSpecies":[{"value":"phaseoli","rank":"pv."}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22},{"wordType":"rank","start":23,"end":26},{"wordType":"infraspecificEpithet","start":27,"end":35}],"id":"ea35594e-41c7-5706-b3b8-bb1b94d11a77","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pathovar. phaseoli

Canonical: Xanthomonas axonopodis pathovar. phaseoli

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Xanthomonas axonopodis pathovar. phaseoli","normalized":"Xanthomonas axonopodis pathovar. phaseoli","canonical":{"stemmed":"Xanthomonas axonopod phaseol","simple":"Xanthomonas axonopodis phaseoli","full":"Xanthomonas axonopodis pathovar. phaseoli"},"cardinality":3,"bacteria":"yes","details":{"infraSpecies":{"genus":"Xanthomonas","species":"axonopodis","infraSpecies":[{"value":"phaseoli","rank":"pathovar."}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22},{"wordType":"rank","start":23,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":41}],"id":"816ce2bc-4cdc-59ab-8900-e4414e8d2125","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pathovar.

Canonical: Xanthomonas axonopodis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Xanthomonas axonopodis pathovar.","normalized":"Xanthomonas axonopodis","canonical":{"stemmed":"Xanthomonas axonopod","simple":"Xanthomonas axonopodis","full":"Xanthomonas axonopodis"},"cardinality":2,"bacteria":"yes","tail":" pathovar.","details":{"species":{"genus":"Xanthomonas","species":"axonopodis"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22}],"id":"851a86de-df67-5fba-b3f7-73937a5edbce","parserVersion":"test_version"}
```

Name: Xanthomonas axonopodis pv.

Canonical: Xanthomonas axonopodis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Xanthomonas axonopodis pv.","normalized":"Xanthomonas axonopodis","canonical":{"stemmed":"Xanthomonas axonopod","simple":"Xanthomonas axonopodis","full":"Xanthomonas axonopodis"},"cardinality":2,"bacteria":"yes","tail":" pv.","details":{"species":{"genus":"Xanthomonas","species":"axonopodis"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22}],"id":"0c0ce6dd-e5ea-5c17-8be3-c381ff662f12","parserVersion":"test_version"}
```

### "Stray" ex is not parsed as species

Name: Pelargonium cucullatum ssp. cucullatum (L.) L'Her. ex [Soland.]

Canonical: Pelargonium cucullatum subsp. cucullatum

Authorship: (L.) L'Her.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pelargonium cucullatum ssp. cucullatum (L.) L'Her. ex [Soland.]","normalized":"Pelargonium cucullatum subsp. cucullatum (L.) L'Her.","canonical":{"stemmed":"Pelargonium cucullat cucullat","simple":"Pelargonium cucullatum cucullatum","full":"Pelargonium cucullatum subsp. cucullatum"},"cardinality":3,"authorship":{"verbatim":"(L.) L'Her.","normalized":"(L.) L'Her.","authors":["L.","L'Her."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Her."]}},"tail":" ex [Soland.]","details":{"infraSpecies":{"genus":"Pelargonium","species":"cucullatum","infraSpecies":[{"value":"cucullatum","rank":"subsp.","authorship":{"verbatim":"(L.) L'Her.","normalized":"(L.) L'Her.","authors":["L.","L'Her."],"originalAuth":{"authors":["L."]},"combinationAuth":{"authors":["L'Her."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22},{"wordType":"rank","start":23,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":38},{"wordType":"authorWord","start":40,"end":42},{"wordType":"authorWord","start":44,"end":50}],"id":"83811b74-a581-5801-aa49-d4eab6775fdb","parserVersion":"test_version"}
```

<!-- not dealing with ex. gr for now -->
Name: Acastella ex gr. rouaulti

Canonical: Acastella

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acastella ex gr. rouaulti","normalized":"Acastella","canonical":{"stemmed":"Acastella","simple":"Acastella","full":"Acastella"},"cardinality":1,"tail":" ex gr. rouaulti","details":{"uninomial":{"uninomial":"Acastella"}},"pos":[{"wordType":"uninomial","start":0,"end":9}],"id":"c1864b52-848a-5de7-8f2d-a3cfe2025c40","parserVersion":"test_version"}
```

### Authoship in upper case

Name: Lecanora strobilinoides GIRALT & GÓMEZ-BOLEA

Canonical: Lecanora strobilinoides

Authorship: Giralt & Gómez-Bolea

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Author in upper case"}],"verbatim":"Lecanora strobilinoides GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Lecanora strobilinoides Giralt \u0026 Gómez-Bolea","canonical":{"stemmed":"Lecanora strobilinoid","simple":"Lecanora strobilinoides","full":"Lecanora strobilinoides"},"cardinality":2,"authorship":{"verbatim":"GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Giralt \u0026 Gómez-Bolea","authors":["Giralt","Gómez-Bolea"],"originalAuth":{"authors":["Giralt","Gómez-Bolea"]}},"details":{"species":{"genus":"Lecanora","species":"strobilinoides","authorship":{"verbatim":"GIRALT \u0026 GÓMEZ-BOLEA","normalized":"Giralt \u0026 Gómez-Bolea","authors":["Giralt","Gómez-Bolea"],"originalAuth":{"authors":["Giralt","Gómez-Bolea"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":23},{"wordType":"authorWord","start":24,"end":30},{"wordType":"authorWord","start":33,"end":44}],"id":"f2bfaa25-c25f-5a31-90c6-a19bd4dc23f4","parserVersion":"test_version"}
```

### Numbers and letters separated with '-' are not parsed as authors

Name: Astatotilapia cf. bloyeti OS-2017

Canonical: Astatotilapia bloyeti

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":4,"warning":"Name comparison"}],"verbatim":"Astatotilapia cf. bloyeti OS-2017","normalized":"Astatotilapia cf. bloyeti","canonical":{"stemmed":"Astatotilapia bloyet","simple":"Astatotilapia bloyeti","full":"Astatotilapia bloyeti"},"cardinality":2,"surrogate":"COMPARISON","tail":" OS-2017","details":{"comparison":{"genus":"Astatotilapia","species":"bloyeti","comparisonMarker":"cf."}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"comparisonMarker","start":14,"end":17},{"wordType":"specificEpithet","start":18,"end":25}],"id":"c841aa1d-78ea-5b6a-93fc-e18c54164144","parserVersion":"test_version"}
```

### Double parenthesis
Name: Eichornia crassipes ( (Martius) ) Solms-Laub.

Canonical: Eichornia crassipes

Authorship: (Martius) Solms-Laub.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Authorship in double parentheses"}],"verbatim":"Eichornia crassipes ( (Martius) ) Solms-Laub.","normalized":"Eichornia crassipes (Martius) Solms-Laub.","canonical":{"stemmed":"Eichornia crassip","simple":"Eichornia crassipes","full":"Eichornia crassipes"},"cardinality":2,"authorship":{"verbatim":"( (Martius) ) Solms-Laub.","normalized":"(Martius) Solms-Laub.","authors":["Martius","Solms-Laub."],"originalAuth":{"authors":["Martius"]},"combinationAuth":{"authors":["Solms-Laub."]}},"details":{"species":{"genus":"Eichornia","species":"crassipes","authorship":{"verbatim":"( (Martius) ) Solms-Laub.","normalized":"(Martius) Solms-Laub.","authors":["Martius","Solms-Laub."],"originalAuth":{"authors":["Martius"]},"combinationAuth":{"authors":["Solms-Laub."]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":19},{"wordType":"authorWord","start":23,"end":30},{"wordType":"authorWord","start":34,"end":45}],"id":"95b90189-29d1-51ca-a1fa-0fb1c19a1fa1","parserVersion":"test_version"}
```

### Numbers at the start/middle of names

Name: Nesomyrmex madecassus_01m

Canonical: Nesomyrmex

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nesomyrmex madecassus_01m","normalized":"Nesomyrmex","canonical":{"stemmed":"Nesomyrmex","simple":"Nesomyrmex","full":"Nesomyrmex"},"cardinality":1,"tail":" madecassus_01m","details":{"uninomial":{"uninomial":"Nesomyrmex"}},"pos":[{"wordType":"uninomial","start":0,"end":10}],"id":"30dd0028-1ad4-5f65-ba5e-3df4963825d2","parserVersion":"test_version"}
```

Name: Hypochrys0des

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Hypochrys0des","cardinality":0,"id":"859c6279-20ea-5e60-9b7d-0c5283e06377","parserVersion":"test_version"}
```

Name: Hypochrys0des Leraut 1981

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Hypochrys0des Leraut 1981","cardinality":0,"id":"c053bbbf-de6c-5b22-a0f9-0803093b9b2d","parserVersion":"test_version"}
```

Name: Phyllodoce mucosa 0ersted, 1843

Canonical: Phyllodoce mucosa

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Phyllodoce mucosa 0ersted, 1843","normalized":"Phyllodoce mucosa","canonical":{"stemmed":"Phyllodoce mucos","simple":"Phyllodoce mucosa","full":"Phyllodoce mucosa"},"cardinality":2,"tail":" 0ersted, 1843","details":{"species":{"genus":"Phyllodoce","species":"mucosa"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":17}],"id":"52695b7b-ebef-5624-9ccf-f9d07cd8133c","parserVersion":"test_version"}
```

Name: Attelabus 0l.

Canonical: Attelabus

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Attelabus 0l.","normalized":"Attelabus","canonical":{"stemmed":"Attelabus","simple":"Attelabus","full":"Attelabus"},"cardinality":1,"tail":" 0l.","details":{"uninomial":{"uninomial":"Attelabus"}},"pos":[{"wordType":"uninomial","start":0,"end":9}],"id":"b9edee54-a7ae-525a-a319-ffeed18cf88a","parserVersion":"test_version"}
```

Name: Acrobothrium 0lsson 1872

Canonical: Acrobothrium

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acrobothrium 0lsson 1872","normalized":"Acrobothrium","canonical":{"stemmed":"Acrobothrium","simple":"Acrobothrium","full":"Acrobothrium"},"cardinality":1,"tail":" 0lsson 1872","details":{"uninomial":{"uninomial":"Acrobothrium"}},"pos":[{"wordType":"uninomial","start":0,"end":12}],"id":"2edfbcca-af28-5498-a762-663e5d5b9f73","parserVersion":"test_version"}
```

Name: Staphylinus haemrrhoidalis 0l. nec Gmel

Canonical: Staphylinus haemrrhoidalis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Staphylinus haemrrhoidalis 0l. nec Gmel","normalized":"Staphylinus haemrrhoidalis","canonical":{"stemmed":"Staphylinus haemrrhoidal","simple":"Staphylinus haemrrhoidalis","full":"Staphylinus haemrrhoidalis"},"cardinality":2,"tail":" 0l. nec Gmel","details":{"species":{"genus":"Staphylinus","species":"haemrrhoidalis"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":26}],"id":"3ef602da-08a5-5acf-8f8a-9c515373ccda","parserVersion":"test_version"}
```

Name: Ea92virus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Ea92virus","cardinality":0,"virus":true,"id":"2465682c-cd5c-5408-859b-8bcc5489125f","parserVersion":"test_version"}
```

### Year without authorship

<!--TODO: collect year information-->
Name: Acarospora cratericola 1929

Canonical: Acarospora cratericola

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola 1929","normalized":"Acarospora cratericola","canonical":{"stemmed":"Acarospora cratericol","simple":"Acarospora cratericola","full":"Acarospora cratericola"},"cardinality":2,"tail":" 1929","details":{"species":{"genus":"Acarospora","species":"cratericola"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22}],"id":"11335046-cf05-5571-84bb-f9c8a4b2d8de","parserVersion":"test_version"}
```

Name: Goggia gemmula 1996

Canonical: Goggia gemmula

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Goggia gemmula 1996","normalized":"Goggia gemmula","canonical":{"stemmed":"Goggia gemmul","simple":"Goggia gemmula","full":"Goggia gemmula"},"cardinality":2,"tail":" 1996","details":{"species":{"genus":"Goggia","species":"gemmula"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":14}],"id":"707ab43c-41bd-56bc-b2aa-96db4913ad35","parserVersion":"test_version"}
```

### Year range

Name: Eurodryas orientalis Herrich-Schäffer 1845-1847

Canonical: Eurodryas orientalis

Authorship: Herrich-Schäffer (1845)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Eurodryas orientalis Herrich-Schäffer 1845-1847","normalized":"Eurodryas orientalis Herrich-Schäffer (1845)","canonical":{"stemmed":"Eurodryas oriental","simple":"Eurodryas orientalis","full":"Eurodryas orientalis"},"cardinality":2,"authorship":{"verbatim":"Herrich-Schäffer 1845-1847","normalized":"Herrich-Schäffer (1845)","year":"(1845)","authors":["Herrich-Schäffer"],"originalAuth":{"authors":["Herrich-Schäffer"],"year":{"year":"1845","isApproximate":true}}},"details":{"species":{"genus":"Eurodryas","species":"orientalis","authorship":{"verbatim":"Herrich-Schäffer 1845-1847","normalized":"Herrich-Schäffer (1845)","year":"(1845)","authors":["Herrich-Schäffer"],"originalAuth":{"authors":["Herrich-Schäffer"],"year":{"year":"1845","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":20},{"wordType":"authorWord","start":21,"end":37},{"wordType":"approximateYear","start":38,"end":42}],"id":"5fbca057-cd1e-5334-b6d3-496559b31818","parserVersion":"test_version"}
```

Name: Tridentella tangeroae Bruce, 1987-92

Canonical: Tridentella tangeroae

Authorship: Bruce (1987)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Tridentella tangeroae Bruce, 1987-92","normalized":"Tridentella tangeroae Bruce (1987)","canonical":{"stemmed":"Tridentella tangero","simple":"Tridentella tangeroae","full":"Tridentella tangeroae"},"cardinality":2,"authorship":{"verbatim":"Bruce, 1987-92","normalized":"Bruce (1987)","year":"(1987)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"1987","isApproximate":true}}},"details":{"species":{"genus":"Tridentella","species":"tangeroae","authorship":{"verbatim":"Bruce, 1987-92","normalized":"Bruce (1987)","year":"(1987)","authors":["Bruce"],"originalAuth":{"authors":["Bruce"],"year":{"year":"1987","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":22,"end":27},{"wordType":"approximateYear","start":29,"end":33}],"id":"6c943756-7f67-51ee-9c06-8f9016538be6","parserVersion":"test_version"}
```

Name: Macroplectra unicolor Moore, 1858/59

Canonical: Macroplectra unicolor

Authorship: Moore (1858)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Macroplectra unicolor Moore, 1858/59","normalized":"Macroplectra unicolor Moore (1858)","canonical":{"stemmed":"Macroplectra unicolor","simple":"Macroplectra unicolor","full":"Macroplectra unicolor"},"cardinality":2,"authorship":{"verbatim":"Moore, 1858/59","normalized":"Moore (1858)","year":"(1858)","authors":["Moore"],"originalAuth":{"authors":["Moore"],"year":{"year":"1858","isApproximate":true}}},"details":{"species":{"genus":"Macroplectra","species":"unicolor","authorship":{"verbatim":"Moore, 1858/59","normalized":"Moore (1858)","year":"(1858)","authors":["Moore"],"originalAuth":{"authors":["Moore"],"year":{"year":"1858","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":21},{"wordType":"authorWord","start":22,"end":27},{"wordType":"approximateYear","start":29,"end":33}],"id":"d6fc4a96-793c-58ce-9926-ec40281062b2","parserVersion":"test_version"}
```

Name: Seryda basirei Druce, 1891/901

Canonical: Seryda basirei

Authorship: Druce (1891)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Years range"}],"verbatim":"Seryda basirei Druce, 1891/901","normalized":"Seryda basirei Druce (1891)","canonical":{"stemmed":"Seryda basire","simple":"Seryda basirei","full":"Seryda basirei"},"cardinality":2,"authorship":{"verbatim":"Druce, 1891/901","normalized":"Druce (1891)","year":"(1891)","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1891","isApproximate":true}}},"details":{"species":{"genus":"Seryda","species":"basirei","authorship":{"verbatim":"Druce, 1891/901","normalized":"Druce (1891)","year":"(1891)","authors":["Druce"],"originalAuth":{"authors":["Druce"],"year":{"year":"1891","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":14},{"wordType":"authorWord","start":15,"end":20},{"wordType":"approximateYear","start":22,"end":26}],"id":"574ff67d-f220-5c14-9634-fcadc3794891","parserVersion":"test_version"}
```

### Year with page number

Name: Recilia truncatus Dash & Viraktamath, 1998a: 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a: 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":22},{"wordType":"authorWord","start":25,"end":36},{"wordType":"year","start":38,"end":43}],"id":"227ada89-45e5-56a9-83ad-47bee641e373","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998: 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998: 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998: 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":22},{"wordType":"authorWord","start":25,"end":36},{"wordType":"year","start":38,"end":42}],"id":"47a39cf1-7be1-5937-b8fa-03a1696c1de6","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998a:29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a:29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a:29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a:29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":22},{"wordType":"authorWord","start":25,"end":36},{"wordType":"year","start":38,"end":43}],"id":"68b51644-5fef-5d5f-819d-f5bf8c9e6051","parserVersion":"test_version"}
```

Name: Recilia truncatus Dash & Viraktamath, 1998a : 29

Canonical: Recilia truncatus

Authorship: Dash & Viraktamath 1998

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Year with latin character"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Recilia truncatus Dash \u0026 Viraktamath, 1998a : 29","normalized":"Recilia truncatus Dash \u0026 Viraktamath 1998","canonical":{"stemmed":"Recilia truncat","simple":"Recilia truncatus","full":"Recilia truncatus"},"cardinality":2,"authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a : 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}},"details":{"species":{"genus":"Recilia","species":"truncatus","authorship":{"verbatim":"Dash \u0026 Viraktamath, 1998a : 29","normalized":"Dash \u0026 Viraktamath 1998","year":"1998","authors":["Dash","Viraktamath"],"originalAuth":{"authors":["Dash","Viraktamath"],"year":{"year":"1998"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":22},{"wordType":"authorWord","start":25,"end":36},{"wordType":"year","start":38,"end":43}],"id":"08507e4f-412c-59c9-b1f2-906dd4b27aa8","parserVersion":"test_version"}
```

### Year in square brackets

Name: Anthoscopus Cabanis [1851]

Canonical: Anthoscopus

Authorship: Cabanis (1851)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"}],"verbatim":"Anthoscopus Cabanis [1851]","normalized":"Anthoscopus Cabanis (1851)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":19},{"wordType":"approximateYear","start":21,"end":25}],"id":"8d86299b-3028-5be2-b2f6-6e4897f4c748","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [185?]

Canonical: Anthoscopus

Authorship: Cabanis (185?)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [185?]","normalized":"Anthoscopus Cabanis (185?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [185?]","normalized":"Cabanis (185?)","year":"(185?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"185?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [185?]","normalized":"Cabanis (185?)","year":"(185?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"185?","isApproximate":true}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":19},{"wordType":"approximateYear","start":21,"end":25}],"id":"3434c072-d015-5f54-ad32-45b01de7fd08","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851?]

Canonical: Anthoscopus

Authorship: Cabanis (1851?)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [1851?]","normalized":"Anthoscopus Cabanis (1851?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":19},{"wordType":"approximateYear","start":21,"end":26}],"id":"6b12b541-b58b-5f11-ba66-bb314b53813f","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851]

Canonical: Anthoscopus

Authorship: Cabanis (1851)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"}],"verbatim":"Anthoscopus Cabanis [1851]","normalized":"Anthoscopus Cabanis (1851)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851]","normalized":"Cabanis (1851)","year":"(1851)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851","isApproximate":true}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":19},{"wordType":"approximateYear","start":21,"end":25}],"id":"8d86299b-3028-5be2-b2f6-6e4897f4c748","parserVersion":"test_version"}
```

Name: Anthoscopus Cabanis [1851?]

Canonical: Anthoscopus

Authorship: Cabanis (1851?)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"},{"quality":2,"warning":"Year with question mark"}],"verbatim":"Anthoscopus Cabanis [1851?]","normalized":"Anthoscopus Cabanis (1851?)","canonical":{"stemmed":"Anthoscopus","simple":"Anthoscopus","full":"Anthoscopus"},"cardinality":1,"authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}},"details":{"uninomial":{"uninomial":"Anthoscopus","authorship":{"verbatim":"Cabanis [1851?]","normalized":"Cabanis (1851?)","year":"(1851?)","authors":["Cabanis"],"originalAuth":{"authors":["Cabanis"],"year":{"year":"1851?","isApproximate":true}}}}},"pos":[{"wordType":"uninomial","start":0,"end":11},{"wordType":"authorWord","start":12,"end":19},{"wordType":"approximateYear","start":21,"end":26}],"id":"6b12b541-b58b-5f11-ba66-bb314b53813f","parserVersion":"test_version"}
```

Name: Trismegistia monodii Ando, 1973 [1974]

Canonical: Trismegistia monodii

Authorship: Ando 1973

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Trismegistia monodii Ando, 1973 [1974]","normalized":"Trismegistia monodii Ando 1973","canonical":{"stemmed":"Trismegistia monodi","simple":"Trismegistia monodii","full":"Trismegistia monodii"},"cardinality":2,"authorship":{"verbatim":"Ando, 1973","normalized":"Ando 1973","year":"1973","authors":["Ando"],"originalAuth":{"authors":["Ando"],"year":{"year":"1973"}}},"tail":" [1974]","details":{"species":{"genus":"Trismegistia","species":"monodii","authorship":{"verbatim":"Ando, 1973","normalized":"Ando 1973","year":"1973","authors":["Ando"],"originalAuth":{"authors":["Ando"],"year":{"year":"1973"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":20},{"wordType":"authorWord","start":21,"end":25},{"wordType":"year","start":27,"end":31}],"id":"f396d2d0-b14e-537f-ae8f-c383310f813e","parserVersion":"test_version"}
```

Name: Zygaena witti Wiegel [1973]

Canonical: Zygaena witti

Authorship: Wiegel (1973)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Year with square brakets"}],"verbatim":"Zygaena witti Wiegel [1973]","normalized":"Zygaena witti Wiegel (1973)","canonical":{"stemmed":"Zygaena witt","simple":"Zygaena witti","full":"Zygaena witti"},"cardinality":2,"authorship":{"verbatim":"Wiegel [1973]","normalized":"Wiegel (1973)","year":"(1973)","authors":["Wiegel"],"originalAuth":{"authors":["Wiegel"],"year":{"year":"1973","isApproximate":true}}},"details":{"species":{"genus":"Zygaena","species":"witti","authorship":{"verbatim":"Wiegel [1973]","normalized":"Wiegel (1973)","year":"(1973)","authors":["Wiegel"],"originalAuth":{"authors":["Wiegel"],"year":{"year":"1973","isApproximate":true}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":13},{"wordType":"authorWord","start":14,"end":20},{"wordType":"approximateYear","start":22,"end":26}],"id":"76eef612-f125-54f9-b241-6b3a9be0a6c6","parserVersion":"test_version"}
```

Name: Deyeuxia coarctata Kunth, 1815 [1816]

Canonical: Deyeuxia coarctata

Authorship: Kunth 1815

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Deyeuxia coarctata Kunth, 1815 [1816]","normalized":"Deyeuxia coarctata Kunth 1815","canonical":{"stemmed":"Deyeuxia coarctat","simple":"Deyeuxia coarctata","full":"Deyeuxia coarctata"},"cardinality":2,"authorship":{"verbatim":"Kunth, 1815","normalized":"Kunth 1815","year":"1815","authors":["Kunth"],"originalAuth":{"authors":["Kunth"],"year":{"year":"1815"}}},"tail":" [1816]","details":{"species":{"genus":"Deyeuxia","species":"coarctata","authorship":{"verbatim":"Kunth, 1815","normalized":"Kunth 1815","year":"1815","authors":["Kunth"],"originalAuth":{"authors":["Kunth"],"year":{"year":"1815"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":19,"end":24},{"wordType":"year","start":26,"end":30}],"id":"2f479365-40be-5181-b194-8a24fc743f73","parserVersion":"test_version"}
```

### Names with broken conversion between encodings

Name: Macrotes cordovaria Guen�e 1857

Canonical: Macrotes cordovaria

Authorship: Guen�e 1857

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Incorrect conversion to UTF-8"}],"verbatim":"Macrotes cordovaria Guen�e 1857","normalized":"Macrotes cordovaria Guen�e 1857","canonical":{"stemmed":"Macrotes cordouar","simple":"Macrotes cordovaria","full":"Macrotes cordovaria"},"cardinality":2,"authorship":{"verbatim":"Guen�e 1857","normalized":"Guen�e 1857","year":"1857","authors":["Guen�e"],"originalAuth":{"authors":["Guen�e"],"year":{"year":"1857"}}},"details":{"species":{"genus":"Macrotes","species":"cordovaria","authorship":{"verbatim":"Guen�e 1857","normalized":"Guen�e 1857","year":"1857","authors":["Guen�e"],"originalAuth":{"authors":["Guen�e"],"year":{"year":"1857"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"authorWord","start":20,"end":26},{"wordType":"year","start":27,"end":31}],"id":"9217d59c-d1e7-5c79-af65-f52623446c15","parserVersion":"test_version"}
```

Name: Fusinus eucos�nius

Canonical: Fusinus eucos�nius

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Incorrect conversion to UTF-8"}],"verbatim":"Fusinus eucos�nius","normalized":"Fusinus eucos�nius","canonical":{"stemmed":"Fusinus eucos�n","simple":"Fusinus eucos�nius","full":"Fusinus eucos�nius"},"cardinality":2,"details":{"species":{"genus":"Fusinus","species":"eucos�nius"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":18}],"id":"157cf8c1-0b0d-5b81-a3a9-f02bdc1413a5","parserVersion":"test_version"}
```

### UTF-8 0xA0 character (NO_BREAK_SPACE)

Name: Byssochlamys fulva Olliver & G. Smith

Canonical: Byssochlamys fulva

Authorship: Olliver & G. Smith

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Byssochlamys fulva Olliver \u0026 G. Smith","normalized":"Byssochlamys fulva Olliver \u0026 G. Smith","canonical":{"stemmed":"Byssochlamys fulu","simple":"Byssochlamys fulva","full":"Byssochlamys fulva"},"cardinality":2,"authorship":{"verbatim":"Olliver \u0026 G. Smith","normalized":"Olliver \u0026 G. Smith","authors":["Olliver","G. Smith"],"originalAuth":{"authors":["Olliver","G. Smith"]}},"details":{"species":{"genus":"Byssochlamys","species":"fulva","authorship":{"verbatim":"Olliver \u0026 G. Smith","normalized":"Olliver \u0026 G. Smith","authors":["Olliver","G. Smith"],"originalAuth":{"authors":["Olliver","G. Smith"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":18},{"wordType":"authorWord","start":19,"end":26},{"wordType":"authorWord","start":29,"end":31},{"wordType":"authorWord","start":32,"end":37}],"id":"83523455-cfe4-5ff9-bc54-841f026576b7","parserVersion":"test_version"}
```

### UTF-8 0x3000 character (IDEOGRAPHIC_SPACE)

Name: Kinosternidae　Agassiz, 1857

Canonical: Kinosternidae

Authorship: Agassiz 1857

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Kinosternidae　Agassiz, 1857","normalized":"Kinosternidae Agassiz 1857","canonical":{"stemmed":"Kinosternidae","simple":"Kinosternidae","full":"Kinosternidae"},"cardinality":1,"authorship":{"verbatim":"Agassiz, 1857","normalized":"Agassiz 1857","year":"1857","authors":["Agassiz"],"originalAuth":{"authors":["Agassiz"],"year":{"year":"1857"}}},"details":{"uninomial":{"uninomial":"Kinosternidae","authorship":{"verbatim":"Agassiz, 1857","normalized":"Agassiz 1857","year":"1857","authors":["Agassiz"],"originalAuth":{"authors":["Agassiz"],"year":{"year":"1857"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":13},{"wordType":"authorWord","start":14,"end":21},{"wordType":"year","start":23,"end":27}],"id":"7e74b6b8-5242-5802-9238-320192f4eaa4","parserVersion":"test_version"}
```

### Punctuation in the end

Name: Melanius:

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Melanius:","cardinality":0,"id":"0a761224-66db-55b4-b6f0-85de52534125","parserVersion":"test_version"}
```

Name: Negalasa fumalis Barnes & McDunnough 1913. Next sentence

Canonical: Negalasa fumalis

Authorship: Barnes & Mc Dunnough 1913

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis Barnes \u0026 McDunnough 1913. Next sentence","normalized":"Negalasa fumalis Barnes \u0026 Mc Dunnough 1913","canonical":{"stemmed":"Negalasa fumal","simple":"Negalasa fumalis","full":"Negalasa fumalis"},"cardinality":2,"authorship":{"verbatim":"Barnes \u0026 McDunnough 1913.","normalized":"Barnes \u0026 Mc Dunnough 1913","year":"1913","authors":["Barnes","Mc Dunnough"],"originalAuth":{"authors":["Barnes","Mc Dunnough"],"year":{"year":"1913"}}},"tail":" Next sentence","details":{"species":{"genus":"Negalasa","species":"fumalis","authorship":{"verbatim":"Barnes \u0026 McDunnough 1913.","normalized":"Barnes \u0026 Mc Dunnough 1913","year":"1913","authors":["Barnes","Mc Dunnough"],"originalAuth":{"authors":["Barnes","Mc Dunnough"],"year":{"year":"1913"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":17,"end":23},{"wordType":"authorWord","start":26,"end":28},{"wordType":"authorWord","start":28,"end":36},{"wordType":"year","start":37,"end":41}],"id":"45b7343f-d42a-52d5-b0a4-25956d46427b","parserVersion":"test_version"}
```

Name: Negalasa fumalis. Next sentence

Canonical: Negalasa

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis. Next sentence","normalized":"Negalasa","canonical":{"stemmed":"Negalasa","simple":"Negalasa","full":"Negalasa"},"cardinality":1,"tail":" fumalis. Next sentence","details":{"uninomial":{"uninomial":"Negalasa"}},"pos":[{"wordType":"uninomial","start":0,"end":8}],"id":"ce740482-fa87-5d84-b335-1c063fd18de1","parserVersion":"test_version"}
```

Name: Negalasa fumalis, continuation of a sentence

Canonical: Negalasa

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis, continuation of a sentence","normalized":"Negalasa","canonical":{"stemmed":"Negalasa","simple":"Negalasa","full":"Negalasa"},"cardinality":1,"tail":" fumalis, continuation of a sentence","details":{"uninomial":{"uninomial":"Negalasa"}},"pos":[{"wordType":"uninomial","start":0,"end":8}],"id":"7862a3d9-ba4d-5f53-a106-ea048e558f1a","parserVersion":"test_version"}
```

Name: Negalasa fumalis Barnes; something else

Canonical: Negalasa fumalis

Authorship: Barnes

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negalasa fumalis Barnes; something else","normalized":"Negalasa fumalis Barnes","canonical":{"stemmed":"Negalasa fumal","simple":"Negalasa fumalis","full":"Negalasa fumalis"},"cardinality":2,"authorship":{"verbatim":"Barnes","normalized":"Barnes","authors":["Barnes"],"originalAuth":{"authors":["Barnes"]}},"tail":"; something else","details":{"species":{"genus":"Negalasa","species":"fumalis","authorship":{"verbatim":"Barnes","normalized":"Barnes","authors":["Barnes"],"originalAuth":{"authors":["Barnes"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":17,"end":23}],"id":"6359dac4-1a88-5b41-86d3-9c01aaee4a2e","parserVersion":"test_version"}
```

Name: Negaprion brevirostris Negaprion brevirostris, the rest of the sentence

Canonical: Negaprion brevirostris

Authorship: Negaprion

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negaprion brevirostris Negaprion brevirostris, the rest of the sentence","normalized":"Negaprion brevirostris Negaprion","canonical":{"stemmed":"Negaprion breuirostr","simple":"Negaprion brevirostris","full":"Negaprion brevirostris"},"cardinality":2,"authorship":{"verbatim":"Negaprion","normalized":"Negaprion","authors":["Negaprion"],"originalAuth":{"authors":["Negaprion"]}},"tail":" brevirostris, the rest of the sentence","details":{"species":{"genus":"Negaprion","species":"brevirostris","authorship":{"verbatim":"Negaprion","normalized":"Negaprion","authors":["Negaprion"],"originalAuth":{"authors":["Negaprion"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"authorWord","start":23,"end":32}],"id":"619b95fa-017d-5b9b-b800-64ebd5ed433b","parserVersion":"test_version"}
```

Name: Negaprion fronto (Jordan and Gilbert, 1882):

Canonical: Negaprion fronto

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Negaprion fronto (Jordan and Gilbert, 1882):","normalized":"Negaprion fronto","canonical":{"stemmed":"Negaprion front","simple":"Negaprion fronto","full":"Negaprion fronto"},"cardinality":2,"tail":" (Jordan and Gilbert, 1882):","details":{"species":{"genus":"Negaprion","species":"fronto"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":16}],"id":"4bb6a543-d757-5fa5-ae8b-a5ac95722e1d","parserVersion":"test_version"}
```

### Names with 'ex' as sp. epithet

<!-- not dealing with this misspelling...-->
Name: Acanthochiton ex quisitus

Canonical: Acanthochiton

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acanthochiton ex quisitus","normalized":"Acanthochiton","canonical":{"stemmed":"Acanthochiton","simple":"Acanthochiton","full":"Acanthochiton"},"cardinality":1,"tail":" ex quisitus","details":{"uninomial":{"uninomial":"Acanthochiton"}},"pos":[{"wordType":"uninomial","start":0,"end":13}],"id":"00392ae2-1bd9-5a14-bea9-9d26f1107892","parserVersion":"test_version"}
```

### Names with Spanish 'y' instead of '&'

Name: Caloptenopsis crassiusculus (Martínez y Fernández-Castillo, 1896)

Canonical: Caloptenopsis crassiusculus

Authorship: (Martínez & Fernández-Castillo 1896)

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Caloptenopsis crassiusculus (Martínez y Fernández-Castillo, 1896)","normalized":"Caloptenopsis crassiusculus (Martínez \u0026 Fernández-Castillo 1896)","canonical":{"stemmed":"Caloptenopsis crassiuscul","simple":"Caloptenopsis crassiusculus","full":"Caloptenopsis crassiusculus"},"cardinality":2,"authorship":{"verbatim":"(Martínez y Fernández-Castillo, 1896)","normalized":"(Martínez \u0026 Fernández-Castillo 1896)","year":"1896","authors":["Martínez","Fernández-Castillo"],"originalAuth":{"authors":["Martínez","Fernández-Castillo"],"year":{"year":"1896"}}},"details":{"species":{"genus":"Caloptenopsis","species":"crassiusculus","authorship":{"verbatim":"(Martínez y Fernández-Castillo, 1896)","normalized":"(Martínez \u0026 Fernández-Castillo 1896)","year":"1896","authors":["Martínez","Fernández-Castillo"],"originalAuth":{"authors":["Martínez","Fernández-Castillo"],"year":{"year":"1896"}}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":27},{"wordType":"authorWord","start":29,"end":37},{"wordType":"authorWord","start":40,"end":58},{"wordType":"year","start":60,"end":64}],"id":"0080ce8d-aba5-512d-8e33-8ee3914e386a","parserVersion":"test_version"}
```

Name: Dicranum saxatile Lagasca y Segura, García & Clemente y Rubio, 1802

Canonical: Dicranum saxatile

Authorship: Lagasca, Segura, García, Clemente & Rubio 1802

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Dicranum saxatile Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Dicranum saxatile Lagasca, Segura, García, Clemente \u0026 Rubio 1802","canonical":{"stemmed":"Dicranum saxatil","simple":"Dicranum saxatile","full":"Dicranum saxatile"},"cardinality":2,"authorship":{"verbatim":"Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Lagasca, Segura, García, Clemente \u0026 Rubio 1802","year":"1802","authors":["Lagasca","Segura","García","Clemente","Rubio"],"originalAuth":{"authors":["Lagasca","Segura","García","Clemente","Rubio"],"year":{"year":"1802"}}},"details":{"species":{"genus":"Dicranum","species":"saxatile","authorship":{"verbatim":"Lagasca y Segura, García \u0026 Clemente y Rubio, 1802","normalized":"Lagasca, Segura, García, Clemente \u0026 Rubio 1802","year":"1802","authors":["Lagasca","Segura","García","Clemente","Rubio"],"originalAuth":{"authors":["Lagasca","Segura","García","Clemente","Rubio"],"year":{"year":"1802"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":25},{"wordType":"authorWord","start":28,"end":34},{"wordType":"authorWord","start":36,"end":42},{"wordType":"authorWord","start":45,"end":53},{"wordType":"authorWord","start":56,"end":61},{"wordType":"year","start":63,"end":67}],"id":"39054306-2722-5119-a040-f8671b5b31a0","parserVersion":"test_version"}
```

Name: Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger & Coronado 1967

Canonical: Carabus hendrichsi

Authorship: Bolvar, Pieltain, Rotger & Coronado 1967

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Spanish 'y' is used instead of '&'"}],"verbatim":"Carabus (Tanaocarabus) hendrichsi Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Carabus (Tanaocarabus) hendrichsi Bolvar, Pieltain, Rotger \u0026 Coronado 1967","canonical":{"stemmed":"Carabus hendrichs","simple":"Carabus hendrichsi","full":"Carabus hendrichsi"},"cardinality":2,"authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado"],"year":{"year":"1967"}}},"details":{"species":{"genus":"Carabus","subGenus":"Tanaocarabus","species":"hendrichsi","authorship":{"verbatim":"Bolvar y Pieltain, Rotger \u0026 Coronado 1967","normalized":"Bolvar, Pieltain, Rotger \u0026 Coronado 1967","year":"1967","authors":["Bolvar","Pieltain","Rotger","Coronado"],"originalAuth":{"authors":["Bolvar","Pieltain","Rotger","Coronado"],"year":{"year":"1967"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"infragenericEpithet","start":9,"end":21},{"wordType":"specificEpithet","start":23,"end":33},{"wordType":"authorWord","start":34,"end":40},{"wordType":"authorWord","start":43,"end":51},{"wordType":"authorWord","start":53,"end":59},{"wordType":"authorWord","start":62,"end":70},{"wordType":"year","start":71,"end":75}],"id":"519c0687-2303-5b8c-a69f-68e2bd055b5e","parserVersion":"test_version"}
```

### Names with unparsed "tail" at the end

Name: Morea (Morea) Burt 2342343242 23424322342 23424234

Canonical: Morea subgen. Morea

Authorship: Burt

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Combination of two uninomials"}],"verbatim":"Morea (Morea) Burt 2342343242 23424322342 23424234","normalized":"Morea subgen. Morea Burt","canonical":{"stemmed":"Morea","simple":"Morea","full":"Morea subgen. Morea"},"cardinality":1,"authorship":{"verbatim":"Burt","normalized":"Burt","authors":["Burt"],"originalAuth":{"authors":["Burt"]}},"tail":" 2342343242 23424322342 23424234","details":{"uninomial":{"uninomial":"Morea","rank":"subgen.","parent":"Morea","authorship":{"verbatim":"Burt","normalized":"Burt","authors":["Burt"],"originalAuth":{"authors":["Burt"]}}}},"pos":[{"wordType":"uninomial","start":0,"end":5},{"wordType":"uninomial","start":7,"end":12},{"wordType":"authorWord","start":14,"end":18}],"id":"ca23679f-f3d8-5194-a406-048f970c4020","parserVersion":"test_version"}
```

Name: Nautilus asterizans von

Canonical: Nautilus asterizans

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nautilus asterizans von","normalized":"Nautilus asterizans","canonical":{"stemmed":"Nautilus asterizans","simple":"Nautilus asterizans","full":"Nautilus asterizans"},"cardinality":2,"tail":" von","details":{"species":{"genus":"Nautilus","species":"asterizans"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19}],"id":"0716f658-c952-5415-b2ad-79a39c2b7b0d","parserVersion":"test_version"}
```

### Discard apostrophes at the start and end of words

Name: Acer 'lanum'

Canonical: Acer

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acer 'lanum'","normalized":"Acer","canonical":{"stemmed":"Acer","simple":"Acer","full":"Acer"},"cardinality":1,"tail":" 'lanum'","details":{"uninomial":{"uninomial":"Acer"}},"pos":[{"wordType":"uninomial","start":0,"end":4}],"id":"2db01ed9-9983-5b33-bc2c-8e272539b928","parserVersion":"test_version"}
```

Name: Labeotropheus trewavasae 'albino

Canonical: Labeotropheus trewavasae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Labeotropheus trewavasae 'albino","normalized":"Labeotropheus trewavasae","canonical":{"stemmed":"Labeotropheus trewauas","simple":"Labeotropheus trewavasae","full":"Labeotropheus trewavasae"},"cardinality":2,"tail":" 'albino","details":{"species":{"genus":"Labeotropheus","species":"trewavasae"}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":24}],"id":"0cb9e0ae-1201-5023-8d20-689d60a3e20c","parserVersion":"test_version"}
```

Name: Labeotropheus trewavasae albino'

Canonical: Labeotropheus trewavasae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Labeotropheus trewavasae albino'","normalized":"Labeotropheus trewavasae","canonical":{"stemmed":"Labeotropheus trewauas","simple":"Labeotropheus trewavasae","full":"Labeotropheus trewavasae"},"cardinality":2,"tail":" albino'","details":{"species":{"genus":"Labeotropheus","species":"trewavasae"}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":24}],"id":"f190cdee-14f0-5174-947d-476dab6baeff","parserVersion":"test_version"}
```

Name: Phedimus takesimensis (Nakai) 't Hart

Canonical: Phedimus takesimensis

Authorship: (Nakai) 't Hart

```json
{"parsed":true,"parseQuality":1,"verbatim":"Phedimus takesimensis (Nakai) 't Hart","normalized":"Phedimus takesimensis (Nakai) 't Hart","canonical":{"stemmed":"Phedimus takesimens","simple":"Phedimus takesimensis","full":"Phedimus takesimensis"},"cardinality":2,"authorship":{"verbatim":"(Nakai) 't Hart","normalized":"(Nakai) 't Hart","authors":["Nakai","'t Hart"],"originalAuth":{"authors":["Nakai"]},"combinationAuth":{"authors":["'t Hart"]}},"details":{"species":{"genus":"Phedimus","species":"takesimensis","authorship":{"verbatim":"(Nakai) 't Hart","normalized":"(Nakai) 't Hart","authors":["Nakai","'t Hart"],"originalAuth":{"authors":["Nakai"]},"combinationAuth":{"authors":["'t Hart"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":21},{"wordType":"authorWord","start":23,"end":28},{"wordType":"authorWord","start":30,"end":32},{"wordType":"authorWord","start":33,"end":37}],"id":"14379aa4-1eb9-5ef7-b355-7e3ef3c1fe5e","parserVersion":"test_version"}
```

### Discard apostrophe with dash (rare, needs further investigation)

<!-- incorrectly parsed, but we will live with it for now-->
Name: Solanum tuberosum wila-k'oy

Canonical: Solanum tuberosum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Solanum tuberosum wila-k'oy","normalized":"Solanum tuberosum","canonical":{"stemmed":"Solanum tuberos","simple":"Solanum tuberosum","full":"Solanum tuberosum"},"cardinality":2,"tail":" wila-k'oy","details":{"species":{"genus":"Solanum","species":"tuberosum"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17}],"id":"3d40713c-3b98-5b38-a3e8-555698722078","parserVersion":"test_version"}
```

<!-- correctly parsed -->
Name: Solanum juzepczukii janck'o-ckaisalla

Canonical: Solanum juzepczukii jancko-ckaisalla

Authorship:

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"Apostrophe is not allowed in canonical"}],"verbatim":"Solanum juzepczukii janck'o-ckaisalla","normalized":"Solanum juzepczukii jancko-ckaisalla","canonical":{"stemmed":"Solanum iuzepczuki iancko-ckaisall","simple":"Solanum juzepczukii jancko-ckaisalla","full":"Solanum juzepczukii jancko-ckaisalla"},"cardinality":3,"details":{"infraSpecies":{"genus":"Solanum","species":"juzepczukii","infraSpecies":[{"value":"jancko-ckaisalla"}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":19},{"wordType":"infraspecificEpithet","start":20,"end":37}],"id":"9ec56934-e986-5392-a531-55d97e5e9dd1","parserVersion":"test_version"}
```

### Possible canonical

Name: Morea (Morea) burtius 2342343242 23424322342 23424234

Canonical: Morea burtius

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Morea (Morea) burtius 2342343242 23424322342 23424234","normalized":"Morea (Morea) burtius","canonical":{"stemmed":"Morea burt","simple":"Morea burtius","full":"Morea burtius"},"cardinality":2,"tail":" 2342343242 23424322342 23424234","details":{"species":{"genus":"Morea","subGenus":"Morea","species":"burtius"}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"infragenericEpithet","start":7,"end":12},{"wordType":"specificEpithet","start":14,"end":21}],"id":"03f59808-c30e-55da-bea5-27aa035feb5d","parserVersion":"test_version"}
```

Name: Verpericola megasoma ""Dall" Pils.

Canonical: Verpericola megasoma

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Verpericola megasoma \"\"Dall\" Pils.","normalized":"Verpericola megasoma","canonical":{"stemmed":"Verpericola megasom","simple":"Verpericola megasoma","full":"Verpericola megasoma"},"cardinality":2,"tail":" \"\"Dall\" Pils.","details":{"species":{"genus":"Verpericola","species":"megasoma"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":20}],"id":"cebb60d9-fc8e-5fa0-874a-ae21819b242b","parserVersion":"test_version"}
```

Name: Verpericola megasoma "Dall" Pils.

Canonical: Verpericola megasoma

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Verpericola megasoma \"Dall\" Pils.","normalized":"Verpericola megasoma","canonical":{"stemmed":"Verpericola megasom","simple":"Verpericola megasoma","full":"Verpericola megasoma"},"cardinality":2,"tail":" \"Dall\" Pils.","details":{"species":{"genus":"Verpericola","species":"megasoma"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":20}],"id":"02011460-ba94-5162-98c9-4064a700c7f8","parserVersion":"test_version"}
```

Name: Moraea spathulata ( (L. f. Klatt

Canonical: Moraea spathulata

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Moraea spathulata ( (L. f. Klatt","normalized":"Moraea spathulata","canonical":{"stemmed":"Moraea spathulat","simple":"Moraea spathulata","full":"Moraea spathulata"},"cardinality":2,"tail":" ( (L. f. Klatt","details":{"species":{"genus":"Moraea","species":"spathulata"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"21cb8638-ff53-534f-b816-1e15ecbb818b","parserVersion":"test_version"}
```

Name: Stewartia micrantha (Chun) Sealy, Bot. Mag. 176: t. 510. 1967.

Canonical: Stewartia micrantha

Authorship: (Chun) Sealy & Bot. Mag.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Stewartia micrantha (Chun) Sealy, Bot. Mag. 176: t. 510. 1967.","normalized":"Stewartia micrantha (Chun) Sealy \u0026 Bot. Mag.","canonical":{"stemmed":"Stewartia micranth","simple":"Stewartia micrantha","full":"Stewartia micrantha"},"cardinality":2,"authorship":{"verbatim":"(Chun) Sealy, Bot. Mag.","normalized":"(Chun) Sealy \u0026 Bot. Mag.","authors":["Chun","Sealy","Bot. Mag."],"originalAuth":{"authors":["Chun"]},"combinationAuth":{"authors":["Sealy","Bot. Mag."]}},"tail":" 176: t. 510. 1967.","details":{"species":{"genus":"Stewartia","species":"micrantha","authorship":{"verbatim":"(Chun) Sealy, Bot. Mag.","normalized":"(Chun) Sealy \u0026 Bot. Mag.","authors":["Chun","Sealy","Bot. Mag."],"originalAuth":{"authors":["Chun"]},"combinationAuth":{"authors":["Sealy","Bot. Mag."]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":19},{"wordType":"authorWord","start":21,"end":25},{"wordType":"authorWord","start":27,"end":32},{"wordType":"authorWord","start":34,"end":38},{"wordType":"authorWord","start":39,"end":43}],"id":"7a4ffc19-61a9-551b-bea2-ebb0f5fe9c5a","parserVersion":"test_version"}
```

Name: Pyrobaculum neutrophilum V24Sta

Canonical: Pyrobaculum neutrophilum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pyrobaculum neutrophilum V24Sta","normalized":"Pyrobaculum neutrophilum","canonical":{"stemmed":"Pyrobaculum neutrophil","simple":"Pyrobaculum neutrophilum","full":"Pyrobaculum neutrophilum"},"cardinality":2,"tail":" V24Sta","details":{"species":{"genus":"Pyrobaculum","species":"neutrophilum"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":24}],"id":"6d0be585-ec54-5662-9d30-1d369ecf2a64","parserVersion":"test_version"}
```

Name: Rana aurora Baird and Girard, 1852; H.B. Shaffer et al., 2004

Canonical: Rana aurora

Authorship: Baird & Girard 1852

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Rana aurora Baird and Girard, 1852; H.B. Shaffer et al., 2004","normalized":"Rana aurora Baird \u0026 Girard 1852","canonical":{"stemmed":"Rana auror","simple":"Rana aurora","full":"Rana aurora"},"cardinality":2,"authorship":{"verbatim":"Baird and Girard, 1852","normalized":"Baird \u0026 Girard 1852","year":"1852","authors":["Baird","Girard"],"originalAuth":{"authors":["Baird","Girard"],"year":{"year":"1852"}}},"tail":"; H.B. Shaffer et al., 2004","details":{"species":{"genus":"Rana","species":"aurora","authorship":{"verbatim":"Baird and Girard, 1852","normalized":"Baird \u0026 Girard 1852","year":"1852","authors":["Baird","Girard"],"originalAuth":{"authors":["Baird","Girard"],"year":{"year":"1852"}}}}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":11},{"wordType":"authorWord","start":12,"end":17},{"wordType":"authorWord","start":22,"end":28},{"wordType":"year","start":30,"end":34}],"id":"f0fa6cd1-8018-5fec-92ad-1bda9ac929ca","parserVersion":"test_version"}
```

Name: Agropyron pectiniforme var. karabaljikji ined.?

Canonical: Agropyron pectiniforme var. karabaljikji

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Agropyron pectiniforme var. karabaljikji ined.?","normalized":"Agropyron pectiniforme var. karabaljikji","canonical":{"stemmed":"Agropyron pectiniform karabaliiki","simple":"Agropyron pectiniforme karabaljikji","full":"Agropyron pectiniforme var. karabaljikji"},"cardinality":3,"tail":" ined.?","details":{"infraSpecies":{"genus":"Agropyron","species":"pectiniforme","infraSpecies":[{"value":"karabaljikji","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"rank","start":23,"end":27},{"wordType":"infraspecificEpithet","start":28,"end":40}],"id":"e951b7d4-0009-54df-9de6-efbb392dc8d6","parserVersion":"test_version"}
```

Name: Staphylococcus hyicus chromogenes Devriese et al. 1978 (Approved Lists 1980).

Canonical: Staphylococcus hyicus chromogenes

Authorship: Devriese et al. 1978

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Staphylococcus hyicus chromogenes Devriese et al. 1978 (Approved Lists 1980).","normalized":"Staphylococcus hyicus chromogenes Devriese et al. 1978","canonical":{"stemmed":"Staphylococcus hyic chromogen","simple":"Staphylococcus hyicus chromogenes","full":"Staphylococcus hyicus chromogenes"},"cardinality":3,"authorship":{"verbatim":"Devriese et al. 1978","normalized":"Devriese et al. 1978","year":"1978","authors":["Devriese et al."],"originalAuth":{"authors":["Devriese et al."],"year":{"year":"1978"}}},"bacteria":"yes","tail":" (Approved Lists 1980).","details":{"infraSpecies":{"genus":"Staphylococcus","species":"hyicus","infraSpecies":[{"value":"chromogenes","authorship":{"verbatim":"Devriese et al. 1978","normalized":"Devriese et al. 1978","year":"1978","authors":["Devriese et al."],"originalAuth":{"authors":["Devriese et al."],"year":{"year":"1978"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":21},{"wordType":"infraspecificEpithet","start":22,"end":33},{"wordType":"authorWord","start":34,"end":42},{"wordType":"authorWord","start":43,"end":49},{"wordType":"year","start":50,"end":54}],"id":"ec17eb44-742c-5325-aca6-e33a0888ef0d","parserVersion":"test_version"}
```

### Treating `& al.` as `et al.`

Name: Adonis cyllenea Boiss. & al.

Canonical: Adonis cyllenea

Authorship: Boiss. et al.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al.","normalized":"Adonis cyllenea Boiss. et al.","canonical":{"stemmed":"Adonis cyllene","simple":"Adonis cyllenea","full":"Adonis cyllenea"},"cardinality":2,"authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"details":{"species":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22},{"wordType":"authorWord","start":23,"end":28}],"id":"a7c2cb28-2ec2-55b5-88a2-6cfd633cbd00","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al

Canonical: Adonis cyllenea

Authorship: Boiss. et al.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al","normalized":"Adonis cyllenea Boiss. et al.","canonical":{"stemmed":"Adonis cyllene","simple":"Adonis cyllenea","full":"Adonis cyllenea"},"cardinality":2,"authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"details":{"species":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22},{"wordType":"authorWord","start":23,"end":27}],"id":"85e122ea-f581-5d4b-a29f-b87c48d0a716","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al. var. paryadrica Boiss.

Canonical: Adonis cyllenea var. paryadrica

Authorship: Boiss.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al. var. paryadrica Boiss.","normalized":"Adonis cyllenea Boiss. et al. var. paryadrica Boiss.","canonical":{"stemmed":"Adonis cyllene paryadric","simple":"Adonis cyllenea paryadrica","full":"Adonis cyllenea var. paryadrica"},"cardinality":3,"authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}},"details":{"infraSpecies":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al.","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"infraSpecies":[{"value":"paryadrica","rank":"var.","authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22},{"wordType":"authorWord","start":23,"end":28},{"wordType":"rank","start":29,"end":33},{"wordType":"infraspecificEpithet","start":34,"end":44},{"wordType":"authorWord","start":45,"end":51}],"id":"6bc790ae-210d-518e-9e20-2d4d517a08ef","parserVersion":"test_version"}
```

Name: Adonis cyllenea Boiss. & al var. paryadrica Boiss.

Canonical: Adonis cyllenea var. paryadrica

Authorship: Boiss.

```json
{"parsed":true,"parseQuality":1,"verbatim":"Adonis cyllenea Boiss. \u0026 al var. paryadrica Boiss.","normalized":"Adonis cyllenea Boiss. et al. var. paryadrica Boiss.","canonical":{"stemmed":"Adonis cyllene paryadric","simple":"Adonis cyllenea paryadrica","full":"Adonis cyllenea var. paryadrica"},"cardinality":3,"authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}},"details":{"infraSpecies":{"genus":"Adonis","species":"cyllenea","authorship":{"verbatim":"Boiss. \u0026 al","normalized":"Boiss. et al.","authors":["Boiss. et al."],"originalAuth":{"authors":["Boiss. et al."]}},"infraSpecies":[{"value":"paryadrica","rank":"var.","authorship":{"verbatim":"Boiss.","normalized":"Boiss.","authors":["Boiss."],"originalAuth":{"authors":["Boiss."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22},{"wordType":"authorWord","start":23,"end":27},{"wordType":"rank","start":28,"end":32},{"wordType":"infraspecificEpithet","start":33,"end":43},{"wordType":"authorWord","start":44,"end":50}],"id":"eb7aee15-e462-5189-8335-a3a323be6907","parserVersion":"test_version"}
```

### Authors do not start with apostrophe

Name: Nereidavus kulkovi 'Kulkov

Canonical: Nereidavus kulkovi

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nereidavus kulkovi 'Kulkov","normalized":"Nereidavus kulkovi","canonical":{"stemmed":"Nereidavus kulkou","simple":"Nereidavus kulkovi","full":"Nereidavus kulkovi"},"cardinality":2,"tail":" 'Kulkov","details":{"species":{"genus":"Nereidavus","species":"kulkovi"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":18}],"id":"6a4999cd-95cc-509d-8e0a-26a0dfcef67d","parserVersion":"test_version"}
```

### Epithets do not start or end with a dash

Name: Abryna -petri Paiva, 1860

Canonical: Abryna

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abryna -petri Paiva, 1860","normalized":"Abryna","canonical":{"stemmed":"Abryna","simple":"Abryna","full":"Abryna"},"cardinality":1,"tail":" -petri Paiva, 1860","details":{"uninomial":{"uninomial":"Abryna"}},"pos":[{"wordType":"uninomial","start":0,"end":6}],"id":"6ccc6217-9084-5b31-81f7-6b4cd7963f65","parserVersion":"test_version"}
```

Name: Abryna petri- Paiva, 1860

Canonical: Abryna

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abryna petri- Paiva, 1860","normalized":"Abryna","canonical":{"stemmed":"Abryna","simple":"Abryna","full":"Abryna"},"cardinality":1,"tail":" petri- Paiva, 1860","details":{"uninomial":{"uninomial":"Abryna"}},"pos":[{"wordType":"uninomial","start":0,"end":6}],"id":"b1e37ace-3ca8-5274-bd93-7333aa3e5223","parserVersion":"test_version"}
```

### names that contain "of"

Name: Musca capraria Trustees of the British Museum (Natural History), 1939

Canonical: Musca capraria

Authorship: Trustees

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Musca capraria Trustees of the British Museum (Natural History), 1939","normalized":"Musca capraria Trustees","canonical":{"stemmed":"Musca caprar","simple":"Musca capraria","full":"Musca capraria"},"cardinality":2,"authorship":{"verbatim":"Trustees","normalized":"Trustees","authors":["Trustees"],"originalAuth":{"authors":["Trustees"]}},"tail":" of the British Museum (Natural History), 1939","details":{"species":{"genus":"Musca","species":"capraria","authorship":{"verbatim":"Trustees","normalized":"Trustees","authors":["Trustees"],"originalAuth":{"authors":["Trustees"]}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":14},{"wordType":"authorWord","start":15,"end":23}],"id":"aa70cf4b-14bb-57a3-9fe1-0a9a544a16da","parserVersion":"test_version"}
```

Name: Nassellarid genera of uncertain affinities

Canonical: Nassellarid genera

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Nassellarid genera of uncertain affinities","normalized":"Nassellarid genera","canonical":{"stemmed":"Nassellarid gener","simple":"Nassellarid genera","full":"Nassellarid genera"},"cardinality":2,"tail":" of uncertain affinities","details":{"species":{"genus":"Nassellarid","species":"genera"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":18}],"id":"ca46eccc-6b42-5faf-be0f-aad069d3e3dd","parserVersion":"test_version"}
```

Name: Natica of nidus

Canonical: Natica

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Natica of nidus","normalized":"Natica","canonical":{"stemmed":"Natica","simple":"Natica","full":"Natica"},"cardinality":1,"tail":" of nidus","details":{"uninomial":{"uninomial":"Natica"}},"pos":[{"wordType":"uninomial","start":0,"end":6}],"id":"6a049500-f407-56e7-80b4-41ab91f64b8c","parserVersion":"test_version"}
```

Name: Neritina chemmoi Reeve var of cornea Linn

Canonical: Neritina chemmoi

Authorship: Reeve

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Neritina chemmoi Reeve var of cornea Linn","normalized":"Neritina chemmoi Reeve","canonical":{"stemmed":"Neritina chemmo","simple":"Neritina chemmoi","full":"Neritina chemmoi"},"cardinality":2,"authorship":{"verbatim":"Reeve","normalized":"Reeve","authors":["Reeve"],"originalAuth":{"authors":["Reeve"]}},"tail":" var of cornea Linn","details":{"species":{"genus":"Neritina","species":"chemmoi","authorship":{"verbatim":"Reeve","normalized":"Reeve","authors":["Reeve"],"originalAuth":{"authors":["Reeve"]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":16},{"wordType":"authorWord","start":17,"end":22}],"id":"d6cbded0-dc9b-5da2-8fb9-8d8b124cc5b4","parserVersion":"test_version"}
```

Name: Wolbachia endosymbiont of Leptogenys gracilis

Canonical: Wolbachia endosymbiont

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Wolbachia endosymbiont of Leptogenys gracilis","normalized":"Wolbachia endosymbiont","canonical":{"stemmed":"Wolbachia endosymbio","simple":"Wolbachia endosymbiont","full":"Wolbachia endosymbiont"},"cardinality":2,"bacteria":"yes","tail":" of Leptogenys gracilis","details":{"species":{"genus":"Wolbachia","species":"endosymbiont"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22}],"id":"ed4bbf5e-068a-518a-8eb3-42ead52b941b","parserVersion":"test_version"}
```

### Names that contain "cv" (cultivar)

Name: Phyllostachys vivax cv aureocaulis

Canonical: Phyllostachys vivax

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Phyllostachys vivax cv aureocaulis","normalized":"Phyllostachys vivax","canonical":{"stemmed":"Phyllostachys uiuax","simple":"Phyllostachys vivax","full":"Phyllostachys vivax"},"cardinality":2,"tail":" cv aureocaulis","details":{"species":{"genus":"Phyllostachys","species":"vivax"}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":19}],"id":"56f7057d-9c5c-5ac7-bc7a-f631fb58f5d6","parserVersion":"test_version"}
```

Name: Rhododendron cv Cilpinense

Canonical: Rhododendron

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Rhododendron cv Cilpinense","normalized":"Rhododendron","canonical":{"stemmed":"Rhododendron","simple":"Rhododendron","full":"Rhododendron"},"cardinality":1,"tail":" cv Cilpinense","details":{"uninomial":{"uninomial":"Rhododendron"}},"pos":[{"wordType":"uninomial","start":0,"end":12}],"id":"abd299df-e4b2-533c-86eb-a4a5e273b934","parserVersion":"test_version"}
```

Name: Ligusticum sinense cv 'chuanxiong' S.H. Qiu & et al.

Canonical: Ligusticum sinense

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ligusticum sinense cv 'chuanxiong' S.H. Qiu \u0026 et al.","normalized":"Ligusticum sinense","canonical":{"stemmed":"Ligusticum sinens","simple":"Ligusticum sinense","full":"Ligusticum sinense"},"cardinality":2,"tail":" cv 'chuanxiong' S.H. Qiu \u0026 et al.","details":{"species":{"genus":"Ligusticum","species":"sinense"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":18}],"id":"73f015c2-6679-5428-b418-6f4487af419d","parserVersion":"test_version"}
```


### "Open taxonomy" with ranks unfinished

Name: Alyxia reinwardti var

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti var","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" var","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"2f0ee2be-8d37-5e43-9eed-776c17f47e93","parserVersion":"test_version"}
```

Name: Alyxia reinwardti var.

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti var.","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" var.","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"aed34708-82ed-52e4-876f-d4468af73fc3","parserVersion":"test_version"}
```

Name: Alyxia reinwardti ssp

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti ssp","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" ssp","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"760486d1-93ed-55c5-ade1-ba2c5b2aa900","parserVersion":"test_version"}
```

Name: Alyxia reinwardti ssp.

Canonical: Alyxia reinwardti

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Alyxia reinwardti ssp.","normalized":"Alyxia reinwardti","canonical":{"stemmed":"Alyxia reinwardt","simple":"Alyxia reinwardti","full":"Alyxia reinwardti"},"cardinality":2,"tail":" ssp.","details":{"species":{"genus":"Alyxia","species":"reinwardti"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"72b5072a-d952-54f8-aea1-5b5bd3c65c45","parserVersion":"test_version"}
```

Name: Alaria spp

Canonical: Alaria

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Alaria spp","normalized":"Alaria","canonical":{"stemmed":"Alaria","simple":"Alaria","full":"Alaria"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Alaria","approximationMarker":"spp"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"approximationMarker","start":7,"end":10}],"id":"5b31e830-ccf6-5918-94c5-75c4db7ef302","parserVersion":"test_version"}
```

Name: Alaria spp.

Canonical: Alaria

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Alaria spp.","normalized":"Alaria","canonical":{"stemmed":"Alaria","simple":"Alaria","full":"Alaria"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Alaria","approximationMarker":"spp."}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"approximationMarker","start":7,"end":11}],"id":"d1cd4f1a-f511-5d5a-8f41-64911995fdec","parserVersion":"test_version"}
```

Name: Xenodon sp

Canonical: Xenodon

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Xenodon sp","normalized":"Xenodon","canonical":{"stemmed":"Xenodon","simple":"Xenodon","full":"Xenodon"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Xenodon","approximationMarker":"sp"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"approximationMarker","start":8,"end":10}],"id":"7b0cb348-7fe9-5248-b396-b0336225ba2a","parserVersion":"test_version"}
```

Name: Xenodon sp.

Canonical: Xenodon

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Xenodon sp.","normalized":"Xenodon","canonical":{"stemmed":"Xenodon","simple":"Xenodon","full":"Xenodon"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Xenodon","approximationMarker":"sp."}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"approximationMarker","start":8,"end":11}],"id":"77b6718f-a26e-5ddf-a4cf-119e972cd015","parserVersion":"test_version"}
```

Name: Formicidae cf.

Canonical: Formicidae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Formicidae cf.","normalized":"Formicidae cf.","canonical":{"stemmed":"Formicidae","simple":"Formicidae","full":"Formicidae"},"cardinality":1,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Formicidae","comparisonMarker":"cf."}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"comparisonMarker","start":11,"end":14}],"id":"61f9ebc4-346e-5857-ab45-38808ff1c960","parserVersion":"test_version"}
```

Name: Formicidae cf

Canonical: Formicidae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name comparison"}],"verbatim":"Formicidae cf","normalized":"Formicidae cf","canonical":{"stemmed":"Formicidae","simple":"Formicidae","full":"Formicidae"},"cardinality":1,"surrogate":"COMPARISON","details":{"comparison":{"genus":"Formicidae","comparisonMarker":"cf"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"comparisonMarker","start":11,"end":13}],"id":"90473425-7ce1-5ec6-8160-737646816ea7","parserVersion":"test_version"}
```

<!-- We do not cover infraspecific comparisons yet -->
Name: Arctostaphylos preglauca cf.

Canonical: Arctostaphylos preglauca

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arctostaphylos preglauca cf.","normalized":"Arctostaphylos preglauca","canonical":{"stemmed":"Arctostaphylos preglauc","simple":"Arctostaphylos preglauca","full":"Arctostaphylos preglauca"},"cardinality":2,"tail":" cf.","details":{"species":{"genus":"Arctostaphylos","species":"preglauca"}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":24}],"id":"246b43d4-9786-5157-8d35-b81a470e6379","parserVersion":"test_version"}
```

Name: Acastoides spp.

Canonical: Acastoides

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Name is approximate"}],"verbatim":"Acastoides spp.","normalized":"Acastoides","canonical":{"stemmed":"Acastoides","simple":"Acastoides","full":"Acastoides"},"cardinality":0,"surrogate":"APPROXIMATION","details":{"approximation":{"genus":"Acastoides","approximationMarker":"spp."}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"approximationMarker","start":11,"end":15}],"id":"9853f0a4-6324-5a7d-8108-e910578e612b","parserVersion":"test_version"}
```

### Ignoring sensu sec

Name: Senecio legionensis sensu Samp., non Lange

Canonical: Senecio legionensis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Senecio legionensis sensu Samp., non Lange","normalized":"Senecio legionensis","canonical":{"stemmed":"Senecio legionens","simple":"Senecio legionensis","full":"Senecio legionensis"},"cardinality":2,"tail":" sensu Samp., non Lange","details":{"species":{"genus":"Senecio","species":"legionensis"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":19}],"id":"948d73b7-499b-5060-ace4-dd061f2f4373","parserVersion":"test_version"}
```

Name: Pseudomonas methanica (Söhngen 1906) sensu. Dworkin and Foster 1956

Canonical: Pseudomonas methanica

Authorship: (Söhngen 1906)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Pseudomonas methanica (Söhngen 1906) sensu. Dworkin and Foster 1956","normalized":"Pseudomonas methanica (Söhngen 1906)","canonical":{"stemmed":"Pseudomonas methanic","simple":"Pseudomonas methanica","full":"Pseudomonas methanica"},"cardinality":2,"authorship":{"verbatim":"(Söhngen 1906)","normalized":"(Söhngen 1906)","year":"1906","authors":["Söhngen"],"originalAuth":{"authors":["Söhngen"],"year":{"year":"1906"}}},"bacteria":"yes","tail":" sensu. Dworkin and Foster 1956","details":{"species":{"genus":"Pseudomonas","species":"methanica","authorship":{"verbatim":"(Söhngen 1906)","normalized":"(Söhngen 1906)","year":"1906","authors":["Söhngen"],"originalAuth":{"authors":["Söhngen"],"year":{"year":"1906"}}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":23,"end":30},{"wordType":"year","start":31,"end":35}],"id":"f4261966-4f80-52c1-a3ff-8eaece507964","parserVersion":"test_version"}
```

Name: Abarema scutifera sensu auct., non (Blanco)Kosterm.

Canonical: Abarema scutifera

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema scutifera sensu auct., non (Blanco)Kosterm.","normalized":"Abarema scutifera","canonical":{"stemmed":"Abarema scutifer","simple":"Abarema scutifera","full":"Abarema scutifera"},"cardinality":2,"tail":" sensu auct., non (Blanco)Kosterm.","details":{"species":{"genus":"Abarema","species":"scutifera"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17}],"id":"59f4b32d-3f8c-569f-bc81-3fe49d708c88","parserVersion":"test_version"}
```

Name: Puya acris Auct.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris Auct.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" Auct.","details":{"species":{"genus":"Puya","species":"acris"}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":10}],"id":"926ec12b-a597-5842-92f2-4b0ae4989df1","parserVersion":"test_version"}
```

Name: Puya acris Auct non L.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris Auct non L.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" Auct non L.","details":{"species":{"genus":"Puya","species":"acris"}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":10}],"id":"6c11df68-9e9d-5e97-b0f0-3609e4f18121","parserVersion":"test_version"}
```

Name: Galium tricorne Stokes, pro parte

Canonical: Galium tricorne

Authorship: Stokes

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Galium tricorne Stokes, pro parte","normalized":"Galium tricorne Stokes","canonical":{"stemmed":"Galium tricorn","simple":"Galium tricorne","full":"Galium tricorne"},"cardinality":2,"authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}},"tail":", pro parte","details":{"species":{"genus":"Galium","species":"tricorne","authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22}],"id":"c4d3da85-86b7-5ca9-925b-6e09ffad3a30","parserVersion":"test_version"}
```

Name: Galium tricorne Stokes,pro parte

Canonical: Galium tricorne

Authorship: Stokes

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Galium tricorne Stokes,pro parte","normalized":"Galium tricorne Stokes","canonical":{"stemmed":"Galium tricorn","simple":"Galium tricorne","full":"Galium tricorne"},"cardinality":2,"authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}},"tail":",pro parte","details":{"species":{"genus":"Galium","species":"tricorne","authorship":{"verbatim":"Stokes","normalized":"Stokes","authors":["Stokes"],"originalAuth":{"authors":["Stokes"]}}}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":15},{"wordType":"authorWord","start":16,"end":22}],"id":"7166cbd9-2b0f-5537-9ac9-98157b60a395","parserVersion":"test_version"}
```

Name: Senecio jacquinianus sec. Rchb.

Canonical: Senecio jacquinianus

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Senecio jacquinianus sec. Rchb.","normalized":"Senecio jacquinianus","canonical":{"stemmed":"Senecio iacquinian","simple":"Senecio jacquinianus","full":"Senecio jacquinianus"},"cardinality":2,"tail":" sec. Rchb.","details":{"species":{"genus":"Senecio","species":"jacquinianus"}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":20}],"id":"e8ad283f-afa8-5fd2-ae8f-bbedf2fb0bb7","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum s.l. (Schultes) Boiss.

Canonical: Acantholimon ulicinum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acantholimon ulicinum s.l. (Schultes) Boiss.","normalized":"Acantholimon ulicinum","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"tail":" s.l. (Schultes) Boiss.","details":{"species":{"genus":"Acantholimon","species":"ulicinum"}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":21}],"id":"cf4b7aa4-b78f-5b79-86c3-9416de24c918","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum s. l. (Schultes) Boiss.

Canonical: Acantholimon ulicinum

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acantholimon ulicinum s. l. (Schultes) Boiss.","normalized":"Acantholimon ulicinum","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"tail":" s. l. (Schultes) Boiss.","details":{"species":{"genus":"Acantholimon","species":"ulicinum"}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":21}],"id":"3a0b0412-f076-5714-8537-62761718ca7c","parserVersion":"test_version"}
```

Name: Acantholimon ulicinum S. L. Schultes

Canonical: Acantholimon ulicinum

Authorship: S. L. Schultes

```json
{"parsed":true,"parseQuality":1,"verbatim":"Acantholimon ulicinum S. L. Schultes","normalized":"Acantholimon ulicinum S. L. Schultes","canonical":{"stemmed":"Acantholimon ulicin","simple":"Acantholimon ulicinum","full":"Acantholimon ulicinum"},"cardinality":2,"authorship":{"verbatim":"S. L. Schultes","normalized":"S. L. Schultes","authors":["S. L. Schultes"],"originalAuth":{"authors":["S. L. Schultes"]}},"details":{"species":{"genus":"Acantholimon","species":"ulicinum","authorship":{"verbatim":"S. L. Schultes","normalized":"S. L. Schultes","authors":["S. L. Schultes"],"originalAuth":{"authors":["S. L. Schultes"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":21},{"wordType":"authorWord","start":22,"end":24},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":28,"end":36}],"id":"702f97e0-792b-5ed4-b2d5-d813544c4139","parserVersion":"test_version"}
```

Name: Amitostigma formosana (S.S.Ying) S.S.Ying

Canonical: Amitostigma formosana

Authorship: (S. S. Ying) S. S. Ying

```json
{"parsed":true,"parseQuality":1,"verbatim":"Amitostigma formosana (S.S.Ying) S.S.Ying","normalized":"Amitostigma formosana (S. S. Ying) S. S. Ying","canonical":{"stemmed":"Amitostigma formosan","simple":"Amitostigma formosana","full":"Amitostigma formosana"},"cardinality":2,"authorship":{"verbatim":"(S.S.Ying) S.S.Ying","normalized":"(S. S. Ying) S. S. Ying","authors":["S. S. Ying","S. S. Ying"],"originalAuth":{"authors":["S. S. Ying"]},"combinationAuth":{"authors":["S. S. Ying"]}},"details":{"species":{"genus":"Amitostigma","species":"formosana","authorship":{"verbatim":"(S.S.Ying) S.S.Ying","normalized":"(S. S. Ying) S. S. Ying","authors":["S. S. Ying","S. S. Ying"],"originalAuth":{"authors":["S. S. Ying"]},"combinationAuth":{"authors":["S. S. Ying"]}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":21},{"wordType":"authorWord","start":23,"end":25},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":27,"end":31},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":35,"end":37},{"wordType":"authorWord","start":37,"end":41}],"id":"fcd831ea-57b6-5151-81e4-86e1c42f4695","parserVersion":"test_version"}
```

Name: Amaurorhinus bewichianus (Wollaston,1860) (s.str.)

Canonical: Amaurorhinus bewichianus

Authorship: (Wollaston 1860)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Amaurorhinus bewichianus (Wollaston,1860) (s.str.)","normalized":"Amaurorhinus bewichianus (Wollaston 1860)","canonical":{"stemmed":"Amaurorhinus bewichian","simple":"Amaurorhinus bewichianus","full":"Amaurorhinus bewichianus"},"cardinality":2,"authorship":{"verbatim":"(Wollaston,1860)","normalized":"(Wollaston 1860)","year":"1860","authors":["Wollaston"],"originalAuth":{"authors":["Wollaston"],"year":{"year":"1860"}}},"tail":" (s.str.)","details":{"species":{"genus":"Amaurorhinus","species":"bewichianus","authorship":{"verbatim":"(Wollaston,1860)","normalized":"(Wollaston 1860)","year":"1860","authors":["Wollaston"],"originalAuth":{"authors":["Wollaston"],"year":{"year":"1860"}}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":24},{"wordType":"authorWord","start":26,"end":35},{"wordType":"year","start":36,"end":40}],"id":"b76e9160-d301-5696-bb87-499328996a7d","parserVersion":"test_version"}
```

Name: Ammodramus caudacutus (s.s.) diversus

Canonical: Ammodramus caudacutus

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ammodramus caudacutus (s.s.) diversus","normalized":"Ammodramus caudacutus","canonical":{"stemmed":"Ammodramus caudacut","simple":"Ammodramus caudacutus","full":"Ammodramus caudacutus"},"cardinality":2,"tail":" (s.s.) diversus","details":{"species":{"genus":"Ammodramus","species":"caudacutus"}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":21}],"id":"2fb79b29-1579-5604-97bd-530c90c245cd","parserVersion":"test_version"}
```

Name: Arenaria serpyllifolia L. s.str.

Canonical: Arenaria serpyllifolia

Authorship: L.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arenaria serpyllifolia L. s.str.","normalized":"Arenaria serpyllifolia L.","canonical":{"stemmed":"Arenaria serpyllifol","simple":"Arenaria serpyllifolia","full":"Arenaria serpyllifolia"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" s.str.","details":{"species":{"genus":"Arenaria","species":"serpyllifolia","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":22},{"wordType":"authorWord","start":23,"end":25}],"id":"8a350298-0dfc-5ad0-9a10-60902587f335","parserVersion":"test_version"}
```

Name: Asplenium trichomanes L. s.lat. - Asplen trich

Canonical: Asplenium trichomanes

Authorship: L.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium trichomanes L. s.lat. - Asplen trich","normalized":"Asplenium trichomanes L.","canonical":{"stemmed":"Asplenium trichoman","simple":"Asplenium trichomanes","full":"Asplenium trichomanes"},"cardinality":2,"authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}},"tail":" s.lat. - Asplen trich","details":{"species":{"genus":"Asplenium","species":"trichomanes","authorship":{"verbatim":"L.","normalized":"L.","authors":["L."],"originalAuth":{"authors":["L."]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":21},{"wordType":"authorWord","start":22,"end":24}],"id":"1687d870-6bea-5573-80ef-4e55eca3199f","parserVersion":"test_version"}
```

Name: Asplenium anisophyllum Kunze, s.l.

Canonical: Asplenium anisophyllum

Authorship: Kunze

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium anisophyllum Kunze, s.l.","normalized":"Asplenium anisophyllum Kunze","canonical":{"stemmed":"Asplenium anisophyll","simple":"Asplenium anisophyllum","full":"Asplenium anisophyllum"},"cardinality":2,"authorship":{"verbatim":"Kunze","normalized":"Kunze","authors":["Kunze"],"originalAuth":{"authors":["Kunze"]}},"tail":", s.l.","details":{"species":{"genus":"Asplenium","species":"anisophyllum","authorship":{"verbatim":"Kunze","normalized":"Kunze","authors":["Kunze"],"originalAuth":{"authors":["Kunze"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"authorWord","start":23,"end":28}],"id":"a0d7a55a-ffad-5243-905e-048177b440df","parserVersion":"test_version"}
```

Name: Abramis Cuvier 1816 sec. Dybowski 1862

Canonical: Abramis

Authorship: Cuvier 1816

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abramis Cuvier 1816 sec. Dybowski 1862","normalized":"Abramis Cuvier 1816","canonical":{"stemmed":"Abramis","simple":"Abramis","full":"Abramis"},"cardinality":1,"authorship":{"verbatim":"Cuvier 1816","normalized":"Cuvier 1816","year":"1816","authors":["Cuvier"],"originalAuth":{"authors":["Cuvier"],"year":{"year":"1816"}}},"tail":" sec. Dybowski 1862","details":{"uninomial":{"uninomial":"Abramis","authorship":{"verbatim":"Cuvier 1816","normalized":"Cuvier 1816","year":"1816","authors":["Cuvier"],"originalAuth":{"authors":["Cuvier"],"year":{"year":"1816"}}}}},"pos":[{"wordType":"uninomial","start":0,"end":7},{"wordType":"authorWord","start":8,"end":14},{"wordType":"year","start":15,"end":19}],"id":"1fddff95-f470-5c36-8bc5-4436fe727bda","parserVersion":"test_version"}
```

Name: Abramis brama subsp. bergi Grib & Vernidub 1935 sec Eschmeyer 2004

Canonical: Abramis brama subsp. bergi

Authorship: Grib & Vernidub 1935

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abramis brama subsp. bergi Grib \u0026 Vernidub 1935 sec Eschmeyer 2004","normalized":"Abramis brama subsp. bergi Grib \u0026 Vernidub 1935","canonical":{"stemmed":"Abramis bram berg","simple":"Abramis brama bergi","full":"Abramis brama subsp. bergi"},"cardinality":3,"authorship":{"verbatim":"Grib \u0026 Vernidub 1935","normalized":"Grib \u0026 Vernidub 1935","year":"1935","authors":["Grib","Vernidub"],"originalAuth":{"authors":["Grib","Vernidub"],"year":{"year":"1935"}}},"tail":" sec Eschmeyer 2004","details":{"infraSpecies":{"genus":"Abramis","species":"brama","infraSpecies":[{"value":"bergi","rank":"subsp.","authorship":{"verbatim":"Grib \u0026 Vernidub 1935","normalized":"Grib \u0026 Vernidub 1935","year":"1935","authors":["Grib","Vernidub"],"originalAuth":{"authors":["Grib","Vernidub"],"year":{"year":"1935"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":13},{"wordType":"rank","start":14,"end":20},{"wordType":"infraspecificEpithet","start":21,"end":26},{"wordType":"authorWord","start":27,"end":31},{"wordType":"authorWord","start":34,"end":42},{"wordType":"year","start":43,"end":47}],"id":"5ac5f7fd-0a42-5133-961e-df94a54fb75f","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., P. P.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., P. P.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", P. P.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":19,"end":23},{"wordType":"authorWord","start":25,"end":33}],"id":"2e18b789-865b-55dc-831b-f1fdd6bf740d","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., p.p.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., p.p.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", p.p.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":19,"end":23},{"wordType":"authorWord","start":25,"end":33}],"id":"bc9b0feb-8a33-5f35-97a9-8ee93220fff8","parserVersion":"test_version"}
```

Name: Abarema clypearia (Jack) Kosterm., p. p.

Canonical: Abarema clypearia

Authorship: (Jack) Kosterm.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abarema clypearia (Jack) Kosterm., p. p.","normalized":"Abarema clypearia (Jack) Kosterm.","canonical":{"stemmed":"Abarema clypear","simple":"Abarema clypearia","full":"Abarema clypearia"},"cardinality":2,"authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}},"tail":", p. p.","details":{"species":{"genus":"Abarema","species":"clypearia","authorship":{"verbatim":"(Jack) Kosterm.","normalized":"(Jack) Kosterm.","authors":["Jack","Kosterm."],"originalAuth":{"authors":["Jack"]},"combinationAuth":{"authors":["Kosterm."]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":19,"end":23},{"wordType":"authorWord","start":25,"end":33}],"id":"1fae34cb-12f4-5600-9589-672199934719","parserVersion":"test_version"}
```

Name: Indigofera phyllogramme var. aphylla R.Vig., p.p.B

Canonical: Indigofera phyllogramme var. aphylla

Authorship: R. Vig.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Indigofera phyllogramme var. aphylla R.Vig., p.p.B","normalized":"Indigofera phyllogramme var. aphylla R. Vig.","canonical":{"stemmed":"Indigofera phyllogramm aphyll","simple":"Indigofera phyllogramme aphylla","full":"Indigofera phyllogramme var. aphylla"},"cardinality":3,"authorship":{"verbatim":"R.Vig.","normalized":"R. Vig.","authors":["R. Vig."],"originalAuth":{"authors":["R. Vig."]}},"tail":", p.p.B","details":{"infraSpecies":{"genus":"Indigofera","species":"phyllogramme","infraSpecies":[{"value":"aphylla","rank":"var.","authorship":{"verbatim":"R.Vig.","normalized":"R. Vig.","authors":["R. Vig."],"originalAuth":{"authors":["R. Vig."]}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":23},{"wordType":"rank","start":24,"end":28},{"wordType":"infraspecificEpithet","start":29,"end":36},{"wordType":"authorWord","start":37,"end":39},{"wordType":"authorWord","start":39,"end":43}],"id":"04bb878e-4442-5b7c-86d7-a41f2f6aefd3","parserVersion":"test_version"}
```

### Unparseable hort. annotations

Name: Asplenium mayi ht.May; Gard.

Canonical: Asplenium mayi

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium mayi ht.May; Gard.","normalized":"Asplenium mayi","canonical":{"stemmed":"Asplenium may","simple":"Asplenium mayi","full":"Asplenium mayi"},"cardinality":2,"tail":" ht.May; Gard.","details":{"species":{"genus":"Asplenium","species":"mayi"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":14}],"id":"74446da2-14ce-5951-95c6-054d29417131","parserVersion":"test_version"}
```

Name: Asplenium mayii ht.May; Gard.

Canonical: Asplenium mayii

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Asplenium mayii ht.May; Gard.","normalized":"Asplenium mayii","canonical":{"stemmed":"Asplenium mayi","simple":"Asplenium mayii","full":"Asplenium mayii"},"cardinality":2,"tail":" ht.May; Gard.","details":{"species":{"genus":"Asplenium","species":"mayii"}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":15}],"id":"00764ac3-b9eb-56bf-9856-6de62459646e","parserVersion":"test_version"}
```

Name: Davallia decora ht.Bull.; Gard.Chr.

Canonical: Davallia decora

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Davallia decora ht.Bull.; Gard.Chr.","normalized":"Davallia decora","canonical":{"stemmed":"Davallia decor","simple":"Davallia decora","full":"Davallia decora"},"cardinality":2,"tail":" ht.Bull.; Gard.Chr.","details":{"species":{"genus":"Davallia","species":"decora"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":15}],"id":"2e6032e9-1a08-5149-8339-5361c84c4a2d","parserVersion":"test_version"}
```

Name: Gymnogramma alstoni ht.Birkenh.; Gard.

Canonical: Gymnogramma alstoni

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Gymnogramma alstoni ht.Birkenh.; Gard.","normalized":"Gymnogramma alstoni","canonical":{"stemmed":"Gymnogramma alston","simple":"Gymnogramma alstoni","full":"Gymnogramma alstoni"},"cardinality":2,"tail":" ht.Birkenh.; Gard.","details":{"species":{"genus":"Gymnogramma","species":"alstoni"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":19}],"id":"77b0759a-2b8f-51ef-8a40-df9268c72cf1","parserVersion":"test_version"}
```

Name: Gymnogramma sprengeriana ht.Wiener Ill.

Canonical: Gymnogramma sprengeriana

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Gymnogramma sprengeriana ht.Wiener Ill.","normalized":"Gymnogramma sprengeriana","canonical":{"stemmed":"Gymnogramma sprengerian","simple":"Gymnogramma sprengeriana","full":"Gymnogramma sprengeriana"},"cardinality":2,"tail":" ht.Wiener Ill.","details":{"species":{"genus":"Gymnogramma","species":"sprengeriana"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":24}],"id":"4e5517fa-4b2c-55f6-8471-76c26ed9983a","parserVersion":"test_version"}
```

### Removing nomenclatural annotations

Name: Amphiprora pseudoduplex (Osada & Kobayasi, 1990) comb. nov.

Canonical: Amphiprora pseudoduplex

Authorship: (Osada & Kobayasi 1990)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Amphiprora pseudoduplex (Osada \u0026 Kobayasi, 1990) comb. nov.","normalized":"Amphiprora pseudoduplex (Osada \u0026 Kobayasi 1990)","canonical":{"stemmed":"Amphiprora pseudoduplex","simple":"Amphiprora pseudoduplex","full":"Amphiprora pseudoduplex"},"cardinality":2,"authorship":{"verbatim":"(Osada \u0026 Kobayasi, 1990)","normalized":"(Osada \u0026 Kobayasi 1990)","year":"1990","authors":["Osada","Kobayasi"],"originalAuth":{"authors":["Osada","Kobayasi"],"year":{"year":"1990"}}},"tail":" comb. nov.","details":{"species":{"genus":"Amphiprora","species":"pseudoduplex","authorship":{"verbatim":"(Osada \u0026 Kobayasi, 1990)","normalized":"(Osada \u0026 Kobayasi 1990)","year":"1990","authors":["Osada","Kobayasi"],"originalAuth":{"authors":["Osada","Kobayasi"],"year":{"year":"1990"}}}}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":23},{"wordType":"authorWord","start":25,"end":30},{"wordType":"authorWord","start":33,"end":41},{"wordType":"year","start":43,"end":47}],"id":"06b58578-d00c-5c90-b77a-bc2325694b51","parserVersion":"test_version"}
```

Name: Methanosarcina barkeri str. fusaro

Canonical: Methanosarcina barkeri

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Methanosarcina barkeri str. fusaro","normalized":"Methanosarcina barkeri","canonical":{"stemmed":"Methanosarcina barker","simple":"Methanosarcina barkeri","full":"Methanosarcina barkeri"},"cardinality":2,"tail":" str. fusaro","details":{"species":{"genus":"Methanosarcina","species":"barkeri"}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22}],"id":"b1d6747d-6aa3-5b7a-a8ed-7ca53c4b19ac","parserVersion":"test_version"}
```

Name: Arthopyrenia hyalospora (Nyl.) R.C. Harris comb. nov.

Canonical: Arthopyrenia hyalospora

Authorship: (Nyl.) R. C. Harris

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Arthopyrenia hyalospora (Nyl.) R.C. Harris comb. nov.","normalized":"Arthopyrenia hyalospora (Nyl.) R. C. Harris","canonical":{"stemmed":"Arthopyrenia hyalospor","simple":"Arthopyrenia hyalospora","full":"Arthopyrenia hyalospora"},"cardinality":2,"authorship":{"verbatim":"(Nyl.) R.C. Harris","normalized":"(Nyl.) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["R. C. Harris"]}},"tail":" comb. nov.","details":{"species":{"genus":"Arthopyrenia","species":"hyalospora","authorship":{"verbatim":"(Nyl.) R.C. Harris","normalized":"(Nyl.) R. C. Harris","authors":["Nyl.","R. C. Harris"],"originalAuth":{"authors":["Nyl."]},"combinationAuth":{"authors":["R. C. Harris"]}}}},"pos":[{"wordType":"genus","start":0,"end":12},{"wordType":"specificEpithet","start":13,"end":23},{"wordType":"authorWord","start":25,"end":29},{"wordType":"authorWord","start":31,"end":33},{"wordType":"authorWord","start":33,"end":35},{"wordType":"authorWord","start":36,"end":42}],"id":"2dcef387-edc3-55a1-9cfc-ee95200bff08","parserVersion":"test_version"}
```

Name: Acanthophis lancasteri WELLS & WELLINGTON (nomen nudum)

Canonical: Acanthophis lancasteri

Authorship: Wells & Wellington

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Author in upper case"}],"verbatim":"Acanthophis lancasteri WELLS \u0026 WELLINGTON (nomen nudum)","normalized":"Acanthophis lancasteri Wells \u0026 Wellington","canonical":{"stemmed":"Acanthophis lancaster","simple":"Acanthophis lancasteri","full":"Acanthophis lancasteri"},"cardinality":2,"authorship":{"verbatim":"WELLS \u0026 WELLINGTON","normalized":"Wells \u0026 Wellington","authors":["Wells","Wellington"],"originalAuth":{"authors":["Wells","Wellington"]}},"tail":" (nomen nudum)","details":{"species":{"genus":"Acanthophis","species":"lancasteri","authorship":{"verbatim":"WELLS \u0026 WELLINGTON","normalized":"Wells \u0026 Wellington","authors":["Wells","Wellington"],"originalAuth":{"authors":["Wells","Wellington"]}}}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":22},{"wordType":"authorWord","start":23,"end":28},{"wordType":"authorWord","start":31,"end":41}],"id":"aa527c3b-972e-56e9-9b8b-0c61c497422d","parserVersion":"test_version"}
```

Name: Acontias lineatus WAGLER 1830: 196 (nomen nudum)

Canonical: Acontias lineatus

Authorship: Wagler 1830

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Author in upper case"},{"quality":2,"warning":"Year with page info"}],"verbatim":"Acontias lineatus WAGLER 1830: 196 (nomen nudum)","normalized":"Acontias lineatus Wagler 1830","canonical":{"stemmed":"Acontias lineat","simple":"Acontias lineatus","full":"Acontias lineatus"},"cardinality":2,"authorship":{"verbatim":"WAGLER 1830: 196","normalized":"Wagler 1830","year":"1830","authors":["Wagler"],"originalAuth":{"authors":["Wagler"],"year":{"year":"1830"}}},"tail":" (nomen nudum)","details":{"species":{"genus":"Acontias","species":"lineatus","authorship":{"verbatim":"WAGLER 1830: 196","normalized":"Wagler 1830","year":"1830","authors":["Wagler"],"originalAuth":{"authors":["Wagler"],"year":{"year":"1830"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"authorWord","start":18,"end":24},{"wordType":"year","start":25,"end":29}],"id":"16afe3dd-7724-5dc0-817c-f6d138d27174","parserVersion":"test_version"}
```

Name: Akeratidae Nomen Nudum

Canonical: Akeratidae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Akeratidae Nomen Nudum","normalized":"Akeratidae","canonical":{"stemmed":"Akeratidae","simple":"Akeratidae","full":"Akeratidae"},"cardinality":1,"tail":" Nomen Nudum","details":{"uninomial":{"uninomial":"Akeratidae"}},"pos":[{"wordType":"uninomial","start":0,"end":10}],"id":"6bd60fba-9b78-5e4e-b904-dda976085fc7","parserVersion":"test_version"}
```

Name: Aster exilis Ell., nomen dubium

Canonical: Aster exilis

Authorship: Ell.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aster exilis Ell., nomen dubium","normalized":"Aster exilis Ell.","canonical":{"stemmed":"Aster exil","simple":"Aster exilis","full":"Aster exilis"},"cardinality":2,"authorship":{"verbatim":"Ell.","normalized":"Ell.","authors":["Ell."],"originalAuth":{"authors":["Ell."]}},"tail":", nomen dubium","details":{"species":{"genus":"Aster","species":"exilis","authorship":{"verbatim":"Ell.","normalized":"Ell.","authors":["Ell."],"originalAuth":{"authors":["Ell."]}}}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":12},{"wordType":"authorWord","start":13,"end":17}],"id":"00884bdf-ca19-5c07-8e48-e1adef987844","parserVersion":"test_version"}
```

Name: Abutilon avicennae Gaertn., nom. illeg.

Canonical: Abutilon avicennae

Authorship: Gaertn.

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Abutilon avicennae Gaertn., nom. illeg.","normalized":"Abutilon avicennae Gaertn.","canonical":{"stemmed":"Abutilon auicenn","simple":"Abutilon avicennae","full":"Abutilon avicennae"},"cardinality":2,"authorship":{"verbatim":"Gaertn.","normalized":"Gaertn.","authors":["Gaertn."],"originalAuth":{"authors":["Gaertn."]}},"tail":", nom. illeg.","details":{"species":{"genus":"Abutilon","species":"avicennae","authorship":{"verbatim":"Gaertn.","normalized":"Gaertn.","authors":["Gaertn."],"originalAuth":{"authors":["Gaertn."]}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":19,"end":26}],"id":"366d9605-0686-5072-b025-6c7b3695f086","parserVersion":"test_version"}
```

Name: Achillea bonarota nom. in herb.

Canonical: Achillea bonarota

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Achillea bonarota nom. in herb.","normalized":"Achillea bonarota","canonical":{"stemmed":"Achillea bonarot","simple":"Achillea bonarota","full":"Achillea bonarota"},"cardinality":2,"tail":" nom. in herb.","details":{"species":{"genus":"Achillea","species":"bonarota"}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17}],"id":"cae8ac71-b3c4-52f7-94cb-31e639081e0d","parserVersion":"test_version"}
```

Name: Aconitum napellus var. formosum (Rchb.) W. D. J. Koch (nom. ambig.)

Canonical: Aconitum napellus var. formosum

Authorship: (Rchb.) W. D. J. Koch

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Aconitum napellus var. formosum (Rchb.) W. D. J. Koch (nom. ambig.)","normalized":"Aconitum napellus var. formosum (Rchb.) W. D. J. Koch","canonical":{"stemmed":"Aconitum napell formos","simple":"Aconitum napellus formosum","full":"Aconitum napellus var. formosum"},"cardinality":3,"authorship":{"verbatim":"(Rchb.) W. D. J. Koch","normalized":"(Rchb.) W. D. J. Koch","authors":["Rchb.","W. D. J. Koch"],"originalAuth":{"authors":["Rchb."]},"combinationAuth":{"authors":["W. D. J. Koch"]}},"tail":" (nom. ambig.)","details":{"infraSpecies":{"genus":"Aconitum","species":"napellus","infraSpecies":[{"value":"formosum","rank":"var.","authorship":{"verbatim":"(Rchb.) W. D. J. Koch","normalized":"(Rchb.) W. D. J. Koch","authors":["Rchb.","W. D. J. Koch"],"originalAuth":{"authors":["Rchb."]},"combinationAuth":{"authors":["W. D. J. Koch"]}}}]}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":17},{"wordType":"rank","start":18,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":31},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWord","start":40,"end":42},{"wordType":"authorWord","start":43,"end":45},{"wordType":"authorWord","start":46,"end":48},{"wordType":"authorWord","start":49,"end":53}],"id":"9f79b2b3-cfd1-541a-9898-b60829134b11","parserVersion":"test_version"}
```

Name: Aesculus canadensis Hort. ex Lavallée

Canonical: Aesculus canadensis

Authorship: Hort. ex Lavallée

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Ex authors are not required"}],"verbatim":"Aesculus canadensis Hort. ex Lavallée","normalized":"Aesculus canadensis Hort. ex Lavallée","canonical":{"stemmed":"Aesculus canadens","simple":"Aesculus canadensis","full":"Aesculus canadensis"},"cardinality":2,"authorship":{"verbatim":"Hort. ex Lavallée","normalized":"Hort. ex Lavallée","authors":["Hort."],"originalAuth":{"authors":["Hort."],"exAuthors":{"authors":["Lavallée"]}}},"details":{"species":{"genus":"Aesculus","species":"canadensis","authorship":{"verbatim":"Hort. ex Lavallée","normalized":"Hort. ex Lavallée","authors":["Hort."],"originalAuth":{"authors":["Hort."],"exAuthors":{"authors":["Lavallée"]}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":19},{"wordType":"authorWord","start":20,"end":25},{"wordType":"authorWord","start":29,"end":37}],"id":"a1c7935f-26c2-5388-a1e2-b5a9508d70ef","parserVersion":"test_version"}
```

Name: × Dialaeliopsis hort.

Canonical: × Dialaeliopsis

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":2,"warning":"Named hybrid"}],"verbatim":"× Dialaeliopsis hort.","normalized":"× Dialaeliopsis","canonical":{"stemmed":"Dialaeliopsis","simple":"Dialaeliopsis","full":"× Dialaeliopsis"},"cardinality":1,"hybrid":"NAMED_HYBRID","tail":" hort.","details":{"uninomial":{"uninomial":"Dialaeliopsis"}},"pos":[{"wordType":"hybridChar","start":0,"end":1},{"wordType":"uninomial","start":2,"end":15}],"id":"5e0197df-26c1-55bc-a5c0-64376c599fa5","parserVersion":"test_version"}
```

### Misc annotations

Name: Velutina haliotoides (Linnaeus, 1758), sensu Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758), sensu Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":", sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":22,"end":30},{"wordType":"year","start":32,"end":36}],"id":"5efd63de-f4ec-55f1-bd5b-494988e58f9b","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 group

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 group","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" group","details":{"infraSpecies":{"genus":"Acarospora","species":"cratericola","infraSpecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":34},{"wordType":"authorWord","start":35,"end":40},{"wordType":"year","start":41,"end":45}],"id":"0f466e31-7e23-5320-ac7e-4c1026bc8af6","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 species group

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 species group","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" species group","details":{"infraSpecies":{"genus":"Acarospora","species":"cratericola","infraSpecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":34},{"wordType":"authorWord","start":35,"end":40},{"wordType":"year","start":41,"end":45}],"id":"a7684260-ed99-5d55-9a35-fd97b67e8933","parserVersion":"test_version"}
```

Name: Acarospora cratericola cratericola Shenk 1974 species complex

Canonical: Acarospora cratericola cratericola

Authorship: Shenk 1974

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Acarospora cratericola cratericola Shenk 1974 species complex","normalized":"Acarospora cratericola cratericola Shenk 1974","canonical":{"stemmed":"Acarospora cratericol cratericol","simple":"Acarospora cratericola cratericola","full":"Acarospora cratericola cratericola"},"cardinality":3,"authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}},"tail":" species complex","details":{"infraSpecies":{"genus":"Acarospora","species":"cratericola","infraSpecies":[{"value":"cratericola","authorship":{"verbatim":"Shenk 1974","normalized":"Shenk 1974","year":"1974","authors":["Shenk"],"originalAuth":{"authors":["Shenk"],"year":{"year":"1974"}}}}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":22},{"wordType":"infraspecificEpithet","start":23,"end":34},{"wordType":"authorWord","start":35,"end":40},{"wordType":"year","start":41,"end":45}],"id":"d227da04-7c89-50f7-8cf1-de09bc5aa903","parserVersion":"test_version"}
```

Name: Parus caeruleus species complex

Canonical: Parus caeruleus

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Parus caeruleus species complex","normalized":"Parus caeruleus","canonical":{"stemmed":"Parus caerule","simple":"Parus caeruleus","full":"Parus caeruleus"},"cardinality":2,"tail":" species complex","details":{"species":{"genus":"Parus","species":"caeruleus"}},"pos":[{"wordType":"genus","start":0,"end":5},{"wordType":"specificEpithet","start":6,"end":15}],"id":"f3752c09-242f-501c-8c8c-0feaf86c4693","parserVersion":"test_version"}
```

### Horticultural annotation

Name: Lachenalia tricolor var. nelsonii (ht.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lachenalia tricolor var. nelsonii (ht.) Baker","normalized":"Lachenalia tricolor var. nelsonii","canonical":{"stemmed":"Lachenalia tricolor nelsoni","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"tail":" (ht.) Baker","details":{"infraSpecies":{"genus":"Lachenalia","species":"tricolor","infraSpecies":[{"value":"nelsonii","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33}],"id":"0f7ce439-6b8d-53db-9ea3-82628f25b9bd","parserVersion":"test_version"}
```

Name: Lachenalia tricolor var. nelsonii (hort.) Baker

Canonical: Lachenalia tricolor var. nelsonii

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Lachenalia tricolor var. nelsonii (hort.) Baker","normalized":"Lachenalia tricolor var. nelsonii","canonical":{"stemmed":"Lachenalia tricolor nelsoni","simple":"Lachenalia tricolor nelsonii","full":"Lachenalia tricolor var. nelsonii"},"cardinality":3,"tail":" (hort.) Baker","details":{"infraSpecies":{"genus":"Lachenalia","species":"tricolor","infraSpecies":[{"value":"nelsonii","rank":"var."}]}},"pos":[{"wordType":"genus","start":0,"end":10},{"wordType":"specificEpithet","start":11,"end":19},{"wordType":"rank","start":20,"end":24},{"wordType":"infraspecificEpithet","start":25,"end":33}],"id":"cc118b05-14ff-5a42-8780-802f60eba565","parserVersion":"test_version"}
```

Name: Puya acris ht.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris ht.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" ht.","details":{"species":{"genus":"Puya","species":"acris"}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":10}],"id":"83c98b8e-f373-57df-92bf-5a39a56d9909","parserVersion":"test_version"}
```

Name: Puya acris hort.

Canonical: Puya acris

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Puya acris hort.","normalized":"Puya acris","canonical":{"stemmed":"Puya acr","simple":"Puya acris","full":"Puya acris"},"cardinality":2,"tail":" hort.","details":{"species":{"genus":"Puya","species":"acris"}},"pos":[{"wordType":"genus","start":0,"end":4},{"wordType":"specificEpithet","start":5,"end":10}],"id":"78228a5e-dcd3-58f9-bf21-b452c378f6ee","parserVersion":"test_version"}
```


### Not parsed OCR errors to get better precision/recall ratio

Name: Mom.alpium (Osbeck, 1778)

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Mom.alpium (Osbeck, 1778)","cardinality":0,"id":"f1452bcf-b779-5d98-bfc8-56455105e3f5","parserVersion":"test_version"}
```

### No parsing -- Genera abbreviated to 3 letters (too rare)

Name: Gen. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Gen. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma","cardinality":0,"id":"54d27b31-2fbd-56e1-85e1-1438970f8953","parserVersion":"test_version"}
```

Name: Genn. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Genn. et n. sp. Kaimatira Pumice Sand, Marton N ~1 Ma","cardinality":0,"id":"8edd1515-a4a1-52c5-ad1b-df7f112e68a9","parserVersion":"test_version"}
```

### No parsing -- incertae sedis

Name: Incertae sedis

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Incertae sedis","cardinality":0,"id":"74d54496-7f1c-52f8-81a9-9a9fb3a25ecb","parserVersion":"test_version"}
```

Name: </i>Hipponicidae<i> incertae sedis</i>

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Hipponicidae incertae sedis","cardinality":0,"id":"0b834a98-b696-5f7b-9d21-1aa17a43b040","parserVersion":"test_version"}
```

Name: incertae sedis

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"incertae sedis","cardinality":0,"id":"14f6de42-21d9-5e67-89cd-a05ebd974a1b","parserVersion":"test_version"}
```

Name: Inc.   sed.

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Inc.   sed.","cardinality":0,"id":"2e1319c9-a44b-531c-8964-67025bbf3b40","parserVersion":"test_version"}
```

Name: inc.sed.

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"inc.sed.","cardinality":0,"id":"dbb95e14-cebc-56a9-a1d2-a70d4b759e8d","parserVersion":"test_version"}
```

Name: inc.   sed.

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"inc.   sed.","cardinality":0,"id":"f5245bf6-a459-5602-9979-02ba9428cf17","parserVersion":"test_version"}
```

Name: Incertaesedis obscuricornis Fairmaire LMH 1893

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Incertaesedis obscuricornis Fairmaire LMH 1893","cardinality":0,"id":"2601fa55-350f-5591-a549-c558284d6e9e","parserVersion":"test_version"}
```

Name: Uropodoideaincertaesedis

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Uropodoideaincertaesedis","cardinality":0,"id":"3bf556bb-ea7c-536e-8b62-93ba329c559d","parserVersion":"test_version"}
```

### No parsing -- bacterium, Candidatus

Name: Acidobacteria bacterium

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Acidobacteria bacterium","cardinality":0,"id":"c982b4fd-c41a-5987-bcc8-989c4164b9ec","parserVersion":"test_version"}
```

Name: Acidimicrobiales bacterium JGI 01_E13

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Acidimicrobiales bacterium JGI 01_E13","cardinality":0,"id":"8b71a29b-4271-5a83-8a92-5dab1d9dc4c3","parserVersion":"test_version"}
```

Name: Acidobacterium ailaaui Myers & King, 2016

Canonical: Acidobacterium ailaaui

Authorship: Myers & King 2016

```json
{"parsed":true,"parseQuality":1,"verbatim":"Acidobacterium ailaaui Myers \u0026 King, 2016","normalized":"Acidobacterium ailaaui Myers \u0026 King 2016","canonical":{"stemmed":"Acidobacterium ailaau","simple":"Acidobacterium ailaaui","full":"Acidobacterium ailaaui"},"cardinality":2,"authorship":{"verbatim":"Myers \u0026 King, 2016","normalized":"Myers \u0026 King 2016","year":"2016","authors":["Myers","King"],"originalAuth":{"authors":["Myers","King"],"year":{"year":"2016"}}},"bacteria":"yes","details":{"species":{"genus":"Acidobacterium","species":"ailaaui","authorship":{"verbatim":"Myers \u0026 King, 2016","normalized":"Myers \u0026 King 2016","year":"2016","authors":["Myers","King"],"originalAuth":{"authors":["Myers","King"],"year":{"year":"2016"}}}}},"pos":[{"wordType":"genus","start":0,"end":14},{"wordType":"specificEpithet","start":15,"end":22},{"wordType":"authorWord","start":23,"end":28},{"wordType":"authorWord","start":31,"end":35},{"wordType":"year","start":37,"end":41}],"id":"b9f4555f-d2e0-5d40-acde-2b546a28a7fc","parserVersion":"test_version"}
```

Name: Candidatus Amesbacteria bacterium GW2011_GWC1_46_24

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Candidatus Amesbacteria bacterium GW2011_GWC1_46_24","cardinality":0,"id":"83382178-94bf-5bf3-a8c8-fdbca4af927c","parserVersion":"test_version"}
```

Name: Candidatus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Candidatus","cardinality":0,"id":"fb9138ac-ae7a-58c9-a912-d31d0a4eeed3","parserVersion":"test_version"}
```

Name: Candidatus Puniceispirillum Oh, Kwon, Kang, Kang, Lee, Kim & Cho, 2010

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Candidatus Puniceispirillum Oh, Kwon, Kang, Kang, Lee, Kim \u0026 Cho, 2010","cardinality":0,"id":"82fde2e2-8e50-5fd0-8ffe-96f34f85505b","parserVersion":"test_version"}
```

Name: Candidatus Halobonum

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Candidatus Halobonum","cardinality":0,"id":"289152c0-1042-5cac-a649-44314b25c857","parserVersion":"test_version"}
```

### No parsing -- 'Not', 'None', 'Unidentified'  phrases

Name: None recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"None recorded","cardinality":0,"id":"54d66439-b10d-50dc-a659-c9bce413ed5d","parserVersion":"test_version"}
```

Name: NONE recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"NONE recorded","cardinality":0,"id":"cedc6de2-aed6-58dc-904f-a14348588f8a","parserVersion":"test_version"}
```

Name: NoNe recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"NoNe recorded","cardinality":0,"id":"39682f61-d0d0-5dc0-bf57-b73ffb97b3ef","parserVersion":"test_version"}
```

Name: None

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"None","cardinality":0,"id":"8cf8696e-6ca6-5ec7-b441-e04a37ea751c","parserVersion":"test_version"}
```

Name: unidentified recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"unidentified recorded","cardinality":0,"id":"4c391bc1-d3f6-5e33-80df-262cbfb09dfe","parserVersion":"test_version"}
```

Name: UniDentiFied recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"UniDentiFied recorded","cardinality":0,"id":"57b55b46-c874-59ae-b3d8-2888d8a3bc1c","parserVersion":"test_version"}
```

Name: not recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"not recorded","cardinality":0,"id":"830df5b1-ef3b-5240-8ecf-4fd74c2fff72","parserVersion":"test_version"}
```

Name: NOT recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"NOT recorded","cardinality":0,"id":"52b51d9e-29db-561c-84ac-cd1592c762c1","parserVersion":"test_version"}
```

Name: Not recorded

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Not recorded","cardinality":0,"id":"025b92f4-2b2c-5593-a02b-66f121b0a42b","parserVersion":"test_version"}
```

Name: Not assigned

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Not assigned","cardinality":0,"id":"19bffdbe-f1c7-5d39-b7b6-3dc96a317c4b","parserVersion":"test_version"}
```

Name: Notassigned

Canonical: Notassigned

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Notassigned","normalized":"Notassigned","canonical":{"stemmed":"Notassigned","simple":"Notassigned","full":"Notassigned"},"cardinality":1,"details":{"uninomial":{"uninomial":"Notassigned"}},"pos":[{"wordType":"uninomial","start":0,"end":11}],"id":"8c07b58a-be4e-5c31-871b-cffe36b9860a","parserVersion":"test_version"}
```

Name: Unnamed clade

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Unnamed clade","cardinality":0,"id":"d510b662-0a4d-5678-a1a7-c58b20d25fa0","parserVersion":"test_version"}
```

Name: Unamed clade

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Unamed clade","cardinality":0,"id":"be6943d3-fa83-5e5d-9515-7cc339473d4d","parserVersion":"test_version"}
```

### No parsing -- genus with apostrophe

Name: Abbott's moray eel

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Abbott's moray eel","cardinality":0,"id":"6a870e4b-5cc5-5226-ac5d-b769521b640f","parserVersion":"test_version"}
```

Name: Chambers' twinpod

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Chambers' twinpod","cardinality":0,"id":"f109486d-9809-5196-b135-75f4cf9d7ef6","parserVersion":"test_version"}
```

Name: Columnea × Alladin's

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Columnea × Alladin's","cardinality":0,"id":"bc01a624-d49e-588d-b49d-253ac7e12939","parserVersion":"test_version"}
```

Name: Hawai'i silversword

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Hawai'i silversword","cardinality":0,"id":"f4ba0445-a5f2-525c-97ce-9316fe16e3cd","parserVersion":"test_version"}
```

### No parsing -- CamelCase 'genus' word

Name: PomaTomus

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"PomaTomus","cardinality":0,"id":"106ff909-e787-52b2-9139-25d0eb7d161e","parserVersion":"test_version"}
```

Name: DizygopUwa stosei

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"DizygopUwa stosei","cardinality":0,"id":"46511ef9-02d8-5f24-8364-b72df3e1494d","parserVersion":"test_version"}
```

Name: Oxytox[idae] Lindermann

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Oxytox[idae] Lindermann","cardinality":0,"id":"39a37760-d9f9-54d6-b49b-f6830e59f34e","parserVersion":"test_version"}
```

Name: ScarabaeinGCsp.

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"ScarabaeinGCsp.","cardinality":0,"id":"c84b775e-cc80-588f-b7bb-0094bab2c6a2","parserVersion":"test_version"}
```

### No parsing -- phytoplasma

Name: Alfalfa witches'-broom phytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Alfalfa witches'-broom phytoplasma","cardinality":0,"id":"b31676ed-c1ed-522c-8380-19a27af11e0d","parserVersion":"test_version"}
```

Name: Allium ampeloprasumphytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Allium ampeloprasumphytoplasma","cardinality":0,"id":"f84e58c5-8e49-5b2d-a4d0-4f1e538c8c7c","parserVersion":"test_version"}
```

Name: Alstroemeria sp. phytoplasma

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"Alstroemeria sp. phytoplasma","cardinality":0,"id":"5348845f-c94a-5c7e-bba1-307e4c07a42d","parserVersion":"test_version"}
```

### Names with spec., nov spec

Name: Lampona spec Platnick, 2000

Canonical: Lampona spec

Authorship: Platnick 2000

```json
{"parsed":true,"parseQuality":1,"verbatim":"Lampona spec Platnick, 2000","normalized":"Lampona spec Platnick 2000","canonical":{"stemmed":"Lampona spec","simple":"Lampona spec","full":"Lampona spec"},"cardinality":2,"authorship":{"verbatim":"Platnick, 2000","normalized":"Platnick 2000","year":"2000","authors":["Platnick"],"originalAuth":{"authors":["Platnick"],"year":{"year":"2000"}}},"details":{"species":{"genus":"Lampona","species":"spec","authorship":{"verbatim":"Platnick, 2000","normalized":"Platnick 2000","year":"2000","authors":["Platnick"],"originalAuth":{"authors":["Platnick"],"year":{"year":"2000"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":12},{"wordType":"authorWord","start":13,"end":21},{"wordType":"year","start":23,"end":27}],"id":"d05d7916-4868-57f6-a97b-c46886f29cd8","parserVersion":"test_version"}
```

Name: Gobiosoma spec (Ginsburg, 1939)

Canonical: Gobiosoma spec

Authorship: (Ginsburg 1939)

```json
{"parsed":true,"parseQuality":1,"verbatim":"Gobiosoma spec (Ginsburg, 1939)","normalized":"Gobiosoma spec (Ginsburg 1939)","canonical":{"stemmed":"Gobiosoma spec","simple":"Gobiosoma spec","full":"Gobiosoma spec"},"cardinality":2,"authorship":{"verbatim":"(Ginsburg, 1939)","normalized":"(Ginsburg 1939)","year":"1939","authors":["Ginsburg"],"originalAuth":{"authors":["Ginsburg"],"year":{"year":"1939"}}},"details":{"species":{"genus":"Gobiosoma","species":"spec","authorship":{"verbatim":"(Ginsburg, 1939)","normalized":"(Ginsburg 1939)","year":"1939","authors":["Ginsburg"],"originalAuth":{"authors":["Ginsburg"],"year":{"year":"1939"}}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":14},{"wordType":"authorWord","start":16,"end":24},{"wordType":"year","start":26,"end":30}],"id":"eb47c188-86fd-54c4-a058-48a980f9419f","parserVersion":"test_version"}
```

Name: Globigerina spec

Canonical: Globigerina spec

Authorship:

```json
{"parsed":true,"parseQuality":1,"verbatim":"Globigerina spec","normalized":"Globigerina spec","canonical":{"stemmed":"Globigerina spec","simple":"Globigerina spec","full":"Globigerina spec"},"cardinality":2,"details":{"species":{"genus":"Globigerina","species":"spec"}},"pos":[{"wordType":"genus","start":0,"end":11},{"wordType":"specificEpithet","start":12,"end":16}],"id":"4f8f7189-42a0-59e2-8d6f-67c3889673d9","parserVersion":"test_version"}
```

Name: Eunotia genuflexa Norpel-Schempp nov spec

Canonical: Eunotia genuflexa

Authorship: Norpel-Schempp

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Eunotia genuflexa Norpel-Schempp nov spec","normalized":"Eunotia genuflexa Norpel-Schempp","canonical":{"stemmed":"Eunotia genuflex","simple":"Eunotia genuflexa","full":"Eunotia genuflexa"},"cardinality":2,"authorship":{"verbatim":"Norpel-Schempp","normalized":"Norpel-Schempp","authors":["Norpel-Schempp"],"originalAuth":{"authors":["Norpel-Schempp"]}},"tail":" nov spec","details":{"species":{"genus":"Eunotia","species":"genuflexa","authorship":{"verbatim":"Norpel-Schempp","normalized":"Norpel-Schempp","authors":["Norpel-Schempp"],"originalAuth":{"authors":["Norpel-Schempp"]}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":17},{"wordType":"authorWord","start":18,"end":32}],"id":"4cc2a699-d38d-5337-8a44-ecc0f79ef138","parserVersion":"test_version"}
```

Name: Ctenotus spec.

Canonical: Ctenotus

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Ctenotus spec.","normalized":"Ctenotus","canonical":{"stemmed":"Ctenotus","simple":"Ctenotus","full":"Ctenotus"},"cardinality":1,"tail":" spec.","details":{"uninomial":{"uninomial":"Ctenotus"}},"pos":[{"wordType":"uninomial","start":0,"end":8}],"id":"991b9ee5-2f56-56e7-a29b-86c47a4901bb","parserVersion":"test_version"}
```

Name: Byrsophlebidae spec. 2

Canonical: Byrsophlebidae

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Byrsophlebidae spec. 2","normalized":"Byrsophlebidae","canonical":{"stemmed":"Byrsophlebidae","simple":"Byrsophlebidae","full":"Byrsophlebidae"},"cardinality":1,"tail":" spec. 2","details":{"uninomial":{"uninomial":"Byrsophlebidae"}},"pos":[{"wordType":"uninomial","start":0,"end":14}],"id":"3b07753b-71e2-5602-9a6e-bf91e672d834","parserVersion":"test_version"}
```

Name: Naviculadicta witkowskii LB & Metzeltin nov spec

Canonical: Naviculadicta witkowskii

Authorship: LB & Metzeltin

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Naviculadicta witkowskii LB \u0026 Metzeltin nov spec","normalized":"Naviculadicta witkowskii LB \u0026 Metzeltin","canonical":{"stemmed":"Naviculadicta witkowski","simple":"Naviculadicta witkowskii","full":"Naviculadicta witkowskii"},"cardinality":2,"authorship":{"verbatim":"LB \u0026 Metzeltin","normalized":"LB \u0026 Metzeltin","authors":["LB","Metzeltin"],"originalAuth":{"authors":["LB","Metzeltin"]}},"tail":" nov spec","details":{"species":{"genus":"Naviculadicta","species":"witkowskii","authorship":{"verbatim":"LB \u0026 Metzeltin","normalized":"LB \u0026 Metzeltin","authors":["LB","Metzeltin"],"originalAuth":{"authors":["LB","Metzeltin"]}}}},"pos":[{"wordType":"genus","start":0,"end":13},{"wordType":"specificEpithet","start":14,"end":24},{"wordType":"authorWord","start":25,"end":27},{"wordType":"authorWord","start":30,"end":39}],"id":"c4dd80b7-984b-51f8-a4ec-573b4b32358b","parserVersion":"test_version"}
```

### HTML tags and entities

Name: Velutina haliotoides (Linnaeus, 1758) <i>sensu</i> Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758) sensu Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":" sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":22,"end":30},{"wordType":"year","start":32,"end":36}],"id":"dc5dc538-23e7-5e4e-83b9-ba7fc4fb22a9","parserVersion":"test_version"}
```

Name: Velutina haliotoides (Linnaeus, 1758), <i>sensu</i> Fabricius, 1780

Canonical: Velutina haliotoides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"},{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Velutina haliotoides (Linnaeus, 1758), sensu Fabricius, 1780","normalized":"Velutina haliotoides (Linnaeus 1758)","canonical":{"stemmed":"Velutina haliotoid","simple":"Velutina haliotoides","full":"Velutina haliotoides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"tail":", sensu Fabricius, 1780","details":{"species":{"genus":"Velutina","species":"haliotoides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":20},{"wordType":"authorWord","start":22,"end":30},{"wordType":"year","start":32,"end":36}],"id":"5efd63de-f4ec-55f1-bd5b-494988e58f9b","parserVersion":"test_version"}
```

Name: <i>Velutina halioides</i> (Linnaeus, 1758)

Canonical: Velutina halioides

Authorship: (Linnaeus 1758)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Velutina halioides (Linnaeus, 1758)","normalized":"Velutina halioides (Linnaeus 1758)","canonical":{"stemmed":"Velutina halioid","simple":"Velutina halioides","full":"Velutina halioides"},"cardinality":2,"authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}},"details":{"species":{"genus":"Velutina","species":"halioides","authorship":{"verbatim":"(Linnaeus, 1758)","normalized":"(Linnaeus 1758)","year":"1758","authors":["Linnaeus"],"originalAuth":{"authors":["Linnaeus"],"year":{"year":"1758"}}}}},"pos":[{"wordType":"genus","start":0,"end":8},{"wordType":"specificEpithet","start":9,"end":18},{"wordType":"authorWord","start":20,"end":28},{"wordType":"year","start":30,"end":34}],"id":"2b3f5800-66c2-535d-8532-a281db56a1b7","parserVersion":"test_version"}
```

Name: Quadrella steyermarkii (Standl.) Iltis &amp; Cornejo

Canonical: Quadrella steyermarkii

Authorship: (Standl.) Iltis & Cornejo

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Quadrella steyermarkii (Standl.) Iltis \u0026 Cornejo","normalized":"Quadrella steyermarkii (Standl.) Iltis \u0026 Cornejo","canonical":{"stemmed":"Quadrella steyermarki","simple":"Quadrella steyermarkii","full":"Quadrella steyermarkii"},"cardinality":2,"authorship":{"verbatim":"(Standl.) Iltis \u0026 Cornejo","normalized":"(Standl.) Iltis \u0026 Cornejo","authors":["Standl.","Iltis","Cornejo"],"originalAuth":{"authors":["Standl."]},"combinationAuth":{"authors":["Iltis","Cornejo"]}},"details":{"species":{"genus":"Quadrella","species":"steyermarkii","authorship":{"verbatim":"(Standl.) Iltis \u0026 Cornejo","normalized":"(Standl.) Iltis \u0026 Cornejo","authors":["Standl.","Iltis","Cornejo"],"originalAuth":{"authors":["Standl."]},"combinationAuth":{"authors":["Iltis","Cornejo"]}}}},"pos":[{"wordType":"genus","start":0,"end":9},{"wordType":"specificEpithet","start":10,"end":22},{"wordType":"authorWord","start":24,"end":31},{"wordType":"authorWord","start":33,"end":38},{"wordType":"authorWord","start":41,"end":48}],"id":"3e33ac5a-3f95-5e61-878d-06318b05c545","parserVersion":"test_version"}
```

Name: Torymus bangalorensis (Mani &amp; Kurian, 1953)

Canonical: Torymus bangalorensis

Authorship: (Mani & Kurian 1953)

```json
{"parsed":true,"parseQuality":3,"qualityWarnings":[{"quality":3,"warning":"HTML tags or entities in the name"}],"verbatim":"Torymus bangalorensis (Mani \u0026 Kurian, 1953)","normalized":"Torymus bangalorensis (Mani \u0026 Kurian 1953)","canonical":{"stemmed":"Torymus bangalorens","simple":"Torymus bangalorensis","full":"Torymus bangalorensis"},"cardinality":2,"authorship":{"verbatim":"(Mani \u0026 Kurian, 1953)","normalized":"(Mani \u0026 Kurian 1953)","year":"1953","authors":["Mani","Kurian"],"originalAuth":{"authors":["Mani","Kurian"],"year":{"year":"1953"}}},"details":{"species":{"genus":"Torymus","species":"bangalorensis","authorship":{"verbatim":"(Mani \u0026 Kurian, 1953)","normalized":"(Mani \u0026 Kurian 1953)","year":"1953","authors":["Mani","Kurian"],"originalAuth":{"authors":["Mani","Kurian"],"year":{"year":"1953"}}}}},"pos":[{"wordType":"genus","start":0,"end":7},{"wordType":"specificEpithet","start":8,"end":21},{"wordType":"authorWord","start":23,"end":27},{"wordType":"authorWord","start":30,"end":36},{"wordType":"year","start":38,"end":42}],"id":"788f0e69-8093-5e89-8ddc-7eb115641304","parserVersion":"test_version"}
```

### Underscores instead of spaces

Name: Oxalis_barrelieri

Canonical: Oxalis barrelieri

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Oxalis_barrelieri","normalized":"Oxalis barrelieri","canonical":{"stemmed":"Oxalis barrelier","simple":"Oxalis barrelieri","full":"Oxalis barrelieri"},"cardinality":2,"details":{"species":{"genus":"Oxalis","species":"barrelieri"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"ad546700-9cae-50d3-9eaf-6adcbbb67bae","parserVersion":"test_version"}
```

Name: Pseudocercospora__dendrobii

Canonical: Pseudocercospora dendrobii

Authorship:

```json
{"parsed":true,"parseQuality":2,"qualityWarnings":[{"quality":2,"warning":"Multiple adjacent space characters"},{"quality":2,"warning":"Non-standard space characters"}],"verbatim":"Pseudocercospora__dendrobii","normalized":"Pseudocercospora dendrobii","canonical":{"stemmed":"Pseudocercospora dendrobi","simple":"Pseudocercospora dendrobii","full":"Pseudocercospora dendrobii"},"cardinality":2,"details":{"species":{"genus":"Pseudocercospora","species":"dendrobii"}},"pos":[{"wordType":"genus","start":0,"end":16},{"wordType":"specificEpithet","start":18,"end":27}],"id":"ae8a4688-2b2a-5974-81bf-1962838a9cbe","parserVersion":"test_version"}
```

Name:   Oxalis_barrelieri

Canonical:

Authorship:

```json
{"parsed":false,"parseQuality":0,"verbatim":"  Oxalis_barrelieri","cardinality":0,"id":"1c4bb48b-d134-54c8-bac1-6771d1f4c9c6","parserVersion":"test_version"}
```

Name: Oxalis barrelieri XXZ_21243

Canonical: Oxalis barrelieri

Authorship:

```json
{"parsed":true,"parseQuality":4,"qualityWarnings":[{"quality":4,"warning":"Unparsed tail"}],"verbatim":"Oxalis barrelieri XXZ_21243","normalized":"Oxalis barrelieri","canonical":{"stemmed":"Oxalis barrelier","simple":"Oxalis barrelieri","full":"Oxalis barrelieri"},"cardinality":2,"tail":" XXZ_21243","details":{"species":{"genus":"Oxalis","species":"barrelieri"}},"pos":[{"wordType":"genus","start":0,"end":6},{"wordType":"specificEpithet","start":7,"end":17}],"id":"8a722b76-cf2f-51d1-b60e-7f9236ddd189","parserVersion":"test_version"}
```