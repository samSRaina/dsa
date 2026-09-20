package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	d := &ListNode{Val: 0, Next: head}
	prev := d
	temp := head

	for temp != nil {
		if temp.Val != val {
			prev = temp
			temp = temp.Next
		} else {
			prev.Next = temp.Next
			temp = temp.Next
		}
	}
	return d.Next
}
