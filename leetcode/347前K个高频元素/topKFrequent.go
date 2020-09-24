package main

import "fmt"

/*
347 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例1 	输入: nums = [1,1,1,2,2,3], k = 2
		输出: [1,2]
示例2 	输入: nums = [1], k = 1
		输出: [1]
*/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	l := make([]int, len(nums)+1)
	for k, v := range m {
		l[k] = v
	}
	return l[len(m)+1-k : len(m)+1]
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
	fmt.Println(topKFrequent([]int{-1, -1}, 1))
}
