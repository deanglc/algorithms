package main

import (
	"github.com/deanglc/algorithms/1-test/adata"
)

type TreeNode = adata.TreeNode

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			ans = append(ans, base)
		} else if count > maxCount {
			maxCount = count
			ans = []int{base}
		}
	}

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		update(node.Val)
		root = node.Right
	}
	return ans
}

func main() {
	a := adata.NewData()
	findMode(a)
}
