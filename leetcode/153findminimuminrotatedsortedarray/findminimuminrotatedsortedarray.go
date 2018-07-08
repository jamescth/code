package main

func findMin(nums []int) int {
	if nums == nil {
		return 0
	}
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	i, min := 1, nums[0]
	for i < l {
		if nums[i] < min {
			min = nums[i]
		}

		i++
	}
	return min
}
