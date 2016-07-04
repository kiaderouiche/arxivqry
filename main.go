package main

import (
	"os"

	"github.com/kiaderouiche/arxivquery/commands"
	"github.com/urfave/cli"
)

var (
	_appName  string = "arxquery"
	_appUsage string = "experimental arxivprog!"
)

func main() {

	app := cli.NewApp()
	app.Name = _appName
	app.Usage = _appUsage

	app.Commands = commands.Commands
	app.Run(os.Args)
}
