package removeduplicates

func removeduplicates(nums []int) int {
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
