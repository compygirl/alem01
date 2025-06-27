package piscine

func StringToIntSlice(str string) []int {
	arr := []rune(str)
	if len(str) == 0 {
		return []int(nil)
	}
	nums := make([]int, len(arr))
	for ind, val := range arr {
		nums[ind] = int(val)
	}

	return nums

	// return []int([]rune(s))
}
