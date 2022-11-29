package main

import (
	"github.com/duan-li/sshgun/cmd"
	"log"
	"os"
)

const name = "sshgun"
const usage = "a tool allows you to run commands over SSH on many servers at once."

func main() {
	app := cmd.Init()
	app.Name = name
	app.Usage = usage

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
