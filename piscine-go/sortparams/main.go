package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	sortRunes(arguments)
	for _, name := range arguments {
		for _, letter := range name {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

func sortRunes(words []string) {
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if words[j] > words[j+1] {
				SwapString(&words[j], &words[j+1])
			}
		}
	}
}

func SwapString(w1 *string, w2 *string) {
	temp := *w1
	*w1 = *w2
	*w2 = temp
}
