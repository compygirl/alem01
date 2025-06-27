package piscine

// func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

// }

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		} else {
			successor := inorder(root)
			root.Data = successor.Data
			root.Right = BTreeDeleteNode(root.Right, successor)
		}
	}
	return root
}

func inorder(node *TreeNode) *TreeNode {
	node = node.Right
	for node != nil && node.Left != nil {
		node = node.Left
	}
	return node
}
