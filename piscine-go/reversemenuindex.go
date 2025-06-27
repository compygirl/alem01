package piscine

func ReverseMenuIndex(menu []string) []string {
	arr := make([]string, len(menu))

	for i := 0; i < len(arr); i++ {
		arr[i] = menu[len(arr)-i-1]
	}
	return arr
}

// func ReverseMenuIndex(menu []string) []string {
// 	arr := make([]string, len(menu))

// 	for i := 0; i < len(arr) / 2; i++ {
// 		arr[i],arr[len(arr)-i-1] = menu[len(arr)-i-1],arr[i]
// 	}
// 	return arr
// }
