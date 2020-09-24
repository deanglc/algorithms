package CircularQueue

//  GO-循环队列  leetcode-346
type MyCircularQueue struct {
	Data []int
	Head int
	Tail int
	Size int
}

/** Initialize your Data structure here. Set the Size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Data: make([]int, k),
		Head: -1,
		Tail: -1,
		Size: k,
	}

}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() == true {
		return false
	}
	if this.IsEmpty() == true {
		this.Head = 0
	}
	this.Tail = (this.Tail + 1) % this.Size
	this.Data[this.Tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() == true {
		return false
	}
	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
		return true
	}
	this.Head = (this.Head + 1) % this.Size
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() == true {
		return -1
	}
	return this.Data[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() == true {
		return -1
	}
	return this.Data[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return ((this.Tail + 1) % this.Size) == this.Head
}
func main() {
	obj := Constructor(3)
	param_1 := obj.EnQueue(1)
	param_11 := obj.EnQueue(2)
	param_111 := obj.EnQueue(3)
	param_1111 := obj.EnQueue(4)
	param_2 := obj.DeQueue()
	param_3 := obj.Front()
	param_4 := obj.Rear()
	param_5 := obj.IsEmpty()
	param_6 := obj.IsFull()
	println(param_1)
	println(param_11)
	println(param_111)
	println(param_1111)
	println(param_2)
	println(param_3)
	println(param_4)
	println(param_5)
	println(param_6)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * MyCircularQueue obj = new MyCircularQueue(k);
 * boolean param_1 = obj.EnQueue(value);
 * boolean param_2 = obj.DeQueue();
 * int param_3 = obj.Front();
 * int param_4 = obj.Rear();
 * boolean param_5 = obj.IsEmpty();
 * boolean param_6 = obj.IsFull();
 */
