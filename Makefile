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

RELEASE_DIR ?= "/tmp"
BUILD_DIR ?= "."
CLIB_DIR ?= "."

all: install

test: deps install
	$(FLAG_MODULE) go test -race ./...

test-build: deps build

deps:
	$(GOCMD) mod download;

tools: deps
	@echo Installing tools from tools.go
	@cat gnparser/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

peg:
	cd ent/parser; \
	peg grammar.peg; \
	goimports -w grammar.peg.go; \

ragel:
	cd ent/internal/preprocess; \
	ragel -Z -G2 virus.rl; \
	ragel -Z -G2 noparse.rl

asset:
	cd io/fs; \
	$(FLAGS_SHARED) go run -tags=dev assets_gen.go

build: peg
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(NO_C) $(GOBUILD) -o $(BUILD_DIR)

install: peg
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(NO_C) $(GOINSTALL)

release: peg dockerhub
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_LINUX) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/gnparser-$(VER)-linux.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_MAC) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/gnparser-$(VER)-mac.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_WIN) $(NO_C) $(GOBUILD); \
	zip -9 $(RELEASE_DIR)/gnparser-$(VER)-win-64.zip gnparser.exe; \
	$(GOCLEAN);

nightly: peg
	cd gnparser; \
	$(GOCLEAN); \
	$(FLAGS_LINUX) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/gnparser-linux.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_MAC) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/gnparser-mac.tar.gz gnparser; \
	$(GOCLEAN); \
	$(FLAGS_WIN) $(NO_C) $(GOBUILD); \
	zip -9 $(RELEASE_DIR)/gnparser-win-64.zip gnparser.exe; \
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
	$(GOBUILD) -buildmode=c-shared -o $(CLIB_DIR)/libgnparser.so;

quality:
	cd tools;\
	$(GOCMD) run quality.go > ../quality.md


.PHONY: man
man: ronn
	@ronn ./man/gnparser.1.ronn --style=dark

.PHONY: ronn
ronn:
	@which ronn > /dev/null || gem install ronn
