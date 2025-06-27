package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a
	*a = temp / *b
	*b = temp % *b
}
