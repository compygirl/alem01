package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		filename := os.Args[1]
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("the mistake is : %v\n", err.Error())
		}

		fmt.Printf("%v", string(file))
	}
}
