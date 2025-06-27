package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func convertToString(x int) string {
	res := ""
	for x > 0 {
		res = string(rune(x%10+'0')) + res
		x = x / 10
	}
	return res
}

func PrintResult(x, y int, quadName string) {
	str := "[" + quadName + "] [" + convertToString(x) + "] [" + convertToString(y) + "]"
	for _, let := range str {
		z01.PrintRune(let)
	}
}

func PrintString(word string) {
	for _, let := range word {
		z01.PrintRune(let)
	}
	// z01.PrintRune('\n')
}

func main() {
	var arr []rune
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, char)
	}
	// str := string(arr)

	x := 0
	y := 0
	for _, char := range arr {
		if char != '\n' && y == 0 {
			x++
		}
		if char == '\n' {
			y++
		}
	}

	if x == 0 || y == 0 {
		PrintString("Not a quad function\n")
		return
	}
	if isEqual(arr, x, y, 'o', 'o', 'o', 'o', '-', '|') {
		PrintResult(x, y, "quadA")
		z01.PrintRune('\n')
		return
	}

	if isEqual(arr, x, y, '/', '\\', '\\', '/', '*', '*') {
		PrintResult(x, y, "quadB")
		z01.PrintRune('\n')
		return
	}

	n := 0
	if isEqual(arr, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		n++
		PrintResult(x, y, "quadC")
	}
	if isEqual(arr, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			PrintString(" || ")
		}
		n++
		PrintResult(x, y, "quadD")
	}
	if isEqual(arr, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			PrintString(" || ")
		}
		n++
		PrintResult(x, y, "quadE")
	}

	if n > 0 {
		z01.PrintRune('\n')
		return
	}

	PrintString("Not a quad function\n")
}

func isEqual(arr []rune, x, y int, c1, c2, c3, c4, hor, ver rune) bool {
	var arrE []rune
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 { // first row
				if j == 0 {
					arrE = append(arrE, c1) // top-left corner
				} else if j == x-1 {
					arrE = append(arrE, c2) // top-right corner
				} else {
					arrE = append(arrE, hor) // top row
				}
			} else if i == y-1 { // bottom row
				if j == 0 {
					arrE = append(arrE, c3) // bottom-left corner
				} else if j == x-1 {
					arrE = append(arrE, c4) // bottom -right corner
				} else {
					arrE = append(arrE, hor) // bottom row
				}
			} else {
				if j == 0 || j == x-1 { // left and right walls
					arrE = append(arrE, ver) // first and last columns
				} else {
					arrE = append(arrE, ' ') // all in the middle
				}
			}
		}
		arrE = append(arrE, '\n')
	}
	// strE := string(arrE)
	if len(arr) != len(arrE) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arrE[i] {
			return false
		}
	}
	// if strE == str {
	// 	return true
	// }
	return true
}
