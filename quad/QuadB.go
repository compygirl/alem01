package main

import "github.com/01-edu/z01"

func main() {
	QuadB(4, 1)
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for r := 0; r < y; r++ {
			for c := 0; c < x; c++ {
				if (r == 0 && c == 0) || (y > 1 && x > 1 && r == y-1 && c == x-1) {
					z01.PrintRune('/')
				} else if (r == 0 && c == x-1) || (r == y-1 && c == 0) {
					z01.PrintRune('\\')
				} else if (r == 0 || r == y-1) || (r != 0 && c == 0) || (r != 0 && c == x-1) {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
