package main

import (
	"strconv"

	. "github.com/jamescth/code/leetcode/treenode"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var ret []string

	var walk func(*TreeNode, string)
	walk = func(root *TreeNode, str string) {
		if root == nil {
			return
		}

		str = str + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ret = append(ret, str)
			return
		}

		str = str + "->"
		walk(root.Left, str)
		walk(root.Right, str)
	}

	walk(root, "")
	return ret

}
