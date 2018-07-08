package main

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	tmp := h

	for {
		if tmp == nil || tmp.Next == nil {
			break
		}

		// walk through dups
		n := tmp.Next
		for n != nil && n.Val == tmp.Val {
			n = n.Next
		}

		tmp.Next = n
		tmp = tmp.Next
	}
	return h
}
