package main

func bubble_sort(nums []int) {
	end := len(nums) - 1

	for {
		if end == 0 {
			break
		}

		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		end--
	}
}
