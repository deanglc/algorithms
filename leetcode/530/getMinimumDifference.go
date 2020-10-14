package _30

import (
	"math"

	"github.com/deanglc/algorithms/1-test/adata"
)

/**
 * 本题与783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = adata.TreeNode

/*
			5
		2		6
	  1   3	       9


*/
func getMinimumDifference(root *TreeNode) (min int) {
	min, pre := math.MaxInt32, -1
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != -1 && curr.Val-pre < min {
			min = curr.Val - pre
		}
		pre = curr.Val
		root = curr.Right
	}
	return min
}

//
func getMinimumDifferenceA(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

//
func getMinimumDifferenceB(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// min是要返回的差的绝对值的最小值
	// preVal记录相邻的上一个节点的值
	// cur是当前节点，用于Morris遍历，初始化为root
	min, preVal, cur := math.MaxInt64, -1, root
	// updateMin用于更新变量min
	updateMin := func(node *TreeNode, num int) {
		if preVal != -1 && node.Val-preVal < min {
			min = node.Val - preVal
		}
		// 无论min是否被赋予新值，均用preVal记录当前节点的值
		preVal = node.Val
	}
	// Morris遍历
	// cur是当前节点
	for cur != nil {
		// 如果cur无左子节点，则更新min，cur指向cur的右子节点
		if cur.Left == nil {
			updateMin(cur, preVal)
			cur = cur.Right
			continue
		}
		// 如果cur有左子节点
		// 先用preNode记录cur的左子节点
		preNode := cur.Left
		// 找到cur左子树上最右的节点，赋值给preNode
		for preNode.Right != nil && preNode.Right != cur {
			preNode = preNode.Right
		}
		// 如果cur左子树上最右的节点（preNode）的right指针指向空，则让其指向cur，cur指向cur的左子节点
		if preNode.Right == nil {
			preNode.Right = cur
			cur = cur.Left
		} else {
			// 如果cur左子树上最右的节点（preNode）的right指针指向cur，让其指向空，更新min，cur指向cur的右子节点
			preNode.Right = nil
			updateMin(cur, preVal)
			cur = cur.Right
		}
	}
	return min
}
