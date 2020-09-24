package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum22(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	var find func(start int, use []int, remain int)
	find = func(start int, use []int, remain int) {
		for i := start; i < len(candidates); i++ {
			c := candidates[i]
			if i > 0 && c == candidates[i-1] {
				continue
			} else if c == remain {
				result = append(result, append([]int(nil), append(use, c)...))
			} else if c < remain {
				find(i+1, append(use, c), remain-c)
			} else {
				return
			}
		}
	}
	find(0, []int{}, target)
	return result
}
func combinationSum2(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	return dfs(candidates, target)
}
func dfs(nums []int, remain int) (result [][]int) {
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		} else if remain < nums[i] {
			break
		} else if remain == nums[i] {
			result = append(result, []int{nums[i]})
			continue
		}
		for _, v := range dfs(nums[i+1:], remain-nums[i]) {
			result = append(result, append([]int{nums[i]}, v...))
		}

	}
	return result
}
