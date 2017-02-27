package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
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
	if len(f.Files) == 0 {
		f.Files = make(map[string]map[string]int)
	}

	if _, ok := f.Files[path]; !ok {
		f.Files[path] = make(map[string]int)
	}
	f.Files[path][date] = words

	return words, nil
}

func (f *FileManager) GetFiles(current bool, search ...string) [][]string {
	sort.Strings(search)

	var result [][]string

	var files []string
	for file := range f.Files {
		files = append(files, file)
	}
	sort.Strings(files)
	today := time.Now().Local().Format("2006-01-02")
	var lastDate string

	for _, file := range files {
		var dates []string

		for date := range f.Files[file] {
			if current || date != today {
				dates = append(dates, date)
			}
		}
		sort.Strings(dates)

		if len(dates) > 0 {
			lastDate = dates[len(dates)-1]
		} else {
			lastDate = today
		}

		if len(search) == 0 {
			result = append(result, []string{file, strconv.Itoa(f.Files[file][lastDate]), lastDate})
		} else if i := sort.SearchStrings(search, file); i < len(search) && search[i] == file {
			result = append(result, []string{file, strconv.Itoa(f.Files[file][lastDate]), lastDate})
		}
	}

	return result
}
