package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func runToInt(chislo int) []rune {
	temp := chislo
	div := 1
	for temp >= 10 {
		temp = temp / 10
		div = div * 10
	}
	temp = chislo
	rem := chislo
	var massivRun []rune

	for rem > 0 {
		massivRun = append(massivRun, rune(rem/div+'0'))
		rem = rem % div
		div /= 10
	}
	return massivRun
}

func main() {
	points := &point{}

	setPoint(points)
	// fmt.Println(string(runToInt(points.x)))
	res := "x = " + string(runToInt(points.x)) + ", y = " + string(runToInt(points.y)) + "\n"
	for _, let := range []rune(res) {
		z01.PrintRune(let)
	}
}
