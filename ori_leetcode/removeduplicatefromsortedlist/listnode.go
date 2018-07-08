package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Insert(root **ListNode, n int) {
	fmt.Printf("%v\n", root)
	if *root == nil {
		*root = &ListNode{Val: n}
		return
	}

	// fmt.Printf("Next %v\n", (*root).Next)
	if (*root).Next == nil {
		(*root).Next = &ListNode{Val: n}
		return
	}

	tmp := (*root).Next
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &ListNode{Val: n}
}

func Show(root *ListNode) {
	for root != nil {
		fmt.Printf("-> %d ", root.Val)
		root = root.Next
	}
	fmt.Printf("\n")
}
