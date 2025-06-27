package piscine

func Capitalize(s string) string {
	w := []rune(s)
	for i := 0; i < len(w); i++ {
		if (i != 0 && !(s[i-1] >= 'a' && s[i-1] <= 'z')) && (!(s[i-1] >= 'A' && s[i-1] <= 'Z')) && (!(s[i-1] >= '0' && s[i-1] <= '9')) {
			if s[i] >= 'a' && s[i] <= 'z' {
				w[i] -= 32
			}
		} else if i != 0 {
			if s[i] >= 'A' && s[i] <= 'Z' {
				w[i] += 32
			}
		} else {
			if s[i] >= 'a' && s[i] <= 'z' {
				w[i] -= 32
			}
		}
	}
	return string(w)
}
