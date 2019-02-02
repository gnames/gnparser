package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGnparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gnparser Suite")
}
