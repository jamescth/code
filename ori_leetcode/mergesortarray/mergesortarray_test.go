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

func TestMerge(t *testing.T) {
	tests := []struct {
		num   int
		nums1 []int
		m     int
		nums2 []int
		n     int
		ret   []int
	}{
		{1, []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tc := range tests {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		if !compare(tc.nums1, tc.ret) {
			t.Errorf("%d failed %v got %v\n", tc.num, tc.nums1, tc.ret)
		}
	}
}
