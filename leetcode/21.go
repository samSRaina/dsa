package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode

	if list1 == nil {
		return list2
	}

	if list1 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
			tail = tail.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
			tail = tail.Next
		}
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return head
}
