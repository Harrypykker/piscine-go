package piscine

func IterativePower(nb int, power int) int {
	if power >= 0 {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	} else {
		return 0
	}
}
