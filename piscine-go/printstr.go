package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	chars := []rune(s)
	for _, ch := range chars {
		z01.PrintRune(ch)
	}
}
