package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb >= 0 && nb < 25 {
		for i := 1; i <= nb; i++ {
			res *= i
		}
	} else {
		res = 0
	}
	return res
}
