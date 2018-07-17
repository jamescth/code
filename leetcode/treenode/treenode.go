package treenode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CompareTreeNode(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val != r2.Val {
		return false
	}

	if !CompareTreeNode(r1.Left, r2.Left) {
		return false
	}
	if !CompareTreeNode(r1.Right, r2.Right) {
		return false
	}
	return true
}

func Insert(root **TreeNode, n int) {
	if *root == nil {
		*root = &TreeNode{Val: n}
		return
	}

	if n < (*root).Val {
		if (*root).Left == nil {
			(*root).Left = &TreeNode{Val: n}
			return
		}
		Insert(&(*root).Left, n)
		return
	}

	if (*root).Right == nil {
		(*root).Right = &TreeNode{Val: n}
		return
	}
	Insert(&(*root).Right, n)
}

func Show(root *TreeNode, level int) {
	if root == nil {
		return
	}
	fmt.Printf("-> val %d l %d\n", root.Val, level)
	Show(root.Left, level+1)
	Show(root.Right, level+1)
}
