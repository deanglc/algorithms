package somanysort_test

import (
	"fmt"
	"testing"

	"github.com/deanglc/algorithms/someCase/somanysort"
)

func TestQuick1Sort(t *testing.T) {
	nums := []int{1, 3, 4, 2}
	somanysort.Quick1Sort(nums, 0, 3)

	// exp := []int{1,2,3,4}

	fmt.Println(nums)
}
func TestQuick2Sort(t *testing.T) {
	nums := []int{1, 3, 4, 2}
	somanysort.Quick2Sort(nums)

	// exp := []int{1,2,3,4}

	fmt.Println(nums)
}
