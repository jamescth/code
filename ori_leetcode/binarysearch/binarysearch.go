package binarysearch

func bs(nums []int, n int) int {

	l := len(nums)
	L, R := 0, l-1
	ret := -1

	for L <= R {
		M := (R + L) / 2
		if n > nums[M] {
			L = M + 1
			continue
		}
		if n < nums[M] {
			R = M - 1
			continue
		}
		if n == nums[M] {
			ret = M
			break
		}
	}
	//fmt.Println(ret)
	return ret
}
