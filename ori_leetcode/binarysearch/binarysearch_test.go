package binarysearch

import "testing"

func TestBS(t *testing.T) {
	tests := []struct {
		num int
		lst []int
		t   int
		idx int
	}{
		{0, []int{1}, 1, 0},
		{1, []int{}, 1, -1},
		{2, []int{1, 2}, 1, 0},
		{3, []int{1, 2, 3, 4, 5, 6}, 1, 0},
		{4, []int{1, 2, 3, 4, 5, 6}, 3, 2},
		{5, []int{1, 2, 3, 4, 5, 6}, 4, 3},
		{6, []int{1, 2, 3, 4, 5, 6}, 6, 5},
		{7, []int{1, 2, 3, 4, 5, 6}, 0, -1},
		{8, []int{1, 2, 3, 4, 5, 6}, 7, -1},
	}
	for _, tc := range tests {
		ret := bs(tc.lst, tc.t)
		if ret != tc.idx {
			t.Errorf("%d failed %v target %d expec %d got %d", tc.num, tc.lst, tc.t, tc.idx, ret)
		}

	}
}
