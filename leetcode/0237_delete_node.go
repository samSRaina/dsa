package leetcode

// perfect solution for given circumstance only
func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next // since it isnt te last node no nil pointer errors
}

// higher runtime but correct
func deleteNode(node *ListNode) {
	prev := node
	for node.Next != nil {
		node.Val = node.Next.Val
		prev = node
		node = node.Next
	}

	prev.Next = nil
}
