package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func linkedlistCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		return false
	}
	return true
}
