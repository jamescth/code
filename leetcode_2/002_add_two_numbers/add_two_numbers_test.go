package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Insert(n int) {

}

func Insert(root **ListNode, n int) {
	if *root == nil {
		*root = &ListNode{Val: n}
		return
	}

	tmp := *root
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp = &ListNode{Val: n}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		num int
		l1  []int
		l2  []int
		ret []int
	}{
		{1, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	}

	for _, tc := range tests {
		var l1, l2, ret *ListNode
		for _, n := range tc.l1 {
			Insert(&l1, n)
		}
		for _, n := range tc.l2 {
			Insert(&l2, n)
		}
		for _, n := range tc.ret {
			Insert(&ret, n)
		}
	}
}
