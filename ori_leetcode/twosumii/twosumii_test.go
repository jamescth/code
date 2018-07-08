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
		ret  []int
	}{
		{1, []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{2, []int{5, 25, 75}, 100, []int{2, 3}},
	}

	for _, tc := range tests {
		if ret := twoSum(tc.nums, tc.k); !compare(ret, tc.ret) {
			t.Errorf("%d fails %v k %d expect %v got %v\n", tc.num, tc.nums, tc.k, tc.ret, ret)
		}
	}
}
