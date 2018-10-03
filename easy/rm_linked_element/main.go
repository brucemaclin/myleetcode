package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	var tmpHead = head
	for tmpHead != nil && tmpHead.Val == val {
		tmpHead = tmpHead.Next
	}
	prev := tmpHead
	node := tmpHead
	for node != nil {
		if node.Val == val {
			prev.Next = node.Next

		} else {
			prev = node
		}
		node = prev.Next
	}
	return tmpHead
}
func main() {

}
