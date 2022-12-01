package cmd

import (
	"github.com/urfave/cli/v2"
)

// Exec Execute the bullet file.
func Exec() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {

		result := "\033[1;40mAll done: \u001B[0m"
		return cli.Exit(result, 0)
	}
}
