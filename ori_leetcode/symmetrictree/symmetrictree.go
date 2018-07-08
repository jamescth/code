package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var walk func(left *TreeNode, right *TreeNode) bool
	walk = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		return walk(left.Left, right.Right) && walk(left.Right, right.Left)
	}

	return walk(root.left, root.right)
}
