package main

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
