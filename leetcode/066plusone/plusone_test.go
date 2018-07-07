package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/util"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret []int
	}{
		{1, []int{1, 2, 3}, []int{1, 2, 4}},
		{2, []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}

	for _, tc := range tests {
		if ret := plusOne(tc.x); !util.CompareIntArray(ret, tc.ret) {
			t.Errorf("%d failed %v expect %v got %v\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
