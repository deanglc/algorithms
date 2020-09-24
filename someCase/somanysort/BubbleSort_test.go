package somanysort_test

import (
	"github.com/deanglc/algorithms/someCase/somanysort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	somanysort.BubbleSort(arr)
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i1, v1 := range arr {
		for i2, v2 := range res {
			if i1 == i2 && v1 != v2 {
				t.Error("冒泡排序出错")
			}
		}
	}
}
