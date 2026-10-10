package leetcode

// the optimized 3 pass solution that i need to understand
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for curr := head; curr != nil; {
		next := curr.Next
		copy := &Node{Val: curr.Val, Next: next}
		curr.Next = copy
		curr = next
	}

	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}

	newHead := head.Next
	for curr := head; curr != nil; {
		copy := curr.Next
		curr.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		curr = curr.Next
	}

	return newHead
}

// my solution using hash function
func copyRandomList(head *Node) *Node {
	D := &Node{}
	tail := D
	set := make(map[*Node]*Node)

	for orig := head; orig != nil; orig = orig.Next {
		tail.Next = &Node{
			Val: orig.Val,
		}
		tail = tail.Next
		set[orig] = tail
	}

	for ; head != nil; head = head.Next {
		set[head].Random = set[head.Random]
	}

	return D.Next
}

// func copyList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
// 	return &Node{
// 		Val:  head.Val,
// 		Next: copyList(head.Next),
// 	}
// }

// AI version
// func copyRandomList(head *Node) *Node {
// 	var n_head *Node
// 	var tail *Node
// 	set := make(map[*Node]*Node)

// 	for head != nil {
// 		node := &Node{
// 			Val: head.Val,
// 		}

// 		if n_head == nil {
// 			n_head = node
// 			tail = node
// 		} else {
// 			tail.Next = node
// 			tail = tail.Next
// 		}
// 		set[head] = tail
// 		head = head.Next
// 	}

// 	temp := head
// 	for temp != nil {
// 		set[temp].Random = set[temp.Random]
// 	}

// 	return n_head
// }
