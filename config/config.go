package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	Goal int
	Dir  string
}

func New(configDir string, goal int) Config {
	var cfg Config

	if configDir == "" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		os.MkdirAll(dir, 0700)
		cfg.Dir = filepath.Join(dir, ".wad")
	} else {
		if configDir[:2] == "~/" {
			usr, _ := user.Current()
			dir := usr.HomeDir
			configDir = strings.Replace(configDir, "~/", dir+"/", 1)
		}
		path, _ := filepath.Abs(configDir)
		os.MkdirAll(path, 0700)
		cfg.Dir = path
	}

	cfg.Goal = goal

	return cfg
}

func (c *Config) Write() {
	configFile := filepath.Join(c.Dir, "config.json")
	config := map[string]int{"goal": c.Goal}
	json, _ := json.Marshal(config)
	ioutil.WriteFile(configFile, json, 0644)
}

// func WadGoal() (int, error) {
// 	var c Config

// 	content, _ := ioutil.ReadFile(filepath.Join(WadDir(), "config.json"))
// 	if err := json.Unmarshal(content, &c); err != nil {
// 		return 1, err
// 	}

// 	return c.Goal, nil
// }
