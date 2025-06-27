package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	chars := []rune(s)
	res := 0
	firstIndex := 0

	if len(chars) > 0 {
		var neg bool = false
		if chars[0] == '-' {
			firstIndex = 1
			neg = true
		} else if chars[0] == '+' {
			firstIndex = 1
		}
		for i := firstIndex; i < len(chars); i++ {
			if chars[i] >= '0' && chars[i] <= '9' {
				res = res*10 + int(chars[i]-'0')
			} else {
				return 0
			}
		}
		if neg {
			res = res * (-1)
		}
	}

	return res
}

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
	z01.PrintRune('\n')
}

func IsNumeric(s string) bool {
	w := []rune(s)
	start := 0
	if w[0] == '-' {
		start = 1
	}
	for i := start; i < len(s); i++ {
		if !(w[i] >= '0' && w[i] <= '9') {
			return false
		}
	}

	return true
}

func printError(warn string) {
	for _, letter := range warn {
		z01.PrintRune(letter)
	}
}

func checkOverflow(num1, num2 int) bool {
	if num1 > 0 && num2 > 0 && num1+num2 < 0 {
		return true
	} else if num1 < 0 && num2 < 0 && num1-num2 > 0 {
		return true
	} else if num1*num2/num1 != num2 {
		return true
	}
	return false
}

func applyOperator(n1, n2 int, op string) int {
	if op == "+" {
		return n1 + n2
	} else if op == "-" {
		return n1 - n2
	} else if op == "/" {
		if n2 == 0 {
			printError("No division by ")
			return 0
		}
		return n1 / n2
	} else if op == "*" {
		return n1 * n2
	} else if op == "%" {
		if n2 == 0 {
			printError("No modulo by ")
			return 0
		}
		return n1 % n2
	} else {
		return 0
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		if IsNumeric(args[0]) && IsNumeric(args[2]) {
			if args[1] == "+" || args[1] == "-" || args[1] == "/" || args[1] == "*" || args[1] == "%" {
				num1 := Atoi(args[0])
				num2 := Atoi(args[2])
				if checkOverflow(num1, num2) {
					return
				}
				PrintNbr(applyOperator(num1, num2, args[1]))
			} else {
				return
			}
		} else {
			return
		}
	}
}
