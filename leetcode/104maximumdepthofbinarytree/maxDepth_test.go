package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/treenode"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		// order matters; This question is for Binary tree, not BST
		{1, []int{7, 15, 9, 20, 3}, 3},
	}

	// t.Skip("This question is for Binary tree, not BST")

	for _, tc := range tests {
		var r1 *treenode.TreeNode
		for _, n := range tc.x {
			treenode.Insert(&r1, n)
		}

		if ret := maxDepth(r1); ret != tc.ret {
			treenode.Show(r1, 0)
			t.Errorf("%d failed ret %d got %d\n", tc.n, tc.ret, ret)
		}
	}
}
