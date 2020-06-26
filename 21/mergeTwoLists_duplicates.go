package _6

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l1
	}
	if l2 == nil {
		return l2
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1, l2.Next)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}


func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode{
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var dummp = &ListNode{Val:0 }
	cur := dummp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummp.Next = l1
			l1 = l1.Next
		} else {
			dummp.Next = l2
			l2 = l2.Next
		}
		dummp = dummp.Next
		if l1 == nil {
			dummp.Next = l2
		}
		if l2 == nil {
			dummp.Next = l1
		}
	}

	return cur.Next
}