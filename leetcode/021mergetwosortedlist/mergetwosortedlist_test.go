package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/util"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   []int
		ret []int
	}{
		{1, []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	}

	for _, tc := range tests {
		var l1, l2, l3 *util.ListNode
		for _, n := range tc.x {
			util.Insert(&l1, n)
		}
		for _, n := range tc.k {
			util.Insert(&l2, n)
		}
		for _, n := range tc.ret {
			util.Insert(&l3, n)
		}

		if ret := mergeTwoLists(l1, l2); !util.CompareListNodes(ret, l3) {
			t.Errorf("%d failed %v k %d expect %v got %v\n", tc.n, tc.x, tc.k, l3, ret)
		}
	}
}
