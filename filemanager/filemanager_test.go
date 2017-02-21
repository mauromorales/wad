package filemanager_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mauromorales/wad/filemanager"
)

var _ = Describe("FileManager", func() {
	var (
		fileManager FileManager
		file        *os.File
		fileContent []byte
	)

	Describe("NewFileManager", func() {
		It("Loads files.json into the struct", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"/path/to/file.txt":{"2017-02-18":458}}`)
			file.Write(fileContent)

			fileManager := NewFileManager(file.Name())

			Expect(len(fileManager.Files)).To(Equal(1))
		})
	})

	Describe("Track", func() {
		var (
			expectedFileStatus map[string]int
		)

		BeforeEach(func() {
			fileManager.Files = map[string]map[string]int{}
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte("one two three")
			file.Write(fileContent)
		})

		AfterEach(func() {
			os.Remove(file.Name())
		})

		It("Starts traking a new file", func() {
			expectedFileStatus = map[string]int{"2017-02-18": 3}

			fileManager.Track(file.Name(), "2017-02-18")

			Expect(len(fileManager.Files)).To(Equal(1))
			Expect(fileManager.Files[file.Name()]).To(Equal(expectedFileStatus))
		})

		It("Updates an existing checkin", func() {
			fileContent = []byte(" four five")
			file.Write(fileContent)
			fileManager.Track(file.Name(), "2017-02-18")
			expectedFileStatus = map[string]int{"2017-02-18": 5}

			fileManager.Track(file.Name(), "2017-02-18")

			Expect(len(fileManager.Files)).To(Equal(1))
			Expect(fileManager.Files[file.Name()]).To(Equal(expectedFileStatus))
		})

		It("Adds a new check in", func() {
			fileManager.Track(file.Name(), "2017-02-18")

			fileContent = []byte(" four five")
			file.Write(fileContent)
			fileManager.Track(file.Name(), "2017-02-19")
			expectedFileStatus = map[string]int{
				"2017-02-18": 3,
				"2017-02-19": 5,
			}

			Expect(len(fileManager.Files)).To(Equal(1))
			Expect(fileManager.Files[file.Name()]).To(Equal(expectedFileStatus))
		})

		It("Errors when the file does not exist", func() {
			_, err := fileManager.Track("/some/non/existing/file.txt", "2017-02-18")
			Expect(err).To(HaveOccurred())
		})
	})
})
