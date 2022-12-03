package hue

const (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m" // error
	Green   = "\033[1;32m%s\033[0m" // success
	Yellow  = "\033[1;33m%s\033[0m" // warn
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m" // fatal
	Teal    = "\033[1;36m%s\033[0m" // info
	White   = "\033[1;37m%s\033[0m"
)
