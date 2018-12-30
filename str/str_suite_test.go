package str_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Str Suite")
}
