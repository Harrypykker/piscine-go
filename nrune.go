package piscine

func NRune(s string, n int) rune {
	first := []rune(s)
	if StrLen2(first) < n || n < 1 {
		return '\x00'
	} else {
		return first[n-1]
	}
}
