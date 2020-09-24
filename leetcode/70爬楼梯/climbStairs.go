package main

/* 70-DP
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数

*/
func climbStairs(n int) int {
	var a, b int = 1, 2
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	println(climbStairs(2))
	println(climbStairs(3))
	println(climbStairs(4))
	println(climbStairs(5))
}
