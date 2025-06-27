package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	flag1 := true
	flag2 := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			flag1 = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			flag2 = false
		}
	}

	return flag1 || flag2
}
