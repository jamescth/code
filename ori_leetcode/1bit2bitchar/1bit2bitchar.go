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
	ret := true

	for i < L {
		if bits[i] == 1 {
			i += 2
			ret = false
		} else {
			ret = true
			i++
		}
	}
	return ret
}
