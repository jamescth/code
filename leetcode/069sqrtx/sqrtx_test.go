package main

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		n   int
		x   int
		ret int
	}{
		{1, 4, 2},
		{2, 8, 2},
		{3, 3, 1},
	}

	for _, tc := range tests {
		if ret := mySqrt(tc.x); ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
