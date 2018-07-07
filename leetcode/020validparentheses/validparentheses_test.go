package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret bool
	}{
		{1, "()", true},
		{2, "()[]{}", true},
		{3, "(]", false},
		{4, "([)]", false},
		{5, "{[]}", true},
	}

	for _, tc := range tests {
		if ret := isValid(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %v got %v\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
