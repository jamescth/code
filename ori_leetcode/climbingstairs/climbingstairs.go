package main

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b

	/*
		// too slow
		if n == 1 || n == 2 {
			return n
		}

		ret := climbStairs(n-2) + climbStairs(n-1)

		return ret
	*/
}
