package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
func permute(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	list := make([]int, 0, n)
	visited := make([]bool, n)
	var dfs func(idx int)
	dfs = func(idx int) {
		if n == idx {
			result = append(result, append([]int(nil), list...))
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			list = append(list, nums[i])
			dfs(idx + 1)
			visited[i] = false
			list = list[:len(list)-1]
		}
	}
	dfs(0)
	return result
}
