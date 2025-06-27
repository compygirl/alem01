package piscine

func ListSize(l *List) int {
	size := 0
	if l == nil {
		return 0
	}
	iter := l.Head
	for iter != nil {
		size++
		iter = iter.Next
	}
	return size
}
