# Changelog

## Unreleased

## [v1.11.3] - dt

- Add [#283]: parsed.Flatten function to generate 'flat' version of parsed data.
- Add: update upload-artifacts in Actions to v4.
- Add [#279]: add `unr.` rank.

## [v1.11.2] - 2025-02-21 Fri

- Add: update modules.
- Add [#282]: remove varning for botanical names with `ex` authors.
- Add [#280]: prol. botanical rank.

## [v1.11.1] - 2024-11-11 Mon

- Add: minor improvements in web-interface.
- Add: make web-interface remember nomenclatural code picked in the previous
  GET query.

## [v1.11.0] - 2024-11-11 Mon

- Add [#265]: add optional nomenclatural code option to parse names with
  an ambiguity according to a particular code.

## [v1.10.4] - 2024-11-07 Thu

- Add [#269]: switch to slog from zerologs.
- Fix [#271]: distinguish between `ex` and `in`.
- Fix [#270]: missing verbatim authorship for names that look similar to
  combination uninomial in ICZN.
- Fix [#268]: if botanical author looks like a combination uninomial,
  make sure its characters are not normalized.

## [v1.10.3] - 2024-10-11 Fri

- Fix [#266]: remove author from species detail for named hybrids.

## [v1.10.2] - 2024-07-30 Tue

- Add [#264]: more exceptions.

## [v1.10.1] - 2024-06-05 Wed

- Add [#263]: add more exceptions with speciific epithets like
  "complex", "do", "spec".

## [v1.10.0] - 2024-06-04 Tue

- Add [#260]: add `candidatus` field for parsed data.
- Add [#232]: parse names like `subgen. Psammophrynopsis Koch, 1953`.

## [v1.9.2] - 2024-05-01 Wed

- Add [#261]: a constructor to return a pool of a given size of gnparsers.

## [v1.9.1] - 2023-10-13 Fri

- Add: update modules.
- Fix [#259]: allow diacritics in any UTF-8 normalization form.
- Fix [#258]: allow authors with 2 dashes in the name.
- Fix [#256]: fix normalization where a misplacced year changes
  the year of original authors.

## [v1.9.0] - 2023-10-12 Thu

- Add: restore backward compatibility by creating a new flag
  `--species-group-cut`.

## [v1.8.0] - 2023-10-11 Wed

- Add [#255]: normalize stemmed canonical of `Aus bus bus` to `Aus bus`.
  **WARNING** this creates some backward incompatibility.
- Add: sorting uses `slices` package.

## [v1.7.5] - 2023-09-26 Tue

- Add: CSV and TSV files provide now verbatim authorship instead of normalized
  one.
- Add: a few more "termination words"
- Fix [#254]: treat `fa` as forma.
- Fix [#253]: process `dem` as an author word for `Von dem Bush` and like.
- Fix [#251]: do not process `y` as `and` for `Rafael Arango y Molina`.
- Fix [#249]: allow `cf` at the end of the strings, cf for infraspecies.
- Fix [#248]: do not escape double quotes for TSV output.
- Fix [#246]: ignore `ms` at the end of the strings.

## [v1.7.4] - 2023-08-22 Tue

- Fix [#243]: parse correctly `Nassa pagoda var. acuta P. P. Carpenter, 1857`.

## [v1.7.3] - 2023-06-17 Sat

- Add [#241]: allow comma before `ex` authors.

## [v1.7.2] - 2023-03-09 Thu

- Add [#240]: add `tr.` `subtr.` as ranks for combo-uninomials.

## [v1.7.1] - 2023-03-07 Tue

- Add: upgrade all modules.

## [v1.7.0] - 2023-03-07 Tue

- Add [#238]: stem takes in account -ii suffix, `macdonaldii` -> `macdonald`.

## [v1.6.9] - 2022-11-10 Thu

- Add [#237]: detect and normalize non-breaking hyphens. In case if other
  non-typical hythens will appear, they will be dealt the same way.

## [v1.6.8] - 2022-10-01 Sat

- Add: update all modules.

## [v1.6.7] - 2022-08-22 Mon

- Add [#231]: more edge cases.
- Add [#230]: Take into account `mihi` annotation.

## [v1.6.6] - 2022-05-15 Sun

- Add [#224]: Creation of Nix packages for gnparser.

## [v1.6.5] - 2022-03-21 Mon

- Add [#223]: Use PEG parser for preprocess instead of RegEx.
  This approach gives 15-17% speed increase.

## [v1.6.4] - 2022-03-19 Sat

- Add [#224]: Parse correctly italian authors with `degli`.

## [v1.6.3] - 2022-02-08 Tue

- Add [#222]: Improve logs for NSQ, switch to zerologs library.

## [v1.6.2] - 2022-02-04 Fri

- Fix [#221]: No parsing for names with `cyanobacterium`.
- Fix [#220]: `Crenarchaeote enrichment culture clone` should stop parsing
  at `enrichment`.
- Fix [#219]: filter out `complex` word during preprocessing for names like
  `Aegla uruguayana complex`.

## [v1.6.1] - 2022-02-01 Sat

- Add: use NSQ logger from sfgrp/lognsq

## [v1.6.0] - 2022-01-22 Sat

- Add [#218]: enable/disable logs for web-services, allow logs aggregation
  with NSQd.

## [v1.5.7] - 2021-11-26 Fri

- Fix: parsed.NormalizeByType preserves period char.

## [v1.5.6] - 2021-11-21 Sun

- Add [#212]: Set year from 'ex' authorship as a year of a name.
  Add 'ex' authors to list of all authors.

- Add [#211]: PR [#214] by @tobymarsden, general approach for `non-...`
  specific epithets.

- Add [#208]: PR [#210] by @tobymarsden, option to preserve diaereses.

- Fix [#213]: Stop generating space between `Mc`, `Mac` and the rest of an
  an author name.

## [v1.5.5] - 2021-11-17 Wed

- Add [#207]: PR [#209] by @tobymarsden, fix parsing of names with `nudum`
  specific epithet.

## [v1.5.4] - 2021-11-14 Sun

- Add: different approach for normalize-by-type for words.
- Add [#205]: allow genera starting with De-, Eu-, Le-, Ne- (by @tobymarsden).
- Add [#203]: allow up to 2 dashes in genera (by @tobymarsden).

## [v1.5.3] - 2021-11-13 Sat

- Add [#202]: add NormalizeMore function for Word.

## [v1.5.2] - 2021-11-10 Wed

- Add [#200]: support for 'div.' rank in uninomial combinations.
- Add [#199]: fixes for several names that were not parsed correctly.
- Add [#198]: parse "Solanum tuberosum wila-k`oyu".
- Add [#97]: do not parse "Cyanophage".
- Add [#85]: parse names with a dagger character.
- Add [#84]: parse "Muscicapa randi Amadon & duPont, 1970".
- Add [#83]: parse authors like 'Laverde-R.'.

## [v1.5.1] - 2021-11-01 Mon

- Add [#191]: support for ambiguous specific epithets

## [v1.5.0] - 2021-10-22 Fri

- Add [#194]: support for cultivars' graft-chymeras (courtesy of @tobymarsden)

## [v1.4.2] - 2021-10-21 Thu

- Add [#196]: parse authors with prefix 'ver'

## [v1.4.1] - 2021-10-07 Thu

- Fix [#195]: parse multinomials where authorshp is not separated by space.

## [v1.4.0] - 2021-09-4 Sat

- Add [#193]: add TSV format for output.
- Add [#190]: support prefixes `do` and `de los` for authors.
- Add [#187]: support `ter` suffix for authors.
- Add [#186]: support non-ASCII apostrophe in authors.

## [v1.3.3] - 2021-09-11 Wed

- Add [#176]: refactoring of hybrid sign treatment (use PEG instead of
  RegEx for normalizing `x`, `X`, and `×`.
- Add [#183]: stop parsing after `nec`, `non`, `fide`, `vide`, treat
  `ms in` as `in` or `ex` for exAuthors.
- Add [#182]: support for authors with prefixes `ten`, `delle`, `dos`.

## [v1.3.2] - 2021-08-02 Mon

- Add [#182]: support `Do`, `Oo`, `Nu` 2-letter genera.
- Add [#53]: exceptions to annotations (`Bottaria nudum` for example).
- Fix: names where sp epithet starts with `cf` can be parsed now.

## [v1.3.1] - 2021-07-17 Sat

- Add [#180]: Zenodo DOI.

## [v1.3.0] - 2021-06-29 Tue

- Add [#179]: cultivars info to README.
- Add [#178]: parse cultivars via REST API.
- Add [#177]: parse botanical cultivars via web.
- Add [#173]: cultivars parsing #174 @tobymarsden.
- Add [#172]: authors initials with a dash like "B.-E.van Wyk".
- Add: tests for cultivars (Toby Marsden)

- Fix [#174]: Hybrid character is missed or wrong in details'
  `Words` section.

## [v1.2.0] - 2021-04-08 Thu

- Add [#169]: option to capitalize first letter of name-strings.
- Add [#166]: support 'fm.' as 'forma'.

## [v1.1.0] - 2021-03-21 Sun

- Add [#163]: support bacterial `Candidatus` names.
- Add [#162]: show PEG AST tree for debugging.
- Add [#161]: add automatic tools dependency.
- Add [#160]: use embed feature of Go v1.16.

## [v1.0.13] - 2021-02-23 Tue

- Add: limit nightly builds to master only.
- Fix [#159]: POST method contains w18rong URL.

## [v1.0.12] - 2021-02-21 Sun

- Add [#154]: parse names with ambiguous `f.` as forma if there
  is a space between authr and `f.`. If there is
  no space, parse as `filius`. Give ambiguity
  warning in both cases.
- Add: PHP example from @barotto about using pipes with gnparser.

## [v1.0.11] - 2021-02-20 Sat

- Fix [#153]: flags `csv=false` and `with_details=false`
  trigger opposite behavior.

## [v1.0.10] - 2021-02-19 Fri

- Add [#152]: change auto-prereleases from nightly to on master submit.
- Add [#151]: do not parse names with `(endo|ecto)?symbiont`.
- Add [#150]: ignore serovar/serotype in bacerital names.
- Add [#149]: support abbreviated subgenus (`Aus (B.) cus`).

## [v1.0.9] - 2021-02-17 Wed

- Add [#146]: unordered flag.
- Add [#145]: better CI/build actions, add nightly binaries.
- Fix [#144]: remove configuration file as it creates more problems than solves.

## [v1.0.8] - 2021-02-15 Mon

- Add: remove config message for CLI app.
- Add: ldflags `-s -w` to decrease binary size.
- Fix: header does not show in CSV format for stream.

## [v1.0.7] - 2021-02-14 Sun

- Add [#143]: `quiet` flag to suppress showing progress output.

- Fix [#142]: stream waits until certain names number is equal the batch size.
- Fix [#141]: config file is not created.

## [v1.0.6] - 2021-02-04 Thu

- Add: update version handling, readme.

## [v1.0.5] - 2021-02-01 Mon

- Add: remove gnlib package.
- Add [#140]: remove config package.

## [v1.0.4] - 2021-01-23 Sat

- Add: cleanup constructor methods names.

## [v1.0.3] - 2021-01-23 Sat

- Add [#139]: make package names less abstract.

## [v1.0.2] - 2021-01-22 Fri

- Fix [#137]: add correct VerbatimID for HTML-containing names.

## [v1.0.1] - 2021-01-20 Wed

- Add [#136]: Man page
- Add [#100]: Switch continuous integration to use GitHub Actions.
- Add [#129]: Make c-binding usable for biodiversity parser.

- Fix [#135]: Changes: SubGenus->Subgenus, InfraSpecies->Infraspecies

## [v1.0.0] - 2021-01-19 Tue

- Add [#127]: Update documentation to v1.0.0.
- Add [#122]: Implement parsing as a stream in addition to batch parsing.
- Add [#126]: Update c-binding to v1.0.0.
- Add [#131]: Add parameters "with_details" and "csv" to REST API.
- Add [#134]: Transoform "positions" section to "words" section.
- Add [#128]: Add more examples to OpenAPI specification.
- Add [#125]: Describe changes from v0.x to 1.x.
- Add [#132]: Add context.Context to control lifespan of `go routines`.
- Add [#115]: Migrate tests from ginkgo to plain tests.
- Add [#109]: Move `web` package to `io`.
- Add [#124]: Document warnings for each quality category.
- Add [#121]: Convert `package` parser to use interfaces.
- Add [#120]: CLI app for newly created functionality.
- Add [#119]: Formatted output for `output.Parsed`.
- Add [#117]: Convert failed parsing results to `output.Parsed`.
- Add [#114]: Convert parsing result to `output.Parsed`.
- Add [#118]: Add `Verbatim` and `Year` fields to the root of `Authorship`.
- Add [#107]: Move `grammar` package to `entity` and rename to `parser`.
- Add [#110]: Move `stemmer` to `entity`.
- Add [#113]: Move `str` package to `entity`.
- Add [#112]: Move `preprocess` package to `entity`.
- Add [#105]: Move `fs` package to `io`.
- Add [#111]: Move `dict` package to `io`.
- Add [#106]: Describe main use-case via interface.
- Add [#104]: Add configuration package.
- Add [#103]: Create an output.Parsed object that can be used in Go and as JSON.
- Add [#101]: Start using gnlib where it makes sense.
- Add [#99]: Move code to GitHub and change links accordingly.
- Add [#95]: Remove dependency on gRPC and protobuf.

## [v0.14.4] - 2020-12-15 Tue

- Add [#96]: Do not parse names starting with "Candidatus".
- Add [#93]: Parse 'y' (Spanish '&') as an author separator.

## [v0.14.3] - 2020-12-13 Sun

- Add [#95]: Remove make dependency on gRPC tooling.
- Add [#94]: Do not parse names with "bacterium" "epithet.

## [v0.14.2] - 2020-05-12 Tue

- Add [#90]: Allow `ß` in names.
- Add [#89]: Support `subspec.` as a rank.
- Add [#82]: Support authors with prefix `zu`.

## [v0.14.1] - 2020-05-07 Thu

- Fix: Change web API from default to Compact format to get correct API output.

## [v0.14.0] - 2020-05-07 Thu

- Add [#81]: Add year range in format "1888/89".
- Add [#80]: Add Cardinality to parser outputs.
- Add [#79]: Make CSV the default format for CLI.
- Add [#78]: Take into account `non-virus` names that look like virus names.

## [v0.13.1] - 2020-03-05 Thu

- Fix [#77]: Memory leak when used as clib.
- Fix [#76]: Non ASCII apostrophe does not show up in canonical.

## [v0.13.0] - 2020-02-12 Wed

- Add [#74]: Simple format output is now in CSV format.
- Add [#73]: Improve speed by using ragel's FSM instead of regex.
- Fix [#75]: Normalize subspecies to `subsp.` instead of `ssp.`.
- Fix [#72]: Surrogate detection by `gnparser.ParseToObject` method.

## [v0.12.0] - 2019-11-18 Mon

- Add [#71]: do not parse 'Unnamed clade...'.
- Add [#69]: gnparser as a shared C library.
- Add: Make dynamic version using ldflags.
- Fix [#70]: parse 'Remera cvancarai' correctly.

## [v0.11.0] - 2019-10-24 Thu

- Add [#68]: add stemmed version of canonical form to outputs.
- Add: benchmarks to gnparser_test.go

## [v0.10.0] - 2019-09-10 Tue

- Add [#67]: field `authorship` of the name for JSON output
- Add [#66]: remove HTML tags during parsing instead of a separate step.
- Add [#61]: handle authors that end with a word "bis".
- Add [#60]: handle correctly deprecated ranks with Greek letters.
- Fix [#62]: parser breaks on `Drepanolejeunea (Spruce) (Steph.)`.

## [v0.9.0] - 2019-08-16 Fri

- Add [#65]: gRPC is able to return a protobuf object now instead of JSON.
  string (only for ParseArray function so far). The same protobuf object is now
  also used by gnparser.ParseToObject function.
- Add [#64]: gRPC method ParseArray that cleans and parses an input from an
  array of names instead of a stream.
- Add [#63]: abbreviation for `form` or `forma` is now `f.` instead of `fm.`.

## [v0.8.0] - 2019-04-10 Wed

- Add [#51]: strings like `Aus (Bus)` are parsed differently for ICN and ICZN
  names. If string inside of parenthesis matches known ICN author
  name is parsed as `Uninomial (Author)`, otherwise it is parsed
  as `Aus subgen. Bus`.

## [v0.7.5] - 2019-03-31 Sun

- Add [#59]: method `ParseToObject` to avoid JSON in Go programs.
- Add [#58]: parse `Aus (Bus)` as `Uninomial (Author)` to prevent botanical
  authors appear as subgenera. We need a better solution for this.
- Add [#57]: warning in cases of an ambiguous `filius`.
- Fix [#56]: bug `Ambrysus-Stål, 1862` breaks parser.

## [v0.7.4] - 2019-02-12 Tue

- Add [#48]: transliteration of diacriticals.
- Add [#43]: notho- (hybrids) rank supported.
- Add [#52]: genera with hyphens with lower or upper char after hyphen.
- Add [#49]: multiple hyphens in specific epithet.

## [v0.7.3] - 2019-02-04 Mon

- Add [#54]: add cleaning functions to gRPC
- Add [#46]: add `supg.` rank
- Add [#45]: add `natio` rank (deprecated ICZN rank)
- Add [#44]: documentation for canonicalName fields
- Add [#42]: tests for command line app

## [v0.7.2] - 2019-02-01 Fri

- Add [#41]: parse/clean multiple names from standard input.

## [v0.7.1] - 2019-01-24 Thu

- Add [#40]: add names with missing parenthesis for combination authors.
- Fix: remove typo for Scala parser URL on the parser web-page.

## [v0.7.0] - 2019-01-23 Wed

- Add [#38]: docker image can do gRPC, REST, CLI
- Add [#37]: flag for cleanup HTML entities and tags,
  underscores are part of parsing.
- Add [#39]: documentation for contributors.
- Add [#31]: continuous integration.
- Add [#36]: substitute underscores to spaces for Newick format.
- Add [#34]: escape HTML entities, remove common tags.
- Add [#33]: Web-based user interface and REST API.

## [v0.6.0] - 2019-01-16 Wed

- Add [#35]: gRPC method to preserve order in output according to input
- Add [#30]: write inline and README documentation.
- Add [#29]: docker and dockerhub support.
- Add [#26]: get all parser rules to CamelCase format.

## [v0.5.1] - 2019-01-15 Tue

- Add: fix Makefile
- Add [#28]: non-ASCII apostrophe support.
- Add [#27]: agamosp. agamossp. agamovar. ranks.
- Add [#25]: reorganize output to be more readable and logical.
- Add [#24]: gRPC server for receiving name-strings and streaming back the
  parsed results.
- Add [#23]: Remove multiple years. Now name can have only one year.
- Add [#22]: Run the parser against 24 million names from global names index and
  fix found problems.
- Add [#21]: Rebuilds tests into `test_data_new.txt` file. It is important for
  making global changes in tests.
- Add [#20]: Pass all tests made for Scala gnparser. Tickets 1-19 are about
  approaching [#20].

## Footnotes

This document follows [changelog guidelines]


[v1.11.3]: https://github.com/gnames/gnparser/compare/v1.11.2...v1.11.3
[v1.11.2]: https://github.com/gnames/gnparser/compare/v1.11.1...v1.11.2
[v1.11.1]: https://github.com/gnames/gnparser/compare/v1.11.0...v1.11.1
[v1.11.0]: https://github.com/gnames/gnparser/compare/v1.10.4...v1.11.0
[v1.10.4]: https://github.com/gnames/gnparser/compare/v1.10.3...v1.10.4
[v1.10.3]: https://github.com/gnames/gnparser/compare/v1.10.2...v1.10.3
[v1.10.2]: https://github.com/gnames/gnparser/compare/v1.10.1...v1.10.2
[v1.10.1]: https://github.com/gnames/gnparser/compare/v1.10.0...v1.10.1
[v1.10.0]: https://github.com/gnames/gnparser/compare/v1.9.2...v1.10.0
[v1.9.2]: https://github.com/gnames/gnparser/compare/v1.9.1...v1.9.2
[v1.9.1]: https://github.com/gnames/gnparser/compare/v1.9.0...v1.9.1
[v1.9.0]: https://github.com/gnames/gnparser/compare/v1.8.0...v1.9.0
[v1.8.0]: https://github.com/gnames/gnparser/compare/v1.7.5...v1.8.0
[v1.7.5]: https://github.com/gnames/gnparser/compare/v1.7.4...v1.7.5
[v1.7.4]: https://github.com/gnames/gnparser/compare/v1.7.3...v1.7.4
[v1.7.3]: https://github.com/gnames/gnparser/compare/v1.7.2...v1.7.3
[v1.7.2]: https://github.com/gnames/gnparser/compare/v1.7.1...v1.7.2
[v1.7.1]: https://github.com/gnames/gnparser/compare/v1.7.0...v1.7.1
[v1.7.0]: https://github.com/gnames/gnparser/compare/v1.6.9...v1.7.0
[v1.6.9]: https://github.com/gnames/gnparser/compare/v1.6.8...v1.6.9
[v1.6.8]: https://github.com/gnames/gnparser/compare/v1.6.7...v1.6.8
[v1.6.7]: https://github.com/gnames/gnparser/compare/v1.6.6...v1.6.7
[v1.6.6]: https://github.com/gnames/gnparser/compare/v1.6.5...v1.6.6
[v1.6.5]: https://github.com/gnames/gnparser/compare/v1.6.4...v1.6.5
[v1.6.4]: https://github.com/gnames/gnparser/compare/v1.6.3...v1.6.4
[v1.6.3]: https://github.com/gnames/gnparser/compare/v1.6.2...v1.6.3
[v1.6.2]: https://github.com/gnames/gnparser/compare/v1.6.1...v1.6.2
[v1.6.1]: https://github.com/gnames/gnparser/compare/v1.6.0...v1.6.1
[v1.6.0]: https://github.com/gnames/gnparser/compare/v1.5.7...v1.6.0
[v1.5.7]: https://github.com/gnames/gnparser/compare/v1.5.6...v1.5.7
[v1.5.6]: https://github.com/gnames/gnparser/compare/v1.5.5...v1.5.6
[v1.5.5]: https://github.com/gnames/gnparser/compare/v1.5.4...v1.5.5
[v1.5.4]: https://github.com/gnames/gnparser/compare/v1.5.3...v1.5.4
[v1.5.3]: https://github.com/gnames/gnparser/compare/v1.5.2...v1.5.3
[v1.5.2]: https://github.com/gnames/gnparser/compare/v1.5.1...v1.5.2
[v1.5.1]: https://github.com/gnames/gnparser/compare/v1.5.0...v1.5.1
[v1.5.0]: https://github.com/gnames/gnparser/compare/v1.4.2...v1.5.0
[v1.4.2]: https://github.com/gnames/gnparser/compare/v1.4.1...v1.4.2
[v1.4.1]: https://github.com/gnames/gnparser/compare/v1.4.0...v1.4.1
[v1.4.0]: https://github.com/gnames/gnparser/compare/v1.3.3...v1.4.0
[v1.3.3]: https://github.com/gnames/gnparser/compare/v1.3.2...v1.3.3
[v1.3.2]: https://github.com/gnames/gnparser/compare/v1.3.1...v1.3.2
[v1.3.1]: https://github.com/gnames/gnparser/compare/v1.3.0...v1.3.1
[v1.3.0]: https://github.com/gnames/gnparser/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/gnames/gnparser/compare/v1.1.0...v1.2.0
[v1.1.0]: https://github.com/gnames/gnparser/compare/v1.0.14...v1.1.0
[v1.0.14]: https://github.com/gnames/gnparser/compare/v1.0.13...v1.0.14
[v1.0.13]: https://github.com/gnames/gnparser/compare/v1.0.12...v1.0.13
[v1.0.12]: https://github.com/gnames/gnparser/compare/v1.0.11...v1.0.12
[v1.0.11]: https://github.com/gnames/gnparser/compare/v1.0.10...v1.0.11
[v1.0.10]: https://github.com/gnames/gnparser/compare/v1.0.9...v1.0.10
[v1.0.9]: https://github.com/gnames/gnparser/compare/v1.0.8...v1.0.9
[v1.0.8]: https://github.com/gnames/gnparser/compare/v1.0.7...v1.0.8
[v1.0.7]: https://github.com/gnames/gnparser/compare/v1.0.6...v1.0.7
[v1.0.6]: https://github.com/gnames/gnparser/compare/v1.0.5...v1.0.6
[v1.0.5]: https://github.com/gnames/gnparser/compare/v1.0.4...v1.0.5
[v1.0.4]: https://github.com/gnames/gnparser/compare/v1.0.3...v1.0.4
[v1.0.3]: https://github.com/gnames/gnparser/compare/v1.0.2...v1.0.3
[v1.0.2]: https://github.com/gnames/gnparser/compare/v1.0.1...v1.0.2
[v1.0.1]: https://github.com/gnames/gnparser/compare/v1.0.0...v1.0.1
[v1.0.0]: https://github.com/gnames/gnparser/compare/v0.14.4...v1.0.0
[v0.14.4]: https://github.com/gnames/gnparser/compare/v0.14.3...v0.14.4
[v0.14.3]: https://github.com/gnames/gnparser/compare/v0.14.2...v0.14.3
[v0.14.2]: https://github.com/gnames/gnparser/compare/v0.14.1...v0.14.2
[v0.14.1]: https://github.com/gnames/gnparser/compare/v0.14.0...v0.14.1
[v0.14.0]: https://github.com/gnames/gnparser/compare/v0.13.1...v0.14.0
[v0.13.1]: https://github.com/gnames/gnparser/compare/v0.13.0...v0.13.1
[v0.13.0]: https://github.com/gnames/gnparser/compare/v0.12.0...v0.13.0
[v0.12.0]: https://github.com/gnames/gnparser/compare/v0.11.0...v0.12.0
[v0.11.0]: https://github.com/gnames/gnparser/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/gnames/gnparser/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/gnames/gnparser/compare/v0.8.0...v0.9.0
[v0.8.0]: https://github.com/gnames/gnparser/compare/v0.7.5...v0.8.0
[v0.7.5]: https://github.com/gnames/gnparser/compare/v0.7.4...v0.7.5
[v0.7.4]: https://github.com/gnames/gnparser/compare/v0.7.3...v0.7.4
[v0.7.3]: https://github.com/gnames/gnparser/compare/v0.7.2...v0.7.3
[v0.7.2]: https://github.com/gnames/gnparser/compare/v0.7.1...v0.7.2
[v0.7.1]: https://github.com/gnames/gnparser/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/gnames/gnparser/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/gnames/gnparser/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/gnames/gnparser/tree/v0.5.1
[#290]: https://github.com/gnames/gnparser/issues/290
[#289]: https://github.com/gnames/gnparser/issues/289
[#288]: https://github.com/gnames/gnparser/issues/288
[#287]: https://github.com/gnames/gnparser/issues/287
[#286]: https://github.com/gnames/gnparser/issues/286
[#285]: https://github.com/gnames/gnparser/issues/285
[#284]: https://github.com/gnames/gnparser/issues/284
[#283]: https://github.com/gnames/gnparser/issues/283
[#282]: https://github.com/gnames/gnparser/issues/282
[#281]: https://github.com/gnames/gnparser/issues/281
[#280]: https://github.com/gnames/gnparser/issues/280
[#279]: https://github.com/gnames/gnparser/issues/279
[#278]: https://github.com/gnames/gnparser/issues/278
[#277]: https://github.com/gnames/gnparser/issues/277
[#276]: https://github.com/gnames/gnparser/issues/276
[#275]: https://github.com/gnames/gnparser/issues/275
[#274]: https://github.com/gnames/gnparser/issues/274
[#273]: https://github.com/gnames/gnparser/issues/273
[#272]: https://github.com/gnames/gnparser/issues/272
[#271]: https://github.com/gnames/gnparser/issues/271
[#270]: https://github.com/gnames/gnparser/issues/270
[#269]: https://github.com/gnames/gnparser/issues/269
[#268]: https://github.com/gnames/gnparser/issues/268
[#267]: https://github.com/gnames/gnparser/issues/267
[#266]: https://github.com/gnames/gnparser/issues/266
[#265]: https://github.com/gnames/gnparser/issues/265
[#264]: https://github.com/gnames/gnparser/issues/264
[#263]: https://github.com/gnames/gnparser/issues/263
[#262]: https://github.com/gnames/gnparser/issues/262
[#261]: https://github.com/gnames/gnparser/issues/261
[#260]: https://github.com/gnames/gnparser/issues/260
[#259]: https://github.com/gnames/gnparser/issues/259
[#258]: https://github.com/gnames/gnparser/issues/258
[#257]: https://github.com/gnames/gnparser/issues/257
[#256]: https://github.com/gnames/gnparser/issues/256
[#255]: https://github.com/gnames/gnparser/issues/255
[#254]: https://github.com/gnames/gnparser/issues/254
[#253]: https://github.com/gnames/gnparser/issues/253
[#252]: https://github.com/gnames/gnparser/issues/252
[#251]: https://github.com/gnames/gnparser/issues/251
[#250]: https://github.com/gnames/gnparser/issues/250
[#249]: https://github.com/gnames/gnparser/issues/249
[#248]: https://github.com/gnames/gnparser/issues/248
[#247]: https://github.com/gnames/gnparser/issues/247
[#246]: https://github.com/gnames/gnparser/issues/246
[#245]: https://github.com/gnames/gnparser/issues/245
[#244]: https://github.com/gnames/gnparser/issues/244
[#243]: https://github.com/gnames/gnparser/issues/243
[#242]: https://github.com/gnames/gnparser/issues/242
[#241]: https://github.com/gnames/gnparser/issues/241
[#240]: https://github.com/gnames/gnparser/issues/240
[#239]: https://github.com/gnames/gnparser/issues/239
[#238]: https://github.com/gnames/gnparser/issues/238
[#237]: https://github.com/gnames/gnparser/issues/237
[#236]: https://github.com/gnames/gnparser/issues/236
[#235]: https://github.com/gnames/gnparser/issues/235
[#234]: https://github.com/gnames/gnparser/issues/234
[#233]: https://github.com/gnames/gnparser/issues/233
[#232]: https://github.com/gnames/gnparser/issues/232
[#231]: https://github.com/gnames/gnparser/issues/231
[#230]: https://github.com/gnames/gnparser/issues/230
[#229]: https://github.com/gnames/gnparser/issues/229
[#228]: https://github.com/gnames/gnparser/issues/228
[#227]: https://github.com/gnames/gnparser/issues/227
[#226]: https://github.com/gnames/gnparser/issues/226
[#225]: https://github.com/gnames/gnparser/issues/225
[#224]: https://github.com/gnames/gnparser/issues/224
[#223]: https://github.com/gnames/gnparser/issues/223
[#222]: https://github.com/gnames/gnparser/issues/222
[#221]: https://github.com/gnames/gnparser/issues/221
[#220]: https://github.com/gnames/gnparser/issues/220
[#219]: https://github.com/gnames/gnparser/issues/219
[#218]: https://github.com/gnames/gnparser/issues/218
[#217]: https://github.com/gnames/gnparser/issues/217
[#216]: https://github.com/gnames/gnparser/issues/216
[#215]: https://github.com/gnames/gnparser/issues/215
[#214]: https://github.com/gnames/gnparser/issues/214
[#213]: https://github.com/gnames/gnparser/issues/213
[#212]: https://github.com/gnames/gnparser/issues/212
[#211]: https://github.com/gnames/gnparser/issues/211
[#210]: https://github.com/gnames/gnparser/issues/210
[#209]: https://github.com/gnames/gnparser/issues/209
[#208]: https://github.com/gnames/gnparser/issues/208
[#207]: https://github.com/gnames/gnparser/issues/207
[#206]: https://github.com/gnames/gnparser/issues/206
[#205]: https://github.com/gnames/gnparser/issues/205
[#204]: https://github.com/gnames/gnparser/issues/204
[#203]: https://github.com/gnames/gnparser/issues/203
[#202]: https://github.com/gnames/gnparser/issues/202
[#201]: https://github.com/gnames/gnparser/issues/201
[#200]: https://github.com/gnames/gnparser/issues/200
[#199]: https://github.com/gnames/gnparser/issues/199
[#198]: https://github.com/gnames/gnparser/issues/198
[#197]: https://github.com/gnames/gnparser/issues/197
[#196]: https://github.com/gnames/gnparser/issues/196
[#195]: https://github.com/gnames/gnparser/issues/195
[#194]: https://github.com/gnames/gnparser/issues/194
[#193]: https://github.com/gnames/gnparser/issues/193
[#192]: https://github.com/gnames/gnparser/issues/192
[#191]: https://github.com/gnames/gnparser/issues/191
[#190]: https://github.com/gnames/gnparser/issues/190
[#189]: https://github.com/gnames/gnparser/issues/189
[#188]: https://github.com/gnames/gnparser/issues/188
[#187]: https://github.com/gnames/gnparser/issues/187
[#186]: https://github.com/gnames/gnparser/issues/186
[#185]: https://github.com/gnames/gnparser/issues/185
[#184]: https://github.com/gnames/gnparser/issues/184
[#183]: https://github.com/gnames/gnparser/issues/183
[#182]: https://github.com/gnames/gnparser/issues/182
[#181]: https://github.com/gnames/gnparser/issues/181
[#180]: https://github.com/gnames/gnparser/issues/180
[#179]: https://github.com/gnames/gnparser/issues/179
[#178]: https://github.com/gnames/gnparser/issues/178
[#177]: https://github.com/gnames/gnparser/issues/177
[#176]: https://github.com/gnames/gnparser/issues/176
[#175]: https://github.com/gnames/gnparser/issues/175
[#174]: https://github.com/gnames/gnparser/issues/174
[#173]: https://github.com/gnames/gnparser/issues/173
[#172]: https://github.com/gnames/gnparser/issues/172
[#171]: https://github.com/gnames/gnparser/issues/171
[#170]: https://github.com/gnames/gnparser/issues/170
[#169]: https://github.com/gnames/gnparser/issues/169
[#168]: https://github.com/gnames/gnparser/issues/168
[#167]: https://github.com/gnames/gnparser/issues/167
[#166]: https://github.com/gnames/gnparser/issues/166
[#165]: https://github.com/gnames/gnparser/issues/165
[#164]: https://github.com/gnames/gnparser/issues/164
[#163]: https://github.com/gnames/gnparser/issues/163
[#162]: https://github.com/gnames/gnparser/issues/162
[#161]: https://github.com/gnames/gnparser/issues/161
[#160]: https://github.com/gnames/gnparser/issues/160
[#159]: https://github.com/gnames/gnparser/issues/159
[#158]: https://github.com/gnames/gnparser/issues/158
[#157]: https://github.com/gnames/gnparser/issues/157
[#156]: https://github.com/gnames/gnparser/issues/156
[#155]: https://github.com/gnames/gnparser/issues/155
[#154]: https://github.com/gnames/gnparser/issues/154
[#153]: https://github.com/gnames/gnparser/issues/153
[#152]: https://github.com/gnames/gnparser/issues/152
[#151]: https://github.com/gnames/gnparser/issues/151
[#150]: https://github.com/gnames/gnparser/issues/150
[#149]: https://github.com/gnames/gnparser/issues/149
[#148]: https://github.com/gnames/gnparser/issues/148
[#147]: https://github.com/gnames/gnparser/issues/147
[#146]: https://github.com/gnames/gnparser/issues/146
[#145]: https://github.com/gnames/gnparser/issues/145
[#144]: https://github.com/gnames/gnparser/issues/144
[#143]: https://github.com/gnames/gnparser/issues/143
[#142]: https://github.com/gnames/gnparser/issues/142
[#141]: https://github.com/gnames/gnparser/issues/141
[#140]: https://github.com/gnames/gnparser/issues/140
[#139]: https://github.com/gnames/gnparser/issues/139
[#138]: https://github.com/gnames/gnparser/issues/138
[#137]: https://github.com/gnames/gnparser/issues/137
[#136]: https://github.com/gnames/gnparser/issues/136
[#135]: https://github.com/gnames/gnparser/issues/135
[#134]: https://github.com/gnames/gnparser/issues/134
[#133]: https://github.com/gnames/gnparser/issues/133
[#132]: https://github.com/gnames/gnparser/issues/132
[#131]: https://github.com/gnames/gnparser/issues/131
[#130]: https://github.com/gnames/gnparser/issues/130
[#129]: https://github.com/gnames/gnparser/issues/129
[#128]: https://github.com/gnames/gnparser/issues/128
[#127]: https://github.com/gnames/gnparser/issues/127
[#126]: https://github.com/gnames/gnparser/issues/126
[#125]: https://github.com/gnames/gnparser/issues/125
[#124]: https://github.com/gnames/gnparser/issues/124
[#123]: https://github.com/gnames/gnparser/issues/123
[#122]: https://github.com/gnames/gnparser/issues/122
[#121]: https://github.com/gnames/gnparser/issues/121
[#120]: https://github.com/gnames/gnparser/issues/120
[#119]: https://github.com/gnames/gnparser/issues/119
[#118]: https://github.com/gnames/gnparser/issues/118
[#117]: https://github.com/gnames/gnparser/issues/117
[#116]: https://github.com/gnames/gnparser/issues/116
[#115]: https://github.com/gnames/gnparser/issues/115
[#114]: https://github.com/gnames/gnparser/issues/114
[#113]: https://github.com/gnames/gnparser/issues/113
[#112]: https://github.com/gnames/gnparser/issues/112
[#111]: https://github.com/gnames/gnparser/issues/111
[#110]: https://github.com/gnames/gnparser/issues/110
[#109]: https://github.com/gnames/gnparser/issues/109
[#108]: https://github.com/gnames/gnparser/issues/108
[#107]: https://github.com/gnames/gnparser/issues/107
[#106]: https://github.com/gnames/gnparser/issues/106
[#105]: https://github.com/gnames/gnparser/issues/105
[#104]: https://github.com/gnames/gnparser/issues/104
[#103]: https://github.com/gnames/gnparser/issues/103
[#102]: https://github.com/gnames/gnparser/issues/102
[#101]: https://github.com/gnames/gnparser/issues/101
[#100]: https://github.com/gnames/gnparser/issues/100
[#99]: https://github.com/gnames/gnparser/issues/99
[#98]: https://github.com/gnames/gnparser/issues/98
[#97]: https://github.com/gnames/gnparser/issues/97
[#96]: https://github.com/gnames/gnparser/issues/96
[#95]: https://github.com/gnames/gnparser/issues/95
[#94]: https://github.com/gnames/gnparser/issues/94
[#93]: https://github.com/gnames/gnparser/issues/93
[#92]: https://github.com/gnames/gnparser/issues/92
[#91]: https://github.com/gnames/gnparser/issues/91
[#90]: https://github.com/gnames/gnparser/issues/90
[#89]: https://github.com/gnames/gnparser/issues/89
[#88]: https://github.com/gnames/gnparser/issues/88
[#87]: https://github.com/gnames/gnparser/issues/87
[#86]: https://github.com/gnames/gnparser/issues/86
[#85]: https://github.com/gnames/gnparser/issues/85
[#84]: https://github.com/gnames/gnparser/issues/84
[#83]: https://github.com/gnames/gnparser/issues/83
[#82]: https://github.com/gnames/gnparser/issues/82
[#81]: https://github.com/gnames/gnparser/issues/81
[#80]: https://github.com/gnames/gnparser/issues/80
[#79]: https://github.com/gnames/gnparser/issues/79
[#78]: https://github.com/gnames/gnparser/issues/78
[#77]: https://github.com/gnames/gnparser/issues/77
[#76]: https://github.com/gnames/gnparser/issues/76
[#75]: https://github.com/gnames/gnparser/issues/75
[#74]: https://github.com/gnames/gnparser/issues/74
[#73]: https://github.com/gnames/gnparser/issues/73
[#72]: https://github.com/gnames/gnparser/issues/72
[#71]: https://github.com/gnames/gnparser/issues/71
[#70]: https://github.com/gnames/gnparser/issues/70
[#69]: https://github.com/gnames/gnparser/issues/69
[#68]: https://github.com/gnames/gnparser/issues/68
[#67]: https://github.com/gnames/gnparser/issues/67
[#66]: https://github.com/gnames/gnparser/issues/66
[#65]: https://github.com/gnames/gnparser/issues/65
[#64]: https://github.com/gnames/gnparser/issues/64
[#63]: https://github.com/gnames/gnparser/issues/63
[#62]: https://github.com/gnames/gnparser/issues/62
[#61]: https://github.com/gnames/gnparser/issues/61
[#60]: https://github.com/gnames/gnparser/issues/60
[#59]: https://github.com/gnames/gnparser/issues/59
[#58]: https://github.com/gnames/gnparser/issues/58
[#57]: https://github.com/gnames/gnparser/issues/57
[#56]: https://github.com/gnames/gnparser/issues/56
[#55]: https://github.com/gnames/gnparser/issues/55
[#54]: https://github.com/gnames/gnparser/issues/54
[#52]: https://github.com/gnames/gnparser/issues/52
[#49]: https://github.com/gnames/gnparser/issues/49
[#48]: https://github.com/gnames/gnparser/issues/48
[#46]: https://github.com/gnames/gnparser/issues/46
[#45]: https://github.com/gnames/gnparser/issues/45
[#44]: https://github.com/gnames/gnparser/issues/44
[#43]: https://github.com/gnames/gnparser/issues/43
[#42]: https://github.com/gnames/gnparser/issues/42
[#41]: https://github.com/gnames/gnparser/issues/41
[#40]: https://github.com/gnames/gnparser/issues/40
[#39]: https://github.com/gnames/gnparser/issues/39
[#38]: https://github.com/gnames/gnparser/issues/38
[#37]: https://github.com/gnames/gnparser/issues/37
[#36]: https://github.com/gnames/gnparser/issues/36
[#35]: https://github.com/gnames/gnparser/issues/35
[#34]: https://github.com/gnames/gnparser/issues/34
[#33]: https://github.com/gnames/gnparser/issues/33
[#32]: https://github.com/gnames/gnparser/issues/32
[#31]: https://github.com/gnames/gnparser/issues/31
[#30]: https://github.com/gnames/gnparser/issues/30
[#29]: https://github.com/gnames/gnparser/issues/29
[#28]: https://github.com/gnames/gnparser/issues/28
[#27]: https://github.com/gnames/gnparser/issues/27
[#26]: https://github.com/gnames/gnparser/issues/26
[#25]: https://github.com/gnames/gnparser/issues/25
[#24]: https://github.com/gnames/gnparser/issues/24
[#23]: https://github.com/gnames/gnparser/issues/23
[#22]: https://github.com/gnames/gnparser/issues/22
[#21]: https://github.com/gnames/gnparser/issues/21
[#20]: https://github.com/gnames/gnparser/issues/20
[changelog guidelines]: https://github.com/olivierlacan/keep-a-changelog
