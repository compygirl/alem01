package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var start1 bool = false
	// var start2 bool = false
	word1 := []rune{}
	var word2 []rune
	var insert bool = false
	var sort bool = false
	if len(args) == 0 {
		printHelp()
	}

	for _, command := range args {

		runesComm := []rune(command)

		if command == "--help" || command == "-h" || len(args) == 0 {
			printHelp()
		} else if "-i" == string(runesComm[0:2]) || "--inser" == string(runesComm[0:7]) {
			insert = true
			for _, symb := range runesComm {
				if symb == '=' {
					start1 = true
				} else if start1 {
					word1 = append(word1, symb)
				}
			}
		} else if "-o" == string(runesComm[0:2]) || "--order" == string(runesComm[0:7]) {
			sort = true
		} else {
			temp := []rune(command)
			word2 = append(word2, temp...)
			if insert {
				word2 = append(word2, word1...)
			}

		}
	}
	if sort {
		sortRunes2(word2)
	}
	printResult(word2)
}

func printHelp() {
	var text string = "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
	runes := []rune(text)
	for _, letter := range runes {
		z01.PrintRune(letter)
	}
}

func sortRunes2(runes []rune) {
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				SwapRune2(&runes[j], &runes[j+1])
			}
		}
	}
}

func printResult(runes []rune) {
	for _, letter := range runes {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func SwapRune2(a *rune, b *rune) {
	temp := *a
	*a = *b
	*b = temp
}
