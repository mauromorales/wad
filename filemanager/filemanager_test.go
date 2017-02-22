package filemanager_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

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
			files := fileManager.GetFiles(true)

			expectedFiles := [][]string{
				[]string{"/file1.txt", "500", "2017-02-19"},
				[]string{"/file2.txt", "100", "2017-02-18"},
			}

			Expect(len(files)).To(Equal(2))
			Expect(files).To(Equal(expectedFiles))
		})

		Context("When the current state is false", func() {
			It("Returns the previous tracked date instead", func() {

				today := time.Now().Local().Format("2006-01-02")
				file, _ = ioutil.TempFile("", "wad")
				fileContent = []byte(fmt.Sprintf(`{"/file1.txt":{"2017-02-18":458,"%v":500}}`, today))
				file.Write(fileContent)

				fileManager := NewFileManager(file.Name())
				files := fileManager.GetFiles(false)

				expectedFiles := [][]string{
					[]string{"/file1.txt", "458", "2017-02-18"},
				}

				Expect(len(files)).To(Equal(1))
				Expect(files).To(Equal(expectedFiles))
			})
		})

		It("Returns only the requested files with their word count from the last day they were tracked", func() {
			file, _ = ioutil.TempFile("", "wad")
			fileContent = []byte(`{"/file1.txt":{"2017-02-18":458,"2017-02-19":500},"/file2.txt":{"2017-02-18":100},"/file3.txt":{"2017-02-19":150}}`)
			file.Write(fileContent)

			fileManager := NewFileManager(file.Name())
			files := fileManager.GetFiles(true, "/file3.txt", "/file1.txt")

			expectedFiles := [][]string{
				[]string{"/file1.txt", "500", "2017-02-19"},
				[]string{"/file3.txt", "150", "2017-02-19"},
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
