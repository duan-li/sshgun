package cmd

import (
	"fmt"
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

		for _, server := range data["bullet"].Servers {
			config := ssh.NewConfig(server.Ip, server.Port, server.Username, server.Password, server.Sudopassword)
			fmt.Println("\033[1;34m=======================================\033[0m")
			fmt.Println("\033[1;34mServer: ", server.Ip, "\u001B[0m")
			fmt.Println("\033[1;34mUsername: ", server.Username, "\u001B[0m")
			fmt.Println("\033[1;34m---------------------------------------\033[0m")
			for _, command := range data["bullet"].Commands {
				fmt.Println("\033[1;32m>>> Command: ", command, "\033[0m")
				out, err := ssh.Runner(config, command)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(out)
			}
		}

		result := "\033[1;40mAll done: " + file + "\u001B[0m"

		return cli.Exit(result, 0)
	}
}
