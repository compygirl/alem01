package piscine

func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr[2]
}
