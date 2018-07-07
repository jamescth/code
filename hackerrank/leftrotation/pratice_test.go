package main

import "testing"

func TestPratice(t *testing.T) {
	tests := []struct {
		n int
		x []int32
		k int32
	}{
		{1, []int32{1, 2, 3, 4, 5}, 4},
	}

	for _, tc := range tests {
		leftRotate(tc.x, tc.k)
	}
}
