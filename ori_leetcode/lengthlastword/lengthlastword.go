package main

func lengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	ret, i := 0, l-1

	for ; i >= 0; i-- {
		if s[i] == ' ' {
			if ret != 0 {
				return ret
			}
			continue
		}

		// got not white space
		ret++
	}

	return ret
}
