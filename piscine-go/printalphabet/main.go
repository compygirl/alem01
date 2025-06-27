package main

import "github.com/01-edu/z01"

func main() {
	for l := 'a'; l <= 'z'; l++ {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
