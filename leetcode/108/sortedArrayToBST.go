package main

import "github.com/deanglc/algorithms/1-test/adata"

type TreeNode = adata.TreeNode

func main() {

}

/*
		[left root right]
left: [0 mid mid-1]
right: [mid+1 mid ]
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
