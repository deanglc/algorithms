package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstringA("abcba")) // 3
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))      // 3
	// fmt.Println(lengthOfLongestSubstring("abcabcab"))    // 3
	// fmt.Println(lengthOfLongestSubstring("abcdabcebba")) // 5
	// fmt.Println(lengthOfLongestSubstring("abcdbe"))      // 4
}
func lengthOfLongestSubstring(s string) (result int) {
	need := make(map[byte]int)
	var left, right = 0, 0
	for i := 0; i < len(s); i++ {
		need[s[i]]++
		if need[s[i]] > 1 {
			result = max(result, i-left) // 对当前无重复存档
			// 重置need
			// 找到最先重复的索引值
			for left < i {
				need[s[left]]--
				left++
				if s[left] == s[i] {
					break
				}
			}
			for j := left; j < i; j++ {
				need[s[j]]--
				if s[j] == s[i] {
					left = j + 1
					break
				}
			}
		}
		right++
	}
	return max(result, right-left)
}

// abcb
func lengthOfLongestSubstringA(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
