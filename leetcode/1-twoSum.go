package leetcode

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
//4ms 3.7mb
func twoSum1(nums []int, target int) []int {
	r := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := r[nums[i]]; ok {
			return []int{i, r[nums[i]]}
		}
		r[target-nums[i]] = i
	}
	return []int{1, -1}

}

// 52ms 2.9mb
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}

	}
	return []int{1}

}

func main() {
	a := twoSum1([]int{2, 2, 7, 8, 9, 11}, 9)
	fmt.Println(a)
}
