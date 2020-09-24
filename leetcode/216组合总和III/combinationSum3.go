package main

/*
Find all possible combinations of k numbers that add up to a number n,
given that only numbers from 1 to 9 can be used and each combination should be
a unique set of numbers.

Note:
1.All numbers will be positive integers.
2.The solution set must not contain duplicate combinations.

*/
// 和为n=remain 个数为k  每个数取值区间 1～9
// todo one more time
func combinationSum3(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}
