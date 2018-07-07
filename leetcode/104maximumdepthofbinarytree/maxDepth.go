package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// self
	ret := 1

	v1 := ret + maxDepth(root.Left)
	v2 := ret + maxDepth(root.Right)

	if v1 > v2 {
		return v1
	}
	return v2
}
