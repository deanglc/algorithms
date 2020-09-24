package main

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

/*		1
	2		3
4	   5  6    7

*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)

}
