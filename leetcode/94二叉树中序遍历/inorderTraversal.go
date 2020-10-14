package main

import (
	"fmt"

	"github.com/deanglc/algorithms/1-test/adata"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 中序非递归  31 21 32 1 33 22 34

// 中序遍历
// 左树左到底 添加由下到上 由左到右
// root
// 右树  添加由上到下

// 迭代法
func inorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		// 左到底
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 操作末端的节点
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		root = curr.Right
	}
	return result
}

// 递归
func inorderRecursiveTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	result = append(inorderRecursiveTraversal(root.Left), result...)
	result = append(result, root.Val)
	result = append(result, inorderRecursiveTraversal(root.Right)...)
	return result
}

// 莫里斯遍历法
func MorrisTraverMid(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}

	// 游标节点初始化为根节点
	cur := root

	// 定义前驱节点
	var pre *TreeNode

	// 当没有遍历到最后情况
	for cur != nil {

		// 当游标节点没有左孩子
		if cur.Left == nil {
			// 加游标节点值加入结果集(visit)
			result = append(result, cur.Val)

			// 因为没有左孩子，进入右孩子
			cur = cur.Right
			continue
		}

		// 当游标有左孩子
		// 1.找到左子树最右节点作为游标节点前驱

		// 得到左子树
		pre = cur.Left

		// 找到左子树最右叶子节点
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		// 最右叶子节点
		if pre.Right == nil {
			// 最右叶子节点与cur连接
			pre.Right = cur

			// 进入左子树
			cur = cur.Left
			continue
		}

		// 最右节点与cur相等（成环情况）
		// visit cur
		result = append(result, cur.Val)

		// 破环
		pre.Right = nil

		// 进入右子树
		cur = cur.Right
	}

	return result
}
func inorderTraversalMLS(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}

type TreeNode = adata.TreeNode

func main() {

	a := adata.NewData()
	fmt.Println(newFunc(a))
	fmt.Println(MorrisTraverMid(a))
}

func newFunc(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := make([]interface{}, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if val, ok := i.(*TreeNode); ok {
			if val.Right != nil {
				stack = append(stack, val.Right)
			}
			stack = append(stack, val.Val)
			if val.Left != nil {
				stack = append(stack, val.Left)
			}
		} else if val, ok := i.(int); ok {
			result = append(result, val)
		}
		// else if 可改为
		// } else {
		// 	result = append(result, i.(int))
		// }
	}
	return result
}
