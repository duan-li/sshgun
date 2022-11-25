package main

import (
	"github.com/duan-li/sshgun/cmd"
	"log"
	"os"
)

const name = "sshgun"
const usage = "a tool for managing ssh connections"

func main() {
	app := cmd.Init()
	app.Name = name
	app.Usage = usage

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
