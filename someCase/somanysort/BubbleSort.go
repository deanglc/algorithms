package somanysort

//
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

}
