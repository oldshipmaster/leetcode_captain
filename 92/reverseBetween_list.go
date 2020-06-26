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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//m<=n
	var m_pre *ListNode
	var m_next *ListNode
	var n_pre *ListNode
	var n_next *ListNode
	if m == n {
		return head
	}
	for head.Next != nil {


	}

}
