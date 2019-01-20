// +build ignore

package main

import (
	"log"

	"github.com/shurcool/vfsgen"
	"gitlab.com/gogna/gnparser/fs"
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
