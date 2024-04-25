package main

import (
	"leetcodego/easy"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{name: "case 1", nums: []int{11, 7, 15, 2}, target: 9},
		{name: "case 2", nums: []int{3, 2, 4}, target: 6},
		{name: "case 3", nums: []int{3, 3}, target: 6},
		{name: "case 4", nums: []int{3, 2, 3}, target: 6},
		{name: "case 5", nums: []int{-1, -2, -3, -4, -5}, target: -8},
		{name: "case 6", nums: []int{0, 4, 3, 0}, target: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.TwoSum(tt.nums, tt.target)

			t.Log(result)
			if len(result) != 2 {
				t.Fatalf("TwoSum failed in case %s by length", tt.name)
			}

			if tt.nums[result[0]]+tt.nums[result[1]] != tt.target || result[1] == result[0] {
				t.Fatalf("TwoSum failed in case %s by sum", tt.name)
			}
		})
	}
}
