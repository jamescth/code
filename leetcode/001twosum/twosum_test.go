package main

import (
	"testing"

	"github.com/jamescth/code/leetcode/util"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret []int
	}{
		{1, []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, tc := range tests {
		if ret := twoSum(tc.x, tc.k); !util.CompareIntArray(ret, tc.ret) {
			t.Errorf("%d failed %v k %d expect %v got %v\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
