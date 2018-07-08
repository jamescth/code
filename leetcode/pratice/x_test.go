package main

import "testing"

func TestBS(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		v   int
		ret int
	}{
		{1, []int{1, 2, 3, 4, 5}, 4, 3},
		{2, []int{1, 3, 5, 6}, 2, 1},
	}

	for _, tc := range tests {
		if ret := binarySearch(tc.x, tc.v); ret != tc.ret {
			t.Errorf("%d failed %v %d expect %d got %d\n", tc.n, tc.x, tc.v, tc.ret, ret)
		}
	}
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		}
	}
	return low
}
