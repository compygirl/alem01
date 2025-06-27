package piscine

func FindNextPrime(nb int) int {
	// Base case
	if nb <= 1 {
		return 2
	} else if IsPrime(nb) {
		return nb
	}

	prime := nb
	found := false

	// Loop continuously until isPrime returns
	// true for a number greater than n
	for !found {
		prime++

		if IsPrime(prime) {
			found = true
		}
	}

	return prime
}

// nb++
// if nb <= 1 {
// 	return 2
// } else {
// 	if IsPrime(nb) {
// 		return nb
// 	}
// }
// for i := 2; i < nb; i++ {
// 	if nb%i == 0 {
// 		nb++
// 		i = 2
// 	} else {
// 		continue
// 	}
// }
// return nb

// if nb <= 1 {
// 	return 2
// } else {
// 	if IsPrime(nb) {
// 		return nb
// 	} else {
// 		for i := nb; i < 10000000000; i++ {
// 			if IsPrime(i) {
// 				return i
// 			}
// 		}
// 		return 0
// 	}
// }
// }
