// +build dev

package dict

import (
	"net/http"
)

var Assets http.FileSystem = http.Dir("./assets")
