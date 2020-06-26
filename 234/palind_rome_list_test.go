package _1

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
			},
			},
		},
	}

	m := isPalindromeListNode( node)
	print(m)

}