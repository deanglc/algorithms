package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newL := new(ListNode)
	curr := newL
	num := 0

	for l1 != nil || l2 != nil || num > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		curr.Val = num % 10
		num = num / 10
	}
	return newL.Next
}
