package piscine

func AlphaCount(s string) int {
	i := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			i++
		}
	}
	return i
}
