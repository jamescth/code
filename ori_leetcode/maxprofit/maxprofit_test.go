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
		{1, []int{7, 1, 5, 3, 6, 4}, 5},
		//{2, []int{3}, 2, 1},
	}

	for _, tc := range tests {
		if ret := maxProfit(tc.nums); ret != tc.ret {
			t.Errorf("%d fails %v expect %v got %d\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
