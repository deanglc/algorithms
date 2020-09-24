package main

//559N叉树的最大深度. N叉树的最大深度 https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
//给定一个 N 叉树，找到其最大深度。最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

type Node struct {
	Val      int
	Children []*Node
}

//测试用例 [1,null,3,2,4,null,5,6]
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var queue = []*Node{root}
	var level int
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			for _, n := range queue[0].Children {
				queue = append(queue, n)
			}
			queue = queue[1:]
		}
		level++
	}
	return level
}
