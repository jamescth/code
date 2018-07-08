package main

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		num int
		n   int
		ret int
	}{
		{1, 10, 4},
		{2, 20, 8},
	}
	for _, tc := range tests {
		if ret := countPrimes(tc.n); ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.num, tc.n, tc.ret, ret)
		}
	}
}
