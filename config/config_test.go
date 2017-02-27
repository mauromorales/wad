package config_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/mauromorales/wad/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("New", func() {

		Context("without a configuration dir", func() {
			It("initializes in the user's home directory", func() {
				c := New("", 500)
				Expect(strings.HasSuffix(c.Dir, ".wad")).To(BeTrue())
			})
		})

		Context("with a configuration dir", func() {
			It("initializes in the given directory", func() {
				c := New("/my/full/path", 500)
				Expect(c.Dir).To(Equal("/my/full/path"))
			})
		})
	})

	Describe("Write", func() {
		It("Saves the current configuration into config.json", func() {
			dir, _ := ioutil.TempDir("", "wad")

			defer os.RemoveAll(dir)

			c := New(dir, 500)
			c.Write()

			content, _ := ioutil.ReadFile(filepath.Join(dir, "config.json"))
			Expect(content).To(Equal([]byte(`{"goal":500}`)))
		})
	})
})
