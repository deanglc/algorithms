// package LCP19
package main

import (
	"fmt"
	"math"
)

func minimumOperations1(leaves string) int {
	count := 0

	left, right := -1, len(leaves)
	if leaves[left+1] != 'r' {
		fmt.Println("222222")
		count++
	}
	if leaves[right-1] != 'r' {
		count++
	}
	for left < right {
		if leaves[left+1] == 'r' {
			left++
			fmt.Println("left++")
		}
		if leaves[right-1] == 'r' {
			right--
			fmt.Println("right--")

		}
		if leaves[left+1] != 'r' && leaves[right-1] != 'r' {
			fmt.Println("over")
			break
		}
	}
	fmt.Println(left, right, count)

	// fmt.Println(leaves[left], leaves[right])
	for i := left + 1; i < right; i++ {
		if leaves[i] == 'r' {
			fmt.Println("i=", i, string(leaves[i]))
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(minimumOperations("ryyryyyrryyyyyryyyrrryyyryryyyyryyrrryryyyryrryrrrryyyrrrrryryyrrrrryyyryyryrryryyryyyyryyrryrryryy"))
	// fmt.Println(minimumOperations("ryryryy"))
}

const inf = math.MaxInt32 // 或 math.MaxInt64

func minimumOperations(leaves string) int {
	n := len(leaves)
	f := make([][3]int, n)
	f[0][0] = boolToInt(leaves[0] == 'y')
	f[0][1] = inf
	f[0][2] = inf
	f[1][2] = inf
	for i := 1; i < n; i++ {
		isRed := boolToInt(leaves[i] == 'r')
		isYellow := boolToInt(leaves[i] == 'y')
		f[i][0] = f[i-1][0] + isYellow
		f[i][1] = min(f[i-1][0], f[i-1][1]) + isRed
		if i >= 2 {
			f[i][2] = min(f[i-1][1], f[i-1][2]) + isYellow
		}
	}
	return f[n-1][2]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
