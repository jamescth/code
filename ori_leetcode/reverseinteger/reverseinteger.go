package reverseinteger

func reverse(x int) int {
	if x == 0 {
		return x
	}

	var ret int
	workInt := x
	if x < 0 {
		workInt = workInt * -1
	}

	for {
		if workInt == 0 {
			break
		}

		ret = ret * 10
		d := workInt % 10
		ret += d
		workInt /= 10

	}

	if ret > 1<<31-1 {
		return 0
	}
	if x < 0 {
		ret *= -1
	}

	return ret
}
