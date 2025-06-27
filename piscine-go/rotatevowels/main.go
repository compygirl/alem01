package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	res := ""
	if len(args) >= 1 {
		for _, word := range args {
			if word != args[len(args)-1] {
				res += word + " "
			} else {
				res += word
			}
		}
		Print(VowelSwitch(res))
	} else {
		z01.PrintRune('\n')
	}
}

func Print(runes []rune) {
	for _, let := range runes {
		z01.PrintRune(let)
	}
	z01.PrintRune('\n')
}

func VowelSwitch(s string) []rune {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	runes := []rune(s)
	var fristIndex bool = false
	var secondtIndex bool = false

	index1 := 0
	index2 := 0

	for i, j := 0, len(runes)-1; i < j; {
		if !fristIndex {
			for _, let := range vowels {
				if runes[i] == let {
					fristIndex = true
					index1 = i
					break
				}
			}
			i++
		}
		if !secondtIndex {
			for _, let := range vowels {
				if runes[j] == let {
					secondtIndex = true
					index2 = j
					break
				}
			}
			j--
		}

		if fristIndex && secondtIndex {
			SwapRune3(&runes[index1], &runes[index2])
			fristIndex = false
			secondtIndex = false
		}

	}
	return runes
}

func SwapRune3(a *rune, b *rune) {
	temp := *a
	*a = *b
	*b = temp
}
