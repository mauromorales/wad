package progresstracker

import (
	"encoding/json"
	"io/ioutil"

	config "github.com/mauromorales/wad/config"
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

func (p *progressTracker) TrackFile(date string, file string, words int) {
	goal, _ := config.WadGoal()

	if _, ok := p.Dates[date]; ok {
		d := p.Dates[date]
		d.Goal = goal
		p.Dates[date] = d
		p.Dates[date].Files[file] = words
	} else {
		p.Dates[date] = Date{goal, make(map[string]int)}
		p.Dates[date].Files[file] = words
	}
}
