package piscine

func Unmatch(a []int) int {
	counter := 0
	for i := 0; i < len(a); i++ {
		counter = 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				counter++
			}
		}
		if counter%2 != 0 {
			return a[i]
		}
	}
	return -1
}
