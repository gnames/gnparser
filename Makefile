GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean

VERSION=`git describe --tags`
VER=`git describe --tags --abbrev=0`
DATE=`date -u '+%Y-%m-%d_%H:%M:%S%Z'`

all: install

t: install
	time gnparser test-data/200k-lines.txt -f simple > t

build: version peg
	cd gnparser; \
	$(GOCLEAN); \
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD)

install: version peg
	cd gnparser; \
	$(GOCLEAN); \
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOINSTALL)

peg:
	cd grammar; \
	peg grammar.peg; \
	goimports -w grammar.peg.go


version:
	echo "package output\n\nconst Version = \"$(VERSION)\"\nconst Build = \"$(DATE)\"\n" \
	> output/version.go
