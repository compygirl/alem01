package piscine

func ListReverse(l *List) {
	// tempList := &List{}
	// iter := l.Head

	// for iter != nil {
	// 	ListPushFront(tempList, iter.Data)
	// 	iter = iter.Next
	// }
	// l.Head = tempList.Head
	// l.Tail = tempList.Tail

	curr := l.Head
	var next *NodeL
	var prev *NodeL
	l.Tail = l.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
