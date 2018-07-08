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
		num int
		k   int
		ret []int
	}{
		{1, 1, []int{1, 1}},
		{1, 3, []int{1, 3, 3, 1}},
		{1, 5, []int{1, 5, 10, 10, 5, 1}},
	}

	for _, tc := range tests {
		if ret := getRow(tc.k); !compare(ret, tc.ret) {
			t.Errorf("%d fails k %d expect %v got %v \n", tc.num, tc.k, tc.ret, ret)
		}
	}
}
