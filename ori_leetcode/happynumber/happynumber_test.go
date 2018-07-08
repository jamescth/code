package main

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		num int
		n   int
		ret bool
	}{
		{1, 19, true},
	}

	for _, tc := range tests {
		if ret := isHappy(tc.n); ret != tc.ret {
			t.Errorf("%d failed %d expect %v got %v\n", tc.num, tc.n, tc.ret, ret)
		}
	}
}
