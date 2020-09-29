package _208环路检测

/*给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-lcci

     		10-9-8

    		11	 7

       1-2-3-4-5-6
slow,fast = head,head 速度slow=1 fast=2
快慢指针运行轨迹 12 24 36 48 510 64 76 88
重置slow = head  fast不变，在相遇点
同速运行轨迹 头8 19 210 311 44
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 71.57% 95.58%
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	// 快慢指针开始动
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 如果传入的是无环链表 则在此结束运行 返回
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 重置slow = head  fast不变，在相遇点
	// 同速运行 直到相遇 相遇点=环路起点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
