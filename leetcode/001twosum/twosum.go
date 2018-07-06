package main

func twoSum(nums []int, target int) []int {

	// a k/v store
	// k: target-nums[i]
	// v: i
	m := make(map[int]int)
	var ret []int
	i, l := 0, len(nums)

	for i < l {
		v := nums[i]
		idx, ok := m[v]
		// fmt.Printf("%d nums %v v %d ok %v\n", i, nums, v, ok)
		if !ok {
			m[target-v] = i
		} else {
			ret = []int{idx, i}
			// fmt.Printf("%d nums %v\n", i, nums)
		}
		i++
	}
	return ret
}
