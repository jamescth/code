package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lessHead := &ListNode{}
	bigHead := &ListNode{}

	lessTail, bigTail := lessHead, bigHead

	for head != nil {
		if head.Val < x {
			lessTail.Next = head
			lessTail = lessTail.Next
		} else {
			bigTail.Next = head
			bigTail = bigTail.Next
		}
		head = head.Next
	}

	lessTail.Next = bigHead.Next
	bigTail.Next = nil
	head = lessHead.Next

	return head
}
