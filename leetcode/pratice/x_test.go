package main

import (
	"strings"
	"testing"
)

func TestBS(t *testing.T) {
	tests := []struct {
		n   int
		x   []int
		v   int
		ret int
	}{
		{1, []int{1, 2, 3, 4, 5}, 4, 3},
		{2, []int{1, 3, 5, 6}, 2, 1},
	}

	for _, tc := range tests {
		if ret := searchInsert(tc.x, tc.v); ret != tc.ret {
			t.Errorf("%d failed %v %d expect %d got %d\n", tc.n, tc.x, tc.v, tc.ret, ret)
		}
	}
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		}
	}
	return low
}

func TestParseIPV6(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret string
	}{
		{1, "192.168.0.0", "192.168.0.0"},
		{2, "2001:db8::1:0:0:1", "2001:db8:0:0:1:0:0:1"},
		{3, "::1", "0:0:0:0:0:0:0:1"},
		{4, "fe80::", "fe80:0:0:0:0:0:0:0"},
		{5, "fe80::1", "fe80:0:0:0:0:0:0:1"},
		{6, "fe80:0:0:0:204:61ff:fe9d:f156", "fe80:0:0:0:204:61ff:fe9d:f156"},
		{7, "fe80:0:0:204:61ff:fe9d:f156", "0:fe80:0:0:204:61ff:fe9d:f156"},
	}

	fn := func(c rune) bool {
		if c == '.' || c == ':' {
			return true
		}
		return false
	}
	_ = fn

	for _, tc := range tests {
		if ret := parseIPV6(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %s got %s\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}

func parseIPV6(addr string) string {
	str := strings.Split(addr, ":")

	// need to calculate how many empty []
	l := len(str)
	needed := 0

	if l < 8 {
		needed = 8 - l + 1
	}
	// fmt.Printf("%v len %d\n", str, len(str))

	var ret string
	var empty int
	for i, r := range str {
		//fmt.Printf("i %d r %v\n", i, r)
		if i == 0 && len(r) == 0 {
			// 1st []int is empty.
			// we are in ::1 case
			needed = 7
			for needed > 0 {
				ret += "0:"
				needed--
			}
			empty++
		} else if i == (l-1) && len(str[i]) == 0 {
			// this is fe00::
			// len is only 3
			ret += "0:"
			empty++
		} else if len(r) == 0 {
			// for non 0 empty []int
			for needed > 0 {
				ret += "0:"
				needed--
			}
			empty++
		} else {
			ret += r + ":"
		}
	}
	ret = ret[:len(ret)-1]
	if empty == 0 && l == 7 {
		ret = "0:" + ret
	}
	return ret
}
