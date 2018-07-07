package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var walk func(root *TreeNode, level int)
	walk = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(ret) {
			ret = append(ret, []int{})
		}

		if level%2 != 0 {
			ret[level] = append(ret[level], root.Val)
		} else {
			ret[level] = append([]int{root.Val}, ret[level]...)
		}

		walk(root.Right, level+1)
		walk(root.Left, level+1)
	}

	walk(root, 0)
	return ret

}
