package web_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var stdout *os.File

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Suite")
}

var _ = BeforeSuite(func() {
	stdout = os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
})

var _ = AfterSuite(func() {
	os.Stdout = stdout
})
