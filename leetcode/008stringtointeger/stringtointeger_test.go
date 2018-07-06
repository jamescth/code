package main

import "testing"

func TestStringToInteger(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret int
	}{
		{1, "42", 42},
		{2, "   -42", -42},
		{3, "4193 with words", 4193},
		{4, "words and 987", 0},
		{5, "-91283472332", -2147483648},
		{6, "+", 0},
		{7, "+-2", 0},
		{8, "9223372036854775808", 2147483647},
	}

	for _, tc := range tests {
		if ret := myAtoi(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
