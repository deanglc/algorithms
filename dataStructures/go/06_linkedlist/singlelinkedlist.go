package _6_linkedlist

import "fmt"

/*
单链表基本操作
*/

type ListNode struct {
	Next *ListNode
	Val  interface{}
}

type LinkedList struct {
	Head   *ListNode
	Length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.Next
}

func (this *ListNode) GetValue() interface{} {
	return this.Next
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// 在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.Next
	p.Next = newNode
	newNode.Next = oldNext
	this.Length++
	return true
}

// 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.Head {
		return false
	}
	cur := this.Head.Next
	pre := this.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur
	this.Length++
	return true
}

// 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.Head, v)
}

// 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.Head
	for nil != cur.Next {
		cur = cur.Next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.Length {
		return nil
	}
	cur := this.Head.Next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

// 删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.Head.Next
	pre := this.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	pre.Next = p.Next
	p = nil
	this.Length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.Head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
