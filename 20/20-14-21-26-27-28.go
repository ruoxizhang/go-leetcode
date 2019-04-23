package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", isValid("()[]{}"))
}

func isValid(s string) bool {
	parenthesisStack := []rune{}

	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			parenthesisStack = append(parenthesisStack, value)
		} else if value == ')' {
			if len(parenthesisStack) == 0 {
				return false
			}
			pop := parenthesisStack[len(parenthesisStack)-1]
			if pop != '(' {
				return false
			}
			parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]
		} else if value == ']' {
			if len(parenthesisStack) == 0 {
				return false
			}
			pop := parenthesisStack[len(parenthesisStack)-1]
			if pop != '[' {
				return false
			}
			parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]
		} else if value == '}' {
			if len(parenthesisStack) == 0 {
				return false
			}
			pop := parenthesisStack[len(parenthesisStack)-1]
			if pop != '{' {
				return false
			}
			parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]
		} else {
			return false
		}
	}

	if len(parenthesisStack) != 0 {
		return false
	}

	return true
}

func longestCommonPrefix(strs []string) string {
	lenStrs := len(strs)
	if lenStrs == 0 {
		return ""
	}

	if len(strs[0]) == 0 {
		return ""
	}

	index := 0

	for true {
		for _, value := range strs {
			if len(value) <= index || value[index] != strs[0][index] {
				return strs[0][0:index]
			}
		}
		index++
	}

	return strs[0]

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	newNode := &head
	l1Node := l1
	l2Node := l2
	for l1Node != nil && l2Node != nil {
		if l1Node.Val < l2Node.Val {
			newNode.Next = l1Node
			newNode = newNode.Next
			l1Node = l1Node.Next
		} else {
			newNode.Next = l2Node
			newNode = newNode.Next
			l2Node = l2Node.Next
		}
	}
	if l1Node != nil {
		newNode.Next = l1Node
	}

	if l2Node != nil {
		newNode.Next = l2Node
	}
	return head.Next
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	index := 0
	for index+1 < len(nums) {
		if nums[index] == nums[index+1] {
			nums = append(nums[:index+1], nums[index+2:]...)
		} else {
			index++
		}
	}
	return len(nums)
}

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	result := searchRange(nums, val)
	if result[0] != -1 {
		nums = append(nums[:result[0]], nums[result[1]+1:]...)
	}

	return len(nums)
}

func searchRange(nums []int, target int) []int {
	lenNums := len(nums)

	if lenNums == 0 {
		return []int{-1, -1}
	}

	if lenNums == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	if target > nums[lenNums-1] {
		return []int{-1, -1}
	}
	if target < nums[0] {
		return []int{-1, -1}
	}

	resultStart := findTheFirst(nums, target)
	if resultStart != -1 {
		resultEnd := findTheLast(nums, target)
		return []int{resultStart, resultEnd}
	}
	return []int{-1, -1}

}

func findTheFirst(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[start] == target && (start == 0 || nums[start-1] != target) {
			return start
		}

		middle := (end-start+1)/2 + start
		if target > nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	if nums[start] == target && (start == 0 || nums[start-1] != target) {
		return start
	}
	if nums[end] == target && (end == 0 || nums[end-1] != target) {
		return end
	}

	return -1
}

func findTheLast(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
			return end
		}

		middle := (end-start+1)/2 + start
		if target < nums[middle] {
			end = middle
		} else {
			start = middle
		}
	}

	if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
		return end
	}
	if nums[start] == target && (start == lenNums-1 || nums[start+1] != target) {
		return start
	}

	return -1
}
