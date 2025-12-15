package problems

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Source: https://leetcode.com/problems/add-two-numbers/description/

Note: This implementation follows a specific logic due to a misunderstanding of the original problem:
1. Reverse both input linked lists.
2. Sum the numbers.
3. Return the reversed sum as the result.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	a := reversedListNumber(l1)
	b := reversedListNumber(l2)

	sum := a + b
	sumBuffer := make([]int, 0, 100)

	for {
		if sum <= 0 {
			break
		}

		dev := sum % 10
		sumBuffer = append(sumBuffer, dev)
		sum /= 10
	}

	sumBufferLen := len(sumBuffer)

	root := &ListNode{}
	current := root
	for i := 0; i < sumBufferLen; i++ {
		current.Val = sumBuffer[i]

		if i == sumBufferLen-1 {
			break
		}

		n := &ListNode{}
		current.Next = n
		current = n
	}

	return root
}

func reversedListNumber(l *ListNode) int {
	var root *ListNode = l
	var buffer = make([]int, 0, 100)

	point := root
	for {
		buffer = append(buffer, point.Val)
		nextPoint := point.Next
		if nextPoint == nil {
			break
		}

		point = nextPoint
	}

	var num int
	for i := len(buffer) - 1; i >= 0; i-- {
		num = (num * 10) + buffer[i]
	}

	return num
}
