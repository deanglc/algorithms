package _01

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

// v1 94.3 20.44
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	origin := root
	for root != nil {
		if root.Val < val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				break
			} else {
				root = root.Right
			}
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{Val: val}
				break
			} else {
				root = root.Left
			}
		}
	}
	return origin
}
