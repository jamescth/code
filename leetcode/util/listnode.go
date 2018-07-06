package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Insert(root **ListNode, n int) {
	// fmt.Printf("root %v %d\n", *root, n)
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

func CompareListNodes(r1, r2 *ListNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil {
		return false
	}
	if r2 == nil {
		return false
	}

	for r1 != nil && r2 != nil {
		if r1.Val != r2.Val {
			return false
		}
		r1, r2 = r1.Next, r2.Next
	}

	if r1 != nil || r2 != nil {
		return false
	}
	return true
}

func Show(root *ListNode) {
	for root != nil {
		fmt.Printf("-> %d ", root.Val)
		root = root.Next
	}
	fmt.Printf("\n")
}
