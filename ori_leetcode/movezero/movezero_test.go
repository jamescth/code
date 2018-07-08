package main

import "testing"

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

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  []int
	}{
		//{1, []int{0, 0, 1}, []int{1, 0, 0}},
		{2, []int{2, 1}, []int{2, 1}},
	}

	for _, tc := range tests {
		moveZeros(tc.nums)
		if !compare(tc.nums, tc.ret) {
			t.Errorf("%d failed %v got %v\n", tc.num, tc.nums, tc.ret)
		}
	}
}
