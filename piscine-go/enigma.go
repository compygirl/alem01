package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// ****c = a
	// d = ***c
	// b = ***d
	// **a = b
	y := *******c
	*******c = ***a
	t := ****d
	****d = y
	g := *b
	*b = t
	***a = g
}
