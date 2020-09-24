package queue_test

import (
	"github.com/deanglc/algorithms/dataStructures/queue"
	"testing"
)

func TestNew(t *testing.T) {
	q := queue.New()
	if !q.IsEmpty() {
		t.Error("err=新队列不为空")
	}

	if q.Length() != 0 {
		t.Error("err=新队列的长度不为0，长度为", q.Length())
	}
}

func TestQueue_Index(t *testing.T) {
	q := queue.New()
	q.Append(1)
	q.Append(2)
	q.Append(3)
	index, err := q.Index(3)
	if err != nil {
		t.Error(err.Error())
	} else if index != 2 {
		t.Error("索引值返回出错")
	}

}

func TestQueue_Pop(t *testing.T) {
	q := queue.New()
	list := []int{1, 2, 3}
	for _, el := range list {
		q.Append(el)
	}
	for index, el := range list {
		if data, _ := q.Pop(); data != el && q.Length() != 2-index {
			t.Errorf("Poped value should be %d, but got %v", el, data)
		}
	}

	_, err := q.Pop()
	if err == nil {
		t.Error("弹出空队列，应报错。但没报错，测试成功")
	}
}
