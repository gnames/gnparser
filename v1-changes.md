# Changes introduced in GNparser during migration from v0.x to v1.x

The [schema of the GNparser v1.x API][v1 spec] is described according to
the [OpenAPI specification].

This document describes changes in GNparser API and input/output format
introduced during the migration of the code from versions 0.x to versions 1.x.

`GNparser` follows the [Semantic Versioning] guidelines. According to
Semantic Versioning, versions 0.x.x are for the original development
of a program, and experimentations with input, output, and API of a program.

When developers understand API, input, and output better, it is time to switch
to v1.x versions. This kind of versions mean that their API and input/output
formats will be backward compatible with v.1.0.0. Only bug fixes and adding new
methods or fields in input/output are allowed. It is not allowed to make
changes to formats and methods that would break programs that depend on the
GNparser.

Moving to v1 for GNparser means that we will try hard to keep API and
input/output format backward compatible for the foreseeable future, and
introduce breaking changes only if there is an important reason to do
that. In case of the introduction of backward incompatibility with v1.0.0, we
will move to versions 2.x.

We use this migration from v.0.x to v1.x as an opportunity to break
compatibility with v0.x versions and use what we learned so far to
mint a stable API and input/output formats for GNparser.

This document describes what kind of changes are introduced during migration
from v0.x to v1.x.

## Table of Content

<!-- vim-markdown-toc GFM -->

* [REST API changes](#rest-api-changes)
* [Command line application interface changes](#command-line-application-interface-changes)
* [Changes in the Output Format](#changes-in-the-output-format)

<!-- vim-markdown-toc -->

## REST API changes

Here we use the main GNparser service `https://parser.globalnames.org` in the
examples.  If you run your own service, make the corresponding substitution of
the domain.

* Adding `https://parser.globalnames.org/api/v1` path.

    Old path `https://parser.globalnames.org/api` still works, but it will now
    run the most recent major version of the API, so if GNparser will move to
    v2, `/api/v1` will continue to serve the v.1.x API, while v2 API will be
    served at `/api/v2/` and `/api`.

* Adding `with_details` parameter.

    By default, the service will not send data that that most users do not
    need. The `details` field of the parsed data will be omitted, as well as
    details of a name's authorship. The `words` list that provides position
    and semantic meaning of every word in a name is excluded as well.

    If such details are important, use `with_details=true` parameter.

* Adding `csv` parameter.

    By default, service will continue to send back data in JSON format. To
    speed up traffic we introduce 'csv=true' parameter, that will serve data in
    a flat CSV format.

    `with_details` parameter is ignored if `csv=true`.

* Change in GET signature.

    v0.x API's GET method follows this example:

    `https://parser.globalnames.api?q=Pardosa|Bubo+bubo`

    v.1.x GET method follows this example:

    `https://parser.globalnames.org/api/Pardosa|Bubo+bubo?csv=true`

## Command line application interface changes

* Parsing large files or using CLI application with STDIN/STDOUT pipes.

    Parsing large files does not happen one name at a time anymore.  Names
    first are collected into "batches," and such batches are sent sequencially
    for a concurrent processing. The resulting parsed data of a batch are
    assembled in the same order as input and send back.

    This approach allows us to keep a very high speed of parsing for a huge
    number of names while keeping the same order of elements in the input and
    the output. Creating a batch takes a bit of overhead. Therefore the bigger
    the batch is, the less noticeable is the overhead. We decided that the
    default batch size will be 50,000 names.

    If the command line application `gnparser` is used inside of a Python,
    Java, Ruby, etc. program and involves STDIN/STDOUT pipes methods, such
    program usually needs to receive results one input entry at a time. To
    achieve that, set the size of a batch to 1:

    ```bash
    gnparser -b 1
    ```

* Parsed output is separated into `base` and `detailed` parts.

    For most use-cases, only a subset of parsed output is needed. Such output
    is served by default, while the more detailed output is omitted. `Base`
    output has a uniform schema, so it is easier to parse.

    If a user requires detailed output, it can be provided using a `--details`
    flag:

    ```bash
    gnparser -d
    ```

* Removal of HTML tags

    Quite often scientific name-strings contain HTML tags, for example
    `<i>Monochamus galloprovincialis</i> (Olivier, 1795)`. The tags are removed
    by default now.

    If user knows that an input never contains HTML tags, there is an option
    `--ignore_tags` that will speed up the parsing slightly.

    ```bash
    gnparser -i
    ```

## Changes in the Output Format

The [schema for GNparser v1.x output][v1 spec] is described according to
the [OpenAPI specification].

* Output has `base` and `details` parts.

    The reason for splitting the output is to decrease IO traffic and to make
    parsed data simpler to understand.

    `Base` part contains canonical forms, normalized and verbatim versions of
    the input name, cardinality, flags (hybrid, bacteria, virus, surrogate).
    The `base` part of the output has a uniform schema for any input, except
    that some fields are omitted when they cannot be generated. For example
    `canonical` field is not generated if the name cannot be parsed. Therefore,
    a missing field is the same as having 'false' or `null`.  The name's
    `authorship` field only contains 'stable' parts, like `verbatim` and
    `normalized` authorship strings, a year string, and a list of authors.

    `Details` part contains components that often can be ignored. It includes
    details of authorships, `details` field of a name, `words` list that
    includes start, end, and meaning of every parsed word in the name. The
    structure of most components in `details` part do not have a fixed schema
    so they change depending on input. The `words` list is an
    exception, its schema does not change.

* `Base` part changes:

  * `canonicalNames` in v.0.x changed to `canonical` in v1.x
  * `authorship` field changed from string to an object with `verbatim`,
    `normalized`, `year`, `authors` fields.
  * `surrogate` in v0.x was a boolean. In v1.x it is an optional field that
    describes the type of a parsed "surrogate" name.
  * `hybrid` v0.x was a boolean, in v1.x it describes the type of a parsed
    hybrid name.
  * `bacteria` in v0.x was a boolean, now it has 3 states:
     absent, 'maybe', 'yes'.
  * `nameStringID` in v0.x changed to `id` in v1.x.

* `Details` part changes

  * The `details` field now provides the type of a name, for example
    `uninomial`, `species`, `hybrid`.
  * The `positions` field is now called `words` and its elements
    include their word value.

[v1 spec]: https://app.swaggerhub.com/apis-docs/dimus/gnparser/1.0.0
[OpenAPI specification]: https://www.openapis.org/
[Semantic Versioning]: https://semver.org/

