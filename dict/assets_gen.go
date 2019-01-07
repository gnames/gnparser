// +build ignore

package main

import (
	"log"

	"github.com/shurcool/vfsgen"
	"gitlab.com/gogna/gnparser/dict"
)

func main() {
	err := vfsgen.Generate(dict.Assets, vfsgen.Options{
		PackageName:  "dict",
		BuildTags:    "!dev",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
