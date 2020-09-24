package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"
imp：有DP那味了
*/
/*
字串长度 = 2 			 	一定是回文串
字串长度 = 3 & s[0]=s[2] 	一定是回文串

*/
// func longestPalindrome(s string) string {
// 	var dp [][]int
// 	n := len(s)
// 	dp = make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = make([]int, n)
// 	}
// 	ans := ""
// 	for l := range s {
// 		for i := range s {
// 			println(l, i, s)
// 			j := i + l
// 			if l == 0 {
// 				dp[i][j] = 1
// 			} else if l == 1 {
// 				if s[i] == s[j] {
// 					dp[i][j] = 1
// 				} else {
// 					dp[i][j] = 0
// 				}
// 			} else {
// 				if dp[i+1][j-1] == 1 && s[i] == s[j] {
// 					dp[i][j] = 1
// 				} else {
// 					dp[i][j] = 0
// 				}
// 			}
// 			if dp[i][j] == 1 && l+1 > len(ans) {
// 				ans = s[i : j+1]
// 			}
// 		}
// 	}
// 	return ans
// }
//
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	fmt.Println(dp)
	return ans
}

func main() {
	fmt.Println(longestPalindrome("dnglclg"))
}
