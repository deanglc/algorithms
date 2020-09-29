package _45二叉树后序遍历

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

// 100% 13.57%
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var l []int

	l = append(l, postorderTraversal(root.Right)...)
	l = append(l, postorderTraversal(root.Left)...)
	l = append(l, root.Val)

	return l
}

// 100% 60.47%
func postorderTraversal2(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

//         1
//     21      22
//   31    32 33   34
// 41  42 43
//    [41,42,31,43,32,21,33,34,22,1]

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	var lastVisit *TreeNode // 记录上一个已弹出node
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		curr := stack[len(stack)-1]
		if curr.Right == nil || curr.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, curr.Val)
			lastVisit = curr
		} else {
			root = curr.Right
		}
	}
	return result
}

// 官方解法2
func postorderTraversal3(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

/* 莫里斯遍历
思路与算法

有一种巧妙的方法可以在线性时间内，只占用常数空间来实现后序遍历。这种方法由 J. H. Morris 在 1979 年的论文「Traversing Binary Trees Simply and Cheaply」中首次提出，因此被称为 Morris 遍历。
Morris 遍历的核心思想是利用树的大量空闲指针，实现空间开销的极限缩减。其后序遍历规则总结如下：

1新建临时节点，令该节点为 root；

2如果当前节点的左子节点为空，则遍历当前节点的右子节点；

3如果当前节点的左子节点不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点；
3如果前驱节点的右子节点为空，将前驱节点的右子节点设置为当前节点，当前节点更新为当前节点的左子节点。
3如果前驱节点的右子节点为当前节点，将它的右子节点重新设为空。倒序输出从当前节点的左子节点到该前驱节点这条路径上的所有节点。当前节点更新为当前节点的右子节点。

4重复步骤 2 和步骤 3，直到遍历结束。

这样我们利用 Morris 遍历的方法，后序遍历该二叉搜索树，即可实现线性时间与常数空间的遍历

复杂度分析：
时间复杂度：O(n)O(n)，其中 nn 是二叉树的节点数。没有左子树的节点只被访问一次，有左子树的节点被访问两次。
空间复杂度：O(1)O(1)。只操作已经存在的指针（树的空闲指针），因此只需要常数的额外空间
*/
func postorderTraversal4(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		path := []int{}
		for ; node != nil; node = node.Right {
			path = append(path, node.Val)
		}
		for i := len(path) - 1; i >= 0; i-- {
			res = append(res, path[i])
		}
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
