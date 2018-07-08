package main

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
		i++
	}
	if max >= 2*smax {
		return idx
	}
	return -1

}
