package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	// only leaf node counts
	if root.Right == nil && root.Left == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
