package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len_a, len_b := getLength(headA), getLength(headB)

	if len_a > len_b {
		headA = headA.Next
		len_a--
	} else if len_a < len_b {
		headB = headB.Next
		len_b--
	}

	// len_a == len_b
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}
	return nil
}

func getLength(root *ListNode) int {
	var ret int
	for root != nil {
		root = root.Next
		ret++
	}
	return ret
}
