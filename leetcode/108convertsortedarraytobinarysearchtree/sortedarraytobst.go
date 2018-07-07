package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	mid := l / 2
	root := &TreeNode{
		Val:   mid,
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return root
}
