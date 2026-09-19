package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	D := &ListNode{
		Next: head,
	}

	slow := D
	fast := D

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return D.Next
}
