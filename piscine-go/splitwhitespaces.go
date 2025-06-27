package piscine

func SplitWhiteSpaces(s string) []string {
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
		} else if len(runes) != 0 {
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
