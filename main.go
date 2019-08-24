package main

import (
	"os"

	"github.com/kiaderouiche/arxivqry/commands"
	"github.com/urfave/cli"
)

func info() {
	app.Name = "arxivqry"
	app.Usage = "Get publication metadata from arxiv.org"
	app.Version = "0.0.1pre2"
	app.Author = "K.I.A.Derouiche"
	app.Email = "kamel.derouiche@gmail.com"
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Commands = commands.Commands
	return app
}

func main() {

	newApp().Run(os.Args)
}
