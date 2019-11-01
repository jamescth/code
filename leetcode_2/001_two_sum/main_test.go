package twosum

import "testing"

func twosum(arr []int, target int) []int {
	l := len(arr)
	if l < 2 {
		return nil
	}

	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[arr[i]] = i
	}

	for i := 0; i < l; i++ {
		sub := target - arr[i]
		if _, ok := m[sub]; ok {
			return []int{i, m[sub]}
		}
	}
	return nil
}

func CompareIntArray(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	tests := []struct {
		num int
		x   []int
		t   int
		ret []int
	}{
		{1, []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, tc := range tests {
		if ret := twosum(tc.x, tc.t); !CompareIntArray(tc.ret, ret) {
			t.Errorf("%d Failed %v target %d expect %v got %v\n", tc.num, tc.x, tc.t, tc.ret, ret)
		}
	}

}
