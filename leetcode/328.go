package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := odd.Next
	p := even

	for p != nil && p.Next != nil {
		odd.Next = p.Next
		p.Next = p.Next.Next

		odd = odd.Next
		p = p.Next
	}

	odd.Next = even

	return head
}
