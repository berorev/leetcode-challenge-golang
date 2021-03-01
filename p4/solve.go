package p4

import (
	"fmt"
)

// Example struct
type Example struct {
	nums1    []int
	nums2    []int
	expected float64
}

// Solve ...
func Solve() {
	// data
	examples := []Example{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.50000,
		},
		{
			nums1:    []int{0, 0},
			nums2:    []int{0, 0},
			expected: 0.00000,
		},
		{
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.00000,
		},
		{
			nums1:    []int{2},
			nums2:    []int{},
			expected: 2.00000,
		},
	}

	// solve
	fmt.Println("-- Median of Two Sorted Arrays")
	for i, example := range examples {
		fmt.Printf("Example #%d : ", i+1)

		result := medianOfTwoSortedArrays(example.nums1, example.nums2)

		if example.expected == result {
			fmt.Println("success")
		} else {
			fmt.Println("failed. expected", example.expected, "but actual", result)
		}
	}
}

func medianOfTwoSortedArrays(nums1, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	mergedLength := len1 + len2

	if mergedLength%2 == 1 {
		targetIndex := mergedLength / 2
		j1 := 0
		j2 := 0
		v := 0
		for i := 0; i <= targetIndex; i++ {
			if j2 >= len2 {
				v = nums1[j1]
				j1++
			} else if j1 >= len1 {
				v = nums2[j2]
				j2++
			} else if nums1[j1] < nums2[j2] {
				v = nums1[j1]
				j1++
			} else if nums1[j1] >= nums2[j2] {
				v = nums2[j2]
				j2++
			}
		}
		return float64(v)
	}

	targetIndex := mergedLength / 2
	j1 := 0
	j2 := 0
	v1 := 0
	v2 := 0
	for i := 0; i <= targetIndex; i++ {
		if j2 >= len2 {
			v1 = nums1[j1]
			j1++
		} else if j1 >= len1 {
			v1 = nums2[j2]
			j2++
		} else if nums1[j1] < nums2[j2] {
			v1 = nums1[j1]
			j1++
		} else if nums1[j1] >= nums2[j2] {
			v1 = nums2[j2]
			j2++
		}
		if i == targetIndex-1 {
			v2 = v1
		}
	}

	return float64(v1+v2) / 2
}
