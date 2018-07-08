package main

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var idx int
	var ret int
	// walk returns total nodes
	var walk func(root *TreeNode, n int) int
	walk = func(root *TreeNode, n int) int {
		if root == nil {
			return 0
		}
		total := walk(root.Left, n)
		idx++
		if idx == n {
			ret = root.Val
			return total
		}
		t1 := walk(root.Right, n)
		return total + t1
	}
	walk(root, k)
	return ret

}
