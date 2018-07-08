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

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  int
	}{
		{1, []int{6, 5, 5}, 5},
	}

	for _, tc := range tests {
		if ret := majorityElement(tc.nums); ret != tc.ret {
			t.Errorf("%d fails %v expect %d got %d\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
