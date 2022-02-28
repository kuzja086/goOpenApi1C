package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func GenerateOpenApi(c *cli.Context) {
	fmt.Println(c.String("path"))
}
