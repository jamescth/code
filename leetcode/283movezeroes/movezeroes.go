package main

func moveZeroes(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}

	i, zeroidx := 0, l-1

	for i <= zeroidx {
		for i <= zeroidx && nums[i] != 0 {
			i++
		}
		if i >= l {
			break
		}
		//fmt.Printf("%d %v\n", i, nums)
		// nums[i] is zero
		for j := i; j < zeroidx; j++ {
			nums[j], nums[j+1] = nums[j+1], nums[j]
			//fmt.Printf("%d %d %v\n", i, j, nums)
		}
		zeroidx--
		if nums[i] != 0 {
			i++
		}
		//fmt.Printf("%d zidx %d %v\n", i, zeroidx, nums)

	}

}
