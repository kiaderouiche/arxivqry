package main

import (
	"os"

	"github.com/kiaderouiche/arxivqry/commands"
	"github.com/urfave/cli"
)

var (
	_appName    string = "arxivqry"
	_appUsage   string = "Get publication metadata from arxiv.org"
	_appVersion string = "0.0.1pre1"
	_appAuthor  string = "K.I.A.Derouiche"
	_appEmail   string = "kamel.derouiche@gmail.com"
)

func main() {

	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = _appName
	app.Usage = _appUsage
	app.Version = _appVersion
	app.Author = _appAuthor
	app.Email = _appEmail
	app.Commands = commands.Commands
	return app
}
