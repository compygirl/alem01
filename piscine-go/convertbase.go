package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	dict := map[rune]int{}
	answer := ""
	for i, c := range baseFrom {
		dict[c] = i
	}

	power := 1
	res := 0

	for i := len(nbr) - 1; i >= 0; i-- {
		res += dict[[]rune(nbr)[i]] * power
		power *= len(baseFrom)
	}

	for res != 0 {
		answer = string(baseTo[res%len(baseTo)]) + answer
		res /= len(baseTo)
	}

	return answer
}
