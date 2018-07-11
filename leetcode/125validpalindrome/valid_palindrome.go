package main

import "strings"

func isPalindrome(s string) bool {
	l := len(s)
	if l < 2 {
		return true
	}

	low, high := 0, l-1
	lowstr := strings.ToLower(s)

	for low < high {
		lb := lowstr[low]
		for low < high && (lb < 'a' || lb > 'z') {
			low++
			lb = lowstr[low]
		}
		hb := lowstr[high]
		for low < high && (hb < 'a' || hb > 'z') {
			high--
			hb = lowstr[high]
		}
		if !(low < high) {
			return true
		}

		if lowstr[low] != lowstr[high] {
			return false
		}

		low++
		high--
	}

	return true
}
