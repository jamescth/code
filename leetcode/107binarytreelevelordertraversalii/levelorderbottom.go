package main

import (
	. "github.com/jamescth/code/leetcode/treenode"
)

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var walktree func(root *TreeNode, level int)

	walktree = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(ret) {
			ret = append([][]int{[]int{}}, ret...)
		}

		newlen := len(ret)
		ret[newlen-1-level] = append(ret[newlen-1-level], root.Val)

		walktree(root.Left, level+1)
		walktree(root.Right, level+1)
	}
	walktree(root, 0)
	return ret
}

/*
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
            store = append([][]int{[]int{}},store...)
		}
		newlen := len(store)
        store[newlen-1- level] = append(store[newlen-1 - level], root.Val)
        //fmt.Println(store)
        dfs(root.Left, level+1)
		dfs(root.Right, level+1)

	}

	dfs(root, 0)

	return store
}
*/
