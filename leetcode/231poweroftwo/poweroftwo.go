package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	// right shift
	for {
		d := n & 1
		n = n >> 1
		if d == 1 {
			if n == 0 {
				return true
			}
			return false
		}
	}
	return false
}
