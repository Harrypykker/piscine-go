package piscine

func StrLen(str string) int {
	for _, r := range str {
		r = r + 1
	}
}
