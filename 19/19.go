package main

import "fmt"

func main() {
	fmt.Print()
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nNext := head
	for i := 0; i < n-1; i++ {
		if nNext.Next != nil {
			nNext = nNext.Next
		} else {
			return nil
		}
	}
	fmt.Printf("%v", nNext.Val)
	preHead := &ListNode{Val: 0, Next: head}
	preRemovedNode := preHead
	for nNext.Next != nil {
		nNext = nNext.Next
		preRemovedNode = preRemovedNode.Next
	}
	removedNode := preRemovedNode.Next
	preRemovedNode.Next = removedNode.Next
	removedNode.Next = nil
	return preHead.Next
}
