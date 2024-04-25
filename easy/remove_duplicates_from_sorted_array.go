package easy

func removeDuplicates(nums []int) int {
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)

			i = 0
		}

		if len(nums)-1 >= i+1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)

			i = 0
		}
	}

	return len(nums)
}
