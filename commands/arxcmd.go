package arxivcmd

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:   "expr-run",
		Usage:  "Exp run command",
		Action: Exprun,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "Test, t",
				Value: "get author data",
			},
		},
	},
}
