package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		num  int
		strs string
		ret  int
	}{
		//{1, " ", 0},
		//{1, "", 0},
		{3, "a ", 1},
		{4, "Hello world", 5},
	}

	for _, tc := range tests {
		if ret := lengthOfLastWord(tc.strs); ret != tc.ret {
			t.Errorf("%d failed %v expect %v got %v\n", tc.num, tc.strs, tc.ret, ret)
		}
	}

}
