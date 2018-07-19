package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret int
	}{
		//{1, []int{4, 5, 1, 2, 3}, 2, 3},
		{2, []int{5, 1, 2, 3, 4}, 1},
	}

	for _, tc := range tests {
		if ret := findMin(tc.x); ret != tc.ret {
			t.Errorf("%d fail %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
