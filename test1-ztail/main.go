package main
import (
	"fmt"
	"os"
)
func main() {
	args := os.Args
	result := 0
	index := args[2]
	for i := 0; i < len(index); i++ {
		result *= 10
		result += int(index[i] - 48)
	}
	if len(args) == 4 {
		file, _ := os.ReadFile(args[3])
		fmt.Printf(string(file[len(file)-result:]))
	} else {
		for i := 3; i < len(args); i++ {
			file, err := os.ReadFile(args[i])
			if err != nil {
				// fmt.Printf("")
				fmt.Printf("%s\n", err)
			} else {
				if i != 3 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %v <==\n", args[i])
				// fmt.Printf("len: %v res: %v\n", len(file), result)
				if len(string(file)) <= result {
					fmt.Printf("%v", string(file))
					// fmt.Println("===================")
					// fmt.Printf("\n")
					os.Exit(1)
				}
				// fmt.Printf("+++++++++++++")
				fmt.Printf(string(file[len(file)-result:]))
			}
		}
	}
}
// Your program output is not correct :
// $ go run . "-c" "23" "src/student/ztail/main.go" "quest8.txt"
// open src/student/ztail/main.go: no such file or directory
// ==> quest8.txt <==
// frv  Wz0s XMl
// exit status 1$
// Expected :
// $ go run . "-c" "23" "src/student/ztail/main.go" "quest8.txt"
// open src/student/ztail/main.go: no such file or directory
// ==> quest8.txt <==
// frv  Wz0s XMl
// exit status 1$
