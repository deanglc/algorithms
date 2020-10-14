package _42

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
		9	8	7
		10		6
1 	2	3	4	5
*/

func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			break
		}
	}
	if f == nil || f.Next == nil {
		return nil
	}
	s = head
	for s != f {
		s = s.Next
		f = f.Next
	}
	return s
}
