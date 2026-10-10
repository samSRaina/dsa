package leetcode

// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head

// 	seen := make(map[*ListNode]struct{})
// 	var cycle bool
// 	for fast != nil && fast.Next != nil && cycle != true {
// 		slow = slow.Next
// 		// if seen[slow] == struct{}{} {
// 		// 	return slow
// 		// } else {
// 		// 	seen[slow] = struct{}{}
// 		// }

// 		// if seen[fast.Next] == struct{}{} {
// 		// 	return fast.Next
// 		// }
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			cycle = true
// 		}
// 	}

// 	if cycle == true {
// 		seen[slow] = struct{}{}
// 		for true {
// 			slow = slow.Next
// 			if seen[slow] == struct{}{} {
// 				return slow
// 			} else {
// 				seen[slow] = struct{}{}
// 			}
// 	}
// 	return nil
// }

func detectCycle(head *ListNode) *ListNode {
	slow := head
	seen := make(map[*ListNode]struct{})

	for slow != nil {
		if seen[slow] == struct{}{} {
			return slow
		} else {
			seen[slow] = struct{}{}
		}
		slow = slow.Next
	}
	return nil
}
