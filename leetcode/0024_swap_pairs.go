package leetcode

func swapPairs(head *ListNode) *ListNode {
	D := &ListNode{
		Next: head,
	}
	prev := D
	temp := head

	for temp != nil && temp.Next != nil {
		after := temp.Next
		temp.Next = after.Next
		after.Next = prev.Next
		prev.Next = after
		prev = temp
		temp = temp.Next
	}

	return D.Next
}
