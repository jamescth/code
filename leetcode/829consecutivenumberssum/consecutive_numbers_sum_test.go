package main

import (
	"testing"
)

func TestConsecutiveNumbersSum(t *testing.T) {
	tests := []struct {
		n   int
		x   int
		ret int
	}{
		{1, 5, 2},
		{2, 9, 3},
		{3, 15, 4},
	}
	for _, tc := range tests {
		if ret := consecutiveNumbersSum(tc.x); ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
