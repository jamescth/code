package longestoflis

import "testing"

func TestBS(t *testing.T) {
	tests := []struct {
		num int
		lst []int
		ret int
	}{
		{0, []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tc := range tests {
		ret := longestOfLIS(tc.lst)
		if ret != tc.ret {
			t.Errorf("%d failed %v expec %d got %d", tc.num, tc.lst, tc.ret, ret)
		}

	}
}
