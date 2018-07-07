package main

import (
	"fmt"

	. "github.com/jamescth/code/leetcode/util"
)

func reversePrint(head *ListNode) {
	if head == nil {
		return
	}

	reversePrint(head.Next)
	fmt.Printf(" %d", head.Val)
}
