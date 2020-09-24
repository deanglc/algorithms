package main

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

func minDepth(root *TreeNode) (result int) {
	// 通过上一层的长度确定下一层的元素
	if root == nil {
		return 0
	}
	result = 1
	queue := make([]*TreeNode, 0) // 存放node
	queue = append(queue, root)
	for len(queue) > 0 {
		// 为什么要取length？
		// 遍历当前层
		l := len(queue)
		for i := 0; i < l; i++ {
			// 凡走过，皆queue.remove
			level := queue[0]
			queue = queue[1:]

			if level.Left == nil && level.Right == nil {
				return result
			}
			// 添加下一层 循环往复 直到叶子节点层
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result++
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
