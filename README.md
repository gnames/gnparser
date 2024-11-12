# Global Names Parser: GNparser written in Go

[![DOI](https://zenodo.org/badge/320967495.svg)](https://zenodo.org/badge/latestdoi/320967495)

Try `GNparser` [online][parser-web].

Try [GNparser with OpenRefine][OpenRefine]

`GNparser` splits scientific names into their semantic elements with an
associated meta information. Parsing is indispensable for matching names
from different data sources, because it can normalize different lexical
variants of names to the same `canonical form`.

This parser, written in Go, is the 3rd iteration of the project. The
first, [biodiversity], had been written in Ruby, the second, [also
gnparser][gnparser-scala], had been written in Scala. This project is
now a substitution for the other two. Scala project is in an archived state,
[biodiversity] now uses Go code for parsing. All three projects were developed
as a part of [Global Names Architecture Project][gna].

To use `GNparser` as a command line tool under Windows, Mac or Linux,
download the [latest release][releases], uncompress it, and copy `gnparser`
binary somewhere in your PATH. On a Mac you might also need to go to
`System Preferences` and security panel select `Allow from other
developers`. Then, after running `gnparser`, click 'Yes' in a dialog box
allowing to run a program from an "unregistered developer".

```bash
tar xvf gnparser-v1.0.0-linux.tar.gz
sudo cp gnparser /usr/local/bin
# for CSV output
gnparser "Homo sapiens Linnaeus"
# for TSV output
gnparser -f tsv "Homo sapiens Linnaeus"
# for JSON output
gnparser -f compact "Homo sapiens Linnaeus"
gnparser -f compact "Homo sapiens Linnaeus" | jq
# or
gnparser -f pretty "Homo sapiens Linnaeus"
gnparser -h
```

<!-- vim-markdown-toc GFM -->

* [Citing](#citing)
* [Introduction](#introduction)
* [Speed](#speed)
* [Features](#features)
* [Use Cases](#use-cases)
  * [Getting the simplest possible canonical form](#getting-the-simplest-possible-canonical-form)
  * [Quickly partition names by the type](#quickly-partition-names-by-the-type)
  * [Normalizing name-strings](#normalizing-name-strings)
  * [Removing authorship from the middle of the name](#removing-authorship-from-the-middle-of-the-name)
  * [Figuring out if names are well-formed](#figuring-out-if-names-are-well-formed)
  * [Creating stable GUIDs for name-strings](#creating-stable-guids-for-name-strings)
  * [Assembling canonical forms etc. from original spelling](#assembling-canonical-forms-etc-from-original-spelling)
* [Tutorials](#tutorials)
* [Installation](#installation)
  * [Install with Homebrew (Mac OS X, Linux)](#install-with-homebrew-mac-os-x-linux)
  * [Linux or Mac OS X](#linux-or-mac-os-x)
  * [Windows](#windows)
  * [Install with Go](#install-with-go)
* [Usage](#usage)
  * [Command Line](#command-line)
  * [Pipes](#pipes)
  * [R language package](#r-language-package)
  * [Ruby Gem](#ruby-gem)
  * [Node.js](#nodejs)
  * [Usage as a REST API Interface or Web-based User Graphical Interface](#usage-as-a-rest-api-interface-or-web-based-user-graphical-interface)
    * [Enabling logs for GNparser's web-service](#enabling-logs-for-gnparsers-web-service)
  * [Use as a Docker image](#use-as-a-docker-image)
  * [Use as a library in Go](#use-as-a-library-in-go)
  * [Use as a shared C library](#use-as-a-shared-c-library)
* [Parsing ambiguities](#parsing-ambiguities)
  * [Names with `filius` (ICN code)](#names-with-filius-icn-code)
  * [Names with subgenus (ICZN code) and genus author (ICN code)](#names-with-subgenus-iczn-code-and-genus-author-icn-code)
* [Authors](#authors)
* [Contributors](#contributors)
* [References](#references)
* [License](#license)

<!-- vim-markdown-toc -->

## Citing

[Zenodo DOI] can be used to cite GNparser

## Introduction

Global Names Parser or `GNparser` is a program written in Go for breaking up
scientific names into their elements. It uses [peg] -- a Parsing
Expression Grammar (PEG) tool.

Many other parsing algorithms for scientific names use regular expressions.
This approach works well for extracting canonical forms in simple cases.
However, for complex scientific names and to parse scientific names into
all semantic elements, regular expressions often fail, unable to overcome
the recursive nature of data embedded in names. By contrast, `GNparser`
is able to deal with the most complex scientific name-strings.

`GNparser` takes a name-string like `Drosophila (Sophophora) melanogaster
Meigen, 1830` and returns parsed components in `CSV`, `TSV` or `JSON` format.
The parsing of scientific names might become surprisingly complex and the
`GNparser's` [test file] is a good source of information about the parser's
capabilities, its input and output.

`GNparser` reached a stable v1. [Differences between v1 and v0][v0diff]

## Speed

Number of names parsed per second on an AMD Ryzen 7 5800H CPU
(8 cores, 16 threads), GNparser v1.3.0:

```bash
gnparser 1_000_000_names.txt -j 200 > /dev/null
```

| Threads | names/sec |
| ------- | --------- |
| 1       | 9,000     |
| 2       | 19,000    |
| 4       | 35,000    |
| 8       | 56,000    |
| 16      | 82,000    |
| 100     | 107,000   |
| 200     | 111,000   |

For simplest output Go `GNparser` is roughly 2 times faster than Scala
`GNparser` and about 100 times faster than pure Ruby implementation. For
JSON formats the parser is approximately 8 times faster than Scala one, due to
more efficient JSON conversion.

## Features

- Fastest parser ever.
- Very easy to install, just placing executable somewhere in the PATH is
  sufficient.
- Extracts all elements from a name, not only canonical forms.
- Works with very complex scientific names, including hybrid formulas.
- Includes RESTful service and interactive web interface.
- Can run as a command line application.
- Can be used as a library in Go projects.
- Can be scaled to many CPUs and computers (if 250 millions names an
  hour is not enough).
- Calculates a stable UUID version 5 ID from the content of a string.
- Provides C-binding to incorporate parser to other [languages][biodiversity].

## Use Cases

### Getting the simplest possible canonical form

Canonical forms of a scientific name are the latinized components without
annotations, authors or dates. They are great for matching lexical variants
of names. Three versions of canonical forms are included:

| Canonical | Example                         | Use                                                    |
| :-------- | :------------------------------ | :----------------------------------------------------- |
| -         | _Spiraea alba var. alba_ Du Roi | Best for disambiguation, but has many lexical variants |
| Full      | _Spiraea alba var. alba_        | Presentation, infraspecies disambiguation              |
| Simple    | _Spiraea alba alba_             | Name matching, presentation                            |
| Stem      | _Spiraea alb alb_               | Best for matching fem./masc. inconsistencies           |

The `canonicalName -> full` is good for presentation, as it keeps more
details.

The `canonicalName -> simple` field is good for matching names from different
sources, because sometimes dataset curators omit hybrid sign in named hybrids,
or remove ranks for infraspecific epithets.

The `canonicalName -> stem` field normalizes `simple` canonical form even
further. It allows to match names with inconsistent gender suffixes in specific
epithets (for example _alba_ vs. _albus_). The normalization is done according
to stemming rules for Latin language described in [Schinke R et al (1996)]. For
example letters `j` are converted to `i`, letters `v` are converted to `u`, and
suffixes are removed from the specific and infraspecific epithets.

If you only care mostly about canonical form of a name you can use default
`--format csv` flag with command line tool.

CSV/TSV output has the following fields:

| Field           | Meaning                                       |
| --------------- | --------------------------------------------- |
| Id              | UUID v5 generated out of Verbatim             |
| Verbatim        | Input name-string without any changes         |
| Cardinality     | 0 - N/A, 1 - Uninomial, 2 - Binomial etc.     |
| CanonicalStem   | Simplest canonical form with removed suffixes |
| CanonicalSimple | Simplest canonical form                       |
| CanonicalFull   | Canonical form with hybrid sign and ranks     |
| Authors         | Authorship of a name                          |
| Year            | Year of the name (if given)                   |
| Quality         | [Parsing quality][quality]                    |

### Quickly partition names by the type

Usually scientific names can be broken into groups according to the number of
elements:

- Uninomial
- Binomial
- Trinomial
- Quadrinomial

The output of `GNparser` contains a `Cardinality` field that tells, when
possible, how many elements are detected in the name.

| Cardinality | Name Type    |
| ----------- | ------------ |
| 0           | Undetermined |
| 1           | Uninomial    |
| 2           | Binomial     |
| 3           | Trinomial    |
| 4           | Quadrinomial |

For hybrid formulas, "approximate" names (with "sp.", "spp." etc.), unparsed
names, as well as names from `BOLD` project cardinality is 0 (Undetermined)

### Normalizing name-strings

There are many inconsistencies in how scientific names may be written.
Use `normalized` field to bring them all to a common form (spelling, spacing,
ranks).

### Removing authorship from the middle of the name

Often data administrators spit name-strings into "name part" and
"authorship part". This practice misses some information when dealing with
names like "_Prosthechea cochleata_ (L.) W.E.Higgins _var. grandiflora_
(Mutel) Christenson". However, if this is the use case, a combination of
`canonicalName -> full` with the authorship from the lowest taxon will do
the job. You can also use the default `--format csv` flag for `gnparser`
command line tool.

### Figuring out if names are well-formed

If there are problems with parsing a name, parser generates `qualityWarnings`
messages and lowers [parsing `quality`][quality] of the name. Quality values
mean the following:

- `"quality": 1` - No problems were detected.
- `"quality": 2` - There were small problems, normalized result
  should still be good.
- `"quality": 3` - There are some significant problems with parsing.
- `"quality": 4` - There were serious problems with the name, and the
  final result is rather doubtful.
- `"quality": 0` - A string could not be recognized as a scientific
  name and parsing failed.

### Creating stable GUIDs for name-strings

`GNparser` uses UUID version 5 to generate its `id` field.
There is algorithmic 1:1 relationship between the name-string and the UUID.
Moreover the same algorithm can be used in any popular language to
generate the same UUID. Such IDs can be used to globally connect information
about name-strings or information associated with name-strings.

More information about UUID version 5 can be found in the [Global Names
blog][uuid5]

### Assembling canonical forms etc. from original spelling

`GNparser` tries to correct problems with spelling, but sometimes it is
important to keep original spelling of the canonical forms or authorship.
The `words` field attaches semantic meaning to every word in the
original name-string and allows users to create canonical forms or other
combinations using the original verbatim spelling of the words. Each element
in `words` contains 3 parts:

1. verbatim value of a word
2. semantic meaning of the word
3. start position of the word
4. end position of the word

The `words` section belongs to additional details. To use it enable
`--details` flag for the command line application.

```bash
gnparser -d "Pardosa moesta Banks, 1892"
```

## Tutorials

- Parsing names from CSV files [tutorial][tutGN]

## Installation

Compiled programs in Go are self-sufficient and small (`GNparser` is only a
few megabytes). As a result the binary file of `gnparser` is all you need to
make it work. You can install it by downloading the [latest version of the
binary][releases] for your operating system **and** CPU architecture, and
placing it in your `PATH`.

### Install with Homebrew (Mac OS X, Linux)

[Homebrew] is a packaging system originally made for Mac OS X. You can use it
now for Mac, Linux, or Windows X WSL (Windows susbsystem for Linux).

1. Install Homebrew according to their [instructions][Homebrew].

2. Install `gnparser` with:

   ```bash
   brew update
   brew tap gnames/gn
   brew install gnparser
   ```

### Linux or Mac OS X

Move `gnparser` executable somewhere in your PATH
(for example `/usr/local/bin`)

```bash
tar xvf gnparser-xxx.tar.gz
sudo mv gnparser /usr/local/bin
```

Please note that Mac OS might prevent running executables that are not
'approved' up officially by Apple. To overcome this problem go to
`System Settings`, search for `developer tools`and under`Allow applications
below to run software locally that does not meet the system's security policy`
add the terminal emulation software that you normally use (e.g. Terminal app).

### Windows

One possible way would be to create a default folder for executables and place
`gnparser` there.

Use `Windows+R` keys
combination and type "`cmd`". In the appeared terminal window type:

```cmd
mkdir C:\bin
copy path_to\gnparser.exe C:\bin
```

[Add `C:\bin` directory to your `PATH`][winpath] `user` and/or `system`
environment variables.

It is also possible to install [Windows Subsystem for Linux][wsl] on Windows
(v10 or v11), and use `gnparser` as a Linux executable.

### Install with Go

If you have Go installed on your computer use

```bash
go get -u github.com/gnames/gnparser/gnparser
```

For development install gnu make and use the following:

```bash
git clone https://github.com/gnames/gnparser.git
cd gnparser
make tools
make install
```

You do need your `PATH` to include `$HOME/go/bin`

## Usage

### Command Line

```bash
gnparser -f pretty "Quadrella steyermarkii (Standl.) Iltis &amp; Cornejo"
```

Relevant flags:

`--help -h`
: Displays help information about the available flags.

`--batch_size -b`
: Sets the maximum number of names processed in a batch. This is ignored
in streaming mode (-s).

`--cultivar -C`
: Deprecated. Use `--nomenclatural-code` instead.

`--capitalize -c`
: Capitalizes the first letter of input name-strings.

`--details -d`
: Provides more detailed output for each parsed name. Ignored for
CSV/TSV formats.

`--diaereses -D`
: Preserves diaereses, e.g. `Leptochloöpsis virgata`. The stemmed
canonical name not include diaereses.

`--format -f`
: Specifies the output format: `csv`, `tsv`, `compact`, or `pretty`.
Defaults to `csv`. CSV and TSV formats include a header row.

`--jobs -j`
: Sets the number of jobs to run concurrently.

`--ignore_tags -i`
: Increases performance by skipping HTML entity and tag processing.
Only use if your input is known to be free of HTML.

`--nomenclatural-code -n`
: Specifies the nomenclatural code (e.g., `botanical`, `zoological`) to use
for parsing in ambiguous cases. For example in `Aus (Bus) cus`: according
to zoological code `Aus` is genus, `Bus` is subgenus, while according
to botanical code `Bus` is the author of `Aus`.

Supported values: `bact`, `bacterial`, `ICNP`, `bot`,
`botanical`, `ICN`, `cult`, `cultivar`, `ICNCP`, `zoo`, `zoological`, `ICZN`.

`--port -p`
: Sets the port for the web-interface and [RESTful API][OpenAPI].

`--species-group-cut`
: Modifies the stemmed canonical form for autonyms and species-group names
by removing the infraspecific epithet. Useful for matching names like
`Aus bus` and `Aus bus bus`.

`--stream -s`
: Enables streaming mode, where names are processed one at a time.
Useful for integrating gnparser with languages other than Go.

`--unordered -u`
: Disables output ordering. The output order may not match the input order.

`--version -V`
: Displays the version number of `GNparser`.

`--web-logs`
: Requires `--port`. Enables output of logs for web-services.

To parse one name:

```bash
# CSV output (default)
gnparser "Parus major Linnaeus, 1788"
# or
gnparser -f csv "Parus major Linnaeus, 1788"

# TSV output
gnparser -f tsv "Parus major Linnaeus, 1788"

# JSON compact format
gnparser "Parus major Linnaeus, 1788" -f compact

# pretty format
gnparser -f pretty "Parus major Linnaeus, 1788"

# to parse a name from the standard input
echo "Parus major Linnaeus, 1788" | gnparser

# to parse a botanical cultivar name
gnparser "Anthurium 'Ace of Spades'" --cultivar
gnparser "Phyllostachys vivax cv aureocaulis" -c

# to parse name that is all in low-case
gnparser "parus major" --capitalize
gnparser "parus major" -c
```

To parse a file:

There is no flag for parsing a file. If parser finds the given file path on
your computer, it will parse the content of the file, assuming that every line
is a new scientific name. If the file path is not found, `GNparser` will try
to parse the "path" as a scientific name.

Parsed results will stream to STDOUT, while progress of the parsing
will be directed to STDERR.

```bash
# to parse with 200 parallel processes
gnparser -j 200 names.txt > names_parsed.csv

# to parse file with more detailed output
gnparser names.txt -d -f compact > names_parsed.txt

# to parse files using pipes
cat names.txt | gnparser -f csv -j 200 > names_parsed.csv

# to parse using `stream` method instead of `batch` method.
cat names.txt | gnparser -s > names_parsed.csv

# to not remove html tags and entities during parsing. You gain a bit of
# performance with this option if your data does not contain HTML tags or
# entities.
gnparser "<i>Pomatomus</i>&nbsp;<i>saltator</i>"
gnparser -i "<i>Pomatomus</i>&nbsp;<i>saltator</i>"
gnparser -i "Pomatomus saltator"
```

If jobs number is set to more than 1, parsing uses several concurrent
processes. This approach increases speed of parsing on multi-CPU
computers. The results are returned in some random order, and reassembled
into the order of input transparently for a user.

Potentially the input file might contain millions of names, therefore creating
one properly formatted JSON output might be prohibitively expensive. Therefore
the parser creates one JSON line per name (when `compact` format is used)

You can use up to 20 times more "threads" than the number of your CPU cores
to reach maximum speed of parsing (`--jobs 200` flag). It is practical
because additional "threads" are very cheap in Go and they try to fill out
every idle gap in the CPU usage.

### Pipes

About any language has an ability to use pipes of the underlying operating
system. From the inside of your program you can make the CLI executable
`GNparser` to listen on a STDIN pipe and produce output into STDOUT pipe. Here
is an example in Ruby:

```ruby
def self.start_gnparser
  io = {}

  ['compact', 'csv'].each do |format|
    stdin, stdout, stderr = Open3.popen3("./gnparser -s --format #{format}")
    io[format.to_sym] = { stdin: stdin, stdout: stdout, stderr: stderr }
  end
end
```

@marcobrt kindly provided an [example in PHP][PHP pipes].

Note that you have to use `--stream -s` flag for this approach to work.

### R language package

For R language it is possible to use [`rgnparser` package][rgnparser]. It
implements mentioned above `pipes` method. It does require `gnparser` app
be installed.

### Ruby Gem

Ruby developers can use `GNparser` functionality via [biodiversity] gem. It
uses C-binding and does not require an installed `gnparser` app.

### Node.js

@tobymarsden created a [wrapper for node.js][node-gnparser]. It uses C-binding
and does not require an installed `gnparser` app.

### Usage as a REST API Interface or Web-based User Graphical Interface

Web-based user interface and API are invoked by `--port` or
`-p` flag. To start web server on `http://0.0.0.0:9000`

```bash
gnparser -p 9000
```

Opening a browser with this address will now show an interactive interface
to parser. API calls would be accessible on `http://0.0.0.0:9000/api/v1/`.

The api is and schema are described fully using [OpenAPI] specification.

Make sure to CGI-escape name-strings for GET requests. An '&' character
needs to be converted to '%26'

- `GET /api?q=Aus+bus|Aus+bus+D.+%26+M.,+1870`
- `POST /api` with request body of JSON array of strings

```ruby
require 'json'
require 'net/http'

uri = URI('https://parser.globalnames.org/api/v1/')
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
request = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json',
                                   'accept' => 'json')
request.body = ['Solanum mariae Särkinen & S.Knapp',
                'Ahmadiago Vánky 2004'].to_json
response = http.request(request)
```

#### Enabling logs for GNparser's web-service

There are several ways to enable logging from a web service.

The following enables web-access logs to be printed to STDERR

```
gnparser -p 80 --web-logs
```

### Use as a Docker image

You need to have [docker runtime installed](https://docs.docker.com/install/)
on your computer for these examples to work.

```bash
# run as a website and a RESTful service
docker run -p 0.0.0.0:80:8080 gnames/gognparser -p 8080 --web-logs

# just parse something
docker run gnames/gognparser "Amaurorhinus bewichianus (Wollaston,1860) (s.str.)"
```

### Use as a library in Go

```go
import (
  "fmt"

  "github.com/gnames/gnparser"
  "github.com/gnames/gnparser/ent/parsed"
)

func Example() {
  names := []string{"Pardosa moesta Banks, 1892", "Bubo bubo"}
  cfg := gnparser.NewConfig()
  gnp := gnparser.New(cfg)
  res := gnp.ParseNames(names)
  fmt.Println(res[0].Authorship.Normalized)
  fmt.Println(res[1].Canonical.Simple)
  fmt.Println(parsed.HeaderCSV(gnp.Format()))
  fmt.Println(res[0].Output(gnp.Format()))
  // Output:
  // Banks 1892
  // Bubo bubo
  // Id,Verbatim,Cardinality,CanonicalStem,CanonicalSimple,CanonicalFull,Authorship,Year,Quality
  // e2fdf10b-6a36-5cc7-b6ca-be4d3b34b21f,"Pardosa moesta Banks, 1892",2,Pardosa moest,Pardosa moesta,Pardosa moesta,Banks 1892,1892,1
}
```

### Use as a shared C library

It is possible to bind `GNparser` functionality with languages that can use
C Application Binary Interface. For example such languages include
Python, Ruby, Rust, C, C++, Java (via JNI).

To compile `GNparser` shared library for your platform/operating system of
choice you need `GNU make` and `GNU gcc compiler` installed:

```bash
make clib
cd binding
cp libgnparser* /path/to/some/project
```

As an example how to use the shared library check this [StackOverflow
question][ruby_ffi_go_usage] and [biodiversity] Ruby gem.

## Parsing ambiguities

Some name-strings cannot be parsed unambiguously without some additional data.

### Names with `filius` (ICN code)

For names like `Aus bus Linn. f. cus` the `f.` is ambiguous. It might mean
that species were described by a son of (`filius`) Linn., or it might mean
that `cus` is `forma` of `bus`. We provide a warning
"Ambiguous f. (filius or forma)" for such cases.

### Names with subgenus (ICZN code) and genus author (ICN code)

For names like `Aus (Bus) L.` or `Aus (Bus) cus L.` the `(Bus)` token would
mean the name of subgenus for ICZN names, but for ICN names it would be an
author of genus `Aus`. We created a list of ICN generic authors using data from
[IRMNG] to distinguish such names from each other. For detected ICN names we
provide a warning "Ambiguity: ICN author or subgenus".

## Authors

- [Dmitry Mozzherin]

## Contributors

- [Toby Marsden]
- [Geoff Ower]
- [Hernan Lucas Pereira]

If you want to submit a bug or add a feature read
[CONTRIBUTING] file.

## References

Mozzherin, D.Y., Myltsev, A.A. & Patterson, D.J. “gnparser”: a powerful parser
for scientific names based on Parsing Expression Grammar. BMC Bioinformatics
18, 279 (2017).[https://doi.org/10.1186/s12859-017-1663-3][gnparser paper]

Rees, T. (compiler) (2019). The Interim Register of Marine and Nonmarine
Genera. Available from `http://www.irmng.org` at VLIZ.
Accessed 2019-04-10

## License

Released under [MIT license]

[CONTRIBUTING]: https://github.com/gnames/gnparser/blob/master/CONTRIBUTING.md
[Dmitry Mozzherin]: https://github.com/dimus
[Geoff Ower]: https://github.com/gdower
[Toby Marsden]: https://github.com/tobymarsden
[Hernan Lucas Pereira]: https://github.com/LocoDelAssembly
[Homebrew]: https://brew.sh/
[IRMNG]: http://www.irmng.org
[MIT license]: https://github.com/gnames/gnparser/raw/master/LICENSE
[Schinke R et al (1996)]: https://caio.ueberalles.net/a_stemming_algorithm_for_latin_text_databases-schinke_et_al.pdf
[Zenodo DOI]: https://zenodo.org/badge/latestdoi/320967495
[biodiversity]: https://github.com/GlobalNamesArchitecture/biodiversity
[export file]: https://github.com/gnames/gnparser/blob/master/binding/main.go
[gna]: http://globalnames.org
[node-gnparser]: https://github.com/amazingplants/node-gnparser
[OpenAPI]: https://apidoc.globalnames.org/gnparser
[gnparser ruby]: https://gitlab.com/gnames/gnparser_rb
[gnparser-scala]: https://github.com/GlobalNamesArchitecture/gnparser
[gnparser.proto]: https://github.com/gnames/gnparser/blob/master/pb/gnparser.proto
[parser-web]: https://parser.globalnames.org
[peg]: https://github.com/pointlander/peg
[quality]: https://github.com/gnames/gnparser/blob/master/quality.md
[releases]: https://github.com/gnames/gnparser/releases/latest
[ruby_ffi_go_usage]: https://stackoverflow.com/questions/58866962/how-to-pass-an-array-of-strings-and-get-an-array-of-strings-in-ruby-using-go-sha
[tutGN]: https://globalnames.org/docs/tut-xsv-gnparser/
[test file]: https://github.com/gnames/gnparser/blob/master/testdata/test_data.md
[uuid5]: http://globalnames.org/news/2015/05/31/gn-uuid-0-5-0
[winpath]: https://www.computerhope.com/issues/ch000549.htm
[wsl]: https://docs.microsoft.com/en-us/windows/wsl/
[gnparser paper]: https://doi.org/10.1186/s12859-017-1663-3
[PHP pipes]: https://gist.github.com/marcobrt/72b2a3d1b0649c1bf738c9fc88f74ec0
[rgnparser]: https://github.com/ropensci/rgnparser
[v0diff]: https://github.com/gnames/gnparser/wiki/Changes-in-v1.0.0
[OpenRefine]: https://github.com/gnames/gnparser/wiki/GNparser-with-OpenRefine
