package main

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret int
	}{
		{1, []int{3, 2, 3, 2}, 3, 2},
		{2, []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{3, []int{2, 2, 2}, 2, 0},
		{4, []int{3}, 2, 1},
		{5, []int{3, 3}, 2, 2},
		{6, []int{4, 5}, 4, 1},
	}

	for _, tc := range tests {
		if ret := removeElement(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d failed %v k %d expect %d got %d\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
