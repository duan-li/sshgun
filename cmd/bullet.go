package cmd

import (
	"fmt"
	bullet "github.com/duan-li/sshgun/bullet"
	ssh "github.com/duan-li/sshgun/ssh"
	"github.com/urfave/cli/v2"
	sshClient "golang.org/x/crypto/ssh"
	"log"
)

const template = `bullet:
  servers:
      - ip: 127.0.0.1
        port: 22
        username: root
        password: root
        sudo: true
        sudopassword: root
      - ip: 127.0.0.2
        port: 22
        username: root
        password: root
        sudo: false
  commands:
      - echo "Hello World"
      - ls -l ~
`

func InitBullet() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if cCtx.NArg() < 1 {
			return fmt.Errorf("Missing template name.")
		}
		file := cCtx.Args().First()

		_, err := bullet.WriteBullet(file, []byte(template))

		if err != nil {
			log.Fatal(err)
		}
		return cli.Exit("New template created: "+file, 0)
	}
}

func Validate() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if cCtx.NArg() < 1 {
			return fmt.Errorf("Missing file name.")
		}
		file := cCtx.Args().First()

		content, err := bullet.ReadBullet(file)

		if err != nil {
			log.Fatal(err)
		}

		_, err1 := bullet.Paser(content)

		if err1 != nil {
			log.Fatal(err1)
		}

		return cli.Exit("Template is valid: "+file, 0)
	}
}

func Check() func(cCtx *cli.Context) error {
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
			sshConfig := ssh.NewSSHConfig(config)
			_, err := sshClient.Dial("tcp", ssh.HostAddress(config), sshConfig)
			if err != nil {
				log.Fatal(err)
			}
		}

		return cli.Exit("Connection is valid: "+file, 0)
	}
}
