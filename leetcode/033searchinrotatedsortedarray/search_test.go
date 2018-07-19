package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret int
	}{
		//{1, []int{4, 5, 1, 2, 3}, 2, 3},
		{2, []int{5, 1, 2, 3, 4}, 1, 1},
	}

	for _, tc := range tests {
		if ret := search(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d fail %v %d expect %d got %d\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
