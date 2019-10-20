
func FindNextPrime(nb int) int {
	check := 0
	for i := 1; i < nb; i++ {
		if nb%i == 0 {
			check = check + 1
			if check >= 2 {
				for j := nb + 1; j < 1000000; j++ {
					check2 := 0
					for k := 1; k <= j; k++ {
						if j%k == 0 {
							check2 = check2 + 1
							if check2 > 2 {
								continue
							}
						}
					}
					return j
				}
			}
		}
	}
	return nb
}
