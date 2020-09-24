package main

import (
	"fmt"

	"github.com/deanglc/algorithms/1-test/adata"
)

/* 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。


 */
type TreeNode = adata.TreeNode

func main() {
	a := adata.NewData()
	fmt.Println(averageOfLevels1(a))
}

// 思路: 遍历每层 累加后/该层个数 append到result
func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		var sum float64
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			sum += float64(level.Val)
			// 添加下一层 循环往复 直到叶子节点层
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, sum/float64(l))
	}
	return result
}

// type TreeNode struct {
// 	Val         int
// 	Left, Right *TreeNode
// }
type data struct {
	sum, count int
}

/*       1
   21		  22
31    32  33       34
0 1
1 2
2 4
*/
func averageOfLevels1(root *TreeNode) []float64 {
	var levelData []data
	// level=深度 从0开始
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)

	}
	dfs(root, 0)

	averages := make([]float64, len(levelData))
	for i, d := range levelData {
		averages[i] = float64(d.sum) / float64(d.count)
	}
	return averages
}
