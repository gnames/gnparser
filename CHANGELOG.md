# Changelog

## Unreleased

## [v0.7.0]

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

[v0.7.0]: https://gitlab.com/gogna/gnparser/compare/v0.6.0...v0.7.0
[v0.6.0]: https://gitlab.com/gogna/gnparser/compare/v0.5.1...v0.6.0
[v0.5.1]: https://gitlab.com/gogna/gnparser/tree/v0.5.1

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
