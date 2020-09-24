package main

import (
	"fmt"

	"github.com/deanglc/algorithms/1-test/adata"
)

func main() {
	a := adata.NewData()
	fmt.Println(sumOfLeftLeaves(a))

}

type TreeNode = adata.TreeNode

func sumOfLeftLeaves(root *TreeNode) (r int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return sumOfLeftLeaves(root.Right) + root.Left.Val
	} else {
		return sumOfLeftLeaves(root.Right) + sumOfLeftLeaves(root.Left)
	}
}
func sumOfLeftLeaves2(root *TreeNode) (r int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			r = r + root.Left.Val
		} else {
			dfs(root.Left)
		}
		dfs(root.Right)
	}
	dfs(root)
	return r
}

// result := make([]int, 0)
// dfs(root, &result)
// return result
// }
// func dfs(root *TreeNode, result *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*result = append(*result, root.Val)
// 	dfs(root.Left, result)
// 	dfs(root.Right, result)
// }
