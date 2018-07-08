package main

import "testing"

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  int
	}{
		//{1, []int{3, 6, 1, 0}, 1},
		//{2, []int{0, 0, 3, 2}, -1},
		{3, []int{1, 0}, 0},
		//{2, []int{5, 25, 75}, 100, []int{2, 3}},
	}

	for _, tc := range tests {
		if ret := dominantIndex(tc.nums); ret != tc.ret {
			t.Errorf("%d fails %v expect %d got %d\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
