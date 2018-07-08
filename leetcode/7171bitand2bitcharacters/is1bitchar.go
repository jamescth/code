package main

func isOneBitCharacter(bits []int) bool {
	L := len(bits)
	if L == 1 {
		if bits[0] == 1 {
			return false
		}
		return true
	}

	i := 0
	// i<L
	var ret bool = true
	for i < L {
		if bits[i] == 1 {
			i = i + 2

			ret = false

		} else {
			i = i + 1
			ret = true
		}

	}

	return ret
}
