package piscine

func Atoi(s string) int {
	chars := []rune(s)
	res := 0
	firstIndex := 0
	if len(chars) > 0 {
		var neg bool = false
		if chars[0] == '-' {
			firstIndex = 1
			neg = true
		} else if chars[0] == '+' {
			firstIndex = 1
		}
		for i := firstIndex; i < len(chars); i++ {
			if chars[i] >= '0' && chars[i] <= '9' {
				res = res*10 + int(chars[i]-'0')
			} else {
				return 0
			}
		}
		if neg {
			res = res * (-1)
		}
	}

	return res
}
