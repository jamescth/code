package main

import "testing"

func TestT(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret int
	}{
		{1, "leetcode", 0},
		{1, "loveleetcode", 2},
	}

	for _, tc := range tests {
		if ret := firstUniqChar(tc.x); ret != tc.ret {
			t.Errorf("%d fail %s expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
