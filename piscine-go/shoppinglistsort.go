package piscine

func SwapString2(w1 *string, w2 *string) {
	temp := *w1
	*w1 = *w2
	*w2 = temp
}

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				SwapString2(&slice[j], &slice[j+1])
			}
		}
	}
	return slice
}
