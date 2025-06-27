package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Parent == nil {
		root = rplc
	} else if node.Parent != nil && node.Parent.Left == node {
		node.Parent.Left = rplc
	} else if node.Parent != nil && node.Parent.Right == node {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
