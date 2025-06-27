package piscine

import (
	"github.com/01-edu/z01"
)

// func main() {
// 	// PrintNbr(-7627169662776005220)
// 	// PrintNbr(500)
// 	// PrintNbr(5040)
// 	// PrintNbr(-1937015452697115600)
// 	// PrintNbr(1059907761690928898)
// 	// PrintNbr(1042387699194471602)
// 	// PrintNbr(-1)
// 	// PrintNbr(-9223372036854775808)
// 	z01.PrintRune('\n')
// }
func PrintNbr(n int) {
	var neg bool = false
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			n = n + 1
			neg = true
		}
		n = (-1) * n
	}

	if n == 0 {
		z01.PrintRune('0')
	} else {
		temp := n
		div := 1
		for temp >= 10 {
			temp = temp / 10
			div = div * 10
		}

		temp = n
		rem := n
		for rem > 0 {
			if div == 1 && neg {
				rem = rem + 1
			}

			curr := rem/div + '0'

			rem = rem % div
			z01.PrintRune(rune(curr))
			if div > 0 {
				for rem == 0 && div > 1 {
					z01.PrintRune('0')
					div = div / 10
				}
			}
			div = div / 10

		}
	}
}
