package main

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		n int
		x []int
	}{
		{0, []int{1, 2}},
	}

	for _, tc := range tests {
		ret := subsets(tc.x)
		fmt.Printf("ret %v\n", ret)
	}
}
