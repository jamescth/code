package main

func compress(chars []byte) int {
	l := len(chars)
	if l < 3 {
		return l
	}

	i := 0
	for i < l {
		for i+1 < l && chars[i] != chars[i+1] {
			i++
		}

		// same char

	}
}
