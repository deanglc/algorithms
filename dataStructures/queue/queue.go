package queue

// 当你想要按顺序处理元素时，使用队列可能是一个很好的选择。
import "github.com/pkg/errors"

// queue vs list difference????
type Queue struct {
	data []interface{}
	len  int
}

func New() *Queue {
	return &Queue{}

}

func (q *Queue) Length() int {
	return q.len
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

// 入队，Enqueue 插在最后一个
func (q *Queue) Append(element interface{}) {
	q.data = append(q.data, element)
	q.len++
}

// 出队 Dequeue 弹出第一个
func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("空队列，无法弹出")
	}

	el := q.data[0]
	q.data = q.data[1:] // 切片，裁掉了第一个元素
	q.len--
	return el, nil
}

// python3 list.index
func (q *Queue) Index(element interface{}) (int, error) {
	for index, el := range q.data {
		if el == element {
			return index, nil
		}
	}
	return -1, errors.New("查无此值")
}
