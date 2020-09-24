package main

/*
给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
// 因为可以排序、交换 所以不用DP
// 利用map的hash/python.set
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		numSet[nums[i]] = true
	}
	longStreak := 0
	for num := range numSet {
		// 当前的num不存在前驱
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			// 得出当前num开始序列的连续长度
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longStreak {
				longStreak = currentStreak
			}
		}
	}

	return longStreak
}
