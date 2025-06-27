package piscine

func BasicJoin(elems []string) string {
	res := ""
	// for _, word := range elems {
	// 	res += word // res = res + word // "" + "Hello!" =  "Hello!"// "Hello!" + " How" // "Hello! How" + " are"//
	// }

	for i := 0; i < len(elems); i++ {
		res += elems[i]
	}
	return res
}
