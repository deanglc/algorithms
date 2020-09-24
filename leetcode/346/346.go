package main

/*给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算其所有整数的移动平均值。
示例
MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3
*/
import "github.com/deanglc/algorithms/dataStructures/CircularQueue"

/* todo
1.int转float
2.优化next代码
3.queue得换个？
*/
type MovingAverage struct {
	deque   CircularQueue.MyCircularQueue
	maxSize int
}

func Constructor(k int) MovingAverage {
	return MovingAverage{
		deque:   CircularQueue.Constructor(k),
		maxSize: k,
	}
}

func (this *MovingAverage) Next(value int) float32 {
	if this.deque.IsFull() {
		this.deque.DeQueue()
	}
	this.deque.EnQueue(value)
	var sum, length int
	for i := 0; i < len(this.deque.Data); i++ {
		if this.deque.Data[i] != 0 {
			length++
			sum = sum + this.deque.Data[i]
		}
	}
	return float32(sum) / float32(length)
}

func main() { // 1 5.5 5.3 6
	m := Constructor(3)
	println(m.Next(1))
	println(m.Next(10))
	println(m.Next(5))
	println(m.Next(3))

}
