package main

import (
	"github.com/urfave/cli"
	"os"
)

var (
	_appName  string = "arxquery"
	_appUsage string = ""
)

func main() {

	app := cli.NewApp()
	app.Name = _appName
	app.Usage =
		app.Run(os.Args)
}
