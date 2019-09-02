package treenode

import "testing"

func TestTreeNode(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		{1, []int{3, 1, 4, 2}, 1},
		{2, []int{5, 3, 6, 2, 4, 1}, 1},
	}

	for _, tc := range tests {
		var t1, t2 *TreeNode
		for _, n := range tc.x {
			Insert(&t1, n)
			Insert(&t2, n)
		}
		Show(t1)
		if !CompareTreeNodes(t1, t2) {
			Show(t1)
			Show(t2)
			t.Errorf("%d failed\n", tc.n)
		}
	}
}
