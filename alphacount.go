package piscine

func AlphaCount(str string) int {
	i := 0
	for letter := range str {
		i++
		fmt.Printf("%c", letter)
	}
	return i
}
