package main

func strStr(haystack, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	if needle == "" {
		return 0
	}

	i := 0
	l := len(haystack)
	for i < l {
		if match(haystack[i:], needle) {
			return i
		}
		i++
	}

	return -1
}

func match(s, s1 string) bool {
	i, l, l1 := 0, len(s), len(s1)

	if l1 > l {
		return false
	}

	for i < l1 {
		if s[i] != s1[i] {
			return false
		}
		i++
	}
	return true
}
