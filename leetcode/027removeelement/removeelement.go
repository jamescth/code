package main

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		// fmt.Printf("nums %v i %d\n", nums, i)
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}

/*
func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	//fmt.Println("new")
	S, E := 0, len(nums)-1
	for {
		for S < len(nums) && nums[S] != val {
			S++
		}

		//fmt.Printf("S %d E %d %v\n", S, E, nums)
		for E >= 0 && nums[E] == val {
			E--
		}

		if S >= E {
			break
		}
		//fmt.Printf(" S %d E %d %v\n", S, E, nums)
		nums[S], nums[E] = nums[E], nums[S]
	}

	//fmt.Printf("   S %d E %d %v\n", S, E, nums)
	return S

}
*/
