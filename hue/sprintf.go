package hue

import "fmt"

func Sblackf(s string) string {
	return fmt.Sprintf(Black, s)
}

func Sredf(s string) string {
	return fmt.Sprintf(Red, s)
}

func Sgreenf(s string) string {
	return fmt.Sprintf(Green, s)
}

func Syellowf(s string) string {
	return fmt.Sprintf(Yellow, s)
}

func Spurplef(s string) string {
	return fmt.Sprintf(Purple, s)
}

func Smagentaf(s string) string {
	return fmt.Sprintf(Magenta, s)
}

func Stealf(s string) string {
	return fmt.Sprintf(Teal, s)
}

func Swhitef(s string) string {
	return fmt.Sprintf(White, s)
}
