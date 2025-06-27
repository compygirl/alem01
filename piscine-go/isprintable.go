package piscine

func IsPrintable(s string) bool {
	w := []rune(s)
	for _, letter := range w {
		if !(letter >= 32 && letter <= 126) {
			return false
		}
	}
	return true
}
