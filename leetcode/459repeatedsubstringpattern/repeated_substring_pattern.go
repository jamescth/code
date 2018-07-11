package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	//fmt.Println(s + s)
	//fmt.Println((s + s)[1 : size*2-1])
	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}
