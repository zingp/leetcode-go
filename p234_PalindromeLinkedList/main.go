package main

import (
	"fmt"
)

/*
Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 如果只有一个元素
    if head == nil || head.Next == nil {
		return true
	}
	
	// 用两个指针，慢指针每次移动一位，快指针每次移动两位，找到中点slow
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	// 反转后半段链表，pre就是反转后的头结点
	var pre *ListNode
	var next *ListNode
	for slow != nil {
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	// 对比反转后的链表和未反转的链表中的每一个元素
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}

	return true
}

// 单链表反转
func reverseList(head *ListNode)(*ListNode){
	if head == nil {
		return head
	}

	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func rangeList(head *ListNode) {
	if head == nil {
		fmt.Println(nil)
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}	
}

// 寻找链表中点的节点
func selectMidListNode(head *ListNode) (*ListNode) {
	// 如果只有一个元素
    if head == nil {
		return head
	}
	slow := head
	fast := head
	for (fast!= nil && fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}



func genListNode(){
	n1 := &ListNode{Val:1,}
	// n2 := &ListNode{Val:1,}
	// n3 := &ListNode{Val:2,}
	// n4 := &ListNode{Val:1,}
	// n5 := &ListNode{Val:5,}
	// n1.Next = n2
	// n2.Next = n3
	// n3.Next = n4
	// n4.Next = n5

	// mid := selectMidListNode(n1)
	// rev := reverseList(n1)
	// rangeList(rev)
	fmt.Println(isPalindrome(n1))

}

func main() {
	genListNode()
}