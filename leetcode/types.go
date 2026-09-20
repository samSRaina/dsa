package leetcode

// List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Tree Node for a BST
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Set type
type ListNodeSet map[*ListNode]struct{}

type IntSet map[int]struct{}
