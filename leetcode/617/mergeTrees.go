package main

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

// 层层遍历 相同位置 相加(null=0)
// func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
// 	if t1 == nil {
// 		return t2
// 	}
// 	if t2 == nil {
// 		return t1
// 	}
// 	newTree := &TreeNode{Val: t1.Val + t2.Val}
//
// 	queue := []*TreeNode{newTree}
// 	queue1 := []*TreeNode{t1}
// 	queue2 := []*TreeNode{t2}
// 	// l1 := make([]int, 0)
// 	// l2 := make([]int, 0)
// 	for len(queue1) > 0 && len(queue2) > 0 {
//
// 	}
// }
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTrees3(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	queue := []*TreeNode{t1, t2}
	for len(queue) > 0 {
		n1 := queue[0]
		n2 := queue[1]
		n1.Val = n1.Val + n2.Val
		if n1.Left != nil && n2.Left != nil {
			queue = append(queue, n1.Left, n2.Left)
		} else if n1.Left == nil && n2.Left != nil {
			n1.Left = n2.Left
		}
		if n1.Right != nil && n2.Right != nil {
			queue = append(queue, n1.Right, n2.Right)
		} else if n1.Right == nil && n2.Right != nil {
			n1.Right = n2.Right
		}
		queue = queue[2:]
	}
	return t1
}
