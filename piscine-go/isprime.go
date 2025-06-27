package piscine

func IsPrime(nb int) bool {
	// if nb <= 1 {
	// 	return false
	// }
	// for i := 2; i < nb; i++ {
	// 	if nb%i == 0 {
	// 		return false
	// 	}
	// }
	// return true
	// Corner cases
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}

	// This is checked so that we can skip
	// middle five numbers in below loop
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
