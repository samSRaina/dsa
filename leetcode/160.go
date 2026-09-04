package leetcode

// 68ms runtime, fml...(but completely self written..)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{}) // set variable of type *listnode

	temp := headA
	for temp != nil {
		seen[temp] = struct{}{}
		temp = temp.Next
	}

	temp = headB
	for temp != nil {
		if _, exists := seen[temp]; exists {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

// method to switch lists using two pointers for length sync eventually..
// Runtime: 100ms
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}

	}

	return p1
}
