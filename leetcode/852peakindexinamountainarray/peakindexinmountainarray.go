package main

func peakIndexInMountainArray(A []int) int {
	max, maxIdx := 0, 0

	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max, maxIdx = A[i], i
		}
	}
	return maxIdx
}
