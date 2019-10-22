package piscine

import "fmt"

func AlphaCount(str string) int {
	for index, letter := range str {
		fmt.PrintRune("letter: %c\n")
	}
}
