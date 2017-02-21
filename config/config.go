package config

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type Config struct {
	Goal int
}

func WadDir() string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	return filepath.Join(dir, ".wad")
}

func WadGoal() (int, error) {
	var c Config

	content, _ := ioutil.ReadFile(filepath.Join(WadDir(), "config.json"))
	if err := json.Unmarshal(content, &c); err != nil {
		return 1, err
	}

	return c.Goal, nil
}
