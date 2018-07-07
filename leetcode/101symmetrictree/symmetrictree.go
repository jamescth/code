package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var walktree func(*TreeNode, *TreeNode) bool
	walktree = func(l *TreeNode, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}

		if !walktree(l.Right, r.Left) {
			return false
		}

		if !walktree(l.Left, r.Right) {
			return false
		}
		return true
	}

	return walktree(root.Left, root.Right)
}
