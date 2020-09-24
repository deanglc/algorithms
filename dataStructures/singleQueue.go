package dataStructures

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

/* refer 百度百科  - 队列=先进先出线性表
https://baike.baidu.com/item/%E9%98%9F%E5%88%97/14580481?fr=aladdin
队列是一种特殊的线性表，特殊之处在于它只允许在表的前端（front）进行删除操作，而在表的后端（rear）进行插入操作，
和栈一样，队列是一种操作受限制的线性表。进行插入操作的端称为队尾，进行删除操作的端称为队头。队列中没有元素时，称为空队列。

队列中插入一个队列元素称为入队，从队列中删除一个队列元素称为出队。因为队列只允许在一端插入，在另一端删除，
所以只有最早进入队列的元素才能最先从队列中删除，故队列又称为先进先出（FIFO—first in first out）线性表
类似 chan
ex:
Q(i) i=3,4,5,6,7,8
front = 2, rear = 8  队列Q拥有 rear-front个元素，即长度，长度为0则为空队列
假设Q队列已满
maxsize=
*/

// Queue 先进先出线性表
type Queue struct {
	MaxSize int
	Array   [4]int // 数组=> 模拟队列
	Front   int    // 指向队列首部
	Rear    int    // 队列尾部
}

// 入队
func (q *Queue) AddQueue(val int) error {
	// 判断队列是否已满(rear = maxsize - 1)
	if q.Rear == q.MaxSize-1 {
		return errors.New("队列已满")
	}
	//
	q.Rear++
	q.Array[q.Rear] = val
	return nil
}

// 显示队列
func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	// this.front
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.Array[i])
	}
	fmt.Println()
}

// 取出队列
func (q *Queue) GetQueue() (val int, err error) {
	// 判断队列是否为空
	if q.Rear == q.Front {
		return -1, errors.New("queue empty")
	}
	q.Front++
	val = q.Array[q.Front]
	return val, err
}

func main() {
	queue := &Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("输入add 表示添加队列数")
		fmt.Println("输入get 表示从队列获取数据")
		fmt.Println("输入show 表示显示队列")
		fmt.Println("输入exit")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要的队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列取出的值为", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
