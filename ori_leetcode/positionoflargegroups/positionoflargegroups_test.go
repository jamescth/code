package main

import (
	"testing"
)

func compare2d(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	l := len(a)

	for i := 0; i < l; i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		ll := len(a[i])
		for j := 0; j < ll; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMaxSet(t *testing.T) {
	tests := []struct {
		num int
		s   string
		ret [][]int
	}{
		{1, "abcdddeeeeaabbbcd", [][]int{[]int{3, 5}, []int{6, 9}, []int{12, 14}}},
		//{2, []int{5, 25, 75}, 100, []int{2, 3}},
	}

	for _, tc := range tests {
		if ret := largeGroupPositions(tc.s); !compare2d(ret, tc.ret) {
			t.Errorf("%d fails %s expect %v got %v\n", tc.num, tc.s, tc.ret, ret)
		}
	}
}
