package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, word := range args {
		if word == "01" || word == "galaxy" || word == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
