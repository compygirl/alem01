package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb; i++ {
		if RecursivePower(i, 2) == nb {
			return i
		}
	}
	return 0
}
