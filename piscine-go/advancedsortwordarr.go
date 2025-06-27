package piscine

func SwapString3(w1 *string, w2 *string) {
	temp := *w1
	*w1 = *w2
	*w2 = temp
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if f(a[j], a[j+1]) == 1 {
				SwapString3(&a[j], &a[j+1])
			}
		}
	}
}
