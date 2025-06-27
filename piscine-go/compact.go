package piscine

func Compact(ptr *[]string) int {
	counter := 0
	var arr []string

	for _, word := range *ptr {
		if word != "" {
			counter++
			arr = append(arr, word)
		}
	}
	*ptr = arr

	return counter
}
