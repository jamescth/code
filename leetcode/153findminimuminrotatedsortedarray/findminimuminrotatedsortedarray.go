package main

func findMin(nums []int) int {
	L := len(nums)
	if L == 0 {
		return -1
	}

	left, right := 0, L-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[right]
}
