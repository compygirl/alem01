package piscine

func Map(f func(int) bool, a []int) []bool {
	bools := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		bools[i] = f(a[i])
	}
	return bools
}
