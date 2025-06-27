package piscine

func Split(s, sep string) []string {
	var words []string
	runes := []rune(s)
	index := 0

	for len(runes) > 0 {
		index = Index(string(runes), sep)

		if index == -1 && len(runes) > 0 {
			words = append(words, string(runes))
			break
		} else {
			words = append(words, string(runes[0:index]))
			runes = runes[index+len(sep):]
		}
	}

	return words
}
