package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

import . "github.com/tj/go-debug"

var debug = Debug("imgops")

var authors = []cli.Author{
	{
		Name:  "Doğan Çelik",
		Email: "dogancelik.com",
	},
}

var Version string

func cliSelect() string {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println(genSelectText())
	ret := ""

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}

		m := getKeyToNameTargets(availableTargets)
		target, mapOk := m[char]
		if mapOk {
			return target
		} else if char == 'i' {
			return defaultTarget
		}
	}

	return ret
}

func cliSearch(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.NewExitError("No file or URL is given", 1)
	}

	srcPath := c.Args().First()
	targets := c.String("targets")

	debug("Path: %s", srcPath)
	debug("Targets: %s", targets)

	var urls []string
	var errUpload error

	// Select flag
	if c.Bool("select") {
		targets = cliSelect()
		if targets == "" {
			return cli.NewExitError("Upload cancelled", 4)
		}
	}

	// Upload
	if isUrl(srcPath) == true {
		debug("Start URL upload")
		urls, errUpload = UploadURL(srcPath, targets)
	} else {
		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
			return cli.NewExitError("File doesn't exist: "+srcPath, 2)
		}
		debug("Start file upload")
		urls, errUpload = UploadFile(srcPath, targets)
	}

	// Upload result
	if errUpload != nil && len(urls) == 0 {
		return cli.NewExitError("Error during upload: "+errUpload.Error(), 3)
	} else {
		if errUpload != nil {
			fmt.Fprintf(os.Stderr, "Unknown targets '%s', will open default page instead", targets)
		}

		for _, url := range urls {
			if c.Bool("return") {
				fmt.Println(url)
			} else {
				open.Start(url)
			}
		}
	}

	return nil
}

func cliMain(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ImgOps CLI"
	app.Usage = "Reverse search images"
	app.Version = Version
	app.Authors = authors
	app.Action = cliMain
	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"a"},
			Usage:   "Search a file or a URL",
			Action:  cliSearch,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "targets, t",
					Value: defaultTarget,
					Usage: "Target website to search at (e.g. Google)",
				},
				cli.BoolFlag{
					Name:  "select, s",
					Usage: "Show a list of targets to select from",
				},
				cli.BoolFlag{
					Name:  "return, r",
					Usage: "Output the result URL",
				},
			},
		},
	}
	app.Run(os.Args)
}
