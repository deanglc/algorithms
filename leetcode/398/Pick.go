package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
	r    *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
		r:    rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (this *Solution) Pick(target int) int {
	var ans, idx = 0, 1

	for i := 0; i < len(this.nums); i++ {
		if target == this.nums[i] {
			if this.r.Intn(idx) == 0 {
				ans = i
			}
			idx++
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
