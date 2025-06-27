package piscine

func DescendAppendRange(max, min int) []int {
	var answer []int
	if max <= min {
		return []int{}
	}
	for i := max; i > min; i-- {
		answer = append(answer, i)
	}
	return answer
}
