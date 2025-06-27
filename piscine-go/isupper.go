package piscine

func IsUpper(s string) bool {
	w := []rune(s)
	for _, letter := range w {
		if !(letter >= 'A' && letter <= 'Z') {
			return false
		}
	}
	return true
}
