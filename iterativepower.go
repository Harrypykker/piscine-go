func iterativepower(nb int, power int) int {
	result := 1
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
}
