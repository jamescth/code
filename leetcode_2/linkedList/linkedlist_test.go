package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
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
	tmp.Next = &ListNode{Val: n}
}

func CompareListNodes(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	} else if l1 == nil {
		return false
	} else if l2 == nil {
		return false
	}

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func Show(root *ListNode) {
	for root != nil {
		fmt.Printf("-> %d", root.Val)
		root = root.Next
	}
	fmt.Printf("\n")
}

func TestListNode(t *testing.T) {
	tests := []struct {
		n int
		x []int
	}{
		{1, []int{1, 2, 3, 4}},
	}

	for _, tc := range tests {
		var l1, l2 *ListNode

		for _, x := range tc.x {
			Insert(&l1, x)
			Insert(&l2, x)
		}

		if !CompareListNodes(l1, l2) {
			Show(l1)
			Show(l2)
			t.Errorf("CompareListNode failed %v %v\n", l1, l2)
		}
	}
}
