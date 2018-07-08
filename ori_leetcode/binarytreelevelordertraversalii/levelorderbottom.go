package binarytreelevelordertraversalii

import "fmt"

// A TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	store := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(store) == level {
			// insert a new element from the start
			store = append([][]int{[]int{}}, store...)
		}
		newlen := len(store)
		store[newlen-1-level] = append(store[newlen-1-level], root.Val)
		fmt.Println(store)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return store
}
