package main

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	i := 0
	max := nums[0]
	tmp := 0
	// total := 0
	for i < l {
		for j := i; j < l; j++ {
			tmp = tmp + nums[j]
			if nums[i] < 0 {
				if max < tmp {
					max = tmp
				}
				tmp = 0
			} else {
				if max < tmp {
					max = tmp
				}
			}
			// fmt.Printf("nums %v i %d j %d tmp %d max %d\n", nums, i, j, tmp, max)
		}
		tmp = 0
		i++
	}
	return max
}
