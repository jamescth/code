package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	tmp := ret
	carry := 0

	// exits when l1, l2 == nil
	for {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		total := v1 + v2 + carry
		if total >= 10 {
			total %= 10
			carry = 1
		} else {
			carry = 0
		}

		tmp.Val = total

		if l1 == nil && l2 == nil {
			break
		}

		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}

	if carry == 1 {
		tmp.Next = &ListNode{Val: carry}
	}

	return ret
}
