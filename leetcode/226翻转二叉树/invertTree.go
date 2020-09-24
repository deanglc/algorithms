package main

// 翻转一棵二叉树。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 以下模版适合于 一层一层的 修改｜遍历｜等等
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
		currNode.Left, currNode.Right = currNode.Right, currNode.Left

	}
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
