package main

import (
	"fmt"
	"math"
	"math/bits"
	"testing"
)

func TestSubsetCount(t *testing.T) {
	nums := []int{2, 4, 6, 10}
	x := math.Pow(2, float64(len(nums))) - 1

	var i uint = 0
	for ; i <= uint(x); i++ {
		// if the bits count is != k, skip
		if bits.OnesCount(i) != 2 {
			continue
		}

		// we have bit counts equal to 2
		// let's get them out
		for j := 0; j < len(nums); j++ {
			shift := 1 << uint(j)

			// if shift is part of i
			// for example:
			// i = 101, both 100 and 001 are part of shift
			if uint(shift)&i > 0 {
				// fmt.Printf("nums[j] %d j %d shift %d i %d\n", nums[j], j, shift, i)
				fmt.Printf("%d ", nums[j])
			}
		}
		fmt.Printf("\n")
	}
}
