package main

import "strings"

func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	// fields is []string for each word
	fields := strings.Fields(str)

	if len(pattern) != len(fields) {
		return false
	}
	i, l := 0, len(pattern)
	// until both strings end
	for i < l {
		// check if pattern seen
		v1, ok := m1[pattern[i]]
		if !ok {
			m1[pattern[i]] = fields[i]
		} else {
			if v1 != fields[i] {
				return false
			}
		}

		v2, ok := m2[fields[i]]
		if !ok {
			m2[fields[i]] = pattern[i]
		} else {
			if v2 != pattern[i] {
				return false
			}
		}

		i++
	}
	return true
}
