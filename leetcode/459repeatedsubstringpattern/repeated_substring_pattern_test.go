package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}
func Test(t *testing.T) {
	tests := []struct {
		n   int
		s   string
		ret bool
	}{
		{1, "abab", true},
		{2, "aba", false},
		{3, "abcabcabcabc", true},
	}

	for _, tc := range tests {
		if ret := repeatedSubstringPattern(tc.s); ret != tc.ret {
			t.Errorf("%d failed %s expect %v got %v\n", tc.n, tc.s, tc.ret, ret)
		}
	}
}
