package piscine

func BasicAtoi2(s string) int {
	chars := []rune(s)
	res := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] >= '0' && chars[i] <= '9' {
			res = res*10 + int(chars[i]-'0')
		} else {
			return 0
		}
	}
	return res
}
