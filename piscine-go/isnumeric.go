package piscine

func IsNumeric(s string) bool {
	w := []rune(s)
	for _, letter := range w {
		if !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
