package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/util"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		m   int
		k   []int
		nn  int
		ret []int
	}{
		{1, []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tc := range tests {
		merge(tc.x, tc.m, tc.k, tc.nn)
		if !util.CompareIntArray(tc.x, tc.ret) {
			t.Errorf("%d failed %v %v expect %v\n", tc.n, tc.x, tc.k, tc.ret)
		}
	}
}
