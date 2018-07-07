package main

import "testing"

func CompareIntSets(a, b [][]int) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}

	fn := func(suba, subb []int) bool {
		lsuba, lsubb := len(suba), len(subb)
		if lsuba != lsubb {
			return false
		}

		for i := 0; i < lsuba; i++ {
			if suba[i] != subb[i] {
				return false
			}
		}
		return true
	}

	for idx := 0; idx < la; idx++ {
		if !fn(a[idx], b[idx]) {
			return false
		}
	}
	return true
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		ret [][]int
	}{
		{1, []int{1, 2, 3}, [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}, []int{3},
			[]int{1, 3}, []int{2, 3}, []int{1, 2, 3}}},
	}

	for _, tc := range tests {
		if ret := subsets(tc.x); !CompareIntSets(ret, tc.ret) {
			t.Errorf("%v\n%v\n", ret, tc.ret)
		}
	}
}
