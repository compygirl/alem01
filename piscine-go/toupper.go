package piscine

func ToUpper(s string) string {
	w := []rune(s)
	for index, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			letter -= 32
			w[index] = letter
		}
	}
	return string(w)
}
