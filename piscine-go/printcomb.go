package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var prev bool = true
	for n := '0'; n <= '9'; n++ {
		for n1 := '1'; n1 <= '9'; n1++ {
			for n2 := '2'; n2 <= '9'; n2++ {
				if n < n1 && n1 < n2 {
					if prev == false {
						z01.PrintRune(',')
						z01.PrintRune(' ')
						prev = true
					}
					z01.PrintRune(n)
					z01.PrintRune(n1)
					z01.PrintRune(n2)
					prev = false

				}
			}
		}
	}
	z01.PrintRune('\n')
}
