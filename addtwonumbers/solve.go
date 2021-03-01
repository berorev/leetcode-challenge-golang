package addtwonumbers

import (
	"berorev/leetcode/utils/arrays"
	"fmt"
)

// ListNode defines singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Example ...
type Example struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

// Append ...
func (n *ListNode) Append(v int) *ListNode {
	next := &ListNode{v, nil}
	if n != nil {
		n.Next = next
	}
	return next
}

func toList(arr []int) *ListNode {
	firstNode := &ListNode{arr[0], nil}
	n := firstNode
	for _, v := range arr[1:] {
		n = n.Append(v)
	}
	return firstNode
}
func toArray(n *ListNode) []int {
	arr := []int{}
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}
	return arr
}

// Solve ...
func Solve() {
	// data
	examples := []Example{
		{
			l1:       toList([]int{2, 4, 3}),
			l2:       toList([]int{5, 6, 4}),
			expected: toList([]int{7, 0, 8}),
		},
		{
			l1:       toList([]int{0}),
			l2:       toList([]int{0}),
			expected: toList([]int{0}),
		},
		{
			l1:       toList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       toList([]int{9, 9, 9, 9}),
			expected: toList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	// solve
	fmt.Println("-- Add Two Numbers")
	for i, example := range examples {
		fmt.Printf("Example #%d : ", i+1)

		result := addTwoNumbers2(example.l1, example.l2)

		expectedArr := toArray(example.expected)
		resultArr := toArray(result)
		if correct := arrays.Equal(expectedArr, resultArr); correct {
			fmt.Println("success")
		} else {
			fmt.Println("failed. expected", expectedArr, "but actual", resultArr)
		}
	}
}

/**
 * runtime: 16ms(22.9%), memory: 5mb(21.23%)
 */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var node, firstNode *ListNode
	rounding := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + rounding

		// node = node.Append(sum % 10)
		next := &ListNode{sum % 10, nil}
		if node != nil {
			node.Next = next
		}
		node = next

		if firstNode == nil {
			firstNode = node
		}
		rounding = sum / 10
	}

	if rounding > 0 {
		// node = node.Append(rounding)
		node.Next = &ListNode{rounding, nil}
	}
	return firstNode
}

/**
 * runtime: 4ms(98.73%), memory: 5mb(21.23%)
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{0, nil}
	node := &dummyNode
	rounding := 0
	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += rounding

		// node = node.Append(sum % 10)
		node.Next = &ListNode{sum % 10, nil}
		node = node.Next

		rounding = sum / 10
	}

	if rounding > 0 {
		// node = node.Append(rounding)
		node.Next = &ListNode{rounding, nil}
	}
	return dummyNode.Next
}
