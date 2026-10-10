package leetcode

func isPalindrome(head *ListNode) bool {
	/*if head == nil || head.Next == nil {
		return true
	}*/
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nxt := slow.Next
		slow.Next = prev
		prev = slow
		slow = nxt
	}
	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true
}

// too much stuff going on:: not a really great solution
// func isPalindrome(head *ListNode) bool {

// 	mid := findMiddleNode(head)
// 	sec := reverseList(mid)
// 	first := head

// 	for sec != nil {
// 		if first.Val != sec.Val {
// 			return false
// 		}

// 		first = first.Next
// 		sec = sec.Next
// 	}

// 	return true
// }

// func findMiddleNode(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	return slow
// }
