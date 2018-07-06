package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		n   int
		x   int
		ret int
	}{
		{1, 123, 321},
		{2, -123, -321},
		{3, 120, 21},
	}

	for _, tc := range tests {
		if ret := reverse(tc.x); ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
