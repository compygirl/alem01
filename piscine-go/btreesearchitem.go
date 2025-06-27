package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	// if root.Data == elem {
	// 	return root
	// }
	// // fmt.Printf("%v %v %v", root.Data, root.Left.Data, root.Right.Data)
	// if root.Data > elem && root.Left != nil {
	// 	return BTreeSearchItem(root.Left, elem)
	// } else if root.Data < elem && root.Right != nil {
	// 	return BTreeSearchItem(root.Right, elem)
	// }
	// return nil

	trvel := root
	// fmt.Printf("%v %v\n", trvel, root)
	for trvel != nil {
		if trvel.Data == elem {
			return trvel
		} else if trvel.Data < elem {
			trvel = trvel.Right
		} else {
			trvel = trvel.Left
		}
	}
	return nil
}
