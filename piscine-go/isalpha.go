package piscine

func IsAlpha(s string) bool {
	w := []rune(s)
	// check := true
	for _, letter := range w {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || (letter >= '0' && letter <= '9') {
		} else {
			return false
		}
	}
	return true
}
