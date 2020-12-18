package output

import "encoding/json"

type Details interface {
	isDetails()
}

type hybridElement interface {
	hybridElement() bool
}
