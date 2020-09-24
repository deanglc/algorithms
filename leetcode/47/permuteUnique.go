package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	list := make([]int, 0, n)
	visited := make([]bool, n)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if n == idx {
			result = append(result, append([]int(nil), list...))
		}
		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			list = append(list, v)
			visited[i] = true
			backtrack(idx + 1)
			visited[i] = false
			list = list[:len(list)-1]
		}
	}
	backtrack(0)
	return result
}
