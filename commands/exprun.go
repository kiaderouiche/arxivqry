package arxivcmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func Exprun(c *cli.Context) error {
	fmt.Println("First running...", c.String("name"))
	return nil
}
