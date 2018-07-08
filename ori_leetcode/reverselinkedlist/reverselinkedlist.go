package main

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
		// for i := 0; i < 3; i++ {
		// fmt.Printf("prev %v newHead %v tmp %v\n", prev, newHead, tmp)
		prev = tmp
		tmp = tmp.Next
		prev.Next = newHead
		newHead = prev
		// fmt.Printf("  prev %v newHead %v tmp %v\n", prev, newHead, tmp)
	}

	return newHead
}
