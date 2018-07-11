package main

import "strings"

// wrong answer
func countSegments(s string) int {
	s = strings.TrimRight(s, " ")
	s = strings.TrimLeft(s, " ")
	return strings.Count(s, " ") + 1
}
