package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	temp := n
	div := 1
	digits := []rune{}

	for temp >= 10 { // 120 > 10  12 > 10 1 >10
		temp /= 10 // 12 1
		div *= 10  // 10 100
	}

	rem := n
	for div > 0 { // 120 > 0 // 20 > 0
		curr := '0' + rem/div               // 120 /100 =1  2
		digits = append(digits, rune(curr)) // 1 2
		rem %= div                          // 20 0
		div /= 10                           // 10
	}
	sortRunes(digits)
	for _, values := range digits {
		z01.PrintRune(values)
	}
}

func sortRunes(runes []rune) {
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				SwapRune(&runes[j], &runes[j+1])
			}
		}
	}
}
