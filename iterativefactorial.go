package piscine

func iterativepower(nb int, power int) int {
	if nb >= 0 {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	} else {
		return 0
	}
}
