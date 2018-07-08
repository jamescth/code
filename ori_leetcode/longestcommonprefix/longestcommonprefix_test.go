package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		num  int
		strs []string
		ret  string
	}{
		{1, []string{"flower", "flow", "flight"}, "fl"},
		{1, []string{"", ""}, ""},
		{1, []string{"c", "c"}, "c"},
	}

	for _, tc := range tests {
		if ret := longestCommonPrefix(tc.strs); ret != tc.ret {
			t.Errorf("%d failed %v expect %v got %v\n", tc.num, tc.strs, tc.ret, ret)
		}
	}

}
