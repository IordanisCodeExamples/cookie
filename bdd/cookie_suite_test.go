package bdd_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cookie Command Suite")
}

var _ = Describe("Cookie Command", func() {
	Describe("Executing the cookie program", func() {
		Context("Given a valid file and date", func() {
			It("Should output the correct cookie", func() {
				cmd := exec.Command("go", "run", "../cmd/cookie/main.go", "-f", "../examples/one.csv", "-d", "2018-12-09")

				var stdout, stderr bytes.Buffer
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr

				err := cmd.Run()
				result := strings.TrimSpace(stdout.String())

				if err != nil {
					Fail(fmt.Sprintf("Command failed with exit status %v. Stderr: %s", err, stderr.String()))
				}
				Expect(result).To(Equal("AtY0laUfhglK3lC7"))
			})
		})

		Context("Given a valid file and date for multiple results", func() {
			It("Should output two cookies", func() {
				cmd := exec.Command("go", "run", "../cmd/cookie/main.go", "-f", "../examples/one.csv", "-d", "2018-12-08")

				var stdout, stderr bytes.Buffer
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr

				err := cmd.Run()
				result := strings.TrimSpace(stdout.String())

				if err != nil {
					Fail(fmt.Sprintf("Command failed with exit status %v. Stderr: %s", err, stderr.String()))
				}
				Expect(result).To(Equal("SAZuXPGUrfbcn5UA\n4sMM2LxV07bPJzwf"))
			})
		})

	})
})
