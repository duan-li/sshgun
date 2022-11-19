package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		data := []byte(template)
		err1 := os.WriteFile(dir+"/"+file, data, 0644)
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println("new template created: ", cCtx.Args().First())
		return nil
	}
}

func Validate() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("validate template: ", cCtx.Args().First())
		return nil
	}
}
