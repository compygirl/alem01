package piscine

func sortRunes2(words []string) {
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if words[j] > words[j+1] {
				SwapString(&words[j], &words[j+1])
			}
		}
	}
}

func SwapString(w1 *string, w2 *string) {
	temp := *w1
	*w1 = *w2
	*w2 = temp
}

func SortWordArr(a []string) {
	sortRunes2(a)
}
