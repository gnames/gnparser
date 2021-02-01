// +build ignore

// quality.go generates a markdown file that describes meaning of each quality
// category.
package main

import (
	"fmt"
	"github.com/gnames/gnparser/ent/parsed"
	"sort"
)

var body = `# Quality categories

## Quality 0

Parsing failed.

## Quality 1

Parsing finished without detecting any problems.`

func main() {
	warnsMap := make(map[int][]string)
	for k, v := range parsed.WarningQualityMap {
		warnsMap[v] = append(warnsMap[v], k.String())
	}

	for _, v := range []int{2, 3, 4} {
		warns := warnsMap[v]
		sort.Strings(warns)
		item := fmt.Sprintf("\n\n## Quality %d\n", v)
		for i := range warns {
			warn := fmt.Sprintf("\n- %s", warns[i])
			item += warn
		}
		body += item
	}
	fmt.Println(body)
}
