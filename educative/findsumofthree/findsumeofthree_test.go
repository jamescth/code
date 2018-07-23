package main

import "testing"

func TestFindSumOfThree(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret bool
	}{
		{1, []int{3, 7, 1, 2, 8, 4, 5}, 20, true},
		{2, []int{3, 7, 1, 2, 8, 4, 5}, 21, false},
	}
	for _, tc := range tests {
		if ret := findSumOfThree(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d failed %v %d expect %v got %v\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
