package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		{1, []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{2, []int{1, 2}, 3},
	}

	for _, tc := range tests {
		if ret := maxSubArray(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
