package main

import (
	"testing"
)

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	l := len(a)

	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMaxSet(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  int
	}{
		{1, []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tc := range tests {
		if ret := maxSubArray(tc.nums); ret != tc.ret {
			t.Errorf("%d fails %v expect %d got %d\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
