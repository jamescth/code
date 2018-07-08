package main

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

	ret max
}
