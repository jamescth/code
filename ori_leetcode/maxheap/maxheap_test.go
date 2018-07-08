package maxheap

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		num  int
		nums IntHeap
		ret  int
	}{
		{0, []int{2, 1, 5}, 1},
		{1, []int{2, 0, 5, 1}, 1},
	}

	for _, tc := range tests {
		ret := maxHeap(tc.nums)
		if ret != tc.ret {
			t.Errorf("%d failed %v expec %d got %d", tc.num, tc.nums, tc.ret, ret)
		}
		t.Log(tc.nums)
	}
}
