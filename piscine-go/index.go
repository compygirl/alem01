package piscine

func Index(s string, toFind string) int {
	word := []rune(s)
	substr := []rune(toFind)
	for i := 0; i < len(word)-len(substr); i++ {
		if s[i:len(substr)+i] == toFind {
			return i
		}
	}
	return -1
}
