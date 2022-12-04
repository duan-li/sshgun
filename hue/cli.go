package hue

import "os"

func info(message string) {
	tealln(message)
}

func warn(message string) {
	yellowln(message)
}

func error(message string) {
	redln(message)
}

func fatal(message string) {
	magentaln(message)
	os.Exit(1)
}

func success(message string) {
	greenln(message)
}
