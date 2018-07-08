package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
	. "github.com/jamescth/code/leetcode/util"
)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}

	var bst func([]int) *TreeNode
	bst = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		mid := len(nums) / 2

		return &TreeNode{
			Val:   nums[mid],
			Left:  bst(nums[:mid]),
			Right: bst(nums[mid+1:]),
		}
	}

	return bst(tmp)
}
