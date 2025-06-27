package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root

	for temp.Left != nil {
		temp = temp.Left
	}
	return temp
}
