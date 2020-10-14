package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(a, target))
}

// func fourSum(nums []int, target int) (result [][]int) {
// 	if len(nums) < 4 {
// 		return nil
// 	}
// 	sort.Ints(nums)
//
// 	return backTrack(nums, target)
//
// }

// func backTrack(nums []int, target int) (result [][]int) {
// 	// l := make([]int, 0, 4)
// 	l := make([]int, 0)
// 	for i, v := range nums {
// 		if len(l) > 4 {
// 			break
// 		}
// 		if target == 0 && len(l) == 4 {
// 			result = append(result, append([]int(nil), l...))
// 			return
// 		}
// 		backTrack(nums[i+1:], target)
// 		if target-v != 0 {
// 			l = append(l, v)
// 			backTrack(nums, target-v)
// 			l = l[:len(l)-1]
// 		}
// 	}
// 	backTrack(nums, 0)
// 	return result
// }
func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
