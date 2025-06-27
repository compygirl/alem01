package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	if isEven(len(args)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

// package main

// import "fmt"

// type student struct {
// 	name string
// 	age  int
// }

// func changeName(pointer student, nameChosen string) {
// 	pointer.name = nameChosen
// 	// (*pointer).name = nameChosen
// 	// pointer -> name = nameCnameChosen
// }

// func addElementsOfStudents(massiv []student) {
// 	for i := 0; i < 10; i++ {
// 		massiv[i].name = "name" + string((i - '0')) //??
// 		massiv[i].age = 30 + i
// 	}
// }

// func main() {
// 	chris := student{}
// 	arr_structs := make([]student, 10)
// 	addElementsOfStudents(arr_structs)
// 	for _, el := range arr_structs {
// 		fmt.Println(el)
// 	}
// 	fmt.Println(chris)
// 	chris.name = "chris"
// 	chris.age = 30
// 	fmt.Println(chris)
// 	changeName(chris, "Lee")
// 	fmt.Println(chris)
// }
