package piscine

func AlphaCount(str string) int {
	i := 0
	for letter := range str {
		if (letter > "a" && letter < z) || (letter > "A" && letter < "Z") {
			i++
		}
	}
	return i
}
