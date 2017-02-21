package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type FileManager struct {
	Files map[string]map[string]int
}

func NewFileManager(filePath string) FileManager {
	var f FileManager

	content, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal(content, &f.Files)
	return f
}

func (f *FileManager) Track(path string, date string) (int, error) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return -1, err
	}

	content := strings.Fields(string(bytes))
	words := len(content)

	if _, ok := f.Files[path]; ok {
		f.Files[path][date] = words
	} else {
		f.Files[path] = make(map[string]int)
		f.Files[path][date] = words
	}

	return words, nil
}
