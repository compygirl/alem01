package piscine

func getSize2(node *NodeI) int {
	if node == nil {
		return 0
	}
	counter := 0
	for node != nil {
		node = node.Next
		counter++
	}
	return counter
}

func selectionSort(l *NodeI) {
	iter1 := l
	for iter1 != nil {
		iter2 := iter1.Next
		for iter2 != nil {
			if iter1.Data > iter2.Data {
				temp := iter2.Data
				iter2.Data = iter1.Data
				iter1.Data = temp
			}
			iter2 = iter2.Next
		}
		iter1 = iter1.Next
	}
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		selectionSort(n2)
		return n2
	} else if n2 == nil {
		selectionSort(n1)
		return n1
	}
	iter := n1
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = n2
	selectionSort(n1)
	return n1
}

// func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
// if n1 == nil {
// 	return n2
// } else if n2 == nil {
// 	return n1
// }
// return nil
// var prehead *NodeI
// var curr *NodeI

// if n1.Data <= n2.Data {
// 	prehead = n1
// 	curr = n1
// } else {
// 	prehead = n2
// 	curr = n2
// }

// for n1 != nil && n2 != nil {
// 	if n1.Data <= n2.Data {
// 		curr.Next = n1
// 	}
// }

// My solution
// iter1 := n1
// iter2 := n2
// size := getSize2(n1) + getSize2(n2)

// var head *NodeI
// var headIter *NodeI
// if n1 == nil {
// 	selectionSort(n2)
// 	return n2
// } else if n2 == nil {
// 	selectionSort(n1)
// 	return n1
// } else {
// 	selectionSort(n1)
// 	selectionSort(n2)
// }

// for size > 0 && iter1 != nil && iter2 != nil {
// 	if iter1.Data < iter2.Data {
// 		if head == nil {
// 			head = iter1
// 			headIter = head
// 		} else {
// 			headIter.Next = iter1
// 			headIter = headIter.Next
// 		}
// 		iter1 = iter1.Next
// 	} else if iter1.Data > iter2.Data {
// 		if head == nil {
// 			head = iter2
// 			headIter = head
// 		} else {
// 			headIter.Next = iter2
// 			headIter = headIter.Next
// 		}
// 		iter2 = iter2.Next
// 	} else {
// 		if head == nil {
// 			head = iter1
// 			head.Next = iter2
// 			headIter = head.Next
// 		} else {
// 			headIter.Next = iter1
// 			headIter = headIter.Next
// 			headIter.Next = iter2
// 			headIter = headIter.Next
// 		}
// 		iter1 = iter1.Next
// 		iter2 = iter2.Next
// 	}
// 	size--
// }
// if size > 0 && iter1 != nil {
// 	headIter.Next = iter1
// } else {
// 	headIter.Next = iter2
// }
// selectionSort(head)
// // for head.Data < 0 {
// // 	head = head.Next
// // }
// return head
// }
