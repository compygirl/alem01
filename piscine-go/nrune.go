package piscine

func NRune(s string, n int) rune {
	if len(s) < n || n < 1 {
		return 0
	} else {
		return []rune(s)[n-1]
	}
}
