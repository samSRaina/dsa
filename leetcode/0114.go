package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	temp := root.Right
	flatten(root.Left)
	root.Right = root.Left
	root.Left = nil

	curr := root.Right
	for curr.Right != nil {
		curr = curr.Right
	}

	curr.Right = temp
	flatten(temp)
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	print(root)
	if root.Left != nil {
		preOrder(root.Left)
	}
	if root.Right != nil {
		preOrder(root.Right)
	}
}

func makeList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	return &ListNode{
		Val:  values[0],
		Next: makeList(values[1:]),
	}
}
