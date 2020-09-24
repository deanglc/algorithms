package main

import "fmt"

func minWindow(s string, t string) (ans string) {
	m := map[byte]int{}
	n := len(s)
	rk := -1
	ans = s
	check := func(tt string) bool {
		//
		for _, v := range tt {
			if m[byte(v)] == 0 {
				return true // 不含全部t的字母
			}
		}
		return false // 含全部t的字母
	}
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		if i == 9 {
			fmt.Println("tellme")
		}
		for rk+1 < n && check(t) {
			m[s[rk+1]]++
			rk++
		}
		if rk-i+1 < len(ans) && !check(t) {
			ans = s[i : rk+1]
			fmt.Println("1111", ans, "rk=", rk)
		}
	}
	return ans
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
