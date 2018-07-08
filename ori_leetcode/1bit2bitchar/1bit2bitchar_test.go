package main

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	tests := []struct {
		num  int
		bits []int
		ret  bool
	}{
		{1, []int{1, 1, 1, 0}, false},
	}

	for _, tc := range tests {
		if ret := isOneBitCharacter(tc.bits); ret != tc.ret {
			t.Errorf("%d failed %v expect %v got %v\n", tc.num, tc.bits, tc.ret, ret)
		}
	}

}
