package _113

import (
	"github.com/deanglc/algorithms/1-test/adata"
)

type TreeNode = adata.TreeNode

// func main() {
//
// 	b := &TreeNode{}
// 	c := &TreeNode{Val: -3}
// 	a := &TreeNode{Val: -2, Left: b, Right: c}
// 	fmt.Println(pathSum1(a, -5))
// }

func pathSum2(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	var dfs func(int, *TreeNode, []int)
	dfs = func(remain int, root *TreeNode, l []int) {
		if root == nil {
			return
		}
		remain -= root.Val
		if isLeafNode(root) && remain == 0 {
			l = append(l, root.Val)
			result = append(result, append([]int(nil), l...))
			return
		} else if !isLeafNode(root) {
			l = append(l, root.Val)
			dfs(remain, root.Left, l)
			dfs(remain, root.Right, l)
		} else {
			return
		}
	}
	dfs(sum, root, []int(nil))
	return result
}

// v1 93.79 && 53.87 4ms 4.6mb
func pathSum1(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	dfs(root, []int(nil), sum, &result)
	return result
}
func dfs(root *TreeNode, tempL []int, remain int, result *[][]int) {
	if root == nil {
		return
	}
	remain = remain - root.Val
	if remain == 0 && isLeafNode(root) {
		tempL = append(tempL, root.Val)
		*result = append(*result, append([]int(nil), tempL...))
	} else if !isLeafNode(root) {
		tempL = append(tempL, root.Val)
		dfs(root.Left, tempL, remain, result)
		dfs(root.Right, tempL, remain, result)
	} else {
		return
	}
}
func isLeafNode(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}
