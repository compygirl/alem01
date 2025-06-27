package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func bubbleSort(l *NodeI) {
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

func ListSort(l *NodeI) *NodeI {
	bubbleSort(l)
	return l
}
