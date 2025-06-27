package main

import "github.com/01-edu/z01"

func main() {
	for l := '0'; l <= '9'; l++ {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
