package main

func singleNumber(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	m := make(map[int]struct{})

	for i := 0; i < l; i++ {
		_, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = struct{}{}
		} else {
			delete(m, nums[i])
		}
	}

	ret := 0
	for k, _ := range m {
		ret = k
	}
	return ret

}
