package main

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
	}{
		{0, []int{-10, -3, 0, 5, 9}},
	}

	for _, tc := range tests {
		ret := sortedArrayToBST(tc.nums)
		Show(ret, 0)
	}
}
