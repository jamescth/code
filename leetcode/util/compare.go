package util

func Compare(a, b []int) bool {
	la, lb := len(a), len(b)

	if la != lb {
		return false
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
