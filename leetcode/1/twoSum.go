package main

import (
	"fmt"
)

// 4ms 3.3mb
func twoSum(nums []int, target int) []int {
	// 声明记得给出长度
	m := make(map[int]int, len(nums))
	for i := range nums {
		if _, exist := m[target-nums[i]]; exist {
			return []int{i, m[target-nums[i]]}
		}
		m[nums[i]] = i
	}
	return []int(nil)
}

func main() {
	// fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))
	fmt.Println(twoSum([]int{2, 7, 8, 9, 0}, 9))
}
