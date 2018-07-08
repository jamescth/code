package main

import "fmt"

// A TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(root **TreeNode, n int) {
	// fmt.Printf("insert %d\n", n)
	if *root == nil {
		*root = &TreeNode{Val: n}
		return
	}

	//fmt.Printf("%d val %d Left %v Right %v\n", n, root.Val, **root.Left, **root.Right)
	if n < (*root).Val {
		if (*root).Left == nil {
			(*root).Left = &TreeNode{Val: n}
			return
		}
		Insert(&(*root).Left, n)
		return
	}

	// Right / n > root.Val
	Insert(&(*root).Right, n)
}

func Show(root *TreeNode, level int) {
	if root == nil {
		return
	}

	fmt.Printf("val %d level %d\n", root.Val, level)

	Show(root.Left, level+1)
	Show(root.Right, level+1)
}
