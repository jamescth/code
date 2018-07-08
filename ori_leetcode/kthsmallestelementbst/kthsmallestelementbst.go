package kthsmallestelementbst

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var idx int
	var ret int
	// walk returns total nodes
	var walk func(root *TreeNode, n int)
	walk = func(root *TreeNode, n int) {
		if root == nil {
			return
		}
		walk(root.Left, n)
		// fmt.Printf("l val %d total %d\n", root.Val, total)
		idx++
		if idx == n {
			// fmt.Println("set ret", root.Val, n)
			ret = root.Val
			return
		}
		walk(root.Right, n)
		// fmt.Printf("r val %d total %d\n", root.Val, t1)
		return
	}
	walk(root, k)
	return ret
}
