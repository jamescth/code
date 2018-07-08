package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		num  int
		strs string
		ret  bool
	}{
		//{1, "{}", true},
		//{2, "{}[]()", true},
		//{3, "{)", false},
		{4, "([)]", false},
	}

	for _, tc := range tests {
		if ret := isValid(tc.strs); ret != tc.ret {
			t.Errorf("%d failed %v expect %v got %v\n", tc.num, tc.strs, tc.ret, ret)
		}
	}

}
