package hue

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
}

func success(message string) {
	greenln(message)
}
