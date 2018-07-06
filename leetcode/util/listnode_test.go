package util

import "testing"

func TestInsert(t *testing.T) {
	tests := []struct {
		n int
		x []int
	}{
		{1, []int{1, 2, 3, 4}},
	}

	for _, tc := range tests {
		var l1, l2 *ListNode
		for _, n := range tc.x {
			Insert(&l1, n)
			Insert(&l2, n)
		}
		if !CompareListNodes(l1, l2) {
			Show(l1)
			Show(l2)
			t.Errorf("CompareListNodes Failed %v %v\n", l1, l2)
		}
	}

}
