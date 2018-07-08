package main

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num int
		n   int
		ret int
	}{
		{0, 43261596, 964176192},
	}

	for _, tc := range tests {
		if ret := reverseBits(tc.n); ret != tc.ret {
			t.Errorf("%d failed %d expect %d got %d\n", tc.num, tc.n, tc.ret, ret)
		}
	}
}
