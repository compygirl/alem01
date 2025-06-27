package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}
