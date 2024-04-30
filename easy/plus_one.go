package easy

func PlusOne(digits []int) []int {
	mxIndex := len(digits) - 1

	digits[mxIndex]++

	for digits[mxIndex] == 10 && mxIndex > 0 {
		digits[mxIndex] = 0
		mxIndex--
		digits[mxIndex]++
	}

	if digits[mxIndex] == 10 {
		digits[mxIndex] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
