package main

import (
	"fmt"
	"sort"
)

/* 39
给定一个无重复元素的数组 candidates和一个目标数target，
找出candidates中所有可以使数字和为target的组合。
candidates中的数字可以无限制重复被选取。
*/
func combinationSum(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	var find func(start int, use []int, remain int)
	find = func(start int, use []int, remain int) {
		for i := start; i < len(candidates); i++ {
			c := candidates[i]
			if c == remain {
				result = append(result, append([]int(nil), append(use, c)...))
				// result = append(result, append(use, c)) todo 为何出现 [2 2 2 2 2 7 3 3]
			} else if c < remain {
				find(i, append(use, c), remain-c)
			} else {
				return
			}
		}
	}
	find(0, []int{}, target)
	return result
}

// python版本实现
func combinationSum11(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	fmt.Println(candidates)
	result = make([][]int, 0)
	find(0, []int{}, target, candidates, &result)
	return result
}

func find(s int, use []int, remain int, candidates []int, result *[][]int) {

	for i := s; i < len(candidates); i++ {
		var c = (candidates)[i]
		if c < remain {
			temp := append(use, c)
			find(i, temp, remain-c, candidates, result)
		}
		if c == remain {
			temp := append(use, c)
			fmt.Println("当前result=", result)
			fmt.Println("添加了", temp)
			*result = append(*result, temp)
		}
		if c > remain {
			return
		}
	}
}
func main() {
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
	// fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	/*
		[7],
		[2,2,3]
	*/
}

//  leetcode-official:https://assets.leetcode-cn.com/solution-static/39/39_fig1.png
func combinationSumOfficial(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(remain, idx int) // target=remain idx=candidates索引位置
	dfs = func(remain, idx int) {
		c := candidates[idx]
		if idx == len(candidates) {
			return
		}
		if remain == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(remain, idx+1)
		// 选择当前数
		if remain-c >= 0 {
			comb = append(comb, c)
			dfs(remain-c, idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

// new
func combinationSum33(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(nums []int, remain int) (result [][]int) {
	for i := 0; i < len(nums); i++ {
		if remain < nums[i] {
			break
		} else if remain == nums[i] {
			result = append(result, []int{nums[i]})
			continue
		}
		for _, v := range dfs(nums[i:], remain-nums[i]) {
			result = append(result, append([]int{nums[i]}, v...))
		}

	}
	return result
}
