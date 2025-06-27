package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	// tempList := &List{}
	// iter := l.Head
	// for iter != nil {
	// 	if iter.Data != data_ref {
	// 		ListPushBack(tempList, iter.Data)
	// 	}
	// 	iter = iter.Next
	// }
	// l.Head = tempList.Head
	// l.Tail = tempList.Tail

	tempHead := l.Head
	l.Head, l.Tail = nil, nil

	iter := tempHead
	last := tempHead
	for iter != nil {
		if iter.Data != data_ref {
			if l.Head == nil {
				l.Head, l.Tail = iter, iter
			} else {
				l.Tail.Next = iter
				l.Tail = l.Tail.Next
				last = iter
			}
		}
		iter = iter.Next
	}
	last.Next = nil
}
