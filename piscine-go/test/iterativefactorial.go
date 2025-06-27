package main

import "fmt"

func main() {
	fmt.Println(IterativeFactorial(0))
}

func IterativeFactorial(nb int) int {
	res := 1
	if nb >= 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			res *= i
		}
	} else {
		res = 0
	}
	return res
}
