package cmd

import (
	"fmt"
	"github.com/duan-li/hue"
	"github.com/duan-li/sshgun/ssh"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
	"syscall"
)

// Exec Execute the bullet file.
func Exec() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		config := ssh.Config{
			Host: cCtx.String("host"),
			Port: cCtx.Int("port"),
			User: cCtx.String("username"),
		}

		if cCtx.Bool("password") == true {
			fmt.Print("Password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				hue.Fatal(err.Error(), true)
			}
			config.Password = string(bytePassword)
		}

		command := cCtx.Args().Get(0)

		if output, err := ssh.Runner(config, command); err != nil {
			return cli.Exit(err, 1)
		} else {
			return cli.Exit(output, 0)
		}

	}
}
