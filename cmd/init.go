package cmd

import (
	"github.com/urfave/cli/v2"
)

func Init() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "exec",
				Usage:  "execute a command directly.",
				Action: Exec(),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "host",
						Usage:   "host address",
						Aliases: []string{"H"},
						Value:   "127.0.0.1",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "host port",
						Value:   22,
					},
					&cli.StringFlag{
						Name:    "user",
						Aliases: []string{"u"},
						Usage:   "username",
					},
					&cli.StringFlag{
						Name:  "password",
						Usage: "password",
					},
					&cli.BoolFlag{
						Name:    "sudo",
						Aliases: []string{"s"},
						Usage:   "Enable sudo command.",
						Value:   false,
					},
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "connection config format `username@host:port`",
					},
				},
			},
			{
				Name:   "run",
				Usage:  "run a command on a ssh connection",
				Action: Run(),
			},
			{
				Name:    "bullet",
				Aliases: []string{"b"},
				Usage:   "bullet templates",
				Subcommands: []*cli.Command{
					{
						Name: "init",
						Usage: "create a bullet from template" +
							"bullet init <template_name>",
						Action: InitBullet(),
					},
					{
						Name: "validate",
						Usage: "validate bullet file" +
							"bullet validate <bullet_file>",
						Action: Validate(),
					},
					{
						Name: "check",
						Usage: "check ssh connection bullet file" +
							"bullet check <bullet_file>",
						Action: Check(),
					},
					{
						Name: "encrypt",
						Usage: "encrypt password in bullet file" +
							"bullet check <bullet_file>",
						Action: Encrypt(),
					},
				},
			},
		},
	}
}
