package piscine

func FirstRune(s string) rune {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	fmt.Print(first)
	return 0
}
