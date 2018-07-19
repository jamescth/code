package main

func search(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] == k {
			return mid
		}

		if nums[left] <= nums[mid] {
			// this can be done as binary search
			if k >= nums[left] && k < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if k > nums[mid] && k <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == k {
		return left
	}
	return -1
}
