package main

func isHappy(n int) bool {
	fn := func(x int) int {
		ret := 0
		for x != 0 {
			c := x % 10
			ret += c * c
			x = x / 10
		}
		// fmt.Printf("%d\n", ret)
		return ret
	}

	// check if cycle, then stop
	slow, fast := fn(n), fn(fn(n))
	for slow != fast {
		slow = fn(slow)
		fast = fn(fn(fast))
	}

	// return true if ret == 1
	if slow == 1 {
		return true
	}
	return false

}
