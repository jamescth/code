package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin, maxLen := 0, 1

	// b is the begin of Palindrome
	// e is the end of Palindrome
	// i is the middle
	// for every loop
	// i-b = e-i and we compare s[b], s[e]

	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		b, e := i, i

		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// loop until no more same char
		}

		// move idx
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// search for the longest palindrome
			// 0...b..i..e...len(s)
		}

		newLen := e + 1 - b
		// need to update
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}

/*
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	walkStr := func(s string) string {
		i, l := 0, len(s)
		R, L := 0, 0
		var max string
		for i < l {

			// repeat chars
			for i+1 < l && s[i] == s[i+1] {
				tmp := i
				for (tmp+1) < l && s[tmp] == s[tmp+1] {
					tmp++
				}
				if len(max) < len(s[i:tmp+1]) {
					max = s[i : tmp+1]
				}
				i = tmp
			}

			// idx R vs. idx L; pick the min length
			length := min(i, l-1-i)
			L, R = i-length, i+length
			stmp := s[L : R+1]
			// fmt.Printf("%s %s i %d length %d L %d R %d\n", s, stmp, i, length, L, R)
			if isPalindromic(stmp) {
				if len(max) < len(stmp) {
					max = stmp
				}
			}

			i++
		}
		return max
	}

	max := ""
	i := 0
	l := len(s)

	for i < l {
		ret := walkStr(s[i:])
		if len(ret) > len(max) {
			max = ret
		}
		i++
	}

	return max
}

func isPalindromic(s string) bool {
	l := len(s)

	i, j := 0, l-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	// fmt.Printf("%s is true\n", s)
	return true
}
*/
