package main

import "fmt"

func main() {
	fmt.Println(swapPairs(createNodeFromArray([]int{1, 2, 3, 4, 5, 6})))
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{Val: 0, Next: nil}
	newHead.Next = head
	tempTail := head
	tempHead := newHead
	for tempTail != nil && tempTail.Next != nil {
		thisNode := tempTail
		tempTail = thisNode.Next.Next
		tempHead.Next = thisNode.Next
		tempHead.Next.Next = thisNode
		thisNode.Next = tempTail
		tempHead = thisNode
		thisNode = tempTail

	}
	if newHead == nil {
		return head
	}
	return newHead.Next
}

func createNodeFromArray(number []int) *ListNode {
	headNode := &ListNode{Val: 0, Next: nil}
	previousNode := headNode
	for _, value := range number {
		previousNode.Next = &ListNode{Val: value, Next: nil}
		previousNode = previousNode.Next
	}
	return headNode.Next
}
