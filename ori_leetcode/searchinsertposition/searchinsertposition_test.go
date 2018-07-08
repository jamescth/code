package searchinsertposition

import "testing"

func TestBS(t *testing.T) {
	tests := []struct {
		num int
		lst []int
		t   int
		idx int
	}{
		{0, []int{1}, 1, 0},
		{1, []int{}, 1, 0},
		{2, []int{1, 2}, 2, 1},
		{3, []int{1, 2}, 0, 0},
		{4, []int{1, 3, 5, 6}, 5, 2},
		{5, []int{1, 3, 5, 6}, 2, 1},
		{6, []int{1, 3, 5, 6}, 7, 4},
		{7, []int{1, 3, 5, 6}, 0, 0},
	}
	for _, tc := range tests {
		ret := searchInsert(tc.lst, tc.t)
		if ret != tc.idx {
			t.Errorf("%d failed %v target %d expec %d got %d", tc.num, tc.lst, tc.t, tc.idx, ret)
		}

	}
}
