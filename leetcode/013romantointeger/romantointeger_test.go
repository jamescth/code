package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret int
	}{
		{1, "III", 3},
		{2, "IV", 4},
		{3, "IX", 9},
		{4, "LVIII", 58},
		{5, "MCMXCIV", 1994},
	}

	for _, tc := range tests {
		if ret := romanToInt(tc.x); ret != tc.ret {
			t.Errorf("%d failed %v expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
