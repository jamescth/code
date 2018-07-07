package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		n   int
		x   []string
		ret string
	}{
		{1, []string{"flower", "flow", "flight"}, "fl"},
		{2, []string{"dog", "racecar", "car"}, ""},
		{3, []string{}, ""},
		{4, []string{"abab", "aba", ""}, ""},
		{3, []string{"a"}, "a"},
		{3, []string{"aa", "aa"}, "aa"},
	}

	for _, tc := range tests {
		if ret := longestCommonPrefix(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %s got %s\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
