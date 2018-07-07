package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		{1, []int{1, 1, 2}, 2},
		{2, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tc := range tests {
		if ret := removeDuplicates(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
