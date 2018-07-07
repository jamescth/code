package main

import (
	"sort"
)

func subsets(nums []int) [][]int {
	ret := [][]int{}

	recur(nums, []int{}, &ret)

	return ret
}

func recur(nums, temp []int, ret *[][]int) {
	// fmt.Printf("1st nums %v\n", nums)
	l := len(nums)
	if l == 0 {
		sort.Ints(temp)

		// Here is where we actuall append
		// we can add condition such as only 2 elements sets
		// if lt := len(temp); lt == 2 {
		*ret = append(*ret, temp)
		//}
		// fmt.Printf("sort ret %v temp %v\n", ret, temp)
		return
	}

	// fmt.Printf("before 2nd nums %v temp %v ret %v\n", nums, temp, ret)
	recur(nums[:l-1], temp, ret)
	// fmt.Printf("before 3rd nums %v temp %v ret %v\n", nums, temp, ret)
	recur(nums[:l-1], append([]int{nums[l-1]}, temp...), ret)
	// fmt.Printf("after 3rd nums %v temp %v ret %v\n", nums, temp, ret)
}
