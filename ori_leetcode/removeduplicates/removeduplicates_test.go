package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  int
	}{
		{1, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		//{2, "", 0},
	}

	for _, tc := range tests {
		ret := removeduplicates(tc.nums)
		if tc.ret != ret {
			t.Errorf("%d []int %v expect %d got %d\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
