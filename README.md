# Global Names Parser: gnparser written in Go

Try in [online][parser-web].

``gnparser`` splits scientific names into their component elements with
associated meta information.  For example, ``"Homo sapiens Linnaeus"`` is
parsed into human readable information as follows:

| Element  | Meaning          | Position
| -------- | ---------------- | --------
| Homo     | genus            | (0,4)
| sapiens  | specificEpithet  | (5,12)
| Linnaeus | author           | (13,21)

This parser, written in Go, is the 3rd iteration of the project. The first,
[biodiversity] had been written in Ruby, the second, [also
gnparser][gnparser-scala], had been written in Scala. This project learned
from previous ones, and, when it matures, it is going to be the a
substitution for other two, and will be the only one that is maintained
further. All three projects were developed as a part of [Global Names
Architecture Project][gna].

Try as a command tool under Windows, Mac or Linux by downloading the [latest
release][releases], uncompressing it, and copying `gnparser` binary somewhere
in your PATH.

```bash
wget https://www.dropbox.com/s/blvmejmp4378cao/gnparser-v0.5.1-linux.tar.gz
tar xvf gnparser-v0.5.1-linux.tar.gz
sudo cp gnparser /usr/local/bin
# for JSON output
gnparser -f pretty "Homo sapiens Linnaeus"
# for very simple text output
gnparser -f simple "Homo sapiens Linnaeus"
gnparser -h
```

- [Global Names Parser: gnparser written in Go](#global-names-parser-gnparser-written-in-go)
	- [Introduction](#introduction)
	- [Speed](#speed)
	- [Features](#features)
	- [Use Cases](#use-cases)
		- [Getting the simplest possible canonical form](#getting-the-simplest-possible-canonical-form)
		- [Normalizing name-strings](#normalizing-name-strings)
		- [Removing authorships in the middle of the name](#removing-authorships-in-the-middle-of-the-name)
		- [Figuring out if names are well-formed](#figuring-out-if-names-are-well-formed)
		- [Creating stable GUIDs for name-strings](#creating-stable-guids-for-name-strings)
		- [Assembling canonical forms etc. from original spelling](#assembling-canonical-forms-etc-from-original-spelling)
	- [Installation](#installation)
		- [Linux or OS X](#linux-or-os-x)
		- [Windows](#windows)
		- [Install with Go](#install-with-go)
	- [Usage](#usage)
		- [Command Line](#command-line)
		- [gRPC server](#grpc-server)
		- [Usage as a REST API Interface](#usage-as-a-rest-api-interface)
		- [Use as a Docker image](#use-as-a-docker-image)
		- [Use as a library in Go](#use-as-a-library-in-go)
	- [Authors](#authors)
	- [Contributors](#contributors)
	- [License](#license)

## Introduction

Global Names Parser or ``gnparser`` is a program written in Go for breaking up
scientific names into their different elements.  It uses [peg] -- a Parsing
Expression Grammar (PEG) tool.

Many other parsing algorithms for scientific names use regular expressions.
This approach works well for extracting canonical forms in simple cases.
However, for complex scientific names and to parse scientific names into
all semantic elements regular expressions often fail, unable to overcome
the recursive nature of data embedded in names. By contrast, ``gnparser``
is able to deal with the most complex scientific name-strings.

``gnparser`` takes a name-string like ``Drosophila (Sophophora) melanogaster
Meigen, 1830`` and returns parsed components in `JSON` format. This behavior is
defined in its tests and the [test file] is a good source of information about
parser's capabilities, its input and output.

## Speed

Number of names parsed per hour on a i7-8750H CPU
(6 cores, 12 threads, at 2.20 GHz), parser v0.5.1

| Threads  | names/hr
| -------- | ------------
| 1        |  48,000,000
| 2        |  63,000,000
| 4        | 128,000,000
| 8        | 202,000,000
| 16       | 248,000,000
| 100      | 293,000,000

For simplest output Go ``gnparser`` is roughly 2 times faster than Scala
``gnparser`` and about 100 times faster than Ruby ``biodiversity`` parser. For
JSON formats the parser is approximately 8 times faster than Scala one, due to
more efficient JSON conversion.

## Features

- Fastest parser ever.
- Very easy to install, just placing executable somewhere in the PATH is
  sufficient.
- Extracts all elements from a name, not only canonical forms.
- Works with very complex scientific names, including hybrids.
- Includes gRPC server that can be used as if a native method call from C++,
  C#, Java, Python, Ruby, PHP, JavaScript, Objective C, Dart.
- Use as a native library from Go projects.
- Can run as a command line application.
- Can be scaled to many CPUs and computers (if 300 millions names an
   hour is not enough).
- Calculates a stable UUID version 5 ID from the content of a string.

## Use Cases

### Getting the simplest possible canonical form

Canonical forms of a scientific name are the latinized components without
annotations, authors or dates. They are great for matching names despite
alternative spellings. Use the ``canonicalName -> simple`` or ``canonicalName
-> full`` fields from parsing results for this use case. ``Full`` version of
canonical form includes infra-specific ranks and hybrid character for named
hybrids.

The ``canonicalName -> simple`` field is good for matching names from different
sources, because sometimes dataset curators omit hybrid sign in named hybrids,
or remove ranks for infraspecific epithets.

The ``canonicalName -> full`` is good for presentation, as it keeps more
details.

If you only care about canonical form of a name you can use ``--format simple``
flag with command line tool or gRPC service.

### Normalizing name-strings

There are many inconsistencies in how scientific names may be written.
Use ``normalized`` field to bring them all to a common form (spelling, spacing,
ranks).

### Removing authorships in the middle of the name

Many data administrators store name-strings in two columns and split them into
"name part" and "authorship part". This practice misses some information when
dealing with names like "*Prosthechea cochleata* (L.) W.E.Higgins *var.
grandiflora* (Mutel) Christenson". However, if this is the use case, a
combination of ``canonicalName -> valueRanked`` with the authorship from the
lowest taxon will do the job. You can also use ``--format simple`` flag for
``gnparse`` command line tool.

### Figuring out if names are well-formed

If there are problems with parsing a name, parser generates ``qualityWarnings``
messages and lowers parsing ``quality`` of the name.  Quality values mean the
following:

- ``"quality": 1`` - No problems were detected
- ``"quality": 2`` - There were small problems, normalized result
  should still be good
- ``"quality": 3`` - There were serious problems with the name, and the
  final result is rather doubtful
- ``"quality": 0`` - A string could not be recognized as a scientific
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
go get -u gitlab.com/gogna/gnparser
cd $GOPATH/srs/gitlab.com/gogna/gnparser
make install
```

You do need your ``PATH`` to include ``$HOME/go/bin``

## Usage

### Command Line

Relevant flags:

``--help -h``
: help information about flags

``--format -f``
: output format. Can be ``compact``, ``pretty``, ``simple``, or ``debug``.
Default is ``compact``.

``--jobs -j``
: number of jobs running concurrently.

``--cleanup -c``
: cleans up input from HTML entities and tags instead of parsing

To parse one name:

```bash
# default compact format
gnparser "Parus major Linnaeus, 1788"

# pretty format
gnparser -f pretty "Parus major Linnaeus, 1788"

# simple pipe-delimited flat format
gnparser -f simple "Parus major Linnaeus, 1788"

# to parse a name from standard input
echo "Parus major Linnaeus, 1788" | gnparser
```

To parse a file:

There is no flag for parsing a file. If parser finds file path on your computer
it will parse the content of the file, assuming every line is a new scientific
name.  If the file path is not found, ``gnparser`` will try to parse the "path"
as a scientific name.

Parsed results will stream to STDOUT, while progress of the parsing
will be directed to STDERR.

```bash
gnparser -j 200 names.txt > names_parsed.txt

# to parse files using pipes
cat names.txt | gnparser -f simple -j 200 > names_parsed.txt

# to clean names from html tags and entities first (no parsing
# or other changes), then parse
cat names.txt | gnparser -c | sed "s/.*|//" | gnparser > names_parsed.txt
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

To cleanup a name (no parsing here, it just removes HTML tags and entities,
and makes no other modifications):

The output contains the original name-string, and "HTML-normalized" one
separated by a pipe ("|") character.

```bash
gnparser -c "<i>Abacopteris glandulosa</i> (Bl.) F&eacute;e &amp; Chin"
```

To cleanup a file of names

```bash
gnparser -j 200 -c names.txt > no_html_names.txt

# using pipes
cat names.txt | gnparser -c -j 200 > no_html_names.txt
```

If you have data that has names with tags or HTML entities, the ``--cleanup
-c`` flag will help to normalize such names for parsing or other purposes.

### gRPC server

Relevant flags:

``--help -h``
: help information about flags

``--grpc -g``
: sets a port to run gRPC server, and starts gnparser in gRPC mode.

``--jobs -j``
: number or workers allocated per gRPC request. Default corresponds to the
  number of CPU threads.

```bash
gnparser -g 8989 -j 20
```

For an example how to use gRPC server check ``gnparser`` [Ruby gem][gnparser
ruby] as well as [gRPC documentation].

### Usage as a REST API Interface

Use web-server REST API as a slower, but more wide-spread alternative to
gRPC server. Web-based user interface and API are invoked by ``--web-port`` or
``-w`` flag. To start web server on ``http://0.0.0.0:9000``

```bash
    gnparser -w 9000
```

Opening a browser with this address will now show an interactive interface
to parser. API calls would be accessibe on ``http://0.0.0.0:9000/api``.

Make sure to CGI-escape name-strings for GET requests. An '&' character
needs to be converted to '%26'

- ``GET /api?q=Aus+bus|Aus+bus+D.+%26+M.,+1870``
- ``POST /api`` with request body of JSON array of strings

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

	"gitlab.com/gogna/gnparser"
)

func main() {
	opts := []gnparser.Option{
		gnparser.Format("simple"),
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
## Authors

- [Dmitry Mozzherin]

## Contributors

- [Geoff Ower]

If you want to submit a bug or add a feature read
[CONTRIBUTING](https://gitlab.com/gogna/gnparser/blob/master/CONTRIBUTING.md) file.

## License
- [Global Names Parser: gnparser written in Go](#global-names-parser-gnparser-written-in-go)
	- [Introduction](#introduction)
	- [Speed](#speed)
	- [Features](#features)
	- [Use Cases](#use-cases)
		- [Getting the simplest possible canonical form](#getting-the-simplest-possible-canonical-form)
		- [Normalizing name-strings](#normalizing-name-strings)
		- [Removing authorships in the middle of the name](#removing-authorships-in-the-middle-of-the-name)
		- [Figuring out if names are well-formed](#figuring-out-if-names-are-well-formed)
		- [Creating stable GUIDs for name-strings](#creating-stable-guids-for-name-strings)
		- [Assembling canonical forms etc. from original spelling](#assembling-canonical-forms-etc-from-original-spelling)
	- [Installation](#installation)
		- [Linux or OS X](#linux-or-os-x)
		- [Windows](#windows)
		- [Install with Go](#install-with-go)
	- [Usage](#usage)
		- [Command Line](#command-line)
		- [gRPC server](#grpc-server)
		- [Usage as a REST API Interface](#usage-as-a-rest-api-interface)
		- [Use as a Docker image](#use-as-a-docker-image)
		- [Use as a library in Go](#use-as-a-library-in-go)
	- [Authors](#authors)
	- [Contributors](#contributors)
	- [License](#license)
Released under [MIT license]

[releases]: https://gitlab.com/gogna/gnparser/releases
[biodiversity]: https://github.com/GlobalNamesArchitecture/biodiversity
[gnparser-scala]: https://github.com/GlobalNamesArchitecture/gnparser
[peg]: https://github.com/pointlander/peg
[gna]: http://globalnames.org
[test file]: https://gitlab.com/gogna/gnparser/raw/master/test-data/test_data.txt
[uuid5]: http://globalnames.org/news/2015/05/31/gn-uuid-0-5-0
[winpath]: https://www.computerhope.com/issues/ch000549.htm
[gnparser ruby]: https://gitlab.com/gnames/gnparser_rb
[gRPC documentation]: https://grpc.io/docs/quickstart
[Dmitry Mozzherin]: https://gitlab.com/dimus
[Geoff Ower]: https://gitlab.com/gdower
[MIT license]: https://gitlab.com/gogna/gnparser/raw/master/LICENSE
[parser-web]: https://parser.globalnames.org
