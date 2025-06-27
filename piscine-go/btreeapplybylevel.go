package piscine

func applyByLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	}
	if level > 0 {
		applyByLevel(root.Left, level-1, f)
		applyByLevel(root.Right, level-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)

	for i := 0; i < h; i++ {
		applyByLevel(root, i, f)
	}
}

// func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	var arr []*TreeNode
// 	f(root.Data)

// 	for
// 	if root.Left != nil {
// 		arr = append(arr, root.Left)
// 	}
// 	if root.Right != nil {
// 		arr = append(arr, root.Right)
// 	}
// }
