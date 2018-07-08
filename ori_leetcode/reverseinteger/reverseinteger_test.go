package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		num int
		i   int
		ret int
	}{
		{1, 321, 123},
		{1, 123, 321}}

	for _, tc := range tests {
		ret := reverse(tc.i)
		if tc.ret != ret {
			t.Errorf("%d int %d expect %d got %d\n", tc.num, tc.i, tc.ret, ret)
		}
	}
}
