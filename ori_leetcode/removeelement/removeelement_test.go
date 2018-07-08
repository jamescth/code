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
		k    int
		ret  int
	}{
		{1, []int{3, 3}, 3, 0},
		{2, []int{3}, 2, 1},
	}

	for _, tc := range tests {
		if ret := removeElement(tc.nums, tc.k); ret != tc.ret {
			t.Errorf("%d fails %v k %d expect %v got %d\n", tc.num, tc.nums, tc.k, tc.ret, ret)
		}
	}
}
