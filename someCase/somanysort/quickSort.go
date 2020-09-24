package somanysort

import (
	"fmt"
)

func Quick1Sort(arr []int, left, right int) {

	val := arr[(left+right)/2]
	i, j := left, right

	for arr[j] > val {
		j--
	}
	for arr[i] < val {
		i++
	}
	fmt.Println(i, j, val)
	arr[i], arr[j] = arr[j], arr[i]
	i++
	j--

	if i < right {
		Quick1Sort(arr, i, right)
	}
	if j > left {
		Quick1Sort(arr, left, j)
	}
}

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1

	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])

}
func main() {
	nums := []int{1, 4, 2, 6, 5, 3}
	Quick1Sort(nums, 0, 5)
	fmt.Println(nums)
}
