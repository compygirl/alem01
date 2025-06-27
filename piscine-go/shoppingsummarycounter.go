package piscine

func SplitWhiteSpaces2(s string) []string {
	var words []string
	var runes []rune
	orig := []rune(s)
	var firstLetter bool = false

	for i := 0; i < len(orig); i++ {
		if orig[i] != ' ' && !firstLetter {
			firstLetter = true
			runes = append(runes, orig[i])
		} else if orig[i] != ' ' && firstLetter {
			runes = append(runes, orig[i])
		} else { //} if len(runes) != 0 {
			words = append(words, string(runes))
			runes = runes[:0]
			firstLetter = false
		}
		if i == len(orig)-1 {
			words = append(words, string(runes))
		}
	}
	return words
}

func ShoppingSummaryCounter(str string) map[string]int {
	dict := map[string]int{}

	if len(str) == 0 {
		dict = map[string]int{"": 1}
	}

	arr := SplitWhiteSpaces2(str)

	for _, word := range arr {
		count, ok := dict[word]
		if ok {
			count++
			dict[word] = count
		} else {
			count = 1
			dict[word] = count
		}
	}

	return dict
}
