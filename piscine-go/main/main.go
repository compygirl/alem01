package main

import (
	"fmt"
	"piscine"
)

// type node struct {
// 	data int
// 	next *node
// }

// func insert(head *node, data int) *node {
// 	n := &node{data: data}

// 	if head == nil {
// 		return n
// 	} else {
// 		n.next = head
// 		return n
// 	}
// }

// func PrintList(head *node) {
// 	for head != nil {
// 		fmt.Print(head.data, " -> ")
// 		head = head.next
// 	}
// 	fmt.Println(nil)
// }

// func main() {
// 	var link *node
// 	link = insert(link, 1)
// 	link = insert(link, 2)
// 	link = insert(link, 3)
// 	link = insert(link, 4)

// 	PrintList(link)
// }

//=======================================================
// REVERSE
//=======================================================
// func main() {
// 	link := &piscine.List{}

// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, 2)
// 	piscine.ListPushBack(link, 3)
// 	piscine.ListPushBack(link, 4)

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)

// 	piscine.ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }
//=======================================================
// BST Lectures
//=======================================================

// type node struct {
// 	left, right *node
// 	data        int
// }

// func insert(root *node, elem int) *node {
// 	if root == nil {
// 		return &node{data: elem}
// 	}

// 	if elem < root.data {
// 		root.left = insert(root.left, elem)
// 	} else {
// 		root.right = insert(root.right, elem)
// 	}
// 	return root
// }

// func PrintTree(root *node) {
// 	if root == nil {
// 		return
// 	}
// 	fmt.Println(root.data)
// 	PrintTree(root.left)
// 	PrintTree(root.right)
// }

// func main() {
// 	var root *node
// 	root = insert(root, 4)
// 	root = insert(root, 2)
// 	root = insert(root, 6)
// 	root = insert(root, 3)
// 	root = insert(root, 5)
// 	root = insert(root, 7)
// 	root = insert(root, 1)
// 	PrintTree(root)
// }
// func main() {
// 	root := &piscine.TreeNode{Data: "4"}
// 	piscine.BTreeInsertData(root, "1")
// 	piscine.BTreeInsertData(root, "7")
// 	piscine.BTreeInsertData(root, "5")
// 	selected := piscine.BTreeSearchItem(root, "7")
// 	fmt.Print("Item selected -> ")
// 	if selected != nil {
// 		fmt.Println(selected.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Parent of selected item -> ")
// 	if selected.Parent != nil {
// 		fmt.Println(selected.Parent.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Left child of selected item -> ")
// 	if selected.Left != nil {
// 		fmt.Println(selected.Left.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Right child of selected item -> ")
// 	if selected.Right != nil {
// 		fmt.Println(selected.Right.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}
// }
func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	node := piscine.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
}
