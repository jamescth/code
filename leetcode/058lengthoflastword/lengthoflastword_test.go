package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret int
	}{
		{1, "hello world", 5},
		{2, "hello world   ", 5},
		{3, "  ", 0},
		{4, "world", 5},
		{5, "a", 1},
	}

	for _, tc := range tests {
		if ret := lengthOfLastWord(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
