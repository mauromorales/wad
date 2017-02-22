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

	Describe("GetFiles", func() {
		It("Returns a list of files with their word count from the last day they were tracked", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"/file1.txt":{"2017-02-18":458,"2017-02-19":500},"/file2.txt":{"2017-02-18":100}}`)
			file.Write(fileContent)

			fileManager := NewFileManager(file.Name())
			files := fileManager.GetFiles()

			expectedFiles := [][]string{
				[]string{"/file1.txt", "500", "2017-02-19"},
				[]string{"/file2.txt", "100", "2017-02-18"},
			}

			Expect(len(files)).To(Equal(2))
			Expect(files).To(Equal(expectedFiles))
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
