package _1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var s []int
	if head == nil {
		return true
	}
	for head.Next != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	s = append(s, head.Val)
	j := len(s)-1
	if j<= 0 {
		return true
	}
	for i := 0; i < len(s); i++ {

		if s[i] == s[j] {
			j--
		} else {
			return false
		}
	}
	return true
}

