package main

import "testing"

func TestNthDigit(t *testing.T) {
	tests := []struct {
		n   int
		x   int
		ret int
	}{
		{1, 1, 1},
		{1, 11, 0},
	}

	for _, tc := range tests {
		ret := findNthDigit(tc.x)
		if ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
