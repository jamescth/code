package main

func isPerfectSquare(num int) bool {
	switch num {
	case 0, 1:
		return true
	}

	i := 0
	for i = 0; i < num; i++ {
		if i*i >= num {
			break
		}
	}

	if num%(i*i) == 0 {
		return true
	}
	return false
}
