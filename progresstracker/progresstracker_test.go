package progresstracker_test

import (
	"io/ioutil"
	"os"

	"github.com/mauromorales/wad/filemanager"

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

			fm := filemanager.NewFileManager(file.Name())

			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			pt.TrackFile(fm, "2017-02-18", "/path/to/file.txt", 458)
			Expect(len(pt.Dates)).To(Equal(1))
		})

		It("Updates an entry in the same day", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{}`)
			file.Write(fileContent)

			fm := filemanager.NewFileManager(file.Name())

			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"2017-02-18": {"goal":500, "files":{"/path/to/file.txt":458}}}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			pt.TrackFile(fm, "2017-02-18", "/path/to/file.txt", 600)
			Expect(len(pt.Dates)).To(Equal(1))
			Expect(pt.Dates["2017-02-18"].Files["/path/to/file.txt"]).To(Equal(600))
		})

		It("Only tracks the difference from the previous tracking", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"/file.txt":{"2017-02-18":458}}`)
			file.Write(fileContent)

			fm := filemanager.NewFileManager(file.Name())

			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"2017-02-18": {"goal":500, "files":{"/file.txt":458}}}`)
			file.Write(fileContent)

			pt, _ := NewProgressTracker(file.Name())

			pt.TrackFile(fm, "2017-02-19", "/file.txt", 700)
			Expect(len(pt.Dates)).To(Equal(2))
			Expect(pt.Dates["2017-02-19"].Files["/file.txt"]).To(Equal(242))
		})
	})
})
