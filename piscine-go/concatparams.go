package piscine

func ConcatParams(args []string) string {
	res := ""
	for _, word := range args {
		if word != args[len(args)-1] {
			res += word + string('\n')
		} else {
			res += word
		}
	}
	return res
}
