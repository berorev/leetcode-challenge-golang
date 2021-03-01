package lswrc

import (
	"fmt"
)

// Example struct
type Example struct {
	s        string
	expected int
}

// Solve ...
func Solve() {
	// data
	examples := []Example{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "dvdf",
			expected: 3,
		},
		{
			s:        "x",
			expected: 1,
		},
		{
			s:        "abc",
			expected: 3,
		},
	}

	// solve
	fmt.Println("-- Longest Substring Without Repeating Characters")
	for i, example := range examples {
		fmt.Printf("Example #%d : ", i+1)

		result := lengthOfLongestSubstring(example.s)

		if example.expected == result {
			fmt.Println("success")
		} else {
			fmt.Println("failed. expected", example.expected, "but actual", result)
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	cur := -1

	for i, v := range s {
		c := byte(v)
		if j, exists := m[c]; exists {
			if cur < j {
				cur = j
			}
		}
		len := i - cur
		if max < len {
			max = len
		}
		m[c] = i
	}

	return max
}
