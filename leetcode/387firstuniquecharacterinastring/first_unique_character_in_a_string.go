package main

import (
	"math"
)

func firstUniqChar(s string) int {
	l := len(s)
	if l < 2 {
		return l - 1
	}

	m := make(map[string]int)
	for i, c := range s {
		b := string(c)
		if _, ok := m[b]; !ok {
			m[b] = i
		} else {
			m[b] = -1
		}
	}

	// fmt.Printf("m %v\n", m)

	var ret int = math.MaxInt32
	for k, v := range m {
		// fmt.Printf("v %d k %s\n", v, k)
		if v != -1 && v < ret {
			ret = v
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
