package piscine

func FirstRune(s string) rune {
	runes := []rune(s)
	safeSubstring := string(runes[0:1])
	fmt.Println(safeSubstring)
}
