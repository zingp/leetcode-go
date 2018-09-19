package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
		return head
	}
	var leftDummy = &ListNode{}
	var rightDummy = &ListNode{}
	left := leftDummy
	right := rightDummy

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = head
		} else {
			right.Next = head
			right = head
		}
		head = head.Next
	}
	right.Next = nil
	left.Next = rightDummy.Next
	return leftDummy.Next  
}

func main() {
	
}