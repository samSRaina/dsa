package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	d := &ListNode{Val: 0, Next: head}
	prev := d

	temp := head
	for temp.Next != nil {
		after := temp.Next
		temp.Next = after.Next
		after.Next = prev.Next
		prev.Next = after
	}
	return d.Next
}
