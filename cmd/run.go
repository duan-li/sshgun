package cmd

import (
	"fmt"
	"github.com/duan-li/hue"
	"github.com/duan-li/sshgun/bullet"
	"github.com/duan-li/sshgun/ssh"
	"github.com/urfave/cli/v2"
	"log"
)

func Run() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if cCtx.NArg() < 1 {
			return fmt.Errorf("Missing file name.")
		}
		file := cCtx.Args().First()

		content, err := bullet.ReadBullet(file)

		if err != nil {
			log.Fatal(err)
		}

		data, err1 := bullet.Paser(content)

		if err1 != nil {
			log.Fatal(err1)
		}
		hue.Blueln("d")
		for _, server := range data["bullet"].Servers {
			config := ssh.NewConfig(server.Ip, server.Port, server.Username, server.Password, server.Sudopassword)

			hue.Info("=======================================")
			hue.Info(fmt.Sprintf("Server: %s", server.Ip))
			hue.Info(fmt.Sprintf("Username: %s", server.Username))
			hue.Info("---------------------------------------")
			for _, command := range data["bullet"].Commands {
				hue.Success(fmt.Sprintf(">>>Command: %s", command), false)
				out, err := ssh.Runner(config, command)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(out)
			}
		}

		result := hue.Sbluef(fmt.Sprintf("All done: %s", file))
		return cli.Exit(result, 0)
	}
}
