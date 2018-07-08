package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)

	if l == 1 {
		return
	}

	k = k % l
	r := l - k
	tmp := append(nums[r:l], nums[:r]...)
	fmt.Printf("k %d r %d l %d %v\n", k, r, l, tmp)
	for i := 0; i < l; i++ {
		nums[i] = tmp[i]
	}
}
