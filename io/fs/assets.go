// +build dev

package fs

import (
	"net/http"
)

// Assets represent virtual filesystem implemented as http.FileSystem.
var Assets http.FileSystem = http.Dir("./assets")
