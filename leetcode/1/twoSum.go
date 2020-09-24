package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
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
