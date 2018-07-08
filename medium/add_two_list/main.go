package  main

 // Definition for singly-linked list.
 type ListNode struct {
 	Val int
      Next *ListNode
  }
//l1：(2 -> 4 -> 3)
// l2:(5 -> 6 -> 4)
//out：7 -> 0 -> 8
// for 342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var over int
	root := l1
	prev1 := l1
	prev2 := l2
	var pprev = l1
	var final *ListNode
	for prev1 != nil && prev2 != nil {
		prev1.Val += prev2.Val +over
		over = 0
		if prev1.Val >= 10 {
			over = 1
			prev1.Val -=10
		}
		pprev = prev1
		prev1 = prev1.Next
		prev2 = prev2.Next
	}
	final = pprev
	for prev1 != nil {
		prev1.Val += over
		over= 0
		if prev1.Val >= 10 {
			over = 1
			prev1.Val -= 10
		}
		final = prev1
		prev1 = prev1.Next
	}

	if prev2!= nil {
		pprev.Next = prev2
		pprev = pprev.Next
		for pprev != nil {
			pprev.Val += over
			over = 0
			if pprev.Val >= 10 {
				over = 1
				pprev.Val -=10
			}
			final = pprev
			pprev = pprev.Next
		}
	}
	if over == 1 {
		final.Next = new(ListNode)
		final.Next.Val = 1
	}
	return root
}
func main()  {

}