package piscine

func BasicAtoi(s string) int {
	chars := []rune(s)
	res := 0
	for i := 0; i < len(chars); i++ {
		// temp :=
		res = res*10 + int(chars[i]-'0')
	}
	return res
}
