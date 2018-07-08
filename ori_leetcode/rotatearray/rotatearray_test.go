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
		{1, []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{1, []int{1, 2}, 3, []int{2, 1}},
		{1, []int{1, 2}, 0, []int{1, 2}},
	}

	for _, tc := range tests {
		if rotate(tc.nums, tc.k); !compare(tc.nums, tc.ret) {
			t.Errorf("%d fails %v k %d expect %v got \n", tc.num, tc.nums, tc.k, tc.ret)
		}
	}
}
