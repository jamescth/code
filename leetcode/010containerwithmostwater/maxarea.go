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

/*
func maxArea(height []int) int {
    	l, r := 0, len(height)-1

	var max, tmp int
	for l < r {
		lheight, rheight := height[l], height[r]
		if lheight > rheight {
			tmp = (r - l) * rheight
		} else {
			tmp = (r - l) * lheight
		}

		if max < tmp {
			max = tmp
		}

		if lheight > rheight {
			for l < r {
				if rheight < height[r] {
					break
				}
				r--
			}
		}

		if rheight >= lheight {
			for l < r {
				if lheight < height[l] {
					break
				}
				l++
			}
		}

	}

    return max
}
*/
