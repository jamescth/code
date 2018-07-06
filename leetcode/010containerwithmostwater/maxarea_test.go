package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		{1, []int{1, 1}, 1},
	}

	for _, tc := range tests {
		if ret := maxArea(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
