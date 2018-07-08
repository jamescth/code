package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

/*
func reverseList(head *ListNode) *ListNode {
    	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	var prev *ListNode
	newHead := head
	newHead.Next = prev
	// walkthrough until no more
	for tmp != nil {
		prev = tmp
		tmp = tmp.Next
		prev.Next = newHead
		newHead = prev
	}

    return newHead
}
*/
