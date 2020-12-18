# Global Names Parser: gnparser written in Go

Try `gnparser` [online][parser-web].

``gnparser`` splits scientific names into their semantic elements with an
associated meta information. For example, ``"Homo sapiens Linnaeus"`` is
parsed into:

| Element  | Meaning         | Position |
| -------- | --------------- | -------- |
| Homo     | genus           | (0,4)    |
| sapiens  | specificEpithet | (5,12)   |
| Linnaeus | author          | (13,21)  |

This parser, written in Go, is the 3rd iteration of the project. The first,
[biodiversity] had been written in Ruby, the second, [also
gnparser][gnparser-scala], had been written in Scala. This project is now
a substitution for the other two. It will be the only one that is maintained
further. All three projects were developed as a part of
[Global Names Architecture Project][gna].

To use `gnparser` as a command line tool under Windows, Mac or Linux,
download the [latest release][releases], uncompress it, and copy `gnparser`
binary somewhere in your PATH.

```bash
wget https://github.com/gnames/gnparser/uploads/55d247b8fbade60116c7e3b650dd978c/gnparser-v0.9.0-linux.tar.gz
tar xvf gnparser-v0.9.0-linux.tar.gz
sudo cp gnparser /usr/local/bin
# for CSV output
gnparser "Homo sapiens Linnaeus"
# for JSON output
gnparser -f compact "Homo sapiens Linnaeus"
# or
gnparser -f pretty "Homo sapiens Linnaeus"
gnparser -h
```

<!-- vim-markdown-toc GFM -->

* [Introduction](#introduction)
* [Speed](#speed)
* [Features](#features)
* [Use Cases](#use-cases)
  * [Getting the simplest possible canonical form](#getting-the-simplest-possible-canonical-form)
  * [Quickly partition names by the type](#quickly-partition-names-by-the-type)
  * [Normalizing name-strings](#normalizing-name-strings)
  * [Removing authorships from the middle of the name](#removing-authorships-from-the-middle-of-the-name)
  * [Figuring out if names are well-formed](#figuring-out-if-names-are-well-formed)
  * [Creating stable GUIDs for name-strings](#creating-stable-guids-for-name-strings)
  * [Assembling canonical forms etc. from original spelling](#assembling-canonical-forms-etc-from-original-spelling)
* [Installation](#installation)
  * [Linux or OS X](#linux-or-os-x)
  * [Windows](#windows)
  * [Install with Go](#install-with-go)
* [Usage](#usage)
  * [Command Line](#command-line)
  * [Pipes](#pipes)
  * [gRPC server](#grpc-server)
  * [Usage as a REST API Interface](#usage-as-a-rest-api-interface)
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

## Introduction

Global Names Parser or ``gnparser`` is a program written in Go for breaking up
scientific names into their elements.  It uses [peg] -- a Parsing
Expression Grammar (PEG) tool.

Many other parsing algorithms for scientific names use regular expressions.
This approach works well for extracting canonical forms in simple cases.
However, for complex scientific names and to parse scientific names into
all semantic elements, regular expressions often fail, unable to overcome
the recursive nature of data embedded in names. By contrast, ``gnparser``
is able to deal with the most complex scientific name-strings.

``gnparser`` takes a name-string like ``Drosophila (Sophophora) melanogaster
Meigen, 1830`` and returns parsed components in `CSV` or `JSON` format. The
parsing of scientific names might become surprisingly complex and the
`gnparser's` [test file] is a good source of information about the parser's
capabilities, its input and output.

## Speed

Number of names parsed per hour on a i7-8750H CPU
(6 cores, 12 threads, at 2.20 GHz), parser v0.5.1

| Threads | names/hr    |
| ------- | ----------- |
| 1       | 48,000,000  |
| 2       | 63,000,000  |
| 4       | 128,000,000 |
| 8       | 202,000,000 |
| 16      | 248,000,000 |
| 100     | 293,000,000 |

For simplest output Go ``gnparser`` is roughly 2 times faster than Scala
``gnparser`` and about 100 times faster than Ruby ``biodiversity`` parser. For
JSON formats the parser is approximately 8 times faster than Scala one, due to
more efficient JSON conversion.

## Features

* Fastest parser ever.
* Very easy to install, just placing executable somewhere in the PATH is
  sufficient.
* Extracts all elements from a name, not only canonical forms.
* Works with very complex scientific names, including hybrid formulas.
* Includes gRPC server that can be used as if a native method call from C++,
* C#, Java, Python, Ruby, PHP, JavaScript, Objective C, Dart.
* Use as a native library from Go projects.
* Can run as a command line application.
* Can be scaled to many CPUs and computers (if 300 millions names an
  hour is not enough).
* Calculates a stable UUID version 5 ID from the content of a string.
* Provides C-binding to incorporate parser into other languages.

## Use Cases

### Getting the simplest possible canonical form

Canonical forms of a scientific name are the latinized components without
annotations, authors or dates. They are great for matching names that differ
in less stable parts. Use the ``canonicalName -> simple`` or ``canonicalName
-> full`` fields from parsing results for this use case. ``Full`` version of
canonical form includes infra-specific ranks and hybrid character for named
hybrids.

The ``canonicalName -> full`` is good for presentation, as it keeps more
details.

The ``canonicalName -> simple`` field is good for matching names from different
sources, because sometimes dataset curators omit hybrid sign in named hybrids,
or remove ranks for infraspecific epithets.

The ``canonicalName -> stem`` field normalizes `simple` canonical form even
further. The normalization is done according to stemming rules for Latin
language described in [Schinke R et al (1996)]. For example letters `j` are
converted to `i`, letters `v` are converted to `u`, and suffixes are removed
from the specific and infraspecific epithets.

If you only care about canonical form of a name you can use ``--format csv``
flag with command line tool.

CSV output has the following fields:

| Field             | Meaning                                         |
| ------------------| ----------------------------------------------- |
| Id                | UUID v5 generated out of Verbatim               |
| Verbatim          | Input name-string without any changes           |
| Cardinality       | 0 - N/A, 1 - Uninomial, 2 - Binomial etc.       |
| CanonicalFull     | Canonical form with hybrid sign and ranks       |
| CanonicalSimple   | Simplest canonical form                         |
| CanonicalStem     | Simplest canonical form with removed suffixes   |
| Authors           | Author string of a name                         |
| Year              | Year of the name (if given)                     |
| Quality           | Parsing quality                                 |

### Quickly partition names by the type

Usually scientific names can be broken into groups accoring by number of
elements:

* Uninomial
* Binomial
* Trinomial
* Quadrinomial

The output of `gnparser` contains a `Cardinality` field that tells, when
possible, how many elements are detected in the name.

| Cardinality  | Name Type    |
| ------------ | ------------ |
| 0            | Undetermined |
| 1            | Uninomial    |
| 2            | Binomial     |
| 3            | Trinomial    |
| 4            | Quadrinomial |

For hybrid formulas, "approximate" names (with "sp.", "spp." etc.), unparsed
names, as well as names from `BOLD` project cardinality is 0 (Undetermined)

### Normalizing name-strings

There are many inconsistencies in how scientific names may be written.
Use ``normalized`` field to bring them all to a common form (spelling, spacing,
ranks).

### Removing authorships from the middle of the name

Many data administrators store name-strings in two columns and split them into
"name part" and "authorship part". This practice misses some information when
dealing with names like "*Prosthechea cochleata* (L.) W.E.Higgins *var.
grandiflora* (Mutel) Christenson". However, if this is the use case, a
combination of ``canonicalName -> full`` with the authorship from the
lowest taxon will do the job. You can also use ``--format csv`` flag for
``gnparse`` command line tool.

### Figuring out if names are well-formed

If there are problems with parsing a name, parser generates ``qualityWarnings``
messages and lowers parsing ``quality`` of the name.  Quality values mean the
following:

* ``"quality": 1`` - No problems were detected
* ``"quality": 2`` - There were small problems, normalized result
  should still be good
* ``"quality": 3`` - There were serious problems with the name, and the
  final result is rather doubtful
* ``"quality": 0`` - A string could not be recognized as a scientific
  name and parsing fails

### Creating stable GUIDs for name-strings

``gnparser`` uses UUID version 5 to generate its ``id`` field.
There is algorithmic 1:1 relationship between the name-string and the UUID.
Moreover the same algorithm can be used in any popular language to
generate the same UUID. Such IDs can be used to globally connect information
about name-strings or information associated with name-strings.

More information about UUID version 5 can be found in the [Global Names
blog][uuid5]

### Assembling canonical forms etc. from original spelling

``gnparser`` tries to correct problems with spelling, but sometimes it is
important to keep original spelling of the canonical forms or authorships.
The ``positions`` field attaches semantic meaning to every word in the
original name-string and allows users to create canonical forms or other
combinations using the original verbatim spelling of the words. Each element
in ``positions`` contains 3 parts:

1. semantic meaning of a word
2. start position of the word
3. end position of the word

For example ``["specificEpithet", 6, 11]`` means that a specific epithet starts
at 6th character and ends *before* 11th character of the string.

## Installation

Compiled programs in Go are self-sufficient and small (``gnparser`` is only a
few megabytes). As a result the binary file of ``gnparser`` is all you need to
make it work. You can install it by downloading the [latest version of the
binary][releases] for your operating system, and placing it in your ``PATH``.

### Linux or OS X

Move ``gnparser`` executable somewhere in your PATH
(for example ``/usr/local/bin``)

```bash
sudo mv path_to/gnparser /usr/local/bin
```

### Windows

One possible way would be to create a default folder for executables and place
``gnparser`` there.

Use ``Windows+R`` keys
combination and type "``cmd``". In the appeared terminal window type:

```cmd
mkdir C:\bin
copy path_to\gnparser.exe C:\bin
```

[Add ``C:\bin`` directory to your ``PATH``][winpath] environment variable.

### Install with Go

If you have Go installed on your computer use

```bash
go get -u github.com/gnames/gnparser
cd $GOPATH/srs/github.com/gnames/gnparser
make install
```

You do need your ``PATH`` to include ``$HOME/go/bin``

## Usage

### Command Line

```bash
gnparser -f pretty "Quadrella steyermarkii (Standl.) Iltis &amp; Cornejo"
```

Relevant flags:

``--help -h``
: help information about flags

``--format -f``
: output format. Can be ``compact``, ``pretty``, ``csv``, or ``debug``.
Default is ``csv``.

CSV format returns a header row and the CSV-compatible parsed result.

``--jobs -j``
: number of jobs running concurrently.

``--nocleanup -n``
: keeps HTML entities and tags if they are present in a name-string. If your
data is clean from HTML tags or entities, you can use this flag to increase
performance.

To parse one name:

```bash
# CSV ouput (default)
gnparser "Parus major Linnaeus, 1788"
# or
gnparser -f csv "Parus major Linnaeus, 1788"

# JSON compact format
gnparser "Parus major Linnaeus, 1788" -f compact

# pretty format
gnparser -f pretty "Parus major Linnaeus, 1788"

# to parse a name from the standard input
echo "Parus major Linnaeus, 1788" | gnparser
```

To parse a file:

There is no flag for parsing a file. If parser finds the given file path on
your computer, it will parse the content of the file, assuming that every line
is a new scientific name. If the file path is not found, ``gnparser`` will try
to parse the "path" as a scientific name.

Parsed results will stream to STDOUT, while progress of the parsing
will be directed to STDERR.

```bash
gnparser -j 200 names.txt > names_parsed.txt

# to parse files using pipes
cat names.txt | gnparser -f csv -j 200 > names_parsed.txt

# to keep html tags and entities during parsing. You gain a bit of performance
# with this option if your data does not contain HTML tags or entities.
gnparser "<i>Pomatomus</i>&nbsp;<i>saltator</i>"
gnparser -n "<i>Pomatomus</i>&nbsp;<i>saltator</i>"
gnparser -n "Pomatomus saltator"
```

To parse a file returning results in the same order as they are given (slower):

```bash
gnparser -j 1 names.txt > names_parsed.txt
```

Potentially the input file might contain millions of names, therefore creating
one properly formatted JSON output might be prohibitively expensive. Therefore
the parser creates one JSON line per name (when ``compact`` format is used)

You can use up to 20 times more "threads" than the number of your CPU cores to
reach maximum speed of parsing (``--jobs 200`` flag). It is practical because
additional threads are very cheap in Go and they try to fill out every idle
gap in the CPU usage.

### Pipes

About any language has an ability to use pipes of the underlying operating
system. From the inside of your program you can make the CLI executable `gnparser`
to listen on a STDIN pipe and produce output into STDOUT pipe. Here is an
example in Ruby:

```ruby
def self.start_gnparser
  io = {}

  ['compact', 'csv'].each do |format|
    stdin, stdout, stderr = Open3.popen3("./gnparser -j 200 --format #{format}")
    io[format.to_sym] = { stdin: stdin, stdout: stdout, stderr: stderr }
  end
end
```

Such arrangement would give you a nearly native performance for large datasets.

### gRPC server

Relevant flags:

``--help -h``
: help information about flags

``--grpc -g``
: sets a port to run gRPC server, and starts gnparser in gRPC mode.

``--jobs -j``
: number or workers allocated per gRPC request. Default corresponds to the
  number of CPU threads. If you have a full control over gRPC server of
  `gnparser`, set this option to 100-300 jobs.

```bash
gnparser -g 8989 -j 200
```

For an example how to use gRPC server check ``gnparser`` [Ruby gem][gnparser
ruby] as well as [gRPC documentation].

It also helps to read [gnparser.proto] file to understand how to deal with
inputs and outputs of gRPC server.

### Usage as a REST API Interface

Use web-server REST API as a slower, but a more wide-spread alternative to
gRPC server. Web-based user interface and API are invoked by ``--web-port`` or
``-w`` flag. To start web server on ``http://0.0.0.0:9000``

```bash
gnparser -w 9000
```

Opening a browser with this address will now show an interactive interface
to parser. API calls would be accessibe on ``http://0.0.0.0:9000/api``.

Make sure to CGI-escape name-strings for GET requests. An '&' character
needs to be converted to '%26'

* ``GET /api?q=Aus+bus|Aus+bus+D.+%26+M.,+1870``
* ``POST /api`` with request body of JSON array of strings

```ruby
require 'json'
require 'net/http'

uri = URI('https://parser.globalnames.org/api')
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
request = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json',
                                   'accept' => 'json')
request.body = ['Solanum mariae Särkinen & S.Knapp',
                'Ahmadiago Vánky 2004'].to_json
response = http.request(request)
```

### Use as a Docker image

You need to have [docker runtime installed](https://docs.docker.com/install/)
on your computer for these examples to work.

```bash
# run as a gRPC server on port 7777
docker run -p 0.0.0.0:7777:7777 gnames/gognparser -g 7777
# run grpc on 'default' 8778 port
docker run -p 0.0.0.0:8778:8778 gnames/gognparser
# to run as a daemon with 50 workers
docker run -d gnames/gognparser -g 7777 -j 50

# run as a website and a RESTful service
docker run -p 0.0.0.0:80:8080 gnames/gognparser -w 8080

# just parse something
docker run gnames/gognparser "Amaurorhinus bewichianus (Wollaston,1860) (s.str.)"
```

### Use as a library in Go

```go
package main

import (
  "fmt"

  "github.com/gnames/gnparser"
)

func main() {
  opts := []gnparser.Option{
    gnparser.Format("csv"),
    gnparser.WorkersNum(100),
  }
  gnp := gnparser.NewGNparser(opts...)
  res, err := gnp.ParseAndFormat("Bubo bubo")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(res)
}
```

To avoid JSON format we provide `gnp.ParseToObject` function.
Use [gnparser.proto] file as a reference of the available object fields.

```go
gnp := NewGNparser()
o := gnp.ParseToObject("Homo sapiens")

fmt.Println(o.Canonical.Simple)
switch d := o.Details.(type) {
case *pb.Parsed_Species:
  fmt.Println(d.Species.Genus)
case *pb.Parsed_Uninomial:
  fmt.Println(d.Uninomial.Value)
...
}
```

### Use as a shared C library

It is possible to bind `gnparser` functionality with languages that can use
C Application Binary Interface. For example such languages include
Python, Ruby, Rust, C, C++, Java (via JNI).

To compile `gnparser` shared library for your platform/operating system of
choice you need `GNU make` and `GNU gcc compiler` installed:

```bash
make clib
cd binding
cp libgnparser* /path/to/some/project
```

As an example how to use the shared library check this [StackOverflow
question][ruby_ffi_go_usage] and [biodiversity] Ruby gem. You can
find shared functions at their [export file].

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
provide a warning "Possible ICN author instead of subgenus".

## Authors

* [Dmitry Mozzherin]

## Contributors

* [Geoff Ower]
* [Hernan Lucas Pereira]

If you want to submit a bug or add a feature read
[CONTRIBUTING] file.

## References

Rees, T. (compiler) (2019). The Interim Register of Marine and Nonmarine
Genera. Available from `http://www.irmng.org` at VLIZ.
Accessed 2019-04-10

## License

Released under [MIT license]

[releases]: https://github.com/gnames/gnparser/-/releases
[biodiversity]: https://github.com/GlobalNamesArchitecture/biodiversity
[gnparser-scala]: https://github.com/GlobalNamesArchitecture/gnparser
[peg]: https://github.com/pointlander/peg
[gna]: http://globalnames.org
[test file]: https://github.com/gnames/gnparser/raw/master/testdata/test_data.txt
[uuid5]: http://globalnames.org/news/2015/05/31/gn-uuid-0-5-0
[winpath]: https://www.computerhope.com/issues/ch000549.htm
[gnparser ruby]: https://gitlab.com/gnames/gnparser_rb
[gRPC documentation]: https://grpc.io/docs/quickstart
[Dmitry Mozzherin]: https://github.com/dimus
[Geoff Ower]: https://github.com/gdower
[Hernan Lucas Pereira]: https://github.com/LocoDelAssembly
[MIT license]: https://github.com/gnames/gnparser/raw/master/LICENSE
[parser-web]: https://parser.globalnames.org
[IRMNG]: http://www.irmng.org
[CONTRIBUTING]: https://github.com/gnames/gnparser/blob/master/CONTRIBUTING.md
[gnparser.proto]: https://github.com/gnames/gnparser/blob/master/pb/gnparser.proto
[Schinke R et al (1996)]: https://caio.ueberalles.net/a_stemming_algorithm_for_latin_text_databases-schinke_et_al.pdf
[ruby_ffi_go_usage]: https://stackoverflow.com/questions/58866962/how-to-pass-an-array-of-strings-and-get-an-array-of-strings-in-ruby-using-go-sha
[export file]: https://github.com/gnames/gnparser/blob/master/binding/main.go
