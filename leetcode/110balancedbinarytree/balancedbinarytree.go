package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getDepth(root.Left)
	r := getDepth(root.Right)

	delta := l - r
	if delta > 1 || delta < -1 {
		return false
	}

	if isBalanced(root.Right) && isBalanced(root.Left) {
		return true
	}
	return false
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	self := 1
	l := getDepth(root.Left)
	r := getDepth(root.Right)
	if l > r {
		return l + self
	} else {
		return r + self
	}
}

/*
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right == nil && root.Left == nil {
		return true
	}

	var walk func(root *TreeNode) (int, bool)
	walk = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		if root.Right == nil && root.Left == nil {
			return 1, true
		}

		l, lb := walk(root.Left)
		r, rb := walk(root.Right)
		var max int
		if l > r {
			max = l + 1
		} else {
			max = r + 1
		}
	    if lb && rb && (l == r || l+1 == r || r+1 == l) {
		    return max,true
	    }
		return max, false
	}

	l, lb := walk(root.Left)
	r, rb := walk(root.Right)

	if lb && rb && (l == r || l+1 == r || r+1 == l) {
		return true
	}
	return false
}
*/
