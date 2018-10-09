package main

import (
	"fmt"
	"unsafe"
)

// 缓存淘汰算法 最近最少使用策略（least recently used）单链表实现
type ListNode struct {
	Val  int
	Next *ListNode
}

func genListNode() *ListNode {
	n1 := &ListNode{Val: 8}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 7}
	n4 := &ListNode{Val: 4}
	// n5 := &ListNode{Val: 5, Next: nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	// n4.Next = n5

	return n1
}

func main() {
	head := genListNode()
	size := unsafe.Sizeof(head.Next)
	fmt.Printf("size:%v\n", size)
}

// 1 首先遍历链表，发现命中val则返回
// 2 未命中
// 		1 如果超过缓存大小，删除最后一个节点，把当前值加到头结点
//		2 如果没超过缓存，直接加到头结点
