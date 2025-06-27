package piscine

func Rot14(s string) string {
	ans := []rune(s)
	// lastLowerLetter := 'z'
	// lastCapLetter := 'Z'

	for i := 0; i < len(ans); i++ {
		if ans[i] >= 'a' && ans[i] <= 'z' {
			if ans[i]+14 > 'z' {
				ans[i] = (ans[i] + 13 - 'z') + 'a'
			} else {
				ans[i] = ans[i] + 14
			}
		} else if ans[i] >= 'A' && ans[i] <= 'Z' {
			if ans[i]+14 > 'Z' {
				ans[i] = (ans[i] + 13 - 'Z') + 'A'
			} else {
				ans[i] = ans[i] + 14
			}
		}
	}
	return string(ans)
}
