package leetcode

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{0, nil}
	carry := 0
	current := &dummyNode

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry

		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10

		current = current.Next

	}

	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}

	return dummyNode.Next
}

// 16 ms 5mb
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{0, nil}
	carry := 0
	current := &dummyNode
	for l1 != nil || l2 != nil {
		var x, y = 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		current = current.Next
	}
	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}
	return dummyNode.Next
}
