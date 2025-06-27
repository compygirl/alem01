package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		max := a[0]
		for _, val := range a {
			if max < val {
				max = val
			}
		}
		return max
	}
}
