package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		n   int
		s   string
		ret bool
	}{
		{1, "A man, a plan, a canal: Panama", true},
		{2, "race a car", false},
	}

	for _, tc := range tests {
		if ret := isPalindrome(tc.s); ret != tc.ret {
			t.Errorf("%d failed %s expect %v got %v\n", tc.n, tc.s, tc.ret, ret)
		}
	}
}
