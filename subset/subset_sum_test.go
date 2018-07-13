package main

import (
	"fmt"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	s := []int{2, 4, 6, 10}

	ret := count_set(s, 16)

	fmt.Println(ret)
}

func count_set(nums []int, total int) int {
	return rec(nums, total, len(nums)-1)
}

var ct int

// nums: array
// total: total amount
// i: index of nums element i

// [2,4,6,10]
// i = 2, total = 5 , should get [2,4], [6]
// var ret []int

func rec(nums []int, total, i int) int {
	ct++
	// fmt.Printf("%d total %d i %d nums %v\n", ct, total, i, nums)
	if total == 0 {
		fmt.Printf("total 0, i %d nums %v\n", i, nums)
		return 1
	} else if total < 0 {
		// fmt.Printf("%d total %d\n", ct, total)
		return 0
	} else if i < 0 {
		return 0
	} else if total < nums[i] {
		// if total < nums[i]
		// we cant find any subset add up to total, remove nums[i]

		// fmt.Printf("%d total %d nums[i] %d i %d\n", ct, total, nums[i], i)
		return rec(nums, total, i-1)
	} else {
		//
		// fmt.Printf("%d total %d -nums[i] %d i %d \n", ct, total, nums[i], i)
		return rec(nums, total-nums[i], i-1) + rec(nums, total, i-1)
	}
}
