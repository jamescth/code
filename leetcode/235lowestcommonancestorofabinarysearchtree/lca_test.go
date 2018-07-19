package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/treenode"
)

func TestLCA(t *testing.T) {
	tests := []struct {
		n         int
		x         []int
		p, q, ret int
	}{
		{1, []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 2, 0, 2},
		{2, []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 0, 2, 2},
		{3, []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 0, 5, 2},
		{4, []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 0, 9, 6},
	}
	for _, tc := range tests {
		var r1 *treenode.TreeNode
		for _, n := range tc.x {
			treenode.Insert(&r1, n)
		}
		// treenode.Show(r1, 0)

		if ret := lca(r1, tc.p, tc.q); ret.Val != tc.ret {
			treenode.Show(r1, 0)

		}
	}

}
