package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	num := 0
	var upper bool = false

	for _, name := range arguments {
		if name == "--upper" {
			upper = true
		} else {
			names := []rune(name)
			num = convertStringToInt(names)
			if num > 26 {
				z01.PrintRune(' ')
			} else {
				if upper {
					z01.PrintRune(rune(64 + num))
				} else {
					z01.PrintRune(rune(96 + num))
				}
			}
		}
	}
	if len(arguments) > 0 {
		z01.PrintRune('\n')
	}
}

func convertStringToInt(runes []rune) int {
	res := 0
	for _, digit := range runes {
		res = res*10 + int(digit-'0')
	}
	return res
}
