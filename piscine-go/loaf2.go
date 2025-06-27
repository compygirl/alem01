package piscine

func LoafOfBread2(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	if str[len(str)-1] == ' ' {
		res := ""
		ans := ""

		space := 0

		for i := len(str) - 1; i >= 0; i-- {
			if str[i] == ' ' {
				space++
			} else {
				break
			}
		}
		for i := 0; i < len(str)-space; i++ {
			if str[i] != ' ' {
				ans += string(str[i])
				if len(ans) == 5 {
					res += ans
					if len(str)-space != i+1 {
						res += " "
					}
					ans = ""
					i++
				}
				if i == len(str)-space-1 && len(ans) < 5 {
					res += ans
					break
				}
			}
		}

		res1 := ""
		if res[len(res)-1] == ' ' {
			for i := 0; i < len(res)-1; i++ {
				res1 += string(res[i])
			}
			return res1 + "\n"
		}

		return res + "\n"
	}

	res := ""
	ans := ""

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			ans += string(str[i])
			if len(ans) == 5 {
				res += ans
				if len(str) != i+1 {
					res += " "
				}
				ans = ""
				i++
			}
			if i == len(str)-1 && len(ans) < 5 {
				res += ans
				break
			}
		}
	}

	res1 := ""
	if res[len(res)-1] == ' ' {
		for i := 0; i < len(res)-1; i++ {
			res1 += string(res[i])
		}
		return res1 + "\n"
	}

	return res + "\n"
}
