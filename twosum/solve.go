package twosum

import (
	"berorev/leetcode/utils/arrays"
	"fmt"
)

// Example struct
type Example struct {
	nums     []int
	target   int
	expected []int
}

// Solve ...
func Solve() {
	// data
	examples := []Example{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	// solve
	for i, example := range examples {
		fmt.Printf("Example #%d : ", i+1)

		result := twoSum(example.nums, example.target)

		if correct := arrays.Equal(example.expected, result); correct {
			fmt.Println("success")
		} else {
			fmt.Println("failed. expected", example.expected, "but actual", result)
		}
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i1, v1 := range nums {
		if i2, exists := m[target-v1]; exists {
			return []int{i2, i1}
		}
		m[v1] = i1
	}
	return []int{}
}
