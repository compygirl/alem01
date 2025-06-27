package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintError(err error) {
	errr := "ERROR: "
	for _, let := range errr {
		z01.PrintRune(let)
	}
	for _, let := range err.Error() {
		z01.PrintRune(let)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			PrintError(err)
		}
		os.Stdin.Close()
		os.Stdout.Close()
	} else {
		files := os.Args[1:]
		for i := range files {
			file, err := ioutil.ReadFile(files[i])
			if err != nil {
				PrintError(err)
				os.Exit(1)
			}

			for _, letter := range file {
				z01.PrintRune(rune(letter))
			}

		}

	}
}
