package piscine

func ToLower(s string) string {
	w := []rune(s)
	for index, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			letter += 32
			w[index] = letter
		}
	}
	return string(w)
}
