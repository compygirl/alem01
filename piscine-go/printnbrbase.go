package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// check uniqueness
	var neg bool = false
	var special bool = false
	index := 0

	runes := []rune{}
	if nbr == -9223372036854775808 {
		nbr = nbr + 1
		neg = true
		nbr = (-1) * nbr
		special = true
	}
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	for _, symbol := range base {
		if symbol == '-' || symbol == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}

	if nbr < 0 {
		neg = true
		nbr *= -1
	}

	for nbr > 0 {
		if special {
			index = nbr%len(base) + 1
			special = false
		} else {
			index = nbr % len(base)
		}
		runes = append(runes, rune(base[index]))
		// runes = append
		nbr = nbr / len(base)
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if neg {
			z01.PrintRune('-')
			z01.PrintRune(runes[i])
			neg = false
		} else {
			z01.PrintRune(runes[i])
		}
	}
}
