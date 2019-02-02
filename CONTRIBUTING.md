# How to contribute to ``gnparser`` project

## **Did you find a bug?**

* **Ensure the bug was not already reported** by searching on GitLab under
  [Issues](https://gitlab.com/gogna/gnparser/issues).

* If you're unable to find an open issue addressing the problem, [open a new
  one](https://gitlab.com/gogna/gnparser/issues/new). Be sure to include a
  **title and clear description**, as much relevant information as possible,
  and a **code sample** or an **executable test case** via
  [https:parser.globalnames.org](https://parser.globalnames.org) demonstrating
  the expected behavior that is not occurring.
* Make sure you **do not put more than one bug report** in the new issue.

## **Do you intend to add a new feature or change an existing one?**

* Suggest your change in the [GlobalNames gitter
  group](https://gitter.im/GlobalNamesArchitecture/GlobalNames), or [create an
  issue](https://gitlab.com/gogna/gnparser/issues/new) that describes your
  suggestion in detail.
* Make sure you **do not put more than one feature or change** in the new issue.


## **Did you write a patch that fixes a bug?**

* Open a new GitHub pull request with the patch.

* Ensure the PR description clearly describes the problem and solution. Include
  the relevant issue number if applicable.

* Clearly state if your PR is a proof of concept and what needs to be done to
  finish it, or, if it is ready to merge patch with tests and documentation
  added.

## **Did you write a client for your favorite language to access ``gnparser`` functionality via gRPC method calls?**

Let us know about your client on [GlobalNames gitter
group](https://gitter.im/GlobalNamesArchitecture/GlobalNames).

## **Do you have questions about the source code?**

* Ask any question on the [GlobalNames gitter
  group](https://gitter.im/GlobalNamesArchitecture/GlobalNames)

## **Would you like to contribute, but do not know how?**

* Read the next section about configuring environment for the project.

## **Setting up ``gnparser`` programming environment**

### Introduction

``gnparser`` uses several external tools and technologies:

1. [Parsing Expression Grammar tool](https://github.com/pointlander/peg) to
   generate parsing code.

2. [Protobuf/gRPC](https://grpc.io/) to facilitate remote method calls from a
   variety of programming languages.

3. [Cobra CLI framework](https://github.com/spf13/cobra) for creating command
   line application.

4. [Ginkgo testing framework](https://github.com/onsi/ginkgo) for
   Behavior-Driven Development and testing.

5. [Virtual File System Generator](https://github.com/shurcooL/vfsgen)
   for including dictionaries and HTML static files or templates.

Most of these projects are installed by one command, but you do need to
setup Protobuf's ``protoc`` binary.

You can find binaries for ``protoc`` for Mac, Linux or Windows on the
[``protobuf`` releases page](https://github.com/protocolbuffers/protobuf/releases)

#### Installation of protoc on Mac

```bash
brew install protobuf

# or

brew upgrade protobuf
```

or

```bash
PROTOC_ZIP=protoc-3.6.1-osx-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
rm -f $PROTOC_ZIP
```

#### Installation of ``protoc`` on Linux

Use our [script for continuous integration](https://gitlab.com/gogna/gnparser/blob/master/scripts/protoc-install.sh)
as a guide.

### Install Go

[Download and install Go](https://golang.org/doc/install) for your operating system. Make sure you
[configured GOPATH environment library](https://github.com/golang/go/wiki/SettingGOPATH).

You need Go v1.11.x or higher, because we use recently introduced Go modules.

### Install ``gnparser`` code

Before Go v1.11 all Go code had to be organized inside of the ``GOPATH``
directory. Now, for projects like ``gnparser`` that use Go modules it is not
necessary, however many tools still behave assuming old ways, so we recommend
to setup ``gnparser`` code traditional way.

```bash
mkdir -p $GOPATH/src/gitlab.com/gogna
cd $GOPATH/src/gitlab.com/gogna
git clone https://gitlab.com/gogna/gnparser.git
# or use URL of your fork on gitlab or github

cd gnparser
# to download all dependencies
make deps

# to make gnparser executable and place it to $GOPATH/bin
make

# now you should be able to use gnparser compiled from the code:
gnparser -f pretty "Pica pica (Linnaeus, 1758)"
```

this should install all the tools and dependencies to test and run ``gnparser``.
You can check it by running

```bash
make test
```

Note that you would need to use ``GO111MODULE=on`` to run tests, or,
alternatively, you can install all dependencies in a 'traditional' way at
``GOPATH/src`` and ``GOPATH/bin``.

```bash
GO111MODULE=on go test
# or
GO111MODULE=on ginkgo

# to run all tests
GO111MODULE=on go test ./...
# or
GO111MODULE=on ginkgo ./...

# to run tests continuously
GO111MODULE=on ginkgo watch
```
