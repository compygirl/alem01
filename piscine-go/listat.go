package piscine

func getSize(l *NodeL) int {
	if l == nil {
		return 0
	}
	size := 0
	for l != nil {
		l = l.Next
		size++
	}
	return size
}

func ListAt(l *NodeL, pos int) *NodeL {
	if pos <= getSize(l)-1 {
		iter := l
		iterInt := 0
		for iter != nil {
			if iterInt == pos {
				return iter
			}
			iter = iter.Next
			iterInt++

		}
	}
	return nil
}
