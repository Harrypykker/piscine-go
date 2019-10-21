package piscine

func IterativeFactorial(nb int) int {
	If nb <= 0 {
		result == 0}
	If nb > 0 {
	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}
