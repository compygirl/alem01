package piscine

func RockAndRoll(n int) string {
	res := ""
	if n%2 == 0 && n%3 != 0 && n >= 0 {
		res += "rock\n"
	} else if n%3 == 0 && n%2 != 0 && n >= 0 {
		res += "roll\n"
	} else if n%2 == 0 && n%3 == 0 && n >= 0 {
		res += "rock and roll\n"
	} else if n < 0 {
		res += "error: number is negative\n"
	} else {
		res += "error: non divisible\n"
	}
	return res
}
