//go:build tools
// +build tools

package main

import (
	_ "github.com/pointlander/peg"
	_ "github.com/spf13/cobra"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "golang.org/x/tools/cmd/goimports"
)
