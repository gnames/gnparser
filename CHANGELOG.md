# Changelog

## Unreleased

- Add [#26]: get all parser rules to camelCase format.

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

[v0.5.1]: https://gitlab.com/gogna/gnparser/tree/v0.5.0

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
