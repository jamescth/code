package main

import "sort"

func lengthOfLIS(nums []int) int {
	if nums == nil {
		return 0
	}

	minTails := []int{}
	for _, n := range nums {
		// search n in minTails
		i := sort.SearchInts(minTails, n)
		// fmt.Printf("n %d tails %v at %d\n", n, minTails, i)
		if i == len(minTails) {
			minTails = append(minTails, n)
			continue
		}

		if i < len(minTails) {
			if minTails[i] > n {
				minTails[i] = n
			}
		}
	}

	return len(minTails)

}
