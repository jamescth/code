package main

import "sort"

func findSumOfThree(nums []int, v int) bool {
	sort.Ints(nums)
	L := len(nums)
	for i := 0; i < L-2; i++ {
		j, k := i+1, L-1

		for j < k {
			t := nums[i] + nums[j] + nums[k]
			if t == v {
				return true
			}

			if t > v {
				k--
			} else {
				j++
			}
		}
	}
	return false
}
