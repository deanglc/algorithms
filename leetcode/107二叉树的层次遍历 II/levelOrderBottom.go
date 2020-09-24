package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result = make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)

		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			list = append(list, curr.Val)
			// 添加下一层的数据
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, list)
	}
	/* 翻转方法1*/
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	// 翻转方法2
	// for i := 0; i < len(levelOrder) / 2; i++ {
	// 	levelOrder[i], levelOrder[len(levelOrder) - 1 - i] = levelOrder[len(levelOrder) - 1 - i], levelOrder[i]
	// }

	return result
}

// 可以不翻转
func levelOrderBottom2(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result = make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			list = append(list, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		// 将最新结果放在队列首即可   append是放在队尾
		result = append([][]int{list}, result...)
	}
	return result
}
