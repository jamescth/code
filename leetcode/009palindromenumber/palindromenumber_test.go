package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		n   int
		x   int
		ret bool
	}{
		{1, 121, true},
		{2, -121, false},
	}

	for _, tc := range tests {
		if ret := isPalindrome(tc.x); ret != tc.ret {
			t.Errorf("%d failed %d expect %v got %v\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
