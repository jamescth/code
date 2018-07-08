package main

import "fmt"

func dominantIndex(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}

	max, smax, idx, i := 0, 0, 0, 0

	for i < l {
		if nums[i] > max {
			smax, max, idx = max, nums[i], i
		} else if nums[i] > smax {
			smax = nums[i]
		}
		fmt.Printf("%d max %d smax %d idx %d\n", i, max, smax, idx)
		i++
	}
	if max >= 2*smax {
		return idx
	}
	return -1
}
