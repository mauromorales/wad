package progresstracker

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/mauromorales/wad/config"
	"github.com/mauromorales/wad/filemanager"
)

type Date struct {
	Goal  int
	Files map[string]int
}

type progressTracker struct {
	Dates map[string]Date
}

func NewProgressTracker(filePath string) (progressTracker, error) {
	var p progressTracker

	content, _ := ioutil.ReadFile(filePath)
	if err := json.Unmarshal(content, &p.Dates); err != nil {
		return p, err
	}

	return p, nil
}

func (p *progressTracker) TrackFile(fm filemanager.FileManager, date string, file string, words int) {
	goal, _ := config.WadGoal()

	if _, ok := p.Dates[date]; ok {
		d := p.Dates[date]
		d.Goal = goal
		p.Dates[date] = d
		if len(fm.GetFiles(false, file)) > 0 {
			lastWordCount, _ := strconv.Atoi(fm.GetFiles(false, file)[0][1])
			p.Dates[date].Files[file] = words - lastWordCount
		} else {
			p.Dates[date].Files[file] = words
		}
	} else {
		p.Dates[date] = Date{goal, make(map[string]int)}
		if len(fm.GetFiles(false, file)) > 0 {
			lastWordCount, _ := strconv.Atoi(fm.GetFiles(false, file)[0][1])
			p.Dates[date].Files[file] = words - lastWordCount
		} else {
			p.Dates[date].Files[file] = words
		}
	}
}
