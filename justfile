# Install just: https://github.com/casey/just

# Project configuration
app := "gnparser"
org := "github.com/gnames/"
version := `git describe --tags`
ver := `git describe --tags --abbrev=0`
date := `date -u '+%Y-%m-%d_%H:%M:%S%Z'`

# Build flags
no_c := "CGO_ENABLED=0"
flags_shared := "GOARCH=amd64"
flags_linux := "GOARCH=amd64 GOOS=linux"
flags_linux_arm := "GOARCH=arm64 GOOS=linux"
flags_mac := "GOARCH=amd64 GOOS=darwin"
flags_mac_arm := "GOARCH=arm64 GOOS=darwin"
flags_win := "GOARCH=amd64 GOOS=windows"
flags_win_arm := "GOARCH=arm64 GOOS=windows"
flags_ld := '-ldflags "-s -w -X ' + org + app + \
            '.Build=' + date + ' -X ' + org + app + \
            '.Version=' + version + '"'
flags_rel := '-trimpath -ldflags "-s -w -X ' + org + app + \
             '.Build=' + date + '"'

# Directories
release_dir := '/tmp'
build_dir := '.'
clib_dir := '.'

# Default recipe
default: install

# Run tests
test: deps install
    go test -shuffle=on -race -count=1 ./...

# Test build
test-build: deps build

# Download dependencies
deps:
    go mod download

# Install tools from tools.go
tools: deps
    @echo Installing tools from tools.go
    @cat {{app}}/tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %

# Generate PEG parsers
peg:
    cd ent/parser && \
    peg grammar.peg && \
    goimports -w grammar.peg.go
    cd ent/internal/preparser && \
    peg grammar.peg && \
    goimports -w grammar.peg.go

# Generate Ragel state machines
ragel:
    cd ent/internal/preprocess && \
    ragel -Z -G2 virus.rl && \
    ragel -Z -G2 noparse.rl

# Generate assets
asset:
    cd io/fs && \
    {{flags_shared}} go run -tags=dev assets_gen.go

# Build the project
build: peg
    cd {{app}} && \
    go clean && \
    {{no_c}} go build {{flags_ld}} -o {{build_dir}}

# Build release version
buildrel: peg
    cd {{app}} && \
    go clean && \
    {{no_c}} go build {{flags_rel}} -o {{build_dir}}

# Install the project
install: peg
    cd {{app}} && \
    go clean && \
    {{no_c}} go install {{flags_ld}}

# Create multi-platform releases
release: peg dockerhub
    #!/usr/bin/env bash
    set -euo pipefail
    cd {{app}}

    # Linux x86
    go clean
    {{flags_linux}} {{no_c}} go build {{flags_ld}}
    tar zcf {{release_dir}}/{{app}}-{{ver}}-linux-x86.tar.gz {{app}}

    # Linux ARM
    go clean
    {{flags_linux_arm}} {{no_c}} go build {{flags_ld}}
    tar zcf {{release_dir}}/{{app}}-{{ver}}-linux-arm.tar.gz {{app}}

    # Mac x86
    go clean
    {{flags_mac}} {{no_c}} go build {{flags_ld}}
    tar zcf {{release_dir}}/{{app}}-{{ver}}-mac-x86.tar.gz {{app}}

    # Mac ARM
    go clean
    {{flags_mac_arm}} {{no_c}} go build {{flags_ld}}
    tar zcf {{release_dir}}/{{app}}-{{ver}}-mac-arm.tar.gz {{app}}

    # Windows x86
    go clean
    {{flags_win}} {{no_c}} go build {{flags_ld}}
    zip -9 {{release_dir}}/{{app}}-{{ver}}-win-x86.zip {{app}}.exe

    # Windows ARM
    go clean
    {{flags_win_arm}} {{no_c}} go build {{flags_ld}}
    zip -9 {{release_dir}}/{{app}}-{{ver}}-win-arm.zip {{app}}.exe

    go clean

# Build with docker-compose
dc: asset build
    docker-compose build

# Build Docker image
docker: build
    docker build -t gnames/go{{app}}:latest -t gnames/go{{app}}:{{version}} .
    cd {{app}} && go clean

# Push to Docker Hub
dockerhub: docker
    docker push gnames/go{{app}}
    docker push gnames/go{{app}}:{{version}}

# Build C library for Darwin (macOS universal binary)
clib_darwin: peg
    #!/usr/bin/env bash
    set -euo pipefail
    cd binding
    go clean
    CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build {{flags_ld}} -buildmode=c-shared -o {{clib_dir}}/lib{{app}}_arm64.so
    CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build {{flags_ld}} -buildmode=c-shared -o {{clib_dir}}/lib{{app}}_amd64.so
    rm lib{{app}}_amd64.h
    mv lib{{app}}_arm64.h lib{{app}}.h
    lipo -create -output {{clib_dir}}/lib{{app}}.so {{clib_dir}}/lib{{app}}_arm64.so {{clib_dir}}/lib{{app}}_amd64.so

# Build C library
clib: peg
    cd binding && \
    go build {{flags_ld}} -buildmode=c-shared -o {{clib_dir}}/lib{{app}}.so

# Generate quality report
quality:
    cd tools && \
    go run quality.go > ../quality.md

# Generate man page
man: ronn
    @ronn ./man/{{app}}.1.ronn --style=dark

# Ensure ronn is installed
ronn:
    @which ronn > /dev/null || gem install ronn
