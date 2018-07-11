package main

import "testing"

func TestT(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		b   []string
		ret string
	}{
		{1, "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
	}

	for _, tc := range tests {
		if ret := mostCommonWord(tc.x, tc.b); ret != tc.ret {
			t.Errorf("%d fail %s %v expect %s got %s\n", tc.n, tc.x, tc.b, tc.ret, ret)
		}
	}
}
