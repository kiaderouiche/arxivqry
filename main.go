package main

import (
	"os"

	"github.com/urfave/cli"
)

func newApp() *cli.App {

	app := cli.NewApp()

	return app
}

func main() {

	newApp().Run(os.Args)
}
