// +build dev

package fs

import (
	"net/http"
)

var Assets http.FileSystem = http.Dir("./assets")
