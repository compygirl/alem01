package piscine

func StrRev(s string) string {
	changableString := []rune(s)
	length := len(changableString)
	for i := 0; i < length/2; i++ {
		SwapRune(&changableString[i], &changableString[length-1-i])
	}
	return string(changableString)
}

func SwapRune(a *rune, b *rune) {
	temp := *a
	*a = *b
	*b = temp
}
