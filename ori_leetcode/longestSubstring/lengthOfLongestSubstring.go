package main

func lengthOfLongestSubstring(s string) int {
	l := len(s)

	fn := func(s string) int {
		leng := len(s)
		if s == "" {
			return 0
		}

		m := make(map[byte]struct{})
		var tmp string
		for i := 0; i < leng; i++ {
			// check for dup
			if _, ok := m[s[i]]; !ok {
				tmp = tmp + string(s[i])
				m[s[i]] = struct{}{}
				continue
			}
			break
		}
		return len(tmp)
	}

	var max int
	for i := 0; i < l; i++ {
		t := fn(s[i:])
		if max < t {
			max = t
		}
	}

	return max
}
