package piscine

func Concat(str1 string, str2 string) string {
	w := []rune(str1)
	w2 := []rune(str2)
	w = append(w, w2...)
	return string(w)
}
