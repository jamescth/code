package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
	}{
		{1, []int{1, 1, 2, 3, 3}},
	}

	for _, tc := range tests {
		var l *ListNode
		for _, n := range tc.nums {
			Insert(&l, n)
		}
		deleteDuplicates(l)
		Show(l)
	}
}
