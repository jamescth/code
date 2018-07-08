package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	var ret []int

	for i := 0; i < len(numbers); i++ {
		d := target - numbers[i]

		//fmt.Printf("%d %v m %v %d\n", i, numbers, m, m[i])
		// we cant get the target
		if _, ok := m[numbers[i]]; !ok {
			// we save the delta. so we can lookup number[i]
			// 9-2 = 7, look for 7 in dict
			m[d] = i + 1
			continue
		}

		// we have a match
		ret = []int{m[numbers[i]], i + 1}
		break
	}

	return ret
}
