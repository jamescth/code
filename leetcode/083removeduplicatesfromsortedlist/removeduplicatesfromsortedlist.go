package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head

	for tmp != nil && tmp.Next != nil {
		t1 := tmp.Next
		for t1 != nil && tmp.Val == t1.Val {
			t1 = t1.Next
		}
		tmp.Next = t1
		tmp = tmp.Next
	}

	return head

}
