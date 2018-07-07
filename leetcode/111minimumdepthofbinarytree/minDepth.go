package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	self := 1

	// only if either left or right is not empty, we need to walkthrough
	if root.Left == nil {
		return self + minDepth(root.Right)
	} else if root.Right == nil {
		return self + minDepth(root.Left)
	}
	// both are not empty
	l := self + minDepth(root.Left)
	r := self + minDepth(root.Right)
	if l < r {
		return l
	}
	return r
}
