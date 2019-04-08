package main

func main() {
	listL := createNodeFromArray([]int{1, 2, 3, 4, 5})
	reverseKGroup(listL, 2)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

func getLength(list *ListNode) int {
	length := 0
	for tempNode := list; tempNode != nil; tempNode = tempNode.Next {
		length++
	}
	return length
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	listLength := getLength(head)
	reverseHead := new(ListNode)
	reverseHead.Next = head

	pre := reverseHead
	tail, next := head, head

	for length := listLength; length >= k; length = length - k {
		for i := 0; i < k; i++ {
			tempNode := next.Next
			next.Next = pre.Next
			pre.Next = next
			next = tempNode
			tail.Next = next
		}
		pre = tail
		tail = next
	}

	return reverseHead.Next
}
