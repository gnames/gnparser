// +build ignore

package main

import (
	"log"

	"github.com/gnames/gnparser/fs"
	"github.com/shurcool/vfsgen"
)

func main() {
	err := vfsgen.Generate(fs.Assets, vfsgen.Options{
		PackageName:  "fs",
		BuildTags:    "!dev",
		VariableName: "Files",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
