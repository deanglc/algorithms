package _35

/* 235 注意审题！！！
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

*/
import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := root
	for {
		if q.Val > ancestor.Val && p.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else if q.Val < ancestor.Val && p.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else {
			return ancestor
		}
	}
}
