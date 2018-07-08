package main

import (
	"testing"
)

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	l := len(a)

	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMaxSet(t *testing.T) {
	tests := []struct {
		num  int
		nums []int
		ret  []int
	}{
		{1, []int{0, 0, -1, 0}, []int{0, 0}},
		{2, []int{-846930886, -1714636915, 424238335, -1649760492}, []int{424238335}},
		{3, []int{756898537, -1973594324, -2038664370, -184803526, 1424268980}, []int{1424268980}},
		{3, []int{24115, -75629, -46517, 30105, 19451, -82188, 99505, 6752, -36716, 54438, -51501, 83871, 11137, -53177, 22294, -21609, -59745, 53635, -98142, 27968, -260, 41594, 16395, 19113, 71006, -97942, 42082, -30767, 85695, -73671}, []int{41594, 16395, 19113, 71006}},
	}

	for _, tc := range tests {
		if ret := maxset(tc.nums); !compare(ret, tc.ret) {
			t.Errorf("%d fails %v expect %v got %v\n", tc.num, tc.nums, tc.ret, ret)
		}
	}
}
