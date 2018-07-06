package main

func reverse(x int) int {
	// x != 0
	var ret int
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}

	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if negative {
		ret *= -1
	}
	return ret
}
