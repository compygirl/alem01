package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		l = &NodeI{Data: data_ref}
	}

	iter := l
	var prev *NodeI
	var next *NodeI
	temp := &NodeI{Data: data_ref}
	for iter != nil {
		if iter.Data >= data_ref {
			if prev != nil {
				next = prev.Next
				prev.Next = temp
				temp.Next = next
			} else {
				temp.Next = iter
				l = temp
			}

			break
		}
		prev = iter
		iter = iter.Next
	}
	if iter == nil {
		prev.Next = temp
	}
	return l
}
