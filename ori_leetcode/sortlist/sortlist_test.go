package main

import "testing"

func TestSortList(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
	}{
		{1, []int{4, 2, 1, 3}},
	}
	for _, tc := range tests {
		var l *ListNode
		for _, n := range tc.nums {
			Insert(&l, n)
		}
		ret := sortList(l)
		Show(ret)
	}
}
