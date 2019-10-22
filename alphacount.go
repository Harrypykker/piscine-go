package piscine

func AlphaCount(str string) int {
	i := 0
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			i++
		}
	}
	return i
}
