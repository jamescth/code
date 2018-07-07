package treenode

import "testing"

func TestNodeInsert(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		{1, []int{3, 1, 4, 2}, 1},
		{2, []int{5, 3, 6, 2, 4, 1}, 1},
	}

	for _, tc := range tests {
		var r1, r2 *TreeNode
		for _, n := range tc.x {
			Insert(&r1, n)
			Insert(&r2, n)
		}
		if !CompareTreeNode(r1, r2) {
			Show(r1, 0)
			Show(r2, 0)
			t.Errorf("%d failed\n", tc.n)
		}
	}
}
