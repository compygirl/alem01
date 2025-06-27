package piscine

func MakeRange(min, max int) []int {
	// var answer := make ([]int, 0)
	if min < max {
		size := max - min
		answer := make([]int, size)
		for i := 0; i < size; i++ {
			answer[i] = i + min
		}
		return answer
	}
	var ans []int
	return ans
}
