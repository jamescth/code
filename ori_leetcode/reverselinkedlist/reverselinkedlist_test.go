package main

import "testing"

func TestReverseList(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
	}{
		{1, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range tests {
		var l *ListNode
		for _, n := range tc.nums {
			Insert(&l, n)
		}
		ret := reverseList(l)
		Show(ret)
	}
}
