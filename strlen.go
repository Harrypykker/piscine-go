package piscine

import "github.com/01-edu/z01"

func StrLen(str string) int {
	for _, r := range str {
		r = r + 1
	}
	z01.PrintRune(r)
}
