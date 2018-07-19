package main

import "github.com/jamescth/code/leetcode/treenode"

func lca(root *treenode.TreeNode, p, q int) *treenode.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p || root.Val == q {
		return root
	}

	if root.Val > p && root.Val > q {
		return lca(root.Left, p, q)
	}
	if root.Val < p && root.Val < q {
		return lca(root.Right, p, q)
	}
	return root
}
