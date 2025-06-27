package piscine

func AppendRange(min, max int) []int {
	var answer []int
	for i := 0; i < max-min; i++ {
		answer = append(answer, i+min)
	}
	return answer
}
