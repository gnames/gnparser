PROJ_NAME = gnparser

VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell date -u '+%Y-%m-%d_%H:%M:%S%Z')

NO_C = CGO_ENABLED=0
FLAGS_SHARED = GOARCH=amd64
FLAGS_LINUX = GOARCH=amd64 GOOS=linux
FLAGS_LINUX_ARM = GOARCH=arm64 GOOS=linux
FLAGS_MAC = GOARCH=amd64 GOOS=darwin
FLAGS_MAC_ARM = GOARCH=arm64 GOOS=darwin
FLAGS_WIN = GOARCH=amd64 GOOS=windows
FLAGS_WIN_ARM = GOARCH=arm64 GOOS=windows
FLAGS_LD=-ldflags "-s -w -X github.com/gnames/$(PROJ_NAME).Build=$(DATE) \
                  -X github.com/gnames/$(PROJ_NAME).Version=$(VERSION)"
FLAGS_REL = -trimpath -ldflags "-s -w \
						-X github.com/gnames/$(PROJ_NAME).Build=$(DATE)"

GOCMD = go
GOBUILD = $(GOCMD) build $(FLAGS_LD)
GOINSTALL = $(GOCMD) install $(FLAGS_LD)
GORELEASE = $(GOCMD) build $(FLAGS_REL)
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get

RELEASE_DIR ?= "/tmp"
BUILD_DIR ?= "."
CLIB_DIR ?= "."

all: install

test: deps install
	$(FLAG_MODULE) go test -shuffle=on -race -count=1 ./...

test-build: deps build

deps:
	$(GOCMD) mod download;

tools: deps
	@echo Installing tools from tools.go
	@cat $(PROJ_NAME)/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

peg:
	cd ent/parser; \
	peg grammar.peg; \
	goimports -w grammar.peg.go; \
	cd ../internal/preparser; \
	peg grammar.peg; \
	goimports -w grammar.peg.go;

ragel:
	cd ent/internal/preprocess; \
	ragel -Z -G2 virus.rl; \
	ragel -Z -G2 noparse.rl

asset:
	cd io/fs; \
	$(FLAGS_SHARED) go run -tags=dev assets_gen.go

build: peg
	cd $(PROJ_NAME); \
	$(GOCLEAN); \
	$(NO_C) $(GOBUILD) -o $(BUILD_DIR)

buildrel: peg
	cd $(PROJ_NAME); \
	$(GOCLEAN); \
	$(NO_C) $(GORELEASE) -o $(BUILD_DIR)

install: peg
	cd $(PROJ_NAME); \
	$(GOCLEAN); \
	$(NO_C) $(GOINSTALL)

release: peg dockerhub
	cd $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_LINUX) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-linux-x86.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_LINUX_ARM) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-linux-arm.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_MAC) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-mac-x86.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_MAC_ARM) $(NO_C) $(GOBUILD); \
	tar zcf $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-mac-arm.tar.gz $(PROJ_NAME); \
	$(GOCLEAN); \
	$(FLAGS_WIN) $(NO_C) $(GOBUILD); \
	zip -9 $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-win-x86.zip $(PROJ_NAME).exe; \
	$(GOCLEAN); \
	$(FLAGS_WIN_ARM) $(NO_C) $(GOBUILD); \
	zip -9 $(RELEASE_DIR)/$(PROJ_NAME)-$(VER)-win-arm.zip $(PROJ_NAME).exe; \
	$(GOCLEAN);

dc: asset build
	docker-compose build;

docker: build
	docker build -t gnames/go$(PROJ_NAME):latest -t gnames/go$(PROJ_NAME):$(VERSION) .; \
	cd $(PROJ_NAME); \
	$(GOCLEAN);

dockerhub: docker
	docker push gnames/go$(PROJ_NAME); \
	docker push gnames/go$(PROJ_NAME):$(VERSION)

clib_darwin: peg
	cd binding; \
	$(GOCLEAN); \
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 $(GOBUILD) -buildmode=c-shared -o $(CLIB_DIR)/lib$(PROJ_NAME)_arm64.so; \
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -buildmode=c-shared -o $(CLIB_DIR)/lib$(PROJ_NAME)_amd64.so; \
	rm lib$(PROJ_NAME)_amd64.h; \
	mv lib$(PROJ_NAME)_arm64.h lib$(PROJ_NAME).h; \
	lipo -create -output $(CLIB_DIR)/lib$(PROJ_NAME).so $(CLIB_DIR)/lib$(PROJ_NAME)_arm64.so $(CLIB_DIR)/lib$(PROJ_NAME)_amd64.so;

clib: peg
	cd binding; \
	$(GOBUILD) -buildmode=c-shared -o $(CLIB_DIR)/lib$(PROJ_NAME).so;

quality:
	cd tools;\
	$(GOCMD) run quality.go > ../quality.md


.PHONY: man
man: ronn
	@ronn ./man/$(PROJ_NAME).1.ronn --style=dark

.PHONY: ronn
ronn:
	@which ronn > /dev/null || gem install ronn
