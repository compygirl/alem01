package piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	menu := food{15, 10, 12}
	if order == "burger" {
		return menu.burger
	} else if order == "chips" {
		return menu.chips
	} else if order == "nuggets" {
		return menu.nuggets
	} else {
		return 404
	}
}
