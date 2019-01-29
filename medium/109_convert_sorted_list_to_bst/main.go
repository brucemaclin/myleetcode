package main

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	preM := preMid(head)
	mid := preM.Next
	preM.Next = nil
	t := &TreeNode{Val: mid.Val}
	t.Left = sortedListToBST(head)
	t.Right = sortedListToBST(mid.Next)
	return t
}

func preMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}
