package main

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rendon/testcli"
)

// Run make install before these tests to get meaningful
// results.

var _ = Describe("Main", func() {
	Describe("--version flag", func() {
		It("returns version", func() {
			c := testcli.Command("gnparser", "-v")
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).To(ContainSubstring("version:"))
		})

		It("ignores other flags returning version", func() {
			c := testcli.Command("gnparser", "-v", "-f", "simple",
				"-j", "200", "-w", "8000")
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).To(ContainSubstring("version:"))
		})
	})

	Describe("-format flag", func() {
		It("formats output", func() {
			c := testcli.Command("gnparser", "Homo sapiens", "-f", "simple")
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).To(ContainSubstring(",Homo sapiens,"))
		})
		It("is ignored with --version", func() {
			c := testcli.Command("gnparser", "Homo sapiens", "-f", "simple", "--version")
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).ToNot(ContainSubstring(",Homo sapiens,"))
			Expect(c.Stdout()).To(ContainSubstring("version:"))
		})
		It("is set to default format if -f value is unknown",
			func() {
				c := testcli.Command("gnparser", "Homo sapiens", "-f", ":)")
				c.Run()
				Expect(c.Success()).To(BeTrue())
				Expect(c.Stdout()).
					To(HavePrefix(`{"parsed":true,"quality":1,`))
			})
	})
	Describe("Stdin", func() {
		It("takes data from Stdin", func() {
			c := testcli.Command("gnparser", "-f", "simple")
			c.SetStdin(strings.NewReader("Homo sapiens"))
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).To(ContainSubstring(",Homo sapiens,"))
		})
		It("takes multiple names from Stdin", func() {
			c := testcli.Command("gnparser", "-f", "simple")
			c.SetStdin(strings.NewReader("Plantago\nBubo L.\n"))
			c.Run()
			Expect(c.Success()).To(BeTrue())
			Expect(c.Stdout()).To(ContainSubstring(",Plantago,"))
			Expect(c.Stdout()).To(ContainSubstring(",Bubo,"))
		})
	})
})
