package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	var prev bool = true
	for n := '9'; n >= '0'; n-- {
		for n1 := '9'; n1 >= '0'; n1-- {
			for n2 := '9'; n2 >= '0'; n2-- {
				for n3 := '9'; n3 >= '0'; n3-- {
					if n == n2 && n1 == n3 {
					} else if (n2 < n) || (n == n2 && n1 > n3) {
						if prev == false {
							z01.PrintRune(',')
							z01.PrintRune(' ')
							prev = true
						}
						z01.PrintRune(n)
						z01.PrintRune(n1)
						z01.PrintRune(' ')
						z01.PrintRune(n2)
						z01.PrintRune(n3)
						prev = false
					}
				}
			}
		}
	}
}
