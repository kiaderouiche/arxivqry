package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:   "exp-run",
		Usage:  "Expriment running",
		Action: Exprun,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "test, t",
				Value: "Run",
			},
		},
	},
}
