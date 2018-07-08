package main

import "fmt"

func maxSubArray(A []int) int {
	l := len(A)
	if l == 1 {
		return A[0]
	}

	max, tmp := A[0], A[0]
	i := 1

	for i < l {
		fmt.Printf("%d tmp %d max %d\n", i, tmp, max)
		if tmp < 0 {
			tmp = A[i]
		} else {
			tmp += A[i]
		}
		if max < tmp {
			max = tmp
		}
		fmt.Printf("%d tmp %d max %d\n", i, tmp, max)
		i++
	}
	return max
}
