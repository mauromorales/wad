package progresstracker_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mauromorales/wad/progresstracker"
)

var _ = Describe("ProgressTracker", func() {
	var (
		file        *os.File
		fileContent []byte
	)

	Describe("NewProgressTracker", func() {
		It("Loads progress.json into the struct", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"2017-02-18": {"goal":500, "files":{"/path/to/file.txt":458}}}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			Expect(len(pt.Dates)).To(Equal(1))
		})
	})

	Describe("TrackFile", func() {
		It("Tracks the number of words for a file on a given day", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			pt.TrackFile("2017-02-18", "/path/to/file.txt", 458)
			Expect(len(pt.Dates)).To(Equal(1))
		})

		It("Updates an entry in the same day", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"2017-02-18": {"goal":500, "files":{"/path/to/file.txt":458}}}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			pt.TrackFile("2017-02-18", "/path/to/file.txt", 600)
			Expect(len(pt.Dates)).To(Equal(1))
			Expect(pt.Dates["2017-02-18"].Files["/path/to/file.txt"]).To(Equal(600))
		})
	})
})
