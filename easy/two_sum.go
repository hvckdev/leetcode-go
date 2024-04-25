package easy

func TwoSum(nums []int, target int) []int {
	realIndexes := make(map[int][]int)

	for i2, num := range nums {
		realIndexes[num] = append(realIndexes[num], i2)
	}

	for _, number := range nums {
		index, ok := realIndexes[target-number]

		if ok {
			if len(index) > 1 {
				return index
			}

			secondIndex, secondOk := realIndexes[number]

			if secondOk && secondIndex[0] != index[0] {
				return []int{index[0], secondIndex[0]}
			}
		}
	}

	return []int{}
}
