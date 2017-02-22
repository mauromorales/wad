package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"

	config "github.com/mauromorales/wad/config"
	fm "github.com/mauromorales/wad/filemanager"
	pt "github.com/mauromorales/wad/progresstracker"
)

func main() {
	app := cli.NewApp()
	app.Name = "wad"
	app.Usage = "A CLI tool to help you write"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "initializes WAD",
			Action: func(c *cli.Context) error {
				os.MkdirAll(config.WadDir(), 0700)

				configFile := filepath.Join(config.WadDir(), "config.json")
				filesFile := filepath.Join(config.WadDir(), "files.json")
				progressFile := filepath.Join(config.WadDir(), "progress.json")

				if _, err := os.Stat(configFile); err == nil {
					return cli.NewExitError("WAD has already been initialized in the given directory.", 1)
				} else {
					config := map[string]int{"goal": 500}

					json, _ := json.Marshal(config)
					ioutil.WriteFile(configFile, json, 0644)

					ioutil.WriteFile(filesFile, []byte("{}"), 0644)
					ioutil.WriteFile(progressFile, []byte("{}"), 0644)
				}

				return nil

			},
		},
		{
			Name:    "files",
			Aliases: []string{"f"},
			Usage:   "track a given file",
			Action: func(c *cli.Context) error {
				filesFile := filepath.Join(config.WadDir(), "files.json")
				fileManager := fm.NewFileManager(filesFile)

				data := fileManager.GetFiles(true)

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"File", "Words", "Last tracked on"})

				for _, v := range data {
					table.Append(v)
				}
				table.Render()

				return nil
			},
		},
		{
			Name:    "track",
			Aliases: []string{"t"},
			Usage:   "track a given file",
			Action: func(c *cli.Context) error {
				filesFile := filepath.Join(config.WadDir(), "files.json")
				fileManager := fm.NewFileManager(filesFile)

				current_time := time.Now().Local()
				file, _ := filepath.Abs(c.Args().First())
				words, err := fileManager.Track(file, current_time.Format("2006-01-02"))
				if err != nil {
					return cli.NewExitError(err, 1)
				}

				jsonContent, _ := json.Marshal(fileManager.Files)
				ioutil.WriteFile(filesFile, jsonContent, 0644)

				progressFile := filepath.Join(config.WadDir(), "progress.json")
				progressTracker, _ := pt.NewProgressTracker(progressFile)
				progressTracker.TrackFile(fileManager, current_time.Format("2006-01-02"), file, words)

				jsonContent, _ = json.Marshal(progressTracker.Dates)
				ioutil.WriteFile(progressFile, jsonContent, 0644)

				return nil
			},
		},
		{
			Name:    "progress",
			Aliases: []string{"p"},
			Usage:   "display the current progress",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "from",
				},
			},
			Action: func(c *cli.Context) error {

				progressFile := filepath.Join(config.WadDir(), "progress.json")
				progressTracker, _ := pt.NewProgressTracker(progressFile)
				data := [][]string{}
				lastGoal := -1

				now := time.Now().Local()
				from := now.AddDate(0, 0, -7)
				for !from.After(now) {
					var total int
					current_date := from.Format("2006-01-02")

					if day, ok := progressTracker.Dates[current_date]; ok {
						for _, words := range day.Files {
							total += words
						}

						var status string
						if total >= day.Goal {
							status = "+"
						} else {
							status = "-"
						}

						lastGoal = day.Goal

						data = append(data, []string{current_date, strconv.Itoa(total), strconv.Itoa(day.Goal), status})
					} else {
						var goal string

						if lastGoal < 0 {
							goal = "n/a"
						} else {
							goal = strconv.Itoa(lastGoal)
						}
						data = append(data, []string{current_date, "0", goal, "n/a"})
					}

					from = from.AddDate(0, 0, 1)
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Day", "Words", "Goal", "Status"})

				for _, v := range data {
					table.Append(v)
				}
				table.Render()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
