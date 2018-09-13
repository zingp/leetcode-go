package main

import (
	"fmt"
)

// ListNode 链表结构体定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断是否是回文链表
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
// 用两个指针，pre代表当前节点的前一个节点，next代表当前节点的后一个节点
func reverseList(head *ListNode) *ListNode {
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

// 打印链表的每一个元素
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
func selectMidListNode(head *ListNode) *ListNode {
	// 如果只有一个元素
	if head == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 生成测试链表
func genListNode() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	// mid := selectMidListNode(n1)
	// rev := reverseList(n1)
	// rangeList(rev)
	fmt.Println(isPalindrome(n1))

}

func main() {
	genListNode()
}

// 在O(1)时间复杂度内删除单链表指定节点
func deleteNode(head *ListNode, node *ListNode)(*ListNode){
	// 如果输入参数有空，直接返回头结点
	if head == nil || node == nil {
		return head
	}
	// 如果删除的是头结点
	if head == node {
		return head.Next
	}

	//如果链表有两个及以上节点，且删除的是尾节点
	if node.Next == nil {
		// 找待删除元素的前驱节点  
		tmp := head
		for  tmp.Next != node {
			tmp = tmp.Next
		}
		// 删除节点
		tmp.Next = nil
	}else {
		// 如果链表有两个及以上节点，且删除的是某个中间结点
		// 将待删除的节点的下一个节点的值 付给待删除的节点
		node.Val = node.Next.Val
		// 将待删除节点的下节点指向下下个节点，即删除待删除节点的下一个节点
		node.Next = node.Next.Next
	}
	// 返回删除节点后的链表头结点
	return head
}

// 反转部分单链表
func reverseBetween(head *ListNode, m int, n int)(* ListNode){
	if head == nil || head.Next == nil {
		return head
	}
	i := 1
	var reverseNewHead *ListNode   // 反转部分链表反转后的头结点，反转过程中这个节点不断变化
	var reverseNewTail *ListNode   // 第m个节点，即反转部分链表反转后的尾结点
	var reversePreNode *ListNode   // 第m-1 个节点，即反转部分链表反转前其头结点的前一个结点
	var reverseNextNode *ListNode  // 反转时，用于挂住下一个节点的节点，当前head的下一个节点
	oldHead := head
	for head != nil {
		if i > n {
			break
		}

		if i == m-1 {
			reversePreNode = head
		}
		if i >= m && i <= n {
			if i == m {
				reverseNewTail = head
			}
			reverseNextNode = head.Next
			head.Next = reverseNewHead
			reverseNewHead = head
			head = reverseNextNode
		}else{
			head = head.Next
		}

		i++
	}
	// 反转部分反转之后，整体反转一次，连接断开的地方
	// 反转部分反转之后，第m个节点的下一个节点应该是第n+1个节点
	reverseNewTail.Next = reverseNextNode
	// 第m-1个节点不是空节点，则他的下一个节点应该反转部分的头结点
	if reversePreNode != nil {
		reversePreNode.Next = reverseNewHead
		return oldHead
	} else {
		// 如果第m-1个节点是空节点说明m=1，从头开始反转；新的头结点就是reverseNewHead
		return reverseNewHead
	}
}

// 旋转单链表
func rotateRight(head *ListNode, k int)(*ListNode) {
	n := 0 
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	// 首尾连接
	cur.Next = head
	m := n - k%n
	for i:=0; i<m; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}