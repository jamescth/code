package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret int
	}{
		{1, "abcabcbb", 3},
		{2, "bbbbb", 1},
		{3, "pwwkew", 3},
	}

	for _, tc := range tests {
		if ret := lengthOfLongestSubstring(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
