package main

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	i := 0
	for i < l {
		for i < l && nums[i] < target {
			i++
		}
		if i >= l {
			return i
		}
		// fmt.Printf("nums %v t %d i %d nums[i] %s\n", nums, target, i, string(nums[i]))
		if nums[i] == target {
			return i
		}
		return i
	}

	// non-reachable
	return 0
}
