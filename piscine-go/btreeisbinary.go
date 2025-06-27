package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Left.Data <= root.Data {
		return BTreeIsBinary(root.Left)
	} else if root.Left != nil && root.Left.Data > root.Data {
		return false
	}

	if root.Right != nil && root.Right.Data >= root.Data {
		return BTreeIsBinary(root.Right)
	} else if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	return true
}
