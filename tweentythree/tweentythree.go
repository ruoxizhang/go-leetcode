package main

func main() {
	lists := make([]*ListNode, 0)
	lists = append(lists, createNodeFromArray([]int{1, 4, 5}))
	lists = append(lists, createNodeFromArray([]int{1, 3, 4}))
	lists = append(lists, createNodeFromArray([]int{2, 6}))
	mergeKLists(lists)
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

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	oldLists := lists
	for len(oldLists) > 1 {
		newLists := make([]*ListNode, 0)
		for index := 0; index < len(oldLists); index = index + 2 {
			if index+1 < len(oldLists) {
				newLists = append(newLists, mergeTwoSortedList(oldLists[index], oldLists[index+1]))
			} else {
				newLists = append(newLists, oldLists[index])
			}
		}
		oldLists = newLists
	}
	return oldLists[0]
}

//list 1 is the longest one
func mergeTwoSortedList(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList = &ListNode{Val: 0, Next: nil}
	currentNewNode := newList

	nextNode1 := list1
	nextNode2 := list2
	for nextNode1 != nil && nextNode2 != nil {
		if nextNode1.Val < nextNode2.Val {
			currentNewNode.Next = nextNode1
			nextNode1 = nextNode1.Next
		} else {
			currentNewNode.Next = nextNode2
			nextNode2 = nextNode2.Next
		}
		currentNewNode = currentNewNode.Next
	}

	if nextNode1 != nil {
		currentNewNode.Next = nextNode1
	}

	if nextNode2 != nil {
		currentNewNode.Next = nextNode2
	}

	return newList.Next
}
