package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}

	num1, err1 := strconv.Atoi(args[1])
	num2, err2 := strconv.Atoi(args[2])

	if err1 == nil && err2 == nil {
		QuadC(num1, num2)
	}
}

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for r := 0; r < y; r++ {
			for c := 0; c < x; c++ {
				if (r == 0 && c == 0) || (r == 0 && c == x-1) {
					z01.PrintRune('A')
				} else if (r == y-1 && c == 0) || (r == y-1 && c == x-1) {
					z01.PrintRune('C')
				} else if (r == 0 || r == y-1) || (r != 0 && c == 0) || (r != 0 && c == x-1) {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
