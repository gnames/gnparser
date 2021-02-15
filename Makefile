VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell date -u '+%Y-%m-%d_%H:%M:%S%Z')

FLAG_MODULE = GO111MODULE=on
FLAGS_SHARED = $(FLAG_MODULE) GOARCH=amd64
NO_C = CGO_ENABLED=0
FLAGS_LINUX = $(FLAGS_SHARED) GOOS=linux
FLAGS_MAC = $(FLAGS_SHARED) GOOS=darwin
FLAGS_WIN = $(FLAGS_SHARED) GOOS=windows
FLAGS_LD=-ldflags "-s -w -X github.com/gnames/gnparser.Build=${DATE} \
                  -X github.com/gnames/gnparser.Version=${VERSION}"
GOCMD = go
GOBUILD = $(GOCMD) build $(FLAGS_LD)
GOINSTALL = $(GOCMD) install $(FLAGS_LD)
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get

all: install

test: deps install
	$(FLAG_MODULE) go test ./...

test-build: deps build

deps:
	$(GOCMD) mod download;

peg:
	cd ent/parser; \
	peg grammar.peg; \
	goimports -w grammar.peg.go; \

ragel:
	cd ent/preprocess; \
	ragel -Z -G2 virus.rl; \
	ragel -Z -G2 noparse.rl

asset:
	cd io/fs; \
	$(FLAGS_SHARED) go run -tags=dev assets_gen.go

build: peg
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(NO_C) $(GOBUILD)

install: peg
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(NO_C) $(GOINSTALL)

release: peg dockerhub
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_LINUX) $(NO_C) $(GOBUILD); \
	tar zcf /tmp/gnparser-$(VER)-linux.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_MAC) $(NO_C) $(GOBUILD); \
	tar zcf /tmp/gnparser-$(VER)-mac.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_WIN) $(NO_C) $(GOBUILD); \
	zip -9 /tmp/gnparser-$(VER)-win-64.zip gnparser.exe; \
	$(GOCLEAN);

dc: asset build
	docker-compose build;

docker: build
	docker build -t gnames/gognparser:latest -t gnames/gognparser:$(VERSION) .; \
	cd gnparser; \
	$(GOCLEAN);

dockerhub: docker
	docker push gnames/gognparser; \
	docker push gnames/gognparser:$(VERSION)

clib: peg
	cd binding; \
	$(GOBUILD) -buildmode=c-shared -o libgnparser.so;

quality:
	cd tools;\
	$(GOCMD) run quality.go > ../quality.md


.PHONY: man
man: ronn
	@ronn ./man/gnparser.1.ronn --style=dark

.PHONY: ronn
ronn:
	@which ronn > /dev/null || gem install ronn
