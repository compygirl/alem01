package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil || root.Right != nil {
		leftLeve := BTreeLevelCount(root.Left)
		rightLevel := BTreeLevelCount(root.Right)
		if leftLeve < rightLevel {
			return 1 + BTreeLevelCount(root.Right)
		} else {
			return 1 + BTreeLevelCount(root.Left)
		}
	} else {
		return 1
	}
}
