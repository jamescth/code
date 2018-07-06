package main

func lengthOfLongestSubstring(s string) int {
	ckRepeat := func(s string) int {
		m := make(map[byte]struct{})

		i, l := 0, len(s)
		for i < l {
			_, ok := m[s[i]]
			if !ok {
				m[s[i]] = struct{}{}
			} else {
				break
			}
			i++
		}

		return i
	}

	lStr, idx := len(s), 0
	max := 0
	for idx < lStr {
		r := ckRepeat(s[idx:])
		if r > max {
			max = r
		}
		idx++
	}

	return max
}
