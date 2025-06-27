package piscine

func CollatzCountdown(start int) int {
	counter := 0
	if start <= 0 {
		return -1
	} else {
		temp := start

		for temp != 1 {
			if temp%2 == 0 {
				temp /= 2
			} else {
				temp = 3*temp + 1
			}
			counter++
		}
	}
	return counter
}
