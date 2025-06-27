package piscine

func TrimAtoi(s string) int {
	temp := 0
	// tens := 1
	var neg bool = false

	for _, letter := range s {
		if temp == 0 && letter == '-' {
			neg = true
		}
		if letter >= '0' && letter <= '9' {
			digit := letter - '0'
			temp = temp*10 + int(digit)
		}

	}

	if neg {
		temp *= (-1)
	}
	return temp
}
