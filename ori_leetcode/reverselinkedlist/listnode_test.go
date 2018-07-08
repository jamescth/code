package main

import "testing"

func TestNodeInsert(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
	}{
		{1, []int{1, 2, 3, 4}},
	}

	for _, tc := range tests {
		var l *ListNode

		for _, n := range tc.nums {
			Insert(&l, n)
		}
		Show(l)
	}
}
