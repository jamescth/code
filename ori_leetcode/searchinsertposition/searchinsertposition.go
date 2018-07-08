package searchinsertposition

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums == nil {
		return 0
	}

	L, R := 0, len(nums)-1

	for L <= R {
		M := (L + R) / 2

		if target > nums[M] {
			L = M + 1
			continue
		}
		if target < nums[M] {
			R = M - 1
			continue
		}
		if target == nums[M] {
			return M
		}
	}

	fmt.Printf("nums %v target %d L %d R %d\n", nums, target, L, R)

	return L
}
