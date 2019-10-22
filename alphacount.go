package piscine

func AlphaCount(str string) int {
	i := 0
	for letter := range str {
		if ([]rune(letter) > "a" && []rune(letter) < "z") || ([]rune(letter) > "A" && []rune(letter) < "Z") {
			i++
		}
	}
	return i
}
