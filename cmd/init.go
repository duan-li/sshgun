package cmd

import (
	"github.com/urfave/cli/v2"
)

func Init() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"exec"},
				Usage:   "run a command on a ssh connection",
				Action:  Run(),
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
				},
			},
		},
	}
}
