package piscine

import "github.com/01-edu/z01.PrintRune"

func AlphaCount(str string) int {
	for index, letter := range str {
		z01.PrintRune("letter: %c\n")
	}
}
