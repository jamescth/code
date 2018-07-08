package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	l := len(nums)
	if l == 0 {
		return nil
	}

	midx := l / 2

	t := &TreeNode{
		Val:   nums[midx],
		Left:  sortedArrayToBST(nums[:midx]),
		Right: sortedArrayToBST(nums[midx+1:]),
	}
	return t
}
