package main

import (
	"reflect"
	"testing"
)

func TestFindMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		k   int
		ret []int
	}{
		{1, []int{-4, 2, -5, 3, 6}, []int{2, 2, 1, 6}},
	}

	for _, tc := range tests {
		if ret := findMaxSlidingWindow(tc.x); !reflect.DeepEqual(ret, tc.ret) {
			t.Errorf("%d failed %v %d expect %v got %v\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
