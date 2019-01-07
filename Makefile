GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get -u

VERSION=`git describe --tags`
VER=`git describe --tags --abbrev=0`
DATE=`date -u '+%Y-%m-%d_%H:%M:%S%Z'`

all: install

init:
	GO111MODULE=on $(GOGET) github.com/pointlander/peg@fa48cc2; \
	GO111MODULE=on $(GOGET) github.com/shurcooL/vfsgen@6a9ea43

version:
	echo "package output\n\nconst Version = \"$(VERSION)\"\nconst Build = \"$(DATE)\"\n" \
	> output/version.go

peg:
	cd grammar; \
	peg grammar.peg; \
	goimports -w grammar.peg.go; \

asset:
	cd dict; \
	go run -tags=dev assets_gen.go

build: version peg
	cd gnparser; \
	$(GOCLEAN); \
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD)

install: version peg
	cd gnparser; \
	$(GOCLEAN); \
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOINSTALL)

release: version peg asset
	cd gnparser; \
	$(GOCLEAN); \
	GGO111MODULE=on OOS=linux GOARCH=amd64 $(GOBUILD) ${LDFLAGS}; \
	tar zcvf /tmp/parser-${VER}-linux.tar.gz gnparser; \
	$(GOCLEAN); \
	GGO111MODULE=on OOS=darwin GOARCH=amd64 $(GOBUILD) ${LDFLAGS}; \
	tar zcvf /tmp/gnparser-${VER}-mac.tar.gz gnparser; \
	$(GOCLEAN); \
	GGO111MODULE=on OOS=windows GOARCH=amd64 $(GOBUILD) ${LDFLAGS}; \
	zip -9 /tmp/gnparser-${VER}-win-64.zip gnparser; \
	$(GOCLEAN);

