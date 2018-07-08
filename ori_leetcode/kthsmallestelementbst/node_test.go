package kthsmallestelementbst

import "testing"

func TestNodeInsert(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  int
	}{
		{0, []int{3, 1, 4, 2}, 1},
		{1, []int{5, 3, 6, 2, 4, 1}, 1},
	}

	for _, tc := range tests {
		//r := &TreeNode{Val: 3}
		var r *TreeNode
		for _, n := range tc.nums {
			Insert(&r, n)
		}
		// Show(r, 0)
	}
}
