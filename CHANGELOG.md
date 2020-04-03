# Changelog

## Unreleased

- Add [#78]: Take into account `non-virus` names that look like virus names.

## [v0.13.1]

- Fix [#77]: Memory leak when used as clib.
- Fix [#76]: Non ASCII apostrophe does not show up in canonical.

## [v0.13.0]

- Add [#74]: Simple format output is now in CSV format.
- Add [#73]: Improve speed by using ragel's FSM instead of regex.
- Fix [#75]: Normalize subspecies to `subsp.` instead of `ssp.`.
- Fix [#72]: Surrogate detection by `gnparser.ParseToObject` method.

## [v0.12.0]

- Add [#71]: do not parse 'Unamed clade...'.
- Add [#69]: gnparser as a shared C library.
- Add: Make dynamic version using ldflags.
- Fix [#70]: parse 'Remera cvancarai' correctly.

## [v0.11.0]

- Add [#68]: add stemmed version of canonical form to outputs.
- Add: benchmarks to gnparser_test.go

## [v0.10.0]

- Add [#67]: field `authorship` of the name for JSON output
- Add [#66]: remove HTML tags during parsing instead of a separate step.
- Add [#61]: handle authors that end with a word "bis".
- Add [#60]: handle correctly deprecated ranks with greek letters.
- Add [#62]: parser breaks on ``Drepanolejeunea (Spruce) (Steph.)``.

## [v0.9.0]

- Add [#65]: gRPC is able to return a protobuf object now instead of JSON.
string (only for ParseArray function so far). The same protobuf object is now
also used by gnparser.ParseToObject function.
- Add [#64]: gRPC method ParseArray that cleans and parses an input from an
array of names instead of a stream.
- Add [#63]: abbreviation for `form` or `forma` is now `f.` instead of `fm.`.

## [v0.8.0]

- Add [#51]: strings like `Aus (Bus)` are parsed differently for ICN and ICZN
             names. If string inside of parenthesis matches known ICN author
             name is parsed as `Uninomial (Author)`, otherwise it is parsed
             as  `Aus subgen. Bus`.

## [v0.7.5]

- Add [#59]: method `ParseToObject` to avoid json in Go programs.
- Add [#58]: parse `Aus (Bus)` as `Uninomial (Author)` to prevent botanical
             authors appear as subgenera. We need a better solution for this.
- Add [#57]: warning in cases of an ambiguous `filius`.
- Fix [#56]: bug `Ambrysus-Stål, 1862` breaks parser.

## [v0.7.4]

- Add [#48]: transliteration of diacriticals.
- Add [#43]: notho- (hybrids) rank supported.
- Add [#52]: genera with hyphens with lower or upper char after hyphen.
- Add [#49]: multiple hyphens in specific epithet.

## [v0.7.3]

- Add [#54]: add cleaning functions to gRPC
- Add [#46]: add ``supg.`` rank
- Add [#45]: add ``natio`` rank (deprecated ICZN rank)
- Add [#44]: documentation for canonicalName fields
- Add [#42]: tests for command line app

## [v0.7.2]

- Add [#41]: parse/clean multiple names from standard input.

## [v0.7.1]

- Add [#40]: add names with missing parenthesis for combination authors.
- Fix: remove typo for Scala parser URL on the parser webpage.

## [v0.7.0]

- Add [#38]: docker image can do gRPC, REST, CLI
- Add [#37]: flag for cleanup HTML entities and tags,
             underscores are part of parsing.
- Add [#39]: documentation for contributors.
- Add [#31]: continuous integration.
- Add [#36]: substitute underscores to spaces for Newick format.
- Add [#34]: escape HTML entities, remove common tags.
- Add [#33]: Web-based user interface and REST API.

## [v0.6.0]

- Add [#35]: gRPC method to preserve order in output according to input
- Add [#30]: write inline and README documentation.
- Add [#29]: docker and dockerhub support.
- Add [#26]: get all parser rules to CamelCase format.

## [v0.5.1]

- Add: fix Makefile
- Add [#28]: non-ASCII apostrophe support.
- Add [#27]: agamosp. agamossp. agamovar. ranks.
- Add [#25]: reorganize output to be more readable and logical.
- Add [#24]: gRPC server for receiving name-strings and streaming back the
             parsed results.
- Add [#23]: Remove multiple years. Now name can have only one year.
- Add [#22]: Run the parser against 24 million names from global names index and
             fix found problems.
- Add [#21]: Rebuilds tests into ``test_data_new.txt`` file. It is important for
             making global changes in tests.
- Add [#20]: Pass all tests made for Scala gnparser. Tickets 1-19 are about
             approaching [#20].

## Footnotes

This document follows [changelog guidelines]

[v0.13.1]: https://gitlab.com/gogna/gnparser/compare/v0.13.0...v0.13.1
[v0.13.0]: https://gitlab.com/gogna/gnparser/compare/v0.12.0...v0.13.0
[v0.12.0]: https://gitlab.com/gogna/gnparser/compare/v0.11.0...v0.12.0
[v0.11.0]: https://gitlab.com/gogna/gnparser/compare/v0.10.0...v0.11.0
[v0.10.0]: https://gitlab.com/gogna/gnparser/compare/v0.9.0...v0.10.0
[v0.9.0]: https://gitlab.com/gogna/gnparser/compare/v0.8.0...v0.9.0
[v0.8.0]: https://gitlab.com/gogna/gnparser/compare/v0.7.5...v0.8.0
[v0.7.5]: https://gitlab.com/gogna/gnparser/compare/v0.7.4...v0.7.5
[v0.7.4]: https://gitlab.com/gogna/gnparser/compare/v0.7.3...v0.7.4
[v0.7.3]: https://gitlab.com/gogna/gnparser/compare/v0.7.2...v0.7.3
[v0.7.2]: https://gitlab.com/gogna/gnparser/compare/v0.7.1...v0.7.2
[v0.7.1]: https://gitlab.com/gogna/gnparser/compare/v0.7.0...v0.7.1
[v0.7.0]: https://gitlab.com/gogna/gnparser/compare/v0.6.0...v0.7.0
[v0.6.0]: https://gitlab.com/gogna/gnparser/compare/v0.5.1...v0.6.0
[v0.5.1]: https://gitlab.com/gogna/gnparser/tree/v0.5.1

[#77]: https://gitlab.com/gogna/gnparser/issues/77
[#76]: https://gitlab.com/gogna/gnparser/issues/76
[#75]: https://gitlab.com/gogna/gnparser/issues/75
[#74]: https://gitlab.com/gogna/gnparser/issues/74
[#73]: https://gitlab.com/gogna/gnparser/issues/73
[#72]: https://gitlab.com/gogna/gnparser/issues/72
[#71]: https://gitlab.com/gogna/gnparser/issues/71
[#70]: https://gitlab.com/gogna/gnparser/issues/70
[#69]: https://gitlab.com/gogna/gnparser/issues/69
[#68]: https://gitlab.com/gogna/gnparser/issues/68
[#67]: https://gitlab.com/gogna/gnparser/issues/67
[#66]: https://gitlab.com/gogna/gnparser/issues/66
[#65]: https://gitlab.com/gogna/gnparser/issues/65
[#64]: https://gitlab.com/gogna/gnparser/issues/64
[#63]: https://gitlab.com/gogna/gnparser/issues/63
[#62]: https://gitlab.com/gogna/gnparser/issues/62
[#61]: https://gitlab.com/gogna/gnparser/issues/61
[#60]: https://gitlab.com/gogna/gnparser/issues/60
[#59]: https://gitlab.com/gogna/gnparser/issues/59
[#58]: https://gitlab.com/gogna/gnparser/issues/58
[#57]: https://gitlab.com/gogna/gnparser/issues/57
[#56]: https://gitlab.com/gogna/gnparser/issues/56
[#55]: https://gitlab.com/gogna/gnparser/issues/55
[#54]: https://gitlab.com/gogna/gnparser/issues/54
[#52]: https://gitlab.com/gogna/gnparser/issues/52
[#49]: https://gitlab.com/gogna/gnparser/issues/49
[#48]: https://gitlab.com/gogna/gnparser/issues/48
[#46]: https://gitlab.com/gogna/gnparser/issues/46
[#45]: https://gitlab.com/gogna/gnparser/issues/45
[#44]: https://gitlab.com/gogna/gnparser/issues/44
[#43]: https://gitlab.com/gogna/gnparser/issues/43
[#42]: https://gitlab.com/gogna/gnparser/issues/42
[#41]: https://gitlab.com/gogna/gnparser/issues/41
[#40]: https://gitlab.com/gogna/gnparser/issues/40
[#39]: https://gitlab.com/gogna/gnparser/issues/39
[#38]: https://gitlab.com/gogna/gnparser/issues/38
[#37]: https://gitlab.com/gogna/gnparser/issues/37
[#36]: https://gitlab.com/gogna/gnparser/issues/36
[#35]: https://gitlab.com/gogna/gnparser/issues/35
[#34]: https://gitlab.com/gogna/gnparser/issues/34
[#33]: https://gitlab.com/gogna/gnparser/issues/33
[#32]: https://gitlab.com/gogna/gnparser/issues/32
[#31]: https://gitlab.com/gogna/gnparser/issues/31
[#30]: https://gitlab.com/gogna/gnparser/issues/30
[#29]: https://gitlab.com/gogna/gnparser/issues/29
[#28]: https://gitlab.com/gogna/gnparser/issues/28
[#27]: https://gitlab.com/gogna/gnparser/issues/27
[#26]: https://gitlab.com/gogna/gnparser/issues/26
[#25]: https://gitlab.com/gogna/gnparser/issues/25
[#24]: https://gitlab.com/gogna/gnparser/issues/24
[#23]: https://gitlab.com/gogna/gnparser/issues/23
[#22]: https://gitlab.com/gogna/gnparser/issues/22
[#21]: https://gitlab.com/gogna/gnparser/issues/21
[#20]: https://gitlab.com/gogna/gnparser/issues/20

[changelog guidelines]: https://github.com/olivierlacan/keep-a-changelog
