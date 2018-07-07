package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		k   string
		ret int
	}{
		{1, "hello", "ll", 2},
		{2, "aaaaa", "bba", -1},
	}

	for _, tc := range tests {
		if ret := strStr(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d failed %s k %s expect %d got %d\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
