package piscine

import "os"
import "github.com/01-edu/z01"

func main() {
	count := os.Args[0]
	p := []rune(count)
	for _, l := range p {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
