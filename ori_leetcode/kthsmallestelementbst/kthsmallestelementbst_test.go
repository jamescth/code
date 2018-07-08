package kthsmallestelementbst

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		k    int
		ret  int
	}{
		{0, []int{5, 3, 6, 2, 4, 1}, 3, 3},
		{1, []int{3, 1, 4, 2}, 1, 1},
	}

	for _, tc := range tests {
		var r *TreeNode
		for _, n := range tc.nums {
			Insert(&r, n)
		}
		if ret := kthSmallest(r, tc.k); ret != tc.ret {
			Show(r, 0)
			t.Errorf("%d failed k %d expect %d got %d\n", tc.num, tc.k, tc.ret, ret)
		}
	}
}
