package piscine

func AlphaCount(s string) int {
	sentence := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if (sentence[i] >= 'a' && sentence[i] <= 'z') || (sentence[i] >= 'A' && sentence[i] <= 'Z') {
			counter++
		}
	}
	// for _, letter := range sentence {
	// 	if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
	// 		counter++
	// 	}
	// }
	return counter
}
