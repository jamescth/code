package main

func maxArea(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	L, R := 0, length-1
	max := 0
	for L < R {
		Lh, Rh := height[L], height[R]
		X := R - L
		s := min(Lh, Rh) * X
		if max < s {
			max = s
		}
		if Lh < Rh {
			L++
		} else {
			R--
		}
		// L++
		// R--
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
