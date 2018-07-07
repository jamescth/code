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

/*
func lengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	s1 := trimTailSpace(s)
	// fmt.Printf("s %s s1 %s\n", s, s1)
	if len(s1) == 0 {
		return 0
	}

	i := len(s1) - 1
	for i >= 0 {
		if s1[i] == ' ' {
			break
		}
		i--
	}
	// fmt.Printf("s %s s1 %s i %d\n", s, s1, i)

	if i < 0 {
		return len(s1)
	}

	return len(s1) - 1 - i
}

func trimTailSpace(s string) string {
	l := len(s)
	i := l - 1
	for i >= 0 {
		if s[i] != ' ' {
			break
		}
		i--
	}
	if i < 0 {
		return ""
	}
	return s[:i+1]
}
*/
