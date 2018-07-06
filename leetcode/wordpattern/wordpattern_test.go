package main

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		k   string
		ret bool
	}{
		{1, "abba", "dog cat cat dog", true},
		{2, "abba", "dog cat cat fish", false},
		{3, "aaaa", "dog cat cat dog", false},
		{4, "abba", "dog dog dog dog", false},
	}

	for _, tc := range tests {
		if ret := wordPattern(tc.x, tc.k); ret != tc.ret {
			t.Errorf("%d failed %s %s expect %v got %v\n", tc.n, tc.x, tc.k, tc.ret, ret)
		}
	}
}
