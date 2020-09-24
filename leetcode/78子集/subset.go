package main

import (
	"fmt"
)

// func backtrack(first int, pos int, curr *[][]int) {
// 	if len(*curr) == pos {
// 		*curr = append(*curr, *curr...)
// 	}
// 	for i := first; i< len(num)
// }
// func subsets(nums []int) (output *[][]int) {
// 	n := len(nums)
// 	for k := 0; k < n+1; k++ {
// 		backtrack(0, k, output)
// 	}
// 	return output
// }
// //
// func main() {
// 	fmt.Println(subsets([]int{1, 2, 3}))
// }
func subsets(nums []int) (result [][]int) {
	result = make([][]int, 0) // result 保存最终结果
	list := make([]int, 0)    // list 保存临时结果
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

func subset(nums []int) (result [][]int) {
	result = make([][]int, 0)
	list := make([]int, 0) // list 保存临时结果

	for i := 0; i < len(nums)+1; i++ {
		Backtrack(0, nums, list, &result)
	}
	return result
}

func Backtrack(pos int, nums []int, list []int, result *[][]int) {

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		Backtrack(i+1, nums, list, result)
		list = list[0 : len(list)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
