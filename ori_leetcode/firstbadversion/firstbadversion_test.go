package main

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		num  int
		vers []bool
		ret  int
	}{
		{1, []bool{false, false, false, true, true}, 3},
		{1, []bool{false, true, true}, 1},
		{1, []bool{true, true}, 0},
	}

	for _, tc := range tests {
		if ret := firstBadVersion(tc.vers); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d", tc.num, tc.vers, tc.ret, ret)
		}
	}

}
