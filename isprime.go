package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	check := 0
	for i := 1; i < nb; i++ {
		if nb%i == 0 {
			check = check + 1
			if check >= 2 {
				return false
			}
		}

	}
	return true
}
