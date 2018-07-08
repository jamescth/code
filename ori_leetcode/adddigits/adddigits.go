package main

func addDigits(num int) int {
	var ret int

	// when addup < 10, break
	var d int
	for {
		d += num % 10
		num = num / 10

		if num == 0 {
			if d < 10 {
				ret = d
				break
			}
			// reset to new value
			num = d
			d = 0
		}
	}
	return ret
}
