package piscine

func JumpOver(str string) string {
	res := ""
	if len(str) < 3 {
		return "\n"
	} else {
		for i := 0; i < len(str); i++ {
			if i%3 == 2 {
				res += string(str[i])
			}
		}
		return res + "\n"
	}
}
