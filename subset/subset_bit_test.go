package main

import (
	"fmt"
	"math"
	"testing"
)

func IntToBitIdx(n int) []int {
	ret := []int{}

	for i := 0; n > 0; i++ {
		if (n & 0x1) == 1 {
			ret = append(ret, i)
		}
		n = n >> 1
	}

	return ret
}

func TestIntToBit(t *testing.T) {
	fmt.Println("TestIntToBit")
	tests := []struct {
		num int
		n   int
	}{
		{1, 7},
		{2, 6},
	}

	for _, tc := range tests {
		fmt.Println(IntToBitIdx(tc.n))
	}
}

func TestSubsetBit(t *testing.T) {
	fmt.Println("TestSubsetBit")
	nums := []int{2, 4, 5, 8}
	num := math.Pow(2, float64(len(nums)))

	for i := 0; i < int(num); i++ {
		//fmt.Printf("%b\n", i)
		for _, v := range IntToBitIdx(i) {
			fmt.Printf("%d ", nums[v])
		}
		fmt.Printf("\n")
	}
	fmt.Println("End TestSubsetBit")
}
