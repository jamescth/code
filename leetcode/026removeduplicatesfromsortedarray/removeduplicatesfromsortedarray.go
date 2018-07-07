package main

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	for i+1 < len(nums) {
		if nums[i] == nums[i+1] {
			// rmIdx := i + 1
			nums = append(nums[:i+1], nums[i+2:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

/*
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	i, end := 0, l-1

	for i < end {
		if nums[i] == nums[i+1] {
			moveIntToEnd(nums[i+1])
			end--
		} else {
			i++
		}
	}
	return i + 1
}

func moveIntToEnd(nums []int) {
	i := 0
	for i+1 < len(nums) {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		i++
	}
	// fmt.Printf("movetoend %v i %d\n", nums, i)
}
*/

/*
func removeDuplicates(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}
	if l == 1 {
		return 1
	}

	ret, idx, iStore := 0, 0, 0

	for idx < l {
		nums[iStore] = nums[idx]
		for idx+1 < l && nums[idx] == nums[idx+1] {
			idx++
		}
		//fmt.Printf("idx %d nums %v\n", idx, nums)
		ret++
		idx++
		iStore++
	}
	return ret
}
*/
