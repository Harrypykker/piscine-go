package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := min - 1; i < max-1; i++ {
		result = append(result, i+1)
	}
	return result
}
