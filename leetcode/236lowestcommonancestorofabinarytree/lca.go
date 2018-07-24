package main

func lca(root *TreeNode, p, q int) *TreeNode {
	// if looking for me, return self
	if root.Val == p || root.Val == q {
		return root
	}

	var left, right *TreeNode

	// look for both childs
	if root.Left != nil {
		left = lca(root.left, p, q)
	}
	if root.Right != nil {
		right = lca(root.Right, p, q)
	}

	// if both childs found, return self
	if left != nil && right != nil {
		return root
	} else {
		// either one of the chidren returned a node, meaning either p or q
		// found on left or right branch. Example: assuming 'p' found in
		// left child, right child returned 'None'. This means 'q' is
		// somewhere below node where 'p' was found we dont need to search
		// all the way, because in such scenarios, node where 'p' found is LCA
		if left != nil {
			return left
		}
		return right
	}
}
