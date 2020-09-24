package binary_search_tree

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 插入
func (n *TreeNode) Insert(newTreeNode *TreeNode) {
	if n.Val < newTreeNode.Val {
		n.Right = newTreeNode
	} else {
		n.Right.Insert(newTreeNode)
	}

	if n.Val > newTreeNode.Val {
		if n.Left == nil {
			n.Left = newTreeNode
		} else {
			n.Left.Insert(newTreeNode)
		}
	}
}

// 获取最小值
func (n *TreeNode) FindMin() (int, error) {
	if n == nil {
		return 0, errors.New("nil 没有最小")
	}
	for n.Left != nil {
		n = n.Left
	}
	return n.Val, nil
}

// 获取最大值
func (n *TreeNode) FindMax() (int, error) {
	if n == nil {
		return 0, errors.New("nil 没有最大")
	}

	for n.Right != nil {
		n = n.Right
	}
	return n.Val, nil
}

// 删除指定元素
func (n *TreeNode) Delete(element int) (*TreeNode, bool) {
	deleted := false

	if n == nil {
		return n, false
	}

	if n.Val < element {
		n.Right, deleted = n.Right.Delete(element)
	} else if n.Val > element {
		n.Left, deleted = n.Left.Delete(element)
	} else if n.Left != nil && n.Right != nil {
		candidate, _ := n.Right.FindMin()
		n.Val = candidate
		n.Right, deleted = n.Right.Delete(candidate)
		deleted = true
	} else {
		if n.Left == nil {
			n = n.Right
		} else {
			n = n.Left
		}
		deleted = true
	}
	return n, deleted
}

//
func (n *TreeNode) Walk(f func(int)) {
	if n == nil {
		return
	}

	n.Left.Walk(f)
	f(n.Val)
	n.Right.Walk(f)

}

type BinarySearchTree struct {
	root    *TreeNode
	nodeNum int
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) TreeNodes() int {
	return b.nodeNum
}

func (b *BinarySearchTree) Insert(element int) {
	node := TreeNode{Val: element}

	if b.root == nil {
		b.root = &node
	} else {
		b.root.Insert(&node)
	}
	b.nodeNum++
}

func (b *BinarySearchTree) Walk(f func(int)) {
	b.root.Walk(f)
}

func (b *BinarySearchTree) Contains(element int) bool {
	n := b.root
	for n != nil {
		if n.Val == element {
			return true
		}

		if n.Val < element {
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return false
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.nodeNum == 0
}

func (b *BinarySearchTree) Find(element int) *TreeNode {
	n := b.root
	for n != nil {
		if n.Val == element {
			return n
		}
		if n.Val < element {
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return nil
}

func (b *BinarySearchTree) FindMin() (int, error) {
	return b.root.FindMin()
}

func (b *BinarySearchTree) FindMax() (int, error) {
	return b.root.FindMax()
}

func (b *BinarySearchTree) Delete(element int) {
	n, deleted := b.root.Delete(element)
	b.root = n
	if deleted {
		b.nodeNum--
	}

}

/*
前-中-后序遍历 由root节点被遍历时机决定
左子树总是优先于右子树
    	1
  	21  	22
  31  32  33  34
前序非递归  1 21 31 32 22 33 34
中序非递归  31 21 32 1 33 22 34
后序非递归  31 32 21 33 34 22 1
    		1
       21  	       22
   31      32    33   34
 41  42 43   44
前序非递归 preorderTraversal [1 21 31 41 42 32 43 44 22 33 34]
中序非递归 inorderTraversal  [41 31 42 21 43 32 44 1 33 22 34]
后序非递归 postorderTraversal[41 42 31 43 44 32 21 33 34 22 1]
分治法     divideAndConquer  [1 21 31 41 42 32 43 44 22 33 34]
BFS层次遍历     levelOrder   [[1] [21 22] [31 32 33 34] [41 42 43 44]]



*/
// 前序递归
func preRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	preRecursiveTraversal(root.Left)
	preRecursiveTraversal(root.Right)
}

// 前序非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val) // 保存结果
			stack = append(stack, root)
			root = root.Left
		}
		// 左到底了，pop最底下那个node 用它的right
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序非递归
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		// 左到底
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop [-1] && result.append && (try right || up)
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后序非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode // 记录上一个已弹出node
	for root != nil || len(stack) != 0 {
		// 左到底
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 查看左底node
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop [-1]
			result = append(result, node.Val)
			// 标记当前这个node已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// dfs 深度遍历-从上到下-左+右
func deepUp2Down(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// dfs 深度遍历- divideAndConquer(分治法) {root+左+右}
// 特点-递归后合并合并！！！
func divideAndConquer(root *TreeNode) (result []int) {
	result = make([]int, 0)
	// nil & leaf
	if root == nil {
		return result
	}
	// 分治(Divide)
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果(Conquer)
	result = append(result, root.Val) // root
	result = append(result, left...)  // 左
	result = append(result, right...) // 右
	return result
}

// BFS层次遍历
func levelOrder(root *TreeNode) (result [][]int) {
	// 通过上一层的长度确定下一层的元素
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0) // 存放node
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0) // 存放node.val
		// 为什么要取length？
		// 遍历当前层
		l := len(queue)
		for i := 0; i < l; i++ {
			// 凡走过，皆queue.remove
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			// 添加下一层 循环往复 直到叶子节点层
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
