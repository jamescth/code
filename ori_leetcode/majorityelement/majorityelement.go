package main

import "fmt"

func majorityElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	i, ret := 0, 0

	for i < l {
		v, ok := m[nums[i]]
		fmt.Printf("%d nums %v ok %v\n", i, nums, ok)
		if !ok {
			m[nums[i]] = 1
		} else {
			v++
			m[nums[i]] = v
			if v > (l-1)/2 {
				ret = nums[i]
				break
			}
		}
		i++
	}
	return ret
}
