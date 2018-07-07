package main

func mySqrt(x int) int {
	i := 0
	if x < 2 {
		return x
	}
	for i*i <= x {
		i++
	}
	return i - 1
}
