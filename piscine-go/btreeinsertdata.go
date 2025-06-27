package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	tn := &TreeNode{Data: data}
	if root == nil {
		return tn
	} else {
		if root.Data < data {
			//
			root.Right = BTreeInsertData(root.Right, data)
			root.Right.Parent = root
		} else {
			// root.Left.Parent = root
			root.Left = BTreeInsertData(root.Left, data)
			root.Left.Parent = root

		}
	}
	return root
}
