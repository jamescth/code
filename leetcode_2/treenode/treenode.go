package treenode

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
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

	if n > (*root).Val {
		if (*root).Right == nil {
			(*root).Right = &TreeNode{Val: n}
			return
		}
		Insert(&(*root).Right, n)
		return
	}
}

func CompareTreeNodes(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !CompareTreeNodes(t1.Left, t2.Left) {
		return false
	}
	if !CompareTreeNodes(t1.Right, t2.Right) {
		return false
	}
	return true
}

func Show(t *TreeNode) {
	var walk func(root *TreeNode, level int)

	var ret [][]int
	walk = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(ret) {
			ret = append(ret, []int{root.Val})
		} else {
			ret[level] = append(ret[level], root.Val)
		}

		fmt.Printf("-> Val %d l %d\n", root.Val, level)
		walk(root.Left, level+1)
		walk(root.Right, level+1)
	}
	walk(t, 0)
	fmt.Println(ret)

}
