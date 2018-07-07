package main

import (
	. "github.com/jamescth/code/leetcode/util"
)

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var ret *ListNode
	ret = minValPop(&l1, &l2)
	tmp := ret

	// no nil lists
	for l1 != nil && l2 != nil {
		t := minValPop(&l1, &l2)
		tmp.Next = t
		tmp = tmp.Next
		/*
			Show(ret)
			Show(tmp)
			Show(l1)
			Show(l2)
		*/
		if l1 == nil || l2 == nil {
			break
		}

		// Next is done in Pop
		//l1 = l1.Next
		//l2 = l2.Next
	}

	// Show(ret)

	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}

	return ret
}

func minValPop(l1, l2 **ListNode) *ListNode {
	/*
		fmt.Printf("Pop\n")
		Show(*l1)
		Show(*l2)
	*/
	if (*l1).Val < (*l2).Val {
		ret := *l1
		*l1 = (*l1).Next
		ret.Next = nil
		return ret
	}
	ret := *l2
	*l2 = (*l2).Next
	ret.Next = nil
	return ret
}

/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root, tmp *ListNode
	if l1.Val < l2.Val {
		root, tmp = l1, l1
		l1 = l1.Next
	} else {
		root, tmp = l2, l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		tmp.Next = l2
	}
	if l2 == nil {
		tmp.Next = l1
	}
	return root
}
*/
