package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret int
	}{
		{1, []int{1, 3, 5, 6}, 5, 2},
		{2, []int{1, 3, 5, 6}, 2, 1},
		{3, []int{1, 3, 5, 6}, 7, 4},
		{4, []int{1, 3, 5, 6}, 0, 0},
	}

	for _, tc := range tests {
		if ret := searchInsert(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d failed %v k %d expect %d got %d\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
