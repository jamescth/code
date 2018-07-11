package maxheap

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		n    int
		nums []int
		ret  int
	}{
		{0, []int{2, 1, 5}, 5},
		{1, []int{2, 0, 5, 1}, 5},
	}

	for _, tc := range tests {
		tmp := make(IntHeap, len(tc.nums))
		copy(tmp, tc.nums)
		// fmt.Println(tmp, tc.nums)

		ret := maxHeap(tmp)
		if ret != tc.ret {
			t.Errorf("%d failed %v expec %d got %d", tc.n, tc.nums, tc.ret, ret)
		}
		t.Log(tc.nums)
	}
}
