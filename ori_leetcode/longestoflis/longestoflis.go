package longestoflis

import (
	"sort"
)

func longestOfLIS(nums []int) int {
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

	/*
		tails := make([]int, 0, len(nums))

		for _, n := range nums {
			at := sort.SearchInts(tails, n)
			fmt.Printf("n %d tails %v at %d\n", n, tails, at)
			if at == len(tails) {
				// n 比 tails 中所有的数都大
				// 把 n 放入 tails 的尾部
				tails = append(tails, n)
			} else if tails[at] > n {
				// tails[at-1] < n < tails[at]
				// tails[at] = n, 不会改变 tail 的递增性，却增加了加入更多数的可能性
				tails[at] = n
			}
		}

		return len(tails)
	*/
}
