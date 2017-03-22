package main

import (
	"fmt"
	"os"

	"github.com/skratchdot/open-golang/open"
)

import . "github.com/tj/go-debug"

var cmdAction string = ""
var debug = Debug("imgops")
var QueryMap map[string]string = getIdMap()

var Version string

func main() {
	fmt.Printf("ImgOps %s by Doğan Çelik (dogancelik.com)\n", Version)

	cmdPath := ""
	cmdAction = ""

	if len(os.Args) > 1 {
		cmdPath = os.Args[1]
	} else {
		fmt.Println("No file or URL")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		cmdAction = os.Args[2]
	}

	debug("Path: %s", cmdPath)
	debug("Action: %s", cmdAction)
	debug("Queries: %v", QueryMap)

	var urls []string
	var errUpload error
	if isUrl(cmdPath) == true {
		debug("Start URL upload")
		urls, errUpload = UploadURL(cmdPath, cmdAction)
	} else {

		if _, err := os.Stat(cmdPath); os.IsNotExist(err) {
			fmt.Println("File doesn't exist:", cmdPath)
			os.Exit(2)
		}
		debug("Start file upload")
		urls, errUpload = UploadFile(cmdPath, cmdAction)
	}

	if errUpload != nil && len(urls) == 0 {
		fmt.Println("Error during upload:", errUpload)
		os.Exit(3)
	} else {
		if errUpload != nil {
			fmt.Printf("Warning wrong action '%s'; will open ImgOps instead\n", cmdAction)
		}

		for _, url := range urls {
			open.Start(url)
		}
	}

}
